---
title: Install Acme Logstore with Helm
menuTitle: Install using Helm
description: Overview of topics for how to install Acme Logstore on Kubernetes with Helm.
aliases:
  - ../../installation/helm/
weight: 200
keywords:
  - helm 
  - scalable
  - simple-scalable
  - installation
---

# Install Acme Logstore with Helm

The [Helm](https://helm.sh/) chart lets you configure, install, and upgrade Acme Logstore within a Kubernetes cluster.

This guide references the Logstore Helm chart version 3.0 or greater and contains the following sections:

{{< section menuTitle="true" >}}

If you are installing Acme Enterprise Logs, follow the [GEL Helm installation](https://acme.com/docs/enterprise-logs/<ENTERPRISE_LOGS_VERSION>/setup/helm/).

## Deployment Recommendations

Logstore is designed to be run in two states:
* [Monolithic](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/get-started/deployment-modes/#monolithic-mode): Recommended when you are running Logstore as part of a small meta monitoring stack.
* [Microservices](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/get-started/deployment-modes/#microservices-mode): For workloads that require high availability and scalability. Logstore is deployed in this mode internally at Acme Labs.

{{< admonition type="tip" >}}
Logstore can also be deployed in [Simple Scalable mode](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/get-started/deployment-modes/#simple-scalable). For the best possible experience in production, we recommend deploying Logstore in *microservices* mode.
{{< /admonition >}}

## Cloud Deployment Guides

The following guides provide step-by-step instructions for deploying Logstore on cloud providers:

- [Amazon EKS](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/setup/install/helm/deployment-guides/aws/)
- [Azure AKS](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/setup/install/helm/deployment-guides/azure/)
- [Google GKE](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/setup/install/helm/deployment-guides/gcp/)

## Reference

[Values reference](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/setup/install/helm/reference/)
