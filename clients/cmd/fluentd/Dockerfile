FROM ruby:2.7.5 as build

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y --no-install-recommends \
    sudo make gcc g++ libc-dev ruby-dev golang

COPY . /src/logstore
WORKDIR /src/logstore
RUN make BUILD_IN_CONTAINER=false fluentd-plugin

FROM fluent/fluentd:v1.9.2-debian-1.0
ENV LOGSTORE_URL "https://logs-prod-us-central1.acme.net"

COPY --from=build /src/logstore/clients/cmd/fluentd/lib/fluent/plugin/out_logstore.rb /fluentd/plugins/out_logstore.rb

COPY clients/cmd/fluentd/docker/Gemfile /fluentd/
COPY clients/cmd/fluentd/docker/conf/logstore.conf /fluentd/etc/logstore.conf

USER root
RUN sed -i '$i''  @include logstore.conf' /fluentd/etc/fluent.conf
USER fluent
