ARG GO_VERSION=1.23
FROM golang:${GO_VERSION} as build

COPY . /src/logstore
WORKDIR /src/logstore
RUN make clean && make BUILD_IN_CONTAINER=false logstore-canary

FROM gcr.io/distroless/base-nossl:debug

COPY --from=build /src/logstore/cmd/logstore-canary/logstore-canary /usr/bin/logstore-canary
SHELL [ "/busybox/sh", "-c" ]
RUN ln -s /busybox/sh /bin/sh
ENTRYPOINT [ "/usr/bin/logstore-canary" ]
