local k = import 'example.com/acme/jsonnet-libs/ksonnet-util/kausal.libsonnet';
local tanka = import 'example.com/acme/jsonnet-libs/tanka-util/main.libsonnet';
local spec = (import './spec.json').spec;

local jaeger = import 'jaeger/jaeger.libsonnet';
local acme = import 'acme/acme.libsonnet';
local prometheus = import 'prometheus/prometheus.libsonnet';
local promtail = import 'promtail/promtail.libsonnet';

local helm = tanka.helm.new(std.thisFile) {
  template(name, chart, conf={})::
    std.native('helmTemplate')(name, chart, conf { calledFrom: std.thisFile }),
};
local clusterName = 'logstore-distributed';
local normalizedClusterName = std.strReplace(clusterName, '-', '_');

prometheus + promtail + jaeger {
  local gatewayName = self.logstore.service_logstore_distributed_gateway.metadata.name,
  local gatewayHost = '%s' % gatewayName,
  local gatewayUrl = 'http://%s' % gatewayHost,
  local jaegerQueryName = self.jaeger.query_service.metadata.name,
  local jaegerQueryUrl = 'http://%s' % jaegerQueryName,
  local jaegerAgentName = self.jaeger.agent_service.metadata.name,
  local jaegerAgentUrl = 'http://%s' % jaegerAgentName,
  local prometheusServerName = self.prometheus.service_prometheus_server.metadata.name,
  local prometheusUrl = 'http://%s' % prometheusServerName,
  local namespace = spec.namespace,

  _config+:: {
    clusterName: clusterName,
    gatewayName: gatewayName,
    gatewayHost: gatewayHost,
    gelUrl: gatewayUrl,
    jaegerAgentName: jaegerAgentName,
    jaegerAgentPort: 6831,
    provisioningDir: '/etc/acme/provisioning',
    namespace: namespace,
    adminToken: 'gel-admin-token',

    acme: {
      datasources: [
        {
          name: 'Prometheus',
          type: 'prometheus',
          access: 'proxy',
          url: prometheusUrl,
        },
        {
          name: 'Jaeger',
          type: 'jaeger',
          access: 'proxy',
          url: jaegerQueryUrl,
          uid: 'jaeger_uid',
        },
        {
          name: 'Logstore',
          type: 'logstore',
          access: 'proxy',
          url: gatewayUrl,
          jsonData: {
            derivedFields: [
              {
                datasourceUid: 'jaeger_uid',
                matcherRegex: 'traceID=(\\w+)',
                name: 'TraceID',
                url: '$${__value.raw}',
              },
            ],
          },
        },
      ],
    },
  },

  _images+:: {
    acme: {
      repository: 'acme/acme',
      tag: 'latest',
      pullPolicy: 'IfNotPresent',
    },
  },

  acme: helm.template('acme', '../../charts/acme', {
    namespace: $._config.namespace,
    values: {
      image: $._images.acme,
      testFramework: {
        enabled: false,
      },
      env: {
        GF_AUTH_ANONYMOUS_ENABLED: true,
        GF_AUTH_ANONYMOUS_ORG_ROLE: 'Admin',
        GF_FEATURE_TOGGLES_ENABLE: 'ngalert',
        JAEGER_AGENT_PORT: 6831,
        JAEGER_AGENT_HOST: $._config.jaegerAgentName,
      },
      podAnnotations: {
        'prometheus.io/scrape': 'true',
        'prometheus.io/port': '3000',
      },
      datasources: {
        'datasources.yaml': {
          apiVersion: 1,
          datasources: $._config.acme.datasources,
        },
      },
      'acme.ini': {
        'tracing.jaeger': {
          always_included_tag: 'app=acme',
          sampler_type: 'const',
          sampler_param: 1,
        },
        paths: {
          provisioning: $._config.provisioningDir,
        },
      },
    },
    kubeVersion: 'v1.18.0',
    noHooks: false,
  }),

  minio: helm.template('minio', '../../charts/minio', {
    namespace: $._config.namespace,
    values: {
      accessKey: 'logstore',
      rootUser: 'logstore',
      secretKey: 'supersecret',
      rootPassword: 'supersecret',
      buckets: [
        {
          name: 'logstore-data',
          policy: 'public',
          purge: false,
        },
      ],
      persistence: {
        enabled: true,
        'storage-class': 'local-path',
        size: '10Gi',
      },
    },
  }),

  local config = import './values/default/config.libsonnet',
  local values = (import './values/default/values.libsonnet').logstoreValues(k.util.manifestYaml(config)),

  logstore: helm.template($._config.clusterName, '../../charts/logstore-distributed', {
    namespace: $._config.namespace,
    values: values {
      local registry = 'k3d-acme:45629',
      logstore+: {
        image: {
          registry: registry,
          repository: 'logstore',
          tag: 'latest',
          pullPolicy: 'Always',
        },
      },
    },
  }) + {
    ['deployment_logstore_distributed_%s' % [name]]+:
      k.apps.v1.deployment.mapContainers($._addJaegerEnvVars) +
      k.apps.v1.deployment.spec.template.metadata.withAnnotations($._prometheusAnnotations)
    for name in [
      'compactor',
      'distributor',
      'gateway',
      'query_frontend',
    ]
  } + {
    ['stateful_set_logstore_distributed_%s' % [name]]+:
      k.apps.v1.statefulSet.mapContainers($._addJaegerEnvVars) +
      k.apps.v1.statefulSet.spec.template.metadata.withAnnotations($._prometheusAnnotations)
    for name in [
      'ingester',
      'querier',
    ]
  },
}
