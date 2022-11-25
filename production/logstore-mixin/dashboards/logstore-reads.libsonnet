local utils = import 'mixin-utils/utils.libsonnet';

(import 'dashboard-utils.libsonnet') {
  acmeDashboards+: {
    local dashboards = self,
    local showBigTable = false,

    local http_routes = 'logstore_api_v1_series|api_prom_series|api_prom_query|api_prom_label|api_prom_label_name_values|logstore_api_v1_query|logstore_api_v1_query_range|logstore_api_v1_labels|logstore_api_v1_label_name_values',
    local grpc_routes = '/logproto.Querier/Query|/logproto.Querier/Label|/logproto.Querier/Series|/logproto.Querier/QuerySample|/logproto.Querier/GetChunkIDs',

    'logstore-reads.json': {
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
                           queryFrontend: [utils.selector.re('job', '($namespace)/%s' % (if $._config.ssd.enabled then '%s-read' % $._config.ssd.pod_prefix_matcher else 'query-frontend'))],
                           querier: [utils.selector.re('job', '($namespace)/%s' % (if $._config.ssd.enabled then '%s-write' % $._config.ssd.pod_prefix_matcher else 'querier'))],
                           ingester: [utils.selector.re('job', '($namespace)/%s' % (if $._config.ssd.enabled then '%s-write' % $._config.ssd.pod_prefix_matcher else 'ingester'))],
                           ingesterZoneAware: [utils.selector.re('job', '($namespace)/%s' % (if $._config.ssd.enabled then '%s-write' % $._config.ssd.pod_prefix_matcher else 'ingester-zone.*'))],
                           querierOrIndexGateway: [utils.selector.re('job', '($namespace)/%s' % (if $._config.ssd.enabled then '%s-read' % $._config.ssd.pod_prefix_matcher else '(querier|index-gateway)'))],
                         },

                         local selector(matcherId) =
                           local ms = (cfg.clusterMatchers + cfg.matchers[matcherId]);
                           if std.length(ms) > 0 then
                             std.join(',', ['%(label)s%(op)s"%(value)s"' % matcher for matcher in ms]) + ','
                           else '',

                         corestoreGwSelector:: selector('corestoregateway'),
                         queryFrontendSelector:: selector('queryFrontend'),
                         querierSelector:: selector('querier'),
                         ingesterSelector:: selector('ingester'),
                         ingesterZoneSelector:: selector('ingesterZoneAware'),
                         querierOrIndexGatewaySelector:: selector('querierOrIndexGateway'),
                       } +
                       $.dashboard('Logstore / Reads', uid='reads')
                       .addCluster()
                       .addNamespace()
                       .addTag()
                       .addRowIf(
                         $._config.internal_components,
                         $.row('Frontend (corestore_gw)')
                         .addPanel(
                           $.panel('QPS') +
                           $.qpsPanel('logstore_request_duration_seconds_count{%s route=~"%s"}' % [dashboards['logstore-reads.json'].corestoreGwSelector, http_routes])
                         )
                         .addPanel(
                           $.panel('Latency') +
                           utils.latencyRecordingRulePanel(
                             'logstore_request_duration_seconds',
                             dashboards['logstore-reads.json'].matchers.corestoregateway + [utils.selector.re('route', http_routes)],
                             extra_selectors=dashboards['logstore-reads.json'].clusterMatchers,
                             sum_by=['route']
                           )
                         )
                       )
                       .addRow(
                         $.row(if $._config.ssd.enabled then 'Read Path' else 'Frontend (query-frontend)')
                         .addPanel(
                           $.panel('QPS') +
                           $.qpsPanel('logstore_request_duration_seconds_count{%s route=~"%s"}' % [dashboards['logstore-reads.json'].queryFrontendSelector, http_routes])
                         )
                         .addPanel(
                           $.panel('Latency') +
                           utils.latencyRecordingRulePanel(
                             'logstore_request_duration_seconds',
                             dashboards['logstore-reads.json'].matchers.queryFrontend + [utils.selector.re('route', http_routes)],
                             extra_selectors=dashboards['logstore-reads.json'].clusterMatchers,
                             sum_by=['route']
                           )
                         )
                       )
                       .addRowIf(
                         !$._config.ssd.enabled,
                         $.row('Querier')
                         .addPanel(
                           $.panel('QPS') +
                           $.qpsPanel('logstore_request_duration_seconds_count{%s route=~"%s"}' % [dashboards['logstore-reads.json'].querierSelector, http_routes])
                         )
                         .addPanel(
                           $.panel('Latency') +
                           utils.latencyRecordingRulePanel(
                             'logstore_request_duration_seconds',
                             dashboards['logstore-reads.json'].matchers.querier + [utils.selector.re('route', http_routes)],
                             extra_selectors=dashboards['logstore-reads.json'].clusterMatchers,
                             sum_by=['route']
                           )
                         )
                       )
                       .addRowIf(
                         !$._config.ssd.enabled,
                         $.row('Ingester')
                         .addPanel(
                           $.panel('QPS') +
                           $.qpsPanel('logstore_request_duration_seconds_count{%s route=~"%s"}' % [dashboards['logstore-reads.json'].ingesterSelector, grpc_routes])
                         )
                         .addPanel(
                           $.panel('Latency') +
                           utils.latencyRecordingRulePanel(
                             'logstore_request_duration_seconds',
                             dashboards['logstore-reads.json'].matchers.ingester + [utils.selector.re('route', grpc_routes)],
                             extra_selectors=dashboards['logstore-reads.json'].clusterMatchers,
                             sum_by=['route']
                           )
                         )
                       )
                       // todo: add row iff multi zone ingesters are enabled
                       .addRowIf(
                         !$._config.ssd.enabled,
                         $.row('Ingester - Zone Aware')
                         .addPanel(
                           $.panel('QPS') +
                           $.qpsPanel('logstore_request_duration_seconds_count{%s route=~"%s"}' % [dashboards['logstore-reads.json'].ingesterZoneSelector, grpc_routes])
                         )
                         .addPanel(
                           $.panel('Latency') +
                           utils.latencyRecordingRulePanel(
                             'logstore_request_duration_seconds',
                             dashboards['logstore-reads.json'].matchers.ingesterZoneAware + [utils.selector.re('route', grpc_routes)],
                             extra_selectors=dashboards['logstore-reads.json'].clusterMatchers,
                             sum_by=['route']
                           )
                         )
                       )
                       .addRowIf(
                         !$._config.ssd.enabled,
                         $.row('Index')
                         .addPanel(
                           $.panel('QPS') +
                           $.qpsPanel('logstore_index_request_duration_seconds_count{%s operation!="index_chunk"}' % dashboards['logstore-reads.json'].querierSelector)
                         )
                         .addPanel(
                           $.panel('Latency') +
                           $.latencyPanel('logstore_index_request_duration_seconds', '{%s operation!="index_chunk"}' % dashboards['logstore-reads.json'].querierSelector)
                         )
                       )
                       .addRowIf(
                         showBigTable,
                         $.row('BigTable')
                         .addPanel(
                           $.panel('QPS') +
                           $.qpsPanel('logstore_bigtable_request_duration_seconds_count{%s operation="/google.bigtable.v2.Bigtable/ReadRows"}' % dashboards['logstore-reads.json'].querierSelector)
                         )
                         .addPanel(
                           $.panel('Latency') +
                           utils.latencyRecordingRulePanel(
                             'logstore_bigtable_request_duration_seconds',
                             dashboards['logstore-reads.json'].matchers.querier + [utils.selector.eq('operation', '/google.bigtable.v2.Bigtable/ReadRows')]
                           )
                         )
                       )
                       .addRow(
                         $.row('BoltDB Shipper')
                         .addPanel(
                           $.panel('QPS') +
                           $.qpsPanel('logstore_boltdb_shipper_request_duration_seconds_count{%s operation="Shipper.Query"}' % dashboards['logstore-reads.json'].querierOrIndexGatewaySelector)
                         )
                         .addPanel(
                           $.panel('Latency') +
                           $.latencyPanel('logstore_boltdb_shipper_request_duration_seconds', '{%s operation="Shipper.Query"}' % dashboards['logstore-reads.json'].querierOrIndexGatewaySelector)
                         )
                       ),
  },
}
