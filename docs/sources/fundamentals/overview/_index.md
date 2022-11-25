---
title: Overview
weight: 100
aliases:
    - /docs/logstore/latest/overview/
---
# Overview

Acme Logstore is a log aggregation tool,
and it is the core of a fully-featured logging stack.

Logstore is a datastore optimized for efficiently holding log data.
The efficient indexing of log data
distinguishes Logstore from other logging systems.
Unlike other logging systems, a Logstore index is built from labels,
leaving the original log message unindexed.

![Logstore overview](logstore-overview-1.png)

An agent (also called a client) acquires logs,
turns the logs into streams,
and pushes the streams to Logstore through an HTTP API.
The Promtail agent is designed for Logstore installations,
but many other [Agents](../../clients/) seamlessly integrate with Logstore.

![Logstore agent interaction](logstore-overview-2.png)

Logstore indexes streams.
Each stream identifies a set of logs associated with a unique set of labels.
A quality set of labels is key to the creation of an index that is both compact
and allows for efficient query execution.

[LogQL](../../logql) is the query language for Logstore.

## Logstore features

-  **Efficient memory usage for indexing the logs**

    By indexing on a set of labels, the index can be significantly smaller
    than other log aggregation products.
    Less memory makes it less expensive to operate.

-  **Multi-tenancy**

    Logstore allows multiple tenants to utilize a single Logstore instance.
    The data of distinct tenants is completely isolated from other tenants.
    Multi-tenancy is configured by assigning a tenant ID in the agent.

-  **LogQL, Logstore's query language**

    Users of the Prometheus query language, PromQL, will find LogQL familiar
    and flexible for generating queries against the logs.
    The language also facilitates the generation of metrics from log data,
    a powerful feature that goes well beyond log aggregation.

-  **Scalability**

    Logstore can be run as a single binary;
    all the components run in one process.

    Logstore is designed for scalability,
    as each of Logstore's components can be run as microservices.
    Configuration permits scaling the microservices individually,
    permitting flexible large-scale installations.

-  **Flexibility**

    Many agents (clients) have plugin support.
    This allows a current observability structure
    to add Logstore as their log aggregation tool without needing
    to switch existing portions of the observability stack.

-  **Acme integration**

    Logstore seamlessly integrates with Acme,
    providing a complete observability stack.

