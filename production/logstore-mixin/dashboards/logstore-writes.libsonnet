local utils = import 'mixin-utils/utils.libsonnet';

(import 'dashboard-utils.libsonnet') {
  acmeDashboards+: {
    local dashboards = self,
    local showBigTable = false,

    'logstore-writes.json': {
                          local cfg = self,

                          showMultiCluster:: true,
                          clusterLabel:: $._config.per_cluster_label,
                          clusterMatchers::
                            if cfg.showMultiCluster then
                              [utils.selector.re(cfg.clusterLabel, '$cluster')]
                            else
                              [],

                          matchers:: {
                            corestoregateway: [utils.selector.re('job', '($namespace)/corestore-gw(-internal)?')],
                            distributor: [utils.selector.re('job', '($namespace)/%s' % (if $._config.ssd.enabled then '%s-write' % $._config.ssd.pod_prefix_matcher else 'distributor'))],
                            ingester: [utils.selector.re('job', '($namespace)/%s' % (if $._config.ssd.enabled then '%s-write' % $._config.ssd.pod_prefix_matcher else 'ingester'))],
                            ingester_zone: [utils.selector.re('job', '($namespace)/%s' % (if $._config.ssd.enabled then '%s-write' % $._config.ssd.pod_prefix_matcher else 'ingester-zone.*'))],
                          },

                          local selector(matcherId) =
                            local ms = cfg.clusterMatchers + cfg.matchers[matcherId];
                            if std.length(ms) > 0 then
                              std.join(',', ['%(label)s%(op)s"%(value)s"' % matcher for matcher in ms]) + ','
                            else '',

                          corestoreGwSelector:: selector('corestoregateway'),
                          distributorSelector:: selector('distributor'),
                          ingesterSelector:: selector('ingester'),
                          ingesterZoneSelector:: selector('ingester_zone'),
                        } +
                        $.dashboard('Logstore / Writes', uid='writes')
                        .addCluster()
                        .addNamespace()
                        .addTag()
                        .addRowIf(
                          $._config.internal_components,
                          $.row('Frontend (corestore_gw)')
                          .addPanel(
                            $.panel('QPS') +
                            $.qpsPanel('logstore_request_duration_seconds_count{%s route=~"api_prom_push|logstore_api_v1_push"}' % dashboards['logstore-writes.json'].corestoreGwSelector)
                          )
                          .addPanel(
                            $.panel('Latency') +
                            utils.latencyRecordingRulePanel(
                              'logstore_request_duration_seconds',
                              dashboards['logstore-writes.json'].clusterMatchers + dashboards['logstore-writes.json'].matchers.corestoregateway + [utils.selector.re('route', 'api_prom_push|logstore_api_v1_push')],
                            )
                          )
                        )
                        .addRow(
                          $.row(if $._config.ssd.enabled then 'Write Path' else 'Distributor')
                          .addPanel(
                            $.panel('QPS') +
                            $.qpsPanel('logstore_request_duration_seconds_count{%s, route=~"api_prom_push|logstore_api_v1_push|/httpgrpc.HTTP/Handle"}' % std.rstripChars(dashboards['logstore-writes.json'].distributorSelector, ','))
                          )
                          .addPanel(
                            $.panel('Latency') +
                            utils.latencyRecordingRulePanel(
                              'logstore_request_duration_seconds',
                              dashboards['logstore-writes.json'].clusterMatchers + dashboards['logstore-writes.json'].matchers.distributor,
                            )
                          )
                        )
                        .addRowIf(
                          !$._config.ssd.enabled,
                          $.row('Ingester - Zone Aware')
                          .addPanel(
                            $.panel('QPS') +
                            $.qpsPanel('logstore_request_duration_seconds_count{%s route="/logproto.Pusher/Push"}' % dashboards['logstore-writes.json'].ingesterZoneSelector)
                          )
                          .addPanel(
                            $.panel('Latency') +
                            utils.latencyRecordingRulePanel(
                              'logstore_request_duration_seconds',
                              dashboards['logstore-writes.json'].clusterMatchers + dashboards['logstore-writes.json'].matchers.ingester_zone + [utils.selector.eq('route', '/logproto.Pusher/Push')],
                            )
                          )
                        )
                        .addRowIf(
                          !$._config.ssd.enabled,
                          $.row('Ingester')
                          .addPanel(
                            $.panel('QPS') +
                            $.qpsPanel('logstore_request_duration_seconds_count{%s route="/logproto.Pusher/Push"}' % dashboards['logstore-writes.json'].ingesterSelector) +
                            $.qpsPanel('logstore_request_duration_seconds_count{%s route="/logproto.Pusher/Push"}' % dashboards['logstore-writes.json'].ingesterSelector)
                          )
                          .addPanel(
                            $.panel('Latency') +
                            utils.latencyRecordingRulePanel(
                              'logstore_request_duration_seconds',
                              dashboards['logstore-writes.json'].clusterMatchers + dashboards['logstore-writes.json'].matchers.ingester + [utils.selector.eq('route', '/logproto.Pusher/Push')],
                            )
                          )
                        )
                        .addRowIf(
                          !$._config.ssd.enabled,
                          $.row('Index')
                          .addPanel(
                            $.panel('QPS') +
                            $.qpsPanel('logstore_index_request_duration_seconds_count{%s operation="index_chunk"}' % dashboards['logstore-writes.json'].ingesterSelector)
                          )
                          .addPanel(
                            $.panel('Latency') +
                            $.latencyPanel('logstore_index_request_duration_seconds', '{%s operation="index_chunk"}' % dashboards['logstore-writes.json'].ingesterSelector)
                          )
                        )
                        .addRowIf(
                          showBigTable,
                          $.row('BigTable')
                          .addPanel(
                            $.panel('QPS') +
                            $.qpsPanel('logstore_bigtable_request_duration_seconds_count{%s operation="/google.bigtable.v2.Bigtable/MutateRows"}' % dashboards['logstore-writes.json'].ingesterSelector)
                          )
                          .addPanel(
                            $.panel('Latency') +
                            utils.latencyRecordingRulePanel(
                              'logstore_bigtable_request_duration_seconds',
                              dashboards['logstore-writes.json'].clusterMatchers + dashboards['logstore-writes.json'].clusterMatchers + dashboards['logstore-writes.json'].matchers.ingester + [utils.selector.eq('operation', '/google.bigtable.v2.Bigtable/MutateRows')]
                            )
                          )
                        )
                        .addRow(
                          $.row('BoltDB Shipper')
                          .addPanel(
                            $.panel('QPS') +
                            $.qpsPanel('logstore_boltdb_shipper_request_duration_seconds_count{%s operation="WRITE"}' % dashboards['logstore-writes.json'].ingesterSelector)
                          )
                          .addPanel(
                            $.panel('Latency') +
                            $.latencyPanel('logstore_boltdb_shipper_request_duration_seconds', '{%s operation="WRITE"}' % dashboards['logstore-writes.json'].ingesterSelector)
                          )
                        ),
  },
}
