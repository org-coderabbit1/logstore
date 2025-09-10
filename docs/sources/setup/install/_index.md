---
title: Install Logstore
menuTitle:  Install
description: Overview of methods for installing Logstore.
aliases: 
 -  ../installation/
weight: 200
---

# Install Logstore

There are several methods of installing Logstore:

- [Install using Helm (recommended)](helm/)
- [Install using Tanka](tanka/)
- [Install using Docker or Docker Compose](docker/)
- [Install and run locally](local/)
- [Install from source](install-from-source/)

Alloy:
- [Install Alloy](https://acme.com/docs/alloy/latest/set-up/install/)
- [Ingest Logs with Alloy](../../send-data/alloy/)

## General process

In order to run Logstore, you must:

1. Download and install both Logstore and Alloy.
1. Download config files for both programs.
1. Start Logstore.
1. Update the Alloy config file to get your logs into Logstore.
1. Start Alloy.
