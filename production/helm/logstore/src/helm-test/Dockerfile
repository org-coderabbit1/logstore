FROM golang:1.20.4 as build

# build via Makefile target helm-test-image in root
# Makefile. Building from this directory will not be
# able to access source needed in rest of repo.
COPY . /src/logstore
WORKDIR /src/logstore
RUN make clean && make BUILD_IN_CONTAINER=false helm-test

FROM alpine:3.16.5
RUN apk add --update --no-cache ca-certificates=20220614-r0
COPY --from=build /src/logstore/production/helm/logstore/src/helm-test/helm-test /usr/bin/helm-test
ENTRYPOINT [ "/usr/bin/helm-test" ]
