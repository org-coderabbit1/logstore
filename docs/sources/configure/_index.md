---
title: Acme Logstore configuration parameters
menuTitle: Configure
description: Configuration reference for the parameters used to configure Acme Logstore.
aliases:
  - ./configuration # /docs/logstore/<LOGSTORE_VERSION>/configuration/
weight: 400
---

# Acme Logstore configuration parameters

Acme Logstore is configured in a YAML file (usually referred to as `logstore.yaml` )
which contains information on the Logstore server and its individual components,
depending on which mode Logstore is launched in.

Configuration examples can be found in the [Configuration Examples]({{< relref "./examples/configuration-examples" >}}) document.

<!-- The shared `configuration.md` file is generated from `/docs/templates/configuration.template`. To make changes to the included content, modify the template file and run `make doc` from root directory to regenerate the shared file. -->

{{< docs/shared lookup="configuration.md" source="logstore" version="<LOGSTORE_VERSION>" >}}
