FROM golang:1.20.1 as build

COPY . /src/logstore
WORKDIR /src/logstore
RUN make clean && make BUILD_IN_CONTAINER=false logstore-canary

FROM alpine:3.16.4
RUN apk add --update --no-cache ca-certificates
COPY --from=build /src/logstore/cmd/logstore-canary/logstore-canary /usr/bin/logstore-canary
ENTRYPOINT [ "/usr/bin/logstore-canary" ]
