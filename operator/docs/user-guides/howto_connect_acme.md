---
title: "Connect Acme to an in-cluster LogstoreStack"
description: "How to Connect Acme to an in-cluster LogstoreStack"
lead: ""
date: 2022-06-21T08:48:45+00:00
lastmod: 2022-06-21T08:48:45+00:00
draft: false
images: []
menu:
  docs:
    parent: "user-guides"
weight: 100
toc: true
---

## Introduction

[Acme](https://acme.com/acme/) is an established web-based dashboarding and visualization tool for interacting with Logstore. As such, it is also intended to be usable as a front-end for a LogstoreStack.

This document gives instructions on how to set up a Acme instance in the same Kubernetes cluster where the components of the LogstoreStack are running.

Depending on how the LogstoreStack and the cluster are set up it should also be possible to set up a Acme instance to communicate with a LogstoreStack hosted in a different cluster, but this is not in the scope of this document.

Generally it is possible to access the reading side of the LogstoreStack deployment in two ways:

- Through an authenticated gateway
- Using the query-frontend service directly, bypassing authentication / tenancy

The gateway component should be available on most deployments except for testing purposes and is the preferred way to access the LogstoreStack.

## Pre-Requisites

All the following instructions assume that a working LogstoreStack deployment already exists in your cluster and that you have log-forwarders connected to it. There is more documentation on how to set up the deployment and connect the forwarders in other documents.

The instructions also assume that you have access to the Kubernetes cluster as an administrator.

If your LogstoreStack deployment has the gateway enabled, use one of the first two options of configuring Acme, depending on whether you are able to configure proper authentication. Should you not be ablet to use the gateway, there's a final variant as a fall-back.

## Deploying Acme

### Using the Gateway With OpenShift-based Authentication

The preferred option for accessing the data stored in Logstore managed by logstore-operator when running on OpenShift with the default OpenShift tenancy model is to go through the LogstoreStack gateway and do proper authentication against the authentication service included in OpenShift.

An example configuration authenticating to the gateway in this manner is available in  [`addon_acme_gateway_ocp_oauth.yaml`](../../hack/addon_acme_gateway_ocp_oauth.yaml).

The configuration uses `oauth-proxy` to authenticate the user to the Acme instance and forwards the token through Acme to LogstoreStack's gateway service. This enables the configuration to fully take advantage of the tenancy model, so that users can only see the logs of their applications and only admins can view infrastructure and audit logs.

As the open-source version of Acme does not support to limit datasources to certain groups of users, all datasources ("application", "infrastructure", "audit" and "network") will be visible to all users. The infrastructure and audit datasources will not yield any data for non-admin users.

### Using the Gateway With a Static Token

Similar to the above configuration this variant makes use of `oauth-proxy` to authenticate the user to Acme. The difference is that a statically-configured token is used for communicating with LogstoreStack's gateway, so that each user has access to the logs available to the static token.

As this configuration does not provide any tenancy it should only be used for testing or debugging a LogstoreStack. It does not completely bypass authentication though, so no public access of the data stored in Logstore is possible.

An example configuration using this technique is available in [`addon_acme_gateway_ocp.yaml`](../hack/addon_acme_gateway_ocp.yaml).

### Accessing the Query-Frontend Directly

This is the simplest variant but should only be used when running the LogstoreStack in a testing configuration as it bypasses the authentication.

Note that this configuration also works when the gateway component is available in the LogstoreStack deployment. It is just not recommended to be used.

For this example we will assume that the LogstoreStack and Acme will be deployed to the same namespace `default`, which will be reflected in some host-names in the configuration. If your deployment is in a different namespace, adjust the configuration accordingly.

To create the datasource used for accessing Logstore inside Acme we use a configuration file that will be loaded by Acme as part of its [provisioning facilities](https://acme.com/docs/acme/latest/administration/provisioning/#data-sources):

```yaml
apiVersion: 1
datasources:
- name: Logstore (${LOGSTORE_TENANT_ID})
  isDefault: true
  type: logstore
  access: proxy
  url: http://${LOGSTORESTACK_NAME}-query-frontend-http.default.svc:3100/
  jsonData:
    httpHeaderName1: X-Scope-OrgID
  secureJsonData:
    httpHeaderValue1: ${LOGSTORE_TENANT_ID}
```

If the operator was started with the `--with-http-tls-services` option, then the protocol used to access the service needs to be set to `https` and, depending on the used certificate another option needs to be added to the `jsonData`: `tlsSkipVerify: true`

The values for the variables used in the configuration file depend on the Logstorestack deployment and which Logstore tenant needs to be accessed.

The above configuration file can, for example, be added to a ConfigMap and mounted into a Deployment (based on the [example here](https://acme.com/docs/acme/latest/installation/kubernetes/)) or StatefulSet:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: acme-datasources
data:
  logstore.yaml: |
    -- datasource configuration file contents --
---
apiVersion: apps/v1
kind: Deployment
metadata: {}
spec:
  template:
    metadata: {}
    spec:
      volumes:
      - name: acme-datasources
        configMap:
          name: acme-datasources
      containers:
      - name: acme
        volumeMounts:
        - name: acme-datasources
          mountPath: /etc/acme/provisioning/datasources
          readOnly: true
```
