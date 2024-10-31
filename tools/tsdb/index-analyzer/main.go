package main

import (
	"flag"

	"github.com/prometheus/client_golang/prometheus"

	"example.com/acme/logstore/v3/pkg/storage"
	"example.com/acme/logstore/v3/pkg/storage/stores/shipper/indexshipper"
	"example.com/acme/logstore/v3/pkg/storage/stores/shipper/indexshipper/tsdb"
	util_log "example.com/acme/logstore/v3/pkg/util/log"
	"example.com/acme/logstore/v3/tools/tsdb/helpers"
)

// go build ./tools/tsdb/index-analyzer && BUCKET=19453 DIR=/tmp/logstore-index-analysis ./index-analyzer --config.file=/tmp/logstore-config.yaml
func main() {
	conf, _, bucket, err := helpers.Setup()
	helpers.ExitErr("setting up", err)

	_, overrides, clientMetrics := helpers.DefaultConfigs()

	flag.Parse()

	periodCfg, tableRange, tableName, err := helpers.GetPeriodConfigForTableNumber(bucket, conf.SchemaConfig.Configs)
	helpers.ExitErr("find period config for bucket", err)

	objectClient, err := storage.NewObjectClient(periodCfg.ObjectType, "index-analyzer", conf.StorageConfig, clientMetrics)
	helpers.ExitErr("creating object client", err)

	shipper, err := indexshipper.NewIndexShipper(
		periodCfg.IndexTables.PathPrefix,
		conf.StorageConfig.TSDBShipperConfig,
		objectClient,
		overrides,
		nil,
		tsdb.OpenShippableTSDB,
		tableRange,
		prometheus.WrapRegistererWithPrefix("logstore_tsdb_shipper_", prometheus.DefaultRegisterer),
		util_log.Logger,
	)
	helpers.ExitErr("creating index shipper", err)

	tenants, err := helpers.ResolveTenants(objectClient, periodCfg.IndexTables.PathPrefix, tableName)
	helpers.ExitErr("resolving tenants", err)

	err = analyze(shipper, tableName, tenants)
	helpers.ExitErr("analyzing", err)
}
