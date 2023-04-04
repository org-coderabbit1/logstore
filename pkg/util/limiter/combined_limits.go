package limiter

import (
	"example.com/acme/logstore/pkg/distributor"
	"example.com/acme/logstore/pkg/ingester"
	"example.com/acme/logstore/pkg/querier"
	"example.com/acme/logstore/pkg/querier/queryrange"
	"example.com/acme/logstore/pkg/ruler"
	"example.com/acme/logstore/pkg/scheduler"
	"example.com/acme/logstore/pkg/storage"
	"example.com/acme/logstore/pkg/storage/stores/indexshipper/compactor"
)

type CombinedLimits interface {
	compactor.Limits
	distributor.Limits
	ingester.Limits
	querier.Limits
	queryrange.Limits
	ruler.RulesLimits
	scheduler.Limits
	storage.StoreLimits
}
