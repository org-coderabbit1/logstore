## Main

## [0.6.2](https://example.com/acme/logstore/compare/operator/v0.6.1...operator/v0.6.2) (2024-09-11)


### Features

* Ingester Stream Limit Improvements ([#13532](https://example.com/acme/logstore/issues/13532)) ([ec34aaa](https://example.com/acme/logstore/commit/ec34aaa1ff2e616ef223631657b63f7dffedd3cc))
* **operator:** Add alert for discarded samples ([#13512](https://example.com/acme/logstore/issues/13512)) ([5f2a02f](https://example.com/acme/logstore/commit/5f2a02f14222dab891b7851e8f48052d6c9b594a))
* **operator:** Add support for Logstore OTLP limits config ([#13446](https://example.com/acme/logstore/issues/13446)) ([d02f435](https://example.com/acme/logstore/commit/d02f435d3bf121b19e15de4f139c95a6d010b25c))
* **operator:** Add support for the volume API ([#13369](https://example.com/acme/logstore/issues/13369)) ([d451e23](https://example.com/acme/logstore/commit/d451e23225047a11b4d5d82900cec4a46d6e7b39))
* **operator:** Enable leader-election ([#13760](https://example.com/acme/logstore/issues/13760)) ([1ba4bff](https://example.com/acme/logstore/commit/1ba4bff005930b173391df35248e6f58e076fa74))
* **operator:** Update Logstore operand to v3.1.0 ([#13422](https://example.com/acme/logstore/issues/13422)) ([cf5f52d](https://example.com/acme/logstore/commit/cf5f52dca0db93847218cdd2c3f4860d983381ae))
* **operator:** Update Logstore operand to v3.1.1 ([#14042](https://example.com/acme/logstore/issues/14042)) ([7ae1588](https://example.com/acme/logstore/commit/7ae1588200396b73a16fadd2610670a5ce5fd747))


### Bug Fixes

* **deps:** update k8s.io/utils digest to 702e33f ([#14033](https://example.com/acme/logstore/issues/14033)) ([b7eecc7](https://example.com/acme/logstore/commit/b7eecc7a693e96f4d0fe0dcd7583ecdc4dd7283f))
* **operator:** add alertmanager client config to ruler template ([#13182](https://example.com/acme/logstore/issues/13182)) ([6148c37](https://example.com/acme/logstore/commit/6148c3760d701768e442186d4e7d574c7dc16c91))
* **operator:** Allow structured metadata only if V13 schema provided ([#13463](https://example.com/acme/logstore/issues/13463)) ([3ac130b](https://example.com/acme/logstore/commit/3ac130b8a152169766cb173718f2312aeb4f694e))
* **operator:** Don't overwrite annotations for LogstoreStack ingress resources ([#13708](https://example.com/acme/logstore/issues/13708)) ([f523530](https://example.com/acme/logstore/commit/f52353060dd936cff587ff2060c8616941695ece))
* **operator:** Improve API documentation for schema version ([#13122](https://example.com/acme/logstore/issues/13122)) ([3a9f50f](https://example.com/acme/logstore/commit/3a9f50f5099a02e662b8ac10ddad0b36cd844161))
* **operator:** Remove duplicate conditions from status ([#13497](https://example.com/acme/logstore/issues/13497)) ([527510d](https://example.com/acme/logstore/commit/527510d1a84a981250047dbabba8d492177b8452))
* **operator:** Set object storage for delete requests when using retention ([#13562](https://example.com/acme/logstore/issues/13562)) ([46de4c1](https://example.com/acme/logstore/commit/46de4c1bc839ef682798bec5003123f7d5f4404b))
* **operator:** Skip updating annotations for serviceaccounts ([#13450](https://example.com/acme/logstore/issues/13450)) ([1b9b111](https://example.com/acme/logstore/commit/1b9b11116b48fb37b7015d27104668412fc04937))
* **operator:** Support v3.1.0 in OpenShift dashboards ([#13430](https://example.com/acme/logstore/issues/13430)) ([8279d59](https://example.com/acme/logstore/commit/8279d59f145df9c9132aeff9e3d46c738650027c))
* **operator:** Watch for CredentialsRequests on CCOAuthEnv only ([#13299](https://example.com/acme/logstore/issues/13299)) ([7fc926e](https://example.com/acme/logstore/commit/7fc926e36ea8fca7bd8e9955c8994574535dbbae))

## [0.6.1](https://example.com/acme/logstore/compare/operator/v0.6.0...operator/v0.6.1) (2024-06-03)


### Features

* prepare 3.0.0 release candidate ([#12348](https://example.com/acme/logstore/issues/12348)) ([664e569](https://example.com/acme/logstore/commit/664e569e14ef55a79cd77bdb49e9ffbe0c55bc37))


### Bug Fixes

* **operator:** Bump golang builder to 1.21.9 ([#12503](https://example.com/acme/logstore/issues/12503)) ([f680ee0](https://example.com/acme/logstore/commit/f680ee0453d1b7d315774591293927b988bca223))
* **operator:** Configure Logstore to use virtual-host-style URLs for S3 AWS endpoints ([#12469](https://example.com/acme/logstore/issues/12469)) ([0084262](https://example.com/acme/logstore/commit/0084262269f4e2cb94d04e0cc0d40e9666177f06))
* **operator:** Improve validation of provided S3 storage configuration ([#12181](https://example.com/acme/logstore/issues/12181)) ([f9350d6](https://example.com/acme/logstore/commit/f9350d6415d45c3cc2f9c0b4f7cd6f8f219907f2))
* **operator:** Use a minimum value for replay memory ceiling ([#13066](https://example.com/acme/logstore/issues/13066)) ([4f3ed77](https://example.com/acme/logstore/commit/4f3ed77cb92c2ffd605743237e609c28f7841728))
* update to build image 0.33.2, fixes bug with promtail windows DNS resolution ([#12732](https://example.com/acme/logstore/issues/12732)) ([759f42d](https://example.com/acme/logstore/commit/759f42dd50bb4896f5e568691ef32245bb8fb25a))
* updated all dockerfiles go1.22 ([#12708](https://example.com/acme/logstore/issues/12708)) ([71a8f2c](https://example.com/acme/logstore/commit/71a8f2c2b11b419bd8c0af1f859671e5d8730448))

## 0.6.0 (2024-03-19)

- [12228](https://example.com/acme/logstore/pull/12228) **xperimental**: Restructure LogstoreStack metrics
- [12164](https://example.com/acme/logstore/pull/12164) **periklis**: Use safe bearer token authentication to scrape operator metrics
- [12216](https://example.com/acme/logstore/pull/12216) **xperimental**: Fix duplicate operator metrics due to ServiceMonitor selector
- [12212](https://example.com/acme/logstore/pull/12212) **xperimental**: Keep credentialMode in status when updating schemas
- [12165](https://example.com/acme/logstore/pull/12165) **JoaoBraveCoding**: Change attribute value used for CCO-based credential mode
- [12157](https://example.com/acme/logstore/pull/12157) **periklis**: Fix managed auth features annotation for community-openshift bundle
- [12104](https://example.com/acme/logstore/pull/12104) **periklis**: Upgrade build and runtime dependencies
- [11928](https://example.com/acme/logstore/pull/11928) **periklis**: Fix remote write client timeout config rename
- [12097](https://example.com/acme/logstore/pull/12097) **btaani**: Fix encoding of blocked query pattern in configuration
- [12106](https://example.com/acme/logstore/pull/12106) **xperimental**: Allow setting explicit CredentialMode in LogstoreStack storage spec
- [11968](https://example.com/acme/logstore/pull/11968) **xperimental**: Extend status to show difference between running and ready
- [12007](https://example.com/acme/logstore/pull/12007) **xperimental**: Extend Azure secret validation
- [12008](https://example.com/acme/logstore/pull/12008) **xperimental**: Support using multiple buckets with AWS STS
- [11964](https://example.com/acme/logstore/pull/11964) **xperimental**: Provide Azure region for managed credentials using environment variable
- [11920](https://example.com/acme/logstore/pull/11920) **xperimental**: Refactor handling of credentials in managed-auth mode
- [11869](https://example.com/acme/logstore/pull/11869) **periklis**: Add support for running with Google Workload Identity
- [11868](https://example.com/acme/logstore/pull/11868) **xperimental**: Integrate support for OpenShift-managed credentials in Azure
- [11854](https://example.com/acme/logstore/pull/11854) **periklis**: Allow custom audience for managed-auth on STS
- [11802](https://example.com/acme/logstore/pull/11802) **xperimental**: Add support for running with Azure Workload Identity
- [11824](https://example.com/acme/logstore/pull/11824) **xperimental**: Improve messages for errors in storage secret
- [11524](https://example.com/acme/logstore/pull/11524) **JoaoBraveCoding**, **periklis**: Add OpenShift cloud credentials support for AWS STS
- [11513](https://example.com/acme/logstore/pull/11513) **btaani**: Add a custom metric that collects Logstorestacks requiring a schema upgrade
- [11718](https://example.com/acme/logstore/pull/11718) **periklis**: Upgrade k8s.io, sigs.k8s.io and openshift deps
- [11671](https://example.com/acme/logstore/pull/11671) **JoaoBraveCoding**: Update mixins to fix structured metadata dashboards
- [11624](https://example.com/acme/logstore/pull/11624) **xperimental**: React to changes in ConfigMap used for storage CA
- [11481](https://example.com/acme/logstore/pull/11481) **JoaoBraveCoding**: Adds AWS STS support
- [11533](https://example.com/acme/logstore/pull/11533) **periklis**: Add serviceaccount per LogstoreStack resource
- [11158](https://example.com/acme/logstore/pull/11158) **btaani**: operator: Add warning for old schema configuration
- [11473](https://example.com/acme/logstore/pull/11473) **JoaoBraveCoding**: Adds structured metadata dashboards
- [11448](https://example.com/acme/logstore/pull/11448) **periklis**: Update Logstore operand to v2.9.3
- [11357](https://example.com/acme/logstore/pull/11357) **periklis**: Fix storing authentication credentials in the Logstore ConfigMap
- [11393](https://example.com/acme/logstore/pull/11393) **periklis**: Add infra annotations for OpenShift based deployments
- [11094](https://example.com/acme/logstore/pull/11094) **periklis**: Add support for blocking queries per tenant
- [11288](https://example.com/acme/logstore/pull/11288) **periklis**: Fix custom CA for object-store in ruler component
- [11091](https://example.com/acme/logstore/pull/11091) **periklis**: Add automatic stream sharding support
- [11022](https://example.com/acme/logstore/pull/11022) **JoaoBraveCoding**: Remove outdated BoltDB dashboards 
- [10932](https://example.com/acme/logstore/pull/10932) **JoaoBraveCoding**: Adds new value v13 to schema
- [11232](https://example.com/acme/logstore/pull/11232) **periklis**: Update dependencies and dev tools
- [11129](https://example.com/acme/logstore/pull/11129) **periklis**: Update deps to secure webhooks for CVE-2023-44487

## 0.5.0 (2023-10-24)

- [10924](https://example.com/acme/logstore/pull/10924) **periklis**: Update Logstore operand to v2.9.2
- [10874](https://example.com/acme/logstore/pull/10874) **periklis**: Bump deps to address CVE-2023-39325 and CVE-2023-44487
- [10854](https://example.com/acme/logstore/pull/10854) **periklis**: Add missing marker/sweeper panels in retention dashboard
- [10717](https://example.com/acme/logstore/pull/10717) **periklis**: Allow SSE settings in AWS S3 object storage secret
- [10715](https://example.com/acme/logstore/pull/10715) **periklis**: Allow endpoint_suffix in azure object storage secret
- [10562](https://example.com/acme/logstore/pull/10562) **periklis**: Add memberlist IPv6 support
- [10720](https://example.com/acme/logstore/pull/10720) **JoaoBraveCoding**: Change default replication factor of 1x.medium to 2
- [10600](https://example.com/acme/logstore/pull/10600) **periklis**: Update Logstore operand to v2.9.1
- [10545](https://example.com/acme/logstore/pull/10545) **xperimental**: Update gateway arguments to enable namespace extraction
- [10558](https://example.com/acme/logstore/pull/10558) **periklis**: Upgrade dashboards for for Logstore v2.9.0
- [10539](https://example.com/acme/logstore/pull/10539) **periklis**: Update Logstore operand to v2.9.0
- [10418](https://example.com/acme/logstore/pull/10418) **btaani**: Use a condition to warn when labels for zone-awareness are empty
- [9468](https://example.com/acme/logstore/pull/9468) **periklis**: Add support for reconciling logstore-mixin dashboards on OpenShift Console
- [9942](https://example.com/acme/logstore/pull/9942) **btaani**: Use a condition to warn when there are no nodes with matching labels for zone-awareness

## 0.4.0 (2023-07-27)

- [10019](https://example.com/acme/logstore/pull/10019) **periklis**: Update Logstore operand to v2.8.3
- [9972](https://example.com/acme/logstore/pull/9972) **JoaoBraveCoding**: Fix OIDC.IssuerCAPath by updating it to type CASpec
- [9931](https://example.com/acme/logstore/pull/9931) **aminesnow**: Custom configuration for LogstoreStack admin groups
- [9971](https://example.com/acme/logstore/pull/9971) **aminesnow**: Add namespace and tenantId labels to RecordingRules
- [9906](https://example.com/acme/logstore/pull/9906) **JoaoBraveCoding**: Add mTLS authentication to tenants
- [9963](https://example.com/acme/logstore/pull/9963) **xperimental**: Fix application tenant alertmanager configuration
- [9795](https://example.com/acme/logstore/pull/9795) **JoaoBraveCoding**: Add initContainer to zone aware components to gatekeep them from starting without the AZ annotation
- [9503](https://example.com/acme/logstore/pull/9503) **shwetaap**: Add Pod annotations with node topology labels to support zone aware scheduling
- [9930](https://example.com/acme/logstore/pull/9930) **periklis**: Use PodAntiAffinity for all components
- [9860](https://example.com/acme/logstore/pull/9860) **xperimental**: Fix update of labels and annotations of PodTemplates
- [9830](https://example.com/acme/logstore/pull/9830) **periklis**: Expose limits config setting cardinality_limit
- [9600](https://example.com/acme/logstore/pull/9600) **periklis**: Add rules labels filters for openshift-logging application tenant
- [9735](https://example.com/acme/logstore/pull/9735) **JoaoBraveCoding** Adjust 1x.extra-small resources according to findings
- [9689](https://example.com/acme/logstore/pull/9689) **xperimental**: Fix availability of demo LogstoreStack size
- [9630](https://example.com/acme/logstore/pull/9630) **jpinsonneau**: Expose per_stream_rate_limit & burst
- [9623](https://example.com/acme/logstore/pull/9623) **periklis**: Fix timeout config constructor when only tenants limits
- [9457](https://example.com/acme/logstore/pull/9457) **Red-GV**: Set seccomp profile to runtime default
- [9448](https://example.com/acme/logstore/pull/9448) **btaani**: Include runtime-config in compiling the SHA1 checksum
- [9511](https://example.com/acme/logstore/pull/9511) **xperimental**: Do not update status after setting degraded condition
- [9405](https://example.com/acme/logstore/pull/9405) **periklis**: Add support for configuring HTTP server timeouts
- [9378](https://example.com/acme/logstore/pull/9378) **aminesnow**: Add zone aware API spec validation
- [9408](https://example.com/acme/logstore/pull/9408) **JoaoBraveCoding**: Add PodAntiAffinity overwrites per component
- [9429](https://example.com/acme/logstore/pull/9429) **aminesnow**: Add default TopologySpreadContraints to Gateway
- [9418](https://example.com/acme/logstore/pull/9418) **JoaoBraveCoding**: Add default TopologySpreadContaints to Querier
- [9383](https://example.com/acme/logstore/pull/9383) **JoaoBraveCoding**: Add default TopologySpreadContaints to Distributor
- [9406](https://example.com/acme/logstore/pull/9406) **aminesnow**: Add label selector to zone awareness TopologySpreadConstraints
- [9366](https://example.com/acme/logstore/pull/9366) **periklis**: Add support for custom tenant topology in rules
- [9315](https://example.com/acme/logstore/pull/9315) **aminesnow**: Add zone awareness spec to LogstoreStack
- [9343](https://example.com/acme/logstore/pull/9343) **JoaoBraveCoding**: Add default PodAntiAffinity to Query Frontend
- [9346](https://example.com/acme/logstore/pull/9346) **periklis**: Enable Route by default on OpenShift clusters
- [9339](https://example.com/acme/logstore/pull/9339) **JoaoBraveCoding**: Add default PodAntiAffinity to Ruler
- [9329](https://example.com/acme/logstore/pull/9329) **JoaoBraveCoding**: Add default PodAntiAffinity to Ingester
- [9262](https://example.com/acme/logstore/pull/9262) **btaani**: Add PodDisruptionBudget to the Ruler
- [9260](https://example.com/acme/logstore/pull/9260) **JoaoBraveCoding**: Add PodDisruptionBudgets to the ingestion path
- [9188](https://example.com/acme/logstore/pull/9188) **aminesnow**: Add PodDisruptionBudgets to the query path
- [9162](https://example.com/acme/logstore/pull/9162) **aminesnow**: Add a PodDisruptionBudget to logstorestack-gateway

## 0.3.0 (2023-04-20)

- [9049](https://example.com/acme/logstore/pull/9049) **alanconway**: Revert 1x.extra-small changes, add 1x.demo
- [8661](https://example.com/acme/logstore/pull/8661) **xuanyunhui**: Add a new Object Storage Type for AlibabaCloud OSS
- [9036](https://example.com/acme/logstore/pull/9036) **periklis**: Update Logstore operand to v2.8.0
- [8978](https://example.com/acme/logstore/pull/8978) **aminesnow**: Add watch for the object storage secret
- [8958](https://example.com/acme/logstore/pull/8958) **periklis**: Align common instance addr with memberlist advertise addr
- [8998](https://example.com/acme/logstore/pull/8998) **periklis**: Remove static placeholder suffix for openshift bundle
- [8930](https://example.com/acme/logstore/pull/8930) **periklis**: Fix makefile target operatorhub
- [8911](https://example.com/acme/logstore/pull/8911) **aminesnow**: Update LogstoreStack annotaion on RulerConfig delete

## 0.2.0 (2023-03-27)

- [8912](https://example.com/acme/logstore/pull/8912) **periklis**: Add missing replaces directives for release v0.2.0
- [8651](https://example.com/acme/logstore/pull/8651) **periklis**: Prepare Community Logstore Operator release v0.2.0
- [8881](https://example.com/acme/logstore/pull/8881) **periklis**: Provide community bundle for openshift community hub
- [8863](https://example.com/acme/logstore/pull/8863) **periklis**: Break the API types out into their own module
- [8878](https://example.com/acme/logstore/pull/8878) **periklis**: Refactor all type validations into own package
- [8875](https://example.com/acme/logstore/pull/8875) **Red-GV**: Remove mutations to non-updatable statefulset fields
- [7451](https://example.com/acme/logstore/pull/7451) **btaani**: Add support for rules configmap sharding
- [8672](https://example.com/acme/logstore/pull/8672) **periklis**: Add support for memberlist bind network configuration
- [8748](https://example.com/acme/logstore/pull/8748) **periklis**: Add alertingrule tenant id label for all rules
- [8743](https://example.com/acme/logstore/pull/8743) **periklis**: Add alerting style guide validation
- [8192](https://example.com/acme/logstore/pull/8192) **jotak**: Allow multiple matchers for multi-tenancy with Network tenant (OpenShift)
- [8800](https://example.com/acme/logstore/pull/8800) **aminesnow**: Promote AlertingRules, RecordingRules and RulerConfig from v1beta1 to v1
- [8792](https://example.com/acme/logstore/pull/8792) **orenc1**: Improve documentation for LogstoreStack installation with ODF Object Storage
- [8791](https://example.com/acme/logstore/pull/8791) **periklis**: Expand OLM skip range for OpenShift Logging 5.7 release (OpenShift)
- [8771](https://example.com/acme/logstore/pull/8771) **periklis**: Update Logstore operand to v2.7.4
- [8776](https://example.com/acme/logstore/pull/8776) **Red-GV**: Update go to v1.20, k8s libraries to v1.26.2, and OpenShift libraries to v4.13
- [8707](https://example.com/acme/logstore/pull/8707) **aminesnow**: Fix gateway's nodeSelector and tolerations
- [8666](https://example.com/acme/logstore/pull/8666) **xperimental**: Fix version inconsistency in generated OpenShift bundle
- [8578](https://example.com/acme/logstore/pull/8578) **xperimental**: Refactor status update to reduce API calls
- [8577](https://example.com/acme/logstore/pull/8577) **Red-GV**: Store gateway tenant information in secret instead of configmap
- [8397](https://example.com/acme/logstore/pull/8397) **periklis**: Update Logstore operand to v2.7.3
- [8308](https://example.com/acme/logstore/pull/8308) **aminesnow**: operator: Cleanup ruler resources when disabled
- [8336](https://example.com/acme/logstore/pull/8336) **periklis**: Update Logstore operand to v2.7.2
- [8253](https://example.com/acme/logstore/pull/8253) **aminesnow**: Add watch on the Alertmanager in openshift-monitoring and decouple it from the user-workload AM
- [8265](https://example.com/acme/logstore/pull/8265) **Red-GV**: Use gRPC compactor service instead of http for retention
- [8038](https://example.com/acme/logstore/pull/8038) **aminesnow**: Add watch on the Alertmanager in OCP's user-workload-monitoring namespace
- [8173](https://example.com/acme/logstore/pull/8173) **periklis**: Remove custom webhook cert mounts for OLM-based deployment (OpenShift)
- [8001](https://example.com/acme/logstore/pull/8001) **aminesnow**: Add API validation to Alertmanager header auth config
- [8087](https://example.com/acme/logstore/pull/8087) **xperimental**: Fix status not updating when state of pods changes

## 0.1.0 (2023-01-10)

- [8068](https://example.com/acme/logstore/pull/8068) **periklis**: Use logstorestack-gateway replicas from size table
- [8068](https://example.com/acme/logstore/pull/8068) **periklis**: Use logstorestack-gateway replicas from size table
- [7839](https://example.com/acme/logstore/pull/7839) **aminesnow**: Configure Alertmanager per-tenant
- [7910](https://example.com/acme/logstore/pull/7910) **periklis**: Update Logstore operand to v2.7.1
- [7815](https://example.com/acme/logstore/pull/7815) **periklis**: Apply delete client changes for compat with release-2.7.x
- [7809](https://example.com/acme/logstore/pull/7809) **xperimental**: Fix histogram-based alerting rules
- [7808](https://example.com/acme/logstore/pull/7808) **xperimental**: Replace fifocache usage by embedded_cache
- [7753](https://example.com/acme/logstore/pull/7753) **periklis**: Check for mandatory CA configmap name in ObjectStorageTLS spec
- [7744](https://example.com/acme/logstore/pull/7744) **periklis**: Fix object storage TLS spec CAKey descriptor
- [7716](https://example.com/acme/logstore/pull/7716) **aminesnow**: Migrate API docs generation tool
- [7710](https://example.com/acme/logstore/pull/7710) **periklis**: Fix LogstoreStackController watches for cluster-scoped resources
- [7682](https://example.com/acme/logstore/pull/7682) **periklis**: Refactor cluster proxy to use configv1.Proxy on OpenShift
- [7711](https://example.com/acme/logstore/pull/7711) **Red-GV**: Remove default value from replicationFactor field
- [7617](https://example.com/acme/logstore/pull/7617) **Red-GV**: Modify ingestionRate for respective shirt size
- [7592](https://example.com/acme/logstore/pull/7592) **aminesnow**: Update API docs generation using gen-crd-api-reference-docs
- [7448](https://example.com/acme/logstore/pull/7448) **periklis**: Add TLS support for compactor delete client
- [7596](https://example.com/acme/logstore/pull/7596) **periklis**: Fix fresh-installs with built-in cert management enabled
- [7064](https://example.com/acme/logstore/pull/7064) **periklis**: Add support for built-in cert management
- [7471](https://example.com/acme/logstore/pull/7471) **aminesnow**: Expose and migrate query_timeout in limits config
- [7437](https://example.com/acme/logstore/pull/7437) **aminesnow**: Fix Custom TLS profile setting for LogstoreStack on OpenShift
- [7415](https://example.com/acme/logstore/pull/7415) **aminesnow**: Add alert relabel config
- [7418](https://example.com/acme/logstore/pull/7418) **Red-GV**: Update golang to v1.19 and k8s dependencies to v0.25.2
- [7322](https://example.com/acme/logstore/pull/7322) **Red-GV**: Configuring server and client HTTP and GRPC TLS options
- [7272](https://example.com/acme/logstore/pull/7272) **aminesnow**: Use cluster monitoring alertmanager by default on openshift clusters
- [7295](https://example.com/acme/logstore/pull/7295) **xperimental**: Add extended-validation for rules on OpenShift
- [6951](https://example.com/acme/logstore/pull/6951) **Red-GV**: Adding operational Logstorestack alerts
- [7254](https://example.com/acme/logstore/pull/7254) **periklis**: Expose Logstore Ruler API via the logstorestack-gateway
- [7214](https://example.com/acme/logstore/pull/7214) **periklis**: Fix ruler GRPC tls client configuration
- [7201](https://example.com/acme/logstore/pull/7201) **xperimental**: Write configuration for per-tenant retention
- [7037](https://example.com/acme/logstore/pull/7037) **xperimental**: Skip enforcing matcher for certain tenants on OpenShift
- [7106](https://example.com/acme/logstore/pull/7106) **xperimental**: Manage global stream-based retention
- [7092](https://example.com/acme/logstore/pull/7092) **aminesnow**: Configure kube-rbac-proxy sidecar to use Intermediate TLS security profile in OCP
- [6870](https://example.com/acme/logstore/pull/6870) **aminesnow**: Configure gateway to honor the global tlsSecurityProfile on Openshift
- [6999](https://example.com/acme/logstore/pull/6999) **Red-GV**: Adding LogstoreStack Gateway alerts
- [7000](https://example.com/acme/logstore/pull/7000) **xperimental**: Configure default node affinity for all pods
- [6923](https://example.com/acme/logstore/pull/6923) **xperimental**: Reconcile owner reference for existing objects
- [6907](https://example.com/acme/logstore/pull/6907) **Red-GV**: Adding valid subscription annotation to operator metadata
- [6479](https://example.com/acme/logstore/pull/6749) **periklis**: Update Logstore operand to v2.6.1
- [6748](https://example.com/acme/logstore/pull/6748) **periklis**: Update go4.org/unsafe/assume-no-moving-gc to latest
- [6741](https://example.com/acme/logstore/pull/6741) **aminesnow**: Golang version to 1.18 and k8s client to 1.24
- [6669](https://example.com/acme/logstore/pull/6669) **xperimental**: Set minimum TLS version to 1.2 to support FIPS
- [6663](https://example.com/acme/logstore/pull/6663) **aminesnow**: Generalize live tail fix to all clusters using TLS
- [6443](https://example.com/acme/logstore/pull/6443) **aminesnow**: Fix live tail of logs not working on OpenShift-based clusters
- [6646](https://example.com/acme/logstore/pull/6646) **periklis**: Update Logstore operand to v2.6.0
- [6594](https://example.com/acme/logstore/pull/6594) **xperimental**: Disable client certificate authentication on gateway
- [6551](https://example.com/acme/logstore/pull/6561) **periklis**: Add operator docs for object storage
- [6549](https://example.com/acme/logstore/pull/6549) **periklis**: Refactor feature gates to use custom resource definition
- [6514](https://example.com/acme/logstore/pull/6514) **Red-GV** Update all pods and containers to be compliant with restricted Pod Security Standard
- [6531](https://example.com/acme/logstore/pull/6531) **periklis**: Use default interface_names for logstorestack clusters (IPv6 Support)
- [6411](https://example.com/acme/logstore/pull/6478) **aminesnow**: Support TLS enabled logstorestack-gateway for vanilla kubernetes deployments
- [6504](https://example.com/acme/logstore/pull/6504) **periklis**: Disable usage report on OpenShift
- [6474](https://example.com/acme/logstore/pull/6474) **periklis**: Bump logstore.acme.com/LogstoreStack from v1beta to v1
- [6411](https://example.com/acme/logstore/pull/6411) **Red-GV**: Extend schema validation in LogstoreStack webhook
- [6334](https://example.com/acme/logstore/pull/6433) **periklis**: Move operator cli flags to component config
- [6224](https://example.com/acme/logstore/pull/6224) **periklis**: Add support for GRPC over TLS for Logstore components
- [5952](https://example.com/acme/logstore/pull/5952) **Red-GV**: Add api to change storage schema version
- [6363](https://example.com/acme/logstore/pull/6363) **periklis**: Allow optional installation of webhooks (Kind)
- [6362](https://example.com/acme/logstore/pull/6362) **periklis**: Allow reduced tenant OIDC authentication requirements
- [6288](https://example.com/acme/logstore/pull/6288) **aminesnow**: Expose only an HTTPS gateway when in openshift mode
- [6195](https://example.com/acme/logstore/pull/6195) **periklis**: Add ruler config support
- [6198](https://example.com/acme/logstore/pull/6198) **periklis**: Add support for custom S3 CA
- [6199](https://example.com/acme/logstore/pull/6199) **Red-GV**: Update GCP secret volume path
- [6125](https://example.com/acme/logstore/pull/6125) **sasagarw**: Add method to get authenticated from GCP
- [5986](https://example.com/acme/logstore/pull/5986) **periklis**: Add support for Logstore Rules reconciliation
- [5987](https://example.com/acme/logstore/pull/5987) **Red-GV**: Update logerr to v2.0.0
- [5907](https://example.com/acme/logstore/pull/5907) **xperimental**: Do not include non-static labels in pod selectors
- [5893](https://example.com/acme/logstore/pull/5893) **periklis**: Align PVC storage size requests for all logstorestack t-shirt sizes
- [5884](https://example.com/acme/logstore/pull/5884) **periklis**: Update Logstore operand to v2.5.0
- [5748](https://example.com/acme/logstore/pull/5748) **Red-GV**: Update Prometheus go client to 12.1
- [5739](https://example.com/acme/logstore/pull/5739) **sasagarw**: Change UUIDs to tenant name in doc
- [5729](https://example.com/acme/logstore/pull/5729) **periklis**: Add missing label matcher for openshift logging tenant mode (OpenShift)
- [5691](https://example.com/acme/logstore/pull/5691) **sasagarw**: Fix immediate reset of degraded condition
- [5704](https://example.com/acme/logstore/pull/5704) **xperimental**: Update operator-sdk to 1.18.1
- [5693](https://example.com/acme/logstore/pull/5693) **periklis**: Replace frontend_worker parallelism with match_max_concurrent
- [5699](https://example.com/acme/logstore/pull/5699) **Red-GV**: Configure boltdb_shipper and schema to use Azure, GCS, and Swift storage
- [5701](https://example.com/acme/logstore/pull/5701) **sasagarw**: Make ReplicationFactor optional in LogstoreStack API
- [5695](https://example.com/acme/logstore/pull/5695) **xperimental**: Update Go to 1.17
- [5615](https://example.com/acme/logstore/pull/5615) **sasagarw**: Document how to connect to LogstoreStack gateway component
- [5655](https://example.com/acme/logstore/pull/5655) **xperimental**: Update Logstore operand to 2.4.2
- [5579](https://example.com/acme/logstore/pull/5579) **Red-GV**: Add playbook for responding to operator alerts
- [5640](https://example.com/acme/logstore/pull/5640) **sasagarw**: Update CSV to point to candidate channel and use openshift-operators-redhat ns (OpenShift)
- [5551](https://example.com/acme/logstore/pull/5551) **sasagarw**: Document how to connect to distributor component
- [5624](https://example.com/acme/logstore/pull/5624) **periklis**: Use tenant name as id for mode openshift-logging (OpenShift)
- [5621](https://example.com/acme/logstore/pull/5621) **periklis**: Use recommended labels for LogstoreStack components
- [5607](https://example.com/acme/logstore/pull/5607) **periklis**: Use logstorestack name as prefix for owned resources
- [5588](https://example.com/acme/logstore/pull/5588) **periklis**: Add RBAC for Prometheus service discovery to Logstore component metrics (OpenShift)
- [5576](https://example.com/acme/logstore/pull/5576) **xperimental**: Change endpoints for generated liveness and readiness probes
- [5560](https://example.com/acme/logstore/pull/5560) **periklis**: Fix service monitor's server name for operator metrics
- [5345](https://example.com/acme/logstore/pull/5345) **ronensc**: Add flag to create Prometheus rules
- [5432](https://example.com/acme/logstore/pull/5432) **Red-GV**: Provide storage configuration for Azure, GCS, and Swift through common_config
- [4975](https://example.com/acme/logstore/pull/4975) **periklis**: Provide saner default for logstore-operator managed chunk_target_size
