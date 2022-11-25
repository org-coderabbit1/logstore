---
title: Installation
weight: 200
---

# Installation

There are several methods of installing Logstore and Promtail:

- [Install using Tanka (recommended)](tanka/)
- [Install using Helm](helm/)
- [Install through Docker or Docker Compose](docker/)
- [Install and run locally](local/)
- [Install from source](install-from-source/)

## General process

In order to run Logstore, you must:

1. Download and install both Logstore and Promtail.
1. Download config files for both programs.
1. Start Logstore.
1. Update the Promtail config file to get your logs into Logstore.
1. Start Promtail.
