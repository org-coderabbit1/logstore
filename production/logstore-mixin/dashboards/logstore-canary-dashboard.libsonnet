local vendor_config = import 'example.com/acme/metricstore/operations/metricstore-mixin/config.libsonnet';
local vendor_utils = import 'example.com/acme/metricstore/operations/metricstore-mixin/dashboards/dashboard-utils.libsonnet';
local g = import 'acme-builder/acme.libsonnet';
local acme = import 'grafonnet/acme.libsonnet';

{
  _config+:: {
    canary+: {
      enabled: false,
    },
  },
  acmeDashboards+: if !$._config.canary.enabled then {} else {
    local dashboard = (
      vendor_utils {
        _config:: vendor_config._config + $._config {
          product: 'Logstore',
          dashboard_prefix: 'Logstore / ',
          tags: ['logstore'],
        },
      }
    ),
    'logstore-canary.json':
      // The dashboard() function automatically adds the "Logstore / " prefix to the dashboard title.
      // This logic is inherited from metricstore-mixin.
      dashboard.dashboard('Canary')
      // We can't make use of simplified template selectors from the logstore dashboard utils until we port the corestore dashboard utils panel/grid functionality.
      .addTemplate('cluster', 'logstore_build_info', $._config.per_cluster_label)
      .addTemplate('namespace', 'logstore_build_info{' + $._config.per_cluster_label + '=~"$cluster"}', 'namespace')
      + {
        // This dashboard uses the new grid system in order to place panels (using gridPos).
        // Because of this we can't use the mixin's addRow() and addPanel().
        schemaVersion: 27,
        rows: null,
        // ugly hack, copy pasta the tag/link
        // code from the logstore-mixin
        tags: $._config.tags,
        links: [
          {
            asDropdown: true,
            icon: 'external link',
            includeVars: true,
            keepTime: true,
            tags: $._config.tags,
            targetBlank: false,
            title: 'Logstore Dashboards',
            type: 'dashboards',
          },
        ],
        panels: [
          // grid row 1
          dashboard.panel('Canary Entries Total') +
          dashboard.newStatPanel('sum(count(logstore_canary_entries_total{' + $._config.per_cluster_label + '=~"$cluster", ' + $._config.per_namespace_label + '=~"$namespace"}))', unit='short') +
          { gridPos: { h: 4, w: 3, x: 0, y: 0 } },

          dashboard.panel('Canary Logs Total') +
          dashboard.newStatPanel('sum(increase(logstore_canary_entries_total{' + $._config.per_cluster_label + '=~"$cluster",' + $._config.per_namespace_label + '=~"$namespace"}[$__range]))', unit='short') +
          { gridPos: { h: 4, w: 3, x: 3, y: 0 } },

          dashboard.panel('Missing') +
          dashboard.newStatPanel('sum(increase(logstore_canary_missing_entries_total{' + $._config.per_cluster_label + '=~"$cluster",' + $._config.per_namespace_label + '=~"$namespace"}[$__range]))', unit='short') +
          { gridPos: { h: 4, w: 3, x: 6, y: 0 } },

          dashboard.panel('Spotcheck Missing') +
          dashboard.newStatPanel('sum(increase(logstore_canary_spot_check_missing_entries_total{' + $._config.per_cluster_label + '=~"$cluster",' + $._config.per_namespace_label + '=~"$namespace"}[$__range]))', unit='short') +
          { gridPos: { h: 4, w: 3, x: 9, y: 0 } },

          // grid row 2
          dashboard.panel('Spotcheck Total') +
          dashboard.newStatPanel('sum(increase(logstore_canary_spot_check_entries_total{' + $._config.per_cluster_label + '=~"$cluster",' + $._config.per_namespace_label + '=~"$namespace"}[$__range]))', unit='short') +
          { gridPos: { h: 4, w: 3, x: 0, y: 4 } },

          dashboard.panel('Metric Test Error %') +
          dashboard.newStatPanel('((sum(logstore_canary_metric_test_expected{' + $._config.per_cluster_label + '=~"$cluster",' + $._config.per_namespace_label + '=~"$namespace"}) - sum(logstore_canary_metric_test_actual{' + $._config.per_cluster_label + '=~"$cluster",' + $._config.per_namespace_label + '=~"$namespace"}))/(sum(logstore_canary_metric_test_actual{' + $._config.per_cluster_label + '=~"$cluster",' + $._config.per_namespace_label + '=~"$namespace"}))) * 100') +
          { gridPos: { h: 4, w: 3, x: 3, y: 4 } },

          dashboard.panel('Missing %') +
          dashboard.newStatPanel('(sum(increase(logstore_canary_missing_entries_total{' + $._config.per_cluster_label + '=~"$cluster",' + $._config.per_namespace_label + '=~"$namespace"}[$__range]))/sum(increase(logstore_canary_entries_total{' + $._config.per_cluster_label + '=~"$cluster",' + $._config.per_namespace_label + '=~"$namespace"}[$__range])))*100') +
          { gridPos: { h: 4, w: 3, x: 6, y: 4 } },

          dashboard.panel('Spotcheck Missing %') +
          dashboard.newStatPanel('(sum(increase(logstore_canary_spot_check_missing_entries_total{' + $._config.per_cluster_label + '=~"$cluster",' + $._config.per_namespace_label + '=~"$namespace"}[$__range]))/sum(increase(logstore_canary_spot_check_entries_total{' + $._config.per_cluster_label + '=~"$cluster",' + $._config.per_namespace_label + '=~"$namespace"}[$__range]))) * 100') +
          { gridPos: { h: 4, w: 3, x: 9, y: 4 } },

          // grid row 3
          dashboard.panel('Metric Test Expected') +
          dashboard.newStatPanel('sum(logstore_canary_metric_test_expected{' + $._config.per_cluster_label + '=~"$cluster",' + $._config.per_namespace_label + '=~"$namespace"})', unit='short') +
          { gridPos: { h: 4, w: 3, x: 0, y: 8 } },

          dashboard.panel('Metric Test Actual') +
          dashboard.newStatPanel('sum(logstore_canary_metric_test_actual{' + $._config.per_cluster_label + '=~"$cluster",' + $._config.per_namespace_label + '=~"$namespace"})', unit='short') +
          { gridPos: { h: 4, w: 3, x: 3, y: 8 } },

          dashboard.panel('Websocket Missing') +
          dashboard.newStatPanel('sum(increase(logstore_canary_websocket_missing_entries_total{' + $._config.per_cluster_label + '=~"$cluster",' + $._config.per_namespace_label + '=~"$namespace"}[$__range]))', unit='short') +
          { gridPos: { h: 4, w: 3, x: 6, y: 8 } },

          dashboard.panel('Websocket Missing %') +
          dashboard.newStatPanel('(sum(increase(logstore_canary_websocket_missing_entries_total{' + $._config.per_cluster_label + '=~"$cluster",' + $._config.per_namespace_label + '=~"$namespace"}[$__range]))/sum(increase(logstore_canary_entries_total{' + $._config.per_cluster_label + '=~"$cluster",' + $._config.per_namespace_label + '=~"$namespace"}[$__range])))*100') +
          { gridPos: { h: 4, w: 3, x: 9, y: 8 } },
          // end of grid

          dashboard.panel('Log Write to read Latency Percentiles') +
          dashboard.queryPanel([
            'histogram_quantile(0.95, sum(rate(logstore_canary_response_latency_seconds_bucket{' + $._config.per_cluster_label + '=~"$cluster",' + $._config.per_namespace_label + '=~"$namespace"}[$__rate_interval])) by (le))',
            'histogram_quantile(0.50, sum(rate(logstore_canary_response_latency_seconds_bucket{' + $._config.per_cluster_label + '=~"$cluster",' + $._config.per_namespace_label + '=~"$namespace"}[$__rate_interval])) by (le))',
          ], ['p95', 'p50']) +
          { gridPos: { h: 6, w: 12, x: 12, y: 0 } },

          acme.heatmapPanel.new(
            'Log Write to Read Latency',
            datasource='$datasource',
            tooltip_showHistogram=true,
            color_colorScheme='interpolateReds',
            legend_show=false,
          ).addTargets(
            [
              acme.prometheus.target(
                'sum(rate(logstore_canary_response_latency_seconds_bucket{' + $._config.per_cluster_label + '=~"$cluster",' + $._config.per_namespace_label + '=~"$namespace"}[$__rate_interval])) by (le)',
                legendFormat='{{le}}',
                format='heatmap',
              ),
            ],
          ) +
          { gridPos: { h: 6, w: 12, x: 12, y: 12 } },

          dashboard.panel('Spot Check Query') +
          dashboard.queryPanel([
            'histogram_quantile(0.99, sum(rate(logstore_canary_spot_check_request_duration_seconds_bucket{' + $._config.per_cluster_label + '=~"$cluster",' + $._config.per_namespace_label + '=~"$namespace"}[$__rate_interval])) by (le))',
            'histogram_quantile(0.50, sum(rate(logstore_canary_spot_check_request_duration_seconds_bucket{' + $._config.per_cluster_label + '=~"$cluster",' + $._config.per_namespace_label + '=~"$namespace"}[$__rate_interval])) by (le))',
          ], ['p99', 'p95']) +
          { gridPos: { h: 6, w: 12, x: 0, y: 14 } },

          dashboard.panel('Metric Test Query') +
          dashboard.queryPanel([
            'histogram_quantile(0.99, sum(rate(logstore_canary_metric_test_request_duration_seconds_bucket{' + $._config.per_cluster_label + '=~"$cluster",' + $._config.per_namespace_label + '=~"$namespace"}[15m])) by (le))',
            'histogram_quantile(0.50, sum(rate(logstore_canary_metric_test_request_duration_seconds_bucket{' + $._config.per_cluster_label + '=~"$cluster",' + $._config.per_namespace_label + '=~"$namespace"}[15m])) by (le))',
          ], ['p99', 'p95'],) +
          { gridPos: { h: 6, w: 12, x: 12, y: 14 } },

          dashboard.panel('Spot Check Missing %') +
          dashboard.queryPanel('topk(20, (sum by (' + $._config.per_cluster_label + ', pod) (increase(logstore_canary_spot_check_missing_entries_total{' + $._config.per_cluster_label + '=~"$cluster",' + $._config.per_namespace_label + '=~"$namespace"}[$__rate_interval]))/sum by (' + $._config.per_cluster_label + ', pod) (increase(logstore_canary_spot_check_entries_total{' + $._config.per_cluster_label + '=~"$cluster",' + $._config.per_namespace_label + '=~"$namespace"}[$__rate_interval])) * 100)) > 0', '') +
          { gridPos: { h: 6, w: 12, x: 0, y: 20 } },

          g.panel('Missing logs') +
          g.queryPanel('topk(20,(sum by (' + $._config.per_cluster_label + ', pod)(increase(logstore_canary_missing_entries_total{' + $._config.per_cluster_label + '=~"$cluster",' + $._config.per_namespace_label + '=~"$namespace"}[$__rate_interval]))/sum by (' + $._config.per_cluster_label + ', pod)(increase(logstore_canary_entries_total{' + $._config.per_cluster_label + '=~"$cluster",' + $._config.per_namespace_label + '=~"$namespace"}[$__rate_interval])))*100) > 0', 'Missing {{ ' + $._config.per_cluster_label + ' }} {{ pod }}') +
          { gridPos: { h: 6, w: 12, x: 12, y: 20 } },

        ],
      },
  },
}
