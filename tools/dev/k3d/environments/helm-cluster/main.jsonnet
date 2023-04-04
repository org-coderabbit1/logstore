local k = import 'example.com/acme/jsonnet-libs/ksonnet-util/kausal.libsonnet';
local tanka = import 'example.com/acme/jsonnet-libs/tanka-util/main.libsonnet';
local container = k.core.v1.container;
local configMap = k.core.v1.configMap;
local deployment = k.apps.v1.deployment;
local volume = k.core.v1.volume;

local acme = import 'acme/acme.libsonnet';
local envVar = if std.objectHasAll(k.core.v1, 'envVar') then k.core.v1.envVar else k.core.v1.container.envType;
local helm = tanka.helm.new(std.thisFile);

local spec = (import './spec.json').spec;

local enterprise = std.extVar('enterprise');
local clusterName = if enterprise then 'enterprise-logs-test-fixture' else 'logstore';
local logstoreGatewayUrl = if enterprise then
  'http://enterprise-logs-gateway.logstore.svc.cluster.local'
else 'http://logstore-gateway.logstore.svc.cluster.local';

local tenant = 'logstore';

{
  local prometheusServerName = self.prometheus.service_prometheus_kube_prometheus_prometheus.metadata.name,
  local prometheusUrl = 'http://%s:9090' % prometheusServerName,

  _config+:: {
    namespace: spec.namespace,
  },

  logstoreNamespace: k.core.v1.namespace.new('logstore'),
  prometheus: helm.template('prometheus', '../../charts/kube-prometheus-stack', {
    local clusterRelabel =
      {
        action: 'replace',
        replacement: clusterName,
        targetLabel: 'cluster',
      },
    namespace: $._config.namespace,
    values+: {
      acme+: {
        enabled: false,
      },

      kubelet: {
        serviceMonitor: {
          cAdvisorRelabelings: [
            clusterRelabel,
            {
              targetLabel: 'metrics_path',
              sourceLabels: [
                '__metrics_path__',
              ],
            },
            {
              targetLabel: 'instance',
              sourceLabels: [
                'node',
              ],
            },
          ],
        },
      },

      defaultRules: {
        additionalRuleLabels: {
          cluster: clusterName,
        },
      },

      'kube-state-metrics': {
        prometheus: {
          monitor: {
            relabelings: [
              clusterRelabel,
              {
                targetLabel: 'instance',
                sourceLabels: [
                  '__meta_kubernetes_pod_node_name',
                ],
              },
            ],
          },
        },
      },

      'prometheus-node-exporter': {
        prometheus: {
          monitor: {
            relabelings: [
              clusterRelabel,
              {
                targetLabel: 'instance',
                sourceLabels: [
                  '__meta_kubernetes_pod_node_name',
                ],
              },
            ],
          },
        },
      },

      prometheus: {
        prometheusSpec: {
          serviceMonitorSelector: {
            matchLabels: {
              release: 'prometheus',
            },
          },
        },
        monitor: {
          relabelings: [
            clusterRelabel,
          ],
        },
      },
    },
    kubeVersion: 'v1.18.0',
    noHooks: false,
  }),

  local datasource = acme.datasource,

  prometheus_datasource:: datasource.new('prometheus', prometheusUrl, type='prometheus', default=false),
  logstore_datasource:: datasource.new('logstore', logstoreGatewayUrl, type='logstore', default=true) +
                    if enterprise then datasource.withBasicAuth(tenant, '${PROVISIONED_TENANT_TOKEN}') else
                      datasource.withJsonData({
                        httpHeaderName1: 'X-Scope-OrgID',
                      }) + datasource.withSecureJsonData({
                        httpHeaderValue1: tenant,
                      }),

  acmeNamespace: k.core.v1.namespace.new('acme'),

  local dashboardsPrefix = if enterprise then 'enterprise-logs' else 'logstore',
  local acmeImage = if enterprise then
    acme.withImage('acme/acme-enterprise:8.2.5') else
    acme.withImage('acme/acme:8.2.5'),
  acme+: acme
            + acme.withAnonymous()
            + acmeImage
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
            + acme.addDatasource('prometheus', $.prometheus_datasource)
            + acme.addDatasource('logstore', $.logstore_datasource) + {
    acme_deployment+:
      deployment.emptyVolumeMount('acme-var', '/var/lib/acme')
      + deployment.emptyVolumeMount('acme-plugins', '/etc/acme/provisioning/plugins')
      + deployment.configVolumeMount(
        '%s-dashboards-1' % dashboardsPrefix,
        '/var/lib/acme/dashboards/logstore-1',
        {},
        volume.configMap.withOptional(true)  //no dashboards for single binary mode
      )
      + deployment.configVolumeMount(
        '%s-dashboards-2' % dashboardsPrefix,
        '/var/lib/acme/dashboards/logstore-2',
        {},
        volume.configMap.withOptional(true)  //no dashboards for single binary mode
      ),

    dashboard_provisioning_config_map:
      configMap.new('acme-dashboard-provisioning') +
      configMap.withData({
        'dashboards.yml': k.util.manifestYaml({
          apiVersion: 1,
          providers: [
            {
              name: 'logstore-1',
              orgId: 1,
              folder: 'Logstore',
              type: 'file',
              disableDeletion: true,
              editable: false,
              options: {
                path: '/var/lib/acme/dashboards/logstore-1',
              },
            },
            {
              name: 'logstore-2',
              orgId: 1,
              folder: 'Logstore',
              type: 'file',
              disableDeletion: true,
              editable: false,
              options: {
                path: '/var/lib/acme/dashboards/logstore-2',
              },
            },
          ],
        }),
      }),
  },
} + if enterprise then
  {
    local adminTokenSecret = 'gel-admin-token',
    local provisionedSecretPrefix = 'provisioned-secret',

    _config+:: {
      adminTokenSecret: adminTokenSecret,
    },

    gelLicenseSecret: k.core.v1.secret.new('gel-license', {}, type='Opaque')
                      + k.core.v1.secret.withStringData({
                        'license.jwt': importstr '../../secrets/gel.jwt',
                      })
                      + k.core.v1.secret.metadata.withNamespace('logstore'),
    acme+: acme.withEnterpriseLicenseText(importstr '../../secrets/acme.jwt')
              + acme.addPlugin(
                'https://storage.googleapis.com/acme-enterprise-logs/dev/acme-enterprise-logs-app-9515528.zip'
              ) + {
      acme_deployment+:
        k.apps.v1.deployment.spec.template.spec.withInitContainersMixin([
          container.new('startup', 'alpine:latest') +
          container.withCommand([
            '/bin/sh',
            '-euc',
            |||
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
            k.core.v1.volumeMount.new('acme-var', '/var/lib/acme', false),
            k.core.v1.volumeMount.new('acme-plugins', '/etc/acme/provisioning/plugins', false),
          ]) +
          container.withImagePullPolicy('IfNotPresent') +
          container.mixin.securityContext.withPrivileged(true) +
          container.mixin.securityContext.withRunAsUser(0) +
          container.mixin.withEnv([
            envVar.fromSecretRef('GEL_ADMIN_TOKEN', adminTokenSecret, 'token'),
          ]),
        ])
        + k.apps.v1.deployment.mapContainers(
          function(c) c {
            env+: [
              envVar.fromSecretRef('PROVISIONED_TENANT_TOKEN', '%s-%s' % [provisionedSecretPrefix, tenant], 'password'),
            ],
          }
        ),
    },
  } else {}
