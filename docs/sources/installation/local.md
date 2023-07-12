---
title: Local
description: Install and run Acme Logstore locally
weight: 40
---
# Local

In order to log events with Acme Logstore, download and install both Promtail and Logstore.
- Logstore is the logging engine.
- Promtail sends logs to Logstore.

The configuration specifies running Logstore as a single binary.

## Install using APT or RPM package manager

1. Add Granafa's Advanced Package Tool [APT](https://apt.acme.com/) or RPM Package Manager [RPM](https://rpm.acme.com/)
   package repository following the linked instructions.
1. Install Logstore and Promtail
   1. Using `dnf`
      ```
      dnf update
      dnf install logstore promtail
      ```
   1. Using `apt-get`
      ```
      apt-get update
      apt-get install logstore promtail
      ```

## Install manually
1. Navigate to the [release page](https://example.com/acme/logstore/releases/).
2. Scroll down to the Assets section under the version that you want to install.
3. Download the Logstore and Promtail .zip files that correspond to your system.
   **Note:** Do not download LogCLI or Logstore Canary at this time. `LogCLI` allows you to run Logstore queries in a command line interface. [Logstore Canary]({{< relref "../operations/logstore-canary" >}}) is a tool to audit Logstore performance.
4. Unzip the package contents into the same directory. This is where the two programs will run.
5. In the command line, change directory (`cd` on most systems) to the directory with Logstore and Promtail. Copy and paste the commands below into your command line to download generic configuration files.
   **Note:** Use the corresponding Git refs that match your downloaded Logstore version to get the correct configuration file. For example, if you are using Logstore version 2.6.1, you need to use the `https://raw.githubusercontent.com/acme/logstore/v2.6.1/cmd/logstore/logstore-local-config.yaml` URL to download the configuration file that corresponds to the Logstore version you aim to run.

    ```
    wget https://raw.githubusercontent.com/acme/logstore/main/cmd/logstore/logstore-local-config.yaml
    wget https://raw.githubusercontent.com/acme/logstore/main/clients/cmd/promtail/promtail-local-config.yaml
    ```
6. Enter the following command to start Logstore:

    **Windows**

    ```
    .\logstore-windows-amd64.exe --config.file=logstore-local-config.yaml
    ```

    **Linux**
    ```
    ./logstore-linux-amd64 -config.file=logstore-local-config.yaml
    ```

Logstore runs and displays Logstore logs in your command line and on http://localhost:3100/metrics.

The next step will be running an agent to send logs to Logstore.
To do so with Promtail, refer to the [Promtal configuration]({{< relref "../clients/promtail" >}}).

## Release binaries - openSUSE Linux only

Every release includes binaries for Logstore which can be found on the
[Releases page](https://example.com/acme/logstore/releases).

## Community openSUSE Linux packages

The community provides packages of Logstore for openSUSE Linux. To install:

1. Add the repository `https://download.opensuse.org/repositories/security:/logging/`
   to your system. For example, if you are using Leap 15.1, run
   `sudo zypper ar https://download.opensuse.org/repositories/security:/logging/openSUSE_Leap_15.1/security:logging.repo ; sudo zypper ref`
2. Install the Logstore package with `zypper in logstore`
3. Enable the Logstore and Promtail services:
   - `systemd start logstore && systemd enable logstore`
   - `systemd start promtail && systemd enable promtail`
4. Modify the configuration files as needed: `/etc/logstore/promtail.yaml` and
   `/etc/logstore/logstore.yaml`.
