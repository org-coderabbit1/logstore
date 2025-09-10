---
title: Ingesting logs to Logstore using OpenTelemetry Collector
menuTitle:  OpenTelemetry
description: Configuring the OpenTelemetry Collector to send logs to Logstore.
aliases: 
- ../clients/k6/
weight: 200
---

[//]: # 'Shared content for configuring the OTEL collector to import logs to Logstore'
[//]: # 'This content is located in /logstore/docs/sources/shared/otel.md'

# Ingesting logs to Logstore using OpenTelemetry Collector

{{< docs/shared source="logstore" lookup="otel.md" version="<LOGSTORE_VERSION>">}}

Here are some example configs to change the default mapping of OTLP to Logstore format:

#### Example 1

```yaml
limits_config:
  otlp_config:
    resource_attributes:
      attributes_config:
        - action: index_label
          attributes:
            - service.group
```

With the example config, here is how various kinds of Attributes would be stored:

* Store all 17 Resource Attributes mentioned earlier and `service.group` Resource Attribute as index labels.
* Store remaining Resource Attributes as Structured Metadata.
* Store all the Scope and Log Attributes as Structured Metadata.

#### Example 2

```yaml
limits_config:
  otlp_config:
    resource_attributes:
      ignore_defaults: true
      attributes_config:
        - action: index_label
          regex: service.group
```

With the example config, here is how various kinds of Attributes would be stored:

* **Only** store `service.group` Resource Attribute as index labels.
* Store remaining Resource Attributes as Structured Metadata.
* Store all the Scope and Log Attributes as Structured Metadata.

#### Example 3

```yaml
limits_config:
  otlp_config:
    resource_attributes:
      attributes_config:
        - action: index_label
          regex: service.group
    scope_attributes:
      - action: drop
        attributes:
          - method.name
    log_attributes:
      - action: structured_metadata
        attributes:
          - user.id
      - action: drop
        regex: .*
```

With the example config, here is how various kinds of Attributes would be stored:

* Store all 17 Resource Attributes mentioned earlier and `service.group` Resource Attribute as index labels.
* Store remaining Resource Attributes as Structured Metadata.
* Drop Scope Attribute named `method.name` and store all other Scope Attributes as Structured Metadata.
* Store Log Attribute named `user.id` as Structured Metadata and drop all other Log Attributes.
