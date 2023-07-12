---
title: Install the monolithic Helm chart
menuTitle: Install monolithic Logstore
description: Install Logstore in monolithic, single binary mode.
aliases:
  - /docs/installation/helm/monolithic
weight: 300
keywords: 
---

# Install the monolithic Helm chart

This Helm Chart installation runs the Acme Logstore *single binary* within a Kubernetes cluster.

If you set the `singleBinary.replicas` value to 1, this chart configures Logstore to run the `all` target in a [monolithic mode]({{< relref "../../../get-started/deployment-modes#monolithic-mode" >}}), designed to work with a filesystem storage. It will also configure meta-monitoring of metrics and logs.
If you set the `singleBinary.replicas` value to 2 or more, this chart configures Logstore to run a *single binary* in a replicated, highly available mode.  When running replicas of a single binary, you must configure object storage.

**Before you begin: Software Requirements**

- Helm 3 or above. See [Installing Helm](https://helm.sh/docs/intro/install/).
- A running Kubernetes cluster

**To deploy Logstore in monolithic mode:**

1. Add [Acme's chart repository](https://example.com/acme/helm-charts) to Helm:

    ```bash
    helm repo add acme https://acme.github.io/helm-charts
    ```

1. Update the chart repository:

    ```bash
    helm repo update
    ```

1. Create the configuration file `values.yaml`:

    - If running a single replica of Logstore, configure the `filesystem` storage:

      ```yaml
      logstore:
        commonConfig:
          replication_factor: 1
        storage:
          type: 'filesystem'
      singleBinary:
        replicas: 1
      ```

    - If running Logstore with a replication factor greater than 1, set the desired number replicas and provide object storage credentials:

      ```yaml
      logstore:
        commonConfig:
          replication_factor: 3
        storage:
          type: 's3'
          s3:
            endpoint: foo.aws.com
            bucketnames: logstore-chunks
            secret_access_key: supersecret
            access_key_id: secret
      singleBinary:
        replicas: 3
      ```

1. Deploy the Logstore cluster using one of these commands.

    - Deploy with the defined configuration:

        ```bash
        helm install --values values.yaml logstore acme/logstore
        ```

    - Deploy with the defined configuration in a custom Kubernetes cluster namespace:

        ```bash
        helm install --values values.yaml logstore --namespace=logstore acme/logstore
        ```
