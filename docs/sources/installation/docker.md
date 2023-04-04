---
title: Install Acme Logstore with Docker or Docker Compose
description: Docker
weight: 30
---
# Install Acme Logstore with Docker or Docker Compose

You can install Acme Logstore and Promtail with Docker or Docker Compose if you are evaluating, testing, or developing Logstore.
For production, we recommend installing with Tanka or Helm.

The configuration acquired with these installation instructions run Logstore as a single binary.

## Prerequisites

- [Docker](https://docs.docker.com/install)
- [Docker Compose](https://docs.docker.com/compose/install) (optional, only needed for the Docker Compose install method)

## Install with Docker

**Linux**

Copy and paste the commands below into your command line.

```bash
wget https://raw.githubusercontent.com/acme/logstore/v2.8.0/cmd/logstore/logstore-local-config.yaml -O logstore-config.yaml
docker run --name logstore -d -v $(pwd):/mnt/config -p 3100:3100 acme/logstore:2.8.0 -config.file=/mnt/config/logstore-config.yaml
wget https://raw.githubusercontent.com/acme/logstore/v2.8.0/clients/cmd/promtail/promtail-docker-config.yaml -O promtail-config.yaml
docker run --name promtail -d -v $(pwd):/mnt/config -v /var/log:/var/log --link logstore acme/promtail:2.8.0 -config.file=/mnt/config/promtail-config.yaml
```

When finished, `logstore-config.yaml` and `promtail-config.yaml` are downloaded in the directory you chose. Docker containers are running Logstore and Promtail using those config files.

Navigate to http://localhost:3100/metrics to view the metrics and http://localhost:3100/ready for readiness.

The image is configured to run by default as user logstore with  UID `10001` and GID `10001`. You can use a different user, specially if you are using bind mounts, by specifying the UID with a `docker run` command and using `--user=UID` with numeric UID suited to your needs.

**Windows**

Copy and paste the commands below into your terminal. Note that you will need to replace the `<placeholders>` in the commands with your local path.

```bash
cd "<local-path>"
wget https://raw.githubusercontent.com/acme/logstore/v2.8.0/cmd/logstore/logstore-local-config.yaml -O logstore-config.yaml
docker run --name logstore -v <local-path>:/mnt/config -p 3100:3100 acme/logstore:2.8.0 --config.file=/mnt/config/logstore-config.yaml
wget https://raw.githubusercontent.com/acme/logstore/v2.8.0/clients/cmd/promtail/promtail-docker-config.yaml -O promtail-config.yaml
docker run -v <local-path>:/mnt/config -v /var/log:/var/log --link logstore acme/promtail:2.8.0 --config.file=/mnt/config/promtail-config.yaml
```

When finished, `logstore-config.yaml` and `promtail-config.yaml` are downloaded in the directory you chose. Docker containers are running Logstore and Promtail using those config files.

Navigate to http://localhost:3100/metrics to view the output.

## Install with Docker Compose

Run the following commands in your command line. They work for Windows or Linux systems.

```bash
wget https://raw.githubusercontent.com/acme/logstore/v2.8.0/production/docker-compose.yaml -O docker-compose.yaml
docker-compose -f docker-compose.yaml up
```
