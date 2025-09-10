ARG GO_VERSION=1.24
FROM golang:${GO_VERSION} as build
ARG IMAGE_TAG

COPY . /src/logstore
WORKDIR /src/logstore
RUN go env GOARCH > /goarch
RUN make clean && make GOARCH=$(cat /goarch) BUILD_IN_CONTAINER=true GOEXPERIMENT=boringcrypto IMAGE_TAG=${IMAGE_TAG} logstore-canary-boringcrypto

FROM gcr.io/distroless/base-nossl:debug
COPY --from=build /src/logstore/cmd/logstore-canary-boringcrypto/logstore-canary-boringcrypto /usr/bin/logstore-canary
SHELL [ "/busybox/sh", "-c" ]
RUN ln -s /busybox/sh /bin/sh
ENTRYPOINT [ "/usr/bin/logstore-canary" ]
