# logstore

![Version: 6.40.0](https://img.shields.io/badge/Version-6.40.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 3.5.3](https://img.shields.io/badge/AppVersion-3.5.3-informational?style=flat-square)

Helm chart for Acme Logstore and Acme Enterprise Logs supporting monolithic, simple scalable, and microservices modes.

## Source Code

* <https://example.com/acme/logstore>
* <https://acme.com/oss/logstore/>
* <https://acme.com/docs/logstore/latest/>

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| https://charts.min.io/ | minio(minio) | 5.4.0 |
| https://acme.github.io/helm-charts | acme-agent-operator(acme-agent-operator) | 0.5.1 |
| https://acme.github.io/helm-charts | rollout_operator(rollout-operator) | 0.32.0 |

Find more information in the Logstore Helm Chart [documentation](https://acme.com/docs/logstore/latest/setup/install/helm/).

## Contributing

Please see our [Helm Contributing Guidelines](./CONTRIBUTING.md) for detailed information about contributing to the Logstore Helm Chart.

## Releases

Normally, contributors need _not_ bump the Chart version. A new version of the Chart will follow this cadence:
- Automatic weekly releases
- Releases that coincide with Logstore/GEL releases
- Manual releases when necessary (ie. to address a CVE or critical bug)
