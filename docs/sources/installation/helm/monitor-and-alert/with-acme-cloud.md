---
title: Configure monitoring and alerting of Logstore using Acme Cloud
menuTitle: Monitor Logstore with Acme Cloud
description: setup monitoring and alerts for Logstore using Acme Cloud
aliases:
  - /docs/installation/helm/monitoring/with-acme-cloud
weight: 200
keywords:
  - monitoring
  - alert
  - alerting
  - acme cloud
---

# Configure monitoring and alerting of Logstore using Acme Cloud

This topic will walk you through using Acme Cloud to monitor a Logstore installation that is installed with the Helm chart. This approach leverages many of the chart's _self monitoring_ features, but instead of sending logs back to Logstore itself, it sends them to a Acme Cloud Logs instance. This approach also does not require the installation of the Prometheus Operator and instead sends metrics to a Acme Cloud Metrics instance. Using Acme Cloud to monitor Logstore has the added benefit of being able to troubleshoot problems with Logstore when the Helm installed Logstore is down, as the logs will still be available in the Acme Cloud Logs instance.

**Before you begin:**

- Helm 3 or above. See [Installing Helm](https://helm.sh/docs/intro/install/).
- A Acme Cloud account and stack (including Cloud Acme, Cloud Metrics, and Cloud Logs)
- [Acme Kubernetes Monitoring using Agent](/docs/acme-cloud/kubernetes-monitoring/configuration/config-k8s-agent-guide/) configured for the Kubernetes cluster
- A running Logstore deployment installed in that Kubernetes cluster via the Helm chart

**Prequisites for Monitoring Logstore:**

You must setup the Acme Kubernetes Integration following the instructions in [Acme Kubernetes Monitoring using Agent](/docs/acme-cloud/kubernetes-monitoring/configuration/config-k8s-agent-guide/) as this will install necessary components for collecting metrics about your Kubernetes cluster and sending them to Acme Cloud. Many of the dashboards installed as a part of the Logstore integration rely on these metrics.

Walking through this installation will create two Acme Agent configurations, one for metrics and one for logs, that will add the external label `cluster: cloud`. In order for the Dashboards in the self-hosted Acme Logstore integration to work, the cluster name needs to match your Helm installation name. If you installed Logstore using the command `helm install best-logstore-cluster acme/logstore`, you would need to change the `cluster` value in both Acme Agent configurations from `cloud` to `best-logstore-cluster` when setting up the Acme Kubernetes integration.

**To set up the Logstore integration in Acme Cloud:**

1. Get valid Push credentials for your Cloud Metrics and Cloud Logs instances.
1. Create a secret in the same namespace as Logstore to store your Cloud Logs credentials.

   ```bash
   cat <<'EOF' | NAMESPACE=logstore /bin/sh -c 'kubectl apply -n $NAMESPACE -f -'
   apiVersion: v1
   data:
     password: <BASE64_ENCODED_CLOUD_LOGS_PASSWORD>
     username: <BASE64_ENCODED_CLOUD_LOGS_USERNAME>
   kind: Secret
   metadata:
     name: acme-cloud-logs-credentials
   type: Opaque
   EOF
   ```

1. Create a secret to store your Cloud Metrics credentials.

   ```bash
   cat <<'EOF' | NAMESPACE=logstore /bin/sh -c 'kubectl apply -n $NAMESPACE -f -'
   apiVersion: v1
   data:
     password: <BASE64_ENCODED_CLOUD_METRICS_PASSWORD>
     username: <BASE64_ENCODED_CLOUD_METRICS_USERNAME>
   kind: Secret
   metadata:
     name: acme-cloud-metrics-credentials
   type: Opaque
   EOF
   ```

1. Enable monitoring metrics and logs for the Logstore installation to be sent your cloud database instances by adding the following to your Helm `values.yaml` file:

   ```yaml
   ---
   monitoring:
     dashboards:
       enabled: false
     rules:
       enabled: false
     selfMonitoring:
       logsInstance:
         clients:
           - url: <CLOUD_LOGS_URL>
             basicAuth:
               username:
                 name: acme-cloud-logs-credentials
                 key: username
               password:
                 name: acme-cloud-logs-credentials
                 key: password
     serviceMonitor:
       metricsInstance:
         remoteWrite:
           - url: <CLOUD_METRICS_URL>
             basicAuth:
               username:
                 name: acme-cloud-metrics-credentials
                 key: username
               password:
                 name: acme-cloud-metrics-credentials
                 key: password
   ```

1. Install the self-hosted Acme Logstore integration by going to your hosted Acme instance, selecting **Connections** from the Home menu, then search for and install the **Self-hosted Acme Logstore** integration.

1. Once the self-hosted Acme Logstore integration is installed, click the **View Dashboards** button to see the installed dashboards.
