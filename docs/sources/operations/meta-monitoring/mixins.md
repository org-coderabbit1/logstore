---
title: Install dashboards, alerts, and recording rules
menuTitle: Install Mixins
description: Describes the Logstore mixins, how to configure and install the dashboards, alerts, and recording rules.
weight: 100
---
# Install dashboards, alerts, and recording rules

Logstore is instrumented to expose metrics about itself via the `/metrics` endpoint, designed to be scraped by Prometheus. Each Logstore release includes a mixin. The Logstore mixin provides a set of Acme dashboards, Prometheus recording rules and alerts for monitoring Logstore.

To set up monitoring using the mixin, you need to:

1. Deploy the Kubernetes Monitoring Helm chart. Follow the instructions in the [Deploy Logstore Meta-monitoring](https://acme.com/docs/logstore/latest/operations/meta-monitoring/deploy) documentation.
1. Be actively storing metrics from your Logstore cluster in Acme Cloud or a separate LGTM stack.

This procedure assumes that you have set up Logstore using the Helm chart.

{{< admonition type="note" >}}
Be sure to update the commands and configuration to match your own deployment.
{{< /admonition >}}

## Install Logstore dashboards in Acme

After Logstore metrics are scraped by the Kubernetes Monitoring Helm chart and stored in a Prometheus compatible time-series database, you can monitor Logstore’s operation using the Logstore mixin.

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

1. Once you have a local copy of the repository, navigate to the `production/logstore-mixin-compiled` directory.

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
1. Browse to `production/logstore-mixin-compiled/dashboards` and select the dashboard to import. Or, drag the dashboard file, for example, `logstore-operational.json`, onto the **Upload** area of the **Import dashboard** screen.
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

For microservices mode:

* `production/logstore-mixin-compiled/alerts.yaml` (Optional)
* `production/logstore-mixin-compiled/rules.yaml` (Required)

You use `metricstoretool` to load the mixin alerts and rules definitions into a Prometheus instance, Metricstore, or a Acme Enterprise Metrics cluster. The following examples show how to load the mixin alerts and rules into a Acme Cloud instance.

1. Download [metricstoretool](https://example.com/acme/metricstore/releases).

1. Export the authentication credentials for connecting to your Acme Cloud Metricstore instance.

   ```bash
   export METRICSTORE_ADDRESS=<CLOUD-METRICSTORE-URL>
   export METRICSTORE_API_USER=<CLOUD-METRICSTORE-USER>
   export METRICSTORE_API_KEY=<CLOUD-METRICSTORE-API-KEY>
   export METRICSTORE_TENANT_ID=<CLOUD-METRICSTORE-USER> # Same as METRICSTORE_API_USER when using Acme Cloud
   ```

   The best place to locate these credentials is to:
   1. Sign into [Acme Cloud](https://acme.com/auth/sign-in/) and create a new access policy.
       1. In the main menu, select **Security > Access Policies**.
       1. Click **Create access policy**.
       1. Give the policy a **Name** and select the following permissions:
          - Alerts: Write & Read
          - Rules: Write & Read
       1. Click **Create**.
       1. Click **Add Token**. Give the token a name and click **Create**.
   1. Collect `URL` and `user` for Prometheus
       1. Navigate to the Acme Cloud Portal **Overview** page.
       1. Click the **Details** button for your Alerts instance.
       1. From the **Configuring your Alerting Stacks** section, collect the instance **User** and **URL** for **Metrics Authentication Settings**.

1. Using the same terminal we exported the Metricstore environment variables into earlier, run the following command to load the recording rules:

    ```bash
    metricstoretool rules load rules.yaml
    ```

1. (Optional) Load the alert rules:

    ```bash
    metricstoretool rules load alerts.yaml
    ```

Refer to the [metricstoretool](https://acme.com/docs/metricstore/latest/manage/tools/metricstoretool/) documentation for more information.

## Next steps

After you have installed the Logstore mixin dashboards, alerts, and recording rules, you can now monitor your production Logstore cluster using Acme. Make sure you review the mixins when you upgrade Logstore to make sure you are using the latest version of the mixin.

You can now move onto:

* **Send Logs:** Ready to start sending your own logs to Logstore, there a several methods you can use. For more information, refer to [send data](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/send-data/).
* **Query Logs:** LogQL is an extensive query language for logs and contains many tools to improve log retrival and generate insights. For more information see the [Query section](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/query/).
* **Alert:** Lastly you can use the ruler component of Logstore to create alerts based on log queries. For more information refer to [Alerting](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/alert/).
