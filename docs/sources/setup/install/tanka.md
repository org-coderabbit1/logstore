---
title: Tanka 
menuTitle: Install using Tanka
description: Describes how to install Logstore using Tanka.
aliases: 
  - ../../installation/tanka/
weight: 300
---
# Tanka

[Tanka](https://tanka.dev) is a reimplementation of
[Ksonnet](https://ksonnet.io) that Acme Labs created after Ksonnet was
deprecated. Tanka is used by Acme Labs to run Acme Logstore in production.

The Tanka installation runs the Logstore cluster in microservices mode.

## Prerequisites

Install the latest version of Tanka (version v0.31.0 or a more recent version) for the `tk env`
commands. Prebuilt binaries for Tanka can be found at the [Tanka releases
URL](https://example.com/acme/tanka/releases).

In your config repo, if you don't have a Tanka application, create a folder and
call `tk init` inside of it. Then create an environment for Logstore and provide the
URL for the Kubernetes API server to deploy to (e.g., `https://localhost:6443`):

```
mkdir <application name>
cd <application name>
tk init
tk env add environments/logstore --namespace=logstore --server=<Kubernetes API server>
```

Install `jsonnet-bundler` (`jb`), find instructions for your platform in Tanka's [installation docs](https://tanka.dev/install#jsonnet-bundler).

## Deploying

Download and install the Logstore and Promtail module using `jb` (version v0.6.0 or a more recent version):

```bash
jb init  # not required if you already ran `tk init`
jb install example.com/acme/logstore/production/ksonnet/logstore@main
jb install example.com/acme/logstore/production/ksonnet/promtail@main
```

Revise the YAML contents of `environments/logstore/main.jsonnet`, updating these variables:

- Update the `username`, `password`, and the relevant `htpasswd` variable values.
- Update the S3 or GCS variable values, depending on your object storage type. See [storage_config](/docs/logstore/<LOGSTORE_VERSION>/configuration/#storage_config) for more configuration details.
- Remove from the configuration the S3 or GCS object storage variables that are not part of your setup.
- Update the Promtail configuration `container_root_path` variable's value to reflect your root path for the Docker daemon. Run `docker info | grep "Root Dir"` to acquire your root path.
- Update the `from` value in the Logstore `schema_config` section to no more than 14 days prior to the current date. The `from` date represents the first day for which the `schema_config` section is valid. For example, if today is `2021-01-15`, set `from` to `2021-01-01`. This recommendation is based on the Logstore default acceptance of log lines up to 14 days in the past. The `reject_old_samples_max_age` configuration variable controls the acceptance range.


```jsonnet
local gateway = import 'logstore/gateway.libsonnet';
local logstore = import 'logstore/logstore.libsonnet';
local promtail = import 'promtail/promtail.libsonnet';

logstore + promtail + gateway {
  _config+:: {
    namespace: 'logstore',
    htpasswd_contents: 'logstore:$apr1$H4yGiGNg$ssl5/NymaGFRUvxIV1Nyr.',

    // S3 variables -- Remove if not using s3
    storage_backend: 's3,dynamodb',
    s3_access_key: 'key',
    s3_secret_access_key: 'secret access key',
    s3_address: 'url',
    s3_bucket_name: 'logstore-test',
    dynamodb_region: 'region',

    // GCS variables -- Remove if not using gcs
    storage_backend: 'bigtable,gcs',
    bigtable_instance: 'instance',
    bigtable_project: 'project',
    gcs_bucket_name: 'bucket',

    //Update the object_store and from fields
    logstore+: {
      schema_config: {
        configs: [{
          from: 'YYYY-MM-DD',
          store: 'boltdb-shipper',
          object_store: 'my-object-storage-backend-type',
          schema: 'v11',
          index: {
            prefix: '%s_index_' % $._config.table_prefix,
            period: '%dh' % $._config.index_period_hours,
          },
        }],
      },
    },

    //Update the container_root_path if necessary
    promtail_config+: {
      clients: [{
        scheme:: 'http',
        hostname:: 'gateway.%(namespace)s.svc' % $._config,
        username:: 'logstore',
        password:: 'password',
        container_root_path:: '/var/lib/docker',
      }],
    },

    replication_factor: 3,
    consul_replicas: 1,
  },
}
```

Run `tk show environments/logstore` to see the manifests that will be deployed to
the cluster. Run `tk apply environments/logstore` to deploy the manifests.
To delete the environment from cluster, run `tk delete environments/logstore`.
