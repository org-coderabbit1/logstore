local acme = import 'grafonnet/acme.libsonnet';
local utils = import 'mixin-utils/utils.libsonnet';

(import 'dashboard-utils.libsonnet') {
  acmeDashboards+: {
    local dashboards = self,
    'logstore-chunks.json': {
                          local cfg = self,
                          labelsSelector:: $._config.per_cluster_label + '="$cluster", job=~"$namespace/%s"' % (
                            if $._config.meta_monitoring.enabled
                            then '(.*ingester.*|%s-write|logstore-single-binary)' % $._config.ssd.pod_prefix_matcher
                            else if $._config.ssd.enabled then '%s-write' % $._config.ssd.pod_prefix_matcher else '(.*ingester.*)'
                          ),
                        } +
                        $.dashboard('Logstore / Chunks', uid='chunks')
                        .addCluster()
                        .addNamespace()
                        .addTag()
                        .addRow(
                          $.row('Active Series / Chunks')
                          .addPanel(
                            $.newQueryPanel('Series') +
                            $.queryPanel('sum(logstore_ingester_memory_chunks{%s})' % dashboards['logstore-chunks.json'].labelsSelector, 'series'),
                          )
                          .addPanel(
                            $.newQueryPanel('Chunks per series') +
                            $.queryPanel(
                              'sum(logstore_ingester_memory_chunks{%s}) / sum(logstore_ingester_memory_streams{%s})' % [
                                dashboards['logstore-chunks.json'].labelsSelector,
                                dashboards['logstore-chunks.json'].labelsSelector,
                              ],
                              'chunks'
                            ),
                          )
                        )
                        .addRow(
                          $.row('Flush Stats')
                          .addPanel(
                            $.newQueryPanel('Utilization', 'percentunit') +
                            $.latencyPanel('logstore_ingester_chunk_utilization', '{%s}' % dashboards['logstore-chunks.json'].labelsSelector, multiplier='1'),
                          )
                          .addPanel(
                            $.newQueryPanel('Age') +
                            $.latencyPanel('logstore_ingester_chunk_age_seconds', '{%s}' % dashboards['logstore-chunks.json'].labelsSelector),
                          ),
                        )
                        .addRow(
                          $.row('Flush Stats')
                          .addPanel(
                            $.newQueryPanel('Log Entries Per Chunk', 'short') +
                            $.latencyPanel('logstore_ingester_chunk_entries', '{%s}' % dashboards['logstore-chunks.json'].labelsSelector, multiplier='1'),
                          )
                          .addPanel(
                            $.newQueryPanel('Index Entries Per Chunk') +
                            $.queryPanel(
                              'sum(rate(logstore_chunk_store_index_entries_per_chunk_sum{%s}[$__rate_interval])) / sum(rate(logstore_chunk_store_index_entries_per_chunk_count{%s}[$__rate_interval]))' % [
                                dashboards['logstore-chunks.json'].labelsSelector,
                                dashboards['logstore-chunks.json'].labelsSelector,
                              ],
                              'Index Entries'
                            ),
                          ),
                        )
                        .addRow(
                          $.row('Flush Stats')
                          .addPanel(
                            $.newQueryPanel('Queue Length') +
                            $.queryPanel('logstore_ingester_flush_queue_length{%(label)s} or corestore_ingester_flush_queue_length{%(label)s}' % { label: dashboards['logstore-chunks.json'].labelsSelector }, '{{pod}}'),
                          )
                          .addPanel(
                            $.newQueryPanel('Flush Rate') +
                            $.newQpsPanel('logstore_ingester_chunk_age_seconds_count{%s}' % dashboards['logstore-chunks.json'].labelsSelector,),
                          ),
                        )
                        .addRow(
                          $.row('Flush Stats')
                          .addPanel(
                            $.newQueryPanel('Chunks Flushed/Second') +
                            $.queryPanel('sum(rate(logstore_ingester_chunks_flushed_total{%s}[$__rate_interval]))' % dashboards['logstore-chunks.json'].labelsSelector, '{{pod}}'),
                          )
                          .addPanel(
                            $.newQueryPanel('Chunk Flush Reason') +
                            $.queryPanel('sum by (reason) (rate(logstore_ingester_chunks_flushed_total{%s}[$__rate_interval])) / ignoring(reason) group_left sum(rate(logstore_ingester_chunks_flushed_total{%s}[$__rate_interval]))' % [dashboards['logstore-chunks.json'].labelsSelector, dashboards['logstore-chunks.json'].labelsSelector], '{{reason}}') + {
                              stack: true,
                              yaxes: [
                                { format: 'short', label: null, logBase: 1, max: 1, min: 0, show: true },
                                { format: 'short', label: null, logBase: 1, max: 1, min: null, show: false },
                              ],
                            },
                          ),
                        )
                        .addRow(
                          $.row('Utilization')
                          .addPanel(
                            acme.heatmapPanel.new(
                              'Chunk Utilization',
                              datasource='$datasource',
                              yAxis_format='percentunit',
                              tooltip_showHistogram=true,
                              color_colorScheme='interpolateSpectral',
                              dataFormat='tsbuckets',
                              yAxis_decimals=0,
                              legend_show=true,
                            ).addTargets(
                              [
                                acme.prometheus.target(
                                  'sum by (le) (rate(logstore_ingester_chunk_utilization_bucket{%s}[$__rate_interval]))' % dashboards['logstore-chunks.json'].labelsSelector,
                                  legendFormat='{{le}}',
                                  format='heatmap',
                                ),
                              ],
                            )
                          )
                        )
                        .addRow(
                          $.row('Utilization')
                          .addPanel(
                            acme.heatmapPanel.new(
                              'Chunk Size Bytes',
                              datasource='$datasource',
                              yAxis_format='bytes',
                              tooltip_showHistogram=true,
                              color_colorScheme='interpolateSpectral',
                              dataFormat='tsbuckets',
                              yAxis_decimals=0,
                              // tooltipDecimals=3,
                              // span=3,
                              legend_show=true,
                            ).addTargets(
                              [
                                acme.prometheus.target(
                                  'sum(rate(logstore_ingester_chunk_size_bytes_bucket{%s}[$__rate_interval])) by (le)' % dashboards['logstore-chunks.json'].labelsSelector,
                                  legendFormat='{{le}}',
                                  format='heatmap',
                                ),
                              ],
                            )
                          )
                        )
                        .addRow(
                          $.row('Utilization')
                          .addPanel(
                            $.newQueryPanel('Chunk Size Quantiles', 'bytes') +
                            $.queryPanel(
                              [
                                'histogram_quantile(0.99, sum(rate(logstore_ingester_chunk_size_bytes_bucket{%s}[$__rate_interval])) by (le))' % dashboards['logstore-chunks.json'].labelsSelector,
                                'histogram_quantile(0.90, sum(rate(logstore_ingester_chunk_size_bytes_bucket{%s}[$__rate_interval])) by (le))' % dashboards['logstore-chunks.json'].labelsSelector,
                                'histogram_quantile(0.50, sum(rate(logstore_ingester_chunk_size_bytes_bucket{%s}[$__rate_interval])) by (le))' % dashboards['logstore-chunks.json'].labelsSelector,
                              ],
                              [
                                'p99',
                                'p90',
                                'p50',
                              ],
                            ),
                          )
                        )
                        .addRow(
                          $.row('Duration')
                          .addPanel(
                            $.newQueryPanel('Chunk Duration hours (end-start)') +
                            $.queryPanel(
                              [
                                'histogram_quantile(0.5, sum(rate(logstore_ingester_chunk_bounds_hours_bucket{%s}[$__rate_interval])) by (le))' % dashboards['logstore-chunks.json'].labelsSelector,
                                'histogram_quantile(0.99, sum(rate(logstore_ingester_chunk_bounds_hours_bucket{%s}[$__rate_interval])) by (le))' % dashboards['logstore-chunks.json'].labelsSelector,
                                'sum(rate(logstore_ingester_chunk_bounds_hours_sum{%s}[$__rate_interval])) / sum(rate(logstore_ingester_chunk_bounds_hours_count{%s}[$__rate_interval]))' % [
                                  dashboards['logstore-chunks.json'].labelsSelector,
                                  dashboards['logstore-chunks.json'].labelsSelector,
                                ],
                              ],
                              [
                                'p50',
                                'p99',
                                'avg',
                              ],
                            ),
                          )
                        ),
  },
}
