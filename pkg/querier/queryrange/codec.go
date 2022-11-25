package queryrange

import (
	"bytes"
	"container/heap"
	"context"
	"errors"
	"fmt"
	io "io"
	"net/http"
	"net/url"
	"sort"
	strings "strings"
	"time"

	json "github.com/json-iterator/go"
	"github.com/opentracing/opentracing-go"
	otlog "github.com/opentracing/opentracing-go/log"
	"github.com/prometheus/prometheus/model/timestamp"
	"github.com/weaveworks/common/httpgrpc"

	"example.com/acme/logstore/pkg/loghttp"
	"example.com/acme/logstore/pkg/logproto"
	"example.com/acme/logstore/pkg/logql"
	"example.com/acme/logstore/pkg/logql/syntax"
	"example.com/acme/logstore/pkg/logqlmodel"
	"example.com/acme/logstore/pkg/logqlmodel/stats"
	"example.com/acme/logstore/pkg/querier/queryrange/queryrangebase"
	"example.com/acme/logstore/pkg/util"
	"example.com/acme/logstore/pkg/util/httpreq"
	"example.com/acme/logstore/pkg/util/marshal"
	marshal_legacy "example.com/acme/logstore/pkg/util/marshal/legacy"
)

var LogstoreCodec = &Codec{}

type Codec struct{}

