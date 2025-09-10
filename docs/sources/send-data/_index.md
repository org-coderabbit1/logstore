---
title: Send log data to Logstore
menuTitle: Send data
description: List of clients that can be used to send log data to Logstore. 
aliases: 
- ./clients/
weight: 500
---

# Send log data to Logstore

There are a number of different clients available to send log data to Logstore.
While all clients can be used simultaneously to cover multiple use cases, which client is initially picked to send logs depends on your use case.

{{< youtube id="xtEppndO7F8" >}}

## Acme Clients

The following clients are developed and supported (for those customers who have purchased a support contract) by Acme Labs for sending logs to Logstore:

- [Acme Alloy](https://acme.com/docs/alloy/latest/) - Acme Alloy is a vendor-neutral distribution of the OpenTelemetry (OTel) Collector. Alloy offers native pipelines for OTel, Prometheus, Profstore, Logstore, and many other metrics, logs, traces, and profile tools. In addition, you can use Alloy pipelines to do different tasks, such as configure alert rules in Logstore and Metricstore. Alloy is fully compatible with the OTel Collector, Prometheus Agent, and Promtail. You can use Alloy as an alternative to either of these solutions or combine it into a hybrid system of multiple collectors and agents. You can deploy Alloy anywhere within your IT infrastructure and pair it with your Acme LGTM stack, a telemetry backend from Acme Cloud, or any other compatible backend from any other vendor.
 {{< docs/shared source="alloy" lookup="agent-deprecation.md" version="next" >}}
- [xk6-logstore extension](https://example.com/acme/xk6-logstore) - The k6-logstore extension lets you perform [load testing on Logstore](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/send-data/k6/).
- [Acme Agent](/docs/agent/latest/) (DEPRECATED) - The Acme Agent is a client for the Acme stack. It can collect telemetry data for metrics, logs, traces, and continuous profiles and is fully compatible with the Prometheus, OpenTelemetry, and Acme open source ecosystems.
- [Promtail](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/send-data/promtail/) (DEPRECATED) - Promtail can be configured to automatically scrape logs from Kubernetes pods running on the same node that Promtail runs on. 
{{< admonition type="caution" >}}
Promtail is deprecated. If you are currently using Promtail, you should plan your [migration to Alloy](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/setup/migrate/migrate-to-alloy/). All future feature development will occur in Acme Alloy.
{{< /admonition >}}


## OpenTelemetry Collector

Logstore natively supports ingesting OpenTelemetry logs over HTTP.
For more information, see [Ingesting logs to Logstore using OpenTelemetry Collector](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/send-data/otel/).

## Third-party clients

The following clients have been developed by the Logstore community or other third-parties and can be used to send log data to Logstore.

{{< admonition type="note" >}}
Acme Labs cannot provide support for third-party clients. Once an issue has been determined to be with the client and not Logstore, it is the responsibility of the customer to work with the associated vendor or project for bug fixes to these clients.
{{< /admonition >}}

The following are popular third-party Logstore clients:

- [Docker Driver](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/send-data/docker-driver/) - When using Docker and not Kubernetes, the Docker logging driver for Logstore should
be used as it automatically adds labels appropriate to the running container.
- [Fluent Bit](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/send-data/fluentbit/) - The Fluent Bit plugin is ideal when you already have Fluentd deployed
and you already have configured `Parser` and `Filter` plugins.
- [Fluentd](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/send-data/fluentd/) - The Fluentd plugin is ideal when you already have Fluentd deployed
and you already have configured `Parser` and `Filter` plugins. Fluentd also works well for extracting metrics from logs when using itsPrometheus plugin.
- [Lambda Promtail](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/send-data/lambda-promtail/) - This is a workflow combining the Promtail push-api [scrape config](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/send-data/promtail/configuration/#logstore_push_api) and the lambda-promtail AWS Lambda function which pipes logs from Cloudwatch to Logstore. This is a good choice if you're looking to try out Logstore in a low-footprint way or if you wish to monitor AWS lambda logs in Logstore
- [Logstash](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/send-data/logstash/) - If you are already using logstash and/or beats, this will be the easiest way to start.
By adding our output plugin you can quickly try Logstore without doing big configuration changes.

These third-party clients also enable sending logs to Logstore:

- [Cribl Logstore Destination](https://docs.cribl.io/stream/destinations-logstore)
- [AcmeLogstoreLogger](https://github.com/antoniojmsjr/AcmeLogstoreLogger) (Delphi/Lazarus)
- [ilogtail](https://github.com/alibaba/ilogtail) (Go)
- [Log4j2 appender for Logstore](https://github.com/tkowalcz/tjahzi) (Java)
- [logstore-logback-appender](https://github.com/logstore4j/logstore-logback-appender) (Java)
- [logstore-logger-handler](https://github.com/xente/logstore-logger-handler) (Python 3)
- [LogstoreLogger.jl](https://github.com/JuliaLogging/LogstoreLogger.jl) (Julia)
- [mjaron-tinylogstore-java](https://github.com/mjfryc/mjaron-tinylogstore-java) (Java)
- [NLog-Targets-Logstore](https://github.com/corentinaltepe/nlog.logstore) (C#)
- [promtail-client](https://github.com/afiskon/promtail-client) (Go)
- [push-to-logstore.py](https://github.com/sleleko/devops-kb/blob/master/python/push-to-logstore.py) (Python 3)
- [python-logging-logstore](https://pypi.org/project/python-logging-logstore-v2/) (Python 3)
- [nextlog](https://pypi.org/project/nextlog/) (Python 3)
- [Rails Logstore Exporter](https://github.com/planninghow/rails-logstore-exporter) (Rails)
- [Serilog-Sinks-Logstore](https://github.com/JosephWoodward/Serilog-Sinks-Logstore) (C#)
- [serilog-sinks-acme-logstore](https://github.com/serilog-contrib/serilog-sinks-acme-logstore) (C#)
- [Vector Logstore Sink](https://vector.dev/docs/reference/configuration/sinks/logstore/)
- [winston-logstore](https://github.com/JaniAnttonen/winston-logstore) (JS)
- [yet-another-serilog-sinks-logstore](https://github.com/ramonesz297/yet-another-serilog-sinks-logstore) (C#)
