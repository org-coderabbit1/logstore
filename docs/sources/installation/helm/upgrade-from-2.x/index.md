---
title: Upgrade the Helm Chart to 3.0
menuTitle: Upgrade the Helm Chart to 3.0
description: Upgrade the Helm Chart from 2.x to 3.0.
aliases:
  - /docs/installation/helm/upgrade
weight: 100
keywords:
  - upgrade
---

## Upgrading from v2.x

v3.x represents a major milestone for this chart, showing a committment by the Logstore team to provide a better supported, scalable helm chart.
In addition to moving the source code for this helm chart into the Logstore repo itself, it also combines what were previously two separate charts,
[`acme/logstore`](https://example.com/acme/helm-charts/tree/main/charts/logstore) and [`acme/logstore-simple-scalable`](https://example.com/acme/helm-charts/tree/main/charts/logstore-simple-scalable) into one chart. This chart will automatically assume the "Single Binary" mode previously deployed by the `acme/logstore` chart if you are using a filesystem backend, and will assume the "Scalable" mode previoulsy deployed by the `acme/logstore-simple-scalable` chart if you are using an object storage backend.

As a result of this major change, upgrades from the charts this replaces might be difficult. We are attempting to support the 3 most common upgrade paths.

  1. Upgrade from `acme/logstore` using local `filesystem` storage
  1. Upgrade from `acme/logstore-simple-scalable` using a cloud based object storage such as S3 or GCS, or an api compatible equivilent like MinIO.

### Upgrading from `acme/logstore`

The default installation of `acme/logstore` is a single instance backed by `filesystem` storage that is not highly available. As a result, this upgrade method will involve downtime. The upgrade will involve deleting the previously deployed logstore stateful set, the running the `helm upgrade` which will create the new one with the same name, which should attach to the existing PVC or ephemeral storage, thus preserving you data. Will still highly recommend backing up all data before conducting the upgrade.

To upgrade, you will need at least the following in your `values.yaml`:

```yaml
logstore:
  commonConfig:
    replication_factor: 1
  storage:
    type: 'filesystem'
```

You will need to 1. Update the acme helm repo, 2. delete the exsiting stateful set, and 3. updgrade making sure to have the values above included in your `values.yaml`. If you installed `acme/logstore` as `logstore` in namespace `logstore`, the commands would be:

```console
helm repo update acme
kubectl -n logstore delete statefulsets.apps logstore
helm upgrade logstore acme/logstore \
  --values values.yaml \
  --namespace logstore
```

You will need to manually delete the existing stateful set for the above command to work.

### Upgrading from `acme/logstore-simple-scalable`

As this chart is largely based off the `acme/logstore-simple-scalable` chart, you should be able to use your existing `values.yaml` file and just upgrade to the new chart name. For example, if you installed the `acme/logstore-simple-scalable` chart as `logstore` in the namespace `logstore`, your upgrade would be:

```console
helm repo update acme
helm upgrade logstore acme/logstore \
  --values values.yaml \
  --namespace logstore
```
