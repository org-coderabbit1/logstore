ARG BUILD_IMAGE=acme/logstore-build-image:0.34.6
ARG GOARCH=amd64
# Directories in this file are referenced from the root of the project not this folder
# This file is intended to be called from the root like so:
# docker build -t acme/logstore-docker-driver -f clients/cmd/docker-driver/Dockerfile .

FROM $BUILD_IMAGE AS build
COPY . /src/logstore
WORKDIR /src/logstore

ARG GOARCH
RUN make clean && make BUILD_IN_CONTAINER=false GOARCH=${GOARCH} clients/cmd/docker-driver/docker-driver

FROM alpine:3.22.1 AS temp

ARG GOARCH

RUN apk add --update --no-cache --arch=${GOARCH} ca-certificates tzdata

FROM --platform=linux/${GOARCH} alpine:3.22.1

COPY --from=temp /etc/ca-certificates.conf /etc/ca-certificates.conf
COPY --from=temp /usr/share/ca-certificates /usr/share/ca-certificates
COPY --from=temp /usr/share/zoneinfo /usr/share/zoneinfo

COPY --from=build /src/logstore/clients/cmd/docker-driver/docker-driver /bin/docker-driver

WORKDIR /bin/
ENTRYPOINT [ "/bin/docker-driver" ]
