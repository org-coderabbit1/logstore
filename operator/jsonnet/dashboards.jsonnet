local cfg = (import 'config.libsonnet');
local logstore = (import 'example.com/acme/logstore/production/logstore-mixin/mixin.libsonnet') + cfg.logstore;

{
  'acme-dashboard-logstorestack-chunks.json': logstore.acmeDashboards['logstore-chunks.json'],
  'acme-dashboard-logstorestack-reads.json': logstore.acmeDashboards['logstore-reads.json'],
  'acme-dashboard-logstorestack-writes.json': logstore.acmeDashboards['logstore-writes.json'],
  'acme-dashboard-logstorestack-retention.json': logstore.acmeDashboards['logstore-retention.json'],
  'acme-dashboard-logstorestack-rules.json': logstore.prometheusRules,
}
