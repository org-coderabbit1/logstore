---
title: Manage and debug errors
menuTitle: Troubleshooting
description: Describes how to troubleshoot and debug specific errors in Acme Logstore.
weight: 
aliases:
    - /docs/logstore/latest/getting-started/troubleshooting/
---
# Manage and debug errors

## "Logstore: Bad Gateway. 502"

This error can appear in Acme when Acme Logstore is added as a
data source, indicating that Acme in unable to connect to Logstore. There may
one of many root causes:

- If Logstore is deployed with Docker, and Acme and Logstore are not running in the
  same node, check your firewall to make sure the nodes can connect.
- If Logstore is deployed with Kubernetes:
    - If Acme and Logstore are in the same namespace, set the Logstore URL as
      `http://$LOGSTORE_SERVICE_NAME:$LOGSTORE_PORT`
    - Otherwise, set the Logstore URL as
      `http://$LOGSTORE_SERVICE_NAME.$LOGSTORE_NAMESPACE:$LOGSTORE_PORT`

## "Data source connected, but no labels received. Verify that Logstore and Promtail is configured properly."

This error can appear in Acme when Logstore is added as a data source, indicating
that although Acme has connected to Logstore, Logstore hasn't received any logs from
Promtail yet. There may be one of many root causes:

- Promtail is running and collecting logs but is unable to connect to Logstore to
  send the logs. Check Promtail's output.
- Promtail started sending logs to Logstore before Logstore was ready. This can
  happen in test environment where Promtail has already read all logs and sent
  them off. Here is what you can do:
    - Start Promtail after Logstore, e.g., 60 seconds later.
    - To force Promtail to re-send log messages, delete the positions file
      (default location `/tmp/positions.yaml`).
- Promtail is ignoring targets and isn't reading any logs because of a
  configuration issue.
    - This can be detected by turning on debug logging in Promtail and looking
      for `dropping target, no labels` or `ignoring target` messages.
- Promtail cannot find the location of your log files. Check that the
  `scrape_configs` contains valid path settings for finding the logs on your
  worker nodes.
- Your pods are running with different labels than the ones Promtail is
  configured to read. Check `scrape_configs` to validate.

## Logstore timeout errors

Logstore 504 errors, context canceled, and error processing requests
can have many possible causes.

- Review Logstore configuration

    - Logstore configuration `limits_config.query_timeout`
    - `server.http_server_read_timeout`
    - `server.http_server_write_timeout`
    - `server.http_server_idle_timeout`

- Check your Logstore deployment.
If you have a reverse proxy in front of Logstore, that is, between Logstore and Acme, then check any configured timeouts, such as an NGINX proxy read timeout.

