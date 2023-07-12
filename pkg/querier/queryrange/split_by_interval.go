package queryrange

import (
	"context"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	otlog "github.com/opentracing/opentracing-go/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/common/model"
	"github.com/weaveworks/common/httpgrpc"

	"example.com/acme/logstore/pkg/util/math"

	"example.com/acme/kit/tenant"

	"example.com/acme/logstore/pkg/logproto"
	"example.com/acme/logstore/pkg/logql/syntax"
	"example.com/acme/logstore/pkg/querier/queryrange/queryrangebase"
	"example.com/acme/logstore/pkg/storage/config"
	"example.com/acme/logstore/pkg/util"
	"example.com/acme/logstore/pkg/util/validation"
)

type logstoreResult struct {
	req queryrangebase.Request
	ch  chan *packedResp
}

type packedResp struct {
	resp queryrangebase.Response
	err  error
}

type SplitByMetrics struct {
	splits prometheus.Histogram
}

func NewSplitByMetrics(r prometheus.Registerer) *SplitByMetrics {
	return &SplitByMetrics{
		splits: promauto.With(r).NewHistogram(prometheus.HistogramOpts{
			Namespace: "logstore",
			Name:      "query_frontend_partitions",
			Help:      "Number of time-based partitions (sub-requests) per request",
			Buckets:   prometheus.ExponentialBuckets(1, 4, 5), // 1 -> 1024
		}),
	}
}

type splitByInterval struct {
	configs  []config.PeriodConfig
	next     queryrangebase.Handler
	limits   Limits
	merger   queryrangebase.Merger
	metrics  *SplitByMetrics
	splitter Splitter
}

type Splitter func(req queryrangebase.Request, interval time.Duration) ([]queryrangebase.Request, error)

// SplitByIntervalMiddleware creates a new Middleware that splits log requests by a given interval.
func SplitByIntervalMiddleware(configs []config.PeriodConfig, limits Limits, merger queryrangebase.Merger, splitter Splitter, metrics *SplitByMetrics) queryrangebase.Middleware {
	if metrics == nil {
		metrics = NewSplitByMetrics(nil)
	}

	return queryrangebase.MiddlewareFunc(func(next queryrangebase.Handler) queryrangebase.Handler {
		return &splitByInterval{
			configs:  configs,
			next:     next,
			limits:   limits,
			merger:   merger,
			metrics:  metrics,
			splitter: splitter,
		}
	})
}

func (h *splitByInterval) Feed(ctx context.Context, input []*logstoreResult) chan *logstoreResult {
	ch := make(chan *logstoreResult)

	go func() {
		defer close(ch)
		for _, d := range input {
			select {
			case <-ctx.Done():
				return
			case ch <- d:
				continue
			}
		}
	}()

	return ch
}

func (h *splitByInterval) Process(
	ctx context.Context,
	parallelism int,
	threshold int64,
	input []*logstoreResult,
	maxSeries int,
) ([]queryrangebase.Response, error) {
	var responses []queryrangebase.Response
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ch := h.Feed(ctx, input)

	// queries with 0 limits should not be exited early
	var unlimited bool
	if threshold == 0 {
		unlimited = true
	}

	// Parallelism will be at least 1
	p := math.Max(parallelism, 1)
	// don't spawn unnecessary goroutines
	if len(input) < parallelism {
		p = len(input)
	}

	// per request wrapped handler for limiting the amount of series.
	next := newSeriesLimiter(maxSeries).Wrap(h.next)
	for i := 0; i < p; i++ {
		go h.loop(ctx, ch, next)
	}

	for _, x := range input {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case data := <-x.ch:
			if data.err != nil {
				return nil, data.err
			}

			responses = append(responses, data.resp)

			// see if we can exit early if a limit has been reached
			if casted, ok := data.resp.(*LogstoreResponse); !unlimited && ok {
				threshold -= casted.Count()

				if threshold <= 0 {
					return responses, nil
				}

			}

		}
	}

	return responses, nil
}

func (h *splitByInterval) loop(ctx context.Context, ch <-chan *logstoreResult, next queryrangebase.Handler) {
	for data := range ch {

		sp, ctx := opentracing.StartSpanFromContext(ctx, "interval")
		data.req.LogToSpan(sp)

		resp, err := next.Do(ctx, data.req)
		sp.Finish()

		select {
		case <-ctx.Done():
			return
		case data.ch <- &packedResp{resp, err}:
		}
	}
}

