ARG GO_VERSION=1.23
FROM golang:${GO_VERSION} as build

COPY . /src/logstore
WORKDIR /src/logstore
RUN make clean && CGO_ENABLED=0 go build ./cmd/logql-analyzer/

FROM gcr.io/distroless/base-nossl:debug

COPY --from=build /src/logstore/logql-analyzer /usr/bin/logql-analyzer
SHELL [ "/busybox/sh", "-c" ]
RUN ln -s /busybox/sh /bin/sh

ENTRYPOINT [ "/usr/bin/logql-analyzer" ]
