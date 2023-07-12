ARG BUILD_IMAGE=acme/logstore-build-image:0.29.3
# Directories in this file are referenced from the root of the project not this folder
# This file is intended to be called from the root like so:
# docker build -t acme/logstore -f cmd/logstore/Dockerfile .

# TODO: add cross-platform support
FROM $BUILD_IMAGE as build
COPY . /src/logstore
WORKDIR /src/logstore
RUN make clean && make BUILD_IN_CONTAINER=false clients/cmd/docker-driver/docker-driver

FROM alpine:3.16.5
RUN apk add --update --no-cache ca-certificates tzdata
COPY --from=build /src/logstore/clients/cmd/docker-driver/docker-driver /bin/docker-driver
WORKDIR /bin/
ENTRYPOINT [ "/bin/docker-driver" ]
