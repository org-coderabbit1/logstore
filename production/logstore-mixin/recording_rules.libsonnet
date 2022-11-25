local utils = import 'mixin-utils/utils.libsonnet';

{
  prometheusRules+:: {
    groups+: [{
      name: 'logstore_rules',
      rules:
        utils.histogramRules('logstore_request_duration_seconds', ['job']) +
        utils.histogramRules('logstore_request_duration_seconds', ['job', 'route']) +
        utils.histogramRules('logstore_request_duration_seconds', ['namespace', 'job', 'route']),
    }],
  },
}
