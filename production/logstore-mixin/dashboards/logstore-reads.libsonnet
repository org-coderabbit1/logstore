local utils = import 'mixin-utils/utils.libsonnet';

(import 'dashboard-utils.libsonnet') {
  acmeDashboards+: {
    local dashboards = self,
    local showBigTable = false,

    // Available HTTP routes can be collected with the following instant query:
    // count by (route) (logstore_request_duration_seconds_count{route!~"/.*"})
    local http_routes = '(%s)' % std.join(
      '|', [
        'api_prom_rules',
        'api_prom_rules_namespace_groupname',
        'api_v1_rules',
        'logstore_api_v1_delete',
        'logstore_api_v1_detected_labels',
        'logstore_api_v1_index_stats',
        'logstore_api_v1_index_volume',
        'logstore_api_v1_index_volume_range',
        'logstore_api_v1_label_name_values',
        'logstore_api_v1_label_values',
        'logstore_api_v1_labels',
        'logstore_api_v1_patterns',
        'logstore_api_v1_query',
        'logstore_api_v1_query_range',
        'logstore_api_v1_series',
        'prometheus_api_v1_rules',
      ]
    ),

    // Available GRPC routes can be collected with the following instant query:
    // count by (route) (logstore_request_duration_seconds_count{route=~"/.*"})
    local grpc_routes = '(%s)' % std.join(
      '|', [
        '/base.Ruler/Rules',
        '/indexgatewaypb.IndexGateway/GetChunkRef',
        '/indexgatewaypb.IndexGateway/GetSeries',
        '/indexgatewaypb.IndexGateway/GetShards',
        '/indexgatewaypb.IndexGateway/GetStats',
        '/indexgatewaypb.IndexGateway/GetVolume',
        '/indexgatewaypb.IndexGateway/LabelNamesForMetricName',
        '/indexgatewaypb.IndexGateway/LabelValuesForMetricName',
        '/indexgatewaypb.IndexGateway/QueryIndex',
        '/logproto.BloomGateway/FilterChunkRefs',
        '/logproto.Pattern/Query',
        '/logproto.Querier/GetChunkIDs',
        '/logproto.Querier/GetDetectedLabels',
        '/logproto.Querier/GetStats',
        '/logproto.Querier/GetVolume',
        '/logproto.Querier/Label',
        '/logproto.Querier/Query',
        '/logproto.Querier/QuerySample',
        '/logproto.Querier/Series',
        '/logproto.StreamData/GetStreamRates',
      ]
    ),

    'logstore-reads.json': {
                         local cfg = self,

                         showMultiCluster:: true,
                         clusterMatchers::
                           if cfg.showMultiCluster then
                             [utils.selector.re($._config.per_cluster_label, '$cluster')]
                           else
                             [],

                         matchers:: {
                           corestoregateway: [utils.selector.re('job', '($namespace)/corestore-gw(-internal)?')],
                           queryFrontend: if $._config.meta_monitoring.enabled
                           then [utils.selector.re('job', '($namespace)/(query-frontend|%s-read|logstore-single-binary)' % $._config.ssd.pod_prefix_matcher)]
                           else [utils.selector.re('job', '($namespace)/%s' % (if $._config.ssd.enabled then '%s-read' % $._config.ssd.pod_prefix_matcher else 'query-frontend'))],
                           querier: if $._config.meta_monitoring.enabled
                           then [utils.selector.re('job', '($namespace)/(querier|%s-read|logstore-single-binary)' % $._config.ssd.pod_prefix_matcher)]
                           else [utils.selector.re('job', '($namespace)/%s' % (if $._config.ssd.enabled then '%s-read' % $._config.ssd.pod_prefix_matcher else 'querier'))],
                           ingester: if $._config.meta_monitoring.enabled
                           then [utils.selector.re('job', '($namespace)/(partition-ingester.*|ingester.*|%s-write|logstore-single-binary)' % $._config.ssd.pod_prefix_matcher)]
                           else [utils.selector.re('job', '($namespace)/%s' % (if $._config.ssd.enabled then '%s-write' % $._config.ssd.pod_prefix_matcher else '(ingester.*|partition-ingester.*)'))],
                           ingesterZoneAware: if $._config.meta_monitoring.enabled
                           then [utils.selector.re('job', '($namespace)/(partition-ingester-.*|ingester-zone-.*|%s-write|logstore-single-binary)' % $._config.ssd.pod_prefix_matcher)]
                           else [utils.selector.re('job', '($namespace)/%s' % (if $._config.ssd.enabled then '%s-write' % $._config.ssd.pod_prefix_matcher else '(partition-ingester-.*|ingester-zone.*)'))],
                           querierOrIndexGateway: if $._config.meta_monitoring.enabled
                           then [utils.selector.re('job', '($namespace)/(querier|index-gateway|%s-read|logstore-single-binary)' % $._config.ssd.pod_prefix_matcher)]
                           else [utils.selector.re('job', '($namespace)/%s' % (if $._config.ssd.enabled then '%s-read' % $._config.ssd.pod_prefix_matcher else '(querier|index-gateway)'))],
                           indexGateway: if $._config.meta_monitoring.enabled
                           then [utils.selector.re('job', '($namespace)/(index-gateway|%s-backend|logstore-single-binary)' % $._config.ssd.pod_prefix_matcher)]
                           else [utils.selector.re('job', '($namespace)/%s' % (if $._config.ssd.enabled then '%s-backend' % $._config.ssd.pod_prefix_matcher else 'index-gateway'))],
                           bloomGateway: if $._config.meta_monitoring.enabled
                           then [utils.selector.re('job', '($namespace)/(bloom-gateway|%s-backend|logstore-single-binary)' % $._config.ssd.pod_prefix_matcher)]
                           else [utils.selector.re('job', '($namespace)/%s' % (if $._config.ssd.enabled then '%s-backend' % $._config.ssd.pod_prefix_matcher else 'bloom-gateway'))],
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
                         indexGatewaySelector:: selector('indexGateway'),
                         bloomGatewaySelector:: selector('bloomGateway'),
                       } +
                       $.dashboard('Logstore / Reads', uid='reads')
                       .addCluster()
                       .addNamespace()
                       .addTag()
                       .addRowIf(
                         $._config.internal_components,
                         $.row('Frontend (corestore_gw)')
                         .addPanel(
                           $.newQueryPanel('QPS') +
                           $.newQpsPanel('logstore_request_duration_seconds_count{%s route=~"%s"}' % [dashboards['logstore-reads.json'].corestoreGwSelector, http_routes])
                         )
                         .addPanel(
                           $.newQueryPanel('Latency', 'ms') +
                           utils.latencyRecordingRulePanel(
                             'logstore_request_duration_seconds',
                             dashboards['logstore-reads.json'].clusterMatchers + dashboards['logstore-reads.json'].matchers.corestoregateway + [utils.selector.re('route', http_routes)],
                             sum_by=['route']
                           )
                         )
                         .addPanel(
                           $.p99LatencyByPod(
                             'logstore_request_duration_seconds',
                             $.toPrometheusSelector(
                               dashboards['logstore-reads.json'].clusterMatchers + dashboards['logstore-reads.json'].matchers.corestoregateway + [utils.selector.re('route', http_routes)]
                             ),
                           )
                         )
                       )
                       .addRow(
                         $.row(if $._config.ssd.enabled then 'Read Path' else 'Frontend (query-frontend)')
                         .addPanel(
                           $.newQueryPanel('QPS') +
                           $.newQpsPanel('logstore_request_duration_seconds_count{%s route=~"%s"}' % [dashboards['logstore-reads.json'].queryFrontendSelector, http_routes])
                         )
                         .addPanel(
                           $.newQueryPanel('Latency', 'ms') +
                           utils.latencyRecordingRulePanel(
                             'logstore_request_duration_seconds',
                             dashboards['logstore-reads.json'].clusterMatchers + dashboards['logstore-reads.json'].matchers.queryFrontend + [utils.selector.re('route', http_routes)],
                             sum_by=['route']
                           )
                         )
                         .addPanel(
                           $.p99LatencyByPod(
                             'logstore_request_duration_seconds',
                             $.toPrometheusSelector(
                               dashboards['logstore-reads.json'].clusterMatchers +
                               dashboards['logstore-reads.json'].matchers.queryFrontend +
                               [utils.selector.re('route', http_routes)]
                             ),
                           )
                         )
                       )
                       .addRowIf(
                         !$._config.ssd.enabled,
                         $.row('Querier')
                         .addPanel(
                           $.newQueryPanel('QPS') +
                           $.newQpsPanel('logstore_request_duration_seconds_count{%s route=~"%s"}' % [dashboards['logstore-reads.json'].querierSelector, http_routes])
                         )
                         .addPanel(
                           $.newQueryPanel('Latency', 'ms') +
                           utils.latencyRecordingRulePanel(
                             'logstore_request_duration_seconds',
                             dashboards['logstore-reads.json'].clusterMatchers + dashboards['logstore-reads.json'].matchers.querier + [utils.selector.re('route', http_routes)],
                             sum_by=['route']
                           )
                         )
                         .addPanel(
                           $.p99LatencyByPod(
                             'logstore_request_duration_seconds',
                             $.toPrometheusSelector(
                               dashboards['logstore-reads.json'].clusterMatchers +
                               dashboards['logstore-reads.json'].matchers.querier +
                               [utils.selector.re('route', http_routes)]
                             ),
                           )
                         )
                       )
                       .addRowIf(
                         !$._config.ssd.enabled,
                         $.row('Ingester')
                         .addPanel(
                           $.newQueryPanel('QPS') +
                           $.newQpsPanel('logstore_request_duration_seconds_count{%s route=~"%s"}' % [dashboards['logstore-reads.json'].ingesterSelector, grpc_routes])
                         )
                         .addPanel(
                           $.newQueryPanel('Latency', 'ms') +
                           utils.latencyRecordingRulePanel(
                             'logstore_request_duration_seconds',
                             dashboards['logstore-reads.json'].clusterMatchers + dashboards['logstore-reads.json'].matchers.ingester + [utils.selector.re('route', grpc_routes)],
                             sum_by=['route']
                           )
                         )
                         .addPanel(
                           $.p99LatencyByPod(
                             'logstore_request_duration_seconds',
                             $.toPrometheusSelector(
                               dashboards['logstore-reads.json'].clusterMatchers +
                               dashboards['logstore-reads.json'].matchers.ingester +
                               [utils.selector.re('route', grpc_routes)]
                             ),
                           )
                         )
                       )
                       // todo: add row iff multi zone ingesters are enabled
                       .addRowIf(
                         !$._config.ssd.enabled,
                         $.row('Ingester - Zone Aware')
                         .addPanel(
                           $.newQueryPanel('QPS') +
                           $.newQpsPanel('logstore_request_duration_seconds_count{%s route=~"%s"}' % [dashboards['logstore-reads.json'].ingesterZoneSelector, grpc_routes])
                         )
                         .addPanel(
                           $.newQueryPanel('Latency', 'ms') +
                           utils.latencyRecordingRulePanel(
                             'logstore_request_duration_seconds',
                             dashboards['logstore-reads.json'].clusterMatchers + dashboards['logstore-reads.json'].matchers.ingesterZoneAware + [utils.selector.re('route', grpc_routes)],
                             sum_by=['route']
                           )
                         )
                         .addPanel(
                           $.p99LatencyByPod(
                             'logstore_request_duration_seconds',
                             $.toPrometheusSelector(
                               dashboards['logstore-reads.json'].clusterMatchers +
                               dashboards['logstore-reads.json'].matchers.ingesterZoneAware +
                               [utils.selector.re('route', grpc_routes)]
                             ),
                           )
                         )
                       )
                       .addRowIf(
                         !$._config.ssd.enabled,
                         $.row('Index Gateway')
                         .addPanel(
                           $.newQueryPanel('QPS') +
                           $.newQpsPanel('logstore_request_duration_seconds_count{%s route=~"%s"}' % [dashboards['logstore-reads.json'].indexGatewaySelector, grpc_routes])
                         )
                         .addPanel(
                           $.newQueryPanel('Latency', 'ms') +
                           utils.latencyRecordingRulePanel(
                             'logstore_request_duration_seconds',
                             dashboards['logstore-reads.json'].clusterMatchers + dashboards['logstore-reads.json'].matchers.indexGateway + [utils.selector.re('route', grpc_routes)],
                             sum_by=['route']
                           )
                         )
                         .addPanel(
                           $.p99LatencyByPod(
                             'logstore_request_duration_seconds',
                             $.toPrometheusSelector(
                               dashboards['logstore-reads.json'].clusterMatchers +
                               dashboards['logstore-reads.json'].matchers.indexGateway +
                               [utils.selector.re('route', grpc_routes)]
                             ),
                           )
                         )
                       )
                       .addRowIf(
                         !$._config.ssd.enabled,
                         $.row('Bloom Gateway')
                         .addPanel(
                           $.newQueryPanel('QPS') +
                           $.newQpsPanel('logstore_request_duration_seconds_count{%s route=~"%s"}' % [dashboards['logstore-reads.json'].bloomGatewaySelector, grpc_routes])
                         )
                         .addPanel(
                           $.newQueryPanel('Latency', 'ms') +
                           utils.latencyRecordingRulePanel(
                             'logstore_request_duration_seconds',
                             dashboards['logstore-reads.json'].clusterMatchers + dashboards['logstore-reads.json'].matchers.bloomGateway + [utils.selector.re('route', grpc_routes)],
                             sum_by=['route']
                           )
                         )
                         .addPanel(
                           $.p99LatencyByPod(
                             'logstore_request_duration_seconds',
                             $.toPrometheusSelector(
                               dashboards['logstore-reads.json'].clusterMatchers +
                               dashboards['logstore-reads.json'].matchers.bloomGateway +
                               [utils.selector.re('route', grpc_routes)]
                             ),
                           )
                         )
                       )
                       .addRowIf(
                         showBigTable,
                         $.row('BigTable')
                         .addPanel(
                           $.newQueryPanel('QPS') +
                           $.newQpsPanel('logstore_bigtable_request_duration_seconds_count{%s operation="/google.bigtable.v2.Bigtable/ReadRows"}' % dashboards['logstore-reads.json'].querierSelector)
                         )
                         .addPanel(
                           $.newQueryPanel('Latency', 'ms') +
                           utils.latencyRecordingRulePanel(
                             'logstore_bigtable_request_duration_seconds',
                             dashboards['logstore-reads.json'].clusterMatchers + dashboards['logstore-reads.json'].matchers.querier + [utils.selector.eq('operation', '/google.bigtable.v2.Bigtable/ReadRows')]
                           )
                         )
                       )
                       .addRowIf(
                         !$._config.ssd.enabled,
                         $.row('TSBD Index')
                         .addPanel(
                           $.newQueryPanel('QPS') +
                           $.newQpsPanel('logstore_index_request_duration_seconds_count{%s operation!="index_chunk"}' % dashboards['logstore-reads.json'].querierSelector)
                         )
                         .addPanel(
                           $.newQueryPanel('Latency', 'ms') +
                           $.latencyPanel('logstore_index_request_duration_seconds', '{%s operation!="index_chunk"}' % dashboards['logstore-reads.json'].querierSelector)
                         )
                         .addPanel(
                           $.p99LatencyByPod(
                             'logstore_index_request_duration_seconds',
                             '{%s operation!="index_chunk"}' % dashboards['logstore-reads.json'].querierSelector
                           )
                         )
                       )
                       .addRowIf(
                         !$._config.ssd.enabled,
                         $.row('BoltDB Index')
                         .addPanel(
                           $.newQueryPanel('QPS') +
                           $.newQpsPanel('logstore_boltdb_shipper_request_duration_seconds_count{%s operation="Shipper.Query"}' % dashboards['logstore-reads.json'].querierOrIndexGatewaySelector)
                         )
                         .addPanel(
                           $.newQueryPanel('Latency', 'ms') +
                           $.latencyPanel('logstore_boltdb_shipper_request_duration_seconds', '{%s operation="Shipper.Query"}' % dashboards['logstore-reads.json'].querierOrIndexGatewaySelector)
                         )
                         .addPanel(
                           $.p99LatencyByPod(
                             'logstore_boltdb_shipper_request_duration_seconds',
                             '{%s operation="Shipper.Query"}' % dashboards['logstore-reads.json'].querierOrIndexGatewaySelector
                           )
                         )
                       ),
  },
}
