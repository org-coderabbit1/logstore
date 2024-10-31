package limists

import (
	"context"
	"time"

	"example.com/acme/logstore/v3/pkg/logql"
)

type TimeRangeLimits interface {
	MaxQueryLookback(context.Context, string) time.Duration
	MaxQueryLength(context.Context, string) time.Duration
}

type Limits interface {
	logql.Limits
	TimeRangeLimits
	QueryTimeout(context.Context, string) time.Duration
	MaxStreamsMatchersPerQuery(context.Context, string) int
	MaxConcurrentTailRequests(context.Context, string) int
	MaxEntriesLimitPerQuery(context.Context, string) int
}
