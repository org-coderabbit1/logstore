FROM golang:1.20.6 as build

COPY . /src/logstore
WORKDIR /src/logstore
RUN make clean && CGO_ENABLED=0 go build ./cmd/logql-analyzer/

FROM alpine:3.16.5

RUN apk add --no-cache ca-certificates

COPY --from=build /src/logstore/logql-analyzer /usr/bin/logql-analyzer

ENTRYPOINT [ "/usr/bin/logql-analyzer" ]
