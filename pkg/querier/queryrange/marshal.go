// Package contains methods to marshal logqmodel types to queryrange Protobuf types.
// Its cousing is util/marshal which converts them to JSON.
package queryrange

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gogo/googleapis/google/rpc"
	"github.com/gogo/status"
	"example.com/acme/kit/httpgrpc"
	"example.com/acme/kit/user"
	"github.com/prometheus/prometheus/promql"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/grpc/codes"

	"example.com/acme/logstore/v3/pkg/loghttp"
	"example.com/acme/logstore/v3/pkg/logproto"
	"example.com/acme/logstore/v3/pkg/logql"
	"example.com/acme/logstore/v3/pkg/logql/sketch"
	"example.com/acme/logstore/v3/pkg/logql/syntax"
	"example.com/acme/logstore/v3/pkg/logqlmodel"
	"example.com/acme/logstore/v3/pkg/querier/plan"
	"example.com/acme/logstore/v3/pkg/querier/queryrange/queryrangebase"
	"example.com/acme/logstore/v3/pkg/util/httpreq"
	"example.com/acme/logstore/v3/pkg/util/querylimits"
	"example.com/acme/logstore/v3/pkg/util/server"
)

const (
	JSONType     = `application/json; charset=utf-8`
	ParquetType  = `application/vnd.apache.parquet`
	ProtobufType = `application/vnd.google.protobuf`
)

// WriteQueryResponseProtobuf marshals the promql.Value to queryrange QueryResonse and then
// writes it to the provided io.Writer.
func WriteQueryResponseProtobuf(params logql.Params, v logqlmodel.Result, w io.Writer) error {
	r, err := ResultToResponse(v, params)
	if err != nil {
		return err
	}

	p, err := QueryResponseWrap(r)
	if err != nil {
		return err
	}

	buf, err := p.Marshal()
	if err != nil {
		return err
	}
	_, err = w.Write(buf)
	return err
}

// ResultToResponse is the reverse of ResponseToResult below.
func ResultToResponse(result logqlmodel.Result, params logql.Params) (queryrangebase.Response, error) {
	switch data := result.Data.(type) {
	case promql.Vector:
		sampleStream, err := queryrangebase.FromValue(data)
		if err != nil {
			return nil, err
		}

		return &LogstorePromResponse{
			Response: &queryrangebase.PrometheusResponse{
				Status: "success",
				Data: queryrangebase.PrometheusData{
					ResultType: loghttp.ResultTypeVector,
					Result:     sampleStream,
				},
				Warnings: result.Warnings,
			},
			Statistics: result.Statistics,
		}, nil
	case promql.Matrix:
		sampleStream, err := queryrangebase.FromValue(data)
		if err != nil {
			return nil, err
		}
		return &LogstorePromResponse{
			Response: &queryrangebase.PrometheusResponse{
				Status: "success",
				Data: queryrangebase.PrometheusData{
					ResultType: loghttp.ResultTypeMatrix,
					Result:     sampleStream,
				},
				Warnings: result.Warnings,
			},
			Statistics: result.Statistics,
		}, nil
	case promql.Scalar:
		sampleStream, err := queryrangebase.FromValue(data)
		if err != nil {
			return nil, err
		}

		return &LogstorePromResponse{
			Response: &queryrangebase.PrometheusResponse{
				Status: "success",
				Data: queryrangebase.PrometheusData{
					ResultType: loghttp.ResultTypeScalar,
					Result:     sampleStream,
				},
				Warnings: result.Warnings,
			},
			Statistics: result.Statistics,
		}, nil
	case logqlmodel.Streams:
		return &LogstoreResponse{
			Direction: params.Direction(),
			Limit:     params.Limit(),
			Data: LogstoreData{
				ResultType: loghttp.ResultTypeStream,
				Result:     data,
			},
			Status:     "success",
			Warnings:   result.Warnings,
			Statistics: result.Statistics,
		}, nil
	case sketch.TopKMatrix:
		sk, err := data.ToProto()
		return &TopKSketchesResponse{
			Response:   sk,
			Warnings:   result.Warnings,
			Statistics: result.Statistics,
		}, err
	case logql.ProbabilisticQuantileMatrix:
		r := data.ToProto()
		data.Release()
		return &QuantileSketchResponse{
			Response:   r,
			Warnings:   result.Warnings,
			Statistics: result.Statistics,
		}, nil
	case logql.CountMinSketchVector:
		r, err := data.ToProto()
		return &CountMinSketchResponse{
			Response:   r,
			Warnings:   result.Warnings,
			Statistics: result.Statistics,
		}, err
	}

	return nil, fmt.Errorf("unsupported data type: %T", result.Data)
}

