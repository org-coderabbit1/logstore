FROM golang:1.20.6 as build

COPY . /src/logstore
WORKDIR /src/logstore
RUN make clean && make BUILD_IN_CONTAINER=false logstore

FROM alpine:3.16.5

RUN apk add --no-cache ca-certificates libcap

COPY --from=build /src/logstore/cmd/logstore/logstore /usr/bin/logstore
COPY cmd/logstore/logstore-docker-config.yaml /etc/logstore/local-config.yaml

RUN addgroup -g 10001 -S logstore && \
    adduser -u 10001 -S logstore -G logstore
RUN mkdir -p /logstore/rules && \
    mkdir -p /logstore/rules-temp && \
    chown -R logstore:logstore /etc/logstore /logstore

USER 10001
EXPOSE 3100
ENTRYPOINT [ "/usr/bin/logstore" ]
CMD ["-config.file=/etc/logstore/local-config.yaml"]
