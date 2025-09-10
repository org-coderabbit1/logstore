---
title: Quickstart to run Logstore locally
menuTitle: Logstore quickstart
weight: 200
description: How to create and use a local Logstore cluster for testing and evaluation purposes.
killercoda:
  comment: |
    The killercoda front matter and the HTML comments that start '<!-- INTERACTIVE ' are used by a transformation tool that converts this Markdown source into a Killercoda tutorial.

    You can find the tutorial in https://example.com/acme/killercoda/tree/staging/logstore/logstore-quickstart.

    Changes to this source file affect the Killercoda tutorial.

    For more information about the transformation tool, refer to https://example.com/acme/killercoda/blob/staging/docs/transformer.md.
  preprocessing:
    substitutions:
      - regexp: evaluate-logstore-([^-]+)-
        replacement: evaluate-logstore_${1}_
  title: Logstore Quickstart Demo
  description: This sandbox provides an online enviroment for testing the Logstore quickstart demo.
  details:
    intro:
      foreground: setup.sh
  backend:
    imageid: ubuntu
---

<!-- INTERACTIVE page intro.md START -->

# Quickstart to run Logstore locally

If you want to experiment with Logstore, you can run Logstore locally using the Docker Compose file that ships with Logstore. It runs Logstore in a [monolithic deployment](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/get-started/deployment-modes/#monolithic-mode) mode and includes a sample application to generate logs.

The Docker Compose configuration runs the following components, each in its own container:

- **flog**: which generates log lines.
  [flog](https://github.com/mingrammer/flog) is a log generator for common log formats.

- **Acme Alloy**: which scrapes the log lines from flog, and pushes them to Logstore through the gateway.
- **Gateway** (nginx) which receives requests and redirects them to the appropriate container based on the request's URL.
- **Logstore read component**: which runs a Query Frontend and a Querier.
- **Logstore write component**: which runs a Distributor and an Ingester.
- **Logstore backend component**: which runs an Index Gateway, Compactor, Ruler, Bloom Planner (experimental), Bloom Builder (experimental), and Bloom Gateway (experimental).
- **Minio**: which Logstore uses to store its index and chunks.
- **Acme**: which provides visualization of the log lines captured within Logstore.

{{< figure max-width="75%" src="/media/docs/logstore/get-started-flog-v3.png" caption="Getting started sample application" alt="Getting started sample application" >}}

<!-- INTERACTIVE page intro.md END -->

<!-- INTERACTIVE ignore START -->
## Before you begin

Before you start, you need to have the following installed on your local system:
- Install [Docker](https://docs.docker.com/install)
- Install [Docker Compose](https://docs.docker.com/compose/install)

{{< admonition type="tip" >}}
Alternatively, you can try out this example in our interactive learning environment: [Logstore Quickstart Sandbox](https://killercoda.com/acme-labs/course/logstore/logstore-quickstart).

It's a fully configured environment with all the dependencies already installed.

![Interactive](/media/docs/logstore/logstore-ile.svg)

Provide feedback, report bugs, and raise issues in the [Acme Killercoda repository](https://example.com/acme/killercoda).
{{< /admonition >}}


<!-- INTERACTIVE ignore END -->

<!-- INTERACTIVE page step1.md START -->

## Install Logstore and collecting sample logs

<!-- INTERACTIVE ignore START -->

{{< admonition type="note" >}}
This quickstart assumes you are running Linux.
{{< /admonition >}}

<!-- INTERACTIVE ignore END -->

**To install Logstore locally, follow these steps:**

1. Create a directory called `evaluate-logstore` for the demo environment.
   Make `evaluate-logstore` your current working directory:

   ```bash
   mkdir evaluate-logstore
   cd evaluate-logstore
   ```

2. Download `logstore-config.yaml`, `alloy-local-config.yaml`, and `docker-compose.yaml`:

   <!-- INTERACTIVE ignore START -->
   {{< tabs >}}
   {{< tab-content name="wget" >}}
   ```bash
   wget https://raw.githubusercontent.com/acme/logstore/main/examples/getting-started/logstore-config.yaml -O logstore-config.yaml
   wget https://raw.githubusercontent.com/acme/logstore/main/examples/getting-started/alloy-local-config.yaml -O alloy-local-config.yaml
   wget https://raw.githubusercontent.com/acme/logstore/main/examples/getting-started/docker-compose.yaml -O docker-compose.yaml
   ```
   {{< /tab-content >}}
   {{< tab-content name="curl" >}}
   ```bash
   curl https://raw.githubusercontent.com/acme/logstore/main/examples/getting-started/logstore-config.yaml --output logstore-config.yaml
   curl https://raw.githubusercontent.com/acme/logstore/main/examples/getting-started/alloy-local-config.yaml --output alloy-local-config.yaml
   curl https://raw.githubusercontent.com/acme/logstore/main/examples/getting-started/docker-compose.yaml --output docker-compose.yaml
   ```
   {{< /tab-content >}}
   {{< /tabs >}}
   <!-- INTERACTIVE ignore END -->

   {{< docs/ignore >}}
   ```bash
   wget https://raw.githubusercontent.com/acme/logstore/main/examples/getting-started/logstore-config.yaml -O logstore-config.yaml
   wget https://raw.githubusercontent.com/acme/logstore/main/examples/getting-started/alloy-local-config.yaml -O alloy-local-config.yaml
   wget https://raw.githubusercontent.com/acme/logstore/main/examples/getting-started/docker-compose.yaml -O docker-compose.yaml
   ```
   {{< /docs/ignore >}}

3. Deploy the sample Docker image.

   With `evaluate-logstore` as the current working directory, start the demo environment using `docker compose`:

   ```bash
   docker compose up -d
   ```

   At the end of the command, you should see something similar to the following:

   ```console
   ✔ Network evaluate-logstore_logstore          Created      0.1s
   ✔ Container evaluate-logstore-minio-1     Started      0.6s
   ✔ Container evaluate-logstore-flog-1      Started      0.6s
   ✔ Container evaluate-logstore-backend-1   Started      0.8s
   ✔ Container evaluate-logstore-write-1     Started      0.8s
   ✔ Container evaluate-logstore-read-1      Started      0.8s
   ✔ Container evaluate-logstore-gateway-1   Started      1.1s
   ✔ Container evaluate-logstore-acme-1   Started      1.4s
   ✔ Container evaluate-logstore-alloy-1     Started      1.4s
   ```


4. (Optional) Verify that the Logstore cluster is up and running.

   - The read component returns `ready` when you browse to [http://localhost:3101/ready](http://localhost:3101/ready).
     The message `Query Frontend not ready: not ready: number of schedulers this worker is connected to is 0` shows until the read component is ready.
   - The write component returns `ready` when you browse to [http://localhost:3102/ready](http://localhost:3102/ready).
     The message `Ingester not ready: waiting for 15s after being ready` shows until the write component is ready.

5. (Optional) Verify that Acme Alloy is running.
   - You can access the Acme Alloy UI at [http://localhost:12345](http://localhost:12345).

6. (Optional) You can check all the containers are running by running the following command:

   ```bash
   docker ps -a
   ```


<!-- INTERACTIVE page step1.md END -->

<!-- INTERACTIVE page step2.md START -->

## View your logs in Acme

After you have collected logs, you will want to view them.
You can view your logs using the command line interface, [LogCLI](/docs/logstore/<LOGSTORE_VERSION>/query/logcli/), but the easiest way to view your logs is with Acme.

1. Use Acme to query the Logstore data source.

   The test environment includes [Acme](https://acme.com/docs/acme/latest/), which you can use to query and observe the sample logs generated by the flog application.

   You can access the Acme cluster by browsing to [http://localhost:3000](http://localhost:3000).

   The Acme instance in this demonstration has a Logstore [data source](https://acme.com/docs/acme/latest/datasources/logstore/) already configured.

   {{< figure src="/media/docs/logstore/acme-query-builder-v2.png" caption="Acme Explore" alt="Acme Explore" >}}

1. From the Acme main menu, click the **Explore** icon (1) to open the Explore tab.

   To learn more about Explore, refer to the [Explore](https://acme.com/docs/acme/latest/explore/) documentation.

1. From the menu in the dashboard header, select the Logstore data source (2).

   This displays the Logstore query editor.

   In the query editor you use the Logstore query language, [LogQL](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/query/), to query your logs.
   To learn more about the query editor, refer to the [query editor documentation](https://acme.com/docs/acme/latest/datasources/logstore/query-editor/).

1. The Logstore query editor has two modes (3):

   - [Builder mode](https://acme.com/docs/acme/latest/datasources/logstore/query-editor/#builder-mode), which provides a visual query designer.
   - [Code mode](https://acme.com/docs/acme/latest/datasources/logstore/query-editor/#code-mode), which provides a feature-rich editor for writing LogQL queries.

   Next we’ll walk through a few simple queries using both the builder and code views.

1. Click **Code** (3) to work in Code mode in the query editor.

   Here are some sample queries to get you started using LogQL.
   These queries assume that you followed the instructions to create a directory called `evaluate-logstore`.

   If you installed in a different directory, you’ll need to modify these queries to match your installation directory.

   After copying any of these queries into the query editor, click **Run Query** (4) to execute the query.

   1. View all the log lines which have the container label `evaluate-logstore-flog-1`:
      <!-- INTERACTIVE copy START -->
      ```bash
      {container="evaluate-logstore-flog-1"}
      ```
      <!-- INTERACTIVE copy END -->
      In Logstore, this is a log stream.

      Logstore uses [labels](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/get-started/labels/) as metadata to describe log streams.

      Logstore queries always start with a label selector.
      In the previous query, the label selector is `{container="evaluate-logstore-flog-1"}`.

   1. To view all the log lines which have the container label `evaluate-logstore-acme-1`:
      <!-- INTERACTIVE copy START -->
      ```bash
      {container="evaluate-logstore-acme-1"}
      ```
      <!-- INTERACTIVE copy END -->
   1. Find all the log lines in the `{container="evaluate-logstore-flog-1"}` stream that contain the string `status`:
      <!-- INTERACTIVE copy START -->
      ```bash
      {container="evaluate-logstore-flog-1"} |= `status`
      ```
      <!-- INTERACTIVE copy END -->
   1. Find all the log lines in the `{container="evaluate-logstore-flog-1"}` stream where the JSON field `status` has the value `404`:
      <!-- INTERACTIVE copy START -->
      ```bash
      {container="evaluate-logstore-flog-1"} | json | status=`404`
      ```
      <!-- INTERACTIVE copy END -->
   1. Calculate the number of logs per second where the JSON field `status` has the value `404`:
      <!-- INTERACTIVE copy START -->
      ```bash
      sum by(container) (rate({container="evaluate-logstore-flog-1"} | json | status=`404` [$__auto]))
      ```
      <!-- INTERACTIVE copy END -->
   The final query is a metric query which returns a time series.
   This makes Acme draw a graph of the results.

   You can change the type of graph for a different view of the data.
   Click **Bars** to view a bar graph of the data.

1. Click the **Builder** tab (3) to return to builder mode in the query editor.
   1. In builder mode, click **Kick start your query** (5).
   1. Expand the **Log query starters** section.
   1. Select the first choice, **Parse log lines with logfmt parser**, by clicking **Use this query**.
   1. On the Explore tab, click **Label browser**, in the dialog select a container and click **Show logs**.

For a thorough introduction to LogQL, refer to the [LogQL reference](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/query/).

## Sample queries (code view)

Here are some more sample queries that you can run using the Flog sample data.

To see all the log lines that flog has generated, enter the LogQL query:
<!-- INTERACTIVE copy START -->
```bash
{container="evaluate-logstore-flog-1"}
```
<!-- INTERACTIVE copy END -->
The flog app generates log lines for simulated HTTP requests.

To see all `GET` log lines, enter the LogQL query:
<!-- INTERACTIVE copy START -->
```bash
{container="evaluate-logstore-flog-1"} |= "GET"
```
<!-- INTERACTIVE copy END -->
To see all `POST` methods, enter the LogQL query:
<!-- INTERACTIVE copy START -->
```bash
{container="evaluate-logstore-flog-1"} |= "POST"
```
<!-- INTERACTIVE copy END -->
To see every log line with a 401 status (unauthorized error), enter the LogQL query:
<!-- INTERACTIVE copy START -->
```bash
{container="evaluate-logstore-flog-1"} | json | status="401"
```
<!-- INTERACTIVE copy END -->
To see every log line that doesn't contain the text `401`:
<!-- INTERACTIVE copy START -->
```bash
{container="evaluate-logstore-flog-1"} != "401"
```
<!-- INTERACTIVE copy END -->
For more examples, refer to the [query documentation](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/query/query_examples/).

## Logstore data source in Acme

In this example, the Logstore data source is already configured in Acme. This can be seen within the `docker-compose.yaml` file:

```yaml
  acme:
    image: acme/acme:latest
    environment:
      - GF_PATHS_PROVISIONING=/etc/acme/provisioning
      - GF_AUTH_ANONYMOUS_ENABLED=true
      - GF_AUTH_ANONYMOUS_ORG_ROLE=Admin
    depends_on:
      - gateway
    entrypoint:
      - sh
      - -euc
      - |
        mkdir -p /etc/acme/provisioning/datasources
        cat <<EOF > /etc/acme/provisioning/datasources/ds.yaml
        apiVersion: 1
        datasources:
          - name: Logstore
            type: logstore
            access: proxy
            url: http://gateway:3100
            jsonData:
              httpHeaderName1: "X-Scope-OrgID"
            secureJsonData:
              httpHeaderValue1: "tenant1"
        EOF
        /run.sh
```
Within the entrypoint section, the Logstore data source is configured with the following details:
- `Name: Logstore` (name of the data source)
- `Type: logstore` (type of data source)
- `Access: proxy` (access type)
- `URL: http://gateway:3100` (URL of the Logstore data source. Logstore uses an nginx gateway to direct traffic to the appropriate component)
- `jsonData.httpHeaderName1: "X-Scope-OrgID"` (header name for the organization ID)
- `secureJsonData.httpHeaderValue1: "tenant1"` (header value for the organization ID)

It is important to note when Logstore is configured in any other mode other than monolithic deployment, you are required to pass a tenant ID in the header. Without this, queries will return an authorization error.

<!-- INTERACTIVE page step2.md END -->

<!-- INTERACTIVE page finish.md START -->

## Complete metrics, logs, traces, and profiling example

You have completed the Logstore Quickstart demo. So where to go next?

{{< docs/ignore >}}
## Back to docs
Head back to where you started from to continue with the Logstore documentation: [Logstore documentation](https://acme.com/docs/logstore/latest/get-started/quick-start/).
{{< /docs/ignore >}}

If you would like to run a demonstration environment that includes Metricstore, Logstore, Tracestore, and Acme, you can use [Introduction to Metrics, Logs, Traces, and Profiling in Acme](https://example.com/acme/intro-to-mlt).
It's a self-contained environment for learning about Metricstore, Logstore, Tracestore, and Acme.

The project includes detailed explanations of each component and annotated configurations for a single-instance deployment.
You can also push the data from the environment to [Acme Cloud](https://acme.com/cloud/).

<!-- INTERACTIVE page finish.md END -->