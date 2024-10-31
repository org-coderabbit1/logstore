---
title: Contributing to Logstore
description: Contributing to Logstore
weight: 200
---

# Contributing to Logstore

Logstore uses [GitHub](https://example.com/acme/logstore) to manage reviews of pull requests:

- If you have a trivial fix or improvement, go ahead and create a pull request.
- If you plan to do something more involved, discuss your ideas on the relevant GitHub issue (creating one if it doesn't exist).

## Steps to contribute

To contribute to Logstore, you must clone it into your `$GOPATH` and add your fork
as a remote.

```bash
$ git clone https://example.com/acme/logstore.git $GOPATH/src/example.com/acme/logstore
$ cd $GOPATH/src/example.com/acme/logstore
$ git remote add fork <FORK_URL>
```

Make your changes, add your changes to a commit, and open a pull request (PR).

```bash
$ git add .
$ git commit -m "docs: fix spelling error"
$ git push -u fork HEAD
```

{{% admonition type="note" %}}
If you downloaded Logstore using `go get`, the message `package example.com/acme/logstore: no Go files in /go/src/example.com/acme/logstore`
is normal and requires no actions to resolve.
{{% /admonition %}}

### Building

While `go install ./cmd/logstore` works, the preferred way to build is by using
`make`:

- `make logstore`: builds Logstore and outputs the binary to `./cmd/logstore/logstore`

- `make promtail`: builds Promtail and outputs the binary to
  `./clients/cmd/promtail/promtail`

- `make logcli`: builds LogCLI and outputs the binary to `./cmd/logcli/logcli`

- `make logstore-canary`: builds Logstore Canary and outputs the binary to
  `./cmd/logstore-canary/logstore-canary`

- `make docker-driver`: builds the Logstore Docker Driver and installs it into
  Docker.

- `make images`: builds all Docker images (optionally suffix the previous binary
  commands with `-image`, e.g., `make logstore-image`).

These commands can be chained together to build multiple binaries in one go.  The following example builds binaries for Logstore, Promtail, and LogCLI.

```bash
$ make logstore promtail logcli
```

## Contribute to the Helm Chart

The official Logstore helm charts can be found in the [Acme Helm Charts Repo](https://example.com/acme/helm-charts).
