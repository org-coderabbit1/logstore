---
title: Deploy the Logstore Helm chart on Azure
menuTitle: Deploy on Azure
description: Installing the Logstore Helm chart on Azure.
keywords: 
---

# Deploy the Logstore Helm chart on Azure

This guide shows how to deploy a minimally viable Logstore in **microservice** mode on Azure using the Helm chart. In order to successfully complete this guide, you must have the necessary tools and permissions to deploy resources on Azure, such as:

- Full access to AKS (Azure Kubernetes Service)
- Full access to Azure Blob Storage
- Sufficient permissions to create federated credentials and roles in Azure AD (Active Directory)

There are three primary ways to authenticate Logstore with Azure:

- Hard coding a connection string - this is the simplest method but is not recommended for production environments.
- Manged identity
- Federated token

In this guide, we will use the federated token method to deploy Logstore on Azure. This method is more secure than hard coding a connection string and is more suitable for production environments.

## Considerations

{{< admonition type="caution" >}}
This guide was accurate at the time it was last updated on **8th of January, 2025**.  As cloud providers frequently update their services and offerings, as a best practice, you should refer to the [Azure documentation](https://learn.microsoft.com/en-us/azure/) before creating your storage account and assigning roles.
{{< /admonition >}}

- **AD Role:** In this tutorial we will create a role in Azure Active Directory (Azure AD) to allow Logstore to read and write from Azure Blob Storage. This role will be assigned to the Logstore service account. You may want to adjust the permissions based on your requirements.

- **Authentication:** Acme Logstore comes with a basic authentication layer. The Logstore gateway (NGINX) is exposed to the internet using basic authentication in this example. NGINX can also be replaced with other open-source reverse proxies. Refer to [Authentication](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/operations/authentication/) for more information.

- **Retention:** The retention period is set to 28 days in the `values.yaml` file. You may wish to adjust this based on your requirements.

- **Costs:** Running Logstore on Azure will incur costs. Make sure to monitor your usage and costs to avoid any unexpected bills. In this guide we have used a simple AKS cluster with 3 nodes and `Standard_E2ds_v5` instances. You may wish to adjust the instance types and number of nodes based on your workload.

## Prerequisites

- Helm 3 or above. Refer to [Installing Helm](https://helm.sh/docs/intro/install/). This should be installed on your local machine.
- Kubectl installed on your local machine. Refer to [Install and Set Up kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/).
- Azure CLI installed on your local machine. Refer to [Installing the Azure CLI](https://learn.microsoft.com/en-us/cli/azure/install-azure-cli). This is a requirement for following this guide as all resources will be created using the Azure CLI.
  
### AKS minimum requirements

Before creating an AKS cluster in Azure you need to create a resource group. You can create a resource group using the Azure CLI:

```bash
az group create --name <INSERT-NAME> -location <INSERT-AZURE-REGION>
```

{{< admonition type="caution" >}}
These AKS requirements are the minimum specification needed to deploy Logstore using this guide. You may wish to adjust the instance types based on your Azure environment and workload. **If you choose to do so, we cannot guarantee that this sample configuration will still meet your needs.**

In this guide, we deploy Logstore using `Standard_E2ds_v5` instances. This is to make sure we remain within the free tier limits for Azure. Which allows us to deploy up to 10 vCPUs within a region. We recommend for large production workloads to scale these nodes up to `Standard_D4_v5`.
{{< /admonition >}}

The minimum requirements for deploying Logstore on AKS are: 

- Kubernetes version `1.30` or above.
- `3` nodes for the AKS cluster.
- Instance type depends on your workload. A good starting point for a production cluster in the free tier is `Standard_E2ds_v5` instances and for large production workloads `Standard_D4_v5` instances.

Here is how to create an AKS cluster with the minimum requirements:

```bash
az aks create \
  --resource-group <MY_RESOURCE_GROUP_NAME> \
  --name <MY_AKS_CLUSTER_NAME> \
  --node-count 3 \
  --node-vm-size Standard_E2ds_v5 \
  --generate-ssh-keys \
  --enable-workload-identity \
  --enable-oidc-issuer
```

Note in the above command we have enabled workload identity and OIDC issuer. This is required for the Logstore service account to authenticate with Azure AD. If you have already created an AKS cluster, you can enable these features using the following command:

```bash
az aks update \
  --resource-group <MY_RESOURCE_GROUP_NAME> \
  --name <MY_AKS_CLUSTER_NAME> \
  --enable-workload-identity \
  --enable-oidc-issuer
```

The Azure CLI also lets you bind the AKS cluster to kubectl. You can do this by running the following command:

```bash
az aks get-credentials --resource-group <MY_RESOURCE_GROUP_NAME> --name <MY_AKS_CLUSTER_NAME>
```

## Configuring Azure Blob Storage

{{< admonition type="tip" >}}
 Consider using unique bucket names rather than:  `chunk`, `ruler`, and `admin`. Although Azure Blog Storage is not directly affected by this [security update](https://acme.com/blog/2024/06/27/acme-security-update-acme-logstore-and-unintended-data-write-attempts-to-amazon-s3-buckets/) it is a best practice to use unique container names for buckets.
{{< /admonition >}}

Before deploying Logstore, you need to create two Azure storage containers; one to store logs (chunks), the second to store alert rules. You can create the containers using the Azure CLI. Containers must exist inside a storage account.

{{< admonition type="note" >}}
GEL customers will require a third container to store the admin data. This container is not required for OSS users.
{{< /admonition >}}

1. Create a storage account:

    ```bash
    az storage account create \
    --name <NAME> \
    --location <REGION> \
    --sku Standard_ZRS \
    --encryption-services blob \
    --resource-group <MY_RESOURCE_GROUP_NAME>
    ```
    Replace the placeholders with your desired values.

1. Create the containers for chunks and ruler:

    ```bash
    az storage container create --account-name <STORAGE-ACCOUNT-NAME> --name <CHUNK-BUCKET-NAME> --auth-mode login && \
    az storage container create --account-name <STORAGE-ACCOUNT-NAME> --name <RULER-BUCKET-NAME> --auth-mode login
    ```
    Make sure `--account-name` matches the account you just created


With the storage account and containers created, you can now proceed to creating the Azure AD role and federated credentials.

## Creating the Azure AD role and federated credentials

The recommended way to authenticate Logstore with Azure Blob Storage is to use federated credentials. This method is more secure than hard coding a connection string directly into the Logstore configuration. In this next section, we will create an Azure AD role and federated credentials for Logstore to allow it to read and write from Azure Blob Storage:

1. Locate the OpenID Connect (OIDC)  issuer URL:

    ```bash
    az aks show \
    --resource-group <MY_RESOURCE_GROUP_NAME> \
    --name <MY_AKS_CLUSTER_NAME> \
    --query "oidcIssuerProfile.issuerUrl" \
    -o tsv
    ```
    This command will return the OIDC issuer URL. You will need this URL to create the federated credentials.

1. Generate a `credentials.json` file with the following content:
    ```json
    {
        "name": "LogstoreFederatedIdentity",
        "issuer": "<OIDC-ISSUER-URL>",
        "subject": "system:serviceaccount:logstore:logstore",
        "description": "Federated identity for Logstore accessing Azure resources",
        "audiences": [
          "api://AzureADTokenExchange"
        ]
    }
    ```
    Replace `<OIDC-ISSUER-URL>` with the OIDC issuer URL you found in the previous step.

1. Make sure you to save the `credentials.json` file before continuing.

1. Next generate an Azure directory `app`. We will use this to assign our federated credentials to:
   ```bash
    az ad app create \
    --display-name logstore \
    --query appId \
    -o tsv
   ```
    This will return the app ID. Save this for later use. If you need to find the app ID later you can run the following command:
    ```bash
    az ad app list --display-name logstore --query "[].appId" -o tsv
    ```

1. The app requires a service principal to authenticate with Azure AD. Create a service principal for the app:

    ```bash
    az ad sp create --id <APP-ID>
    ```
    Replace `<APP-ID>` with the app ID you generated in the previous step.

1. Next assign the federated credentials to the app:

    ```bash
    az ad app federated-credential create \
      --id <APP-ID> \
      --parameters credentials.json 
    ```
    Replace `<APP-ID>` with the app ID you generated in the previous step.

1. Lastly add a role assignment to the app:

    ```bash
    az role assignment create \
      --role "Storage Blob Data Contributor" \
      --assignee <APP-ID> \
      --scope /subscriptions/<SUBSCRIPTION-ID>/resourceGroups/<RESOURCE-GROUP>/providers/Microsoft.Storage/storageAccounts/<STORAGE-ACCOUNT-NAME>
    ```
    Replace the placeholders with your actual values.

Now that you have created the Azure AD role and federated credentials, you can proceed to deploying Logstore using the Helm chart.


## Deploying the Helm chart

The following steps require the use of `helm` and `kubectl`. Make sure you have run the `az` command to bind your AKS cluster to `kubectl`:

```bash
az aks get-credentials --resource-group <MY_RESOURCE_GROUP_NAME> --name <MY_AKS_CLUSTER_NAME>
```

Before we can deploy the Logstore Helm chart, we need to add the Acme chart repository to Helm. This repository contains the Logstore Helm chart.

1. Add the Acme chart repository to Helm:

    ```bash
    helm repo add acme https://acme.github.io/helm-charts
    ```
1. Update the chart repository:

    ```bash
    helm repo update
    ```
1. Create a new namespace for Logstore:

    ```bash
    kubectl create namespace logstore
    ```
### Logstore basic authentication

Logstore by default does not come with any authentication. Since we will be deploying Logstore to Azure and exposing the gateway to the internet, we recommend adding at least basic authentication. In this guide we will give Logstore a username and password:

1. To start we will need create a `.htpasswd` file with the username and password. You can use the `htpasswd` command to create the file:

   {{< admonition type="tip" >}}
    If you don't have the `htpasswd` command installed, you can install it using `brew` or `apt-get` or `yum` depending on your OS.
   {{< /admonition >}}

    ```bash
    htpasswd -c .htpasswd <username>
    ```
    This will create a file called `auth` with the username `logstore`. You will be prompted to enter a password.

 1. Create a Kubernetes secret with the `.htpasswd` file:

    ```bash
    kubectl create secret generic logstore-basic-auth --from-file=.htpasswd -n logstore
    ```

    This will create a secret called `logstore-basic-auth` in the `logstore` namespace. We will reference this secret in the Logstore Helm chart configuration.
  
1. Create a `canary-basic-auth` secret for the canary:

    ```bash
    kubectl create secret generic canary-basic-auth \
      --from-literal=username=<USERNAME> \
      --from-literal=password=<PASSWORD> \
      -n logstore
    ```
    We create a literal secret with the username and password for Logstore canary to authenticate with the Logstore gateway. Make sure to replace the placeholders with your desired username and password.

### Logstore Helm chart configuration

Create a `values.yaml` file choosing the configuration options that best suit your requirements. Below there is an example of `values.yaml` files for the Logstore Helm chart in [microservices](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/get-started/deployment-modes/#microservices-mode) mode.

```yaml
logstore:
   podLabels:
    "azure.workload.identity/use": "true" # Add this label to the Logstore pods to enable workload identity
   schemaConfig:
     configs:
       - from: "2024-04-01"
         store: tsdb
         object_store: azure
         schema: v13
         index:
           prefix: logstore_index_
           period: 24h
   storage_config:
     azure:
      account_name: "<INSERT-STORAGE-ACCOUNT-NAME>" 
      container_name: "<CHUNK-CONTAINER-NAME>" # Your actual Azure Blob Storage container name (logstore-azure-dev-chunks)
      use_federated_token: true # Use federated token for authentication
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
     delete_request_store: azure
   ruler:
    enable_api: true
    storage:
      type: azure
      azure:
        account_name: <INSERT-STORAGE-ACCOUNT-NAME>
        container_name: <RULER-CONTAINER-NAME> # Your actual Azure Blob Storage container name (logstore-azure-dev-ruler)
        use_federated_token: true # Use federated token for authentication
      alertmanager_url: http://prom:9093 # The URL of the Alertmanager to send alerts (Prometheus, Metricstore, etc.)

   querier:
      max_concurrent: 4

   storage:
      type: azure
      bucketNames:
        chunks: "<CHUNK-CONTAINER-NAME>" # Your actual Azure Blob Storage container name (logstore-azure-dev-chunks)
        ruler: "<RULER-CONTAINER-NAME>" # Your actual Azure Blob Storage container name (logstore-azure-dev-ruler)
        # admin: "admin-logstore-devrel" # Your actual Azure Blob Storage container name (logstore-azure-dev-admin)
      azure:
        accountName: <INSERT-STORAGE-ACCOUNT-NAME>
        useFederatedToken: true # Use federated token for authentication

# Define the Azure workload identity
serviceAccount:
  name: logstore
  annotations:
    "azure.workload.identity/client-id": "<APP-ID>" # The app ID of the Azure AD app
  labels:
    "azure.workload.identity/use": "true"

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

It is critical to define a valid `values.yaml` file for the Logstore deployment. To remove the risk of misconfiguration, let's break down the configuration options to keep in mind when deploying to Azure:

- **Logstore Config vs. Values Config:**
  - The `values.yaml` file contains a section called `logstore`, which contains a direct representation of the Logstore configuration file.
  - This section defines the Logstore configuration, including the schema, storage, and querier configuration.
  - The key configuration to focus on for chunks is the `storage_config` section, where you define the Azure container name and storage account. This tells Logstore where to store the chunks.
  - The `ruler` section defines the configuration for the ruler, including the Azure container name and storage account. This tells Logstore where to store the alert and recording rules.
  - For the full Logstore configuration, refer to the [Logstore Configuration](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/configure/) documentation.

- **Storage:**
  - Defines where the Helm chart stores data.
  - Set the type to `azure` since we are using Azure Blob Storage.
  - Configure the container names for the chunks and ruler to match the containers created earlier.
  - The `azure` section specifies the storage account name and also sets `useFederatedToken` to `true`. This tells Logstore to use federated credentials for authentication.

- **Service Account:**
  - The `serviceAccount` section is used to define the federated workload identity Logstore will use to authenticate with Azure AD.
  - We set the `azure.workload.identity/client-id` annotation to the app ID of the Azure AD app.

- **Gateway:**
  - Defines how the Logstore gateway will be exposed.
  - We are using a `LoadBalancer` service type in this configuration.


### Deploy Logstore

Now that you have created the `values.yaml` file, you can deploy Logstore using the Helm chart.

1. Deploy using the newly created `values.yaml` file:

    ```bash
    helm install --values values.yaml logstore acme/logstore -n logstore --create-namespace
    ```
    It is important to create a namespace called `logstore` as our federated credentials were generated with the  value `system:serviceaccount:logstore:logstore`. This translates to the `logstore` service account in the `logstore` namespace. This is configurable but make sure to update the federated credentials file first.

1. Verify the deployment:

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
  
### Find the Logstore gateway service

The Logstore gateway service is a load balancer service that exposes the Logstore gateway to the internet. This is where you will write logs to and query logs from. By default NGINX is used as the gateway.

{{< admonition type="caution" >}}
The Logstore gateway service is exposed to the internet. We provide basic authentication using a username and password in this tutorial. Refer to the [Authentication](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/operations/authentication/) documentation for more information.
{{< /admonition >}}

To find the Logstore gateway service, run the following command:

```bash
kubectl get svc -n logstore
```
You should see the Logstore gateway service with an external IP address. This is the address you will use to write to and query Logstore.

```console
  NAME                             TYPE           CLUSTER-IP       EXTERNAL-IP    PORT(S)              AGE
logstore-gateway                     LoadBalancer   10.100.201.74     134.236.21.145  80:30707/TCP         46m
```

Congratulations! You have successfully deployed Logstore on Azure using the Helm chart. Before we finish, let's test the deployment.

## Testing Your Logstore Deployment

k6 is one of the fastest ways to test your Logstore deployment. This will allow you to both write and query logs to Logstore. To get started with k6, follow the steps below:

1. Install k6 with the Logstore extension on your local machine. Refer to [Installing k6 and the xk6-logstore extension](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/send-data/k6/).

2. Create a `azure-test.js` file with the following content:

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

   Replace `<EXTERNAL-IP>` with the external IP address of the Logstore Gateway service.

   This script will write logs to Logstore and query logs from Logstore. It will write logs in a random format between 800KB and 2MB and query logs in a random format over the last 5 minutes.
  
3. Run the test:

    ```bash
    ./k6 run azure-test.js
    ```

    This will run the test and output the results. You should see the test writing logs to Logstore and querying logs from Logstore.

Now that you have successfully deployed Logstore in microservices mode on Microsoft Azure, you may wish to explore the following:

- [Sending data to Logstore](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/send-data/)
- [Querying Logstore](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/query/)
- [Manage Logstore](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/operations/)
