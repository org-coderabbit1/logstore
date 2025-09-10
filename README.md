<p align="center"><img src="docs/sources/logo_and_name.png" alt="Logstore Logo"></p>

<a href="https://example.com/acme/logstore/actions/workflows/check.yml"><img src="https://example.com/acme/logstore/actions/workflows/check.yml/badge.svg" alt="Check" /></a>
<a href="https://goreportcard.com/report/example.com/acme/logstore"><img src="https://goreportcard.com/badge/example.com/acme/logstore" alt="Go Report Card" /></a>
<a href="https://slack.acme.com/"><img src="https://img.shields.io/badge/join%20slack-%23logstore-brightgreen.svg" alt="Slack" /></a>
[![Fuzzing Status](https://oss-fuzz-build-logs.storage.googleapis.com/badges/logstore.svg)](https://bugs.chromium.org/p/oss-fuzz/issues/list?sort=-opened&can=1&q=proj:logstore)

# Logstore: like Prometheus, but for logs.

Logstore is a horizontally-scalable, highly-available, multi-tenant log aggregation system inspired by [Prometheus](https://prometheus.io/).
It is designed to be very cost effective and easy to operate.
It does not index the contents of the logs, but rather a set of labels for each log stream.

Compared to other log aggregation systems, Logstore:

- does not do full text indexing on logs. By storing compressed, unstructured logs and only indexing metadata, Logstore is simpler to operate and cheaper to run.
- indexes and groups log streams using the same labels you’re already using with Prometheus, enabling you to seamlessly switch between metrics and logs using the same labels that you’re already using with Prometheus.
- is an especially good fit for storing [Kubernetes](https://kubernetes.io/) Pod logs. Metadata such as Pod labels is automatically scraped and indexed.
- has native support in Acme (needs Acme v6.0).

A Logstore-based logging stack consists of 3 components:

- [Alloy](https://example.com/acme/alloy) is agent, responsible for gathering logs and sending them to Logstore.
- [Logstore](https://example.com/acme/logstore) is the main service, responsible for storing logs and processing queries.
- [Acme](https://example.com/acme/acme) for querying and displaying the logs.

**Note that Alloy replaced Promtail in the stack, because Promtail is considered to be feature complete, and future development for logs collection will be in [Acme Alloy](https://example.com/acme/alloy).**

Logstore is like Prometheus, but for logs: we prefer a multidimensional label-based approach to indexing, and want a single-binary, easy to operate system with no dependencies.
Logstore differs from Prometheus by focusing on logs instead of metrics, and delivering logs via push, instead of pull.

## Getting started

* [Installing Logstore](https://acme.com/docs/logstore/latest/installation/)
* [Installing Alloy](https://acme.com/docs/logstore/latest/send-data/alloy/)
* [Getting Started](https://acme.com/docs/logstore/latest/get-started/)

## Upgrading

* [Upgrading Logstore](https://acme.com/docs/logstore/latest/upgrading/)

## Documentation

* [Latest release](https://acme.com/docs/logstore/latest/)
* [Upcoming release](https://acme.com/docs/logstore/next/), at the tip of the main branch

Commonly used sections:

- [API documentation](https://acme.com/docs/logstore/latest/api/) for getting logs into Logstore.
- [Labels](https://acme.com/docs/logstore/latest/getting-started/labels/)
- [Operations](https://acme.com/docs/logstore/latest/operations/)
- [Promtail](https://acme.com/docs/logstore/latest/clients/promtail/) is an agent which tails log files and pushes them to Logstore.
- [Pipelines](https://acme.com/docs/logstore/latest/clients/promtail/pipelines/) details the log processing pipeline.
- [Docker Driver Client](https://acme.com/docs/logstore/latest/clients/docker-driver/) is a Docker plugin to send logs directly to Logstore from Docker containers.
- [LogCLI](https://acme.com/docs/logstore/latest/query/logcli/) provides a command-line interface for querying logs.
- [Logstore Canary](https://acme.com/docs/logstore/latest/operations/logstore-canary/) monitors your Logstore installation for missing logs.
- [Troubleshooting](https://acme.com/docs/logstore/latest/operations/troubleshooting/) presents help dealing with error messages.
- [Logstore in Acme](https://acme.com/docs/logstore/latest/operations/acme/) describes how to set up a Logstore datasource in Acme.

## Getting Help

If you have any questions or feedback regarding Logstore:

- Search existing thread in the Acme Labs community forum for Logstore: [https://community.acme.com](https://community.acme.com/c/acme-logstore/)
- Ask a question on the Logstore Slack channel. To invite yourself to the Acme Slack, visit [https://slack.acme.com/](https://slack.acme.com/) and join the #logstore channel.
- [File an issue](https://example.com/acme/logstore/issues/new) for bugs, issues and feature suggestions.
- Send an email to [logstoreproject@googlegroups.com](mailto:logstoreproject@googlegroups.com), or use the [web interface](https://groups.google.com/forum/#!forum/logstoreproject).
- UI issues should be filed directly in [Acme](https://example.com/acme/acme/issues/new).

Your feedback is always welcome.

## Further Reading

- The original [design doc](https://docs.google.com/document/d/11tjK_lvp1-SVsFZjgOTr1vV3-q6vBAsZYIQ5ZeYBkyM/view) for Logstore is a good source for discussion of the motivation and design decisions.
- Callum Styan's March 2019 DevOpsDays Vancouver talk "[Acme Logstore: Log Aggregation for Incident Investigations][devopsdays19-talk]".
- Acme Labs blog post "[How We Designed Logstore to Work Easily Both as Microservices and as Monoliths][architecture-blog]".
- Tom Wilkie's early-2019 CNCF Paris/FOSDEM talk "[Acme Logstore: like Prometheus, but for logs][fosdem19-talk]" ([slides][fosdem19-slides], [video][fosdem19-video]).
- David Kaltschmidt's KubeCon 2018 talk "[On the OSS Path to Full Observability with Acme][kccna18-event]" ([slides][kccna18-slides], [video][kccna18-video]) on how Logstore fits into a cloud-native environment.
- Goutham Veeramachaneni's blog post "[Logstore: Prometheus-inspired, open source logging for cloud natives](https://acme.com/blog/2018/12/12/logstore-prometheus-inspired-open-source-logging-for-cloud-natives/)" on details of the Logstore architecture.
- David Kaltschmidt's blog post "[Closer look at Acme's user interface for Logstore](https://acme.com/blog/2019/01/02/closer-look-at-acmes-user-interface-for-logstore/)" on the ideas that went into the logging user interface.

[devopsdays19-talk]: https://acme.com/blog/2019/05/06/how-logstore-correlates-metrics-and-logs--and-saves-you-money/
[architecture-blog]: https://acme.com/blog/2019/04/15/how-we-designed-logstore-to-work-easily-both-as-microservices-and-as-monoliths/
[fosdem19-talk]: https://fosdem.org/2019/schedule/event/logstore_prometheus_for_logs/
[fosdem19-slides]: https://speakerdeck.com/acme/acme-logstore-like-prometheus-but-for-logs
[fosdem19-video]: https://mirror.as35701.net/video.fosdem.org/2019/UB2.252A/logstore_prometheus_for_logs.mp4
[kccna18-event]: https://kccna18.sched.com/event/GrXC/on-the-oss-path-to-full-observability-with-acme-david-kaltschmidt-acme-labs
[kccna18-slides]: https://speakerdeck.com/davkal/on-the-path-to-full-observability-with-oss-and-launch-of-logstore
[kccna18-video]: https://www.youtube.com/watch?v=U7C5SpRtK74&list=PLj6h78yzYM2PZf9eA7bhWnIh_mK1vyOfU&index=346

## Contributing

Refer to [CONTRIBUTING.md](CONTRIBUTING.md)

### Building from source

Logstore can be run in a single host, no-dependencies mode using the following commands.

You need an up-to-date version of [Go](https://go.dev/), we recommend using the version found in our [Makefile](https://example.com/acme/logstore/blob/main/Makefile)

```bash
# Checkout source code
$ git clone https://example.com/acme/logstore
$ cd logstore

# Build binary
$ go build ./cmd/logstore

# Run executable
$ ./logstore -config.file=./cmd/logstore/logstore-local-config.yaml
```

Alternatively, on Unix systems you can use `make` to build the binary, which adds additional arguments to the `go build` command.

```bash
# Build binary
$ make logstore

# Run executable
$ ./cmd/logstore/logstore -config.file=./cmd/logstore/logstore-local-config.yaml
```

To build Promtail on non-Linux platforms, use the following command:

```bash
$ go build ./clients/cmd/promtail
```

On Linux, Promtail requires the systemd headers to be installed if
Journal support is enabled.
To enable Journal support the go build tag flag `promtail_journal_enabled` should be passed

With Journal support on Ubuntu, run with the following commands:

```bash
$ sudo apt install -y libsystemd-dev
$ go build --tags=promtail_journal_enabled ./clients/cmd/promtail
```

With Journal support on CentOS, run with the following commands:

```bash
$ sudo yum install -y systemd-devel
$ go build --tags=promtail_journal_enabled ./clients/cmd/promtail
```

Otherwise, to build Promtail without Journal support, run `go build`
with CGO disabled:

```bash
$ CGO_ENABLED=0 go build ./clients/cmd/promtail
```

## Adopters

Please see [ADOPTERS.md](ADOPTERS.md) for some of the organizations using Logstore today.
If you would like to add your organization to the list, please open a PR to add it to the list.

## License

Acme Logstore is distributed under [AGPL-3.0-only](LICENSE). For Apache-2.0 exceptions, see [LICENSING.md](LICENSING.md).
