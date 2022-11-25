local k = import 'ksonnet-util/kausal.libsonnet',
      statefulSet = k.apps.v1.statefulSet;

{
  _config+:: {
    headless_service_name: error 'must provide a name for the headless memberlist service under $._config.headless_service_name',
    http_listen_port: error 'must provide http listen port under $._config.http_listen_port',
    logstore: error 'must provide logstore config under $._config.logstore',

    commonArgs: {
      'config.file': '/etc/logstore/config.yaml',
    },

    config_hash_mixin:
      statefulSet.mixin.spec.template.metadata.withAnnotationsMixin({
        config_hash: std.md5(std.toString($._config.logstore)),
      }),
  },

  local configMap = k.core.v1.configMap,

  config_file:
    configMap.new('logstore') +
    configMap.withData({
      'config.yaml': k.util.manifestYaml($._config.logstore),
    }),

  local deployment = k.apps.v1.deployment,

}
