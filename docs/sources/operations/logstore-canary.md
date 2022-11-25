---
title: Logstore Canary
weight: 60
---
# Logstore Canary

![canary](../canary.png)

Logstore Canary is a standalone app that audits the log-capturing performance of
a Acme Logstore cluster.

Logstore Canary generates artificial log lines.
These log lines are sent to the Logstore cluster.
Logstore Canary communicates with the Logstore cluster to capture metrics about the
artificial log lines,
such that Logstore Canary forms information about the performance of the
Logstore cluster.
The information is available as Prometheus time series metrics.

![block_diagram](../logstore-canary-block.png)

Logstore Canary writes a log to a file and stores the timestamp in an internal
array. The contents look something like this:

```nohighlight
1557935669096040040 ppppppppppppppppppppppppppppppppppppppppppppppppppppppppppp
```

The relevant part of the log entry is the timestamp; the `p`s are just filler
bytes to make the size of the log configurable.

An agent (like Promtail) should be configured to read the log file and ship it
to Logstore.

Meanwhile, Logstore Canary will open a WebSocket connection to Logstore and will tail
the logs it creates. When a log is received on the WebSocket, the timestamp
in the log message is compared to the internal array.

If the received log is:

- The next in the array to be received, it is removed from the array and the
  (current time - log timestamp) is recorded in the `response_latency`
  histogram. This is the expected behavior for well behaving logs.
- Not the next in the array to be received, it is removed from the array, the
  response time is recorded in the `response_latency` histogram, and the
  `out_of_order_entries` counter is incremented.
- Not in the array at all, it is checked against a separate list of received
  logs to either increment the `duplicate_entries` counter or the
  `unexpected_entries` counter.

In the background, Logstore Canary also runs a timer which iterates through all of
the entries in the internal array. If any of the entries are older than the
duration specified by the `-wait` flag (defaulting to 60s), they are removed
from the array and the `websocket_missing_entries` counter is incremented. An
additional query is then made directly to Logstore for any missing entries to
determine if they are truly missing or only missing from the WebSocket. If
missing entries are not found in the direct query, the `missing_entries` counter
is incremented.

### Additional Queries

#### Spot Check

Starting with version 1.6.0, the canary will spot check certain results over time
to make sure they are present in Logstore, this is helpful for testing the transition
of inmemory logs in the ingester to the store to make sure nothing is lost.

`-spot-check-interval` and `-spot-check-max` are used to tune this feature,
`-spot-check-interval` will pull a log entry from the stream at this interval
and save it in a separate list up to `-spot-check-max`.

Every `-spot-check-query-rate`, Logstore will be queried for each entry in this list and
`logstore_canary_spot_check_entries_total` will be incremented, if a result
is missing `logstore_canary_spot_check_missing_entries_total` will be incremented.

The defaults of `15m` for `spot-check-interval` and `4h` for `spot-check-max`
means that after 4 hours of running the canary will have a list of 16 entries
it will query every minute (default `spot-check-query-rate` interval is 1m),
so be aware of the query load this can put on Logstore if you have a lot of canaries.

__NOTE:__ if you are using `out-of-order-percentage` to test ingestion of out-of-order
log lines be sure not to set the two out of order time range flags too far in the past.
The defaults are already enough to test this functionality properly, and setting them
too far in the past can cause issues with the spot check test.


When using `out-of-order-percentage` you also need to make use of pipeline stages
in your Promtail configuration in order to set the timestamps correctly as the logs are pushed
to Logstore. The `client/promtail/pipelines` docs have examples of how to do this.

#### Metric Test

Logstore Canary will run a metric query `count_over_time` to
verify that the rate of logs being stored in Logstore corresponds to the rate they are being
created by Logstore Canary.

`-metric-test-interval` and `-metric-test-range` are used to tune this feature, but
by default every `15m` the canary will run a `count_over_time` instant-query to Logstore
for a range of `24h`.

If the canary has not run for `-metric-test-range` (`24h`) the query range is adjusted
to the amount of time the canary has been running such that the rate can be calculated
since the canary was started.

The canary calculates what the expected count of logs would be for the range
(also adjusting this based on canary runtime) and compares the expected result with
the actual result returned from Logstore.  The _difference_ is stored as the value in
the gauge `logstore_canary_metric_test_deviation`

