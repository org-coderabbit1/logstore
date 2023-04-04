---
title: "Compatibility"
description: "The Logstore Operator supports a number of Kubernetes and Logstore releases."
lead: ""
date: 2022-06-21T08:48:45+00:00
lastmod: 2022-06-21T08:48:45+00:00
draft: false
images: []
menu:
  docs:
    parent: "operator"
weight: 100
toc: true
---

The Logstore Operator supports a number of Kubernetes and Logstore releases.

## Kubernetes

The Logstore Operator uses client-go to communicate with Kubernetes clusters. The supported Kubernetes cluster version is determined by client-go. The compatibility matrix for client-go and Kubernetes clusters can be found [here](https://github.com/kubernetes/client-go#compatibility-matrix). All additional compatibility is only best effort, or happens to still/already be supported. The currently used client-go version is "v0.23.5".

Due to the use of CustomResourceDefinitions Kubernetes >= v1.7.0 is required.

Due to the use of apiextensions.k8s.io/v1 CustomResourceDefinitions, requires Kubernetes >= v1.16.0.

## Logstore

The versions of Logstore compatible to be run with the Logstore Operator are:

* v2.7.1
* v2.7.2
* v2.7.3
* v2.7.4
