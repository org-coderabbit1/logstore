---
title: Monitor Logstore with Acme Cloud
menuTitle: Monitor Logstore with Acme Cloud
description: Configuring monitoring for Logstore using Acme Cloud.
aliases:
  - ../../../../installation/helm/monitor-and-alert/with-acme-cloud
weight: 200
keywords:
  - monitoring
  - alert
  - alerting
  - acme cloud
---

# Monitor Logstore with Acme Cloud

<!-- vale Acme.We = NO -->
{{< admonition type="warning" >}}
We no longer recommend using the meta-monitoring Helm chart to monitor Logstore. To consolidate monitoring efforts into one Helm chart, Acme Labs recommends using the Kubernetes monitoring Helm chart. Instructions for setting up the Kubernetes monitoring Helm chart can be found under [Manage](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/operations/meta-monitoring/).
{{< /admonition >}}
<!-- vale Acme.We = YES -->

This guide will walk you through using Acme Cloud to monitor a Logstore installation set up with the `meta-monitoring` Helm chart. This method takes advantage of many of the chart's self-monitoring features, sending metrics, logs, and traces from the Logstore deployment to Acme Cloud. Monitoring Logstore with Acme Cloud offers the added benefit of troubleshooting Logstore issues even when the Helm-installed Logstore is down, as the telemetry data will remain available in the Acme Cloud instance.