It's expected that there will be some deviation, the method of creating an expected
calculation based on the query rate compared to actual query data is imperfect
and will lead to a deviation of a few log entries.

It's not expected for there to be a deviation of more than 3-4 log entries.

### Control

Logstore Canary responds to two endpoints to allow dynamic suspending/resuming of the
canary process.  This can be useful if you'd like to quickly disable or reenable the
canary.  To stop or start the canary issue an HTTP GET request against the `/suspend` or
`/resume` endpoints.

## Installation

### Binary

Logstore Canary is provided as a pre-compiled binary as part of the
[Logstore Releases](https://example.com/acme/logstore/releases) on GitHub.

### Docker

Logstore Canary is also provided as a Docker container image:

```bash
# change tag to the most recent release
$ docker pull acme/logstore-canary:2.0.0
```

### Kubernetes

To run on Kubernetes, you can do something simple like:

`kubectl run logstore-canary --generator=run-pod/v1
--image=acme/logstore-canary:latest --restart=Never --image-pull-policy=IfNotPresent
--labels=name=logstore-canary -- -addr=logstore:3100`

Or you can do something more complex like deploy it as a DaemonSet, there is a
Tanka setup for this in the `production` folder, you can import it using
`jsonnet-bundler`:

```shell
jb install example.com/acme/logstore-canary/production/ksonnet/logstore-canary
```

Then in your Tanka environment's `main.jsonnet` you'll want something like
this:

```jsonnet
local logstore_canary = import 'logstore-canary/logstore-canary.libsonnet';

logstore_canary {
  logstore_canary_args+:: {
    addr: "logstore:3100",
    port: 80,
    labelname: "instance",
    interval: "100ms",
    size: 1024,
    wait: "3m",
  },
  _config+:: {
    namespace: "default",
  }
}
```
#### Examples

Standalone Pod Implementation of logstore-canary

```
---
apiVersion: v1
kind: Pod
metadata:
  labels:
    app: logstore-canary
    name: logstore-canary
  name: logstore-canary
spec:
  containers:
  - args:
    - -addr=logstore:3100
    image: acme/logstore-canary:latest
    imagePullPolicy: IfNotPresent
    name: logstore-canary
    resources: {}
---
apiVersion: v1
kind: Service
metadata:
  name: logstore-canary
  labels:
    app: logstore-canary
spec:
  type: ClusterIP
  selector:
    app: logstore-canary
  ports:
  - name: metrics
    protocol: TCP
    port: 3500
    targetPort: 3500
```

DaemonSet Implementation of logstore-canary

```
---
kind: DaemonSet
apiVersion: extensions/v1beta1
metadata:
  labels:
    app: logstore-canary
    name: logstore-canary
  name: logstore-canary
spec:
  template:
    metadata:
      name: logstore-canary
      labels:
        app: logstore-canary
    spec:
      containers:
      - args:
        - -addr=logstore:3100
        image: acme/logstore-canary:latest
        imagePullPolicy: IfNotPresent
        name: logstore-canary
        resources: {}
---
apiVersion: v1
kind: Service
metadata:
  name: logstore-canary
  labels:
    app: logstore-canary
spec:
  type: ClusterIP
  selector:
    app: logstore-canary
  ports:
  - name: metrics
    protocol: TCP
    port: 3500
    targetPort: 3500
```


### From Source

If the other options are not sufficient for your use case, you can compile
`logstore-canary` yourself:

```bash
# clone the source tree
$ git clone https://example.com/acme/logstore

# build the binary
$ make logstore-canary

# (optionally build the container image)
$ make logstore-canary-image
```

## Configuration

The address of Logstore must be passed in with the `-addr` flag, and if your Logstore
server uses TLS, `-tls=true` must also be provided. Note that using TLS will
cause the WebSocket connection to use `wss://` instead of `ws://`.

The `-labelname` and `-labelvalue` flags should also be provided, as these are
used by Logstore Canary to filter the log stream to only process logs for the
current instance of the canary. Ensure that the values provided to the flags are
unique to each instance of Logstore Canary. Acme Labs' Tanka config
accomplishes this by passing in the pod name as the label value.

If Logstore Canary reports a high number of `unexpected_entries`, Logstore Canary may
not be waiting long enough and the value for the `-wait` flag should be
increased to a larger value than 60s.

__Be aware__ of the relationship between `pruneinterval` and the `interval`.
For example, with an interval of 10ms (100 logs per second) and a prune interval
of 60s, you will write 6000 logs per minute. If those logs were not received
over the WebSocket, the canary will attempt to query Logstore directly to see if
they are completely lost. __However__ the query return is limited to 1000
results so you will not be able to return all the logs even if they did make it
to Logstore.

__Likewise__, if you lower the `pruneinterval` you risk causing a denial of
service attack as all your canaries attempt to query for missing logs at
whatever your `pruneinterval` is defined at.

All options:

```
  -addr string
    	The Logstore server URL:Port, e.g. logstore:3100
  -buckets int
    	Number of buckets in the response_latency histogram (default 10)
  -ca-file string
    	Client certificate authority for optional use with TLS connection to Logstore
  -cert-file string
    	Client PEM encoded X.509 certificate for optional use with TLS connection to Logstore
  -insecure
    	Allow insecure TLS connections
  -interval duration
    	Duration between log entries (default 1s)
  -key-file string
    	Client PEM encoded X.509 key for optional use with TLS connection to Logstore
  -labelname string
    	The label name for this instance of logstore-canary to use in the log selector (default "name")
  -labelvalue string
    	The unique label value for this instance of logstore-canary to use in the log selector (default "logstore-canary")
  -max-wait duration
    	Duration to keep querying Logstore for missing websocket entries before reporting them missing (default 5m0s)
  -metric-test-interval duration
    	The interval the metric test query should be run (default 1h0m0s)
  -metric-test-range duration
    	The range value [24h] used in the metric test instant-query. Note: this value is truncated to the running time of the canary until this value is reached (default 24h0m0s)
  -out-of-order-max duration
    	Maximum amount of time to go back for out of order entries (in seconds). (default 1m0s)
  -out-of-order-min duration
    	Minimum amount of time to go back for out of order entries (in seconds). (default 30s)
  -out-of-order-percentage int
    	Percentage (0-100) of log entries that should be sent out of order.
  -pass string
    	Logstore password. This credential should have both read and write permissions to Logstore endpoints
  -port int
    	Port which logstore-canary should expose metrics (default 3500)
  -pruneinterval duration
    	Frequency to check sent vs received logs, also the frequency which queries for missing logs will be dispatched to logstore (default 1m0s)
  -push
    	Push the logs directly to given Logstore address
  -query-timeout duration
    	How long to wait for a query response from Logstore (default 10s)
  -size int
    	Size in bytes of each log line (default 100)
  -spot-check-initial-wait duration
    	How long should the spot check query wait before starting to check for entries (default 10s)
  -spot-check-interval duration
    	Interval that a single result will be kept from sent entries and spot-checked against Logstore, e.g. 15min default one entry every 15 min will be saved and then queried again every 15min until spot-check-max is reached (default 15m0s)
  -spot-check-max duration
    	How far back to check a spot check entry before dropping it (default 4h0m0s)
  -spot-check-query-rate duration
    	Interval that the canary will query Logstore for the current list of all spot check entries (default 1m0s)
  -streamname string
    	The stream name for this instance of logstore-canary to use in the log selector (default "stream")
  -streamvalue string
    	The unique stream value for this instance of logstore-canary to use in the log selector (default "stdout")
  -tenant-id string
    	Tenant ID to be set in X-Scope-OrgID header.
  -tls
    	Does the logstore connection use TLS?
  -user string
    	Logstore username.
  -version
    	Print this builds version information
  -wait duration
    	Duration to wait for log entries on websocket before querying logstore for them (default 1m0s)
  -write-max-backoff duration
    	Maximum backoff time between retries  (default 5m0s)
  -write-max-retries int
    	Maximum number of retries when push a log entry  (default 10)
  -write-min-backoff duration
    	Initial backoff time before first retry  (default 500ms)
  -write-timeout duration
    	How long to wait write response from Logstore (default 10s)
```
