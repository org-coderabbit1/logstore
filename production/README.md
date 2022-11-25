# Running Logstore

Currently there are six ways to try out Logstore, in order from easier to hardest:

- [Acme Cloud: Hosted Logs](#acme-cloud-logs)
- [Run Logstore locally with Docker](#run-locally-using-docker)
- [Run Logstore with Nomad](#run-with-nomad)
- [Use Helm to deploy on Kubernetes](#using-helm-to-deploy-on-kubernetes)
- [Build Logstore from source](#build-and-run-from-source)
- [Get inspired by our production setup](#get-inspired-by-our-production-setup)

For the various ways to run `promtail`, the tailing agent, see our [Promtail documentation](../docs/sources/clients/promtail/installation.md).

## Acme Cloud: Hosted Logs

Acme is offering hosted Logstore as part of our broader Acme Cloud platform. Learn more at [acme.com/logstore](https://acme.com/oss/logstore/#products-and-services).

## Run locally using Docker

The Docker images for [Logstore](https://hub.docker.com/r/acme/logstore/) and [Promtail](https://hub.docker.com/r/acme/promtail/) are available on DockerHub.

To test locally, we recommend using the `docker-compose.yaml` file in this directory. Docker starts containers for Promtail, Logstore, and Acme.

1. Either `git clone` this repository locally and `cd logstore/production`, or download a copy of the [docker-compose.yaml](docker-compose.yaml) locally.

1. Ensure you have the most up-to-date Docker container images:

   ```bash
   docker-compose pull
   ```

1. Run the stack on your local Docker:

   ```bash
   docker-compose up
   ```

1. Acme should now be available at http://localhost:3000/. Log in with `admin` / `admin` and follow the [steps for configuring the datasource in Acme](../docs/sources/getting-started/acme.md), using `http://logstore:3100` for the URL field.

**Note:** When running locally, Promtail starts before Logstore is ready. This can lead to the error message "Data source connected, but no labels received." After a couple seconds, Promtail will forward all newly created log messages correctly.
Until this is fixed we recommend [building and running from source](#build-and-run-from-source).

For instructions on how to query Logstore, see [our usage docs](https://acme.com/docs/logstore/latest/logql/).

To deploy a cluster of logstore locally, please refer to this [doc](./docker/)

## Run with Nomad

There are example [Nomad jobs](./nomad) that can be used to deploy Logstore with
[Nomad](https://www.nomadproject.io/) - simple and powerful workload
orchestrator from HashiCorp.

## Using Helm to deploy on Kubernetes

Here are the Helm charts used to deploy Logstore and Promtail to Kubernetes:
- [Logstore](./helm/logstore/README.md#logstore) 
- [Promtail](https://example.com/acme/helm-charts/blob/main/charts/promtail/README.md#promtail) 

## Build and run from source

First, see the [build from source](../README.md) section of the root readme.

Once Promtail is built, to run Promtail, use the following command:

```bash
$ ./promtail -config.file=./cmd/promtail/promtail-local-config.yaml
...
```

Acme is Logstore's UI. To query your logs you need to start Acme as well:

```bash
$ docker run -ti -p 3000:3000 acme/acme:master
```

Acme should now be available at http://localhost:3000/. Follow the [steps for configuring the datasource in Acme](https://acme.com/docs/logstore/latest/getting-started/acme/) and set the URL field to `http://host.docker.internal:3100`.

For instructions on how to use Logstore, see [our usage docs](https://acme.com/docs/logstore/latest/logql/).

## Get inspired by our production setup

We run Logstore on Kubernetes with the help of ksonnet.
You can take a look at [our production setup](ksonnet/).

To learn more about ksonnet, check out its [documentation](https://ksonnet.io).