func (h *splitByInterval) Do(ctx context.Context, r queryrangebase.Request) (queryrangebase.Response, error) {
	tenantIDs, err := tenant.TenantIDs(ctx)
	if err != nil {
		return nil, httpgrpc.Errorf(http.StatusBadRequest, err.Error())
	}

	interval := validation.MaxDurationOrZeroPerTenant(tenantIDs, h.limits.QuerySplitDuration)
	// skip split by if unset
	if interval == 0 {
		return h.next.Do(ctx, r)
	}

	intervals, err := h.splitter(r, interval)
	if err != nil {
		return nil, err
	}

	h.metrics.splits.Observe(float64(len(intervals)))

	// no interval should not be processed by the frontend.
	if len(intervals) == 0 {
		return h.next.Do(ctx, r)
	}

	if sp := opentracing.SpanFromContext(ctx); sp != nil {
		sp.LogFields(otlog.Int("n_intervals", len(intervals)))
	}

	if len(intervals) == 1 {
		return h.next.Do(ctx, intervals[0])
	}

	var limit int64
	switch req := r.(type) {
	case *LogstoreRequest:
		limit = int64(req.Limit)
		if req.Direction == logproto.BACKWARD {
			for i, j := 0, len(intervals)-1; i < j; i, j = i+1, j-1 {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	case *LogstoreSeriesRequest, *LogstoreLabelNamesRequest, *logproto.IndexStatsRequest, *logproto.VolumeRequest:
		// Set this to 0 since this is not used in Series/Labels/Index Request.
		limit = 0
	default:
		return nil, httpgrpc.Errorf(http.StatusBadRequest, "unknown request type")
	}

	input := make([]*logstoreResult, 0, len(intervals))
	for _, interval := range intervals {
		input = append(input, &logstoreResult{
			req: interval,
			ch:  make(chan *packedResp),
		})
	}

	maxSeriesCapture := func(id string) int { return h.limits.MaxQuerySeries(ctx, id) }
	maxSeries := validation.SmallestPositiveIntPerTenant(tenantIDs, maxSeriesCapture)
	maxParallelism := MinWeightedParallelism(ctx, tenantIDs, h.configs, h.limits, model.Time(r.GetStart()), model.Time(r.GetEnd()))
	resps, err := h.Process(ctx, maxParallelism, limit, input, maxSeries)
	if err != nil {
		return nil, err
	}
	return h.merger.MergeResponse(resps...)
}

func splitByTime(req queryrangebase.Request, interval time.Duration) ([]queryrangebase.Request, error) {
	var reqs []queryrangebase.Request

	switch r := req.(type) {
	case *LogstoreRequest:
		util.ForInterval(interval, r.StartTs, r.EndTs, false, func(start, end time.Time) {
			reqs = append(reqs, &LogstoreRequest{
				Query:     r.Query,
				Limit:     r.Limit,
				Step:      r.Step,
				Interval:  r.Interval,
				Direction: r.Direction,
				Path:      r.Path,
				StartTs:   start,
				EndTs:     end,
			})
		})
	case *LogstoreSeriesRequest:
		// metadata queries have end time inclusive.
		// Set endTimeInclusive to true so that ForInterval keeps a gap of 1ms between splits to
		// avoid querying duplicate data in adjacent queries.
		util.ForInterval(interval, r.StartTs, r.EndTs, true, func(start, end time.Time) {
			reqs = append(reqs, &LogstoreSeriesRequest{
				Match:   r.Match,
				Path:    r.Path,
				StartTs: start,
				EndTs:   end,
				Shards:  r.Shards,
			})
		})
	case *LogstoreLabelNamesRequest:
		// metadata queries have end time inclusive.
		// Set endTimeInclusive to true so that ForInterval keeps a gap of 1ms between splits to
		// avoid querying duplicate data in adjacent queries.
		util.ForInterval(interval, r.StartTs, r.EndTs, true, func(start, end time.Time) {
			reqs = append(reqs, &LogstoreLabelNamesRequest{
				Path:    r.Path,
				StartTs: start,
				EndTs:   end,
				Query:   r.Query,
			})
		})
	case *logproto.IndexStatsRequest:
		startTS := model.Time(r.GetStart()).Time()
		endTS := model.Time(r.GetEnd()).Time()
		util.ForInterval(interval, startTS, endTS, true, func(start, end time.Time) {
			reqs = append(reqs, &logproto.IndexStatsRequest{
				From:     model.TimeFromUnix(start.Unix()),
				Through:  model.TimeFromUnix(end.Unix()),
				Matchers: r.GetMatchers(),
			})
		})
	case *logproto.VolumeRequest:
		startTS := model.Time(r.GetStart()).Time()
		endTS := model.Time(r.GetEnd()).Time()
		util.ForInterval(interval, startTS, endTS, true, func(start, end time.Time) {
			reqs = append(reqs, &logproto.VolumeRequest{
				From:         model.TimeFromUnix(start.Unix()),
				Through:      model.TimeFromUnix(end.Unix()),
				Matchers:     r.GetMatchers(),
				Limit:        r.Limit,
				TargetLabels: r.TargetLabels,
			})
		})
	default:
		return nil, nil
	}
	return reqs, nil
}

// maxRangeVectorAndOffsetDuration returns the maximum range vector and offset duration within a LogQL query.
func maxRangeVectorAndOffsetDuration(q string) (time.Duration, time.Duration, error) {
	expr, err := syntax.ParseExpr(q)
	if err != nil {
		return 0, 0, err
	}

	if _, ok := expr.(syntax.SampleExpr); !ok {
		return 0, 0, nil
	}

	var maxRVDuration, maxOffset time.Duration
	expr.Walk(func(e interface{}) {
		if r, ok := e.(*syntax.LogRange); ok {
			if r.Interval > maxRVDuration {
				maxRVDuration = r.Interval
			}
			if r.Offset > maxOffset {
				maxOffset = r.Offset
			}
		}
	})
	return maxRVDuration, maxOffset, nil
}

// reduceSplitIntervalForRangeVector reduces the split interval for a range query based on the duration of the range vector.
// Large range vector durations will not be split into smaller intervals because it can cause the queries to be slow by over-processing data.
func reduceSplitIntervalForRangeVector(r queryrangebase.Request, interval time.Duration) (time.Duration, error) {
	maxRange, _, err := maxRangeVectorAndOffsetDuration(r.GetQuery())
	if err != nil {
		return 0, err
	}
	if maxRange > interval {
		return maxRange, nil
	}
	return interval, nil
}

func splitMetricByTime(r queryrangebase.Request, interval time.Duration) ([]queryrangebase.Request, error) {
	var reqs []queryrangebase.Request

	interval, err := reduceSplitIntervalForRangeVector(r, interval)
	if err != nil {
		return nil, err
	}

	logstoreReq := r.(*LogstoreRequest)

	// step align start and end time of the query. Start time is rounded down and end time is rounded up.
	stepNs := r.GetStep() * 1e6
	startNs := logstoreReq.StartTs.UnixNano()
	start := time.Unix(0, startNs-startNs%stepNs)

	endNs := logstoreReq.EndTs.UnixNano()
	if mod := endNs % stepNs; mod != 0 {
		endNs += stepNs - mod
	}
	end := time.Unix(0, endNs)

	logstoreReq = logstoreReq.WithStartEnd(util.TimeToMillis(start), util.TimeToMillis(end)).(*LogstoreRequest)

	// step is >= configured split interval, let us just split the query interval by step
	if logstoreReq.Step >= interval.Milliseconds() {
		util.ForInterval(time.Duration(logstoreReq.Step*1e6), logstoreReq.StartTs, logstoreReq.EndTs, false, func(start, end time.Time) {
			reqs = append(reqs, &LogstoreRequest{
				Query:     logstoreReq.Query,
				Limit:     logstoreReq.Limit,
				Step:      logstoreReq.Step,
				Interval:  logstoreReq.Interval,
				Direction: logstoreReq.Direction,
				Path:      logstoreReq.Path,
				StartTs:   start,
				EndTs:     end,
			})
		})

		return reqs, nil
	}

	for start := logstoreReq.StartTs; start.Before(logstoreReq.EndTs); start = nextIntervalBoundary(start, r.GetStep(), interval).Add(time.Duration(r.GetStep()) * time.Millisecond) {
		end := nextIntervalBoundary(start, r.GetStep(), interval)
		if end.Add(time.Duration(r.GetStep())*time.Millisecond).After(logstoreReq.EndTs) || end.Add(time.Duration(r.GetStep())*time.Millisecond) == logstoreReq.EndTs {
			end = logstoreReq.EndTs
		}
		reqs = append(reqs, &LogstoreRequest{
			Query:     logstoreReq.Query,
			Limit:     logstoreReq.Limit,
			Step:      logstoreReq.Step,
			Interval:  logstoreReq.Interval,
			Direction: logstoreReq.Direction,
			Path:      logstoreReq.Path,
			StartTs:   start,
			EndTs:     end,
		})
	}

	return reqs, nil
}

// Round up to the step before the next interval boundary.
func nextIntervalBoundary(t time.Time, step int64, interval time.Duration) time.Time {
	stepNs := step * 1e6
	nsPerInterval := interval.Nanoseconds()
	startOfNextInterval := ((t.UnixNano() / nsPerInterval) + 1) * nsPerInterval
	// ensure that target is a multiple of steps away from the start time
	target := startOfNextInterval - ((startOfNextInterval - t.UnixNano()) % stepNs)
	if target == startOfNextInterval {
		target -= stepNs
	}
	return time.Unix(0, target)
}
