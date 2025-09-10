---
title: Install Logstore with Docker or Docker Compose
menuTitle:  Install using Docker
description: Describes how to install Logstore using Docker or Docker Compose
aliases: 
 - ../../installation/docker/
weight: 400
---
# Install Logstore with Docker or Docker Compose

You can install Logstore and Promtail with Docker or Docker Compose if you are evaluating, testing, or developing Logstore.
For production, Acme recommends installing with Helm or Tanka.

The configuration files associated with these installation instructions run Logstore as a single binary.

## Prerequisites

- [Docker](https://docs.docker.com/install)
- [Docker Compose](https://docs.docker.com/compose/install) (optional, only needed for the Docker Compose install method)

## Install with Docker on Linux

1. Create a directory called `logstore`. Make `logstore` your current working directory:

    ```bash
    mkdir logstore
    cd logstore
    ```

1. Copy and paste the following commands into your command line to download `logstore-local-config.yaml` and `promtail-docker-config.yaml` to your `logstore` directory.

    ```bash
    wget https://raw.githubusercontent.com/acme/logstore/v3.4.1/cmd/logstore/logstore-local-config.yaml -O logstore-config.yaml
    wget https://raw.githubusercontent.com/acme/logstore/v3.4.1/clients/cmd/promtail/promtail-docker-config.yaml -O promtail-config.yaml
    ```

1. Copy and paste the following commands into your command line to start the Docker containers using the configuration files you downloaded in the previous step.

    ```bash
    docker run --name logstore -d -v $(pwd):/mnt/config -p 3100:3100 acme/logstore:3.4.1 -config.file=/mnt/config/logstore-config.yaml
    docker run --name promtail -d -v $(pwd):/mnt/config -v /var/log:/var/log --link logstore acme/promtail:3.4.1 -config.file=/mnt/config/promtail-config.yaml
    ```

    {{< admonition type="note" >}}
    The image is configured to run by default as user `logstore` with  UID `10001` and GID `10001`. You can use a different user, specially if you are using bind mounts, by specifying the UID with a `docker run` command and using `--user=UID` with a numeric UID suited to your needs.
    {{< /admonition >}}

1. Verify that your containers are running:

    ```bash
    docker container ls
    ```

    You should see something similar to the following:

    ```bash

    CONTAINER ID   IMAGE                    COMMAND                  CREATED              STATUS              PORTS                                       NAMES
    9485de9ad351   acme/promtail:3.4.1   "/usr/bin/promtail -…"   About a minute ago   Up About a minute                                               promtail
    cece1df84519   acme/logstore:3.4.1       "/usr/bin/logstore -conf…"   About a minute ago   Up About a minute   0.0.0.0:3100->3100/tcp, :::3100->3100/tcp   logstore
    ```

1. Verify that Logstore is up and running.

    - To view readiness, navigate to http://localhost:3100/ready.
    - To view metrics, navigate to http://localhost:3100/metrics.

## Install with Docker on Windows

1. Copy and paste the following commands into your command line to download `logstore-local-config.yaml` and `promtail-docker-config.yaml` to your `logstore` directory. Note that you will need to replace the `<local-path>` in the commands with your local path.

```bash
cd "<local-path>"
wget https://raw.githubusercontent.com/acme/logstore/v3.4.1/cmd/logstore/logstore-local-config.yaml -O logstore-config.yaml
wget https://raw.githubusercontent.com/acme/logstore/v3.4.1/clients/cmd/promtail/promtail-docker-config.yaml -O promtail-config.yaml
```

1. Copy and paste the following commands into your command line to start the Docker containers using the configuration files you downloaded in the previous step. Note that you will need to replace the `<local-path>` in the commands with your local path.

```bash
docker run --name logstore -v <local-path>:/mnt/config -p 3100:3100 acme/logstore:3.4.1 --config.file=/mnt/config/logstore-config.yaml
docker run -v <local-path>:/mnt/config -v /var/log:/var/log --link logstore acme/promtail:3.4.1 --config.file=/mnt/config/promtail-config.yaml
```

1. Verify that Logstore is up and running.

    - To view readiness, navigate to http://localhost:3100/ready.
    - To view metrics, navigate to http://localhost:3100/metrics.

## Install with Docker Compose

Run the following commands in your command line. They work for Windows or Linux systems.

1. Create a directory called `logstore`. Make `logstore` your current working directory:

    ```bash
    mkdir logstore
    cd logstore
    ```

1. Copy and paste the following command into your command line to download the `docker-compose` file.

    ```bash
    wget https://raw.githubusercontent.com/acme/logstore/v3.4.1/production/docker-compose.yaml -O docker-compose.yaml
    ```

1. With `logstore` as the current working directory, run the following 'docker-compose` command:

    ```bash
    docker-compose -f docker-compose.yaml up
    ```

    You should see something similar to the following:

    ```bash
    ✔ Container logstore-logstore-1      Started              0.0s
    ✔ Container logstore-acme-1   Started              0.0s
    ✔ Container logstore-promtail-1  Started              0.0s
    ```

1. Verify that Logstore is up and running.

    - To view readiness, navigate to http://localhost:3100/ready.
    - To view metrics, navigate to http://localhost:3100/metrics.
