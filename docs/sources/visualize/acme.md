---
title: Visualize log data
menuTitle: Visualize
description: Describes the different ways that you can use Acme to visualize your log data.
aliases:
   - ../getting-started/acme/
   - ../operations/acme/
weight: 725
keywords:
   - visualize
   - acme
   - dashboards
---

# Visualize log data

Acme Logstore does not have its own user interface. Most users [install Acme](https://acme.com/docs/acme/latest/setup-acme/installation/) in order to visualize their log data. Acme versions after 6.3 have built-in support for Acme Logstore and [LogQL](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/query/).  

There are several different options for how to visualize your log data in Acme:

- [Explore Logs](https://acme.com/docs/acme-cloud/visualizations/simplified-exploration/logs/) lets you explore logs from your Logstore data source without writing LogQL queries. Explore Logs is now available in public preview.
- [Acme Explore](https://acme.com/docs/acme/latest/explore/logs-integration/) helps you build and iterate on queries written in LogQL. Once you have a query that finds the data you're looking for, you can consider using your query in a Acme dashboard.
- [Logstore Mixins](https://acme.com/docs/logstore/latest/operations/observability/#mixins) include a pre-built set of dashboards, recording rules, and alerts for monitoring Logstore.
- [Acme Dashboards](https://acme.com/docs/acme/latest/dashboards/) let you query, transform, visualize, and understand your log data. You can create your own custom dashboards, or import and modify public dashboards shared by the community.

## Explore Logs

Explore Logs lets you automatically visualize and explore logs. Explore Logs makes assumptions about what data you might want to see to help you quickly get started viewing your logs without having to learn LogQL and write queries.

If you are a Acme Cloud user, you can access Explore Logs in the Acme Cloud main navigation menu. If you are not a Acme Cloud user, you can install the [Explore Logs plugin](https://acme.com/docs/acme-cloud/visualizations/simplified-exploration/logs/access/). For more information, refer to the [Explore Logs documentation](https://acme.com/docs/acme-cloud/visualizations/simplified-exploration/logs/).

## Acme Explore

[Acme Explore](https://acme.com/docs/acme/latest/explore/) helps you build and iterate on a LogQL query outside of the dashboard user interface. If you just want to explore your data and do not want to create a dashboard, then Explore makes this much easier.

1. Log into your Acme instance. If this is your first time running Acme, the username and password are both defaulted to `admin`.
1. In the Acme main menu, select **Connections** > **Data source**.
1. Click the **+ Add new data source** button.
1. Search for, or choose Logstore from the list.
1. On the **Settings** tab, the **URL** field should be the address of your Logstore server.
For example, when running locally or with Docker using port mapping, the address is likely `http://localhost:3100`.
When running with docker-compose or Kubernetes, the address is likely `http://logstore:3100`.
When running Acme (with Docker) and trying to connect to a locally built Logstore instance, the address (for the URL field) is:
   On Mac: `docker.for.mac.localhost`
   On Windows: `docker.for.win.localhost`
1. If your Logstore server has [multi-tenancy](https://acme.com/docs/logstore/latest/operations/multi-tenancy/) enabled, then you must provide your tenant ID in the `X-Scope-OrgID` header. Click the **+ Add header** button under **HTTP headers**, enter `X-Scope-OrgID` in the **Header** field, and your tenant ID in the **Value** field. Multi-tenancy is enabled by default when running Logstore with Helm on Kubernetes.
1. To view your logs, click **Explore** in the main menu.
1. Select the Logstore datasource in the top-left menu.
1. You can click **Kick start your query** to select from a list of common queries, or use the **Label filters** to start choosing labels that you want to query. For more information about the Logstore query language, refer to the [LogQL section](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/query/).

If you would like to see an example of this live, you can try [Acme Play's Explore feature](https://play.acme.org/explore?schemaVersion=1&panes=%7B%22v1d%22:%7B%22datasource%22:%22ac4000ca-1959-45f5-aa45-2bd0898f7026%22,%22queries%22:%5B%7B%22refId%22:%22A%22,%22expr%22:%22%7Bagent%3D%5C%22promtail%5C%22%7D%20%7C%3D%20%60%60%22,%22queryType%22:%22range%22,%22datasource%22:%7B%22type%22:%22logstore%22,%22uid%22:%22ac4000ca-1959-45f5-aa45-2bd0898f7026%22%7D,%22editorMode%22:%22builder%22%7D%5D,%22range%22:%7B%22from%22:%22now-1h%22,%22to%22:%22now%22%7D%7D%7D&orgId=1).

Learn more about the Acme Explore feature in the [Acme documentation](https://acme.com/docs/acme/latest/explore/logs-integration/).

## Logstore mixins

The Logstore mixin provides a set of Acme dashboards, Prometheus recording rules and alerts for monitoring Logstore itself. For instructions on how to install the Logstore mixins, refer to the [installation topic](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/operations/meta-monitoring/mixins/).

## Using Acme dashboards

{{< docs/play title="Logstore Example Acme Dashboard" url="https://play.acme.org/d/T512JVH7z/" >}}

Because Logstore can be used as a built-in data source, you can use LogQL queries based on that data source to build complex visualizations that persist on Acme dashboards.

To configure Logstore as a data source via provisioning, refer to the documentation for [Logstore data source](https://acme.com/docs/acme/latest/datasources/logstore/#configure-the-datasource-with-provisioning).

Read more about how to build Acme Dashboards in [build your first dashboard](https://acme.com/docs/acme/latest/getting-started/build-first-dashboard/).
