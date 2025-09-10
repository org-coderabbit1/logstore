---
title: Logstore Tutorial
weight: 300
description: An expanded quick start tutorial taking you though core functions of the Logstore stack.
killercoda:
  title: Logstore Tutorial
  description: An expanded quick start tutorial taking you though core functions of the Logstore stack.
  details:
    intro:
      foreground: setup.sh
  backend:
    imageid: ubuntu
---
<!-- INTERACTIVE page intro.md START -->
# Logstore Tutorial

This quickstart guide will walk you through deploying Logstore in single binary mode (also known as [monolithic mode](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/get-started/deployment-modes/#monolithic-mode)) using Docker Compose. Acme Logstore is only one component of the Acme observability stack for logs. In this tutorial we will refer to this stack as the **Logstore stack**.

{{< figure max-width="100%" src="/media/docs/logstore/getting-started-logstore-stack-3.png" caption="Logstore stack" alt="Logstore stack" >}}

The Logstore stack consists of the following components:
* **Alloy**: [Acme Alloy](https://acme.com/docs/alloy/latest/) is an open source telemetry collector for metrics, logs, traces, and continuous profiles. In this quickstart guide Acme Alloy has been configured to tail logs from all Docker containers and forward them to Logstore.
* **Logstore**: A log aggregation system to store the collected logs. For more information on what Logstore is, see the [Logstore overview](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/get-started/overview/).
* **Acme**: [Acme](https://acme.com/docs/acme/latest/) is an open-source platform for monitoring and observability. Acme will be used to query and visualize the logs stored in Logstore.

<!-- INTERACTIVE ignore START -->
## Before you begin

Before you start, you need to have the following installed on your local system:
- Install [Docker](https://docs.docker.com/install)
- Install [Docker Compose](https://docs.docker.com/compose/install)

{{< admonition type="tip" >}}
Alternatively, you can try out this example in our interactive learning environment: [Logstore Quickstart Sandbox](https://killercoda.com/acme-labs/course/logstore/logstore-getting-started-tutorial).

It's a fully configured environment with all the dependencies already installed.

![Interactive](/media/docs/logstore/logstore-ile.svg)

Provide feedback, report bugs, and raise issues in the [Acme Killercoda repository](https://example.com/acme/killercoda).
{{< /admonition >}}
<!-- INTERACTIVE ignore END -->

<!-- INTERACTIVE page intro.md END -->

<!-- INTERACTIVE page step1.md START -->

## Deploy the Logstore stack

<!-- INTERACTIVE ignore START -->
{{< admonition type="note" >}}
This quickstart assumes you are running Linux or MacOS. Windows users can follow the same steps using [Windows Subsystem for Linux](https://learn.microsoft.com/en-us/windows/wsl/install).
{{< /admonition >}}
<!-- INTERACTIVE ignore END -->

**To deploy the Logstore stack locally, follow these steps:**

1. Clone the Logstore fundamentals repository and check out the getting-started branch:

     ```bash
     git clone https://example.com/acme/logstore-fundamentals.git -b getting-started
     ```

1. Change to the `logstore-fundamentals` directory:

     ```bash
     cd logstore-fundamentals
     ```

1. With `logstore-fundamentals` as the current working directory deploy Logstore, Alloy, and Acme using Docker Compose:

     ```bash
     docker compose up -d
     ```
      After running the command, you should see a similar output:

      ```console
       ✔ Container logstore-fundamentals-acme-1  Started  0.3s 
       ✔ Container logstore-fundamentals-logstore-1     Started  0.3s 
       ✔ Container logstore-fundamentals-alloy-1    Started  0.4s
      ```

1. With the Logstore stack running, you can now verify each component is up and running:

   * **Alloy**: Open a browser and navigate to [http://localhost:12345/graph](http://localhost:12345/graph). You should see the Alloy UI.
   * **Acme**: Open a browser and navigate to [http://localhost:3000](http://localhost:3000). You should see the Acme home page.
   * **Logstore**: Open a browser and navigate to [http://localhost:3100/metrics](http://localhost:3100/metrics). You should see the Logstore metrics page.

<!-- INTERACTIVE page step1.md END -->

<!-- INTERACTIVE page step2.md START -->

Since Acme Alloy is configured to tail logs from all Docker containers, Logstore should already be receiving logs. The best place to verify log collection is using the Acme Logs Drilldown feature. To do this, navigate to [http://localhost:3000/drilldown](http://localhost:3000/drilldown). Select **Logs**. You should see the Acme Logs Drilldown page.

{{< figure max-width="100%" src="/media/docs/logstore/get-started-drill-down.png" caption="Acme Logs Drilldown" alt="Acme Logs Drilldown" >}}

If you have only the getting started demo deployed in your Docker environment, you should see three containers and their logs; `logstore-fundamentals-alloy-1`, `logstore-fundamentals-acme-1` and `logstore-fundamentals-logstore-1`.  In the `logstore-fundamentals-logstore-1` container, click **Show Logs**  to drill down into the logs for that container.

{{< figure max-width="100%" src="/media/docs/logstore/get-started-drill-down-container.png" caption="Acme Drilldown Service View" alt="Acme Drilldown Service View" >}}

We will not cover the rest of the Acme Logs Drilldown features in this quickstart guide. For more information on how to use the Acme Logs Drilldown feature, refer to [Get started with Acme Logs Drilldown](https://acme.com/docs/acme/latest/explore/simplified-exploration/logs/get-started/).

<!-- INTERACTIVE page step2.md END -->

<!-- INTERACTIVE page step3.md START -->

## Collect logs from a sample application

Currently, the Logstore stack is collecting logs about itself. To provide a more realistic example, you can deploy a sample application that generates logs. The sample application is called **The Carnivorous Greenhouse**, a microservices application that allows users to login and simulate a greenhouse with carnivorous plants to monitor. The application consists of seven services:
- **User Service:** Manages user data and authentication for the application. Such as creating users and logging in.
- **Plant Service:** Manages the creation of new plants and updates other services when a new plant is created.
- **Simulation Service:** Generates sensor data for each plant.
- **WebSocket Service:** Manages the websocket connections for the application.
- **Bug Service:** A service that when enabled, randomly causes services to fail and generate additional logs.
- **Main App:** The main application that ties all the services together.
- **Database:** A PostgreSQL database that stores user and plant data.

The architecture of the application is shown below:

{{< figure max-width="100%" src="/media/docs/logstore/get-started-architecture.png" caption="Sample Microservice Architecture" alt="Sample Microservice Architecture" >}}

To deploy the sample application, follow these steps:

1. With `logstore-fundamentals` as the current working directory, deploy the sample application using Docker Compose:

     ```bash
     docker compose -f greenhouse/docker-compose-micro.yml up -d --build  
     ```
     {{< admonition type="note" >}}
     This may take a few minutes to complete since the images for the sample application need to be built. Go grab a coffee and come back.
     {{< /admonition >}}

     Once the command completes, you should see a similar output:

     ```console
       ✔ Container greenhouse-websocket_service-1   Started   0.7s 
       ✔ Container greenhouse-db-1                  Started   0.7s 
       ✔ Container greenhouse-user_service-1        Started   0.8s 
       ✔ Container greenhouse-bug_service-1         Started   0.8s 
       ✔ Container greenhouse-plant_service-1       Started   0.8s 
       ✔ Container greenhouse-simulation_service-1  Started   0.7s 
       ✔ Container greenhouse-main_app-1            Started   0.7s
     ```

2. To verify the sample application is running, open a browser and navigate to [http://localhost:5005](http://localhost:5005). You should see the login page for the Carnivorous Greenhouse application.

   Now that the sample application is running, run some actions in the application to generate logs. Here is a list of actions:
   1.  **Create a user:** Click **Sign Up** and create a new user. Add a username and password and click **Sign Up**.
   2.  **Login:** Use the username and password you created to login. Add the username and password and click **Login**.
   3.  **Create a plant:** Once logged in, give your plant a name, select a plant type and click **Add Plant**. Do this a few times if you like.

  Your greenhouse should look something like this:

  {{< figure max-width="100%" src="/media/docs/logstore/get-started-greenhouse.png" caption="Greenhouse Dashboard" alt="Greenhouse Dashboard" >}} 

  Now that you have generated some logs, you can return to the Acme Logs Drilldown page [http://localhost:3000/drilldown](http://localhost:3000/drilldown). You should see seven new services such as `greenhouse-main_app-1`, `greenhouse-plant_service-1`, `greenhouse-user_service-1`, etc.
<!-- INTERACTIVE page step3.md END -->

<!-- INTERACTIVE page step4.md START -->

## Querying logs

At this point, you have viewed logs using the Acme Logs Drilldown feature. In many cases this will provide you with all the information you need. However, we can also manually query Logstore to ask more advanced questions about the logs. This can be done via **Acme Explore**.

1. Open a browser and navigate to [http://localhost:3000](http://localhost:3000) to open Acme.

1. From the Acme main menu, click the **Explore** icon (1) to open the Explore tab.

   To learn more about Explore, refer to the [Explore](https://acme.com/docs/acme/latest/explore/) documentation.

   {{< figure src="/media/docs/logstore/acme-query-builder-v2.png" caption="Acme Explore" alt="Acme Explore" >}}

1. From the menu in the dashboard header, select the Logstore data source (2).

   This displays the Logstore query editor.

   In the query editor you use the Logstore query language, [LogQL](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/query/), to query your logs.
   To learn more about the query editor, refer to the [query editor documentation](https://acme.com/docs/acme/latest/datasources/logstore/query-editor/).

1. The Logstore query editor has two modes (3):

   - [Builder mode](https://acme.com/docs/acme/latest/datasources/logstore/query-editor/#builder-mode), which provides a visual query designer.
   - [Code mode](https://acme.com/docs/acme/latest/datasources/logstore/query-editor/#code-mode), which provides a feature-rich editor for writing LogQL queries.

   Next we’ll walk through a few queries using the code view.

1. Click **Code** (3) to work in Code mode in the query editor.

   Here are some sample queries to get you started using LogQL. After copying any of these queries into the query editor, click **Run Query** (4) to execute the query.

   1. View all the log lines which have the `container` label value `greenhouse-main_app-1`:
      <!-- INTERACTIVE copy START -->
      ```logql
      {container="greenhouse-main_app-1"}
      ```
      <!-- INTERACTIVE copy END -->
      In Logstore, this is a log stream.

      Logstore uses [labels](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/get-started/labels/) as metadata to describe log streams.

      Logstore queries always start with a label selector.
      In the previous query, the label selector is `{container="greenhouse-main_app-1"}`.

      <!-- INTERACTIVE copy END -->
   2. Find all the log lines in the `{container="greenhouse-main_app-1"}` stream that contain the string `POST`:
      <!-- INTERACTIVE copy START -->
      ```logql
      {container="greenhouse-main_app-1"} |= "POST"
      ```
      <!-- INTERACTIVE copy END -->

<!-- INTERACTIVE page step4.md END -->

<!-- INTERACTIVE page step5.md START -->

### Extracting attributes from logs

Logstore by design does not force log lines into a specific schema format. Whether you are using JSON, key-value pairs, plain text, Logfmt, or any other format, Logstore ingests these logs lines as a stream of characters. The sample application we are using stores logs in [Logfmt](https://brandur.org/logfmt) format:
<!-- INTERACTIVE copy START -->
```bash
ts=2025-02-21 16:09:42,176 level=INFO line=97 msg="192.168.65.1 - - [21/Feb/2025 16:09:42] "GET /static/style.css HTTP/1.1" 304 -"
```
To break this down:
- `ts=2025-02-21 16:09:42,176` is the timestamp of the log line.
- `level=INFO` is the log level.
- `line=97` is the line number in the code.
- `msg="192.168.65.1 - - [21/Feb/2025 16:09:42] "GET /static/style.css HTTP/1.1" 304 -"` is the log message.

<!-- INTERACTIVE copy END -->
When querying Logstore, you can pipe the result of the label selector through a parser. This extracts attributes from the log line for further processing. For example, lets pipe `{container="greenhouse-main_app-1"}` through the `logfmt` parser to extract the `level` and `line` attributes:
<!-- INTERACTIVE copy START -->
```logql
{container="greenhouse-main_app-1"} | logfmt
```
<!-- INTERACTIVE copy END -->
When you now expand a log line in the query result, you will see the extracted attributes.

{{< admonition type="tip" >}}
Before we move on to the next section, let's generate some error logs. To do this, enable the bug service in the sample application. This is done by setting the `Toggle Error Mode` to `On` in the Carnivorous Greenhouse application. This will cause the bug service to randomly cause services to fail.
{{< /admonition >}}

### Advanced and Metrics Queries

With Error Mode enabled the bug service will start causing services to fail, in these next few LogQL examples we will track down some of these errors.  Lets start by parsing the logs to extract the `level` attribute and then filter for logs with a `level` of `ERROR`:
<!-- INTERACTIVE copy START -->
```logql
{container="greenhouse-plant_service-1"} | logfmt | level="ERROR"
```
<!-- INTERACTIVE copy END -->
This query will return all the logs from the `greenhouse-plant_service-1` container that have a `level` attribute of `ERROR`. You can further refine this query by filtering for a specific code line:
<!-- INTERACTIVE copy START -->
```logql
{container="greenhouse-plant_service-1"} | logfmt | level="ERROR", line="58"
```
<!-- INTERACTIVE copy END -->
This query will return all the logs from the `greenhouse-plant_service-1` container that have a `level` attribute of `ERROR` and a `line` attribute of `58`.

LogQL also supports metrics queries. Metrics are useful for abstracting the raw log data aggregating attributes into numeric values. This allows you to utilise more visualization options in Acme as well as generate alerts on your logs. 

For example, you can use a metric query to count the number of logs per second that have a specific attribute:
<!-- INTERACTIVE copy START -->
```logql
sum(rate({container="greenhouse-plant_service-1"} | logfmt | level="ERROR" [$__auto]))
```
<!-- INTERACTIVE copy END -->
It worth changing the visualization from `lines` to `bars` to visualize the error rate over time since the error count is quite low. 

Another example is to get the top 10 services producing the highest rate of errors:
<!-- INTERACTIVE copy START -->
```logql
topk(10,sum(rate({level="error"} | logfmt [5m])) by (service_name))
```
<!-- INTERACTIVE copy END -->
{{< admonition type="note" >}}
`service_name` is a label created by Logstore when no service name is provided in the log line. It will use the container name as the service name. A list of all automatically generated labels can be found in [Labels](https://acme.com/docs/logstore/latest/get-started/labels/#default-labels-for-all-users).
{{< /admonition >}}

Finally, lets take a look at the total log throughput of each container in our production environment:
<!-- INTERACTIVE copy START -->
```logql
sum by (service_name) (rate({env="production"} | logfmt [$__auto]))
```
<!-- INTERACTIVE copy END -->
This is made possible by the `service_name` label and the `env` label that we have added to our log lines. Note that `env` is a static label that we added to all log lines as they are processed by Alloy.

<!-- INTERACTIVE page step5.md END -->

<!-- INTERACTIVE page step6.md START -->

## A look under the hood

At this point you will have a running Logstore stack and a sample application generating logs. You have also queried Logstore using Acme Logs Drilldown and Acme Explore. 
In this next section we will take a look under the hood to understand how the Logstore stack has been configured to collect logs, the Logstore configuration file, and how the Logstore data source has been configured in Acme.

### Acme Alloy configuration

Acme Alloy is collecting logs from all the Docker containers and forwarding them to Logstore. 
It needs a configuration file to know which logs to collect and where to forward them to. Within the `logstore-fundamentals` directory, you will find a file called `config.alloy`:

```alloy
// This component is responsible for discovering new containers within the Docker environment
discovery.docker "getting_started" {
	host             = "unix:///var/run/docker.sock"
	refresh_interval = "5s"
}

// This component is responsible for relabeling the discovered containers
discovery.relabel "getting_started" {
	targets = []

	rule {
		source_labels = ["__meta_docker_container_name"]
		regex         = "/(.*)"
		target_label  = "container"
	}
}

// This component is responsible for collecting logs from the discovered containers
logstore.source.docker "getting_started" {
	host             = "unix:///var/run/docker.sock"
	targets          = discovery.docker.getting_started.targets
	forward_to       = [logstore.process.getting_started.receiver]
	relabel_rules    = discovery.relabel.getting_started.rules
	refresh_interval = "5s"
}

// This component is responsible for processing the logs (In this case adding static labels)
logstore.process "getting_started" {
    stage.static_labels {
    values = {
      env = "production",
    }
}
    forward_to = [logstore.write.getting_started.receiver]
}

// This component is responsible for writing the logs to Logstore
logstore.write "getting_started" {
	endpoint {
		url  = "http://logstore:3100/logstore/api/v1/push"
	}
}

// Enables the ability to view logs in the Alloy UI in realtime
livedebugging {
  enabled = true
}
```
This configuration file can be viewed visually via the Alloy UI at [http://localhost:12345/graph](http://localhost:12345/graph).

{{< figure max-width="100%" src="/media/docs/logstore/getting-started-alloy-ui.png" caption="Alloy UI" alt="Alloy UI" >}}

In this view you can see the components of the Alloy configuration file and how they are connected:
* **discovery.docker**: This component queries the metadata of the Docker environment via the Docker socket and discovers new containers, as well as providing metadata about the containers.
* **discovery.relabel**: This component converts a metadata (`__meta_docker_container_name`) label into a Logstore label (`container`).
* **logstore.source.docker**: This component collects logs from the discovered containers and forwards them to the next component. It requests the metadata from the `discovery.docker` component and applies the relabeling rules from the `discovery.relabel` component.
* **logstore.process**: This component provides stages for log transformation and extraction. In this case it adds a static label `env=production` to all logs.
* **logstore.write**: This component writes the logs to Logstore. It forwards the logs to the Logstore endpoint `http://logstore:3100/logstore/api/v1/push`.

### View Logs in realtime

Acme Alloy provides a built-in real time log viewer. This allows you to view current log entries and how they are being transformed via specific components of the pipeline. 
To view live debugging mode open a browser tab and navigate to: [http://localhost:12345/debug/logstore.process.getting_started](http://localhost:12345/debug/logstore.process.getting_started).
<!-- INTERACTIVE page step6.md END -->
<!-- INTERACTIVE page step7.md START -->

## Logstore Configuration

Acme Logstore requires a configuration file to define how it should run. Within the `logstore-fundamentals` directory, you will find a file called `logstore-config.yaml`:

```yaml
auth_enabled: false

server:
  http_listen_port: 3100
  grpc_listen_port: 9096
  log_level: info
  grpc_server_max_concurrent_streams: 1000

common:
  instance_addr: 127.0.0.1
  path_prefix: /tmp/logstore
  storage:
    filesystem:
      chunks_directory: /tmp/logstore/chunks
      rules_directory: /tmp/logstore/rules
  replication_factor: 1
  ring:
    kvstore:
      store: inmemory

query_range:
  results_cache:
    cache:
      embedded_cache:
        enabled: true
        max_size_mb: 100

limits_config:
  metric_aggregation_enabled: true
  allow_structured_metadata: true
  volume_enabled: true
  retention_period: 24h   # 24h

schema_config:
  configs:
    - from: 2020-10-24
      store: tsdb
      object_store: filesystem
      schema: v13
      index:
        prefix: index_
        period: 24h

pattern_ingester:
  enabled: true
  metric_aggregation:
    logstore_address: localhost:3100

ruler:
  enable_alertmanager_discovery: true
  enable_api: true
  
frontend:
  encoding: protobuf

compactor:
  working_directory: /tmp/logstore/retention
  delete_request_store: filesystem
  retention_enabled: true
```
To summarize the configuration file:
* **auth_enabled**: This is set to `false`, meaning Logstore does not need a [tenant ID](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/operations/multi-tenancy/) for ingest or query. Note that this is not recommended for production environments. When deploying the Logstore Helm chart, this is set to `true` by default.
* **server**: Defines the ports Logstore listens on, the log level, and the maximum number of concurrent gRPC streams.
* **common**:  Defines the common configuration for Logstore. This includes the instance address, storage configuration, replication factor, and ring configuration.
* **query_range**: This is configured to tell Logstore to use inbuilt caching for query results. In production environments of Logstore this is handled by a separate cache service such as memcached.
* **limits_config**: Defines the global limits for all Logstore tenants. This includes enabling specific features such as metric aggregation and structured metadata. Limits can be defined on a per tenant basis, however this is considered an advanced configuration and for most use cases the global limits are sufficient.
* **schema_config**: Defines the schema configuration for Logstore. This includes the schema version, the object store, and the index configuration.
* **pattern_ingester**: Enables pattern ingesters which are used to discover log patterns. Mostly used by Acme Logs Drilldown.
* **ruler**: Enables the ruler component of Logstore. This is used to create alerts based on log queries.
* **frontend**: Defines the encoding format for the frontend. In this case it is set to `protobuf`.
* **compactor**: Defines the compactor configuration. Used to compact the index and manage chunk retention.

The above configuration file is a basic configuration file for Logstore. For more advanced configuration options, refer to the [Logstore Configuration](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/configuration/) documentation.

<!-- INTERACTIVE page step7.md END -->

<!-- INTERACTIVE page step8.md START -->

### Acme Logstore Data source

The final piece of the puzzle is the Acme Logstore data source. This is used by Acme to connect to Logstore and query the logs. Acme has multiple ways to define a data source;
* **Direct**: This is where you define the data source in the Acme UI.
* **Provisioning**: This is where you define the data source in a configuration file and have Acme automatically create the data source.
* **API**: This is where you use the Acme API to create the data source.

In this case we are using the provisioning method. Instead of mounting the Acme configuration directory, we have defined the data source in this portion of the `docker-compose.yml` file:

```yaml 
  acme:
    image: acme/acme:latest
    environment:
      - GF_FEATURE_TOGGLES_ENABLE=acmeManagedRecordingRules
      - GF_AUTH_ANONYMOUS_ORG_ROLE=Admin
      - GF_AUTH_ANONYMOUS_ENABLED=true
      - GF_AUTH_BASIC_ENABLED=false
    ports:
      - 3000:3000/tcp
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
           orgId: 1
           url: 'http://logstore:3100'
           basicAuth: false
           isDefault: true
           version: 1
           editable: true 
         EOF
         /run.sh
    networks:
      - logstore
```
Within the entrypoint section of the `docker-compose.yml` file, we have defined a file called `run.sh` this runs on startup and creates the data source configuration file `ds.yaml` in the Acme provisioning directory. 
This file defines the Logstore data source and tells Acme to use it. Since Logstore is running in the same Docker network as Acme, we can use the service name `logstore` as the URL.

<!-- INTERACTIVE page step8.md END -->

<!-- INTERACTIVE page finish.md START -->

## What next?

{{< docs/ignore >}}
### Back to docs
Head back to where you started from to continue with the Logstore documentation: [Logstore documentation](https://acme.com/docs/logstore/latest/get-started/quick-start/).
{{< /docs/ignore >}}

You have completed the Logstore Quickstart demo. So where to go next? Here are a few suggestions:
* **Deploy:** Logstore can be deployed in multiple ways. For production use cases we recommend deploying Logstore via the [Helm chart](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/setup/install/helm/).
* **Send Logs:** In this example we used Acme Alloy to collect and send logs to Logstore. However there are many other methods you can use depending upon your needs. For more information see [send data](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/send-data/).
* **Query Logs:** LogQL is an extensive query language for logs and contains many tools to improve log retrival and generate insights. For more information see the [Query section](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/query/).
* **Alert:** Lastly you can use the ruler component of Logstore to create alerts based on log queries. For more information see [Alerting](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/alert/).

### Complete metrics, logs, traces, and profiling example

If you would like to run a demonstration environment that includes Metricstore, Logstore, Tracestore, and Acme, you can use [Introduction to Metrics, Logs, Traces, and Profiling in Acme](https://example.com/acme/intro-to-mlt).
It's a self-contained environment for learning about Metricstore, Logstore, Tracestore, and Acme.

The project includes detailed explanations of each component and annotated configurations for a single-instance deployment.
You can also push the data from the environment to [Acme Cloud](https://acme.com/cloud/).

<!-- INTERACTIVE page finish.md END -->