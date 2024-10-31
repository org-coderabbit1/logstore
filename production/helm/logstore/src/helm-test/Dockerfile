ARG GO_VERSION=1.23
FROM golang:${GO_VERSION} as build

# build via Makefile target helm-test-image in root
# Makefile. Building from this directory will not be
# able to access source needed in rest of repo.
COPY . /src/logstore
WORKDIR /src/logstore
RUN make clean && make BUILD_IN_CONTAINER=false helm-test

FROM gcr.io/distroless/base-nossl:debug
COPY --from=build /src/logstore/production/helm/logstore/src/helm-test/helm-test /usr/bin/helm-test
ENTRYPOINT [ "/usr/bin/helm-test" ]
