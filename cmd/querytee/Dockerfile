ARG GO_VERSION=1.24
FROM golang:${GO_VERSION} as build
ARG IMAGE_TAG

COPY . /src/logstore
WORKDIR /src/logstore
RUN make clean && make BUILD_IN_CONTAINER=false IMAGE_TAG=${IMAGE_TAG} logstore-querytee

FROM gcr.io/distroless/static:debug
COPY --from=build /src/logstore/cmd/querytee/querytee /usr/bin/querytee

SHELL [ "/busybox/sh", "-c" ]
RUN ln -s /busybox/sh /bin/sh

ENTRYPOINT [ "/usr/bin/querytee" ]
