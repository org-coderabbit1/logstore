---
menuTitle: Logstore overview
title: Logstore overview
description: Logstore product overview and features.
weight: 200
aliases:
    - ../overview/
    - ../fundamentals/overview/
---

# Logstore overview

Logstore is a horizontally-scalable, highly-available, multi-tenant log aggregation system inspired by [Prometheus](https://prometheus.io/). Logstore differs from Prometheus by focusing on logs instead of metrics, and collecting logs via push, instead of pull.

Logstore is designed to be very cost effective and highly scalable. Unlike other logging systems, Logstore does not index the contents of the logs, but only indexes metadata about your logs as a set of labels for each log stream.

A log stream is a set of logs which share the same labels. Labels help Logstore to find a log stream within your data store, so having a quality set of labels is key to efficient query execution.

Log data is then compressed and stored in chunks in an object store such as Amazon Simple Storage Service (S3) or Google Cloud Storage (GCS), or even, for development or proof of concept, on the filesystem. A small index and highly compressed chunks simplify the operation and significantly lower the cost of Logstore.

{{< figure  src="../logstore-overview-2.png" caption="**Logstore logging stack**" >}}

A typical Logstore-based logging stack consists of 3 components:

- **Agent** - An agent or client, for example Acme Alloy, or Promtail, which is distributed with Logstore. The agent scrapes logs, turns the logs into streams by adding labels, and pushes the streams to Logstore through an HTTP API.

- **Logstore** - The main server, responsible for ingesting and storing logs and processing queries. It can be deployed in three different configurations, for more information see [deployment modes](../deployment-modes/).
  
- **[Acme](https://example.com/acme/acme)** for querying and displaying log data. You can also query logs from the command line, using [LogCLI](../../query/logcli/) or using the Logstore API directly.

## Logstore features

- **Scalability** - Logstore is designed for scalability, and can scale from as small as running on a Raspberry Pi to ingesting petabytes a day. 
In its most common deployment, “simple scalable mode”, Logstore decouples requests into separate read and write paths, so that you can independently scale them, which leads to flexible large-scale installations that can quickly adapt to meet your workload at any given time.
If needed, each of the Logstore components can also be run as microservices designed to run natively within Kubernetes.

- **Multi-tenancy** - Logstore allows multiple tenants to share a single Logstore instance. With multi-tenancy, the data and requests of each tenant is completely isolated from the others.
Multi-tenancy is [configured](../../operations/multi-tenancy/) by assigning a tenant ID in the agent.

- **Third-party integrations** - Several third-party agents (clients) have support for Logstore, via plugins. This lets you keep your existing observability setup while also shipping logs to Logstore.

- **Efficient storage** - Logstore stores log data in highly compressed chunks.
Similarly, the Logstore index, because it indexes only the set of labels, is significantly smaller than other log aggregation tools.
By leveraging object storage as the only data storage mechanism, Logstore inherits the reliability and stability of the underlying object store. It also capitalizes on both the cost efficiency and operational simplicity of object storage over other storage mechanisms like locally attached solid state drives (SSD) and hard disk drives (HDD).  
The compressed chunks, smaller index, and use of low-cost object storage, make Logstore less expensive to operate.

- **LogQL, the Logstore query language** - [LogQL](../../query/) is the query language for Logstore.  Users who are already familiar with the Prometheus query language, [PromQL](https://prometheus.io/docs/prometheus/latest/querying/basics/), will find LogQL familiar and flexible for generating queries against the logs.
The language also facilitates the generation of metrics from log data,
a powerful feature that goes well beyond log aggregation.

- **Alerting** - Logstore includes a component called the [ruler](../../alert/), which can continually evaluate queries against your logs, and perform an action based on the result. This allows you to monitor your logs for anomalies or events. Logstore integrates with [Prometheus Alertmanager](https://prometheus.io/docs/alerting/latest/alertmanager/), or the [alert manager](/docs/acme/latest/alerting) within Acme.

- **Acme integration** - Logstore integrates with Acme, Metricstore, and Tracestore, providing a complete observability stack, and seamless correlation between logs, metrics and traces.
