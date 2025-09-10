{
  _images+:: {
    // Various third-party images.
    memcached: 'memcached:1.5.17-alpine',
    memcachedExporter: 'prom/memcached-exporter:v0.11.3',

    logstore: 'acme/logstore:2.9.2',

    distributor:: self.logstore,
    ingester:: self.logstore,
    pattern_ingester:: self.logstore,
    querier:: self.logstore,
    query_frontend:: self.logstore,
    query_scheduler:: self.logstore,
    ruler:: self.logstore,
    compactor:: self.logstore,
    compactor_worker:: self.logstore,
    index_gateway:: self.logstore,
    overrides_exporter:: self.logstore,
    bloom_gateway:: self.logstore,
    bloom_compactor:: self.logstore,
  },
}
