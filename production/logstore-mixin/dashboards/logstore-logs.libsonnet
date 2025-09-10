local logstoreLogs = (import './dashboard-logstore-logs.json');
local template = import 'grafonnet/template.libsonnet';


(import 'dashboard-utils.libsonnet') {
  local containerTemplate =
    template.new(
      'container',
      '$datasource',
      'label_values(kube_pod_container_info{' + $._config.per_cluster_label + '="$cluster", ' + $._config.per_namespace_label + '="$namespace"}, container)',
      sort=1,
      multi=true,
    ),

  local levelTemplate =
    template.custom(
      'level',
      'debug,info,warn,error',
      '',
      multi=true,
    ),

  local logTemplate =
    template.text(
      'filter',
      label='LogQL Filter'
    ),

  acmeDashboards+: {
    local dashboards = self,

    'logstore-logs.json': {
                        local cfg = self,

                        showMultiCluster:: true,

                      } + logstoreLogs +
                      $.dashboard('Logstore / Logs', uid='logs')
                      .addCluster()
                      .addNamespace()
                      .addTag()
                      .addLog() +
                      {
                        panels: [
                          p {
                            targets: [
                              e {
                                expr: if dashboards['logstore-logs.json'].showMultiCluster
                                then std.strReplace(super.expr, 'cluster="$cluster"', $._config.per_cluster_label + '="$cluster"')
                                else std.strReplace(super.expr, 'cluster="$cluster", ', ''),
                              }
                              for e in p.targets
                            ],
                          }
                          for p in super.panels
                        ],
                        templating+: {
                          list+: [
                            containerTemplate,
                            levelTemplate,
                            logTemplate,
                          ],
                        },
                      },
  },
}
