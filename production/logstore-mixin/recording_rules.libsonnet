local utils = import 'mixin-utils/utils.libsonnet';

{
  prometheusRules+:: {
    groups+: [{
      name: 'logstore_rules',
      rules:
        utils.histogramRules('logstore_request_duration_seconds', [$._config.per_cluster_label, 'job'], $._config.recording_rules_range_interval) +
        utils.histogramRules('logstore_request_duration_seconds', [$._config.per_cluster_label, 'job', 'route'], $._config.recording_rules_range_interval) +
        utils.histogramRules('logstore_request_duration_seconds', [$._config.per_cluster_label, 'namespace', 'job', 'route'], $._config.recording_rules_range_interval),
    }],
  },
}