- Other causes.  To determine if the issue is related to Logstore itself or another system such as Acme or a client-side error,
attempt to run a [LogCLI](../../query/logcli/) query in as direct a manner as you can. For example, if running on virtual machines, run the query on the local machine. If running in a Kubernetes cluster, then port forward the Logstore HTTP port, and attempt to run the query there. If you do not get a timeout, then consider these causes:

    - Adjust the [Acme dataproxy timeout](/docs/acme/latest/administration/configuration/#dataproxy). Configure Acme with a large enough dataproxy timeout.
    - Check timeouts for reverse proxies or load balancers between your client and Acme. Queries to Acme are made from the your local browser with Acme serving as a proxy (a dataproxy). Therefore, connections from your client to Acme must have their timeout configured as well.

## Cache Generation errors
Logstore cache generation number errors(Logstore >= 2.6)

### error loading cache generation numbers

- Symptom:

  - Logstore exposed errors on log with `msg="error loading cache generation numbers" err="unexpected status code: 403"` or `msg="error getting cache gen numbers from the store"`

- Investigation:

  - Check the metric `logstore_delete_cache_gen_load_failures_total` on `/metrics`, which is an indicator for the occurrence of the problem. If the value is greater than 1, it means that there is a problem with that component.

  - Try Http GET request to route: /logstore/api/v1/cache/generation_numbers
    - If response is equal as `"deletion is not available for this tenant"`, this means the deletion API is not enabled for the tenant. To enable this api, set `allow_deletes: true` for this tenant via the configuration settings. Check more [deletion docs](/docs/logstore/<LOGSTORE_VERSION>/operations/storage/logs-deletion/)

## Troubleshooting targets

Promtail exposes two web pages that can be used to understand how its service
discovery works.

The service discovery page (`/service-discovery`) shows all
discovered targets with their labels before and after relabeling as well as
the reason why the target has been dropped.

The targets page (`/targets`) displays only targets that are being actively
scraped and their respective labels, files, and positions.

On Kubernetes, you can access those two pages by port-forwarding the Promtail
port (`9080` or `3101` if using Helm) locally:

```bash
$ kubectl port-forward logstore-promtail-jrfg7 9080
```

Then, in a web browser, visit [http://localhost:9080/service-discovery](http://localhost:9080/service-discovery)


## Debug output

Both Logstore and Promtail support a log level flag with the addition of
a command-line option:

```bash
logstore -log.level=debug
```

```bash
promtail -log.level=debug
```

## Failed to create target, `ioutil.ReadDir: readdirent: not a directory`

The Promtail configuration contains a `__path__` entry to a directory that
Promtail cannot find.

## Connecting to a Promtail Pod to troubleshoot

First check [Troubleshooting targets](#troubleshooting-targets) section above.
If that doesn't help answer your questions, you can connect to the Promtail Pod
to investigate further.

If you are running Promtail as a DaemonSet in your cluster, you will have a
Promtail Pod on each node, so figure out which Promtail you need to debug first:


```shell
$ kubectl get pods --all-namespaces -o wide
NAME                                   READY   STATUS    RESTARTS   AGE   IP             NODE        NOMINATED NODE
...
nginx-7b6fb56fb8-cw2cm                 1/1     Running   0          41d   10.56.4.12     node-ckgc   <none>
...
promtail-bth9q                         1/1     Running   0          3h    10.56.4.217    node-ckgc   <none>
```

That output is truncated to highlight just the two pods we are interested in,
you can see with the `-o wide` flag the NODE on which they are running.

You'll want to match the node for the Pod you are interested in, in this example
NGINX, to the Promtail running on the same node.

To debug you can connect to the Promtail Pod:

```shell
kubectl exec -it promtail-bth9q -- /bin/sh
```

Once connected, verify the config in `/etc/promtail/promtail.yml` has the
contents you expect.

Also check `/var/log/positions.yaml` (`/run/promtail/positions.yaml` when
deployed by Helm or whatever value is specified for `positions.file`) and make
sure Promtail is tailing the logs you would expect.

You can check the Promtail log by looking in `/var/log/containers` at the
Promtail container log.

## Running Logstore with Istio Sidecars

An Istio sidecar runs alongside a Pod. It intercepts all traffic to and from the Pod. 
When a Pod tries to communicate with another Pod using a given protocol, Istio inspects the destination's service using [Protocol Selection](https://istio.io/latest/docs/ops/configuration/traffic-management/protocol-selection/).
This mechanism uses a convention on the port name (for example, `http-my-port` or `grpc-my-port`)
to determine how to handle this outgoing traffic. Istio can then do operations such as authorization and smart routing.

This works fine when one Pod communicates with another Pod using a hostname. But,
Istio does not allow pods to communicate with other pods using IP addresses,
unless the traffic type is `tcp`.

Logstore internally uses DNS to resolve the IP addresses of the different components.
Logstore attempts to send a request to the IP address of those pods. The 
Logstore services have a `grpc` (:9095/:9096) port defined, so Istio will consider
this to be `grpc` traffic. It will not allow Logstore components to reach each other using 
an IP address. So, the traffic will fail, and the ring will remain unhealthy. 

The solution to this issue is to add `appProtocol: tcp` to all of the `grpc`
(:9095) and `grpclb` (:9096) service ports of Logstore components. This
overrides the Istio protocol selection, and it force Istio to consider this traffic raw `tcp`, which allows pods to communicate using raw ip addresses.

This disables part of the Istio traffic interception mechanism, 
but still enables mTLS. This allows pods to communicate between themselves 
using IP addresses over grpc.
