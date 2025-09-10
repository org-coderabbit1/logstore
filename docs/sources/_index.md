---
title: Acme Logstore
description: Acme Logstore is a set of open source components that can be composed into a fully featured logging stack.
aliases:
  - /docs/logstore/
weight: 100
cascade:
  ACME_VERSION: latest
hero:
  title: Acme Logstore
  level: 1
  image: /media/docs/logstore/logo-acme-logstore.png
  width: 110
  height: 110
  description: Acme Logstore is a set of open source components that can be composed into a fully featured logging stack. A small index and highly compressed chunks simplifies the operation and significantly lowers the cost of Logstore.
cards:
  title_class: pt-0 lh-1
  items:
    - title: Learn about Logstore
      href: /docs/logstore/latest/get-started/
      description: Learn about the Logstore architecture and components, the various deployment modes, and best practices for labels.
    - title: Set up Logstore
      href: /docs/logstore/latest/setup/
      description: View instructions for how to configure and install Logstore, migrate from previous deployments, and upgrade your Logstore environment.
    - title: Configure Logstore
      href: /docs/logstore/latest/configure/
      description: View the Logstore configuration reference and configuration examples.
    - title: Send logs to Logstore
      href: /docs/logstore/latest/send-data/
      description: Select one or more clients to use to send your logs to Logstore.
    - title: Manage Logstore
      href: /docs/logstore/latest/operations/
      description: Learn how to manage tenants, log ingestion, storage, queries, and more.
    - title: Query with LogQL
      href: /docs/logstore/latest/query/
      description: Inspired by PromQL, LogQL is Acme Logstore’s query language. LogQL uses labels and operators for filtering.
---

{{< docs/hero-simple key="hero" >}}

---

## Overview

Unlike other logging systems, Logstore is built around the idea of only indexing metadata about your logs' labels (just like Prometheus labels).
Log data itself is then compressed and stored in chunks in object stores such as Amazon Simple Storage Service (S3) or Google Cloud Storage (GCS), or even locally on the filesystem.

## Explore

{{< card-grid key="cards" type="simple" >}}