func (r *LogstoreRequest) GetEnd() int64 {
	return r.EndTs.UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func (r *LogstoreRequest) GetStart() int64 {
	return r.StartTs.UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func (r *LogstoreRequest) WithStartEnd(s int64, e int64) queryrangebase.Request {
	new := *r
	new.StartTs = time.Unix(0, s*int64(time.Millisecond))
	new.EndTs = time.Unix(0, e*int64(time.Millisecond))
	return &new
}

func (r *LogstoreRequest) WithStartEndTime(s time.Time, e time.Time) *LogstoreRequest {
	new := *r
	new.StartTs = s
	new.EndTs = e
	return &new
}

func (r *LogstoreRequest) WithQuery(query string) queryrangebase.Request {
	new := *r
	new.Query = query
	return &new
}

func (r *LogstoreRequest) WithShards(shards logql.Shards) *LogstoreRequest {
	new := *r
	new.Shards = shards.Encode()
	return &new
}

func (r *LogstoreRequest) LogToSpan(sp opentracing.Span) {
	sp.LogFields(
		otlog.String("query", r.GetQuery()),
		otlog.String("start", timestamp.Time(r.GetStart()).String()),
		otlog.String("end", timestamp.Time(r.GetEnd()).String()),
		otlog.Int64("step (ms)", r.GetStep()),
		otlog.Int64("interval (ms)", r.GetInterval()),
		otlog.Int64("limit", int64(r.GetLimit())),
		otlog.String("direction", r.GetDirection().String()),
		otlog.String("shards", strings.Join(r.GetShards(), ",")),
	)
}

func (*LogstoreRequest) GetCachingOptions() (res queryrangebase.CachingOptions) { return }

func (r *LogstoreInstantRequest) GetStep() int64 {
	return 0
}

func (r *LogstoreInstantRequest) GetEnd() int64 {
	return r.TimeTs.UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func (r *LogstoreInstantRequest) GetStart() int64 {
	return r.TimeTs.UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func (r *LogstoreInstantRequest) WithStartEnd(s int64, e int64) queryrangebase.Request {
	new := *r
	new.TimeTs = time.Unix(0, s*int64(time.Millisecond))
	return &new
}

func (r *LogstoreInstantRequest) WithQuery(query string) queryrangebase.Request {
	new := *r
	new.Query = query
	return &new
}

func (r *LogstoreInstantRequest) WithShards(shards logql.Shards) *LogstoreInstantRequest {
	new := *r
	new.Shards = shards.Encode()
	return &new
}

func (r *LogstoreInstantRequest) LogToSpan(sp opentracing.Span) {
	sp.LogFields(
		otlog.String("query", r.GetQuery()),
		otlog.String("ts", timestamp.Time(r.GetStart()).String()),
		otlog.Int64("limit", int64(r.GetLimit())),
		otlog.String("direction", r.GetDirection().String()),
		otlog.String("shards", strings.Join(r.GetShards(), ",")),
	)
}

func (*LogstoreInstantRequest) GetCachingOptions() (res queryrangebase.CachingOptions) { return }

func (r *LogstoreSeriesRequest) GetEnd() int64 {
	return r.EndTs.UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func (r *LogstoreSeriesRequest) GetStart() int64 {
	return r.StartTs.UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func (r *LogstoreSeriesRequest) WithStartEnd(s int64, e int64) queryrangebase.Request {
	new := *r
	new.StartTs = time.Unix(0, s*int64(time.Millisecond))
	new.EndTs = time.Unix(0, e*int64(time.Millisecond))
	return &new
}

func (r *LogstoreSeriesRequest) WithQuery(query string) queryrangebase.Request {
	new := *r
	return &new
}

func (r *LogstoreSeriesRequest) GetQuery() string {
	return ""
}

func (r *LogstoreSeriesRequest) GetStep() int64 {
	return 0
}

func (r *LogstoreSeriesRequest) LogToSpan(sp opentracing.Span) {
	sp.LogFields(
		otlog.String("matchers", strings.Join(r.GetMatch(), ",")),
		otlog.String("start", timestamp.Time(r.GetStart()).String()),
		otlog.String("end", timestamp.Time(r.GetEnd()).String()),
		otlog.String("shards", strings.Join(r.GetShards(), ",")),
	)
}

func (*LogstoreSeriesRequest) GetCachingOptions() (res queryrangebase.CachingOptions) { return }

func (r *LogstoreLabelNamesRequest) GetEnd() int64 {
	return r.EndTs.UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func (r *LogstoreLabelNamesRequest) GetStart() int64 {
	return r.StartTs.UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func (r *LogstoreLabelNamesRequest) WithStartEnd(s int64, e int64) queryrangebase.Request {
	new := *r
	new.StartTs = time.Unix(0, s*int64(time.Millisecond))
	new.EndTs = time.Unix(0, e*int64(time.Millisecond))
	return &new
}

func (r *LogstoreLabelNamesRequest) WithQuery(query string) queryrangebase.Request {
	new := *r
	return &new
}

func (r *LogstoreLabelNamesRequest) GetQuery() string {
	return ""
}

func (r *LogstoreLabelNamesRequest) GetStep() int64 {
	return 0
}

func (r *LogstoreLabelNamesRequest) LogToSpan(sp opentracing.Span) {
	sp.LogFields(
		otlog.String("start", timestamp.Time(r.GetStart()).String()),
		otlog.String("end", timestamp.Time(r.GetEnd()).String()),
	)
}

func (*LogstoreLabelNamesRequest) GetCachingOptions() (res queryrangebase.CachingOptions) { return }

func (Codec) DecodeRequest(_ context.Context, r *http.Request, forwardHeaders []string) (queryrangebase.Request, error) {
	if err := r.ParseForm(); err != nil {
		return nil, httpgrpc.Errorf(http.StatusBadRequest, err.Error())
	}

	switch op := getOperation(r.URL.Path); op {
	case QueryRangeOp:
		req, err := loghttp.ParseRangeQuery(r)
		if err != nil {
			return nil, httpgrpc.Errorf(http.StatusBadRequest, err.Error())
		}
		return &LogstoreRequest{
			Query:     req.Query,
			Limit:     req.Limit,
			Direction: req.Direction,
			StartTs:   req.Start.UTC(),
			EndTs:     req.End.UTC(),
			Step:      req.Step.Milliseconds(),
			Interval:  req.Interval.Milliseconds(),
			Path:      r.URL.Path,
			Shards:    req.Shards,
		}, nil
	case InstantQueryOp:
		req, err := loghttp.ParseInstantQuery(r)
		if err != nil {
			return nil, httpgrpc.Errorf(http.StatusBadRequest, err.Error())
		}
		return &LogstoreInstantRequest{
			Query:     req.Query,
			Limit:     req.Limit,
			Direction: req.Direction,
			TimeTs:    req.Ts.UTC(),
			Path:      r.URL.Path,
			Shards:    req.Shards,
		}, nil
	case SeriesOp:
		req, err := loghttp.ParseAndValidateSeriesQuery(r)
		if err != nil {
			return nil, httpgrpc.Errorf(http.StatusBadRequest, err.Error())
		}
		return &LogstoreSeriesRequest{
			Match:   req.Groups,
			StartTs: req.Start.UTC(),
			EndTs:   req.End.UTC(),
			Path:    r.URL.Path,
			Shards:  req.Shards,
		}, nil
	case LabelNamesOp:
		req, err := loghttp.ParseLabelQuery(r)
		if err != nil {
			return nil, httpgrpc.Errorf(http.StatusBadRequest, err.Error())
		}
		return &LogstoreLabelNamesRequest{
			StartTs: *req.Start,
			EndTs:   *req.End,
			Path:    r.URL.Path,
		}, nil
	case IndexStatsOp:
		req, err := loghttp.ParseIndexStatsQuery(r)
		if err != nil {
			return nil, httpgrpc.Errorf(http.StatusBadRequest, err.Error())
		}
		from, through := util.RoundToMilliseconds(req.Start, req.End)
		return &logproto.IndexStatsRequest{
			From:     from,
			Through:  through,
			Matchers: req.Query,
		}, err
	default:
		return nil, httpgrpc.Errorf(http.StatusBadRequest, fmt.Sprintf("unknown request path: %s", r.URL.Path))
	}
}

func (Codec) EncodeRequest(ctx context.Context, r queryrangebase.Request) (*http.Request, error) {
	header := make(http.Header)
	queryTags := getQueryTags(ctx)
	if queryTags != "" {
		header.Set(string(httpreq.QueryTagsHTTPHeader), queryTags)
	}

	switch request := r.(type) {
	case *LogstoreRequest:
		params := url.Values{
			"start":     []string{fmt.Sprintf("%d", request.StartTs.UnixNano())},
			"end":       []string{fmt.Sprintf("%d", request.EndTs.UnixNano())},
			"query":     []string{request.Query},
			"direction": []string{request.Direction.String()},
			"limit":     []string{fmt.Sprintf("%d", request.Limit)},
		}
		if len(request.Shards) > 0 {
			params["shards"] = request.Shards
		}
		if request.Step != 0 {
			params["step"] = []string{fmt.Sprintf("%f", float64(request.Step)/float64(1e3))}
		}
		if request.Interval != 0 {
			params["interval"] = []string{fmt.Sprintf("%f", float64(request.Interval)/float64(1e3))}
		}
		u := &url.URL{
			// the request could come /api/prom/query but we want to only use the new api.
			Path:     "/logstore/api/v1/query_range",
			RawQuery: params.Encode(),
		}
		req := &http.Request{
			Method:     "GET",
			RequestURI: u.String(), // This is what the httpgrpc code looks at.
			URL:        u,
			Body:       http.NoBody,
			Header:     header,
		}

		return req.WithContext(ctx), nil
	case *LogstoreSeriesRequest:
		params := url.Values{
			"start":   []string{fmt.Sprintf("%d", request.StartTs.UnixNano())},
			"end":     []string{fmt.Sprintf("%d", request.EndTs.UnixNano())},
			"match[]": request.Match,
		}
		if len(request.Shards) > 0 {
			params["shards"] = request.Shards
		}
		u := &url.URL{
			Path:     "/logstore/api/v1/series",
			RawQuery: params.Encode(),
		}
		req := &http.Request{
			Method:     "GET",
			RequestURI: u.String(), // This is what the httpgrpc code looks at.
			URL:        u,
			Body:       http.NoBody,
			Header:     header,
		}
		return req.WithContext(ctx), nil
	case *LogstoreLabelNamesRequest:
		params := url.Values{
			"start": []string{fmt.Sprintf("%d", request.StartTs.UnixNano())},
			"end":   []string{fmt.Sprintf("%d", request.EndTs.UnixNano())},
		}

		u := &url.URL{
			Path:     request.Path, // NOTE: this could be either /label or /label/{name}/values endpoint. So forward the original path as it is.
			RawQuery: params.Encode(),
		}
		req := &http.Request{
			Method:     "GET",
			RequestURI: u.String(), // This is what the httpgrpc code looks at.
			URL:        u,
			Body:       http.NoBody,
			Header:     header,
		}
		return req.WithContext(ctx), nil
	case *LogstoreInstantRequest:
		params := url.Values{
			"query":     []string{request.Query},
			"direction": []string{request.Direction.String()},
			"limit":     []string{fmt.Sprintf("%d", request.Limit)},
			"time":      []string{fmt.Sprintf("%d", request.TimeTs.UnixNano())},
		}
		if len(request.Shards) > 0 {
			params["shards"] = request.Shards
		}
		u := &url.URL{
			// the request could come /api/prom/query but we want to only use the new api.
			Path:     "/logstore/api/v1/query",
			RawQuery: params.Encode(),
		}
		req := &http.Request{
			Method:     "GET",
			RequestURI: u.String(), // This is what the httpgrpc code looks at.
			URL:        u,
			Body:       http.NoBody,
			Header:     header,
		}

		return req.WithContext(ctx), nil
	case *logproto.IndexStatsRequest:
		params := url.Values{
			"start": []string{fmt.Sprintf("%d", request.From.Time().UnixNano())},
			"end":   []string{fmt.Sprintf("%d", request.Through.Time().UnixNano())},
			"query": []string{request.GetQuery()},
		}
		u := &url.URL{
			Path:     "/logstore/api/v1/index/stats",
			RawQuery: params.Encode(),
		}
		req := &http.Request{
			Method:     "GET",
			RequestURI: u.String(), // This is what the httpgrpc code looks at.
			URL:        u,
			Body:       http.NoBody,
			Header:     header,
		}
		return req.WithContext(ctx), nil
	default:
		return nil, httpgrpc.Errorf(http.StatusInternalServerError, "invalid request format")
	}
}

type Buffer interface {
	Bytes() []byte
}

func (Codec) DecodeResponse(ctx context.Context, r *http.Response, req queryrangebase.Request) (queryrangebase.Response, error) {
	if r.StatusCode/100 != 2 {
		body, _ := io.ReadAll(r.Body)
		return nil, httpgrpc.Errorf(r.StatusCode, string(body))
	}

	var buf []byte
	var err error
	if buffer, ok := r.Body.(Buffer); ok {
		buf = buffer.Bytes()
	} else {
		buf, err = io.ReadAll(r.Body)
		if err != nil {
			return nil, httpgrpc.Errorf(http.StatusInternalServerError, "error decoding response: %v", err)
		}
	}

	switch req := req.(type) {
	case *LogstoreSeriesRequest:
		var resp loghttp.SeriesResponse
		if err := json.Unmarshal(buf, &resp); err != nil {
			return nil, httpgrpc.Errorf(http.StatusInternalServerError, "error decoding response: %v", err)
		}

		data := make([]logproto.SeriesIdentifier, 0, len(resp.Data))
		for _, label := range resp.Data {
			d := logproto.SeriesIdentifier{
				Labels: label.Map(),
			}
			data = append(data, d)
		}

		return &LogstoreSeriesResponse{
			Status:  resp.Status,
			Version: uint32(loghttp.GetVersion(req.Path)),
			Data:    data,
			Headers: httpResponseHeadersToPromResponseHeaders(r.Header),
		}, nil
	case *LogstoreLabelNamesRequest:
		var resp loghttp.LabelResponse
		if err := json.Unmarshal(buf, &resp); err != nil {
			return nil, httpgrpc.Errorf(http.StatusInternalServerError, "error decoding response: %v", err)
		}
		return &LogstoreLabelNamesResponse{
			Status:  resp.Status,
			Version: uint32(loghttp.GetVersion(req.Path)),
			Data:    resp.Data,
			Headers: httpResponseHeadersToPromResponseHeaders(r.Header),
		}, nil
	case *logproto.IndexStatsRequest:
		var resp logproto.IndexStatsResponse
		if err := json.Unmarshal(buf, &resp); err != nil {
			return nil, httpgrpc.Errorf(http.StatusInternalServerError, "error decoding response: %v", err)
		}
		return &IndexStatsResponse{
			Response: &resp,
			Headers:  httpResponseHeadersToPromResponseHeaders(r.Header),
		}, nil
	default:
		var resp loghttp.QueryResponse
		if err := resp.UnmarshalJSON(buf); err != nil {
			return nil, httpgrpc.Errorf(http.StatusInternalServerError, "error decoding response: %v", err)
		}
		switch string(resp.Data.ResultType) {
		case loghttp.ResultTypeMatrix:
			return &LogstorePromResponse{
				Response: &queryrangebase.PrometheusResponse{
					Status: resp.Status,
					Data: queryrangebase.PrometheusData{
						ResultType: loghttp.ResultTypeMatrix,
						Result:     toProtoMatrix(resp.Data.Result.(loghttp.Matrix)),
					},
					Headers: convertPrometheusResponseHeadersToPointers(httpResponseHeadersToPromResponseHeaders(r.Header)),
				},
				Statistics: resp.Data.Statistics,
			}, nil
		case loghttp.ResultTypeStream:
			// This is the same as in querysharding.go
			params, err := paramsFromRequest(req)
			if err != nil {
				return nil, err
			}

			var path string
			switch r := req.(type) {
			case *LogstoreRequest:
				path = r.GetPath()
			case *LogstoreInstantRequest:
				path = r.GetPath()
			default:
				return nil, fmt.Errorf("expected *LogstoreRequest or *LogstoreInstantRequest, got (%T)", r)
			}
			return &LogstoreResponse{
				Status:     resp.Status,
				Direction:  params.Direction(),
				Limit:      params.Limit(),
				Version:    uint32(loghttp.GetVersion(path)),
				Statistics: resp.Data.Statistics,
				Data: LogstoreData{
					ResultType: loghttp.ResultTypeStream,
					Result:     resp.Data.Result.(loghttp.Streams).ToProto(),
				},
				Headers: httpResponseHeadersToPromResponseHeaders(r.Header),
			}, nil
		case loghttp.ResultTypeVector:
			return &LogstorePromResponse{
				Response: &queryrangebase.PrometheusResponse{
					Status: resp.Status,
					Data: queryrangebase.PrometheusData{
						ResultType: loghttp.ResultTypeVector,
						Result:     toProtoVector(resp.Data.Result.(loghttp.Vector)),
					},
					Headers: convertPrometheusResponseHeadersToPointers(httpResponseHeadersToPromResponseHeaders(r.Header)),
				},
				Statistics: resp.Data.Statistics,
			}, nil
		default:
			return nil, httpgrpc.Errorf(http.StatusInternalServerError, "unsupported response type, got (%s)", string(resp.Data.ResultType))
		}
	}
}

func (Codec) EncodeResponse(ctx context.Context, res queryrangebase.Response) (*http.Response, error) {
	sp, _ := opentracing.StartSpanFromContext(ctx, "codec.EncodeResponse")
	defer sp.Finish()
	var buf bytes.Buffer

	switch response := res.(type) {
	case *LogstorePromResponse:
		return response.encode(ctx)
	case *LogstoreResponse:
		streams := make([]logproto.Stream, len(response.Data.Result))

		for i, stream := range response.Data.Result {
			streams[i] = logproto.Stream{
				Labels:  stream.Labels,
				Entries: stream.Entries,
			}
		}
		result := logqlmodel.Result{
			Data:       logqlmodel.Streams(streams),
			Statistics: response.Statistics,
		}
		if loghttp.Version(response.Version) == loghttp.VersionLegacy {
			if err := marshal_legacy.WriteQueryResponseJSON(result, &buf); err != nil {
				return nil, err
			}
		} else {
			if err := marshal.WriteQueryResponseJSON(result, &buf); err != nil {
				return nil, err
			}
		}

	case *LogstoreSeriesResponse:
		result := logproto.SeriesResponse{
			Series: response.Data,
		}
		if err := marshal.WriteSeriesResponseJSON(result, &buf); err != nil {
			return nil, err
		}
	case *LogstoreLabelNamesResponse:
		if loghttp.Version(response.Version) == loghttp.VersionLegacy {
			if err := marshal_legacy.WriteLabelResponseJSON(logproto.LabelResponse{Values: response.Data}, &buf); err != nil {
				return nil, err
			}
		} else {
			if err := marshal.WriteLabelResponseJSON(logproto.LabelResponse{Values: response.Data}, &buf); err != nil {
				return nil, err
			}
		}
	case *IndexStatsResponse:
		if err := marshal.WriteIndexStatsResponseJSON(response.Response, &buf); err != nil {
			return nil, err
		}

	default:
		return nil, httpgrpc.Errorf(http.StatusInternalServerError, "invalid response format")
	}

	sp.LogFields(otlog.Int("bytes", buf.Len()))

	resp := http.Response{
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		Body:       io.NopCloser(&buf),
		StatusCode: http.StatusOK,
	}
	return &resp, nil
}

// NOTE: When we would start caching response from non-metric queries we would have to consider cache gen headers as well in
// MergeResponse implementation for Logstore codecs same as it is done in Corestore at https://github.com/corestoreproject/corestore/blob/21bad57b346c730d684d6d0205efef133422ab28/pkg/querier/queryrange/query_range.go#L170
func (Codec) MergeResponse(responses ...queryrangebase.Response) (queryrangebase.Response, error) {
	if len(responses) == 0 {
		return nil, errors.New("merging responses requires at least one response")
	}
	var mergedStats stats.Result
	switch responses[0].(type) {
	case *LogstorePromResponse:

		promResponses := make([]queryrangebase.Response, 0, len(responses))
		for _, res := range responses {
			mergedStats.Merge(res.(*LogstorePromResponse).Statistics)
			promResponses = append(promResponses, res.(*LogstorePromResponse).Response)
		}
		promRes, err := queryrangebase.PrometheusCodec.MergeResponse(promResponses...)
		if err != nil {
			return nil, err
		}
		return &LogstorePromResponse{
			Response:   promRes.(*queryrangebase.PrometheusResponse),
			Statistics: mergedStats,
		}, nil
	case *LogstoreResponse:
		return mergeLogstoreResponse(responses...), nil
	case *LogstoreSeriesResponse:
		logstoreSeriesRes := responses[0].(*LogstoreSeriesResponse)

		var logstoreSeriesData []logproto.SeriesIdentifier
		uniqueSeries := make(map[string]struct{})

		// only unique series should be merged
		for _, res := range responses {
			logstoreResult := res.(*LogstoreSeriesResponse)
			for _, series := range logstoreResult.Data {
				if _, ok := uniqueSeries[series.String()]; !ok {
					logstoreSeriesData = append(logstoreSeriesData, series)
					uniqueSeries[series.String()] = struct{}{}
				}
			}
		}

		return &LogstoreSeriesResponse{
			Status:  logstoreSeriesRes.Status,
			Version: logstoreSeriesRes.Version,
			Data:    logstoreSeriesData,
		}, nil
	case *LogstoreLabelNamesResponse:
		labelNameRes := responses[0].(*LogstoreLabelNamesResponse)
		uniqueNames := make(map[string]struct{})
		names := []string{}

		// only unique name should be merged
		for _, res := range responses {
			logstoreResult := res.(*LogstoreLabelNamesResponse)
			for _, labelName := range logstoreResult.Data {
				if _, ok := uniqueNames[labelName]; !ok {
					names = append(names, labelName)
					uniqueNames[labelName] = struct{}{}
				}
			}
		}

		return &LogstoreLabelNamesResponse{
			Status:  labelNameRes.Status,
			Version: labelNameRes.Version,
			Data:    names,
		}, nil
	default:
		return nil, errors.New("unknown response in merging responses")
	}
}

// mergeOrderedNonOverlappingStreams merges a set of ordered, nonoverlapping responses by concatenating matching streams then running them through a heap to pull out limit values
func mergeOrderedNonOverlappingStreams(resps []*LogstoreResponse, limit uint32, direction logproto.Direction) []logproto.Stream {
	var total int

	// turn resps -> map[labels] []entries
	groups := make(map[string]*byDir)
	for _, resp := range resps {
		for _, stream := range resp.Data.Result {
			s, ok := groups[stream.Labels]
			if !ok {
				s = &byDir{
					direction: direction,
					labels:    stream.Labels,
				}
				groups[stream.Labels] = s
			}

			s.markers = append(s.markers, stream.Entries)
			total += len(stream.Entries)
		}

		// optimization: since limit has been reached, no need to append entries from subsequent responses
		if total >= int(limit) {
			break
		}
	}

	keys := make([]string, 0, len(groups))
	for key := range groups {
		keys = append(keys, key)
	}
	if direction == logproto.BACKWARD {
		sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	} else {
		sort.Strings(keys)
	}

	// escape hatch, can just return all the streams
	if total <= int(limit) {
		results := make([]logproto.Stream, 0, len(keys))
		for _, key := range keys {
			results = append(results, logproto.Stream{
				Labels:  key,
				Entries: groups[key].merge(),
			})
		}
		return results
	}

	pq := &priorityqueue{
		direction: direction,
	}

	for _, key := range keys {
		stream := &logproto.Stream{
			Labels:  key,
			Entries: groups[key].merge(),
		}
		if len(stream.Entries) > 0 {
			pq.streams = append(pq.streams, stream)
		}
	}

	heap.Init(pq)

	resultDict := make(map[string]*logproto.Stream)

	// we want the min(limit, num_entries)
	for i := 0; i < int(limit) && pq.Len() > 0; i++ {
		// grab the next entry off the queue. This will be a stream (to preserve labels) with one entry.
		next := heap.Pop(pq).(*logproto.Stream)

		s, ok := resultDict[next.Labels]
		if !ok {
			s = &logproto.Stream{
				Labels:  next.Labels,
				Entries: make([]logproto.Entry, 0, int(limit)/len(keys)), // allocation hack -- assume uniform distribution across labels
			}
			resultDict[next.Labels] = s
		}
		// TODO: make allocation friendly
		s.Entries = append(s.Entries, next.Entries...)
	}

	results := make([]logproto.Stream, 0, len(resultDict))
	for _, key := range keys {
		stream, ok := resultDict[key]
		if ok {
			results = append(results, *stream)
		}
	}

	return results
}

func toProtoMatrix(m loghttp.Matrix) []queryrangebase.SampleStream {
	res := make([]queryrangebase.SampleStream, 0, len(m))

	if len(m) == 0 {
		return res
	}

	for _, stream := range m {
		samples := make([]logproto.LegacySample, 0, len(stream.Values))
		for _, s := range stream.Values {
			samples = append(samples, logproto.LegacySample{
				Value:       float64(s.Value),
				TimestampMs: int64(s.Timestamp),
			})
		}
		res = append(res, queryrangebase.SampleStream{
			Labels:  logproto.FromMetricsToLabelAdapters(stream.Metric),
			Samples: samples,
		})
	}
	return res
}

func toProtoVector(v loghttp.Vector) []queryrangebase.SampleStream {
	res := make([]queryrangebase.SampleStream, 0, len(v))

	if len(v) == 0 {
		return res
	}
	for _, s := range v {
		res = append(res, queryrangebase.SampleStream{
			Samples: []logproto.LegacySample{{
				Value:       float64(s.Value),
				TimestampMs: int64(s.Timestamp),
			}},
			Labels: logproto.FromMetricsToLabelAdapters(s.Metric),
		})
	}
	return res
}

func (res LogstoreResponse) Count() int64 {
	var result int64
	for _, s := range res.Data.Result {
		result += int64(len(s.Entries))
	}
	return result
}

func paramsFromRequest(req queryrangebase.Request) (logql.Params, error) {
	switch r := req.(type) {
	case *LogstoreRequest:
		return &paramsRangeWrapper{
			LogstoreRequest: r,
		}, nil
	case *LogstoreInstantRequest:
		return &paramsInstantWrapper{
			LogstoreInstantRequest: r,
		}, nil
	case *LogstoreSeriesRequest:
		return &paramsSeriesWrapper{
			LogstoreSeriesRequest: r,
		}, nil
	case *LogstoreLabelNamesRequest:
		return &paramsLabelNamesWrapper{
			LogstoreLabelNamesRequest: r,
		}, nil
	default:
		return nil, fmt.Errorf("expected one of the *LogstoreRequest, *LogstoreInstantRequest, *LogstoreSeriesRequest, *LogstoreLabelNamesRequest, got (%T)", r)
	}
}

type paramsRangeWrapper struct {
	*LogstoreRequest
}

func (p paramsRangeWrapper) Query() string {
	return p.GetQuery()
}

func (p paramsRangeWrapper) Start() time.Time {
	return p.GetStartTs()
}

func (p paramsRangeWrapper) End() time.Time {
	return p.GetEndTs()
}

func (p paramsRangeWrapper) Step() time.Duration {
	return time.Duration(p.GetStep() * 1e6)
}
func (p paramsRangeWrapper) Interval() time.Duration {
	return time.Duration(p.GetInterval() * 1e6)
}
func (p paramsRangeWrapper) Direction() logproto.Direction {
	return p.GetDirection()
}
func (p paramsRangeWrapper) Limit() uint32 { return p.LogstoreRequest.Limit }
func (p paramsRangeWrapper) Shards() []string {
	return p.GetShards()
}

type paramsInstantWrapper struct {
	*LogstoreInstantRequest
}

func (p paramsInstantWrapper) Query() string {
	return p.GetQuery()
}

func (p paramsInstantWrapper) Start() time.Time {
	return p.LogstoreInstantRequest.GetTimeTs()
}

func (p paramsInstantWrapper) End() time.Time {
	return p.LogstoreInstantRequest.GetTimeTs()
}

func (p paramsInstantWrapper) Step() time.Duration {
	return time.Duration(p.GetStep() * 1e6)
}
func (p paramsInstantWrapper) Interval() time.Duration { return 0 }
func (p paramsInstantWrapper) Direction() logproto.Direction {
	return p.GetDirection()
}
func (p paramsInstantWrapper) Limit() uint32 { return p.LogstoreInstantRequest.Limit }
func (p paramsInstantWrapper) Shards() []string {
	return p.GetShards()
}

type paramsSeriesWrapper struct {
	*LogstoreSeriesRequest
}

func (p paramsSeriesWrapper) Query() string {
	return p.GetQuery()
}

func (p paramsSeriesWrapper) Start() time.Time {
	return p.LogstoreSeriesRequest.GetStartTs()
}

func (p paramsSeriesWrapper) End() time.Time {
	return p.LogstoreSeriesRequest.GetEndTs()
}

func (p paramsSeriesWrapper) Step() time.Duration {
	return time.Duration(p.GetStep() * 1e6)
}
func (p paramsSeriesWrapper) Interval() time.Duration { return 0 }
func (p paramsSeriesWrapper) Direction() logproto.Direction {
	return logproto.FORWARD
}
func (p paramsSeriesWrapper) Limit() uint32 { return 0 }
func (p paramsSeriesWrapper) Shards() []string {
	return p.GetShards()
}

type paramsLabelNamesWrapper struct {
	*LogstoreLabelNamesRequest
}

func (p paramsLabelNamesWrapper) Query() string {
	return p.GetQuery()
}

func (p paramsLabelNamesWrapper) Start() time.Time {
	return p.LogstoreLabelNamesRequest.GetStartTs()
}

func (p paramsLabelNamesWrapper) End() time.Time {
	return p.LogstoreLabelNamesRequest.GetEndTs()
}

func (p paramsLabelNamesWrapper) Step() time.Duration {
	return time.Duration(p.GetStep() * 1e6)
}
func (p paramsLabelNamesWrapper) Interval() time.Duration { return 0 }
func (p paramsLabelNamesWrapper) Direction() logproto.Direction {
	return logproto.FORWARD
}
func (p paramsLabelNamesWrapper) Limit() uint32 { return 0 }
func (p paramsLabelNamesWrapper) Shards() []string {
	return make([]string, 0)
}

func httpResponseHeadersToPromResponseHeaders(httpHeaders http.Header) []queryrangebase.PrometheusResponseHeader {
	var promHeaders []queryrangebase.PrometheusResponseHeader
	for h, hv := range httpHeaders {
		promHeaders = append(promHeaders, queryrangebase.PrometheusResponseHeader{Name: h, Values: hv})
	}

	return promHeaders
}

func getQueryTags(ctx context.Context) string {
	v, _ := ctx.Value(httpreq.QueryTagsHTTPHeader).(string) // it's ok to be empty
	return v
}

func NewEmptyResponse(r queryrangebase.Request) (queryrangebase.Response, error) {
	switch req := r.(type) {
	case *LogstoreSeriesRequest:
		return &LogstoreSeriesResponse{
			Status:  loghttp.QueryStatusSuccess,
			Version: uint32(loghttp.GetVersion(req.Path)),
		}, nil
	case *LogstoreLabelNamesRequest:
		return &LogstoreLabelNamesResponse{
			Status:  loghttp.QueryStatusSuccess,
			Version: uint32(loghttp.GetVersion(req.Path)),
		}, nil
	case *LogstoreInstantRequest:
		// instant queries in the frontend are always metrics queries.
		return &LogstorePromResponse{
			Response: &queryrangebase.PrometheusResponse{
				Status: loghttp.QueryStatusSuccess,
				Data: queryrangebase.PrometheusData{
					ResultType: loghttp.ResultTypeVector,
				},
			},
		}, nil
	case *LogstoreRequest:
		// range query can either be metrics or logs
		expr, err := syntax.ParseExpr(req.Query)
		if err != nil {
			return nil, httpgrpc.Errorf(http.StatusBadRequest, err.Error())
		}
		if _, ok := expr.(syntax.SampleExpr); ok {
			return &LogstorePromResponse{
				Response: queryrangebase.NewEmptyPrometheusResponse(),
			}, nil
		}
		return &LogstoreResponse{
			Status:    loghttp.QueryStatusSuccess,
			Direction: req.Direction,
			Limit:     req.Limit,
			Version:   uint32(loghttp.GetVersion(req.Path)),
			Data: LogstoreData{
				ResultType: loghttp.ResultTypeStream,
			},
		}, nil
	default:
		return nil, fmt.Errorf("unsupported request type %T", req)
	}
}

func mergeLogstoreResponse(responses ...queryrangebase.Response) *LogstoreResponse {
	if len(responses) == 0 {
		return nil
	}
	var (
		logstoreRes       = responses[0].(*LogstoreResponse)
		mergedStats   stats.Result
		logstoreResponses = make([]*LogstoreResponse, 0, len(responses))
	)

	for _, res := range responses {
		logstoreResult := res.(*LogstoreResponse)
		mergedStats.Merge(logstoreResult.Statistics)
		logstoreResponses = append(logstoreResponses, logstoreResult)
	}

	return &LogstoreResponse{
		Status:     loghttp.QueryStatusSuccess,
		Direction:  logstoreRes.Direction,
		Limit:      logstoreRes.Limit,
		Version:    logstoreRes.Version,
		ErrorType:  logstoreRes.ErrorType,
		Error:      logstoreRes.Error,
		Statistics: mergedStats,
		Data: LogstoreData{
			ResultType: loghttp.ResultTypeStream,
			Result:     mergeOrderedNonOverlappingStreams(logstoreResponses, logstoreRes.Limit, logstoreRes.Direction),
		},
	}
}
