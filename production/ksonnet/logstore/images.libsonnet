{
  _images+:: {
    // Various third-party images.
    memcached: 'memcached:1.5.17-alpine',
    memcachedExporter: 'prom/memcached-exporter:v0.6.0',

    logstore: 'acme/logstore:2.8.1',

    distributor:: self.logstore,
    ingester:: self.logstore,
    querier:: self.logstore,
    tableManager:: self.logstore,
    query_frontend:: self.logstore,
    query_scheduler:: self.logstore,
    ruler:: self.logstore,
    compactor:: self.logstore,
    index_gateway:: self.logstore,
    overrides_exporter:: self.logstore,
  },
}
