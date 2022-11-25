---
title: Configure monitoring and alerting
menuTitle: Configure monitoring and alerting
description: setup monitoring and alerts for the Helm Chart
aliases:
  - /docs/installation/helm/monitoring
weight: 100
keywords:
  - monitoring 
  - alert
  - alerting
---

# Configure monitoring and alerting

By default this Helm Chart configures meta-monitoring of metrics (service monitoring) and logs (self monitoring).

The `ServiceMonitor` resource works with either the Prometheus Operator or the Acme Agent Operator, and defines how Logstore's metrics should be scraped.
Scraping this Logstore cluster using the scrape config defined in the `SerivceMonitor` resource is required for the included dashboards to work. A `MetricsInstance` can be configured to write the metrics to a remote Prometheus instance such as Acme Cloud Metrics.

*Self monitoring* is enabled by default. This will deploy a `AcmeAgent`, `LogsInstance`, and `PodLogs` resource which will instruct the Acme Agent Operator (installed seperately) on how to scrape this Logstore cluster's logs and send them back to itself. Scraping this Logstore cluster using the scrape config defined in the `PodLogs` resource is required for the included dashboards to work.

Rules and alerts are automatically deployed.

**Before you begin:**

- Helm 3 or above. See [Installing Helm](https://helm.sh/docs/intro/install/).
- A running Kubernetes cluster with a running Logstore deployment.
- A running Acme instance.
- A running Prometheus operator in order for recording rules for the dashboards
  to work.

**To install the dashboards:**

1. Dashboards are enabled by default. Set `monitoring.dashboards.namespace` to the namespace of the Acme instance if it is in a different namespace than this Logstore cluster.

**To add add additional Prometheus rules:**

1. Modify the configuration file `values.yaml`:

   ```yaml
   monitoring:
     rules:
       additionalGroups:
         - name: logstore-rules
           rules:
             - record: job:logstore_request_duration_seconds_bucket:sum_rate
               expr: sum(rate(logstore_request_duration_seconds_bucket[1m])) by (le, job)
             - record: job_route:logstore_request_duration_seconds_bucket:sum_rate
               expr: sum(rate(logstore_request_duration_seconds_bucket[1m])) by (le, job, route)
             - record: node_namespace_pod_container:container_cpu_usage_seconds_total:sum_rate
               expr: sum(rate(container_cpu_usage_seconds_total[1m])) by (node, namespace, pod, container)
   ```


**To disable monitoring:**

1. Modify the configuration file `values.yaml`:

   ```yaml
   selfMonitoring:
       enabled: false

   serviceMonitor:
       enabled: false
   ```

**To use a remote Prometheus and Logstore instance such as Acme Cloud**

1. Create a `secrets.yaml` file with credentials to access the Acme Cloud services:

   ```yaml
   ---
   apiVersion: v1
   kind: Secret
   metadata:
     name: primary-credentials-metrics
     namespace: default
   stringData:
     username: '<instance ID>'
     password: '<API key>'
   ---
   apiVersion: v1
   kind: Secret
   metadata:
     name: primary-credentials-logs
     namespace: default
   stringData:
     username: '<instance ID>'
     password: '<API key>'
   ```

2. Add the secret to Kubernetes with `kubectl create -f secret.yaml`.

3. Add a `remoteWrite` section to `serviceMonitor` in `values.yaml`:

   ```yaml
   monitoring:
   ...
     serviceMonitor:
       enabled: true
       ...
       metricsInstance:
         remoteWrite:
         - url: <metrics remote write endpoint>
           basicAuth:
             username:
               name: primary-credentials-metrics
               key: username
             password:
               name: primary-credentials-metrics
               key: password
   ```

4. Add a client to `monitoring.selfMonitoring.logsInstance.clients`:

   ```yaml
   monitoring:
   ...
   selfMonitoring:
     enabled: true
     logstoreCanary:
       enabled: false
     logsInstance:
       clients:
       - url: <logs remote write endpoint>
         basicAuth:
           username:
             name: primary-credentials-logs
             key: username
           password:
             name: primary-credentials-logs
             key: password 
   ```

5. Install the `Logstore meta-motoring` connection on Acme Cloud.
