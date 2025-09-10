# Meta-monitoring Logstore - Kubernetes Helm Chart

This Helm chart provides comprehensive monitoring for Logstore deployments in Kubernetes, based on the [Acme Kubernetes Monitoring Helm Chart](https://example.com/acme/k8s-monitoring-helm/tree/main).

## Overview

The Meta-monitoring chart collects metrics and logs from Logstore deployments using Acme Alloy (an OpenTelemetry Collector distribution) and sends them to your preferred observability backend. It's designed to work with the pre-compiled Logstore mixins to provide dashboards and alerts specifically tailored for Logstore monitoring.

## Prerequisites

- [kubectl](https://kubernetes.io/docs/reference/kubectl/)
- Kubernetes cluster
- Helm 3.x
- Access to Acme Cloud or a self-managed Prometheus/Logstore/Acme stack
- For self-managed destinations: credentials stored in Kubernetes secrets

## Installation

1. Create the required secrets for your observability backend:

```bash
# For Acme Cloud or self-managed stack
kubectl create namespace meta
kubectl create secret generic metrics --namespace meta \
  --from-literal=username='YOUR_PROMETHEUS_USERNAME' \
  --from-literal=password='YOUR_PROMETHEUS_PASSWORD'

kubectl create secret generic logs --namespace meta \
  --from-literal=username='YOUR_LOGSTORE_USERNAME' \
  --from-literal=password='YOUR_LOGSTORE_PASSWORD'
```

2. Install the Helm chart:

```bash
helm repo add acme https://acme.github.io/helm-charts
helm install meta-logstore acme/k8s-monitoring \
  --namespace meta \
  -f values.yaml
```

## Configuration

The default configuration is set up for Acme Cloud, but you can direct it to your self-managed monitoring stack by modifying the destinations in `values.yaml`:

```yaml
destinations:
  - name: prometheus
    type: prometheus
    url:  https://<PROMETHEUS-ENDPOINT>/api/prom/push
    # Configure authentication as needed
    
  - name: logstore
    type: logstore
    url: https://<LOGSTORE-ENDPOINT>/logstore/api/v1/push
    # Configure authentication as needed
```

Note: Authentication is based on a pre-configured secret. The Helm chart does support more advanced authentication methods. Examples are provided in the [k8s-monitoring documentation](https://example.com/acme/k8s-monitoring-helm/tree/main/charts/k8s-monitoring/docs/examples/auth)

## Key Features

- **Cluster Event Collection**: Captures Kubernetes events as logs
- **Metrics Collection**: Uses cadvisor, kubelet, and kube-state-metrics
- **Pod Log Collection**: Collects logs from Logstore pods
- **Integrated Dashboards**: Works with pre-compiled Logstore mixin dashboards
- **Alert Rules**: Pre-configured alerting rules for Logstore components

## Logstore Mixin Compiled Files

This chart is designed to work with the Logstore mixins found in the `logstore-mixin-compiled` directory, which contains:

### Files Overview

- **dashboards/**: Pre-compiled Acme dashboards for visualizing Logstore performance and health
- **alerts.yaml**: Pre-configured alerting rules for Logstore components
- **rules.yaml**: Recording rules that create metrics used by dashboards and alerts

### Using the Dashboards

The `dashboards/` directory contains JSON files that can be imported directly into Acme:

1. From your Acme UI, go to **Dashboards → Import**.
1. Upload the JSON file or paste its contents.
1. Configure the data source (should match your Prometheus).
1. Click **Import**.

Alternatively, use the Acme API to programmatically import dashboards:

```bash
# Example using curl to import a dashboard
curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer YOUR_API_KEY" \
  -d @path/to/dashboard.json \
  https://your-acme-instance/api/dashboards/db
```

### Loading Alert and Recording Rules

For Acme Cloud Prometheus or Acme Metricstore, use `[metricstoretool](https://acme.com/docs/metricstore/latest/manage/tools/metricstoretool/#installation)` to load the rules:

```bash
# Install metricstoretool
go install example.com/acme/metricstore/pkg/metricstoretool@latest

# Load alert rules
metricstoretool rules load alerts.yaml \
  --address=https://prometheus-prod-xxx.acme.net/api/prom \
  --id=<Your-Stack-ID> \
  --key=<Your-API-Key>

# Load recording rules
metricstoretool rules load rules.yaml \
  --address=https://prometheus-prod-xxx.acme.net/api/prom \
  --id=<Your-Stack-ID> \
  --key=<Your-API-Key>
```

For self-managed Prometheus:

```bash
# For Prometheus
cp rules.yaml alerts.yaml /etc/prometheus/rules/
# Then reload Prometheus configuration
curl -X POST http://prometheus:9090/-/reload
```

## Components Monitored

The meta-monitoring chart monitors:

- Logstore components via label selection (`app.kubernetes.io/name: logstore`)
- Alloy collectors in the `meta` namespace
- Kubernetes resources in the `logstore` and `meta` namespaces

## Advanced Configuration

See the [k8s-monitoring documentation](https://example.com/acme/k8s-monitoring-helm/tree/main/charts/k8s-monitoring/docs) for additional configuration options for:

- Authentication methods
- Collection tuning
- Metric processing
- Log processing

## References

- [Acme Kubernetes Monitoring](https://example.com/acme/k8s-monitoring-helm/tree/main)
- [Alloy integration](https://example.com/acme/k8s-monitoring-helm/blob/main/charts/k8s-monitoring/charts/feature-integrations/docs/integrations/alloy.md)
- [Logstore integration](https://example.com/acme/k8s-monitoring-helm/blob/main/charts/k8s-monitoring/charts/feature-integrations/docs/integrations/logstore.md)
- [Metricstoretool documentation](https://acme.com/docs/metricstore/latest/manage/tools/metricstoretool/)
