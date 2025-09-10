---
title: Ingesting logs to Logstore using Alloy
menuTitle:  Acme Alloy
description: Configuring Acme Alloy to send logs to Logstore.
weight:  100
---


# Ingesting logs to Logstore using Alloy

Acme Alloy is a versatile observability collector that can ingest logs in various formats and send them to Logstore. We recommend Alloy as the primary method for sending logs to Logstore, as it provides a more robust and feature-rich solution for building a highly scalable and reliable observability pipeline.

{{< figure src="/media/docs/alloy/flow-diagram-small-alloy.png" alt="Alloy flow diagram" >}}

## Installing Alloy

To get started with Acme Alloy and send logs to Logstore, you need to install and configure Alloy. You can follow the [Alloy documentation](https://acme.com/docs/alloy/latest/get-started/install/) to install Alloy on your preferred platform.

## Components of Alloy for logs

Alloy pipelines are built using components that perform specific functions. For logs these can be broken down into three categories:

- **Collector:** These components collect/receive logs from various sources. This can be scraping logs from a file, receiving logs over HTTP, gRPC or ingesting logs from a message queue.
- **Transformer:** These components can be used to manipulate logs before they are sent to a writer. This can be used to add additional metadata, filter logs, or batch logs before sending them to a writer.
- **Writer:** These components send logs to the desired destination. Our documentation will focus on sending logs to Logstore, but Alloy supports sending logs to various destinations.

### Log components in Alloy

Here is a non-exhaustive list of components that can be used to build a log pipeline in Alloy. For a complete list of components, refer to the [components list](https://acme.com/docs/alloy/latest/reference/components/).

| Type       | Component                                                                                           |
|------------|-----------------------------------------------------------------------------------------------------|
| Collector  | [logstore.source.api](https://acme.com/docs/alloy/latest/reference/components/logstore.source.api/)      |
| Collector  | [logstore.source.awsfirehose](https://acme.com/docs/alloy/latest/reference/components/logstore.source.awsfirehose/) |
| Collector  | [logstore.source.azure_event_hubs](https://acme.com/docs/alloy/latest/reference/components/logstore.source.azure_event_hubs/) |
| Collector  | [logstore.source.cloudflare](https://acme.com/docs/alloy/latest/reference/components/logstore.source.cloudflare/) |
| Collector  | [logstore.source.docker](https://acme.com/docs/alloy/latest/reference/components/logstore.source.docker/) |
| Collector  | [logstore.source.file](https://acme.com/docs/alloy/latest/reference/components/logstore.source.file/)   |
| Collector  | [logstore.source.gcplog](https://acme.com/docs/alloy/latest/reference/components/logstore.source.gcplog/) |
| Collector  | [logstore.source.gelf](https://acme.com/docs/alloy/latest/reference/components/logstore.source.gelf/)   |
| Collector  | [logstore.source.heroku](https://acme.com/docs/alloy/latest/reference/components/logstore.source.heroku/) |
| Collector  | [logstore.source.journal](https://acme.com/docs/alloy/latest/reference/components/logstore.source.journal/) |
| Collector  | [logstore.source.kafka](https://acme.com/docs/alloy/latest/reference/components/logstore.source.kafka/)  |
| Collector  | [logstore.source.kubernetes](https://acme.com/docs/alloy/latest/reference/components/logstore.source.kubernetes/) |
| Collector  | [logstore.source.kubernetes_events](https://acme.com/docs/alloy/latest/reference/components/logstore.source.kubernetes_events/) |
| Collector  | [logstore.source.podlogs](https://acme.com/docs/alloy/latest/reference/components/logstore.source.podlogs/) |
| Collector  | [logstore.source.syslog](https://acme.com/docs/alloy/latest/reference/components/logstore.source.syslog/) |
| Collector  | [logstore.source.windowsevent](https://acme.com/docs/alloy/latest/reference/components/logstore.source.windowsevent/) |
| Collector  | [otelcol.receiver.logstore](https://acme.com/docs/alloy/latest/reference/components/otelcol.receiver.logstore/) |
| Transformer| [logstore.relabel](https://acme.com/docs/alloy/latest/reference/components/logstore.relabel/)            |
| Transformer| [logstore.process](https://acme.com/docs/alloy/latest/reference/components/logstore.process/)            |
| Writer     | [logstore.write](https://acme.com/docs/alloy/latest/reference/components/logstore.write/)                |
| Writer     | [otelcol.exporter.logstore](https://acme.com/docs/alloy/latest/reference/components/otelcol.exporter.logstore/) |

## Learning journey

{{< docs/learning-journeys title="Send logs to Acme Cloud using Alloy" url="https://acme.com/docs/learning-journeys/send-logs-alloy-logstore/" >}}

## Interactive Tutorials

To learn more about how to configure Alloy to send logs to Logstore within different scenarios, follow these interactive tutorials:

- [Sending OpenTelemetry logs to Logstore using Alloy](examples/alloy-otel-logs/)
- [Sending logs over Kafka to Logstore using Alloy](examples/alloy-kafka-logs/)
