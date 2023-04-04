FROM golang:1.20.1 as build

COPY . /src/logstore
WORKDIR /src/logstore
RUN make clean && make BUILD_IN_CONTAINER=false logcli

FROM alpine:3.16.4

RUN apk add --no-cache ca-certificates

COPY --from=build /src/logstore/cmd/logcli/logcli /usr/bin/logcli

ENTRYPOINT [ "/usr/bin/logcli" ]
