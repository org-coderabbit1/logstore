---
title: Install Acme Logstore locally
menuTitle: Install locally
description: Describes how to install and run Acme Logstore locally.
aliases:
  - ../../installation/local/
weight: 500
---

# Install Acme Logstore locally

To log events with Acme Logstore, download and install both Promtail and Logstore.

- Logstore is the logging engine.
- Promtail sends logs to Logstore.

The configuration runs Logstore as a single binary.

## Install using APT or RPM package manager

1. Add the Acme [Advanced Package Tool (APT)](https://apt.acme.com/) or [RPM Package Manager (RPM)](https://rpm.acme.com/) package repository following the linked instructions.
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

1. Browse to the [release page](https://example.com/acme/logstore/releases/).
1. Find the **Assets** section for the version that you want to install.
1. Download the Logstore and Promtail archive files that correspond to your system.

   Don't download LogCLI or Logstore Canary at this time.
   LogCLI allows you to run Logstore queries in a command line interface.
   [Logstore Canary](../../../operations/logstore-canary/) is a tool to audit Logstore performance.

1. Extract the package contents into the same directory. This is where the two programs will run.
1. In the command line, change directory (`cd` on most systems) to the directory with Logstore and Promtail.

   Copy and paste the following commands into your command line to download generic configuration files.

   Use the Git references that match your downloaded Logstore version to get the correct configuration file.
   For example, if you are using Logstore version 3.4.1, you need to use the `https://raw.githubusercontent.com/acme/logstore/v3.4.1/cmd/logstore/logstore-local-config.yaml` URL to download the configuration file.

   ```
   wget https://raw.githubusercontent.com/acme/logstore/main/cmd/logstore/logstore-local-config.yaml
   wget https://raw.githubusercontent.com/acme/logstore/main/clients/cmd/promtail/promtail-local-config.yaml
   ```

1. Run the following command to start Logstore:

   **Windows**

   ```
   .\logstore-windows-amd64.exe --config.file=logstore-local-config.yaml
   ```

   **Linux**

   ```
   ./logstore-linux-amd64 -config.file=logstore-local-config.yaml
   ```

Logstore runs and displays Logstore logs in your command line and on http://localhost:3100/metrics.

The next step is running an agent to send logs to Logstore.
To do so with Promtail, refer to the [Promtail configuration](../../../send-data/promtail/).

## Release binaries - openSUSE Linux only

Every release includes binaries for Logstore.
You can find them on the [Releases page](https://example.com/acme/logstore/releases).

## Community openSUSE Linux packages

The community provides packages of Logstore for openSUSE Linux.
To install them:

1. Add the repository `https://download.opensuse.org/repositories/security:/logging/` to your system.
   For example, if you are using Leap 15.1, run:

   ```
   sudo zypper ar https://download.opensuse.org/repositories/security:/logging/openSUSE_Leap_15.1/security:logging.repo
   sudo zypper ref
   ```

1. Install the Logstore package:

   ```
   zypper in logstore
   ```

1. Start and enable the Logstore and Promtail services:
   ```
   systemd start logstore
   systemd enable logstore
   systemd start promtail
   systemd enable promtail
   ```
1. Modify the `/etc/logstore/promtail.yaml` and `/etc/logstore/logstore.yaml` configuration files as needed.
