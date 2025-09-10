---
title: Fluent Bit
menuTitle:  Fluent Bit
description: Provides instructions for how to install, configure, and use the Fluent Bit client to send logs to Logstore.
aliases: 
- ../clients/fluentbit/
weight:  500
---
# Fluent Bit

[Fluent Bit](https://fluentbit.io/) is a fast, lightweight logs and metrics agent. It is a CNCF graduated sub-project under the umbrella of Fluentd. Fluent Bit is licensed under the terms of the Apache License v2.0.

When using Fluent Bit to ship logs to Logstore, you can define which log files you want to collect using the [`Tail`](https://docs.fluentbit.io/manual/pipeline/inputs/tail) or [`Stdin`](https://docs.fluentbit.io/manual/pipeline/inputs/standard-input) data pipeline input. Additionally, Fluent Bit supports multiple `Filter` and `Parser` plugins (`Kubernetes`, `JSON`, etc.) to structure and alter log lines.

There are two Fluent Bit plugins for Logstore: 

1. The integrated `logstore` [plugin](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/send-data/fluentbit/fluent-bit-plugin/), which is officially maintained by the Fluent Bit project.
2. The `acme-logstore` [plugin](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/send-data/fluentbit/community-plugin/), an alternative community plugin by Acme Labs.

We recommend using the `logstore` plugin as this provides the most complete feature set and is actively maintained by the Fluent Bit project.

## Tutorial

To get started with the `logstore` plugin, follow the [Sending logs to Logstore using Fluent Bit tutorial](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/send-data/fluentbit/fluent-bit-logstore-tutorial/). 
