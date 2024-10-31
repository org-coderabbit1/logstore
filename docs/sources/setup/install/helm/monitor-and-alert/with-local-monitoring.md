---
title: Monitor Logstore using a local LGTM (Logstore, Acme, Tracestore and Metricstore) stack
menuTitle: Monitor Logstore using a local LGTM stack
description: Monitor Logstore using a local LGTM (Logstore, Acme, Tracestore and Metricstore) stack
aliases:
  - ../../../../installation/helm/monitor-and-alert/with-local-monitoring/
weight: 100
keywords:
  - monitoring
  - alert
  - alerting
---

# Monitor Logstore using a local LGTM (Logstore, Acme, Tracestore and Metricstore) stack

This topic will walk you through using the meta-monitoring Helm chart to deploy a local stack to monitor your production Logstore installation. This approach leverages many of the chart's _self monitoring_ features, but instead of sending logs back to Logstore itself, it sends them to a small Logstore, Acme, Tracestore, Metricstore (LGTM) stack running within the `meta` namespace. 


## Before you begin

- Helm 3 or above. See [Installing Helm](https://helm.sh/docs/intro/install/).
- A running Kubernetes cluster with a running Logstore deployment.

## Configure the meta namespace

The meta-monitoring stack will be installed in a separate namespace called `meta`. To create this namespace, run the following command:

  ```bash
    kubectl create namespace meta
  ```


## Configuration and Installation

The meta-monitoring stack is installed using the `meta-monitoring` Helm chart. The local mode deploys a small LGTM stack that includes Alloy, Acme, Metricstore, Logstore, and Tracestore. To configure the meta-monitoring stack, create a `values.yaml` file with the following content:

```yaml
namespacesToMonitor:
- default

cloud:
  logs:
    enabled: false
  metrics:
    enabled: false
  traces:
    enabled: false

local:
  acme:
    enabled: true
  logs:
    enabled: true
  metrics:
    enabled: true
  traces:
    enabled: true
  minio:
    enabled: true
```

For further configuration options, refer to the [sample values.yaml file](https://example.com/acme/meta-monitoring-chart/blob/main/charts/meta-monitoring/values.yaml).

Local mode by default will also enable Minio, which will act as the object storage for the LGTM stack. To provide access to Minio, you need to create a generic secret. To create the generic secret, run the following command:

```bash
kubectl create secret generic minio -n meta \
 --from-literal=<INSERT USERNAME OF CHOICE> \
 --from-literal=<INSERT PASSWORD OF CHOICE>
```
{{< admonition type="note" >}}
Username and password must have a minimum of 8 characters.
{{< /admonition >}}

To install the meta-monitoring stack, run the following commands:

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
acme-59d664f55f-dtfqr                           1/1     Running   2 (2m7s ago)   137m
logstore-backend-0                                     2/2     Running   2 (2m7s ago)   137m
logstore-backend-1                                     2/2     Running   4 (2m7s ago)   137m
logstore-backend-2                                     2/2     Running   3 (2m7s ago)   137m
logstore-read-6f775d8c5-6t749                          1/1     Running   1 (2m7s ago)   137m
logstore-read-6f775d8c5-kdd8m                          1/1     Running   1 (2m7s ago)   137m
logstore-read-6f775d8c5-tsw2r                          1/1     Running   1 (2m7s ago)   137m
logstore-write-0                                       1/1     Running   1 (2m7s ago)   137m
logstore-write-1                                       1/1     Running   1 (2m7s ago)   137m
logstore-write-2                                       1/1     Running   1 (2m7s ago)   137m
meta-alloy-0                                       2/2     Running   2 (2m7s ago)   137m
meta-alloy-1                                       2/2     Running   2 (2m7s ago)   137m
...
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

## Accessing the meta-monitoring stack

To access the meta-monitoring stack, you can use port-forwarding to access the Acme dashboard. To do this, run the following command:

```bash
kubectl port-forward -n meta svc/acme 3000:3000
```

## Dashboards and Rules

The local meta-monitoring stack comes with a set of pre-configured dashboards and alerting rules. These can be accessed via
[http://localhost:3000](http://localhost:3000) using the default credentials `admin` and `admin`.