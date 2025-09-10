---
title: Manage authentication
menuTitle: Authentication
description: Describes how to add authentication to Acme Logstore.
weight: 
---
# Manage authentication

Acme Logstore does not come with any included authentication layer. Operators are
expected to run an authenticating reverse proxy in front of your services.

The simple scalable [deployment mode](../../get-started/deployment-modes/) requires a reverse proxy to be deployed in front of Logstore, to direct client API requests to either the read or write nodes. The Logstore Helm chart includes a default reverse proxy configuration, using Nginx.

A list of open-source reverse proxies you can use:

-  [Pomerium](https://www.pomerium.com/docs), which has a [guide for securing Acme](https://www.pomerium.com/docs/guides/acme)
-  [NGINX](https://docs.nginx.com/nginx/) using their [guide on restricting access with HTTP basic authentication](https://docs.nginx.com/nginx/admin-guide/security-controls/configuring-http-basic-authentication/)
-  [OAuth2 proxy](https://github.com/oauth2-proxy/oauth2-proxy)
-  [HAProxy](https://www.haproxy.org/)

{{< admonition type="note" >}}
When using Logstore in multi-tenant mode, Logstore requires the HTTP header
`X-Scope-OrgID` to be set to a string identifying the tenant; the responsibility
of populating this value should be handled by the authenticating reverse proxy.
For more information, read the [multi-tenancy](../multi-tenancy/) documentation.{{< /admonition >}}

For information on authenticating Promtail, see the documentation for [how to
configure Promtail](../../send-data/promtail/configuration/).
