local k = import 'example.com/acme/jsonnet-libs/ksonnet-util/kausal.libsonnet';
local tanka = import 'example.com/acme/jsonnet-libs/tanka-util/main.libsonnet';

local acme = import 'acme/acme.libsonnet';
local envVar = if std.objectHasAll(k.core.v1, 'envVar') then k.core.v1.envVar else k.core.v1.container.envType;
local helm = tanka.helm.new(std.thisFile);

local spec = (import './spec.json').spec;

{
  local prometheusServerName = self.prometheus.service_prometheus_kube_prometheus_prometheus.metadata.name,
  local prometheusUrl = 'http://%s:9090' % prometheusServerName,

  local logstoreGatewayHost = self.logstore.service_enterprise_logs_gateway.metadata.name,
  local logstoreGatewayUrl = 'http://%s' % logstoreGatewayHost,

  local licenseClusterName = 'enterprise-logs-test-fixture',
  local provisionedSecretPrefix = 'provisioned-secret',
  local adminTokenSecret = 'gel-admin-token',

  local tenant = 'logstore',

  _config+:: {
    clusterName: licenseClusterName,
    namespace: spec.namespace,
    adminTokenSecret: adminTokenSecret,
    adminApiUrl: logstoreGatewayUrl,
  },

  logstore: helm.template($._config.clusterName, '../../../../../production/helm/logstore', {
    namespace: $._config.namespace,
    values: {
      logstore+: {},
      enterprise+: {
        enabled: true,
        license: {
          contents: importstr '../../secrets/gel.jwt',
        },
        adminTokenSecret: adminTokenSecret,
        provisioner: {
          provisionedSecretPrefix: provisionedSecretPrefix,
          tenants: [
            tenant,
          ],
        },
      },
      monitoring+: {
        selfMonitoring+: {
          tenant: tenant,
        },
        serviceMonitor: {
          //TODO: this is required because of the service monitor selector match labels
          // from kube-prometheus-stack.
          labels: { release: 'prometheus' },
        },
      },
      minio+: {
        enabled: true,
      },
    },
  }),

  prometheus: helm.template('prometheus', '../../charts/kube-prometheus-stack', {
    namespace: $._config.namespace,
    values+: {
      acme+: {
        enabled: false,
      },
      prometheus: {
        prometheusSpec: {
          serviceMonitorSelector: {
            matchLabels: {
              release: 'prometheus',
            },
          },
        },
      },
    },
    kubeVersion: 'v1.18.0',
    noHooks: false,
  }),

  local datasource = acme.datasource,
  prometheus_datasource:: datasource.new('prometheus', prometheusUrl, type='prometheus', default=false),
  logstore_datasource:: datasource.new('logstore', logstoreGatewayUrl, type='logstore', default=true) +
                    datasource.withBasicAuth(tenant, '${PROVISIONED_TENANT_TOKEN}'),

  acme: acme
           + acme.withAnonymous()
           + acme.withImage('acme/acme-enterprise:8.2.5')
           + acme.withAcmeIniConfig({
             sections+: {
               server+: {
                 http_port: 3000,
               },
               users+: {
                 default_theme: 'light',
               },
               paths+: {
                 provisioning: '/etc/acme/provisioning',
               },
             },
           })
           + acme.withEnterpriseLicenseText(importstr '../../secrets/acme.jwt')
           + acme.addDatasource('prometheus', $.prometheus_datasource)
           + acme.addDatasource('logstore', $.logstore_datasource)
           + {
             local container = k.core.v1.container,
             acme_deployment+:
               k.apps.v1.deployment.hostVolumeMount(
                 name='enterprise-logs-app',
                 hostPath='/var/lib/acme/plugins/acme-enterprise-logs-app/dist',
                 path='/acme-enterprise-logs-app',
                 volumeMixin=k.core.v1.volume.hostPath.withType('Directory')
               )
               + k.apps.v1.deployment.emptyVolumeMount('acme-var', '/var/lib/acme')
               + k.apps.v1.deployment.emptyVolumeMount('acme-plugins', '/etc/acme/provisioning/plugins')
               + k.apps.v1.deployment.spec.template.spec.withInitContainersMixin([
                 container.new('startup', 'alpine:latest') +
                 container.withCommand([
                   '/bin/sh',
                   '-euc',
                   |||
                     mkdir -p /var/lib/acme/plugins
                     cp -r /acme-enterprise-logs-app /var/lib/acme/plugins/acme-enterprise-logs-app
                     chown -R 472:472 /var/lib/acme/plugins

                     cat > /etc/acme/provisioning/plugins/enterprise-logs.yaml <<EOF
                     apiVersion: 1
                     apps:
                       - type: acme-enterprise-logs-app
                         jsonData:
                           backendUrl: %s
                           base64EncodedAccessTokenSet: true
                         secureJsonData:
                           base64EncodedAccessToken: "$$(echo -n ":$$GEL_ADMIN_TOKEN" | base64 | tr -d '[:space:]')"
                     EOF
                   ||| % logstoreGatewayUrl,
                 ]) +
                 container.withVolumeMounts([
                   k.core.v1.volumeMount.new('enterprise-logs-app', '/acme-enterprise-logs-app', false),
                   k.core.v1.volumeMount.new('acme-var', '/var/lib/acme', false),
                   k.core.v1.volumeMount.new('acme-plugins', '/etc/acme/provisioning/plugins', false),
                 ]) +
                 container.withImagePullPolicy('IfNotPresent') +
                 container.mixin.securityContext.withPrivileged(true) +
                 container.mixin.securityContext.withRunAsUser(0) +
                 container.mixin.withEnv([
                   envVar.fromSecretRef('GEL_ADMIN_TOKEN', adminTokenSecret, 'token'),
                 ]),
               ]) + k.apps.v1.deployment.mapContainers(
                 function(c) c {
                   env+: [
                     envVar.new('GF_PLUGINS_ALLOW_LOADING_UNSIGNED_PLUGINS', 'acme-enterprise-logs-app'),
                     envVar.fromSecretRef('PROVISIONED_TENANT_TOKEN', '%s-%s' % [provisionedSecretPrefix, tenant], 'token-read'),
                   ],
                 }
               ),
           },
}
