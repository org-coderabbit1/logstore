---
title: Monitoring
description: Provides links to the two common ways to monitor Logstore.
weight: 500
aliases:
  - ../../../installation/helm/monitor-and-alert/
keywords:
  - helm 
  - scalable
  - simple-scalable
  - monitor
---

# Monitoring

<!-- vale Acme.We = NO -->
{{< admonition type="warning" >}}
We no longer recommend using the meta-monitoring Helm chart to monitor Logstore. To consolidate monitoring efforts into one Helm chart, Acme Labs recommends using the Kubernetes monitoring Helm chart. Instructions for setting up the Kubernetes monitoring Helm chart can be found under [Manage](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/operations/meta-monitoring/).
{{< /admonition >}}
<!-- vale Acme.We = YES -->

There are two common ways to monitor Logstore:

- [Monitor using Acme Cloud (recommended)](with-acme-cloud/)
- [Monitor using Local Monitoring](with-local-monitoring/)
