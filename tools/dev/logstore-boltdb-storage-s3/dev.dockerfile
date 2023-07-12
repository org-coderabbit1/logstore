FROM golang:1.20.4
ENV CGO_ENABLED=0
RUN go install github.com/go-delve/delve/cmd/dlv@v1.20.2

FROM alpine:3.16.4

RUN     mkdir /logstore
WORKDIR /logstore
ADD     ./logstore ./
ADD     ./.src ./src
COPY --from=0 /go/bin/dlv ./
