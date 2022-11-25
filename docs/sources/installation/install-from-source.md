---
title: Build from source
weight: 50
---
# Build from source

Clone the Acme Logstore repository and use the provided `Makefile`
to build Logstore from source.

## Prerequisites

- [Go](https://golang.org/), version 1.14 or later;
set your `$GOPATH` environment variable
- `make`
- Docker (for updating protobuf and yacc files)

## Build locally

1. Clone Logstore to `$GOPATH/src/example.com/acme/logstore`:

    ```bash
    git clone https://example.com/acme/logstore $GOPATH/src/example.com/acme/logstore
    ```

2. With a current working directory of `$GOPATH/src/example.com/acme/logstore`:

    ```bash
    make logstore
    ```

The built executable will be in `$GOPATH/src/example.com/acme/logstore/cmd/logstore/logstore`.
