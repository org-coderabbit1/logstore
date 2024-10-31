---
title: Install the microservice Helm chart
menuTitle: Install microservice Logstore
description: Installing Logstore in microservice mode using the Helm chart.
weight: 300
keywords: 
---

# Install the microservice Helm chart

This Helm Chart deploys Acme Logstore on Kubernetes.

This Helm chart deploys Logstore to run Logstore in [microservice mode](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/get-started/deployment-modes/#microservices-mode) within a Kubernetes cluster. The microservices deployment mode runs components of Logstore as distinct processes.

The default Helm chart deploys the following components:
- **Compactor component** (1 replica): Compacts and processes stored data.
- **Distributor component** (3 replicas, maxUnavailable: 2): Distributes incoming requests. Up to 2 replicas can be unavailable during updates.
- **IndexGateway component** (2 replicas, maxUnavailable: 1): Handles indexing. Up to 1 replica can be unavailable during updates.
- **Ingester component** (3 replicas): Handles ingestion of data.
- **Querier component** (3 replicas, maxUnavailable: 2): Processes queries. Up to 2 replicas can be unavailable during updates.
- **QueryFrontend component** (2 replicas, maxUnavailable: 1): Manages frontend queries. Up to 1 replica can be unavailable during updates.
- **QueryScheduler component** (2 replicas): Schedules queries.

{{< admonition type="note" >}}
We do not recommend running in Microservice mode with `filesystem` storage. For the purpose of this guide, we will use MinIO as the object storage to provide a complete example. 
{{< /admonition >}}


## Prerequisites

- Helm 3 or above. See [Installing Helm](https://helm.sh/docs/intro/install/).
- A running Kubernetes cluster (must have at least 3 nodes).



## Deploying the Helm chart for development and testing


1. Add [Acme's chart repository](https://example.com/acme/helm-charts) to Helm:

   ```bash
   helm repo add acme https://acme.github.io/helm-charts
   ```

1. Update the chart repository:

   ```bash
   helm repo update
   ```

1. Create the configuration file `values.yaml`. The example below illustrates how to deploy Logstore in test mode using MinIO as storage:

     ```yaml
     logstore:
        schemaConfig:
          configs:
            - from: "2024-04-01"
              store: tsdb
              object_store: s3
              schema: v13
              index:
                prefix: logstore_index_
                period: 24h
        ingester:
          chunk_encoding: snappy
        querier:
          # Default is 4, if you have enough memory and CPU you can increase, reduce if OOMing
          max_concurrent: 4
        pattern_ingester:
          enabled: true
        limits_config:
          allow_structured_metadata: true
          volume_enabled: true
          retention_period: 672h
        compactor:
          retention_enabled: true 
          delete_request_store: s3

     deploymentMode: Distributed

     ingester:
       replicas: 3 # To ensure data durability with replication
     querier:
       replicas: 3 # Improve query performance via parallelism
       maxUnavailable: 2
     queryFrontend:
       replicas: 2
       maxUnavailable: 1
     queryScheduler:
       replicas: 2
     distributor:
       replicas: 3 
       maxUnavailable: 2
     compactor:
       replicas: 1
     indexGateway:
       replicas: 2
       maxUnavailable: 1

     bloomPlanner:
       replicas: 0
     bloomBuilder:
       replicas: 0
     bloomGateway:
       replicas: 0
    
     backend:
        replicas: 0
     read:
        replicas: 0
     write:
        replicas: 0

     singleBinary:
        replicas: 0

      # This exposes the Logstore gateway so it can be written to and queried externaly
     gateway:
        service:
          type: LoadBalancer


     # Enable minio for storage
     minio:
       enabled: true
     ```

1. Install or upgrade the Logstore deployment.
     - To install:
        ```bash
       helm install --values values.yaml logstore acme/logstore
       ```
    - To upgrade:
       ```bash
       helm upgrade --values values.yaml logstore acme/logstore
       ```
       

1. Verify that Logstore is running:
    ```bash
    kubectl get pods -n logstore
    ```
    The output should an output similar to the following:

    ```bash
      logstore-canary-8thrx                      1/1     Running   0          167m
      logstore-canary-h965l                      1/1     Running   0          167m
      logstore-canary-th8kb                      1/1     Running   0          167m
      logstore-chunks-cache-0                    2/2     Running   0          167m
      logstore-compactor-0                       1/1     Running   0          167m
      logstore-compactor-1                       1/1     Running   0          167m
      logstore-distributor-7c9bb8f4dd-bcwc5      1/1     Running   0          167m
      logstore-distributor-7c9bb8f4dd-jh9h8      1/1     Running   0          167m
      logstore-distributor-7c9bb8f4dd-np5dw      1/1     Running   0          167m
      logstore-gateway-77bc447887-qgc56          1/1     Running   0          167m
      logstore-index-gateway-0                   1/1     Running   0          167m
      logstore-index-gateway-1                   1/1     Running   0          166m
      logstore-ingester-zone-a-0                 1/1     Running   0          167m
      logstore-ingester-zone-b-0                 1/1     Running   0          167m
      logstore-ingester-zone-c-0                 1/1     Running   0          167m
      logstore-minio-0                           1/1     Running   0          167m
      logstore-querier-bb8695c6d-bv9x2           1/1     Running   0          167m
      logstore-querier-bb8695c6d-bz2rw           1/1     Running   0          167m
      logstore-querier-bb8695c6d-z9qf8           1/1     Running   0          167m
      logstore-query-frontend-6659566b49-528j5   1/1     Running   0          167m
      logstore-query-frontend-6659566b49-84jtx   1/1     Running   0          167m
      logstore-query-frontend-6659566b49-9wfr7   1/1     Running   0          167m
      logstore-query-scheduler-f6dc4b949-fknfk   1/1     Running   0          167m
      logstore-query-scheduler-f6dc4b949-h4nwh   1/1     Running   0          167m
      logstore-query-scheduler-f6dc4b949-scfwp   1/1     Running   0          167m
      logstore-results-cache-0                   2/2     Running   0          167m
    ```

## Object Storage Configuration

After testing Logstore with [MinIO](https://min.io/docs/minio/kubernetes/upstream/index.html), we recommend configuring Logstore with an object storage provider. The following examples shows how to configure Logstore with different object storage providers:

{{< admonition type="caution" >}}
When deploying Logstore using S3 Storage **DO NOT** use the default bucket names;  `chunk`, `ruler` and `admin`. Choose a unique name for each bucket. For more information see the following [security update](https://acme.com/blog/2024/06/27/acme-security-update-acme-logstore-and-unintended-data-write-attempts-to-amazon-s3-buckets/). This caution does not apply when you are using MinIO. When using MinIO we recommend using the default bucket names.
{{< /admonition >}}


{{< collapse title="S3" >}}

```yaml
# Example configuration for Logstore with S3 storage

logstore:
  schemaConfig:
    configs:
      - from: 2024-04-01
        store: tsdb
        object_store: s3
        schema: v13
        index:
          prefix: logstore_index_
          period: 24h
  storage_config:
    aws:
      region: <AWS region your bucket is in, for example, `eu-west-2`>
      bucketnames: <Your AWS bucket for chunk, for example, `aws-logstore-dev-chunk`>
      s3forcepathstyle: false
  ingester:
      chunk_encoding: snappy
  pattern_ingester:
      enabled: true
  limits_config:
    allow_structured_metadata: true
    volume_enabled: true
    retention_period: 672h # 28 days retention
  querier:
      max_concurrent: 4

  storage:
    type: s3
    bucketNames:
        chunks: <Your AWS bucket for chunk, for example, `aws-logstore-dev-chunk`>
        ruler: <Your AWS bucket for ruler, for example,  `aws-logstore-dev-ruler`>
        admin: <Your AWS bucket for admin, for example,  `aws-logstore-dev-admin`>
    s3:
      # s3 URL can be used to specify the endpoint, access key, secret key, and bucket name this works well for S3 compatible storage or if you are hosting Logstore on-premises and want to use S3 as the storage backend. Either use the s3 URL or the individual fields below (AWS endpoint, region, secret).
      s3: s3://access_key:secret_access_key@custom_endpoint/bucket_name
      # AWS endpoint URL
      endpoint: <your-endpoint>
      # AWS region where the S3 bucket is located
      region: <your-region>
      # AWS secret access key
      secretAccessKey: <your-secret-access-key>
      # AWS access key ID
      accessKeyId: <your-access-key-id>
      # AWS signature version (e.g., v2 or v4)
      signatureVersion: <your-signature-version>
      # Forces the path style for S3 (true/false)
      s3ForcePathStyle: false
      # Allows insecure (HTTP) connections (true/false)
      insecure: false
      # HTTP configuration settings
      http_config: {}

  deploymentMode: Distributed

  # Disable minio storage
  minio:
      enabled: false

  ingester:
    replicas: 3
  querier:
    replicas: 3
    maxUnavailable: 2
  queryFrontend:
    replicas: 2
    maxUnavailable: 1
  queryScheduler:
    replicas: 2
  distributor:
    replicas: 3
    maxUnavailable: 2
  compactor:
    replicas: 1
  indexGateway:
    replicas: 2
    maxUnavailable: 1

  bloomPlanner:
    replicas: 0
  bloomBuilder:
    replicas: 0
  bloomGateway:
    replicas: 0

  backend:
    replicas: 0
  read:
    replicas: 0
  write:
    replicas: 0

  singleBinary:
    replicas: 0

```
{{< /collapse >}}

{{< collapse title="Azure" >}}

```yaml
# Example configuration for Logstore with Azure Blob Storage

logstore:
  schemaConfig:
    configs:
      - from: "2024-04-01"
        store: tsdb
        object_store: azure
        schema: v13
        index:
          prefix: logstore_index_
          period: 24h
  ingester:
    chunk_encoding: snappy
  tracing:
    enabled: true
  querier:
    max_concurrent: 4

  storage:
    type: azure
    azure:
      # Name of the Azure Blob Storage account
      accountName: <your-account-name>
      # Key associated with the Azure Blob Storage account
      accountKey: <your-account-key>
      # Comprehensive connection string for Azure Blob Storage account (Can be used to replace endpoint, accountName, and accountKey)
      connectionString: <your-connection-string>
      # Flag indicating whether to use Azure Managed Identity for authentication
      useManagedIdentity: false
      # Flag indicating whether to use a federated token for authentication
      useFederatedToken: false
      # Client ID of the user-assigned managed identity (if applicable)
      userAssignedId: <your-user-assigned-id>
      # Timeout duration for requests made to the Azure Blob Storage account (in seconds)
      requestTimeout: <your-request-timeout>
      # Domain suffix of the Azure Blob Storage service endpoint (e.g., core.windows.net)
      endpointSuffix: <your-endpoint-suffix>
    bucketNames:
      chunks: "chunks"
      ruler: "ruler"
      admin: "admin"
deploymentMode: Distributed

ingester:
  replicas: 3
querier:
  replicas: 3
  maxUnavailable: 2
queryFrontend:
  replicas: 2
  maxUnavailable: 1
queryScheduler:
  replicas: 2
distributor:
  replicas: 3
  maxUnavailable: 2
compactor:
  replicas: 1
indexGateway:
  replicas: 2
  maxUnavailable: 1

bloomPlanner:
  replicas: 0
bloomBuilder:
  replicas: 0
bloomGateway:
  replicas: 0

backend:
  replicas: 0
read:
  replicas: 0
write:
  replicas: 0

singleBinary:
  replicas: 0

```
{{< /collapse >}}

To configure other storage providers, refer to the [Helm Chart Reference]({{< relref "../reference" >}}).

## Deploying the Logstore Helm chart to a Production Environment

{{< admonition type="note" >}}
We are actively working on providing more guides for deploying Logstore in production. 
{{< /admonition >}}

We recommend running Logstore at scale within a cloud environment like AWS, Azure, or GCP. The below guides will show you how to deploy a minimally viable production environment.
- [Deploy Logstore on AWS]({{< relref "../deployment-guides/aws" >}})

## Next Steps 
* Configure an agent to [send log data to Logstore](/docs/logstore/<LOGSTORE_VERSION>/send-data/).
* Monitor the Logstore deployment using the [Meta Monitoring Helm chart](/docs/logstore/<LOGSTORE_VERSION>/setup/install/helm/monitor-and-alert/)
