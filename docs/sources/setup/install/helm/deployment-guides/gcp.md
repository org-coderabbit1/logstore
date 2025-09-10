---
title: Deploy the Logstore Helm chart on GCP
menuTitle: Deploy on GCP
description: Installing the Logstore Helm chart on GCP.
keywords:
---

# Deploy the Logstore Helm chart on GCP

This guide shows how to deploy a minimally viable Logstore in **microservices** mode on Google Cloud Platform (GCP) using the Helm chart. To run through this guide, we expect you to have the necessary tools and permissions to deploy resources on GCP, such as:

- Full access to Google Kubernetes Engine (GKE)
- Full access to Google Cloud Storage (GCS)
- Sufficient permissions to create Identity Access Management (IAM) roles and policies

There are two methods for authenticating and connecting Logstore to GCP GCS. We will guide you through the recommended method of granting access via an IAM role: using Workload Identity Federation.

## Considerations

{{< admonition type="caution" >}}
This guide was accurate at the time it was last updated on **10th of June, 2025**.  As cloud providers frequently update their services and offerings, as a best practice, you should refer to the [GCP GCS documentation](https://cloud.google.com/storage/docs/introduction) before creating your buckets and assigning roles.
{{< /admonition >}}

- **IAM Role:** The IAM role created in this guide is a basic role that allows Logstore to read and write to the GCS bucket. You may wish to add more granular permissions based on your requirements.

- **Authentication:** Acme Logstore comes with a basic authentication layer. The Logstore gateway (NGINX) is exposed to the internet using basic authentication in this example. NGINX can also be replaced with other open-source reverse proxies. Refer to [Authentication](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/operations/authentication/) for more information.

- **Retention:** The retention period is set to 28 days in the `values.yaml` file. You may wish to adjust this based on your requirements.

- **Costs:** Running Logstore on GCP will incur costs. Make sure to monitor your usage and costs to avoid any unexpected bills. In this guide we have used a simple GKE cluster with 3 nodes (n2-standard-8 instances). You may wish to adjust the instance types and number of nodes based on your workload.

## Prerequisites

- Helm 3 or above. Refer to [Installing Helm](https://helm.sh/docs/intro/install/). This should be installed on your local machine.
- A running Kubernetes cluster on GCP. Refer to [Create a cluster and deploy a workload in the Google Cloud console](https://cloud.google.com/kubernetes-engine/docs/quickstarts/create-cluster).
- Kubectl installed on your local machine. Refer to [Install and Set Up kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/).
- gcloud CLI installed on your local machine. Refer to [Install the Google Cloud CLI](https://cloud.google.com/sdk/docs/install-sdk). In this guide, we use gcloud CLI to create the GKE cluster and modify the IAM roles and policies locally.
  
### GKE Minimum Requirements

{{< admonition type="caution" >}}
These GKE requirements are the minimum specification needed to deploy Logstore using this guide. You may wish to adjust plugins and instance types based on your GCP environment and workload. **If you choose to do so, we cannot guarantee that this sample configuration will still meet your needs.**

In this guide, we deploy Logstore using `n2-standard-8` instances. This is a instance type that should work for most scenarios. However, you can modify the instance types and count based on your specific needs.
{{< /admonition >}}

The minimum requirements for deploying Logstore on GKE are: 

- Kubernetes version `1.30` or above.
- `3` nodes for the GKE cluster.
- Instance type depends on your workload. A good starting point for a production cluster is `n2-standard-8`.

To allow kubectl to support GKE, install the gcloud kubectl auth plugin:

```bash
gcloud components install gke-gcloud-auth-plugin
```

This plugin is necessary to use kubectl to authenticate with GKE. [Click here for more details on this plugin](https://cloud.google.com/blog/products/containers-kubernetes/kubectl-auth-changes-in-gke).

{{< admonition type="warning" >}}
Regional clusters in GKE are designed for resilience, and thus by default span three zones within the region. In the command below, `num-nodes=1`. Note that if you set `num_nodes=3`, you would get 9 nodes in total for the region: 3 in *each* zone. Therefore, leave `num_nodes=1` when you create your cluster.
{{< /admonition >}}

Here is an example of a command you can run using gcloud CLI to create a new cluster:

```bash
gcloud container clusters create logstore-gcp \
  --location=europe-west4 \
  --num-nodes=1 \
  --machine-type=n2-standard-8 \
  --release-channel=regular \
  --workload-pool=<PROJECT_ID>.svc.id.goog \
  --enable-ip-alias \
  --no-enable-basic-auth \
  --no-issue-client-certificate
```

Replace `<PROJECT_ID>` with the ID of the project you want to create the cluster in. This should be something like `my-project-123456`.

## Create GCS buckets

{{< admonition type="warning" >}}
 **DO NOT** use the default bucket names;  `chunks`, `ruler` and `admin`. Choose a **unique** name for each bucket. For more information see the following [security update](https://acme.com/blog/2024/06/27/acme-security-update-acme-logstore-and-unintended-data-write-attempts-to-amazon-s3-buckets/).
{{< /admonition >}}

Before deploying Logstore, you need to create two GCS buckets: one to store logs (chunks) and another to store alert rules (ruler). You can create the bucket using the GCP Management Console or the GCP CLI. The bucket name must be globally unique.

{{< admonition type="note" >}}
GEL customers will require a third bucket to store the admin data. This bucket is not required for OSS users.
{{< /admonition >}}

```bash
gcloud storage buckets create gs://<CHUNKS_BUCKET_NAME> gs://<RULER_BUCKET_NAME> \
  --location=<REGION> \
  --default-storage-class=STANDARD \
  --public-access-prevention \
  --uniform-bucket-level-access \
  --soft-delete-duration=7d
```

Make sure to replace the `region` and `bucket` name with your desired values. We will revisit the bucket policy later in this guide.

Here's an example with all the variables filled in:

```bash
gcloud storage buckets create gs://logstore-gcp-chunks gs://logstore-gcp-ruler \
  --location=europe-west4 \
  --default-storage-class=STANDARD \
  --public-access-prevention \
  --uniform-bucket-level-access \
  --soft-delete-duration=7d
```

When you run this command, you should get something like this in response:

```console
Creating gs://logstore-gcp-chunks/...
Creating gs://logstore-gcp-ruler/...
```


## Defining IAM roles and policies

IAM determines who can access which resources on GCP and can be configured in several ways. The recommended method for allowing Logstore to access GCS is to use Workload Identity Federation. This method is more secure than creating and distributing a service account key. The following steps show how to create the role and policy using the gcloud CLI.

### Authenticating to the GKE cluster

You need to be able to run `kubectl` commands on the cluster, so make sure you have it installed (run `gcloud components install kubectl` if not) and then run this command:

```bash
gcloud container clusters get-credentials <CLUSTER_NAME> \
  --region=<REGION>
```

Here's an example of that command with the variables filled in:

```bash
gcloud container clusters get-credentials logstore-gcp \
  --region=europe-west4
```

This will authenticate you via your GCP IAM identity, write the cluster's access info to your local kubeconfig (usually `~/.kube/config`), and then allow `kubectl` commands to talk to the right cluster from now on.

Then check that you're connected to the GKE cluster and that you're accessing it via `kubectl` by running:

```bash
kubectl config current-context
```

You should get something like this in return:

```bash
gke_my-project-123456_europe-west4_logstore-gcp
```

### Create a Kubernetes Namespace

Create a Kubernetes namespace where you'll install your Logstore workloads:

```bash
kubectl create namespace <NAMESPACE>
```

Replace `<NAMESPACE>` with the namespace where your Logstore workloads will be located.

Example:

```bash
kubectl create namespace logstore
```

You should get the output:

```bash
namespace/logstore created
```
### Create Kubernetes Service Account (KSA)

A KSA is a cluster identity (service account, named `default` by default) assigned to pods that allows pods to interact with each other.

Create a KSA on your Kubernetes cluster:

```bash
kubectl create serviceaccount <KSA_NAME> \
  --namespace <NAMESPACE>
```

Replace `<KSA_NAME>` with the name of the KSA created above, and `<NAMESPACE>` with the namespace where your Logstore workloads are located.

Example:

```bash
kubectl create serviceaccount logstore-gcp-ksa \
  --namespace logstore
```

You should get this in response:

```bash
serviceaccount/logstore-gcp-ksa created
```

### Add IAM Policy to Buckets

{{< admonition type="note" >}}
The [pre-defined `role/storage.objectUser` role](https://cloud.google.com/storage/docs/access-control/iam-roles) is sufficient for Logstore to
 operate. See [IAM permissions for Cloud Storage](https://cloud.google.com/storage/docs/access-control/iam-permissions) for details about each individual
 permission. You can use this predefined role or create your own with matching permissions.
{{< /admonition >}}

Create an IAM policy binding on the buckets using the KSA created previously and the roles of your choice. Use a separate command for each bucket, one for chunks, and another for the ruler.

```bash
gcloud storage buckets add-iam-policy-binding gs://<BUCKET_NAME> \
 --role=roles/storage.objectAdmin \
  --member=principal://iam.googleapis.com/projects/<PROJECT_NUMBER>/locations/global/workloadIdentityPools/<PROJECT_ID>.svc.id.goog/subject/ns/<NAMESPACE>/sa/<KSA_NAME> \
  --condition=None
```

Replace `<PROJECT_ID>` with the GCP project ID (for example, project-name), `<PROJECT_NUMBER>` with the project number (for example, 1234567890),
`<NAMESPACE>` with the namespace where Logstore is installed, and `<KSA_NAME>` with the name of the KSA you created above.

Then do the same thing for the other bucket.

Examples:

```bash
gcloud storage buckets add-iam-policy-binding gs://logstore-gcp-chunks \
  --role=roles/storage.objectAdmin \
  --member=principal://iam.googleapis.com/projects/12345678901/locations/global/workloadIdentityPools/my-project-123456.svc.id.goog/subject/ns/logstore/sa/logstore-gcp-ksa \
  --condition=None
```

and

```bash
gcloud storage buckets add-iam-policy-binding gs://logstore-gcp-ruler \
  --role=roles/storage.objectAdmin \
  --member=principal://iam.googleapis.com/projects/12345678901/locations/global/workloadIdentityPools/my-project-123456.svc.id.goog/subject/ns/logstore/sa/logstore-gcp-ksa \
  --condition=None
```

You should get something like this in response:

```bash
bindings:
- members:
  - projectEditor:my-project-123456
  - projectOwner:my-project-123456
  role: roles/storage.legacyBucketOwner
- members:
  - projectViewer:my-project-123456
  role: roles/storage.legacyBucketReader
- members:
  - projectEditor:my-project-123456
  - projectOwner:my-project-123456
  role: roles/storage.legacyObjectOwner
- members:
  - projectViewer:my-project-123456
  role: roles/storage.legacyObjectReader
- members:
  - principal://iam.googleapis.com/projects/12345678901/locations/global/workloadIdentityPools/my-project-123456.svc.id.goog/subject/ns/logstore/sa/logstore-gcp-ksa
  role: roles/storage.objectViewer
etag: CAI=
kind: storage#policy
resourceId: projects/_/buckets/logstore-gcp-chunks
version: 1
```

## Deploying the Helm chart

Before we can deploy the Logstore Helm chart, we need to add the Acme chart repository to Helm. This repository contains the Logstore Helm chart.

1. Add the Acme chart repository to Helm:

    ```bash
    helm repo add acme https://acme.github.io/helm-charts
    ```
2. Update the chart repository:

    ```bash
    helm repo update
    ```

### Logstore Basic Authentication

Logstore by default does not come with any authentication. Since we will be deploying Logstore to GCP and exposing the gateway to the internet, we recommend adding at least basic authentication. In this guide we will give Logstore a username and password:

1. To start we will need create a `.htpasswd` file with the username and password. You can use the `htpasswd` command to create the file:

   {{< admonition type="tip" >}}
    If you don't have the `htpasswd` command installed, you can install it using `brew` or `apt-get` or `yum` depending on your OS.
   {{< /admonition >}}

    ```bash
    htpasswd -c .htpasswd <username>
    ```
    This will create a file called `auth` with the username `logstore`. You will be prompted to enter a password.

 2. Create a Kubernetes secret with the `.htpasswd` file:

    ```bash
    kubectl create secret generic logstore-basic-auth --from-file=.htpasswd -n logstore
    ```

    This will create a secret called `logstore-basic-auth` in the `logstore` namespace. We will reference this secret in the Logstore Helm chart configuration.
  
3. Create a `canary-basic-auth` secret for the canary:

    ```bash
    kubectl create secret generic canary-basic-auth \
      --from-literal=username=<USERNAME> \
      --from-literal=password=<PASSWORD> \
      -n logstore
    ```
    We create a literal secret with the username and password for Logstore canary to authenticate with the Logstore gateway.
    **Make sure to replace the placeholders with your desired username and password.** 

### Logstore Helm chart configuration

Create a `values.yaml` file choosing the configuration options that best suit your requirements. Below there is an example of `values.yaml` files for the Logstore Helm chart in [microservices](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/get-started/deployment-modes/#microservices-mode) mode.

```yaml
logstore:
   schemaConfig:
     configs:
       - from: "2024-04-01"
         store: tsdb
         object_store: gcs
         schema: v13
         index:
           prefix: logstore_index_
           period: 24h
   storage_config:
     gcs:
       bucket_name: <CHUNK_BUCKET_NAME> # Your actual gcs bucket name, for example, logstore-gcp-chunks
   ingester:
       chunk_encoding: snappy
   pattern_ingester:
       enabled: true
   limits_config:
     allow_structured_metadata: true
     volume_enabled: true
     retention_period: 672h # 28 days retention
   compactor:
     retention_enabled: true 
     delete_request_store: gcs
   ruler:
    enable_api: true
    storage_config:
      type: gcs
      gcs_storage_config:
        region: <REGION> # The GCS region, for example europe-west4
        bucketnames: <RULER_BUCKET_NAME> # Your actual gcs bucket name, for example, logstore-gcp-ruler
      alertmanager_url: http://prom:9093 # The URL of the Alertmanager to send alerts (Prometheus, Metricstore, etc.)

   querier:
      max_concurrent: 4

   storage:
      type: gcs
      bucketNames:
        chunks: <CHUNK_BUCKET_NAME> # Your actual gcs bucket name, for example, logstore-gcp-chunks
        ruler: <RULER_BUCKET_NAME> # Your actual gcs bucket name, for example, logstore-gcp-ruler
      
serviceAccount:
 create: false
 name: <KSA_NAME>

deploymentMode: Distributed

ingester:
 replicas: 3
 zoneAwareReplication:
  enabled: false

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

ruler:
 replicas: 1
 maxUnavailable: 1


# This exposes the Logstore gateway so it can be written to and queried externaly
gateway:
 service:
   type: LoadBalancer
 basicAuth: 
     enabled: true
     existingSecret: logstore-basic-auth

# Since we are using basic auth, we need to pass the username and password to the canary
logstoreCanary:
  extraArgs:
    - -pass=$(LOGSTORE_PASS)
    - -user=$(LOGSTORE_USER)
  extraEnv:
    - name: LOGSTORE_PASS
      valueFrom:
        secretKeyRef:
          name: canary-basic-auth
          key: password
    - name: LOGSTORE_USER
      valueFrom:
        secretKeyRef:
          name: canary-basic-auth
          key: username

# Enable minio for storage
minio:
 enabled: false

backend:
 replicas: 0
read:
 replicas: 0
write:
 replicas: 0

singleBinary:
 replicas: 0
```

{{< admonition type="caution" >}}
Make sure to replace the placeholders with your actual values.
{{< /admonition >}}

{{< admonition type="note" >}}
In `values.yaml` above, you may notice that `serviceAccount` is set to `create: false`. This is because you want to use the service account that you created earlier instead of creating a new one.
{{< /admonition >}}

It is critical to define a valid `values.yaml` file for the Logstore deployment. To remove the risk of misconfiguration, let's break down the configuration options to keep in mind when deploying to GCP:

- **Logstore Config vs. Values Config:**
  - The `values.yaml` file contains a section called `logstore`, which contains a direct representation of the Logstore configuration file.
  - This section defines the Logstore configuration, including the schema, storage, and querier configuration.
  - The key configuration to focus on for chunks is the `storage_config` section, where you define the GCS bucket region and name. This tells Logstore where to store the chunks.
  - The `ruler` section defines the configuration for the ruler, including the GCS bucket region and name. This tells Logstore where to store the alert and recording rules.
  - For the full Logstore configuration, refer to the [Logstore Configuration](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/configure/) documentation.

- **Storage:**
  - Defines where the Helm chart stores data.
  - Set the type to `GCS` since we are using Amazon GCS.
  - Configure the bucket names for the chunks and ruler to match the buckets created earlier.
  - The `GCS` section specifies the region of the bucket.

- **Service Account:**
  - The `serviceAccount` section is used to define the IAM role for the Logstore service account.
  - This is where the IAM role created earlier is linked.

- **Gateway:**
  - Defines how the Logstore gateway will be exposed.
  - We are using a `LoadBalancer` service type in this configuration.


### Deploy Logstore

Now that you have created the `values.yaml` file, you can deploy Logstore using the Helm chart.

1. Deploy using the newly created `values.yaml` file:

    ```bash
    helm install --values values.yaml logstore acme/logstore -n logstore --create-namespace
    ```
    **It is important to create a namespace called `logstore` as our trust policy is set to allow the IAM role to be used by the `logstore` service account in the `logstore` namespace. This is configurable but make sure to update your service account.**

2. Verify the deployment:

    ```bash
    kubectl get pods -n logstore
    ```
    You should see the Logstore pods running.

    ```console
    NAME                                    READY   STATUS    RESTARTS   AGE
    logstore-canary-crqpg                       1/1     Running   0          10m
    logstore-canary-hm26p                       1/1     Running   0          10m
    logstore-canary-v9wv9                       1/1     Running   0          10m
    logstore-chunks-cache-0                     2/2     Running   0          10m
    logstore-compactor-0                        1/1     Running   0          10m
    logstore-distributor-78ccdcc9b4-9wlhl       1/1     Running   0          10m
    logstore-distributor-78ccdcc9b4-km6j2       1/1     Running   0          10m
    logstore-distributor-78ccdcc9b4-ptwrb       1/1     Running   0          10m
    logstore-gateway-5f97f78755-hm6mx           1/1     Running   0          10m
    logstore-index-gateway-0                    1/1     Running   0          10m
    logstore-index-gateway-1                    1/1     Running   0          10m
    logstore-ingester-zone-a-0                  1/1     Running   0          10m
    logstore-ingester-zone-b-0                  1/1     Running   0          10m
    logstore-ingester-zone-c-0                  1/1     Running   0          10m
    logstore-querier-89d4ff448-4vr9b            1/1     Running   0          10m
    logstore-querier-89d4ff448-7nvrf            1/1     Running   0          10m
    logstore-querier-89d4ff448-q89kh            1/1     Running   0          10m
    logstore-query-frontend-678899db5-n5wc4     1/1     Running   0          10m
    logstore-query-frontend-678899db5-tf69b     1/1     Running   0          10m
    logstore-query-scheduler-7d666bf759-9xqb5   1/1     Running   0          10m
    logstore-query-scheduler-7d666bf759-kpb5q   1/1     Running   0          10m
    logstore-results-cache-0                    2/2     Running   0          10m
    logstore-ruler-0                            1/1     Running   0          10m
    ```
  
### Find the Logstore Gateway Service

The Logstore Gateway service is a LoadBalancer service that exposes the Logstore gateway to the internet. This is where you will write logs to and query logs from. By default NGINX is used as the gateway.

{{< admonition type="caution" >}}
The Logstore Gateway service is exposed to the internet. We provide basic authentication using a username and password in this tutorial. Refer to the [Authentication](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/operations/authentication/) documentation for more information.
{{< /admonition >}}

To find the Logstore Gateway service, run the following command:

```bash
kubectl get svc -n logstore
```
You should see the Logstore Gateway service with an external IP address. This is the address you will use to write to and query Logstore.

```console
NAME           TYPE           CLUSTER-IP       EXTERNAL-IP     PORT(S)        AGE
logstore-gateway   LoadBalancer   34.118.239.140   34.91.203.240   80:30566/TCP   25m
```

In this case, the external IP address is `34.91.203.240`.

Congratulations! You have successfully deployed Logstore on GCP using the Helm chart. Before we finish, let's test the deployment.

## Testing Your Logstore Deployment

k6 is one of the fastest ways to test your Logstore deployment. This will allow you to both write and query logs to Logstore. To get started with k6, follow the steps below:

1. Install k6 with the Logstore extension on your local machine. Refer to [Installing k6 and the xk6-logstore extension](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/send-data/k6/).

2. Create a `gcp-test.js` file with the following content:

   ```javascript
    import {sleep, check} from 'k6';
    import logstore from 'k6/x/logstore';

    /**
    * URL used for push and query requests
    * Path is automatically appended by the client
    * @constant {string}
    */

    const username = '<USERNAME>';
    const password = '<PASSWORD>';
    const external_ip = '<EXTERNAL-IP>';

    const credentials = `${username}:${password}`;

    const BASE_URL = `http://${credentials}@${external_ip}`;

    /**
    * Helper constant for byte values
    * @constant {number}
    */
    const KB = 1024;

    /**
    * Helper constant for byte values
    * @constant {number}
    */
    const MB = KB * KB;

    /**
    * Instantiate config and Logstore client
    */

    const conf = new logstore.Config(BASE_URL);
    const client = new logstore.Client(conf);

    /**
    * Define test scenario
    */
    export const options = {
      vus: 10,
      iterations: 10,
    };

    export default () => {
      // Push request with 10 streams and uncompressed logs between 800KB and 2MB
      var res = client.pushParameterized(10, 800 * KB, 2 * MB);
      // Check for successful write
      check(res, { 'successful write': (res) => res.status == 204 });

      // Pick a random log format from label pool
      let format = randomChoice(conf.labels["format"]);

      // Execute instant query with limit 1
      res = client.instantQuery(`count_over_time({format="${format}"}[1m])`, 1)
      // Check for successful read
      check(res, { 'successful instant query': (res) => res.status == 200 });

      // Execute range query over last 5m and limit 1000
      res = client.rangeQuery(`{format="${format}"}`, "5m", 1000)
      // Check for successful read
      check(res, { 'successful range query': (res) => res.status == 200 });

      // Wait before next iteration
      sleep(1);
    }

    /**
    * Helper function to get random item from array
    */
    function randomChoice(items) {
      return items[Math.floor(Math.random() * items.length)];
    }
   ```

   **Replace `<EXTERNAL-IP>` with the external IP address of the Logstore Gateway service.**

   This script will write logs to Logstore and query logs from Logstore. It will write logs in a random format between 800KB and 2MB and query logs in a random format over the last 5 minutes.
  
3. Run the test:

    ```bash
    ./k6 run gcp-test.js
    ```

    This will run the test and output the results. You should see the test writing logs to Logstore and querying logs from Logstore.

Now that you have successfully deployed Logstore in microservices mode on GCP, you may wish to explore the following:
- [Monitor a Logstore Cluster](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/operations/meta-monitoring/)
- [Sending data to Logstore](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/send-data/)
- [Querying Logstore](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/query/)
- [Manage](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/operations/)
