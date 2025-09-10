local logstoreOperational = (import './dashboard-logstore-operational.json');
local utils = import 'mixin-utils/utils.libsonnet';

(import 'dashboard-utils.libsonnet') {
  acmeDashboards+: {
    local dashboards = self,

    'logstore-operational.json': {
                               local cfg = self,

                               showAnnotations:: true,
                               showLinks:: true,
                               showMultiCluster:: true,

                               hiddenRows:: [
                                 'Cassandra',
                                 if $._config.ssd.enabled then 'Ingester',
                                 if !$._config.ssd.enabled then 'Backend Path',
                                 if !$._config.operational.memcached then 'Memcached',
                                 if !$._config.operational.consul then 'Consul',
                                 if !$._config.operational.bigTable then 'Big Table',
                                 if !$._config.operational.dynamo then 'Dynamo',
                                 if !$._config.operational.gcs then 'GCS',
                                 if !$._config.operational.s3 then 'S3',
                                 if !$._config.operational.azureBlob then 'Azure Blob',
                                 if !$._config.operational.boltDB then 'BoltDB Shipper',
                               ],

                               hiddenPanels:: if $._config.promtail.enabled then [] else [
                                 'Bad Words',
                               ],

                               jobMatchers:: {
                                 corestoregateway: [utils.selector.re('job', '($namespace)/corestore-gw(-internal)?')],
                                 distributor: if $._config.meta_monitoring.enabled
                                 then [utils.selector.re('job', '($namespace)/(distributor|%s-write|logstore-single-binary)' % $._config.ssd.pod_prefix_matcher)]
                                 else [utils.selector.re('job', '($namespace)/%s' % (if $._config.ssd.enabled then '%s-write' % $._config.ssd.pod_prefix_matcher else 'distributor'))],
                                 ingester: if $._config.meta_monitoring.enabled
                                 then [utils.selector.re('job', '($namespace)/(partition-ingester.*|ingester.*|%s-write|logstore-single-binary)' % $._config.ssd.pod_prefix_matcher)]
                                 else [utils.selector.re('job', '($namespace)/%s' % (if $._config.ssd.enabled then '%s-write' % $._config.ssd.pod_prefix_matcher else '(ingester.*|partition-ingester.*)'))],
                                 querier: if $._config.meta_monitoring.enabled
                                 then [utils.selector.re('job', '($namespace)/(querier|%s-read|logstore-single-binary)' % $._config.ssd.pod_prefix_matcher)]
                                 else [utils.selector.re('job', '($namespace)/%s' % (if $._config.ssd.enabled then '%s-read' % $._config.ssd.pod_prefix_matcher else 'querier'))],
                                 queryFrontend: if $._config.meta_monitoring.enabled
                                 then [utils.selector.re('job', '($namespace)/(query-frontend|%s-read|logstore-single-binary)' % $._config.ssd.pod_prefix_matcher)]
                                 else [utils.selector.re('job', '($namespace)/%s' % (if $._config.ssd.enabled then '%s-read' % $._config.ssd.pod_prefix_matcher else 'query-frontend'))],
                                 backend: [utils.selector.re('job', '($namespace)/%s-backend' % $._config.ssd.pod_prefix_matcher)],
                               },

                               podMatchers:: {
                                 corestoregateway: [utils.selector.re('pod', 'corestore-gw')],
                                 distributor: if $._config.meta_monitoring.enabled
                                 then [utils.selector.re('pod', '(distributor|%s-write|logstore-single-binary)' % $._config.ssd.pod_prefix_matcher)]
                                 else [utils.selector.re('pod', '%s' % (if $._config.ssd.enabled then '%s-write.*' % $._config.ssd.pod_prefix_matcher else 'distributor.*'))],
                                 ingester: if $._config.meta_monitoring.enabled
                                 then [utils.selector.re('pod', '(partition-ingester.*|ingester.*|%s-write|logstore-single-binary)' % $._config.ssd.pod_prefix_matcher)]
                                 else [utils.selector.re('pod', '%s' % (if $._config.ssd.enabled then '%s-write.*' % $._config.ssd.pod_prefix_matcher else '(ingester.*|partition-ingester.*)'))],
                                 querier: if $._config.meta_monitoring.enabled
                                 then [utils.selector.re('pod', '(querier|%s-read|logstore-single-binary)' % $._config.ssd.pod_prefix_matcher)]
                                 else [utils.selector.re('pod', '%s' % (if $._config.ssd.enabled then '%s-read.*' % $._config.ssd.pod_prefix_matcher else 'querier.*'))],
                                 backend: [utils.selector.re('pod', '%s-backend.*' % $._config.ssd.pod_prefix_matcher)],
                               },
                             }
                             + logstoreOperational + {
                               annotations:
                                 if dashboards['logstore-operational.json'].showAnnotations
                                 then super.annotations
                                 else {},

                               links:
                                 if dashboards['logstore-operational.json'].showLinks then
                                   super.links
                                 else [],

                               local matcherStr(matcherId, matcher='job', sep=',') =
                                 if matcher == 'job' then
                                   if std.length(dashboards['logstore-operational.json'].jobMatchers[matcherId]) > 0 then
                                     std.join(',', ['%(label)s%(op)s"%(value)s"' % matcher for matcher in dashboards['logstore-operational.json'].jobMatchers[matcherId]]) + sep
                                   else error 'no job matchers'
                                 else if matcher == 'pod' then
                                   if std.length(dashboards['logstore-operational.json'].podMatchers[matcherId]) > 0 then
                                     std.join(',', ['%(label)s%(op)s"%(value)s"' % matcher for matcher in dashboards['logstore-operational.json'].podMatchers[matcherId]]) + sep
                                   else error 'no pod matchers'
                                 else error 'matcher must be either job or container',

                               local replaceClusterMatchers(expr) =
                                 if dashboards['logstore-operational.json'].showMultiCluster
                                 // Replace the recording rules cluster label with the per-cluster label
                                 then std.strReplace(
                                   // Replace the cluster label for equality matchers with the per-cluster label
                                   std.strReplace(
                                     // Replace the cluster label for regex matchers with the per-cluster label
                                     std.strReplace(
                                       expr,
                                       'cluster=~"$cluster"',
                                       $._config.per_cluster_label + '=~"$cluster"'
                                     ),
                                     'cluster="$cluster"',
                                     $._config.per_cluster_label + '="$cluster"'
                                   ),
                                   'cluster_job',
                                   $._config.per_cluster_label + '_job'
                                 )
                                 else
                                   std.strReplace(
                                     std.strReplace(
                                       std.strReplace(
                                         expr,
                                         ', ' + $._config.per_cluster_label + '="$cluster"',
                                         ''
                                       ),
                                       ', ' + $._config.per_cluster_label + '=~"$cluster"',
                                       ''
                                     ),
                                     $._config.per_cluster_label + '="$cluster",',
                                     ''
                                   ),

                               local replaceBackendMatchers(expr) =
                                 std.strReplace(
                                   std.strReplace(
                                     std.strReplace(
                                       expr,
                                       $._config.per_instance_label + '=~"backend.*"',
                                       matcherStr('backend', matcher='pod', sep='')
                                     ),
                                     'job="$namespace/backend",',
                                     matcherStr('backend')
                                   ),
                                   'job="$namespace/backend"',
                                   std.rstripChars(matcherStr('backend'), ',')
                                 ),

                               local replaceQuerierMatchers(expr) =
                                 std.strReplace(
                                   std.strReplace(
                                     std.strReplace(
                                       expr,
                                       $._config.per_instance_label + '=~"querier.*"',
                                       matcherStr('querier', matcher='pod', sep='')
                                     ),
                                     'job="$namespace/querier",',
                                     matcherStr('querier')
                                   ),
                                   'job="$namespace/querier"',
                                   std.rstripChars(matcherStr('querier'), ',')
                                 ),

                               local replaceMatchers(expr) =
                                 std.strReplace(
                                   std.strReplace(
                                     std.strReplace(
                                       std.strReplace(
                                         std.strReplace(
                                           std.strReplace(
                                             std.strReplace(
                                               std.strReplace(
                                                 std.strReplace(
                                                   std.strReplace(
                                                     std.strReplace(
                                                       std.strReplace(
                                                         expr,
                                                         $._config.per_instance_label + '=~"ingester.*"',
                                                         matcherStr('ingester', matcher='pod', sep='')
                                                       ),
                                                       $._config.per_instance_label + '=~"distributor.*"',
                                                       matcherStr('distributor', matcher='pod', sep='')
                                                     ),
                                                     'job="$namespace/corestore-gw",',
                                                     matcherStr('corestoregateway')
                                                   ),
                                                   'job="$namespace/corestore-gw"',
                                                   std.rstripChars(matcherStr('corestoregateway'), ',')
                                                 ),
                                                 'job=~"($namespace)/corestore-gw",',
                                                 matcherStr('corestoregateway')
                                               ),
                                               'job="$namespace/distributor",',
                                               matcherStr('distributor')
                                             ),
                                             'job="$namespace/distributor"',
                                             std.rstripChars(matcherStr('distributor'), ',')
                                           ),
                                           'job=~"($namespace)/distributor",',
                                           matcherStr('distributor')
                                         ),
                                         'job=~"($namespace)/distributor"',
                                         std.rstripChars(matcherStr('distributor'), ',')
                                       ),
                                       'job="$namespace/ingester",',
                                       matcherStr('ingester')
                                     ),
                                     'job="$namespace/ingester"',
                                     std.rstripChars(matcherStr('ingester'), ',')
                                   ),
                                   'job=~"($namespace)/ingester",',
                                   matcherStr('ingester'),
                                 ),

                               local replaceAllMatchers(expr) =
                                 replaceBackendMatchers(
                                   replaceQuerierMatchers(
                                     replaceMatchers(expr)
                                   )
                                 ),

                               local selectDatasource(ds) =
                                 if ds == null || ds == '' then ds
                                 else if ds == '$datasource' then '$datasource'
                                 else '$logstore_datasource',

                               local isRowHidden(row) =
                                 std.member(dashboards['logstore-operational.json'].hiddenRows, row),

                               local isPanelHidden(panelTitle) =
                                 std.member(dashboards['logstore-operational.json'].hiddenPanels, panelTitle),

                               local replaceCorestoreGateway(expr, replacement) = if $._config.internal_components then
                                 expr
                               else
                                 std.strReplace(
                                   expr,
                                   'job=~"$namespace/corestore-gw(-internal)?"',
                                   matcherStr(replacement, matcher='job', sep='')
                                 ),

                               local removeInternalComponents(title, expr) = if (title == 'Queries/Second') then
                                 replaceCorestoreGateway(expr, 'queryFrontend')
                               else if (title == 'Pushes/Second') then
                                 replaceCorestoreGateway(expr, 'distributor')
                               else if (title == 'Push Latency') then
                                 replaceCorestoreGateway(expr, 'distributor')
                               else
                                 replaceAllMatchers(expr),

                               panels: [
                                 p {
                                   datasource: selectDatasource(super.datasource),
                                   targets: if std.objectHas(p, 'targets') then [
                                     e {
                                       expr: removeInternalComponents(p.title, replaceClusterMatchers(e.expr)),
                                     }
                                     for e in p.targets
                                   ] else [],
                                   panels: if std.objectHas(p, 'panels') then [
                                     sp {
                                       datasource: selectDatasource(super.datasource),
                                       targets: if std.objectHas(sp, 'targets') then [
                                         e {
                                           expr: removeInternalComponents(p.title, replaceClusterMatchers(e.expr)),
                                         }
                                         for e in sp.targets
                                       ] else [],
                                       panels: if std.objectHas(sp, 'panels') then [
                                         ssp {
                                           datasource: selectDatasource(super.datasource),
                                           targets: if std.objectHas(ssp, 'targets') then [
                                             e {
                                               expr: removeInternalComponents(p.title, replaceClusterMatchers(e.expr)),
                                             }
                                             for e in ssp.targets
                                           ] else [],
                                         }
                                         for ssp in sp.panels
                                         if !(isPanelHidden(ssp.title))
                                       ] else [],
                                     }
                                     for sp in p.panels
                                     if !(isPanelHidden(sp.title))
                                   ] else [],
                                   title: if !($._config.ssd.enabled && p.type == 'row') then p.title else
                                     if p.title == 'Distributor' then 'Write Path'
                                     else if p.title == 'Querier' then 'Read Path'
                                     else p.title,
                                 }
                                 for p in super.panels
                                 if !(p.type == 'row' && isRowHidden(p.title)) && !(isPanelHidden(p.title))
                               ],
                             } +
                             $.dashboard('Logstore / Operational', uid='operational')
                             // The queries in this dashboard don't make use of the cluster template label selector
                             // but we keep it here to allow selecting a namespace specific to a certain cluster, the
                             // namespace template variable selectors query uses the cluster value.
                             .addLog()
                             .addCluster()
                             .addNamespace()
                             .addTag(),
  },
}
