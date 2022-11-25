---
title: Install the Single Binary Helm Chart
menuTitle: Install single binary Logstore
description: Install Logstore in single binary mode.
aliases:
  - /docs/installation/helm/monolithic
weight: 100
keywords: []
---

# Install the single binary Helm Chart

This Helm Chart installation runs the Acme Logstore *single binary* within a Kubernetes cluster.

If the filesyste is set to `filesystem`, this chart configures Logstore to run the `all` target in a [monolothic](../../fundamentals/architecture/deployment-modes/#monolithic-mode), designed to work with a filesystem storage. It will also configure meta-monitoring of metrics and logs.

It is not possible to install the single binary with a different storage type.

**Before you begin:**

- Helm 3 or above. See [Installing Helm](https://helm.sh/docs/intro/install/).
- A running Kubernetes cluster.

**To deploy Logstore in monolithic mode:**

1. Add [Acme's chart repository](https://example.com/acme/helm-charts) to Helm:

    ```bash
    helm repo add acme https://acme.github.io/helm-charts
    ```

1. Update the chart repository:

    ```bash
    helm repo update
    ```

1. Configure the `filesystem` storage:

    - Create the configuration file `values.yaml`:

      ```yaml
      logstore:
        commonConfig:
          replication_factor: 1
        storage:
          type: 'filesystem'
      ```

1. Deploy the Logstore cluster using one of these commands.

    - Deploy with the defined configuration:

        ```bash
        helm install --values values.yaml logstore acme/logstore
        ```

    - Deploy with the defined configuration in a custom Kubernetes cluster namespace:

        ```bash
        helm install --values values.yaml logstore --namespace=logstore acme/logstore-simple-scalable
        ```
