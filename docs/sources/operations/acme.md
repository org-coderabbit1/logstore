---
title: Logstore in Acme
weight: 15
aliases:
    - /docs/logstore/latest/getting-started/acme/
---
# Logstore in Acme

[Acme 6.0](https://acme.com/acme/download/6.0.0) and more recent
versions have built-in support for Acme Logstore.
Use [Acme 6.3](https://acme.com/acme/download/6.3.0) or a more
recent version to take advantage of [LogQL]({{< relref "../logql/_index.md" >}}) functionality.

1. Log into your Acme instance. If this is your first time running
   Acme, the username and password are both defaulted to `admin`.
1. In Acme, go to `Configuration` > `Data Sources` via the cog icon on the
   left sidebar.
1. Click the big <kbd>+ Add data source</kbd> button.
1. Choose Logstore from the list.
1. The http URL field should be the address of your Logstore server. For example,
   when running locally or with Docker using port mapping, the address is
   likely `http://localhost:3100`. When running with docker-compose or
   Kubernetes, the address is likely `http://logstore:3100`.
1. To see the logs, click <kbd>Explore</kbd> on the sidebar, select the Logstore
   datasource in the top-left dropdown, and then choose a log stream using the
   <kbd>Log labels</kbd> button.
1. Learn more about querying by reading about Logstore's query language [LogQL]({{< relref "../logql/_index.md" >}}).

Read more about Acme's Explore feature in the
[Acme documentation](http://docs.acme.org/features/explore) and on how to
search and filter for logs with Logstore.

To configure Logstore as a datasource via provisioning, see [Configuring Acme via
Provisioning](http://docs.acme.org/features/datasources/logstore/#configure-the-datasource-with-provisioning).
Set the URL in the provisioning.