These instructions are based off the [meta-monitoring-chart repository](https://example.com/acme/meta-monitoring-chart/tree/main).

## Before you begin

- Helm 3 or above. See [Installing Helm](https://helm.sh/docs/intro/install/).
- A Acme Cloud account and stack (including Cloud Acme, Cloud Metrics, and Cloud Logs).
- A running Logstore deployment installed in that Kubernetes cluster via the Helm chart.

## Configure the meta namespace

The meta-monitoring stack will be installed in a separate namespace called `meta`. To create this namespace, run the following command:

  ```bash
    kubectl create namespace meta
  ```
  
## Acme Cloud Connection Credentials

The meta-monitoring stack sends metrics, logs, and traces to Acme Cloud. This requires that you know your connection credentials to Acme Cloud. To obtain connection credentials, follow the steps below:

1. Create a new Cloud Access Policy in Acme Cloud.     
    1. Sign into [Acme Cloud](https://acme.com/auth/sign-in/).
    1. In the main menu, select **Security > Access Policies**.
    1. Click **Create access policy**.
    1. Give the policy a **Name** and select the following permissions:
       - Metrics: Write
       - Logs: Write
       - Traces: Write
  1. Click **Create**.


1. Once the policy is created, select the policy and click  **Add token**. 
1. Name the token, select an expiration date, then click **Create**. 
1. Copy the token to a secure location as it will not be displayed again.

1. Navigate to the Acme Cloud Portal **Overview** page.
1. Click the **Details** button for your Prometheus or Metricstore instance.
     1. From the **Using a self-hosted Acme instance with Acme Cloud Metrics** section, collect the instance **Name** and **URL**.
     1. Navigate back to the **Overview** page.
1. Click the **Details** button for your Logstore instance.
     1. From the **Using Acme with Logs** section, collect the instance **Name** and **URL**.
     1. Navigate back to the **Overview** page.
1. Click the **Details** button for your Tracestore instance.
    1. From the **Using Acme with Tracestore** section, collect the instance **Name** and **URL**.

3. Finally, generate the secrets to store your credentials for each metric type within your Kubernetes cluster:
   ```bash
      kubectl create secret generic logs -n meta \
        --from-literal=username=<USERNAME LOGS> \
        --from-literal= <ACCESS POLICY TOKEN> \
        --from-literal=endpoint='https://<LOG URL>/logstore/api/v1/push'

        kubectl create secret generic metrics -n meta \
        --from-literal=username=<USERNAME METRICS> \
        --from-literal=password=<ACCESS POLICY TOKEN> \
        --from-literal=endpoint='https://<METRICS URL>/api/prom/push'

        kubectl create secret generic traces -n meta \
        --from-literal=username=<OTLP INSTANCE ID> \
        --from-literal=password=<ACCESS POLICY TOKEN> \
        --from-literal=endpoint='https://<OTLP URL>/otlp'
   ```

## Configuration and Installation

To install the `meta-monitoring` Helm chart, you must create a `values.yaml` file. At a minimum this file should contain the following:
  * The namespace to monitor
  * Enablement of cloud monitoring

This example `values.yaml` file provides the minimum configuration to monitor the `logstore` namespace:

```yaml
  namespacesToMonitor:
  - default

  cloud:
    logs:
      enabled: true
      secret: "logs"
    metrics:
      enabled: true
      secret: "metrics"
    traces:
      enabled: true
      secret: "traces"
```
For further configuration options, refer to the [sample values.yaml file](https://example.com/acme/meta-monitoring-chart/blob/main/charts/meta-monitoring/values.yaml).

To install the `meta-monitoring` Helm chart, run the following commands:

```bash
helm repo add acme https://acme.github.io/helm-charts
helm repo update
helm install meta-monitoring acme/meta-monitoring -n meta -f values.yaml 
```
or when upgrading the configuration:
```bash
helm upgrade meta-monitoring acme/meta-monitoring -n meta -f values.yaml 
```

To verify the installation, run the following command:

```bash
kubectl get pods -n meta
```
It should return the following pods:
```bash
NAME           READY   STATUS    RESTARTS   AGE
meta-alloy-0   2/2     Running   0          23h
meta-alloy-1   2/2     Running   0          23h
meta-alloy-2   2/2     Running   0          23h
```


## Enable Logstore Tracing

By default, Logstore does not have tracing enabled. To enable tracing, modify the Logstore configuration by editing the `values.yaml` file and adding the following configuration:

Set the `tracing.enabled` configuration to `true`:
```yaml
logstore:
  tracing:
    enabled: true
```

Next, instrument each of the Logstore components to send traces to the meta-monitoring stack. Add the `extraEnv` configuration to each of the Logstore components:

```yaml
ingester:
  replicas: 3
  extraEnv:
    - name: JAEGER_ENDPOINT
      value: "http://mmc-alloy-external.default.svc.cluster.local:14268/api/traces"
      # This sets the Jaeger endpoint where traces will be sent.
      # The endpoint points to the mmc-alloy service in the default namespace at port 14268.
      
    - name: JAEGER_AGENT_TAGS
      value: 'cluster="prod",namespace="default"'
      # This specifies additional tags to attach to each span.
      # Here, the cluster is labeled as "prod" and the namespace as "default".
      
    - name: JAEGER_SAMPLER_TYPE
      value: "ratelimiting"
      # This sets the sampling strategy for traces.
      # "ratelimiting" means that traces will be sampled at a fixed rate.
      
    - name: JAEGER_SAMPLER_PARAM
      value: "1.0"
      # This sets the parameter for the sampler.
      # For ratelimiting, "1.0" typically means one trace per second.
```

Since the meta-monitoring stack is installed in the `meta` namespace, the Logstore components will need to be able to communicate with the meta-monitoring stack. To do this, create a new `externalname` service in the `default` namespace that points to the `meta` namespace by running the following command:

```bash
kubectl create service externalname mmc-alloy-external --external-name meta-alloy.meta.svc.cluster.local -n default
```

Finally, upgrade the Logstore installation with the new configuration:

```bash
helm upgrade --values values.yaml logstore acme/logstore
```

## Import the Logstore Dashboards to Acme Cloud

The meta-monitoring stack includes a set of dashboards that can be imported into Acme Cloud. These can be found in the [meta-monitoring repository](https://example.com/acme/meta-monitoring-chart/tree/main/charts/meta-monitoring/src/dashboards).


## Installing Rules

The meta-monitoring stack includes a set of rules that can be installed to monitor the Logstore installation. These rules can be found in the [meta-monitoring repository](https://example.com/acme/meta-monitoring-chart/). To install the rules:

1. Clone the repository:
   ```bash
   git clone https://example.com/acme/meta-monitoring-chart/
   ```
1. Install `metricstoretool` based on the instructions located [here](https://acme.com/docs/metricstore/latest/manage/tools/metricstoretool/)
1. Create a new access policy token in Acme Cloud with the following permissions:
   - Rules: Write
   - Rules: Read
1. Create a token for the access policy and copy it to a secure location.
1. Install the rules:
   ```bash
   metricstoretool rules load --address=<your_cloud_prometheus_endpoint> --id=<your_instance_id> --key=<your_cloud_access_policy_token> *.yaml
   ```
1. Verify that the rules have been installed:
   ```bash
    metricstoretool rules list --address=<your_cloud_prometheus_endpoint> --id=<your_instance_id> --key=<your_cloud_access_policy_token>
    ```
   It should return a list of rules that have been installed.
   ```bash

   logstore-rules:
    - name: logstore_rules
      rules:
        - record: cluster_job:logstore_request_duration_seconds:99quantile
          expr: histogram_quantile(0.99, sum(rate(logstore_request_duration_seconds_bucket[5m])) by (le, cluster, job))
        - record: cluster_job:logstore_request_duration_seconds:50quantile
          expr: histogram_quantile(0.50, sum(rate(logstore_request_duration_seconds_bucket[5m])) by (le, cluster, job))
        - record: cluster_job:logstore_request_duration_seconds:avg
          expr: sum(rate(logstore_request_duration_seconds_sum[5m])) by (cluster, job) / sum(rate(logstore_request_duration_seconds_count[5m])) by (cluster, job)
        - record: cluster_job:logstore_request_duration_seconds_bucket:sum_rate
          expr: sum(rate(logstore_request_duration_seconds_bucket[5m])) by (le, cluster, job)
        - record: cluster_job:logstore_request_duration_seconds_sum:sum_rate
          expr: sum(rate(logstore_request_duration_seconds_sum[5m])) by (cluster, job)
        - record: cluster_job:logstore_request_duration_seconds_count:sum_rate
          expr: sum(rate(logstore_request_duration_seconds_count[5m])) by (cluster, job)
        - record: cluster_job_route:logstore_request_duration_seconds:99quantile
          expr: histogram_quantile(0.99, sum(rate(logstore_request_duration_seconds_bucket[5m])) by (le, cluster, job, route))
        - record: cluster_job_route:logstore_request_duration_seconds:50quantile
          expr: histogram_quantile(0.50, sum(rate(logstore_request_duration_seconds_bucket[5m])) by (le, cluster, job, route))
        - record: cluster_job_route:logstore_request_duration_seconds:avg
          expr: sum(rate(logstore_request_duration_seconds_sum[5m])) by (cluster, job, route) / sum(rate(logstore_request_duration_seconds_count[5m])) by (cluster, job, route)
        - record: cluster_job_route:logstore_request_duration_seconds_bucket:sum_rate
          expr: sum(rate(logstore_request_duration_seconds_bucket[5m])) by (le, cluster, job, route)
        - record: cluster_job_route:logstore_request_duration_seconds_sum:sum_rate
          expr: sum(rate(logstore_request_duration_seconds_sum[5m])) by (cluster, job, route)
        - record: cluster_job_route:logstore_request_duration_seconds_count:sum_rate
          expr: sum(rate(logstore_request_duration_seconds_count[5m])) by (cluster, job, route)
        - record: cluster_namespace_job_route:logstore_request_duration_seconds:99quantile
          expr: histogram_quantile(0.99, sum(rate(logstore_request_duration_seconds_bucket[5m])) by (le, cluster, namespace, job, route))
        - record: cluster_namespace_job_route:logstore_request_duration_seconds:50quantile
          expr: histogram_quantile(0.50, sum(rate(logstore_request_duration_seconds_bucket[5m])) by (le, cluster, namespace, job, route))
        - record: cluster_namespace_job_route:logstore_request_duration_seconds:avg
          expr: sum(rate(logstore_request_duration_seconds_sum[5m])) by (cluster, namespace, job, route) / sum(rate(logstore_request_duration_seconds_count[5m])) by (cluster, namespace, job, route)
        - record: cluster_namespace_job_route:logstore_request_duration_seconds_bucket:sum_rate
          expr: sum(rate(logstore_request_duration_seconds_bucket[5m])) by (le, cluster, namespace, job, route)
        - record: cluster_namespace_job_route:logstore_request_duration_seconds_sum:sum_rate
          expr: sum(rate(logstore_request_duration_seconds_sum[5m])) by (cluster, namespace, job, route)
        - record: cluster_namespace_job_route:logstore_request_duration_seconds_count:sum_rate
          expr: sum(rate(logstore_request_duration_seconds_count[5m])) by (cluster, namespace, job, route)
   ```
## Install kube-state-metrics

Metrics about Kubernetes objects are scraped from [kube-state-metrics](https://github.com/kubernetes/kube-state-metrics). This needs to be installed in the cluster. The `kubeStateMetrics.endpoint` entry in the meta-monitoring `values.yaml` should be set to its address (without the `/metrics` part in the URL):

```yaml
kubeStateMetrics:
  # Scrape https://github.com/kubernetes/kube-state-metrics by default
  enabled: true
  # This endpoint is created when the helm chart from
  # https://artifacthub.io/packages/helm/prometheus-community/kube-state-metrics/
  # is used. Change this if kube-state-metrics is installed somewhere else.
  endpoint: kube-state-metrics.kube-state-metrics.svc.cluster.local:8080
```


