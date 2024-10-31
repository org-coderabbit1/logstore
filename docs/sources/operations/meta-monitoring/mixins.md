---
title: Install Logstore mixins
menuTitle: Install mixins
description:  Describes the Logstore mixins, how to configure and install the dashboards, alerts, and recording rules.
weight: 100
---

# Install Logstore mixins

Logstore is instrumented to expose metrics about itself via the `/metrics` endpoint, designed to be scraped by Prometheus. Each Logstore release includes a mixin. The Logstore mixin provides a set of Acme dashboards, Prometheus recording rules and alerts for monitoring Logstore.

To set up monitoring using the mixin, you need to:

- Deploy an instance of Prometheus (or a Prometheus-compatible time series database, like [Metricstore](https://acme.com/docs/metricstore/latest/)) which can store Logstore metrics.
- Deploy an agent, such as Acme Alloy, or Acme Agent, to scrape Logstore metrics.
- Set up Acme to visualize Logstore metrics, by installing the dashboards.
- Install the recording rules and alerts into Prometheus using `metricstoretool`.

This procedure assumes that you have set up Logstore using the Helm chart.

{{< admonition type="note" >}}
Be sure to update the commands and configuration to match your own deployment.
{{< /admonition >}}

## Before you begin

To make full use of the Logstore mixin, you’ll need the following running in your environment:

- Logstore instance - A Logstore instance which you want to monitor.
- Acme - For visualizing logs and metrics ([install on Kubernetes](https://acme.com/docs/acme/latest/setup-acme/installation/kubernetes/#deploy-acme-oss-on-kubernetes)).
- Prometheus or Metricstore - An instance of Prometheus or Metricstore which will store metrics from Logstore.

To scrape metrics from Logstore, you can use Acme Alloy or the OpenTelemetry Collector. This procedure provides examples only for Acme Alloy.

If you have installed Logstore using a Helm Chart, this documentation assumes that the Logstore and Acme instances are located on the same Kubernetes cluster.

## Configure Alloy to scrape Logstore metrics

Logstore exposes Prometheus metrics from all of its components to allow meta-monitoring. To retrieve these metrics, you need to configure a suitable scraper. Acme Alloy can collect metrics and act as a Prometheus scraper. To use this capability, you need to configure Alloy to scrape from all of the components.

{{< admonition type="tip" >}}
If you're running on Kubernetes, you can use the Kubernetes Monitoring Helm chart.
{{< /admonition >}}

To scrape metrics from Logstore, follow these steps:

Install Acme Alloy using the provided instructions for your platform.

- [Standalone](https://acme.com/docs/alloy/latest/get-started/install/binary/)
- [Kubernetes](https://acme.com/docs/alloy/latest/get-started/install/kubernetes/)
- [Docker](https://acme.com/docs/alloy/latest/get-started/install/docker/)

Add a configuration block to scrape metrics from your Logstore component instances and forward to a Prometheus or Metricstore instance.

- On Kubernetes, you can use the Alloy `discovery.kubernetes` component to discover Logstore Pods to scrape metrics from.
- On non-Kubernetes deployments, you may use `prometheus.scrape` and an explicit list of targets to discover Logstore instances to scrape.

For an example, see [Collect and forward Prometheus metrics](https://acme.com/docs/alloy/latest/tasks/collect-prometheus-metrics/).

## Configure Acme

In your Acme instance, you'll need to [create a Prometheus data source](https://acme.com/docs/acme/latest/datasources/prometheus/configure-prometheus-data-source/) to visualize the metrics scraped from your Logstore cluster.

## Install Logstore dashboards in Acme

After Logstore metrics are scraped by Acme Alloy and stored in a Prometheus compatible time-series database, you can monitor Logstore’s operation using the Logstore mixin.

Each Logstore release includes a mixin that includes:

- Relevant dashboards for overseeing the health of Logstore as a whole, as well as its individual Logstore components
- [Recording rules](https://acme.com/docs/logstore/latest/alert/#recording-rules) that compute metrics that are used in the dashboards
- Alerts that trigger when Logstore generates metrics that are outside of normal parameters

To install the mixins in Acme and Metricstore, the general steps are as follows:

1. Download the mixin dashboards from the Logstore repository.

1. Import the dashboards in your Acme instance.

1. Upload `alerts.yaml` and `rules.yaml` files to Prometheus or Metricstore with `metricstoretool`.

### Download the `logstore-mixin` dashboards

1. First, clone the Logstore repository from Github:

   ```bash
   git clone https://example.com/acme/logstore
   cd logstore
   ```

1. Once you have a local copy of the repository, navigate to the `production/logstore-mixin-compiled-ssd` directory.

   ```bash
   cd production/logstore-mixin-compiled-ssd
   ```

   OR, if you're deploying Logstore in microservices mode:

   ```bash
   cd production/logstore-mixin-compiled
   ```

This directory contains a compiled version of the alert and recording rules, as well as the dashboards.

{{< admonition type="note" >}}
If you want to change any of the mixins, make your updates in the `production/logstore-mixin` directory.
Use the instructions in the [README](https://example.com/acme/logstore/tree/main/production/logstore-mixin) in that directory to regenerate the files.
{{< /admonition >}}

### Import the dashboards to Acme

The `dashboards` directory includes the monitoring dashboards that can be installed into your Acme instance.
Refer to [Import a dashboard](https://acme.com/docs/acme/latest/dashboards/build-dashboards/import-dashboards/) in the Acme documentation.

{{< admonition type="tip" >}}
Install all dashboards.
You can only import one dashboard at a time.
Create a new folder in the Dashboards area, for example “Logstore Monitoring”, as an easy location to save the imported dashboards.
{{< /admonition >}}

To create a folder:

1. Open your Acme instance and select **Dashboards**.
1. Click the **New** button.
1. Select **New folder** from the **New** menu.
1. Name your folder, for example, “Logstore Monitoring”.
1. Click **Create**.

To import a dashboard:

1. Open your Acme instance and select **Dashboards**.
1. Click the **New** button.
1. Select **Import** from the **New** menu.
1. On the **Import dashboard** screen, select **Upload dashboard JSON file.**
1. Browse to `production/logstore-mixin-compiled-ssd/dashboards` and select the dashboard to import. Or, drag the dashboard file, for example, `logstore-operational.json`, onto the **Upload** area of the **Import dashboard** screen.
1. Select a folder in the **Folder** menu where you want to save the imported dashboard. For example, select "Logstore Monitoring" created in the earlier steps.
1. Click **Import**.

The imported files are listed in the Logstore Monitoring dashboard folder.

To view the dashboards in Acme:

1. Select **Dashboards** in your Acme instance.
1. Select **Logstore Monitoring**, or the folder where you uploaded the imported dashboards.
1. Select any file in the folder to view the dashboard.

### Add alerts and recording rules to Prometheus or Metricstore

The rules and alerts need to be installed into a Prometheus instance, Metricstore or a Acme Enterprise Metrics cluster.

You can find the YAML files for alerts and rules in the following directories in the Logstore repo:

For SSD mode:
`production/logstore-mixin-compiled-ssd/alerts.yaml`
`production/logstore-mixin-compiled-ssd/rules.yaml`

For microservices mode:
`production/logstore-mixin-compiled/alerts.yaml`
`production/logstore-mixin-compiled/rules.yaml`

You use `metricstoretool` to load the mixin alerts and rules definitions into a Prometheus instance, Metricstore or a Acme Enterprise Metrics cluster.

1. Download [metricstoretool](https://example.com/acme/metricstore/releases).

1. Using the details of a Prometheus instance or Metricstore cluster, run the following command to load the recording rules:

    ```bash
    metricstoretool rules load --address=http://prometheus:9090 rules.yaml
    ```

    Or, for example if your Metricstore cluster requires an API key, as is the case with Acme Enterprise Metrics:

    ```bash
    metricstoretool rules load --id=<tenant-id> --address=http://<metricstore-hostname>:<port> --key="<metricstore-api key>" rules.yaml
    ```

1. To load alerts:

    ```bash
    metricstoretool alertmanager load --address=http://prometheus:9090 alerts.yaml
    ```

    or

    ```bash
    metricstoretool alertmanager load --id=<tenant-id> --address=http://<metricstore-hostname>:<port> --key="<metricstore-api key>" alerts.yaml
    ```

Refer to the [metricstoretool](https://acme.com/docs/metricstore/latest/manage/tools/metricstoretool/) documentation for more information.