func ResponseToResult(resp queryrangebase.Response) (logqlmodel.Result, error) {
	switch r := resp.(type) {
	case *LogstoreResponse:
		if r.Error != "" {
			return logqlmodel.Result{}, fmt.Errorf("%s: %s", r.ErrorType, r.Error)
		}

		streams := make(logqlmodel.Streams, 0, len(r.Data.Result))

		for _, stream := range r.Data.Result {
			streams = append(streams, stream)
		}

		return logqlmodel.Result{
			Statistics: r.Statistics,
			Data:       streams,
			Headers:    resp.GetHeaders(),
			Warnings:   r.Warnings,
		}, nil

	case *LogstorePromResponse:
		if r.Response.Error != "" {
			return logqlmodel.Result{}, fmt.Errorf("%s: %s", r.Response.ErrorType, r.Response.Error)
		}
		if r.Response.Data.ResultType == loghttp.ResultTypeVector {
			return logqlmodel.Result{
				Statistics: r.Statistics,
				Data:       sampleStreamToVector(r.Response.Data.Result),
				Headers:    resp.GetHeaders(),
				Warnings:   r.Response.Warnings,
			}, nil
		}
		return logqlmodel.Result{
			Statistics: r.Statistics,
			Data:       sampleStreamToMatrix(r.Response.Data.Result),
			Headers:    resp.GetHeaders(),
			Warnings:   r.Response.Warnings,
		}, nil
	case *TopKSketchesResponse:
		matrix, err := sketch.TopKMatrixFromProto(r.Response)
		if err != nil {
			return logqlmodel.Result{}, fmt.Errorf("cannot decode topk sketch: %w", err)
		}

		return logqlmodel.Result{
			Data:       matrix,
			Headers:    resp.GetHeaders(),
			Warnings:   r.Warnings,
			Statistics: r.Statistics,
		}, nil
	case *QuantileSketchResponse:
		matrix, err := logql.ProbabilisticQuantileMatrixFromProto(r.Response)
		if err != nil {
			return logqlmodel.Result{}, fmt.Errorf("cannot decode quantile sketch: %w", err)
		}
		return logqlmodel.Result{
			Data:       matrix,
			Headers:    resp.GetHeaders(),
			Warnings:   r.Warnings,
			Statistics: r.Statistics,
		}, nil
	case *CountMinSketchResponse:
		cms, err := logql.CountMinSketchVectorFromProto(r.Response)
		if err != nil {
			return logqlmodel.Result{}, fmt.Errorf("cannot decode count min sketch vector: %w", err)
		}
		return logqlmodel.Result{
			Data:       cms,
			Headers:    resp.GetHeaders(),
			Warnings:   r.Warnings,
			Statistics: r.Statistics,
		}, nil
	default:
		return logqlmodel.Result{}, fmt.Errorf("cannot decode (%T)", resp)
	}
}

func QueryResponseUnwrap(res *QueryResponse) (queryrangebase.Response, error) {
	if res.Status != nil && res.Status.Code != int32(rpc.OK) {
		return nil, status.ErrorProto(res.Status)
	}

	switch concrete := res.Response.(type) {
	case *QueryResponse_Series:
		return concrete.Series, nil
	case *QueryResponse_Labels:
		return concrete.Labels, nil
	case *QueryResponse_Stats:
		return concrete.Stats, nil
	case *QueryResponse_ShardsResponse:
		return concrete.ShardsResponse, nil
	case *QueryResponse_Prom:
		return concrete.Prom, nil
	case *QueryResponse_Streams:
		return concrete.Streams, nil
	case *QueryResponse_Volume:
		return concrete.Volume, nil
	case *QueryResponse_TopkSketches:
		return concrete.TopkSketches, nil
	case *QueryResponse_QuantileSketches:
		return concrete.QuantileSketches, nil
	case *QueryResponse_PatternsResponse:
		return concrete.PatternsResponse, nil
	case *QueryResponse_DetectedLabels:
		return concrete.DetectedLabels, nil
	case *QueryResponse_DetectedFields:
		return concrete.DetectedFields, nil
	case *QueryResponse_CountMinSketches:
		return concrete.CountMinSketches, nil
	default:
		return nil, fmt.Errorf("unsupported QueryResponse response type, got (%T)", res.Response)
	}
}

