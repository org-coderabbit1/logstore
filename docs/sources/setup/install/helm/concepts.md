---
title: Helm chart components
menuTitle: Helm chart components
description: A short introduction to the components installed with the Logstore Helm Chart.
aliases:
  - ../../../installation/helm/concepts/
weight: 100
keywords:
  - dashboards
  - gateway
  - caching
---

# Helm chart components

This section describes the components installed by the Helm Chart.

## 3 methods of deployment

The Logstore chart supports three methods of deployment:
- [Monolithic]({{< relref "./install-monolithic" >}}) 
- [Simple Scalable]({{< relref "./install-scalable" >}})
- [Microservice]({{< relref "./install-microservices" >}})

By default, the chart installs in [Simple Scalable]({{< relref "./install-scalable" >}}) mode. This is the recommended method for most users. To understand the differences between deployment methods, see the [Logstore deployment modes]({{< relref "../../../get-started/deployment-modes" >}}) documentation.

## Monitoring Logstore

The Logstore Helm chart does not deploy self-monitoring by default. Logstore clusters can be monitored using the meta-monitoring stack, which monitors the logs, metrics, and traces of the Logstore cluster. There are two deployment options for this stack, see the installation instructions within [Monitoring]({{< relref "./monitor-and-alert" >}}).

{{< admonition type="note" >}}
The meta-monitoring stack replaces the monitoring section of the Logstore helm chart which is now **DEPRECATED**. See the [Monitoring]({{< relref "./monitor-and-alert" >}}) section for more information.
{{< /admonition >}}


## Canary

This chart installs the [Logstore Canary app]({{< relref "../../../operations/logstore-canary" >}}) by default. This is another tool to verify the Logstore deployment is in a healthy state. It can be disabled by setting `logstoreCanary.enabled=false`.

## Gateway

By default and inspired by Acme's [Tanka setup](https://example.com/acme/logstore/blob/main/production/ksonnet/logstore), the chart
installs the gateway component which is an NGINX that exposes the Logstore API and automatically proxies requests to the correct
Logstore components (read or write, or single instance in the case of filesystem storage).
The gateway must be enabled if an Ingress is required, since the Ingress exposes the gateway only.
If the gateway is enabled, Acme and log shipping agents, such as Promtail, should be configured to use the gateway.
If NetworkPolicies are enabled, they are more restrictive if the gateway is enabled.

## Caching

By default, this chart configures in-memory caching. If that caching does not work for your deployment, you should setup [memcache]({{< relref "../../../operations/caching" >}}).
