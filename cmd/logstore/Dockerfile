ARG GO_VERSION=1.23
FROM golang:${GO_VERSION} as build

COPY . /src/logstore
WORKDIR /src/logstore
RUN make clean && make BUILD_IN_CONTAINER=false logstore

FROM gcr.io/distroless/base-nossl:debug

COPY --from=build /src/logstore/cmd/logstore/logstore /usr/bin/logstore
COPY cmd/logstore/logstore-docker-config.yaml /etc/logstore/local-config.yaml

SHELL [ "/busybox/sh", "-c" ]

RUN addgroup -g 10001 -S logstore && \
    adduser -u 10001 -S logstore -G logstore
RUN mkdir -p /logstore/rules && \
    mkdir -p /logstore/rules-temp && \
    chown -R logstore:logstore /etc/logstore /logstore && \
    ln -s /busybox/sh /bin/sh
USER 10001
EXPOSE 3100
ENTRYPOINT [ "/usr/bin/logstore" ]
CMD ["-config.file=/etc/logstore/local-config.yaml"]
