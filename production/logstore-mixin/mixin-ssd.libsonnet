(import 'mixin.libsonnet') + {
  acmeDashboardFolder: 'Logstore SSD',

  _config+:: {
    internal_components: false,

    // By default the helm chart uses the Acme Agent instead of promtail
    promtail+: {
      enabled: false,
    },

    ssd+: {
      enabled: true,
    },
  },
}
