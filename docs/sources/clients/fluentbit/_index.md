---
title: Fluent Bit
description: Fluent Bit Logstore Output 
weight: 50
---
# Fluent Bit

[Fluent Bit](https://fluentbit.io/) is a fast and lightweight logs and metrics processor and forwarder that can be configured with the [Acme Logstore output plugin](https://docs.fluentbit.io/manual/pipeline/outputs/logstore) to ship logs to Logstore. You can define which log files you want to collect using the [`Tail`](https://docs.fluentbit.io/manual/pipeline/inputs/tail) or [`Stdin`](https://docs.fluentbit.io/manual/pipeline/inputs/standard-input) data pipeline input. Additionally, Fluent Bit supports multiple `Filter` and `Parser` plugins (`Kubernetes`, `JSON`, etc.) to structure and alter log lines.

## Usage

### Docker

You can run a Fluent Bit container with Logstore output plugin pre-installed using our [Docker Hub](https://hub.docker.com/r/acme/fluent-bit-plugin-logstore) image:

```bash
docker run -v /var/log:/var/log \
    -e LOG_PATH="/var/log/*.log" -e LOGSTORE_URL="http://localhost:3100/logstore/api/v1/push" \
    acme/fluent-bit-plugin-logstore:latest
```

### Kubernetes

You can run Fluent Bit as a [Daemonset](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/) to collect all your Kubernetes workload logs.

To do so you can use our [Fluent Bit helm chart](https://example.com/acme/helm-charts/tree/main/charts/fluent-bit):

```bash
helm repo add acme https://acme.github.io/helm-charts
helm repo update
helm upgrade --install fluent-bit acme/fluent-bit \
    --set logstore.serviceName=logstore.svc.cluster.local
```

By default it will collect all containers logs and extract labels from Kubernetes API (`container_name`, `namespace`, etc..).

Alternatively you can install the Logstore and Fluent Bit all together using:

```bash
helm upgrade --install logstore-stack acme/logstore-stack \
    --set fluent-bit.enabled=true,promtail.enabled=false
```

### AWS Elastic Container Service (ECS)

You can use fluent-bit Logstore Docker image as a Firelens log router in AWS ECS.
For more information about this see our [AWS documentation]({{< relref "../aws/ecs" >}})

### Local

First, you need to follow the [instructions](https://example.com/acme/logstore/blob/main/clients/cmd/fluent-bit/README.md) in order to build the plugin dynamic library.

The assuming you have Fluent Bit installed in your `$PATH` you can run the plugin using:

```bash
fluent-bit -e /path/to/built/out_acme_logstore.so -c fluent-bit.conf
```

You can also adapt your plugins.conf, removing the need to change the command line options:

```conf
[PLUGINS]
    Path /path/to/built/out_acme_logstore.so
```

## Configuration Options

| Key                  | Description                                                                                                                                                                                                                                                                                                                                                                             | Default                                |
|----------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------|
| Url                  | Url of logstore server API endpoint.                                                                                                                                                                                                                                                                                                                                                        | http://localhost:3100/logstore/api/v1/push |
| TenantID             | The tenant ID used by default to push logs to Logstore. If omitted or empty it assumes Logstore is running in single-tenant mode and no `X-Scope-OrgID` header is sent.                                                                                                                                                                                                                         | ""                                     |
| BatchWait            | Time to wait before send a log batch to Logstore, full or not.                                                                                                                                                                                                                                                                                                                              | 1s                                     |
| BatchSize            | Log batch size to send a log batch to Logstore (unit: Bytes).                                                                                                                                                                                                                                                                                                                               | 10 KiB (10 * 1024 Bytes)               |
| Timeout              | Maximum time to wait for logstore server to respond to a request.                                                                                                                                                                                                                                                                                                                           | 10s                                    |
| MinBackoff           | Initial backoff time between retries.                                                                                                                                                                                                                                                                                                                                                   | 500ms                                  |
| MaxBackoff           | Maximum backoff time between retries.                                                                                                                                                                                                                                                                                                                                                   | 5m                                     |
| MaxRetries           | Maximum number of retries when sending batches. Setting it to `0` will retry indefinitely.                                                                                                                                                                                                                                                                                                                                        | 10                                     |
| Labels               | labels for API requests.                                                                                                                                                                                                                                                                                                                                                                | {job="fluent-bit"}                     |
| LogLevel             | LogLevel for plugin logger.                                                                                                                                                                                                                                                                                                                                                             | "info"                                 |
| RemoveKeys           | Specify removing keys.                                                                                                                                                                                                                                                                                                                                                                  | none                                   |
| AutoKubernetesLabels | If set to true, it will add all Kubernetes labels to Logstore labels                                                                                                                                                                                                                                                                                                                        | false                                  |
| LabelKeys            | Comma separated list of keys to use as stream labels. All other keys will be placed into the log line. LabelKeys is deactivated when using `LabelMapPath` label mapping configuration.                                                                                                                                                                                                  | none                                   |
| LineFormat           | Format to use when flattening the record to a log line. Valid values are "json" or "key_value". If set to "json" the log line sent to Logstore will be the fluentd record (excluding any keys extracted out as labels) dumped as json. If set to "key_value", the log line will be each item in the record concatenated together (separated by a single space) in the format <key>=<value>. | json                                   |
| DropSingleKey        | If set to true and after extracting label_keys a record only has a single key remaining, the log line sent to Logstore will just be the value of the record key.                                                                                                                                                                                                                            | true                                   |
| LabelMapPath         | Path to a json file defining how to transform nested records.                                                                                                                                                                                                                                                                                                                           | none                                   |
| Buffer               | Enable buffering mechanism                                                                                                                                                                                                                                                                                                                                                              | false                                  |
| BufferType           | Specify the buffering mechanism to use (currently only dque is implemented).                                                                                                                                                                                                                                                                                                            | dque                                   |
| DqueDir              | Path to the directory for queued logs                                                                                                                                                                                                                                                                                                                                                   | /tmp/flb-storage/logstore                  |
| DqueSegmentSize      | Segment size in terms of number of records per segment                                                                                                                                                                                                                                                                                                                                  | 500                                    |
| DqueSync             | Whether to fsync each queue change. Specify no fsync with "normal", and fsync with "full".                                                                                                                                                                                                                                                                                                                                                      | "normal"                                  |
| DqueName             | Queue name, must be uniq per output                                                                                                                                                                                                                                                                                                                                                     | dque                                   |

### Labels

Labels are used to [query logs]({{< relref "../../query" >}}) `{container_name="nginx", cluster="us-west1"}`, they are usually metadata about the workload producing the log stream (`instance`, `container_name`, `region`, `cluster`, `level`).  In Logstore labels are indexed consequently you should be cautious when choosing them (high cardinality label values can have performance drastic impact).

You can use `Labels`, `RemoveKeys` , `LabelKeys` and `LabelMapPath` to how the output plugin will perform labels extraction.

### AutoKubernetesLabels

If set to true, it will add all Kubernetes labels to Logstore labels automatically and ignore parameters `LabelKeys`, LabelMapPath.

### LabelMapPath

When using the `Parser` and `Filter` plugins Fluent Bit can extract and add data to the current record/log data. While Logstore labels are key value pair, record data can be nested structures.
You can pass a JSON file that defines how to extract labels from each record. Each json key from the file will be matched with the log record to find label values. Values from the configuration are used as label names.

Considering the record below :

```json
{
  "kubernetes": {
    "container_name": "promtail",
    "pod_name": "promtail-xxx",
    "namespace_name": "prod",
    "labels" : {
        "team": "x-men"
    }
  },
  "HOSTNAME": "docker-desktop",
  "log" : "a log line",
  "time": "20190926T152206Z"
}
```

and a LabelMap file as follow :

```json
{
  "kubernetes": {
    "container_name": "container",
    "pod_name": "pod",
    "namespace_name": "namespace",
    "labels" : {
        "team": "team"
    }
  }
}
```

The labels extracted will be `{team="x-men", container="promtail", pod="promtail-xxx", namespace="prod"}`.

If you don't want the `kubernetes` and `HOSTNAME` fields to appear in the log line you can use the `RemoveKeys` configuration field. (e.g. `RemoveKeys kubernetes,HOSTNAME`).

### Buffering

Buffering refers to the ability to store the records somewhere, and while they are processed and delivered, still be able to store more. The Logstore output plugin can be blocked by the Logstore client because of its design:

- If the BatchSize is over the limit, the output plugin pauses receiving new records until the pending batch is successfully sent to the server
- If the Logstore server is unreachable (retry 429s, 500s and connection-level errors), the output plugin blocks new records until the Logstore server is available again, and the pending batch is successfully sent to the server or as long as the maximum number of attempts has been reached within configured back-off mechanism

The blocking state with some of the input plugins is not acceptable, because it can have an undesirable side effect on the part that generates the logs. Fluent Bit implements a buffering mechanism that is based on parallel processing. Therefore, it cannot send logs in order. There are two ways of handling the out-of-order logs: 

- Configure Logstore to [accept out-of-order writes]({{< relref "../../configure#accept-out-of-order-writes" >}}).

- Configure the Logstore output plugin to use the buffering mechanism based on [`dque`](https://github.com/joncrlsn/dque), which is compatible with the Logstore server strict time ordering:

    ```properties
    [Output]
        Name acme-logstore
        Match *
        Url http://localhost:3100/logstore/api/v1/push
        Buffer true
        DqueSegmentSize 8096
        DqueDir /tmp/flb-storage/buffer
        DqueName logstore.0
    ```

### Configuration examples

To configure the Logstore output plugin add this section to fluent-bit.conf

```properties
[Output]
    Name acme-logstore
    Match *
    Url http://localhost:3100/logstore/api/v1/push
    BatchWait 1s
    BatchSize 30720
    # (30KiB)
    Labels {test="fluent-bit-go", lang="Golang"}
    RemoveKeys key1,key2
    LabelKeys key3,key4
    LineFormat key_value
```

```properties
[Output]
    Name acme-logstore
    Match *
    Url http://localhost:3100/logstore/api/v1/push
    BatchWait 1s
    BatchSize 30720 # (30KiB)
    AutoKubernetesLabels true
    RemoveKeys key1,key2
```

A full [example configuration file](https://example.com/acme/logstore/blob/main/clients/cmd/fluent-bit/fluent-bit.conf) is also available in the Logstore repository.

### Running multiple plugin instances

You can run multiple plugin instances in the same fluent-bit process, for example if you want to push to different Logstore servers or route logs into different Logstore tenant IDs. To do so, add additional `[Output]` sections.
