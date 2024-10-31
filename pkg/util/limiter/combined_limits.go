package limiter

import (
	bloombuilder "example.com/acme/logstore/v3/pkg/bloombuild/builder"
	bloomplanner "example.com/acme/logstore/v3/pkg/bloombuild/planner"
	"example.com/acme/logstore/v3/pkg/bloomgateway"
	"example.com/acme/logstore/v3/pkg/compactor"
	"example.com/acme/logstore/v3/pkg/distributor"
	"example.com/acme/logstore/v3/pkg/indexgateway"
	"example.com/acme/logstore/v3/pkg/ingester"
	"example.com/acme/logstore/v3/pkg/pattern"
	querier_limits "example.com/acme/logstore/v3/pkg/querier/limits"
	queryrange_limits "example.com/acme/logstore/v3/pkg/querier/queryrange/limits"
	"example.com/acme/logstore/v3/pkg/ruler"
	scheduler_limits "example.com/acme/logstore/v3/pkg/scheduler/limits"
	"example.com/acme/logstore/v3/pkg/storage"
	"example.com/acme/logstore/v3/pkg/storage/bucket"
)

type CombinedLimits interface {
	compactor.Limits
	distributor.Limits
	ingester.Limits
	querier_limits.Limits
	queryrange_limits.Limits
	ruler.RulesLimits
	scheduler_limits.Limits
	storage.StoreLimits
	indexgateway.Limits
	bloomgateway.Limits
	bloomplanner.Limits
	bloombuilder.Limits
	pattern.Limits
	bucket.SSEConfigProvider
}
