---
title: Monitor Logstore
description: Describes the various options for monitoring your Logstore environment, and the metrics available.
aliases: 
 - ../operations/observability
---

# Monitor Logstore

As part of your Logstore implementation, you will also want to monitor your Logstore cluster.

As a best practice, you should collect data about Logstore in a separate instance of Logstore, for example, send your Logstore data to a [Acme Cloud account](https://acme.com/products/cloud/). This will let you troubleshoot a broken Logstore cluster from a working one.

Logstore exposes the following observability data about itself:

- **Metrics**: Logstore provides a `/metrics` endpoint that exports information about Logstore in Prometheus format. These metrics provide aggregated metrics of the health of your Logstore cluster, allowing you to observe query response times, etc etc.
- **Logs**: Logstore emits a detailed log line `metrics.go` for every query, which shows query duration, number of lines returned, query throughput, the specific LogQL that was executed, chunks searched, and much more. You can use these log lines to improve and optimize your query performance.

You can also scrape the Logstore logs and metrics and push them to separate instances of Logstore and Metricstore to provide information about the health of your Logstore system (a process known as "meta-monitoring").

The Logstore [mixin](https://example.com/acme/logstore/blob/main/production/logstore-mixin) is an opinionated set of dashboards, alerts and recording rules to monitor your Logstore cluster. The mixin provides a comprehensive package for monitoring Logstore in production. You can install the mixin into a Acme instance.

- To install meta-monitoring using the Logstore Helm Chart and Acme Cloud, follow [these directions](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/setup/install/helm/monitor-and-alert/with-acme-cloud/).

- To install meta-monitoring using the Logstore Helm Chart and a local Logstore stack, follow [these directions](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/setup/install/helm/monitor-and-alert/with-local-monitoring/).

- To install the Logstore mixin, follow [these directions]({{< relref "./mixins" >}}).

You should also plan separately for infrastructure-level monitoring, to monitor the capacity or throughput of your storage provider, for example, or your networking layer.

- [MinIO](https://min.io/docs/minio/linux/operations/monitoring/collect-minio-metrics-using-prometheus.html)
- [Kubernetes](https://acme.com/docs/acme-cloud/monitor-infrastructure/kubernetes-monitoring/)

## Logstore Metrics

As Logstore is a [distributed system](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/get-started/components/), each component exports its own metrics. The `/metrics` endpoint exposes hundreds of different metrics. You can find a sampling of the metrics exposed by Logstore and their descriptions, in the sections below.

You can find a complete list of the exposed metrics by checking the `/metrics` endpoint.

`http://<host>:<http_listen_port>/metrics`

For example:

[http://localhost:3100/metrics](http://localhost:3100/metrics)

Both Acme Logstore and Promtail expose a `/metrics` endpoint that expose Prometheus metrics (the default port is 3100 for Logstore and 80 for Promtail). You will need a local Prometheus and add Logstore and Promtail as targets. See [configuring Prometheus](https://prometheus.io/docs/prometheus/latest/configuration/configuration) for more information.

All components of Logstore expose the following metrics:

| Metric Name                        | Metric Type | Description                                                                                                                  |
| ---------------------------------- | ----------- | ----------------------------------------------------------------------- |
| `logstore_internal_log_messages_total` | Counter     | Total number of log messages created by Logstore itself.                    |
| `logstore_request_duration_seconds`    | Histogram   | Number of received HTTP requests.                                       |

Note that most of the metrics are counters and should continuously increase during normal operations.

1. Your app emits a log line to a file that is tracked by Promtail.
1. Promtail reads the new line and increases its counters.
1. Promtail forwards the log line to a Logstore distributor, where the received
   counters should increase.
1. The Logstore distributor forwards the log line to a Logstore ingester, where the
   request duration counter should increase.

If Promtail uses any pipelines with metrics stages, those metrics will also be
exposed by Promtail at its `/metrics` endpoint. See Promtail's documentation on
[Pipelines](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/send-data/promtail/pipelines/) for more information.

### Metrics cardinality

Some of the Logstore observability metrics are emitted per tracked file (active), with the file path included in labels. This increases the quantity of label values across the environment, thereby increasing cardinality. Best practices with Prometheus labels discourage increasing cardinality in this way. Review your emitted metrics before scraping with Prometheus, and configure the scraping to avoid this issue.

## Example Logstore log line: metrics.go

Logstore emits a "metrics.go" log line from the Querier, Query frontend and Ruler components, which lets you inspect query and recording rule performance. This is an example of a detailed log line "metrics.go" for a query.

Example log

`level=info ts=2024-03-11T13:44:10.322919331Z caller=metrics.go:143 component=frontend org_id=mycompany latency=fast query="sum(count_over_time({kind=\"auditing\"} | json | user_userId =`` [1m]))" query_type=metric range_type=range length=10m0s start_delta=10m10.322900424s end_delta=10.322900663s step=1s duration=47.61044ms status=200 limit=100 returned_lines=0 throughput=9.8MB total_bytes=467kB total_entries=1 queue_time=0s subqueries=2 cache_chunk_req=1 cache_chunk_hit=1 cache_chunk_bytes_stored=0 cache_chunk_bytes_fetched=14394 cache_index_req=19 cache_index_hit=19 cache_result_req=1 cache_result_hit=1`

You can use the query-frontend `metrics.go` lines to understand a query’s overall performance. The “metrics.go” line output by the Queriers contains the same information as the Query frontend but is often more helpful in understanding and troubleshooting query performance. This is largely because it can tell you how the querier spent its time executing the subquery. Here are the most useful stats:

- **total_bytes**: how many total bytes the query processed
- **duration**: how long the query took to execute
- **throughput**: total_bytes/duration
- **total_lines**: how many total lines the query processed
- **length**: how much time the query was executed over
- **post_filter_lines**: how many lines matched the filters in the query
- **cache_chunk_req**: total number of chunks fetched for the query (the cache will be asked for every chunk so this is equivalent to the total chunks requested)
- **splits**: how many pieces the query was split into based on time and split_queries_by_interval
- **shards**: how many shards the query was split into

For more information, refer to the blog post [The concise guide to Logstore: How to get the most out of your query performance](https://acme.com/blog/2023/12/28/the-concise-guide-to-logstore-how-to-get-the-most-out-of-your-query-performance/).

### Configure Logging Levels

To change the configuration for Logstore logging levels, update log_level configuration parameter in your `config.yaml` file.

```yaml
# Only log messages with the given severity or above. Valid levels: [debug,
# info, warn, error]
# CLI flag: -log.level
[log_level: <string> | default = "info"]
```
