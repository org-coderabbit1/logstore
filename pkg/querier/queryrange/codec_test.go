package queryrange

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	strings "strings"
	"testing"
	"time"

	"github.com/prometheus/common/model"
	"github.com/stretchr/testify/require"

	"example.com/acme/logstore/pkg/loghttp"
	"example.com/acme/logstore/pkg/logproto"
	"example.com/acme/logstore/pkg/logqlmodel/stats"
	"example.com/acme/logstore/pkg/querier/queryrange/queryrangebase"
	"example.com/acme/logstore/pkg/util"
)

func init() {
	time.Local = nil // for easier tests comparison
}

var (
	start = testTime //  Marshalling the time drops the monotonic clock so we can't use time.Now
	end   = start.Add(1 * time.Hour)
)

func Test_codec_DecodeRequest(t *testing.T) {
	tests := []struct {
		name       string
		reqBuilder func() (*http.Request, error)
		want       queryrangebase.Request
		wantErr    bool
	}{
		{"wrong", func() (*http.Request, error) { return http.NewRequest(http.MethodGet, "/bad?step=bad", nil) }, nil, true},
		{"query_range", func() (*http.Request, error) {
			return http.NewRequest(http.MethodGet,
				fmt.Sprintf(`/query_range?start=%d&end=%d&query={foo="bar"}&step=10&limit=200&direction=FORWARD`, start.UnixNano(), end.UnixNano()), nil)
		}, &LogstoreRequest{
			Query:     `{foo="bar"}`,
			Limit:     200,
			Step:      10000, // step is expected in ms
			Direction: logproto.FORWARD,
			Path:      "/query_range",
			StartTs:   start,
			EndTs:     end,
		}, false},
		{"query_range", func() (*http.Request, error) {
			return http.NewRequest(http.MethodGet,
				fmt.Sprintf(`/query_range?start=%d&end=%d&query={foo="bar"}&interval=10&limit=200&direction=BACKWARD`, start.UnixNano(), end.UnixNano()), nil)
		}, &LogstoreRequest{
			Query:     `{foo="bar"}`,
			Limit:     200,
			Step:      14000, // step is expected in ms; calculated default if request param not present
			Interval:  10000, // interval is expected in ms
			Direction: logproto.BACKWARD,
			Path:      "/query_range",
			StartTs:   start,
			EndTs:     end,
		}, false},
		{"series", func() (*http.Request, error) {
			return http.NewRequest(http.MethodGet,
				fmt.Sprintf(`/series?start=%d&end=%d&match={foo="bar"}`, start.UnixNano(), end.UnixNano()), nil)
		}, &LogstoreSeriesRequest{
			Match:   []string{`{foo="bar"}`},
			Path:    "/series",
			StartTs: start,
			EndTs:   end,
		}, false},
		{"labels", func() (*http.Request, error) {
			return http.NewRequest(http.MethodGet,
				fmt.Sprintf(`/label?start=%d&end=%d`, start.UnixNano(), end.UnixNano()), nil)
		}, &LogstoreLabelNamesRequest{
			Path:    "/label",
			StartTs: start,
			EndTs:   end,
		}, false},
		{"label_values", func() (*http.Request, error) {
			return http.NewRequest(http.MethodGet,
				fmt.Sprintf(`/label/test/values?start=%d&end=%d&query={foo="bar"}`, start.UnixNano(), end.UnixNano()), nil)
		}, &LogstoreLabelNamesRequest{
			Path:    "/label/test/values",
			StartTs: start,
			EndTs:   end,
			Query:   `{foo="bar"}`,
		}, false},
		{"index_stats", func() (*http.Request, error) {
			return LogstoreCodec.EncodeRequest(context.Background(), &logproto.IndexStatsRequest{
				From:     model.TimeFromUnixNano(start.UnixNano()),
				Through:  model.TimeFromUnixNano(end.UnixNano()),
				Matchers: `{job="foo"}`,
			})
		}, &logproto.IndexStatsRequest{
			From:     model.TimeFromUnixNano(start.UnixNano()),
			Through:  model.TimeFromUnixNano(end.UnixNano()),
			Matchers: `{job="foo"}`,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := tt.reqBuilder()
			if err != nil {
				t.Fatal(err)
			}
			got, err := LogstoreCodec.DecodeRequest(context.TODO(), req, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("codec.DecodeRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, got, tt.want)
		})
	}
}

func Test_codec_DecodeResponse(t *testing.T) {
	tests := []struct {
		name    string
		res     *http.Response
		req     queryrangebase.Request
		want    queryrangebase.Response
		wantErr bool
	}{
		{"500", &http.Response{StatusCode: 500, Body: io.NopCloser(strings.NewReader("some error"))}, nil, nil, true},
		{"no body", &http.Response{StatusCode: 200, Body: io.NopCloser(badReader{})}, nil, nil, true},
		{"bad json", &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(""))}, nil, nil, true},
		{"not success", &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"status":"fail"}`))}, nil, nil, true},
		{"unknown", &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"status":"success"}`))}, nil, nil, true},
		{
			"matrix", &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(matrixString))}, nil,
			&LogstorePromResponse{
				Response: &queryrangebase.PrometheusResponse{
					Status: loghttp.QueryStatusSuccess,
					Data: queryrangebase.PrometheusData{
						ResultType: loghttp.ResultTypeMatrix,
						Result:     sampleStreams,
					},
				},
				Statistics: statsResult,
			}, false,
		},
		{
			"matrix-empty-streams",
			&http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(matrixStringEmptyResult))},
			nil,
			&LogstorePromResponse{
				Response: &queryrangebase.PrometheusResponse{
					Status: loghttp.QueryStatusSuccess,
					Data: queryrangebase.PrometheusData{
						ResultType: loghttp.ResultTypeMatrix,
						Result:     make([]queryrangebase.SampleStream, 0), // shouldn't be nil.
					},
				},
				Statistics: statsResult,
			}, false,
		},
		{
			"vector-empty-streams",
			&http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(vectorStringEmptyResult))},
			nil,
			&LogstorePromResponse{
				Response: &queryrangebase.PrometheusResponse{
					Status: loghttp.QueryStatusSuccess,
					Data: queryrangebase.PrometheusData{
						ResultType: loghttp.ResultTypeVector,
						Result:     make([]queryrangebase.SampleStream, 0), // shouldn't be nil.
					},
				},
				Statistics: statsResult,
			}, false,
		},
		{
			"streams v1", &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(streamsString))},
			&LogstoreRequest{Direction: logproto.FORWARD, Limit: 100, Path: "/logstore/api/v1/query_range"},
			&LogstoreResponse{
				Status:    loghttp.QueryStatusSuccess,
				Direction: logproto.FORWARD,
				Limit:     100,
				Version:   uint32(loghttp.VersionV1),
				Data: LogstoreData{
					ResultType: loghttp.ResultTypeStream,
					Result:     logStreams,
				},
				Statistics: statsResult,
			}, false,
		},
		{
			"streams legacy", &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(streamsString))},
			&LogstoreRequest{Direction: logproto.FORWARD, Limit: 100, Path: "/api/prom/query_range"},
			&LogstoreResponse{
				Status:    loghttp.QueryStatusSuccess,
				Direction: logproto.FORWARD,
				Limit:     100,
				Version:   uint32(loghttp.VersionLegacy),
				Data: LogstoreData{
					ResultType: loghttp.ResultTypeStream,
					Result:     logStreams,
				},
				Statistics: statsResult,
			}, false,
		},
		{
			"series", &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(seriesString))},
			&LogstoreSeriesRequest{Path: "/logstore/api/v1/series"},
			&LogstoreSeriesResponse{
				Status:  "success",
				Version: uint32(loghttp.VersionV1),
				Data:    seriesData,
			}, false,
		},
		{
			"labels legacy", &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(labelsString))},
			&LogstoreLabelNamesRequest{Path: "/api/prom/label"},
			&LogstoreLabelNamesResponse{
				Status:  "success",
				Version: uint32(loghttp.VersionLegacy),
				Data:    labelsData,
			}, false,
		},
		{
			"index stats", &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(indexStatsString))},
			&logproto.IndexStatsRequest{},
			&IndexStatsResponse{
				Response: &logproto.IndexStatsResponse{
					Streams: 1,
					Chunks:  2,
					Bytes:   3,
					Entries: 4,
				},
			}, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LogstoreCodec.DecodeResponse(context.TODO(), tt.res, tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("codec.DecodeResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_codec_EncodeRequest(t *testing.T) {
	// we only accept LogstoreRequest.
	got, err := LogstoreCodec.EncodeRequest(context.TODO(), &queryrangebase.PrometheusRequest{})
	require.Error(t, err)
	require.Nil(t, got)

	ctx := context.Background()
	toEncode := &LogstoreRequest{
		Query:     `{foo="bar"}`,
		Limit:     200,
		Step:      86400000, // nanoseconds
		Interval:  10000000, // nanoseconds
		Direction: logproto.FORWARD,
		Path:      "/query_range",
		StartTs:   start,
		EndTs:     end,
	}
	got, err = LogstoreCodec.EncodeRequest(ctx, toEncode)
	require.NoError(t, err)
	require.Equal(t, ctx, got.Context())
	require.Equal(t, "/logstore/api/v1/query_range", got.URL.Path)
	require.Equal(t, fmt.Sprintf("%d", start.UnixNano()), got.URL.Query().Get("start"))
	require.Equal(t, fmt.Sprintf("%d", end.UnixNano()), got.URL.Query().Get("end"))
	require.Equal(t, `{foo="bar"}`, got.URL.Query().Get("query"))
	require.Equal(t, fmt.Sprintf("%d", 200), got.URL.Query().Get("limit"))
	require.Equal(t, `FORWARD`, got.URL.Query().Get("direction"))
	require.Equal(t, "86400.000000", got.URL.Query().Get("step"))
	require.Equal(t, "10000.000000", got.URL.Query().Get("interval"))

	// testing a full roundtrip
	req, err := LogstoreCodec.DecodeRequest(context.TODO(), got, nil)
	require.NoError(t, err)
	require.Equal(t, toEncode.Query, req.(*LogstoreRequest).Query)
	require.Equal(t, toEncode.Step, req.(*LogstoreRequest).Step)
	require.Equal(t, toEncode.Interval, req.(*LogstoreRequest).Interval)
	require.Equal(t, toEncode.StartTs, req.(*LogstoreRequest).StartTs)
	require.Equal(t, toEncode.EndTs, req.(*LogstoreRequest).EndTs)
	require.Equal(t, toEncode.Direction, req.(*LogstoreRequest).Direction)
	require.Equal(t, toEncode.Limit, req.(*LogstoreRequest).Limit)
	require.Equal(t, "/logstore/api/v1/query_range", req.(*LogstoreRequest).Path)
}

func Test_codec_series_EncodeRequest(t *testing.T) {
	got, err := LogstoreCodec.EncodeRequest(context.TODO(), &queryrangebase.PrometheusRequest{})
	require.Error(t, err)
	require.Nil(t, got)

	ctx := context.Background()
	toEncode := &LogstoreSeriesRequest{
		Match:   []string{`{foo="bar"}`},
		Path:    "/series",
		StartTs: start,
		EndTs:   end,
	}
	got, err = LogstoreCodec.EncodeRequest(ctx, toEncode)
	require.NoError(t, err)
	require.Equal(t, ctx, got.Context())
	require.Equal(t, "/logstore/api/v1/series", got.URL.Path)
	require.Equal(t, fmt.Sprintf("%d", start.UnixNano()), got.URL.Query().Get("start"))
	require.Equal(t, fmt.Sprintf("%d", end.UnixNano()), got.URL.Query().Get("end"))
	require.Equal(t, `{foo="bar"}`, got.URL.Query().Get("match[]"))

	// testing a full roundtrip
	req, err := LogstoreCodec.DecodeRequest(context.TODO(), got, nil)
	require.NoError(t, err)
	require.Equal(t, toEncode.Match, req.(*LogstoreSeriesRequest).Match)
	require.Equal(t, toEncode.StartTs, req.(*LogstoreSeriesRequest).StartTs)
	require.Equal(t, toEncode.EndTs, req.(*LogstoreSeriesRequest).EndTs)
	require.Equal(t, "/logstore/api/v1/series", req.(*LogstoreSeriesRequest).Path)
}

func Test_codec_labels_EncodeRequest(t *testing.T) {
	ctx := context.Background()
	toEncode := &LogstoreLabelNamesRequest{
		Path:    "/logstore/api/v1/labels",
		StartTs: start,
		EndTs:   end,
	}
	got, err := LogstoreCodec.EncodeRequest(ctx, toEncode)
	require.NoError(t, err)
	require.Equal(t, ctx, got.Context())
	require.Equal(t, "/logstore/api/v1/labels", got.URL.Path)
	require.Equal(t, fmt.Sprintf("%d", start.UnixNano()), got.URL.Query().Get("start"))
	require.Equal(t, fmt.Sprintf("%d", end.UnixNano()), got.URL.Query().Get("end"))

	// testing a full roundtrip
	req, err := LogstoreCodec.DecodeRequest(context.TODO(), got, nil)
	require.NoError(t, err)
	require.Equal(t, toEncode.StartTs, req.(*LogstoreLabelNamesRequest).StartTs)
	require.Equal(t, toEncode.EndTs, req.(*LogstoreLabelNamesRequest).EndTs)
	require.Equal(t, "/logstore/api/v1/labels", req.(*LogstoreLabelNamesRequest).Path)

	// Test labels values endpoint
	toEncode = &LogstoreLabelNamesRequest{
		Path:    "/logstore/api/v1/labels/__name__/values",
		StartTs: start,
		EndTs:   end,
		Query:   `{foo="bar"}`,
	}
	got, err = LogstoreCodec.EncodeRequest(ctx, toEncode)
	require.NoError(t, err)
	require.Equal(t, ctx, got.Context())
	require.Equal(t, "/logstore/api/v1/labels/__name__/values", got.URL.Path)
	require.Equal(t, fmt.Sprintf("%d", start.UnixNano()), got.URL.Query().Get("start"))
	require.Equal(t, fmt.Sprintf("%d", end.UnixNano()), got.URL.Query().Get("end"))
	require.Equal(t, `{foo="bar"}`, got.URL.Query().Get("query"))

	// testing a full roundtrip
	req, err = LogstoreCodec.DecodeRequest(context.TODO(), got, nil)
	require.NoError(t, err)
	require.Equal(t, toEncode.StartTs, req.(*LogstoreLabelNamesRequest).StartTs)
	require.Equal(t, toEncode.EndTs, req.(*LogstoreLabelNamesRequest).EndTs)
	require.Equal(t, toEncode.Query, req.(*LogstoreLabelNamesRequest).Query)
	require.Equal(t, "/logstore/api/v1/labels/__name__/values", req.(*LogstoreLabelNamesRequest).Path)
}

func Test_codec_labels_DecodeRequest(t *testing.T) {
	ctx := context.Background()
	u, err := url.Parse(`/logstore/api/v1/labels/__name__/values?start=1575285010000000010&end=1575288610000000010&query={foo="bar"}`)
	require.NoError(t, err)

	r := &http.Request{URL: u}
	req, err := LogstoreCodec.DecodeRequest(context.TODO(), r, nil)
	require.NoError(t, err)
	require.Equal(t, start, req.(*LogstoreLabelNamesRequest).StartTs)
	require.Equal(t, end, req.(*LogstoreLabelNamesRequest).EndTs)
	require.Equal(t, `{foo="bar"}`, req.(*LogstoreLabelNamesRequest).Query)
	require.Equal(t, "/logstore/api/v1/labels/__name__/values", req.(*LogstoreLabelNamesRequest).Path)

	got, err := LogstoreCodec.EncodeRequest(ctx, req)
	require.NoError(t, err)
	require.Equal(t, ctx, got.Context())
	require.Equal(t, "/logstore/api/v1/labels/__name__/values", got.URL.Path)
	require.Equal(t, fmt.Sprintf("%d", start.UnixNano()), got.URL.Query().Get("start"))
	require.Equal(t, fmt.Sprintf("%d", end.UnixNano()), got.URL.Query().Get("end"))
	require.Equal(t, `{foo="bar"}`, got.URL.Query().Get("query"))
}

func Test_codec_index_stats_EncodeRequest(t *testing.T) {
	from, through := util.RoundToMilliseconds(start, end)
	toEncode := &logproto.IndexStatsRequest{
		From:     from,
		Through:  through,
		Matchers: `{job="foo"}`,
	}
	got, err := LogstoreCodec.EncodeRequest(context.Background(), toEncode)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%d", from.UnixNano()), got.URL.Query().Get("start"))
	require.Equal(t, fmt.Sprintf("%d", through.UnixNano()), got.URL.Query().Get("end"))
	require.Equal(t, `{job="foo"}`, got.URL.Query().Get("query"))
}

func Test_codec_EncodeResponse(t *testing.T) {
	tests := []struct {
		name    string
		res     queryrangebase.Response
		body    string
		wantErr bool
	}{
		{"error", &badResponse{}, "", true},
		{"prom", &LogstorePromResponse{
			Response: &queryrangebase.PrometheusResponse{
				Status: loghttp.QueryStatusSuccess,
				Data: queryrangebase.PrometheusData{
					ResultType: loghttp.ResultTypeMatrix,
					Result:     sampleStreams,
				},
			},
			Statistics: statsResult,
		}, matrixString, false},
		{
			"logstore v1",
			&LogstoreResponse{
				Status:    loghttp.QueryStatusSuccess,
				Direction: logproto.FORWARD,
				Limit:     100,
				Version:   uint32(loghttp.VersionV1),
				Data: LogstoreData{
					ResultType: loghttp.ResultTypeStream,
					Result:     logStreams,
				},
				Statistics: statsResult,
			}, streamsString, false,
		},
		{
			"logstore legacy",
			&LogstoreResponse{
				Status:    loghttp.QueryStatusSuccess,
				Direction: logproto.FORWARD,
				Limit:     100,
				Version:   uint32(loghttp.VersionLegacy),
				Data: LogstoreData{
					ResultType: loghttp.ResultTypeStream,
					Result:     logStreams,
				},
				Statistics: statsResult,
			}, streamsStringLegacy, false,
		},
		{
			"logstore series",
			&LogstoreSeriesResponse{
				Status:  "success",
				Version: uint32(loghttp.VersionV1),
				Data:    seriesData,
			}, seriesString, false,
		},
		{
			"logstore labels",
			&LogstoreLabelNamesResponse{
				Status:  "success",
				Version: uint32(loghttp.VersionV1),
				Data:    labelsData,
			}, labelsString, false,
		},
		{
			"logstore labels legacy",
			&LogstoreLabelNamesResponse{
				Status:  "success",
				Version: uint32(loghttp.VersionLegacy),
				Data:    labelsData,
			}, labelsLegacyString, false,
		},
		{
			"index stats",
			&IndexStatsResponse{
				Response: &logproto.IndexStatsResponse{
					Streams: 1,
					Chunks:  2,
					Bytes:   3,
					Entries: 4,
				},
			}, indexStatsString, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LogstoreCodec.EncodeResponse(context.TODO(), tt.res)
			if (err != nil) != tt.wantErr {
				t.Errorf("codec.EncodeResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				require.Equal(t, 200, got.StatusCode)
				body, err := io.ReadAll(got.Body)
				require.Nil(t, err)
				bodyString := string(body)
				require.JSONEq(t, tt.body, bodyString)
			}
		})
	}
}

func Test_codec_MergeResponse(t *testing.T) {
	tests := []struct {
		name      string
		responses []queryrangebase.Response
		want      queryrangebase.Response
		wantErr   bool
	}{
		{"empty", []queryrangebase.Response{}, nil, true},
		{"unknown response", []queryrangebase.Response{&badResponse{}}, nil, true},
		{
			"prom",
			[]queryrangebase.Response{
				&LogstorePromResponse{
					Response: &queryrangebase.PrometheusResponse{
						Status: loghttp.QueryStatusSuccess,
						Data: queryrangebase.PrometheusData{
							ResultType: loghttp.ResultTypeMatrix,
							Result:     sampleStreams,
						},
					},
				},
			},
			&LogstorePromResponse{
				Statistics: stats.Result{Summary: stats.Summary{Splits: 1}},
				Response: &queryrangebase.PrometheusResponse{
					Status: loghttp.QueryStatusSuccess,
					Data: queryrangebase.PrometheusData{
						ResultType: loghttp.ResultTypeMatrix,
						Result:     sampleStreams,
					},
				},
			},
			false,
		},
		{
			"logstore backward",
			[]queryrangebase.Response{
				&LogstoreResponse{
					Status:    loghttp.QueryStatusSuccess,
					Direction: logproto.BACKWARD,
					Limit:     100,
					Version:   1,
					Data: LogstoreData{
						ResultType: loghttp.ResultTypeStream,
						Result: []logproto.Stream{
							{
								Labels: `{foo="bar", level="error"}`,
								Entries: []logproto.Entry{
									{Timestamp: time.Unix(0, 2), Line: "2"},
									{Timestamp: time.Unix(0, 1), Line: "1"},
								},
							},
							{
								Labels: `{foo="bar", level="debug"}`,
								Entries: []logproto.Entry{
									{Timestamp: time.Unix(0, 6), Line: "6"},
									{Timestamp: time.Unix(0, 5), Line: "5"},
								},
							},
						},
					},
				},
				&LogstoreResponse{
					Status:    loghttp.QueryStatusSuccess,
					Direction: logproto.BACKWARD,
					Limit:     100,
					Version:   1,
					Data: LogstoreData{
						ResultType: loghttp.ResultTypeStream,
						Result: []logproto.Stream{
							{
								Labels: `{foo="bar", level="error"}`,
								Entries: []logproto.Entry{
									{Timestamp: time.Unix(0, 10), Line: "10"},
									{Timestamp: time.Unix(0, 9), Line: "9"},
									{Timestamp: time.Unix(0, 9), Line: "9"},
								},
							},
							{
								Labels: `{foo="bar", level="debug"}`,
								Entries: []logproto.Entry{
									{Timestamp: time.Unix(0, 16), Line: "16"},
									{Timestamp: time.Unix(0, 15), Line: "15"},
								},
							},
						},
					},
				},
			},
			&LogstoreResponse{
				Status:     loghttp.QueryStatusSuccess,
				Direction:  logproto.BACKWARD,
				Limit:      100,
				Version:    1,
				Statistics: stats.Result{Summary: stats.Summary{Splits: 2}},
				Data: LogstoreData{
					ResultType: loghttp.ResultTypeStream,
					Result: []logproto.Stream{
						{
							Labels: `{foo="bar", level="error"}`,
							Entries: []logproto.Entry{
								{Timestamp: time.Unix(0, 10), Line: "10"},
								{Timestamp: time.Unix(0, 9), Line: "9"},
								{Timestamp: time.Unix(0, 9), Line: "9"},
								{Timestamp: time.Unix(0, 2), Line: "2"},
								{Timestamp: time.Unix(0, 1), Line: "1"},
							},
						},
						{
							Labels: `{foo="bar", level="debug"}`,
							Entries: []logproto.Entry{
								{Timestamp: time.Unix(0, 16), Line: "16"},
								{Timestamp: time.Unix(0, 15), Line: "15"},
								{Timestamp: time.Unix(0, 6), Line: "6"},
								{Timestamp: time.Unix(0, 5), Line: "5"},
							},
						},
					},
				},
			},
			false,
		},
		{
			"logstore backward limited",
			[]queryrangebase.Response{
				&LogstoreResponse{
					Status:    loghttp.QueryStatusSuccess,
					Direction: logproto.BACKWARD,
					Limit:     6,
					Version:   1,
					Data: LogstoreData{
						ResultType: loghttp.ResultTypeStream,
						Result: []logproto.Stream{
							{
								Labels: `{foo="bar", level="error"}`,
								Entries: []logproto.Entry{
									{Timestamp: time.Unix(0, 10), Line: "10"},
									{Timestamp: time.Unix(0, 9), Line: "9"},
									{Timestamp: time.Unix(0, 9), Line: "9"},
								},
							},
							{
								Labels: `{foo="bar", level="debug"}`,
								Entries: []logproto.Entry{
									{Timestamp: time.Unix(0, 16), Line: "16"},
									{Timestamp: time.Unix(0, 15), Line: "15"},
								},
							},
						},
					},
				},
				&LogstoreResponse{
					Status:    loghttp.QueryStatusSuccess,
					Direction: logproto.BACKWARD,
					Limit:     6,
					Version:   1,
					Data: LogstoreData{
						ResultType: loghttp.ResultTypeStream,
						Result: []logproto.Stream{
							{
								Labels: `{foo="bar", level="error"}`,
								Entries: []logproto.Entry{
									{Timestamp: time.Unix(0, 2), Line: "2"},
									{Timestamp: time.Unix(0, 1), Line: "1"},
								},
							},
							{
								Labels: `{foo="bar", level="debug"}`,
								Entries: []logproto.Entry{
									{Timestamp: time.Unix(0, 6), Line: "6"},
									{Timestamp: time.Unix(0, 5), Line: "5"},
								},
							},
						},
					},
				},
			},
			&LogstoreResponse{
				Status:     loghttp.QueryStatusSuccess,
				Direction:  logproto.BACKWARD,
				Limit:      6,
				Version:    1,
				Statistics: stats.Result{Summary: stats.Summary{Splits: 2}},
				Data: LogstoreData{
					ResultType: loghttp.ResultTypeStream,
					Result: []logproto.Stream{
						{
							Labels: `{foo="bar", level="error"}`,
							Entries: []logproto.Entry{
								{Timestamp: time.Unix(0, 10), Line: "10"},
								{Timestamp: time.Unix(0, 9), Line: "9"},
								{Timestamp: time.Unix(0, 9), Line: "9"},
							},
						},
						{
							Labels: `{foo="bar", level="debug"}`,
							Entries: []logproto.Entry{
								{Timestamp: time.Unix(0, 16), Line: "16"},
								{Timestamp: time.Unix(0, 15), Line: "15"},
								{Timestamp: time.Unix(0, 6), Line: "6"},
							},
						},
					},
				},
			},
			false,
		},
		{
			"logstore forward",
			[]queryrangebase.Response{
				&LogstoreResponse{
					Status:    loghttp.QueryStatusSuccess,
					Direction: logproto.FORWARD,
					Limit:     100,
					Version:   1,
					Data: LogstoreData{
						ResultType: loghttp.ResultTypeStream,
						Result: []logproto.Stream{
							{
								Labels: `{foo="bar", level="error"}`,
								Entries: []logproto.Entry{
									{Timestamp: time.Unix(0, 1), Line: "1"},
									{Timestamp: time.Unix(0, 2), Line: "2"},
								},
							},
							{
								Labels: `{foo="bar", level="debug"}`,
								Entries: []logproto.Entry{
									{Timestamp: time.Unix(0, 5), Line: "5"},
									{Timestamp: time.Unix(0, 6), Line: "6"},
								},
							},
						},
					},
				},
				&LogstoreResponse{
					Status:    loghttp.QueryStatusSuccess,
					Direction: logproto.FORWARD,
					Limit:     100,
					Version:   1,
					Data: LogstoreData{
						ResultType: loghttp.ResultTypeStream,
						Result: []logproto.Stream{
							{
								Labels: `{foo="bar", level="error"}`,
								Entries: []logproto.Entry{
									{Timestamp: time.Unix(0, 9), Line: "9"},
									{Timestamp: time.Unix(0, 10), Line: "10"},
								},
							},
							{
								Labels: `{foo="bar", level="debug"}`,
								Entries: []logproto.Entry{
									{Timestamp: time.Unix(0, 15), Line: "15"},
									{Timestamp: time.Unix(0, 15), Line: "15"},
									{Timestamp: time.Unix(0, 16), Line: "16"},
								},
							},
						},
					},
				},
			},
			&LogstoreResponse{
				Status:     loghttp.QueryStatusSuccess,
				Direction:  logproto.FORWARD,
				Limit:      100,
				Version:    1,
				Statistics: stats.Result{Summary: stats.Summary{Splits: 2}},
				Data: LogstoreData{
					ResultType: loghttp.ResultTypeStream,
					Result: []logproto.Stream{
						{
							Labels: `{foo="bar", level="debug"}`,
							Entries: []logproto.Entry{
								{Timestamp: time.Unix(0, 5), Line: "5"},
								{Timestamp: time.Unix(0, 6), Line: "6"},
								{Timestamp: time.Unix(0, 15), Line: "15"},
								{Timestamp: time.Unix(0, 15), Line: "15"},
								{Timestamp: time.Unix(0, 16), Line: "16"},
							},
						},
						{
							Labels: `{foo="bar", level="error"}`,
							Entries: []logproto.Entry{
								{Timestamp: time.Unix(0, 1), Line: "1"},
								{Timestamp: time.Unix(0, 2), Line: "2"},
								{Timestamp: time.Unix(0, 9), Line: "9"},
								{Timestamp: time.Unix(0, 10), Line: "10"},
							},
						},
					},
				},
			},
			false,
		},
		{
			"logstore forward limited",
			[]queryrangebase.Response{
				&LogstoreResponse{
					Status:    loghttp.QueryStatusSuccess,
					Direction: logproto.FORWARD,
					Limit:     5,
					Version:   1,
					Data: LogstoreData{
						ResultType: loghttp.ResultTypeStream,
						Result: []logproto.Stream{
							{
								Labels: `{foo="bar", level="error"}`,
								Entries: []logproto.Entry{
									{Timestamp: time.Unix(0, 1), Line: "1"},
									{Timestamp: time.Unix(0, 2), Line: "2"},
								},
							},
							{
								Labels: `{foo="bar", level="debug"}`,
								Entries: []logproto.Entry{
									{Timestamp: time.Unix(0, 5), Line: "5"},
									{Timestamp: time.Unix(0, 6), Line: "6"},
								},
							},
						},
					},
				},
				&LogstoreResponse{
					Status:    loghttp.QueryStatusSuccess,
					Direction: logproto.FORWARD,
					Limit:     5,
					Version:   1,
					Data: LogstoreData{
						ResultType: loghttp.ResultTypeStream,
						Result: []logproto.Stream{
							{
								Labels: `{foo="bar", level="error"}`,
								Entries: []logproto.Entry{
									{Timestamp: time.Unix(0, 9), Line: "9"},
									{Timestamp: time.Unix(0, 10), Line: "10"},
								},
							},
							{
								Labels: `{foo="bar", level="debug"}`,
								Entries: []logproto.Entry{
									{Timestamp: time.Unix(0, 15), Line: "15"},
									{Timestamp: time.Unix(0, 15), Line: "15"},
									{Timestamp: time.Unix(0, 16), Line: "16"},
								},
							},
						},
					},
				},
			},
			&LogstoreResponse{
				Status:     loghttp.QueryStatusSuccess,
				Direction:  logproto.FORWARD,
				Limit:      5,
				Version:    1,
				Statistics: stats.Result{Summary: stats.Summary{Splits: 2}},
				Data: LogstoreData{
					ResultType: loghttp.ResultTypeStream,
					Result: []logproto.Stream{
						{
							Labels: `{foo="bar", level="debug"}`,
							Entries: []logproto.Entry{
								{Timestamp: time.Unix(0, 5), Line: "5"},
								{Timestamp: time.Unix(0, 6), Line: "6"},
							},
						},
						{
							Labels: `{foo="bar", level="error"}`,
							Entries: []logproto.Entry{
								{Timestamp: time.Unix(0, 1), Line: "1"},
								{Timestamp: time.Unix(0, 2), Line: "2"},
								{Timestamp: time.Unix(0, 9), Line: "9"},
							},
						},
					},
				},
			},
			false,
		},
		{
			"logstore series",
			[]queryrangebase.Response{
				&LogstoreSeriesResponse{
					Status:  "success",
					Version: 1,
					Data: []logproto.SeriesIdentifier{
						{
							Labels: map[string]string{"filename": "/var/hostlog/apport.log", "job": "varlogs"},
						},
						{
							Labels: map[string]string{"filename": "/var/hostlog/test.log", "job": "varlogs"},
						},
					},
				},
				&LogstoreSeriesResponse{
					Status:  "success",
					Version: 1,
					Data: []logproto.SeriesIdentifier{
						{
							Labels: map[string]string{"filename": "/var/hostlog/apport.log", "job": "varlogs"},
						},
						{
							Labels: map[string]string{"filename": "/var/hostlog/other.log", "job": "varlogs"},
						},
					},
				},
			},
			&LogstoreSeriesResponse{
				Statistics: stats.Result{Summary: stats.Summary{Splits: 2}},
				Status:     "success",
				Version:    1,
				Data: []logproto.SeriesIdentifier{
					{
						Labels: map[string]string{"filename": "/var/hostlog/apport.log", "job": "varlogs"},
					},
					{
						Labels: map[string]string{"filename": "/var/hostlog/test.log", "job": "varlogs"},
					},
					{
						Labels: map[string]string{"filename": "/var/hostlog/other.log", "job": "varlogs"},
					},
				},
			},
			false,
		},
		{
			"logstore labels",
			[]queryrangebase.Response{
				&LogstoreLabelNamesResponse{
					Status:  "success",
					Version: 1,
					Data:    []string{"foo", "bar", "buzz"},
				},
				&LogstoreLabelNamesResponse{
					Status:  "success",
					Version: 1,
					Data:    []string{"foo", "bar", "buzz"},
				},
				&LogstoreLabelNamesResponse{
					Status:  "success",
					Version: 1,
					Data:    []string{"foo", "blip", "blop"},
				},
			},
			&LogstoreLabelNamesResponse{
				Statistics: stats.Result{Summary: stats.Summary{Splits: 3}},
				Status:     "success",
				Version:    1,
				Data:       []string{"foo", "bar", "buzz", "blip", "blop"},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LogstoreCodec.MergeResponse(tt.responses...)
			if (err != nil) != tt.wantErr {
				t.Errorf("codec.MergeResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}

type badResponse struct{}

func (badResponse) Reset()                                                 {}
func (badResponse) String() string                                         { return "noop" }
func (badResponse) ProtoMessage()                                          {}
func (badResponse) GetHeaders() []*queryrangebase.PrometheusResponseHeader { return nil }

type badReader struct{}

func (badReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("")
}

var (
	statsResultString = `"stats" : {
		"ingester" : {
			"store": {
				"chunk":{
					"compressedBytes": 1,
					"decompressedBytes": 2,
					"decompressedLines": 3,
					"headChunkBytes": 4,
					"headChunkLines": 5,
					"totalDuplicates": 8
				},
				"chunksDownloadTime": 0,
				"totalChunksRef": 0,
				"totalChunksDownloaded": 0
			},
			"totalBatches": 6,
			"totalChunksMatched": 7,
			"totalLinesSent": 9,
			"totalReached": 10
		},
		"querier": {
			"store" : {
				"chunk": {
					"compressedBytes": 11,
					"decompressedBytes": 12,
					"decompressedLines": 13,
					"headChunkBytes": 14,
					"headChunkLines": 15,
					"totalDuplicates": 19
				},
				"chunksDownloadTime": 16,
				"totalChunksRef": 17,
				"totalChunksDownloaded": 18
			}
		},
		"cache": {
			"chunk": {
				"entriesFound": 0,
				"entriesRequested": 0,
				"entriesStored": 0,
				"bytesReceived": 0,
				"bytesSent": 0,
				"requests": 0,
				"downloadTime": 0
			},
			"index": {
				"entriesFound": 0,
				"entriesRequested": 0,
				"entriesStored": 0,
				"bytesReceived": 0,
				"bytesSent": 0,
				"requests": 0,
				"downloadTime": 0
			},
			"result": {
				"entriesFound": 0,
				"entriesRequested": 0,
				"entriesStored": 0,
				"bytesReceived": 0,
				"bytesSent": 0,
				"requests": 0,
				"downloadTime": 0
			}
		},
		"summary": {
			"bytesProcessedPerSecond": 20,
			"execTime": 22,
			"linesProcessedPerSecond": 23,
			"queueTime": 21,
			"shards": 0,
			"splits": 0,
			"subqueries": 0,
			"totalBytesProcessed": 24,
			"totalEntriesReturned": 10,
			"totalLinesProcessed": 25
		}
	},`
	matrixString = `{
	"data": {
	  ` + statsResultString + `
	  "resultType": "matrix",
	  "result": [
		{
		  "metric": {
			"filename": "\/var\/hostlog\/apport.log",
			"job": "varlogs"
		  },
		  "values": [
			  [
				1568404331.324,
				"0.013333333333333334"
			  ]
			]
		},
		{
		  "metric": {
			"filename": "\/var\/hostlog\/syslog",
			"job": "varlogs"
		  },
		  "values": [
				[
					1568404331.324,
					"3.45"
				],
				[
					1568404331.339,
					"4.45"
				]
			]
		}
	  ]
	},
	"status": "success"
  }`
	matrixStringEmptyResult = `{
	"data": {
	  ` + statsResultString + `
	  "resultType": "matrix",
	  "result": []
	},
	"status": "success"
  }`
	vectorStringEmptyResult = `{
	"data": {
	  ` + statsResultString + `
	  "resultType": "vector",
	  "result": []
	},
	"status": "success"
  }`

	sampleStreams = []queryrangebase.SampleStream{
		{
			Labels:  []logproto.LabelAdapter{{Name: "filename", Value: "/var/hostlog/apport.log"}, {Name: "job", Value: "varlogs"}},
			Samples: []logproto.LegacySample{{Value: 0.013333333333333334, TimestampMs: 1568404331324}},
		},
		{
			Labels:  []logproto.LabelAdapter{{Name: "filename", Value: "/var/hostlog/syslog"}, {Name: "job", Value: "varlogs"}},
			Samples: []logproto.LegacySample{{Value: 3.45, TimestampMs: 1568404331324}, {Value: 4.45, TimestampMs: 1568404331339}},
		},
	}
	streamsString = `{
		"status": "success",
		"data": {
			` + statsResultString + `
			"resultType": "streams",
			"result": [
				{
					"stream": {
						"test": "test"
					},
					"values":[
						[ "123456789012345", "super line" ]
					]
				},
				{
					"stream": {
						"test": "test2"
					},
					"values":[
						[ "123456789012346", "super line2" ]
					]
				}
			]
		}
	}`
	streamsStringLegacy = `{
		` + statsResultString + `"streams":[{"labels":"{test=\"test\"}","entries":[{"ts":"1970-01-02T10:17:36.789012345Z","line":"super line"}]},{"labels":"{test=\"test2\"}","entries":[{"ts":"1970-01-02T10:17:36.789012346Z","line":"super line2"}]}]}`
	logStreams = []logproto.Stream{
		{
			Labels: `{test="test"}`,
			Entries: []logproto.Entry{
				{
					Line:      "super line",
					Timestamp: time.Unix(0, 123456789012345).UTC(),
				},
			},
		},
		{
			Labels: `{test="test2"}`,
			Entries: []logproto.Entry{
				{
					Line:      "super line2",
					Timestamp: time.Unix(0, 123456789012346).UTC(),
				},
			},
		},
	}
	seriesString = `{
		"status": "success",
		"data": [
			{"filename": "/var/hostlog/apport.log", "job": "varlogs"},
			{"filename": "/var/hostlog/test.log", "job": "varlogs"}
		]
	}`
	seriesData = []logproto.SeriesIdentifier{
		{
			Labels: map[string]string{"filename": "/var/hostlog/apport.log", "job": "varlogs"},
		},
		{
			Labels: map[string]string{"filename": "/var/hostlog/test.log", "job": "varlogs"},
		},
	}
	labelsString = `{
		"status": "success",
		"data": [
			"foo",
			"bar"
		]
	}`
	labelsLegacyString = `{
		"values": [
			"foo",
			"bar"
		]
	}`
	indexStatsString = `{
		"streams": 1,
		"chunks": 2,
		"bytes": 3,
		"entries": 4
		}`
	labelsData  = []string{"foo", "bar"}
	statsResult = stats.Result{
		Summary: stats.Summary{
			BytesProcessedPerSecond: 20,
			QueueTime:               21,
			ExecTime:                22,
			LinesProcessedPerSecond: 23,
			TotalBytesProcessed:     24,
			TotalLinesProcessed:     25,
			TotalEntriesReturned:    10,
		},
		Querier: stats.Querier{
			Store: stats.Store{
				Chunk: stats.Chunk{
					CompressedBytes:   11,
					DecompressedBytes: 12,
					DecompressedLines: 13,
					HeadChunkBytes:    14,
					HeadChunkLines:    15,
					TotalDuplicates:   19,
				},
				ChunksDownloadTime:    16,
				TotalChunksRef:        17,
				TotalChunksDownloaded: 18,
			},
		},

		Ingester: stats.Ingester{
			Store: stats.Store{
				Chunk: stats.Chunk{
					CompressedBytes:   1,
					DecompressedBytes: 2,
					DecompressedLines: 3,
					HeadChunkBytes:    4,
					HeadChunkLines:    5,
					TotalDuplicates:   8,
				},
			},
			TotalBatches:       6,
			TotalChunksMatched: 7,
			TotalLinesSent:     9,
			TotalReached:       10,
		},

		Caches: stats.Caches{
			Chunk:  stats.Cache{},
			Index:  stats.Cache{},
			Result: stats.Cache{},
		},
	}
)

func BenchmarkResponseMerge(b *testing.B) {
	const (
		resps         = 10
		streams       = 100
		logsPerStream = 1000
	)

	for _, tc := range []struct {
		desc  string
		limit uint32
		fn    func([]*LogstoreResponse, uint32, logproto.Direction) []logproto.Stream
	}{
		{
			"mergeStreams unlimited",
			uint32(streams * logsPerStream),
			mergeStreams,
		},
		{
			"mergeOrderedNonOverlappingStreams unlimited",
			uint32(streams * logsPerStream),
			mergeOrderedNonOverlappingStreams,
		},
		{
			"mergeStreams limited",
			uint32(streams*logsPerStream - 1),
			mergeStreams,
		},
		{
			"mergeOrderedNonOverlappingStreams limited",
			uint32(streams*logsPerStream - 1),
			mergeOrderedNonOverlappingStreams,
		},
	} {
		input := mkResps(resps, streams, logsPerStream, logproto.FORWARD)
		b.Run(tc.desc, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tc.fn(input, tc.limit, logproto.FORWARD)
			}
		})
	}
}

func mkResps(nResps, nStreams, nLogs int, direction logproto.Direction) (resps []*LogstoreResponse) {
	for i := 0; i < nResps; i++ {
		r := &LogstoreResponse{}
		for j := 0; j < nStreams; j++ {
			stream := logproto.Stream{
				Labels: fmt.Sprintf(`{foo="%d"}`, j),
			}
			// split nLogs evenly across all responses
			for k := i * (nLogs / nResps); k < (i+1)*(nLogs/nResps); k++ {
				stream.Entries = append(stream.Entries, logproto.Entry{
					Timestamp: time.Unix(int64(k), 0),
					Line:      fmt.Sprintf("%d", k),
				})

				if direction == logproto.BACKWARD {
					for x, y := 0, len(stream.Entries)-1; x < len(stream.Entries)/2; x, y = x+1, y-1 {
						stream.Entries[x], stream.Entries[y] = stream.Entries[y], stream.Entries[x]
					}
				}
			}
			r.Data.Result = append(r.Data.Result, stream)
		}
		resps = append(resps, r)
	}
	return resps
}

type buffer struct {
	buff []byte
	io.ReadCloser
}

func (b *buffer) Bytes() []byte {
	return b.buff
}

func Benchmark_CodecDecodeLogs(b *testing.B) {
	ctx := context.Background()
	resp, err := LogstoreCodec.EncodeResponse(ctx, &LogstoreResponse{
		Status:    loghttp.QueryStatusSuccess,
		Direction: logproto.BACKWARD,
		Version:   uint32(loghttp.VersionV1),
		Limit:     1000,
		Data: LogstoreData{
			ResultType: loghttp.ResultTypeStream,
			Result:     generateStream(),
		},
	})
	require.Nil(b, err)

	buf, err := io.ReadAll(resp.Body)
	require.Nil(b, err)
	reader := bytes.NewReader(buf)
	resp.Body = &buffer{
		ReadCloser: io.NopCloser(reader),
		buff:       buf,
	}
	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		_, _ = reader.Seek(0, io.SeekStart)
		result, err := LogstoreCodec.DecodeResponse(ctx, resp, &LogstoreRequest{
			Limit:     100,
			StartTs:   start,
			EndTs:     end,
			Direction: logproto.BACKWARD,
			Path:      "/logstore/api/v1/query_range",
		})
		require.Nil(b, err)
		require.NotNil(b, result)
	}
}

func Benchmark_CodecDecodeSamples(b *testing.B) {
	ctx := context.Background()
	resp, err := LogstoreCodec.EncodeResponse(ctx, &LogstorePromResponse{
		Response: &queryrangebase.PrometheusResponse{
			Status: loghttp.QueryStatusSuccess,
			Data: queryrangebase.PrometheusData{
				ResultType: loghttp.ResultTypeMatrix,
				Result:     generateMatrix(),
			},
		},
	})
	require.Nil(b, err)

	buf, err := io.ReadAll(resp.Body)
	require.Nil(b, err)
	reader := bytes.NewReader(buf)
	resp.Body = io.NopCloser(reader)
	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		_, _ = reader.Seek(0, io.SeekStart)
		result, err := LogstoreCodec.DecodeResponse(ctx, resp, &LogstoreRequest{
			Limit:     100,
			StartTs:   start,
			EndTs:     end,
			Direction: logproto.BACKWARD,
			Path:      "/logstore/api/v1/query_range",
		})
		require.Nil(b, err)
		require.NotNil(b, result)
	}
}

func generateMatrix() (res []queryrangebase.SampleStream) {
	for i := 0; i < 100; i++ {
		s := queryrangebase.SampleStream{
			Labels:  []logproto.LabelAdapter{},
			Samples: []logproto.LegacySample{},
		}
		for j := 0; j < 1000; j++ {
			s.Samples = append(s.Samples, logproto.LegacySample{
				Value:       float64(j),
				TimestampMs: int64(j),
			})
		}
		res = append(res, s)
	}
	return res
}

func generateStream() (res []logproto.Stream) {
	for i := 0; i < 1000; i++ {
		s := logproto.Stream{
			Labels: fmt.Sprintf(`{foo="%d", buzz="bar", cluster="us-central2", namespace="logstore-dev", container="query-frontend"}`, i),
		}
		for j := 0; j < 10; j++ {
			s.Entries = append(s.Entries, logproto.Entry{Timestamp: time.Now(), Line: fmt.Sprintf("%d\nyolo", j)})
		}
		res = append(res, s)
	}
	return res
}