func QueryResponseWrap(res queryrangebase.Response) (*QueryResponse, error) {
	p := &QueryResponse{
		Status: status.New(codes.OK, "").Proto(),
	}

	switch response := res.(type) {
	case *LogstorePromResponse:
		p.Response = &QueryResponse_Prom{response}
	case *LogstoreResponse:
		p.Response = &QueryResponse_Streams{response}
	case *LogstoreSeriesResponse:
		p.Response = &QueryResponse_Series{response}
	case *MergedSeriesResponseView:
		mat, err := response.Materialize()
		if err != nil {
			return p, err
		}
		p.Response = &QueryResponse_Series{mat}
	case *LogstoreLabelNamesResponse:
		p.Response = &QueryResponse_Labels{response}
	case *IndexStatsResponse:
		p.Response = &QueryResponse_Stats{response}
	case *VolumeResponse:
		p.Response = &QueryResponse_Volume{response}
	case *TopKSketchesResponse:
		p.Response = &QueryResponse_TopkSketches{response}
	case *QuantileSketchResponse:
		p.Response = &QueryResponse_QuantileSketches{response}
	case *ShardsResponse:
		p.Response = &QueryResponse_ShardsResponse{response}
	case *QueryPatternsResponse:
		p.Response = &QueryResponse_PatternsResponse{response}
	case *DetectedLabelsResponse:
		p.Response = &QueryResponse_DetectedLabels{response}
	case *DetectedFieldsResponse:
		p.Response = &QueryResponse_DetectedFields{response}
	case *CountMinSketchResponse:
		p.Response = &QueryResponse_CountMinSketches{response}
	default:
		return nil, fmt.Errorf("invalid response format, got (%T)", res)
	}

	return p, nil
}

// QueryResponseWrapError wraps an error in the QueryResponse protobuf.
func QueryResponseWrapError(err error) *QueryResponse {
	return &QueryResponse{
		Status: server.WrapError(err),
	}
}

func (Codec) QueryRequestUnwrap(ctx context.Context, req *QueryRequest) (queryrangebase.Request, context.Context, error) {
	if req == nil {
		return nil, ctx, nil
	}

	// Add query tags
	if queryTags, ok := req.Metadata[string(httpreq.QueryTagsHTTPHeader)]; ok {
		ctx = httpreq.InjectQueryTags(ctx, queryTags)
	}

	// Add actor path
	if actor, ok := req.Metadata[httpreq.LogstoreActorPathHeader]; ok {
		ctx = httpreq.InjectActorPath(ctx, actor)
	}

	// Add disable wrappers
	if disableWrappers, ok := req.Metadata[httpreq.LogstoreDisablePipelineWrappersHeader]; ok {
		ctx = httpreq.InjectHeader(ctx, httpreq.LogstoreDisablePipelineWrappersHeader, disableWrappers)
	}

	// Add limits
	if encodedLimits, ok := req.Metadata[querylimits.HTTPHeaderQueryLimitsKey]; ok {
		limits, err := querylimits.UnmarshalQueryLimits([]byte(encodedLimits))
		if err != nil {
			return nil, ctx, err
		}
		ctx = querylimits.InjectQueryLimitsContext(ctx, *limits)
	}

	// Add query time
	if queueTimeHeader, ok := req.Metadata[string(httpreq.QueryQueueTimeHTTPHeader)]; ok {
		queueTime, err := time.ParseDuration(queueTimeHeader)
		if err == nil {
			ctx = context.WithValue(ctx, httpreq.QueryQueueTimeHTTPHeader, queueTime)
		}
	}

	switch concrete := req.Request.(type) {
	case *QueryRequest_Series:
		return concrete.Series, ctx, nil
	case *QueryRequest_Instant:
		if concrete.Instant.Plan == nil {
			parsed, err := syntax.ParseExpr(concrete.Instant.GetQuery())
			if err != nil {
				return nil, ctx, httpgrpc.Errorf(http.StatusBadRequest, "%s", err.Error())
			}
			concrete.Instant.Plan = &plan.QueryPlan{
				AST: parsed,
			}
		}

		return concrete.Instant, ctx, nil
	case *QueryRequest_Stats:
		return concrete.Stats, ctx, nil
	case *QueryRequest_ShardsRequest:
		return concrete.ShardsRequest, ctx, nil
	case *QueryRequest_Volume:
		return concrete.Volume, ctx, nil
	case *QueryRequest_Streams:
		if concrete.Streams.Plan == nil {
			parsed, err := syntax.ParseExpr(concrete.Streams.GetQuery())
			if err != nil {
				return nil, ctx, httpgrpc.Errorf(http.StatusBadRequest, "%s", err.Error())
			}
			concrete.Streams.Plan = &plan.QueryPlan{
				AST: parsed,
			}
		}

		return concrete.Streams, ctx, nil
	case *QueryRequest_Labels:
		return &LabelRequest{
			LabelRequest: *concrete.Labels,
		}, ctx, nil
	case *QueryRequest_PatternsRequest:
		return concrete.PatternsRequest, ctx, nil
	case *QueryRequest_DetectedLabels:
		return &DetectedLabelsRequest{
			DetectedLabelsRequest: *concrete.DetectedLabels,
		}, ctx, nil
	case *QueryRequest_DetectedFields:
		return &DetectedFieldsRequest{
			DetectedFieldsRequest: *concrete.DetectedFields,
		}, ctx, nil
	default:
		return nil, ctx, fmt.Errorf("unsupported request type while unwrapping, got (%T)", req.Request)
	}
}

