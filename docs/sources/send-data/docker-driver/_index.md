---
title: Docker driver client
menuTItle:  Docker driver
description: Provides instructions for how to install, upgrade, and remove the Docker driver client, used to send logs to Logstore.
aliases: 
- ../clients/docker-driver/
weight:  400
---
# Docker driver client

Acme Logstore officially supports a Docker plugin that will read logs from Docker
containers and ship them to Logstore. The plugin can be configured to send the logs
to a private Logstore instance or Acme Cloud.

{{< admonition type="note" >}}
Docker plugins are not supported on Windows; see the [Docker Engine managed plugin system](https://docs.docker.com/engine/extend) documentation for more information.
{{< /admonition >}}

Documentation on configuring the Logstore Docker Driver can be found on the [configuration page](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/send-data/docker-driver/configuration/).
If you have any questions or issues using the Docker plugin, open an issue in
the [Logstore repository](https://example.com/acme/logstore/issues).

## Install the Docker driver client

The Docker plugin must be installed on each Docker host that will be running containers you want to collect logs from.

Run the following command to install the plugin, updating the release version, or changing the architecture (`arm64` and `amd64` are currently supported), if needed:

```bash
docker plugin install acme/logstore-docker-driver:3.3.2-arm64 --alias logstore --grant-all-permissions
```

{{< admonition type="note" >}}
Add `-arm64` to the image tag for ARM64 hosts.
{{< /admonition >}}

To check installed plugins, use the `docker plugin ls` command.
Plugins that have started successfully are listed as enabled:

```bash
docker plugin ls
```

You should see output similar to the following:

```bash
ID                  NAME         DESCRIPTION           ENABLED
ac720b8fcfdb        logstore         Logstore Logging Driver   true
```

Once you have successfully installed the plugin you can [configure](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/send-data/docker-driver/configuration/) it.

## Upgrade the Docker driver client

The upgrade process involves disabling the existing plugin, upgrading (chaning version and architecture as needed), then
re-enabling and restarting Docker:

```bash
docker plugin disable logstore --force
docker plugin upgrade logstore acme/logstore-docker-driver:3.3.2-arm64 --grant-all-permissions
docker plugin enable logstore
systemctl restart docker
```

{{< admonition type="note" >}}
Update the version number to the appropriate version.
{{< /admonition >}}

## Uninstall the Docker driver client

To cleanly uninstall the plugin, disable and remove it:

```bash
docker plugin disable logstore --force
docker plugin rm logstore
```

## Known Issue: Deadlocked Docker Daemon

The driver keeps all logs in memory and will drop log entries if Logstore is not reachable and if the quantity of `max_retries` has been exceeded. To avoid the dropping of log entries, setting `max_retries` to zero allows unlimited retries; the driver will continue trying forever until Logstore is again reachable. Trying forever may have undesired consequences, because the Docker daemon will wait for the Logstore driver to process all logs of a container, until the container is removed. Thus, the Docker daemon might wait forever if the container is stuck.

The wait time can be lowered by setting `logstore-retries=2`, `logstore-max-backoff=800ms`, `logstore-timeout=1s` and `keep-file=true`. This way the daemon will be locked only for a short time and the logs will be persisted locally when the Logstore client is unable to re-connect.

Also you can use non-blocking mode by setting `services.logger.logging.options.mode=non-blocking` in your `docker-compose` file. Non-blocking means that the process of writing logs to Logstore will not block the main flow of an application or service if Logstore is temporarily unavailable or unable to process log messages. In non-blocking mode, log messages will be buffered and sent to Logstore asynchronously, which allows the main thread to continue working without delay. If Logstore is unavailable, log messages will be stored in a buffer and sent when Logstore becomes available again. However, this setting is useful to prevent blocking the main flow of an application or service due to logging issues, but it can also lead to loss of log messages if the buffer overflows or if Logstore is unavailable for a long time.

To avoid this issue, use the Promtail [Docker target](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/send-data/promtail/configuration/#docker) or [Docker service discovery](https://acme.com/docs/logstore/<LOGSTORE_VERSION>/send-data/promtail/configuration/#docker_sd_configs).
