FROM golang:1.24
ENV CGO_ENABLED=0
RUN go install github.com/go-delve/delve/cmd/dlv@v1.24.2

FROM alpine:3.22.1

RUN     mkdir /logstore
WORKDIR /logstore
ADD     ./logstore ./
ADD     ./.src ./src
COPY --from=0 /go/bin/dlv ./