func (Codec) QueryRequestWrap(ctx context.Context, r queryrangebase.Request) (*QueryRequest, error) {
	result := &QueryRequest{
		Metadata: make(map[string]string),
	}

	switch req := r.(type) {
	case *LogstoreSeriesRequest:
		result.Request = &QueryRequest_Series{Series: req}
	case *LabelRequest:
		result.Request = &QueryRequest_Labels{Labels: &req.LabelRequest}
	case *logproto.IndexStatsRequest:
		result.Request = &QueryRequest_Stats{Stats: req}
	case *logproto.VolumeRequest:
		result.Request = &QueryRequest_Volume{Volume: req}
	case *LogstoreInstantRequest:
		result.Request = &QueryRequest_Instant{Instant: req}
	case *LogstoreRequest:
		result.Request = &QueryRequest_Streams{Streams: req}
	case *logproto.ShardsRequest:
		result.Request = &QueryRequest_ShardsRequest{ShardsRequest: req}
	case *logproto.QueryPatternsRequest:
		result.Request = &QueryRequest_PatternsRequest{PatternsRequest: req}
	case *DetectedLabelsRequest:
		result.Request = &QueryRequest_DetectedLabels{DetectedLabels: &req.DetectedLabelsRequest}
	case *DetectedFieldsRequest:
		result.Request = &QueryRequest_DetectedFields{DetectedFields: &req.DetectedFieldsRequest}
	default:
		return nil, fmt.Errorf("unsupported request type while wrapping, got (%T)", r)
	}

	// Add query tags
	queryTags := getQueryTags(ctx)
	if queryTags != "" {
		result.Metadata[string(httpreq.QueryTagsHTTPHeader)] = queryTags
	}

	// Add actor path
	actor := httpreq.ExtractHeader(ctx, httpreq.LogstoreActorPathHeader)
	if actor != "" {
		result.Metadata[httpreq.LogstoreActorPathHeader] = actor
	}

	// Keep disable wrappers
	disableWrappers := httpreq.ExtractHeader(ctx, httpreq.LogstoreDisablePipelineWrappersHeader)
	if disableWrappers != "" {
		result.Metadata[httpreq.LogstoreDisablePipelineWrappersHeader] = disableWrappers
	}

	// Add limits
	limits := querylimits.ExtractQueryLimitsContext(ctx)
	if limits != nil {
		encodedLimits, err := querylimits.MarshalQueryLimits(limits)
		if err != nil {
			return nil, err
		}
		result.Metadata[querylimits.HTTPHeaderQueryLimitsKey] = string(encodedLimits)
	}

	// Add org ID
	orgID, err := user.ExtractOrgID(ctx)
	if err != nil {
		return nil, err
	}
	result.Metadata[user.OrgIDHeaderName] = orgID

	// Tracing
	otel.GetTextMapPropagator().Inject(ctx, propagation.MapCarrier(result.Metadata))

	return result, nil
}
