(import 'dashboard-utils.libsonnet') {
  local compactor_pod_matcher = if $._config.meta_monitoring.enabled
  then $._config.per_instance_label + '=~"(.*compactor.*|%s-backend.*|logstore-single-binary)"' % $._config.ssd.pod_prefix_matcher
  else if $._config.ssd.enabled then 'container="logstore", ' + $._config.per_instance_label + '=~"%s-read.*"' % $._config.ssd.pod_prefix_matcher else 'container="compactor"',
  local compactor_job_matcher = if $._config.meta_monitoring.enabled
  then '"(.*compactor|%s-backend.*|logstore-single-binary)"' % $._config.ssd.pod_prefix_matcher
  else if $._config.ssd.enabled then '%s-backend' % $._config.ssd.pod_prefix_matcher else 'compactor',
  acmeDashboards+::
    {
      'logstore-retention.json':
        ($.dashboard('Logstore / Retention', uid='retention'))
        .addCluster()
        .addNamespace()
        .addTag()
        .addLog()
        .addRow(
          $.row('Resource Usage')
          .addPanel(
            $.CPUUsagePanel('CPU', compactor_pod_matcher),
          )
          .addPanel(
            $.memoryWorkingSetPanel('Memory (workingset)', compactor_pod_matcher),
          )
          .addPanel(
            $.goHeapInUsePanel('Memory (go heap inuse)', compactor_job_matcher),
          )

        )
        .addRow(
          $.row('Compaction')
          .addPanel(
            $.fromNowPanel('Last Compact Tables Operation Success', 'logstore_boltdb_shipper_compact_tables_operation_last_successful_run_timestamp_seconds')
          )
          .addPanel(
            $.newQueryPanel('Compact Tables Operations Duration', 's') +
            $.queryPanel(['logstore_boltdb_shipper_compact_tables_operation_duration_seconds{%s}' % $.namespaceMatcher()], ['duration']),
          )
        )
        .addRow(
          $.row('')
          .addPanel(
            $.newQueryPanel('Number of times Tables were skipped during Compaction') +
            $.queryPanel(['sum(logstore_compactor_locked_table_successive_compaction_skips{%s})' % $.namespaceMatcher()], ['{{table_name}}']),
          )
          .addPanel(
            $.newQueryPanel('Compact Tables Operations Per Status') +
            $.queryPanel(['sum by (status)(rate(logstore_boltdb_shipper_compact_tables_operation_total{%s}[$__rate_interval]))' % $.namespaceMatcher()], ['{{success}}']),
          )
        )
        .addRow(
          $.row('Retention')
          .addPanel(
            $.fromNowPanel('Last Mark Operation Success', 'logstore_compactor_apply_retention_last_successful_run_timestamp_seconds')
          )
          .addPanel(
            $.newQueryPanel('Mark Operations Duration', 's') +
            $.queryPanel(['logstore_compactor_apply_retention_operation_duration_seconds{%s}' % $.namespaceMatcher()], ['duration']),
          )
          .addPanel(
            $.newQueryPanel('Mark Operations Per Status') +
            $.queryPanel(['sum by (status)(rate(logstore_compactor_apply_retention_operation_total{%s}[$__rate_interval]))' % $.namespaceMatcher()], ['{{success}}']),
          )
        )
        .addRow(
          $.row('Per Table Marker')
          .addPanel(
            $.newQueryPanel('Processed Tables Per Action') +
            $.queryPanel(['count by(action)(logstore_boltdb_shipper_retention_marker_table_processed_total{%s})' % $.namespaceMatcher()], ['{{action}}']) +
            $.withStacking,
          )
          .addPanel(
            $.newQueryPanel('Modified Tables') +
            $.queryPanel(['count by(table,action)(logstore_boltdb_shipper_retention_marker_table_processed_total{%s , action=~"modified|deleted"})' % $.namespaceMatcher()], ['{{table}}-{{action}}']) +
            $.withStacking,
          )
          .addPanel(
            $.newQueryPanel('Marks Creation Rate Per Table') +
            $.queryPanel(['sum by (table)(rate(logstore_boltdb_shipper_retention_marker_count_total{%s}[$__rate_interval])) >0' % $.namespaceMatcher()], ['{{table}}']) +
            $.withStacking,
          )
        )
        .addRow(
          $.row('')
          .addPanel(
            $.newQueryPanel('Marked Chunks (24h)') +
            $.statPanel('sum (increase(logstore_boltdb_shipper_retention_marker_count_total{%s}[24h]))' % $.namespaceMatcher(), 'short')
          )
          .addPanel(
            $.newQueryPanel('Mark Table Latency') +
            $.latencyPanel('logstore_boltdb_shipper_retention_marker_table_processed_duration_seconds', '{%s}' % $.namespaceMatcher())
          )
        )
        .addRow(
          $.row('Sweeper')
          .addPanel(
            $.newQueryPanel('Delete Chunks (24h)') +
            $.statPanel('sum (increase(logstore_boltdb_shipper_retention_sweeper_chunk_deleted_duration_seconds_count{%s}[24h]))' % $.namespaceMatcher(), 'short')
          )
          .addPanel(
            $.newQueryPanel('Delete Latency') +
            $.latencyPanel('logstore_boltdb_shipper_retention_sweeper_chunk_deleted_duration_seconds', '{%s}' % $.namespaceMatcher())
          )
        )
        .addRow(
          $.row('')
          .addPanel(
            $.newQueryPanel('Sweeper Lag', 's') +
            $.queryPanel(['time() - (logstore_boltdb_shipper_retention_sweeper_marker_file_processing_current_time{%s} > 0)' % $.namespaceMatcher()], ['lag']),
          )
          .addPanel(
            $.newQueryPanel('Marks Files to Process') +
            $.queryPanel(['sum(logstore_boltdb_shipper_retention_sweeper_marker_files_current{%s})' % $.namespaceMatcher()], ['count']),
          )
          .addPanel(
            $.newQueryPanel('Delete Rate Per Status') +
            $.queryPanel(['sum by (status)(rate(logstore_boltdb_shipper_retention_sweeper_chunk_deleted_duration_seconds_count{%s}[$__rate_interval]))' % $.namespaceMatcher()], ['{{status}}']),
          )
        )
        .addRow(
          $.row('Logs')
          .addPanel(
            $.logPanel('Compactor Logs', '{%s}' % $.jobMatcher(compactor_job_matcher)),
          )
        ),
    },
}
