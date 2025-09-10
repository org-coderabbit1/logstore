local utils = import 'mixin-utils/utils.libsonnet';

(import 'dashboard-utils.libsonnet') {
  acmeDashboards+: {
    local dashboards = self,
    local showBigTable = false,

    'logstore-writes.json': {
                          local cfg = self,

                          showMultiCluster:: true,
                          clusterMatchers::
                            if cfg.showMultiCluster then
                              [utils.selector.re($._config.per_cluster_label, '$cluster')]
                            else
                              [],

                          matchers:: {
                            corestoregateway: [utils.selector.re('job', '($namespace)/corestore-gw(-internal)?')],
                            distributor: if $._config.meta_monitoring.enabled
                            then [utils.selector.re('job', '($namespace)/(distributor|%s-write|logstore-single-binary)' % $._config.ssd.pod_prefix_matcher)]
                            else [utils.selector.re('job', '($namespace)/%s' % (if $._config.ssd.enabled then '%s-write' % $._config.ssd.pod_prefix_matcher else 'distributor'))],
                            ingester: if $._config.meta_monitoring.enabled
                            then [utils.selector.re('job', '($namespace)/(partition-ingester.*|ingester.*|%s-write|logstore-single-binary)' % $._config.ssd.pod_prefix_matcher)]
                            else [utils.selector.re('job', '($namespace)/%s' % (if $._config.ssd.enabled then '%s-write' % $._config.ssd.pod_prefix_matcher else '(ingester.*|partition-ingester.*)'))],
                            ingester_zone: if $._config.meta_monitoring.enabled
                            then [utils.selector.re('job', '($namespace)/(partition-ingester-.*|ingester-zone-.*|%s-write|logstore-single-binary)' % $._config.ssd.pod_prefix_matcher)]
                            else [utils.selector.re('job', '($namespace)/%s' % (if $._config.ssd.enabled then '%s-write' % $._config.ssd.pod_prefix_matcher else '(ingester-zone.*|partition-ingester-.*)'))],
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
                            $.newQueryPanel('QPS') +
                            $.newQpsPanel('logstore_request_duration_seconds_count{%s route=~"api_prom_push|logstore_api_v1_push|otlp_v1_logs"}' % dashboards['logstore-writes.json'].corestoreGwSelector)
                          )
                          .addPanel(
                            $.newQueryPanel('Latency', 'ms') +
                            utils.latencyRecordingRulePanel(
                              'logstore_request_duration_seconds',
                              dashboards['logstore-writes.json'].clusterMatchers +
                              dashboards['logstore-writes.json'].matchers.corestoregateway +
                              [utils.selector.re('route', 'api_prom_push|logstore_api_v1_push|otlp_v1_logs')],
                            )
                          )
                          .addPanel(
                            $.p99LatencyByPod(
                              'logstore_request_duration_seconds',
                              $.toPrometheusSelector(
                                dashboards['logstore-writes.json'].clusterMatchers +
                                dashboards['logstore-writes.json'].matchers.corestoregateway +
                                [utils.selector.re('route', 'api_prom_push|logstore_api_v1_push|otlp_v1_logs')],
                              ),
                            )
                          )
                        )
                        .addRow(
                          $.row(if $._config.ssd.enabled then 'Write Path' else 'Distributor')
                          .addPanel(
                            $.newQueryPanel('QPS') +
                            $.newQpsPanel('logstore_request_duration_seconds_count{%s, route=~"api_prom_push|logstore_api_v1_push|otlp_v1_logs|/httpgrpc.HTTP/Handle"}' % std.rstripChars(dashboards['logstore-writes.json'].distributorSelector, ','))
                          )
                          .addPanel(
                            $.newQueryPanel('Latency', 'ms') +
                            utils.latencyRecordingRulePanel(
                              'logstore_request_duration_seconds',
                              dashboards['logstore-writes.json'].clusterMatchers +
                              dashboards['logstore-writes.json'].matchers.distributor +
                              [utils.selector.re('route', 'api_prom_push|logstore_api_v1_push|otlp_v1_logs|/httpgrpc.HTTP/Handle')],
                            )
                          )
                          .addPanel(
                            $.p99LatencyByPod(
                              'logstore_request_duration_seconds',
                              $.toPrometheusSelector(
                                dashboards['logstore-writes.json'].clusterMatchers +
                                dashboards['logstore-writes.json'].matchers.distributor +
                                [utils.selector.re('route', 'api_prom_push|logstore_api_v1_push|otlp_v1_logs|/httpgrpc.HTTP/Handle')],
                              ),
                            )
                          )
                        )
                        .addRowIf(
                          $._config.tsdb,
                          $.row((if $._config.ssd.enabled then 'Write Path' else 'Distributor') + ' - Structured Metadata')
                          .addPanel(
                            $.newQueryPanel('Per Total Received Bytes') +
                            $.queryPanel('sum (rate(logstore_distributor_structured_metadata_bytes_received_total{%s}[$__rate_interval])) / sum(rate(logstore_distributor_bytes_received_total{%s}[$__rate_interval]))' % [dashboards['logstore-writes.json'].distributorSelector, dashboards['logstore-writes.json'].distributorSelector], 'bytes')
                          )
                          .addPanel(
                            $.newQueryPanel('Per Tenant') +
                            $.queryPanel('sum by (tenant) (rate(logstore_distributor_structured_metadata_bytes_received_total{%s}[$__rate_interval])) / ignoring(tenant) group_left sum(rate(logstore_distributor_structured_metadata_bytes_received_total{%s}[$__rate_interval]))' % [dashboards['logstore-writes.json'].distributorSelector, dashboards['logstore-writes.json'].distributorSelector], '{{tenant}}') + {
                              stack: true,
                              yaxes: [
                                { format: 'short', label: null, logBase: 1, max: 1, min: 0, show: true },
                                { format: 'short', label: null, logBase: 1, max: 1, min: null, show: false },
                              ],
                            },
                          )
                        )
                        .addRowIf(
                          !$._config.ssd.enabled,
                          $.row('Ingester - Zone Aware')
                          .addPanel(
                            $.newQueryPanel('QPS') +
                            $.newQpsPanel('logstore_request_duration_seconds_count{%s route="/logproto.Pusher/Push"}' % dashboards['logstore-writes.json'].ingesterZoneSelector)
                          )
                          .addPanel(
                            $.newQueryPanel('Latency', 'ms') +
                            utils.latencyRecordingRulePanel(
                              'logstore_request_duration_seconds',
                              dashboards['logstore-writes.json'].clusterMatchers +
                              dashboards['logstore-writes.json'].matchers.ingester_zone +
                              [utils.selector.eq('route', '/logproto.Pusher/Push')],
                            )
                          )
                          .addPanel(
                            $.p99LatencyByPod(
                              'logstore_request_duration_seconds',
                              $.toPrometheusSelector(
                                dashboards['logstore-writes.json'].clusterMatchers +
                                dashboards['logstore-writes.json'].matchers.ingester_zone +
                                [utils.selector.eq('route', '/logproto.Pusher/Push')],
                              ),
                            )
                          )
                        )
                        .addRowIf(
                          !$._config.ssd.enabled,
                          $.row('Ingester')
                          .addPanel(
                            $.newQueryPanel('QPS') +
                            $.newQpsPanel('logstore_request_duration_seconds_count{%s route="/logproto.Pusher/Push"}' % dashboards['logstore-writes.json'].ingesterSelector)
                          )
                          .addPanel(
                            $.newQueryPanel('Latency', 'ms') +
                            utils.latencyRecordingRulePanel(
                              'logstore_request_duration_seconds',
                              dashboards['logstore-writes.json'].clusterMatchers +
                              dashboards['logstore-writes.json'].matchers.ingester +
                              [utils.selector.eq('route', '/logproto.Pusher/Push')],
                            )
                          )
                          .addPanel(
                            $.p99LatencyByPod(
                              'logstore_request_duration_seconds',
                              $.toPrometheusSelector(
                                dashboards['logstore-writes.json'].clusterMatchers +
                                dashboards['logstore-writes.json'].matchers.ingester +
                                [utils.selector.eq('route', '/logproto.Pusher/Push')]
                              ),
                            )
                          )
                        )
                        .addRowIf(
                          !$._config.ssd.enabled,
                          $.row('Index')
                          .addPanel(
                            $.newQueryPanel('QPS') +
                            $.newQpsPanel('logstore_index_request_duration_seconds_count{%s operation="index_chunk"}' % dashboards['logstore-writes.json'].ingesterSelector)
                          )
                          .addPanel(
                            $.newQueryPanel('Latency', 'ms') +
                            $.latencyPanel('logstore_index_request_duration_seconds', '{%s operation="index_chunk"}' % dashboards['logstore-writes.json'].ingesterSelector)
                          )
                          .addPanel(
                            $.p99LatencyByPod(
                              'logstore_index_request_duration_seconds',
                              '{%s operation="index_chunk"}' % dashboards['logstore-writes.json'].ingesterSelector,
                            )
                          )
                        )
                        .addRowIf(
                          showBigTable,
                          $.row('BigTable')
                          .addPanel(
                            $.newQueryPanel('QPS') +
                            $.newQpsPanel('logstore_bigtable_request_duration_seconds_count{%s operation="/google.bigtable.v2.Bigtable/MutateRows"}' % dashboards['logstore-writes.json'].ingesterSelector)
                          )
                          .addPanel(
                            $.newQueryPanel('Latency', 'ms') +
                            utils.latencyRecordingRulePanel(
                              'logstore_bigtable_request_duration_seconds',
                              dashboards['logstore-writes.json'].clusterMatchers +
                              dashboards['logstore-writes.json'].matchers.ingester +
                              [utils.selector.eq('operation', '/google.bigtable.v2.Bigtable/MutateRows')],
                            )
                          )
                          .addPanel(
                            $.p99LatencyByPod(
                              'logstore_bigtable_request_duration_seconds',
                              $.toPrometheusSelector(
                                dashboards['logstore-writes.json'].clusterMatchers +
                                dashboards['logstore-writes.json'].matchers.ingester +
                                [utils.selector.eq('operation', '/google.bigtable.v2.Bigtable/MutateRows')],
                              ),
                            )
                          )
                        )
                        .addRowIf(
                          !$._config.ssd.enabled,
                          $.row('BoltDB Index')
                          .addPanel(
                            $.newQueryPanel('QPS') +
                            $.newQpsPanel('logstore_boltdb_shipper_request_duration_seconds_count{%s operation="WRITE"}' % dashboards['logstore-writes.json'].ingesterSelector)
                          )
                          .addPanel(
                            $.newQueryPanel('Latency', 'ms') +
                            $.latencyPanel('logstore_boltdb_shipper_request_duration_seconds', '{%s operation="WRITE"}' % dashboards['logstore-writes.json'].ingesterSelector)
                          )
                          .addPanel(
                            $.p99LatencyByPod(
                              'logstore_boltdb_shipper_request_duration_seconds',
                              '{%s operation="WRITE"}' % dashboards['logstore-writes.json'].ingesterSelector,
                            )
                          )
                        ),
  },
}
