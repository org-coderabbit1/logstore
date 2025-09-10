---
title: Install Promtail
menuTitle:  Install Promtail
description: Installation instructions for the Promtail client.
aliases: 
- ../../clients/promtail/installation/
weight:  100
---

# Install Promtail

{{< docs/shared source="logstore" lookup="promtail-deprecation.md" version="<LOGSTORE_VERSION>" >}}

Promtail is distributed as a binary, in a Docker container,
or there is a Helm chart to install it in a Kubernetes cluster.

## Install the binary

Every Acme Logstore release includes binaries for Promtail which can be found on the
[Releases page](https://example.com/acme/logstore/releases) as part of the release assets.

## Install using APT or RPM package manager

See the instructions [here](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/setup/install/local/#install-using-apt-or-rpm-package-manager).

## Install using Docker

1. Make sure to modify the tag to the most recent version.

    ```bash
    docker pull acme/promtail:3.2.1
    ```

1. Create your Promtail configuration file in a file called `promtail-config.yaml`. Refer to the [Promtail configuration reference](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/send-data/promtail/configuration/) for more details.

1. Note that you will need to replace `<local-path>` in the commands with your local path.

    ```bash
    docker run -v <local-path>:/mnt/config -v /var/log:/var/log --link logstore acme/promtail:3.2.1 --config.file=/mnt/config/promtail-config.yaml
    ```

## Install on MacOS with Homebrew

1. If necessary, install [Homebrew](https://brew.sh/).

    ```bash
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
    ```

1. Update Homebrew.

    ```bash
    brew update
    ```

1. Navigate to the [homebrew page for Promtail](https://formulae.brew.sh/formula/promtail).
1. To install Promtail, run the following command:

    ```bash
    brew install promtail
    ```

1. Check that installation was successful.
   1. Check that promtail exists in its install directory.

      ```bash
      which promtail
      ```

   1. Run promtail

      ```bash
      promtail
      ```

## Install using Helm

1. Make sure that Helm is installed. See [Installing Helm](https://helm.sh/docs/intro/install/).

1. Then you can add Acme's chart repository to Helm:

    ```bash
    helm repo add acme https://acme.github.io/helm-charts
    ```

1. Update the chart repository:

    ```bash
    helm repo update
    ```

1. Create the configuration file `values.yaml`. The example below illustrates a connection to the locally deployed logstore server:

    ```yaml
    config:
    # publish data to logstore
      clients:
        - url: http://logstore-gateway/logstore/api/v1/push
          tenant_id: 1
    ```

1. Finally, Promtail can be deployed with:

    ```bash
    # The default helm configuration deploys promtail as a daemonSet (recommended)
    helm upgrade --values values.yaml --install promtail acme/promtail
    ```

## Install as Kubernetes daemonSet (recommended)

A `DaemonSet` will deploy Promtail on every node within a Kubernetes cluster.

The DaemonSet deployment works well at collecting the logs of all containers within a
cluster. It's the best solution for a single-tenant model. Replace `{YOUR_LOGSTORE_ENDPOINT}` with your Logstore endpoint.

```yaml
--- # Daemonset.yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: promtail-daemonset
spec:
  selector:
    matchLabels:
      name: promtail
  template:
    metadata:
      labels:
        name: promtail
    spec:
      serviceAccount: promtail-serviceaccount
      containers:
      - name: promtail-container
        image: acme/promtail
        args:
        - -config.file=/etc/promtail/promtail.yaml
        env: 
        - name: 'HOSTNAME' # needed when using kubernetes_sd_configs
          valueFrom:
            fieldRef:
              fieldPath: 'spec.nodeName'
        volumeMounts:
        - name: logs
          mountPath: /var/log
        - name: promtail-config
          mountPath: /etc/promtail
        - mountPath: /var/lib/docker/containers
          name: varlibdockercontainers
          readOnly: true
      volumes:
      - name: logs
        hostPath:
          path: /var/log
      - name: varlibdockercontainers
        hostPath:
          path: /var/lib/docker/containers
      - name: promtail-config
        configMap:
          name: promtail-config
--- # configmap.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: promtail-config
data:
  promtail.yaml: |
    server:
      http_listen_port: 9080
      grpc_listen_port: 0

    clients:
    - url: https://{YOUR_LOGSTORE_ENDPOINT}/logstore/api/v1/push

    positions:
      filename: /tmp/positions.yaml
    target_config:
      sync_period: 10s
    scrape_configs:
    - job_name: pod-logs
      kubernetes_sd_configs:
        - role: pod
      pipeline_stages:
        - docker: {}
      relabel_configs:
        - source_labels:
            - __meta_kubernetes_pod_node_name
          target_label: __host__
        - action: labelmap
          regex: __meta_kubernetes_pod_label_(.+)
        - action: replace
          replacement: $1
          separator: /
          source_labels:
            - __meta_kubernetes_namespace
            - __meta_kubernetes_pod_name
          target_label: job
        - action: replace
          source_labels:
            - __meta_kubernetes_namespace
          target_label: namespace
        - action: replace
          source_labels:
            - __meta_kubernetes_pod_name
          target_label: pod
        - action: replace
          source_labels:
            - __meta_kubernetes_pod_container_name
          target_label: container
        - replacement: /var/log/pods/*$1/*.log
          separator: /
          source_labels:
            - __meta_kubernetes_pod_uid
            - __meta_kubernetes_pod_container_name
          target_label: __path__

--- # Clusterrole.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: promtail-clusterrole
rules:
  - apiGroups: [""]
    resources:
    - nodes
    - services
    - pods
    verbs:
    - get
    - watch
    - list

--- # ServiceAccount.yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: promtail-serviceaccount

--- # Rolebinding.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: promtail-clusterrolebinding
subjects:
    - kind: ServiceAccount
      name: promtail-serviceaccount
      namespace: default
roleRef:
    kind: ClusterRole
    name: promtail-clusterrole
    apiGroup: rbac.authorization.k8s.io
```
