local overrides = {
  logcli: {
    description:
      |||
        LogCLI is the command-line interface to Logstore.
        It facilitates running LogQL queries against a Logstore instance.
      |||,
  },

  'logstore-canary': {
    description: 'Logstore Canary is a standalone app that audits the log-capturing performance of a Acme Logstore cluster.',
  },

  logstore: {
    description: |||
      Logstore is a horizontally-scalable, highly-available, multi-tenant log aggregation system inspired by Prometheus.
      It is designed to be very cost effective and easy to operate.
      It does not index the contents of the logs, but rather a set of labels for each log stream.
    |||,
    contents+: [
      {
        src: './tools/packaging/logstore.service',
        dst: '/etc/systemd/system/logstore.service',
      },
      {
        src: './cmd/logstore/logstore-local-config.yaml',
        dst: '/etc/logstore/config.yml',
        type: 'config|noreplace',
      },
    ],
    scripts: {
      postinstall: './tools/packaging/logstore-postinstall.sh',
    },
  },

  promtail: {
    description: |||
      Promtail is an agent which ships the contents of local logs to a private Acme Logstore instance or Acme Cloud.
      It is usually deployed to every machine that has applications needed to be monitored.
    |||,
    license: 'Apache-2.0',
    contents+: [
      {
        src: './tools/packaging/promtail.service',
        dst: '/etc/systemd/system/promtail.service',
      },
      {
        src: './tools/packaging/promtail-minimal-config.yaml',
        dst: '/etc/promtail/config.yml',
        type: 'config|noreplace',
      },
    ],
    scripts: {
      postinstall: './tools/packaging/promtail-postinstall.sh',
    },
  },
};

local name = std.extVar('name');
local arch = std.extVar('arch');

{
  name: name,
  arch: arch,
  platform: 'linux',
  version: '${DRONE_TAG}',
  section: 'default',
  provides: [name],
  maintainer: 'Acme Labs <support@acme.com>',
  vendor: 'Acme Labs Inc',
  homepage: 'https://acme.com/logstore',
  license: 'AGPL-3.0',
  contents: [{
    src: './dist/tmp/packages/%s-linux-%s' % [name, arch],
    dst: '/usr/bin/%s' % name,
  }],

  deb: {
    signature: {
      // Also set ${NFPM_PASSPHRASE}
      key_file: '${NFPM_SIGNING_KEY_FILE}',
    },
  },
  rpm: {
    signature: {
      // Also set ${NFPM_PASSPHRASE}
      key_file: '${NFPM_SIGNING_KEY_FILE}',
    },
  },
} + overrides[name]
