local k = import 'example.com/acme/jsonnet-libs/ksonnet-util/kausal.libsonnet';
local tanka = import 'example.com/acme/jsonnet-libs/tanka-util/main.libsonnet';
local configMap = k.core.v1.configMap;

local spec = (import './spec.json').spec;

{
  _config+:: {
    namespace: spec.namespace,
  },

  logstoreNamespace: k.core.v1.namespace.new('logstore'),

  gelLicenseSecret: k.core.v1.secret.new('gel-license', {}, type='Opaque')
                    + k.core.v1.secret.withStringData({
                      'license.jwt': importstr '../../secrets/gel.jwt',
                    })
                    + k.core.v1.secret.metadata.withNamespace('logstore'),
  local acmeCloudCredentials = import '../../secrets/acme-cloud-credentials.json',
  acmeCloudMetricsCredentials: k.core.v1.secret.new('acme-cloud-metrics-credentials', {}, type='Opaque')
                                  + k.core.v1.secret.withStringData({
                                    username: '%d' % acmeCloudCredentials.metrics.username,
                                    password: acmeCloudCredentials.metrics.password,
                                  })
                                  + k.core.v1.secret.metadata.withNamespace('logstore'),
  acmeCloudLogsCredentials: k.core.v1.secret.new('acme-cloud-logs-credentials', {}, type='Opaque')
                               + k.core.v1.secret.withStringData({
                                 username: '%d' % acmeCloudCredentials.logs.username,
                                 password: acmeCloudCredentials.logs.password,
                               })
                               + k.core.v1.secret.metadata.withNamespace('logstore'),


}
