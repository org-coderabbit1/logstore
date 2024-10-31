local config = import 'config.libsonnet';
local k = import 'ksonnet-util/kausal.libsonnet';

// backwards compatibility with ksonnet
local envVar = if std.objectHasAll(k.core.v1, 'envVar') then k.core.v1.envVar else k.core.v1.container.envType;

config {
  namespace: k.core.v1.namespace.new($._config.namespace),

  local container = k.core.v1.container,

  logstore_canary_args:: {
    labelvalue: '$(POD_NAME)',
  },

  logstore_canary_container::
    container.new('logstore-canary', $._images.logstore_canary) +
    k.util.resourcesRequests('10m', '20Mi') +
    container.withPorts(k.core.v1.containerPort.new(name='http-metrics', port=80)) +
    container.withArgsMixin(k.util.mapToFlags($.logstore_canary_args)) +
    container.withEnv([
      envVar.fromFieldPath('HOSTNAME', 'spec.nodeName'),
      envVar.fromFieldPath('POD_NAME', 'metadata.name'),
    ]),

  local daemonSet = k.apps.v1.daemonSet,

  logstore_canary_daemonset:
    daemonSet.new('logstore-canary', [$.logstore_canary_container]),
}
