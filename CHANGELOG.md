# Changelog

## Unreleased

### Bug Fixes

* **promtail:** fix parser for azureeventhubs message without time field ([#14218](https://example.com/acme/logstore/pull/14218))

## [3.1.1](https://example.com/acme/logstore/compare/v3.1.0...v3.1.1) (2024-08-08)


### Features

* **logstore:** add ability to disable AWS S3 dual stack endpoints usage ([#13795](https://example.com/acme/logstore/issues/13795)) ([464ac73](https://example.com/acme/logstore/commit/464ac736a6fb70b673ee3cec21049b18d353cadb))


### Bug Fixes

* **deps:** bumped dependencies versions to resolve CVEs ([#13789](https://example.com/acme/logstore/issues/13789)) ([34206cd](https://example.com/acme/logstore/commit/34206cd2d6290566034710ae6c2d08af8804bc91))

## [3.1.0](https://example.com/acme/logstore/compare/v3.0.0...v3.1.0) (2024-07-02)


### ⚠ BREAKING CHANGES

* update helm chart to support distributed mode and 3.0 ([#12067](https://example.com/acme/logstore/issues/12067))

### Features

* Add a version of the mixin dashboards for meta monitoring ([#12700](https://example.com/acme/logstore/issues/12700)) ([ec1a057](https://example.com/acme/logstore/commit/ec1a057a323ed1bd8de448e714a672b64140b691))
* Add backoff to flush op ([#13140](https://example.com/acme/logstore/issues/13140)) ([9767807](https://example.com/acme/logstore/commit/9767807680cb4149c7b56345c531b62105a1b976))
* add detected-fields command to logcli ([#12739](https://example.com/acme/logstore/issues/12739)) ([210ea93](https://example.com/acme/logstore/commit/210ea93a690b1b9746b3ff62bbd5d217a3bc8e8e))
* Add ingester_chunks_flush_failures_total ([#12925](https://example.com/acme/logstore/issues/12925)) ([cc3694e](https://example.com/acme/logstore/commit/cc3694eecddaab579d08328cdab78a7d8a7bd720))
* add logstoretool ([#12166](https://example.com/acme/logstore/issues/12166)) ([7b7d3d4](https://example.com/acme/logstore/commit/7b7d3d4cd2c979c778d3741156f0d765a9e531b2))
* Add metrics for number of patterns detected & evicted ([#12918](https://example.com/acme/logstore/issues/12918)) ([bc53b33](https://example.com/acme/logstore/commit/bc53b337218425af5b5ce69dcef56e27afec6647))
* Add new Drain tokenizer that splits on most punctuation ([#13143](https://example.com/acme/logstore/issues/13143)) ([6a0fdd0](https://example.com/acme/logstore/commit/6a0fdd088091fc37e3e9424c78a2d6d587dbbb33))
* Add pattern ingester support in SSD mode ([#12685](https://example.com/acme/logstore/issues/12685)) ([19bfef4](https://example.com/acme/logstore/commit/19bfef48cbad57468591e8214c4a5f390091f1e1))
* add profile tagging to ingester ([#13068](https://example.com/acme/logstore/issues/13068)) ([00d3c7a](https://example.com/acme/logstore/commit/00d3c7a52d9f2b48fccb0cd5b105a2577b3d0305))
* add recalculateOwnedStreams to check stream ownership if the ring is changed ([#13103](https://example.com/acme/logstore/issues/13103)) ([e7689b2](https://example.com/acme/logstore/commit/e7689b248dbe549b2ac61a0e335d8b5b999cc47d))
* Add step param to Patterns Query API ([#12703](https://example.com/acme/logstore/issues/12703)) ([7b8533e](https://example.com/acme/logstore/commit/7b8533e435cf9d0466d3b147b2b3e0f6b3613fe9))
* Add tokenizer interface for Drain Training ([#13069](https://example.com/acme/logstore/issues/13069)) ([797bb64](https://example.com/acme/logstore/commit/797bb641736a2355b4f8503c147fc0c8a814f19a))
* add toleration for bloom components ([#12653](https://example.com/acme/logstore/issues/12653)) ([fcb2b0a](https://example.com/acme/logstore/commit/fcb2b0a16a7692ee0a705ce239375843a63246c7))
* Add utf8 support to Pattern Lexer to support utf8 chars ([#13085](https://example.com/acme/logstore/issues/13085)) ([f6f8bab](https://example.com/acme/logstore/commit/f6f8babf83f3d90f4e6f3f9b732fe22382861f47))
* add warnings to metadata context directly ([#12579](https://example.com/acme/logstore/issues/12579)) ([c4ac8cc](https://example.com/acme/logstore/commit/c4ac8cc009a75b616f867701c440797f655bcd1b))
* Added getting started video ([#12975](https://example.com/acme/logstore/issues/12975)) ([8442dca](https://example.com/acme/logstore/commit/8442dca9d2341471996a73a011f206630c67e857))
* Added Interactive Sandbox to Quickstart tutorial ([#12701](https://example.com/acme/logstore/issues/12701)) ([97212ea](https://example.com/acme/logstore/commit/97212eadf15c2b5ee2cd59b7c1df71f6177cfe7e))
* Added video and updated Acme Agent -> Alloy ([#13032](https://example.com/acme/logstore/issues/13032)) ([1432a3e](https://example.com/acme/logstore/commit/1432a3e84a7e5df18b8dc0e217121fd78da9e75e))
* API: Expose optional label matcher for label names API ([#11982](https://example.com/acme/logstore/issues/11982)) ([8084259](https://example.com/acme/logstore/commit/808425953fa8a8eca3199b3664e43ceba362747a))
* area/promtail: Added support to install wget on promtail docker image to support docker healthcheck ([#11711](https://example.com/acme/logstore/issues/11711)) ([ffe684c](https://example.com/acme/logstore/commit/ffe684c330bcd65f9b07a02d6f93bb475106becc))
* **blooms:** Add counter metric for blocks that are not available at query time ([#12968](https://example.com/acme/logstore/issues/12968)) ([d6374bc](https://example.com/acme/logstore/commit/d6374bc2ce3041005842edd353a3bb010f467abe))
* **blooms:** Add in-memory LRU cache for meta files ([#12862](https://example.com/acme/logstore/issues/12862)) ([fcd544c](https://example.com/acme/logstore/commit/fcd544c2d9d52b62d09e31c532a5cd2115f4d2bc))
* **blooms:** Blooms/v2 encoding multipart series ([#13093](https://example.com/acme/logstore/issues/13093)) ([fbe7c55](https://example.com/acme/logstore/commit/fbe7c559b5ed153fb46a1965c24180011a558b85))
* **blooms:** compute chunks once ([#12664](https://example.com/acme/logstore/issues/12664)) ([bc78d13](https://example.com/acme/logstore/commit/bc78d13d9b736bb9313403569d0f69e85663afce))
* **blooms:** ignore individual bloom-gw failures ([#12863](https://example.com/acme/logstore/issues/12863)) ([4c9b22f](https://example.com/acme/logstore/commit/4c9b22f11077b560d21f086a84d42176e9196d5b))
* **blooms:** ingester aware bounded impl ([#12840](https://example.com/acme/logstore/issues/12840)) ([7bbd8b5](https://example.com/acme/logstore/commit/7bbd8b5087d637ac592403c5daafda35353fe13d))
* **bloom:** Skip attempts to filter chunks for which blooms have not been built ([#12961](https://example.com/acme/logstore/issues/12961)) ([a1b1eeb](https://example.com/acme/logstore/commit/a1b1eeb09583f04a36ebdb96f716f3f285b90adf))
* **blooms:** limit bloom size during creation ([#12796](https://example.com/acme/logstore/issues/12796)) ([eac5622](https://example.com/acme/logstore/commit/eac56224b8e228a694090ffaee47300b23eeb13b))
* **blooms:** record time spent resolving shards ([#12636](https://example.com/acme/logstore/issues/12636)) ([9c25985](https://example.com/acme/logstore/commit/9c25985b970865f054dfa9243cbe984d921df3c8))
* **blooms:** Separate page buffer pools for series pages and bloom pages ([#12992](https://example.com/acme/logstore/issues/12992)) ([75ccf21](https://example.com/acme/logstore/commit/75ccf2160bfe647b1cb3daffb98869e9c1c44130))
* Boilerplate for new bloom build planner and worker components. ([#12989](https://example.com/acme/logstore/issues/12989)) ([8978ecf](https://example.com/acme/logstore/commit/8978ecf0c85dfbe18b52632112e5be20eff411cf))
* **cache:** Add `Cache-Control: no-cache` support for Logstore instant queries. ([#12896](https://example.com/acme/logstore/issues/12896)) ([88e545f](https://example.com/acme/logstore/commit/88e545fc952d6ff55c61d079db920f00abc04865))
* **canary:** Add test to check query results with and without cache. ([#13104](https://example.com/acme/logstore/issues/13104)) ([71507a2](https://example.com/acme/logstore/commit/71507a2b640ad071d88ee894e80235f93be73c3d))
* Detected labels from store ([#12441](https://example.com/acme/logstore/issues/12441)) ([587a6d2](https://example.com/acme/logstore/commit/587a6d20e938f4f58e5a49563a3c267762cf89eb))
* **detected-labels:** include labels with cardinality > 1 ([#13128](https://example.com/acme/logstore/issues/13128)) ([8be8364](https://example.com/acme/logstore/commit/8be8364435bb83dd134580ba6fc1f0bdb5474356))
* **detectedFields:** add parser to response ([#12872](https://example.com/acme/logstore/issues/12872)) ([2b3ae48](https://example.com/acme/logstore/commit/2b3ae48d9be63183907dfd7163af6a980360c853))
* **detectedFields:** Support multiple parsers to be returned for a single field ([#12899](https://example.com/acme/logstore/issues/12899)) ([19fef93](https://example.com/acme/logstore/commit/19fef9355fdd46911611dbec25df0f5a4e397d31))
* Enable log volume endpoint by default ([#12628](https://example.com/acme/logstore/issues/12628)) ([397aa56](https://example.com/acme/logstore/commit/397aa56e157cbf733da548474a4bcae773e82362))
* Enable log volume endpoint by default in helm ([#12690](https://example.com/acme/logstore/issues/12690)) ([e39677f](https://example.com/acme/logstore/commit/e39677f97b4ba27c90d9f8d2991441095e55b06e))
* Generic logline placeholder replacement and tokenization ([#12799](https://example.com/acme/logstore/issues/12799)) ([4047902](https://example.com/acme/logstore/commit/40479029d74d588268956190d956a088aed682e1))
* **helm:** Allow extraObject items as multiline strings ([#12397](https://example.com/acme/logstore/issues/12397)) ([af5be90](https://example.com/acme/logstore/commit/af5be900764acfe4bff54ceef164a4f660990f8a))
* **helm:** Support for PVC Annotations for Non-Distributed Modes ([#12023](https://example.com/acme/logstore/issues/12023)) ([efdae3d](https://example.com/acme/logstore/commit/efdae3df14c47d627eb99e91466e0451db6e16f6))
* improve performance of `first_over_time` and `last_over_time` queries by sharding them ([#11605](https://example.com/acme/logstore/issues/11605)) ([f66172e](https://example.com/acme/logstore/commit/f66172eed17f9418ab22615537c7b65b09de96e5))
* improve syntax parser for pattern ([#12489](https://example.com/acme/logstore/issues/12489)) ([48dae44](https://example.com/acme/logstore/commit/48dae4417cca75a40d6a3bf16b0d976714e8db81))
* include the stream we failed to create in the stream limit error message ([#12437](https://example.com/acme/logstore/issues/12437)) ([ec81991](https://example.com/acme/logstore/commit/ec81991f4d7f6d83a34dffb073d60c330c69e94d))
* Increase drain max depth from 8 -> 30 ([#13063](https://example.com/acme/logstore/issues/13063)) ([d0a2859](https://example.com/acme/logstore/commit/d0a285926b7257d54cf948ba644c619a4b49a871))
* Introduce `index audit` to `logstoretool` ([#13008](https://example.com/acme/logstore/issues/13008)) ([47f0236](https://example.com/acme/logstore/commit/47f0236ea8f33a67a0a1abf6e6d6b3582661c4ba))
* logstore/main.go: Log which config file path is used on startup ([#12985](https://example.com/acme/logstore/issues/12985)) ([7a3338e](https://example.com/acme/logstore/commit/7a3338ead82e4c577652ab86e9a55faf200ac05a))
* new stream count limiter ([#13006](https://example.com/acme/logstore/issues/13006)) ([1111595](https://example.com/acme/logstore/commit/1111595179c77f9303ebdfd362f14b1ac50044cb))
* Optimize log parsing performance by using unsafe package ([#13223](https://example.com/acme/logstore/issues/13223)) ([9f31b25](https://example.com/acme/logstore/commit/9f31b25253502f035cfb6a831bcea7f778f427dd))
* parameterise the MaximumEventAgeInSeconds, LogGroupName, and IAMRoleName for lambda-promtail CloudFormation template ([#12728](https://example.com/acme/logstore/issues/12728)) ([8892dc8](https://example.com/acme/logstore/commit/8892dc89231ebe7b05fc1c4e0b7647f328f9c1ce))
* **promtail:** Support of RFC3164 aka BSD Syslog ([#12810](https://example.com/acme/logstore/issues/12810)) ([be41525](https://example.com/acme/logstore/commit/be4152576e6d8cb280fd65604199db7157981f07))
* Querier: Split gRPC client into two. ([#12726](https://example.com/acme/logstore/issues/12726)) ([7b6f057](https://example.com/acme/logstore/commit/7b6f0577c3277b84230f0f2deba747b01ca2b2fa))
* **reporting:** Report cpu usage ([#12970](https://example.com/acme/logstore/issues/12970)) ([87288d3](https://example.com/acme/logstore/commit/87288d37f9e9c1e90295bf785adbc4bfdb66fb30))
* split detected fields queries ([#12491](https://example.com/acme/logstore/issues/12491)) ([6c33809](https://example.com/acme/logstore/commit/6c33809015bef8078b17dcb6b0701e930132f042))
* Support negative numbers in LogQL ([#13091](https://example.com/acme/logstore/issues/13091)) ([6df81db](https://example.com/acme/logstore/commit/6df81db978b0157ab96fa0629a311f919dad1e8a))
* Tune Patterns query drain instance ([#13137](https://example.com/acme/logstore/issues/13137)) ([30df31e](https://example.com/acme/logstore/commit/30df31e28b5c360ffed2dea3b47f515e4e24146d))
* Update getting started demo to Logstore 3.0 ([#12723](https://example.com/acme/logstore/issues/12723)) ([282e385](https://example.com/acme/logstore/commit/282e38548ceb96b1c518010c47b8eabf4317e8fd))
* update helm chart to support distributed mode and 3.0 ([#12067](https://example.com/acme/logstore/issues/12067)) ([79b876b](https://example.com/acme/logstore/commit/79b876b65d55c54f4d532e98dc24743dea8bedec))
* Update Logstore monitoring docs to new meta monitoring helm ([#13176](https://example.com/acme/logstore/issues/13176)) ([b4d44f8](https://example.com/acme/logstore/commit/b4d44f89f997e59c84e69ed075341bb6e1371d08))
* Updated best practises for labels ([#12749](https://example.com/acme/logstore/issues/12749)) ([6ebfbe6](https://example.com/acme/logstore/commit/6ebfbe658bbd92e3599ca4aff3bcfdd302d3cc32))
* Updated SS and microservices deployment docs ([#13083](https://example.com/acme/logstore/issues/13083)) ([1b80458](https://example.com/acme/logstore/commit/1b80458e2eff2d41b9126a7529ee32ae1e269f05))


### Bug Fixes

* `codec` not initialized in downstream roundtripper ([#12873](https://example.com/acme/logstore/issues/12873)) ([b6049f6](https://example.com/acme/logstore/commit/b6049f6792492d5753626e5845b0094199463966))
* Add a missing `continue` in fuse which may cause incorrect bloom test result ([#12650](https://example.com/acme/logstore/issues/12650)) ([0d1ebeb](https://example.com/acme/logstore/commit/0d1ebebd3afe9504506aaed0b7827318eb2d9cfe))
* Add copyString function to symbolizer to avoid retaining  memory ([#13146](https://example.com/acme/logstore/issues/13146)) ([86b119a](https://example.com/acme/logstore/commit/86b119ac7ba206d294eb257f99c308fe8452bd58))
* add detected_level info when the info word appears on log message ([#13218](https://example.com/acme/logstore/issues/13218)) ([c9bfa3e](https://example.com/acme/logstore/commit/c9bfa3ebbf362b3d056879f0ef5f3e656f28c500))
* Add missing Helm helper logstore.hpa.apiVersion ([#12755](https://example.com/acme/logstore/issues/12755)) ([3070ea7](https://example.com/acme/logstore/commit/3070ea70bb05bffced6a8304f506b03ed4c8e2aa))
* Add missing OTLP endpoint to nginx config ([#12709](https://example.com/acme/logstore/issues/12709)) ([8096748](https://example.com/acme/logstore/commit/8096748f1f205e766deab9438c4b2bc587facfc5))
* add missing parentheses in meta monitoring dashboards ([#12802](https://example.com/acme/logstore/issues/12802)) ([151d0a5](https://example.com/acme/logstore/commit/151d0a58ac9f5aa67f944e6729720f5f70d07e27))
* add retry middleware to the "limited" query roundtripper ([#13161](https://example.com/acme/logstore/issues/13161)) ([bb864b3](https://example.com/acme/logstore/commit/bb864b3ad63d61f5b091a9cc04246da2f44b2157))
* allow cluster label override in bloom dashboards ([#13012](https://example.com/acme/logstore/issues/13012)) ([987e551](https://example.com/acme/logstore/commit/987e551f9e21b9a612dd0b6a3e60503ce6fe13a8))
* **blooms:** bloomshipper no longer returns empty metas on fetch ([#13130](https://example.com/acme/logstore/issues/13130)) ([ad279e5](https://example.com/acme/logstore/commit/ad279e518cb252ef7e26283ec16540846dbd3acf))
* **blooms:** Clean block directories recursively on startup ([#12895](https://example.com/acme/logstore/issues/12895)) ([7b77e31](https://example.com/acme/logstore/commit/7b77e310982147162777f9febfbcd98ec8a8c383))
* **blooms:** Correctly return unfiltered chunks for series that are not mapped to any block ([#12774](https://example.com/acme/logstore/issues/12774)) ([c36b114](https://example.com/acme/logstore/commit/c36b1142c7acd6a13a3634ddbef71254040cff73))
* **blooms:** Deduplicate filtered series and chunks ([#12791](https://example.com/acme/logstore/issues/12791)) ([3bf2d1f](https://example.com/acme/logstore/commit/3bf2d1fea08593bdf10dc8a6827998a6d8a8243c))
* **blooms:** Disable metas cache on bloom gateway ([#12959](https://example.com/acme/logstore/issues/12959)) ([00bdd2f](https://example.com/acme/logstore/commit/00bdd2f5b703991b280317ceff0fcf2eed1847d9))
* **blooms:** Do not fail requests when fetching metas from cache fails ([#12838](https://example.com/acme/logstore/issues/12838)) ([667076d](https://example.com/acme/logstore/commit/667076d9359c56118f1149f31a94c8a44bc171c7))
* **blooms:** dont break iterator conventions ([#12808](https://example.com/acme/logstore/issues/12808)) ([1665e85](https://example.com/acme/logstore/commit/1665e853a0a6aa63f535bcc5a4bb67775723cc87))
* **blooms:** Fix `partitionSeriesByDay` function ([#12900](https://example.com/acme/logstore/issues/12900)) ([738c274](https://example.com/acme/logstore/commit/738c274a5828aab4d88079c38400ddc705c0cb5d))
* **blooms:** Fix a regression introduced with [#12774](https://example.com/acme/logstore/issues/12774) ([#12776](https://example.com/acme/logstore/issues/12776)) ([ecefb49](https://example.com/acme/logstore/commit/ecefb495084a59d25778af520041766e087598ba))
* **blooms:** Fix findGaps when ownership goes to MaxUInt64 and that is covered by existing meta ([#12558](https://example.com/acme/logstore/issues/12558)) ([0ee2a61](https://example.com/acme/logstore/commit/0ee2a6126ae40a1d666f500c19efd639763f1bae))
* **blooms:** Fully deduplicate chunks from FilterChunkRef responses ([#12807](https://example.com/acme/logstore/issues/12807)) ([a0f358f](https://example.com/acme/logstore/commit/a0f358fcc8295d93ee38b67738e8d90045c50dab))
* **blooms:** Handle not found metas gracefully ([#12853](https://example.com/acme/logstore/issues/12853)) ([37c8822](https://example.com/acme/logstore/commit/37c88220b3a7f8268c48f1bf37f4eb11cdba1b5f))
* **blooms:** Reset error on LazyBloomIter.Seek ([#12806](https://example.com/acme/logstore/issues/12806)) ([76ba24e](https://example.com/acme/logstore/commit/76ba24e3d8ce5e3c872442ce9d64505605ef0f53))
* change log level since this is a known case ([#13029](https://example.com/acme/logstore/issues/13029)) ([ca030a5](https://example.com/acme/logstore/commit/ca030a5c4335b0258e83aebd8779ea4d348003f3))
* close res body ([#12444](https://example.com/acme/logstore/issues/12444)) ([616977a](https://example.com/acme/logstore/commit/616977a942b63fb2ee7545e155abe246f6175308))
* Correctly encode step when translating proto to http internally ([#13171](https://example.com/acme/logstore/issues/13171)) ([740551b](https://example.com/acme/logstore/commit/740551bb31e0c1806de8d87f02fa4f507aa24092))
* crrect initialization of a few slices ([#12674](https://example.com/acme/logstore/issues/12674)) ([0eba448](https://example.com/acme/logstore/commit/0eba448fc70b78ca7cd612831c9d3be116faa7a2))
* Defer closing blocks iter after checking error from loadWorkForGap ([#12934](https://example.com/acme/logstore/issues/12934)) ([cb1f5d9](https://example.com/acme/logstore/commit/cb1f5d9fca2908bd31a3c6bef38d49fe084d2939))
* Do not filter out chunks for store when `From==Through` and `From==start`  ([#13117](https://example.com/acme/logstore/issues/13117)) ([d9cc513](https://example.com/acme/logstore/commit/d9cc513fd2decf96d047d388136417c03ccdc682))
* **docs:** broken link in getting started readme ([#12736](https://example.com/acme/logstore/issues/12736)) ([425a2d6](https://example.com/acme/logstore/commit/425a2d690c13592abf32f2ed2475676c3422ac51))
* **docs:** Move promtail configuration to the correct doc ([#12737](https://example.com/acme/logstore/issues/12737)) ([1161846](https://example.com/acme/logstore/commit/1161846e19105e2669a5b388998722c23bd0f2f4))
* Ensure Drain patterns are valid for LogQL pattern match filter ([#12815](https://example.com/acme/logstore/issues/12815)) ([fd2301f](https://example.com/acme/logstore/commit/fd2301fd62b18eb345bc43868b40343efc1a1f10))
* errors reported by the race detector ([#13174](https://example.com/acme/logstore/issues/13174)) ([2b19dac](https://example.com/acme/logstore/commit/2b19dac82a97b1d75075eb87a4f7fdfed003c072)), closes [#8586](https://example.com/acme/logstore/issues/8586)
* Fix bloom deleter PR after merge ([#13167](https://example.com/acme/logstore/issues/13167)) ([c996349](https://example.com/acme/logstore/commit/c99634978cb189744946e6dc388f0cc4183e98f2))
* Fix compactor matcher in the logstore-deletion dashboard ([#12790](https://example.com/acme/logstore/issues/12790)) ([a03846b](https://example.com/acme/logstore/commit/a03846b4367cbb5a0aa445e539d92ae41e3f481a))
* Fix duplicate enqueue item problem in bloom download queue when do sync download ([#13114](https://example.com/acme/logstore/issues/13114)) ([f98ff7f](https://example.com/acme/logstore/commit/f98ff7f58400b5f5a425fae003fb959bfb8c6454))
* Fix for how the loop sync is done ([#12941](https://example.com/acme/logstore/issues/12941)) ([5cd850e](https://example.com/acme/logstore/commit/5cd850e0d02151c6f9c6285189b887b4929cfa12))
* Fix incorrect sorting of chunks in bloom-filtered response since `ChunkRef.Cmp` method is used in reverse ([#12999](https://example.com/acme/logstore/issues/12999)) ([670cd89](https://example.com/acme/logstore/commit/670cd89aa8ffb8b852bca05fd0adb554e93ce796))
* Fix indentation of query range values in helm ([#12577](https://example.com/acme/logstore/issues/12577)) ([9823f20](https://example.com/acme/logstore/commit/9823f2030a294e6dc9c50d6f956a7691df5d53df))
* Fix log level detection ([#12651](https://example.com/acme/logstore/issues/12651)) ([6904a65](https://example.com/acme/logstore/commit/6904a6520d3b5599404b339577c7c3311e635da9))
* Fix panic on requesting out-of-order Pattern samples ([#13010](https://example.com/acme/logstore/issues/13010)) ([2171f64](https://example.com/acme/logstore/commit/2171f6409f7157888df9637a635664c67b7ca844))
* fix parsing of default per tenant otlp config ([#12836](https://example.com/acme/logstore/issues/12836)) ([7cc9a93](https://example.com/acme/logstore/commit/7cc9a9386a8f89dbec6a25435180ed4625ae6490))
* fix setting of info log level when trying to detect level from log lines ([#12635](https://example.com/acme/logstore/issues/12635)) ([0831802](https://example.com/acme/logstore/commit/0831802a99243f9fe61f6cc8795739bf67e8d8e9))
* Fix the logstoretool imports ([#12673](https://example.com/acme/logstore/issues/12673)) ([6dce988](https://example.com/acme/logstore/commit/6dce98870d8c5c7054b3444d2fe4e66dad262a53))
* Fixes read & backend replicas settings ([#12828](https://example.com/acme/logstore/issues/12828)) ([d751134](https://example.com/acme/logstore/commit/d7511343bcdfe77a6213599827ce0093b2949c18))
* helm: Set compactor addr for distributed mode. ([#12748](https://example.com/acme/logstore/issues/12748)) ([521d40a](https://example.com/acme/logstore/commit/521d40a96a5c1c65c786c73ec374580fe767dd3b))
* **helm:** Fix GEL image tag, bucket name and proxy URLs ([#12878](https://example.com/acme/logstore/issues/12878)) ([67ed2f7](https://example.com/acme/logstore/commit/67ed2f7092c8c0d97ba0bec08fde7ede65faa33f))
* **helm:** fix query-frontend and ruler targetPort 'http-metrics' in Service template ([#13024](https://example.com/acme/logstore/issues/13024)) ([1ab9d27](https://example.com/acme/logstore/commit/1ab9d271c354caf0ba589691e6477fb9a19039f0))
* **helm:** fix queryScheduler servicemonitor ([#12753](https://example.com/acme/logstore/issues/12753)) ([8101e21](https://example.com/acme/logstore/commit/8101e21f9973b8261de0ee3eb34fa4d7b88ddaac))
* **helm:** fixed ingress paths mapping ([#12932](https://example.com/acme/logstore/issues/12932)) ([5ada92b](https://example.com/acme/logstore/commit/5ada92b190c671055bb09ca2dd234b6bac49289e))
* **helm:** only default bucket names when using minio ([#12548](https://example.com/acme/logstore/issues/12548)) ([2e32ec5](https://example.com/acme/logstore/commit/2e32ec52d8766c0a5a75be30585402f1dce52cc5))
* **helm:** Removed duplicate bucketNames from documentation and fixed key name `deploymentMode` ([#12641](https://example.com/acme/logstore/issues/12641)) ([0d8ff9e](https://example.com/acme/logstore/commit/0d8ff9ee7929b8facbdb469abe344c320d3bd5ce))
* incorrect compactor matcher in logstore-deletion dashboard mixin ([#12567](https://example.com/acme/logstore/issues/12567)) ([006f88c](https://example.com/acme/logstore/commit/006f88cef19d4d1fe14a40287ccdf534f6975475))
* **indexstats:** do not collect stats from "IndexStats" lookups for other query types ([#12978](https://example.com/acme/logstore/issues/12978)) ([1f5291a](https://example.com/acme/logstore/commit/1f5291a4a3bd3c98c190d9a5dda32bbd78f18c3b))
* Ingester zoneAwareReplication ([#12659](https://example.com/acme/logstore/issues/12659)) ([9edb0ce](https://example.com/acme/logstore/commit/9edb0ce140c4fe716a62e81e0fce747d92954f4c))
* Introduce feature flag for [last|first]_over_time sharding. ([#13067](https://example.com/acme/logstore/issues/13067)) ([6e45550](https://example.com/acme/logstore/commit/6e4555010eab5a2b12caf9af2df5f0991362d754))
* Invalidate caches when pipeline wrappers are disabled ([#12903](https://example.com/acme/logstore/issues/12903)) ([a772ed7](https://example.com/acme/logstore/commit/a772ed705c6506992cd1f2364b11fa60c1879f57))
* **ksonnet:** Do not generate rbac for consul if you are using memberlist ([#12688](https://example.com/acme/logstore/issues/12688)) ([2d62fca](https://example.com/acme/logstore/commit/2d62fca05d6ec82196b46c956733c89439660754))
* lambda-promtail, update s3 filename regex to allow finding of log files from AWS GovCloud regions ([#12482](https://example.com/acme/logstore/issues/12482)) ([7a81d26](https://example.com/acme/logstore/commit/7a81d264a4ba54efdb1d79d382fd4188c036aaee))
* logstore version prefix in Makefile ([#12514](https://example.com/acme/logstore/issues/12514)) ([dff72d2](https://example.com/acme/logstore/commit/dff72d2a52094fb2a831b5930cbfc67759b0978d))
* logstore-operational.libsonnet ([#12789](https://example.com/acme/logstore/issues/12789)) ([51a841f](https://example.com/acme/logstore/commit/51a841f20dbcbcb233836373ee246fb723ef70ba))
* make detected fields work for both json and proto ([#12682](https://example.com/acme/logstore/issues/12682)) ([f68d1f7](https://example.com/acme/logstore/commit/f68d1f7fafa1ec55e90d3a253ef2ee8bb9c2e342))
* make the tsdb filenames correctly reproducible from the identifier ([#12536](https://example.com/acme/logstore/issues/12536)) ([ec888ec](https://example.com/acme/logstore/commit/ec888ec8a564c7a93937c785c0540e7d2bcde20e))
* Missing password for Logstore-Canary when logstore.auth_enabled is true ([#12411](https://example.com/acme/logstore/issues/12411)) ([68b23dc](https://example.com/acme/logstore/commit/68b23dc2b5c74b9175d5e24fb445748c422cb7b6))
* mixin generation when cluster label is changed ([#12613](https://example.com/acme/logstore/issues/12613)) ([1ba7a30](https://example.com/acme/logstore/commit/1ba7a303566610363c0c36c87e7bc6bb492dfc93))
* **mixin:** dashboards $__auto fix ([#12707](https://example.com/acme/logstore/issues/12707)) ([91ef72f](https://example.com/acme/logstore/commit/91ef72f742fe1f8621af15d8190c5c0d4d613ab9))
* Mixins - Add missing log datasource on logstore-deletion ([#13011](https://example.com/acme/logstore/issues/13011)) ([1948899](https://example.com/acme/logstore/commit/1948899999107e7f27f4b9faace64942abcdb41f))
* **mixins:** Align logstore-writes mixins with logstore-reads ([#13022](https://example.com/acme/logstore/issues/13022)) ([757b776](https://example.com/acme/logstore/commit/757b776de39bf0fc0c6d1dd74e4a245d7a99023a))
* **nix:** lambda-promtail vendor hash ([#12763](https://example.com/acme/logstore/issues/12763)) ([ae180d6](https://example.com/acme/logstore/commit/ae180d6e070946eb5359ecd63a9e01e02f160ce3))
* not owned stream count ([#13030](https://example.com/acme/logstore/issues/13030)) ([4901a5c](https://example.com/acme/logstore/commit/4901a5c452fa6822a645f56e20e704db9366182a))
* **operator:** add alertmanager client config to ruler template ([#13182](https://example.com/acme/logstore/issues/13182)) ([6148c37](https://example.com/acme/logstore/commit/6148c3760d701768e442186d4e7d574c7dc16c91))
* **operator:** Bump golang builder to 1.21.9 ([#12503](https://example.com/acme/logstore/issues/12503)) ([f680ee0](https://example.com/acme/logstore/commit/f680ee0453d1b7d315774591293927b988bca223))
* **operator:** Configure Logstore to use virtual-host-style URLs for S3 AWS endpoints ([#12469](https://example.com/acme/logstore/issues/12469)) ([0084262](https://example.com/acme/logstore/commit/0084262269f4e2cb94d04e0cc0d40e9666177f06))
* **operator:** Improve API documentation for schema version ([#13122](https://example.com/acme/logstore/issues/13122)) ([3a9f50f](https://example.com/acme/logstore/commit/3a9f50f5099a02e662b8ac10ddad0b36cd844161))
* **operator:** Use a minimum value for replay memory ceiling ([#13066](https://example.com/acme/logstore/issues/13066)) ([4f3ed77](https://example.com/acme/logstore/commit/4f3ed77cb92c2ffd605743237e609c28f7841728))
* Optimize regular initialization ([#12926](https://example.com/acme/logstore/issues/12926)) ([a46d14f](https://example.com/acme/logstore/commit/a46d14fb05ea14dd39095d2d71cd037acc2dfc51))
* **orFilters:** fix multiple or filters would get wrong filtertype ([#13169](https://example.com/acme/logstore/issues/13169)) ([9981e9e](https://example.com/acme/logstore/commit/9981e9e40d4eda1a88d1aee0483cec1c098b92c7))
* **otel:** Map 500 errors to 503 ([#13173](https://example.com/acme/logstore/issues/13173)) ([b31e04e](https://example.com/acme/logstore/commit/b31e04e3f1b7424cc52b518dc974a382a25bf045))
* **packaging:** Require online network in systemd unit file for Logstore and Promtail ([#12741](https://example.com/acme/logstore/issues/12741)) ([57f78b5](https://example.com/acme/logstore/commit/57f78b574ac9aa16f8322fb0edc4c7f0ec3cf6fa))
* panics when ingester response is nil ([#12946](https://example.com/acme/logstore/issues/12946)) ([3cc28aa](https://example.com/acme/logstore/commit/3cc28aaf0ec08373fb104327827e6a062807e7ff))
* promtail race fixes ([#12656](https://example.com/acme/logstore/issues/12656)) ([4e04d07](https://example.com/acme/logstore/commit/4e04d07168a8c5cb7086ced8486c6d584faa1045))
* promtail; clean up metrics generated from logs after a config reload. ([#11882](https://example.com/acme/logstore/issues/11882)) ([39a7181](https://example.com/acme/logstore/commit/39a7181a600e9dc848dd3c0b0163c07242a46278))
* **promtail:** Fix bug with Promtail config reloading getting stuck indefinitely ([#12795](https://example.com/acme/logstore/issues/12795)) ([4d761ac](https://example.com/acme/logstore/commit/4d761acd85b90cbdcafdf8d2547f0db14f6ae4dd))
* **promtail:** Fix UDP receiver on syslog transport ([#10708](https://example.com/acme/logstore/issues/10708)) ([a00f1f1](https://example.com/acme/logstore/commit/a00f1f1b0b8f536f2cdac2f8857eb40c716aa696))
* **promtail:** Handle docker logs when a log is split in multiple frames ([#12374](https://example.com/acme/logstore/issues/12374)) ([c0113db](https://example.com/acme/logstore/commit/c0113db4e8c4647188db6477d2ab265eda8dbb6c))
* properly return http status codes from ingester to querier for RPC function calls ([#13134](https://example.com/acme/logstore/issues/13134)) ([691b174](https://example.com/acme/logstore/commit/691b1741386716095a4926cea5d5bb53caa88d9a))
* **query sharding:** Generalize avg -> sum/count sharding using existing binop mapper ([#12599](https://example.com/acme/logstore/issues/12599)) ([11e7687](https://example.com/acme/logstore/commit/11e768726fb25f905de880ad2f5495b0f7fba156))
* **regression:** reverts acme/logstore[#13039](https://example.com/acme/logstore/issues/13039) to prevent use-after-free corruptions ([#13162](https://example.com/acme/logstore/issues/13162)) ([41c5ee2](https://example.com/acme/logstore/commit/41c5ee21fc80177b50e74515ca568223e86ae56a))
* Remove Hardcoded Bucket Name from EventBridge Example CloudFormation Template ([#12609](https://example.com/acme/logstore/issues/12609)) ([8c18463](https://example.com/acme/logstore/commit/8c18463285f214ba5b0b9a127bbe0071a2ec7d69))
* remove unneccessary disk panels for ssd read path ([#13014](https://example.com/acme/logstore/issues/13014)) ([8d9fb68](https://example.com/acme/logstore/commit/8d9fb68ae5d4f26ddc2ae184a1cb6a3b2a2c2127))
* remove unused parameter causing lint error ([#12801](https://example.com/acme/logstore/issues/12801)) ([33e82ec](https://example.com/acme/logstore/commit/33e82ec133b133e79666f7eec7d8d69954aa2aa3))
* **spans:** corrects early-close for a few spans ([#12887](https://example.com/acme/logstore/issues/12887)) ([93aaf29](https://example.com/acme/logstore/commit/93aaf29e681053a1d23dcf855cfe92af8415260d))
* temporarily moving from alloy -> alloy dev ([#13062](https://example.com/acme/logstore/issues/13062)) ([7ffe0fb](https://example.com/acme/logstore/commit/7ffe0fb6490e171e0100cb35ce6fde9377eff237))
* Track bytes discarded by ingester. ([#12981](https://example.com/acme/logstore/issues/12981)) ([88c6711](https://example.com/acme/logstore/commit/88c671162f70e075f6aa43599aa560fe7b4b5627))
* Update expected patterns when pruning ([#13079](https://example.com/acme/logstore/issues/13079)) ([2923a7d](https://example.com/acme/logstore/commit/2923a7d95818055a6ae9557d4b2f733b1af826f3))
* update to build image 0.33.2, fixes bug with promtail windows DNS resolution ([#12732](https://example.com/acme/logstore/issues/12732)) ([759f42d](https://example.com/acme/logstore/commit/759f42dd50bb4896f5e568691ef32245bb8fb25a))
* updated all dockerfiles go1.22 ([#12708](https://example.com/acme/logstore/issues/12708)) ([71a8f2c](https://example.com/acme/logstore/commit/71a8f2c2b11b419bd8c0af1f859671e5d8730448))
* Updated Logstore Otlp Ingest Configuration ([#12648](https://example.com/acme/logstore/issues/12648)) ([ff88f3c](https://example.com/acme/logstore/commit/ff88f3c3088a235eef5153a9d6414c161797a180))
* upgrade old plugin for the logstore-operational dashboard. ([#13016](https://example.com/acme/logstore/issues/13016)) ([d3c9cec](https://example.com/acme/logstore/commit/d3c9cec22891b45ed1cb93a9eacc5dad6a117fc5))
* Use an intermediate env variable in GH workflow ([#12905](https://example.com/acme/logstore/issues/12905)) ([772616c](https://example.com/acme/logstore/commit/772616cd8f5cbac70374dd4a53f1714fb49a7a3b))
* Use to the proper config names in warning messages ([#12114](https://example.com/acme/logstore/issues/12114)) ([4a05964](https://example.com/acme/logstore/commit/4a05964d5520d46d149f2a4e4709eee36c7fb418))
* **workflows:** don't run metric collector on forks ([#12687](https://example.com/acme/logstore/issues/12687)) ([7253444](https://example.com/acme/logstore/commit/72534449a07cd9f410973f2d01772024e8e4b7ba))


### Performance Improvements

* **blooms:** Resolve bloom blocks on index gateway and shard by block address ([#12720](https://example.com/acme/logstore/issues/12720)) ([5540c92](https://example.com/acme/logstore/commit/5540c92d50fe25356231e05995d24a7ca342084b))
* Improve Detected labels API ([#12816](https://example.com/acme/logstore/issues/12816)) ([e7fdeb9](https://example.com/acme/logstore/commit/e7fdeb974aff62c5775b9f98ebb2228000b28c8d))
* Introduce fixed size memory pool for bloom querier ([#13039](https://example.com/acme/logstore/issues/13039)) ([fc26431](https://example.com/acme/logstore/commit/fc264310ce64fc082965a5d7f036e45a5a399c61))
* Replace channel check with atomic bool in tailer.send() ([#12976](https://example.com/acme/logstore/issues/12976)) ([4a5edf1](https://example.com/acme/logstore/commit/4a5edf1a2af9e8af1842dc8d9b5482659d61031e))
* TSDB: Add fast-path to `inversePostingsForMatcher` ([#12679](https://example.com/acme/logstore/issues/12679)) ([402d1d7](https://example.com/acme/logstore/commit/402d1d7c48ab4eb77835f4ebb9ef7cabf1dd7449))


## [3.0.1](https://example.com/acme/logstore/compare/v3.0.0...v3.0.1) (2024-08-09)


### Bug Fixes

* **deps:** bumped dependencies versions to resolve CVEs ([#13833](https://example.com/acme/logstore/pull/13833)) ([e13011d](https://example.com/acme/logstore/commit/e13011d91a77501ca4f659df9cf33f23085d3a35))
* Fix nil pointer dereference in bloomstore initialisation ([#12869](https://example.com/acme/logstore/issues/12869)) ([167b468](https://example.com/acme/logstore/commit/167b468598bc70bbed6eed44826d3f9b85e1e0b8)), closes [#12270](https://example.com/acme/logstore/issues/12270)


## [3.0.0](https://example.com/acme/logstore/compare/v2.9.6...v3.0.0) (2024-04-08)

Starting with the 3.0 release we began using [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) and [release-please](https://github.com/googleapis/release-please) to generate the changelog. As a result the format has changed slightly from previous releases.

### Features

* **helm:** configurable API version for PodLog CRD ([#10812](https://example.com/acme/logstore/issues/10812)) ([d1dee91](https://example.com/acme/logstore/commit/d1dee9150b0e69941b2bd3ce4b23afead174ea29))
* **lambda/promtail:** support dropping labels ([#10755](https://example.com/acme/logstore/issues/10755)) ([ec54c72](https://example.com/acme/logstore/commit/ec54c723ebbeeda88000dde188d539ecfe05dad8))
* **logstash:** clients logstash output structured metadata support ([#10899](https://example.com/acme/logstore/issues/10899)) ([32f1ec2](https://example.com/acme/logstore/commit/32f1ec2fda5057732a2b20b98942aafec112c4ba))
* **logstore**: Allow custom usage trackers for ingested and discarded bytes metric. [11840](https://example.com/acme/logstore/pull/11840)
* **logstore**: feat: Support split align and caching for instant metric query results [11814](https://example.com/acme/logstore/pull/11814)
* **logstore**: Helm: Allow the definition of resources for AcmeAgent pods. [11851](https://example.com/acme/logstore/pull/11851)
* **logstore**: Ruler: Add the ability to disable the `X-Scope-OrgId` tenant identification header in remote write requests. [11819](https://example.com/acme/logstore/pull/11819)
* **logstore**: Add profiling integrations to tracing instrumentation. [11633](https://example.com/acme/logstore/pull/11633)
* **logstore**: Add a metrics.go log line for requests from querier to ingester [11571](https://example.com/acme/logstore/pull/11571)
* **logstore**: support GET for /ingester/shutdown [11477](https://example.com/acme/logstore/pull/11477)
* **logstore**: bugfix(memcached): Make memcached batch fetch truly context aware. [11363](https://example.com/acme/logstore/pull/11363)
* **logstore**: Helm: Add extraContainers to the write pods. [11319](https://example.com/acme/logstore/pull/11319)
* **logstore**: Inflight-logging: Add extra metadata to inflight requests logging. [11243](https://example.com/acme/logstore/pull/11243)
* **logstore**: Use metrics namespace for more metrics. [11025](https://example.com/acme/logstore/pull/11025).
* **logstore**: Change default of metrics.namespace. [11110](https://example.com/acme/logstore/pull/11110).
* **logstore**: Helm: Allow topologySpreadConstraints [11086](https://example.com/acme/logstore/pull/11086)
* **logstore**: Storage: Allow setting a constant prefix for all created keys [10096](https://example.com/acme/logstore/pull/10096)
* **logstore**: Remove already deprecated `store.max-look-back-period`. [11038](https://example.com/acme/logstore/pull/11038)
* **logstore**: Support Logstore ruler to notify WAL writes to remote storage. [10906](https://example.com/acme/logstore/pull/10906)
* **logstore**: Helm: allow AcmeAgent tolerations [10613](https://example.com/acme/logstore/pull/10613)
* **logstore**: Storage: remove signatureversionv2 from s3. [10295](https://example.com/acme/logstore/pull/10295)
* **logstore**: Dynamic client-side throttling to avoid object storage rate-limits (GCS only) [10140](https://example.com/acme/logstore/pull/10140)
* **logstore**: Removes already deprecated `-querier.engine.timeout` CLI flag and corresponding YAML setting as well as the `querier.query_timeout` YAML setting. [10302](https://example.com/acme/logstore/pull/10302)
* **logstore** Tracing: elide small traces for Stats call. [10308](https://example.com/acme/logstore/pull/10308)
* **logstore** Shard `avg_over_time` range aggregations. [10373](https://example.com/acme/logstore/pull/10373)
* **logstore** Remove deprecated config `-s3.sse-encryption` in favor or `-s3.sse.*` settings. [10377](https://example.com/acme/logstore/pull/10377)
* **logstore** Remove deprecated `ruler.wal-cleaer.period` [10378](https://example.com/acme/logstore/pull/10378)
* **logstore** Remove `experimental.ruler.enable-api` in favour of `ruler.enable-api` [10380](https://example.com/acme/logstore/pull/10380)
* **logstore** Remove deprecated `split_queries_by_interval` and `forward_headers_list` configuration options in the `query_range` section [10395](https://example.com/acme/logstore/pull/10395/)
* **logstore** Add `logstore_distributor_ingester_append_timeouts_total` metric, remove `logstore_distributor_ingester_append_failures_total` metric [10456](https://example.com/acme/logstore/pull/10456)
* **logstore** Remove configuration `use_boltdb_shipper_as_backup` [10534](https://example.com/acme/logstore/pull/10534)
* **logstore** Enable embedded cache if no other cache is explicitly enabled. [10620](https://example.com/acme/logstore/pull/10620)
* **logstore** Remove legacy ingester shutdown handler `/ingester/flush_shutdown`. [10655](https://example.com/acme/logstore/pull/10655)
* **logstore** Remove `ingester.max-transfer-retries` configuration option in favor of using the WAL. [10709](https://example.com/acme/logstore/pull/10709)
* **logstore** Deprecate write dedupe cache as this is not required by the newer single store indexes (tsdb and boltdb-shipper). [10736](https://example.com/acme/logstore/pull/10736)
* **logstore** Embedded cache: Updates the metric prefix from `querier_cache_` to `logstore_embeddedcache_` and removes duplicate metrics. [10693](https://example.com/acme/logstore/pull/10693)
* **logstore** Removes `shared_store` and `shared_store_key_prefix` from tsdb, boltdb shipper and compactor configs and their corresponding CLI flags. [10840](https://example.com/acme/logstore/pull/10840)
* **logstore** Config: Better configuration defaults to provide a better experience for users out of the box. [10793](https://example.com/acme/logstore/pull/10793)
* **logstore** Config: Removes `querier.worker-parallelism` and updates default value of `querier.max-concurrent` to 4. [10785](https://example.com/acme/logstore/pull/10785)
* **logstore** Add support for case-insensitive logql functions [10733](https://example.com/acme/logstore/pull/10733)
* **logstore** Native otlp ingestion support [10727](https://example.com/acme/logstore/pull/10727)
* Refactor to not use global logger in modules [11051](https://example.com/acme/logstore/pull/11051)
* **logstore** do not wrap requests but send pure Protobuf from frontend v2 via scheduler to querier when `-frontend.encoding=protobuf`. [10956](https://example.com/acme/logstore/pull/10956)
* **logstore** shard `quantile_over_time` range queries using probabilistic data structures. [10417](https://example.com/acme/logstore/pull/10417)
* **logstore** Config: Adds `frontend.max-query-capacity` to tune per-tenant query capacity. [11284](https://example.com/acme/logstore/pull/11284)
* **kaviraj,ashwanthgoli** Support caching /series and /labels query results [11539](https://example.com/acme/logstore/pull/11539)
* **logstore** Force correct memcached timeout when fetching chunks. [11545](https://example.com/acme/logstore/pull/11545)
* **logstore** Results Cache: Adds `query_length_served` cache stat to measure the length of the query served from cache. [11589](https://example.com/acme/logstore/pull/11589)
* **logstore** Query Frontend: Allow customisable splitting of queries which overlap the `query_ingester_within` window to reduce query pressure on ingesters. [11535](https://example.com/acme/logstore/pull/11535)
* **logstore** Cache: atomically check background cache size limit correctly. [11654](https://example.com/acme/logstore/pull/11654)
* **logstore** Metadata cache: Adds `frontend.max-metadata-cache-freshness` to configure the time window for which metadata results are not cached. This helps avoid returning inaccurate results by not caching recent results. [11682](https://example.com/acme/logstore/pull/11682)
* **logstore** Cache: extending #11535 to align custom ingester query split with cache keys for correct caching of results. [11679](https://example.com/acme/logstore/pull/11679)
* **logstore** otel: Add support for per tenant configuration for mapping otlp data to logstore format [11143](https://example.com/acme/logstore/pull/11143)
* **logstore** Config: Adds `frontend.log-query-request-headers` to enable logging of request headers in query logs. [11499](https://example.com/acme/logstore/pull/11284)
* **logstore** Ruler: Add support for filtering results of `/prometheus/api/v1/rules` endpoint by rule_name, rule_group, file and type. [11817](https://example.com/acme/logstore/pull/11817)
* **logstore** Metadata: Introduces a separate split interval of `split_recent_metadata_queries_by_interval` for `recent_metadata_query_window` to help with caching recent metadata query results. [11897](https://example.com/acme/logstore/pull/11897)
* **logstore** Ksonnet: Introduces memory limits to the compactor configuration to avoid unbounded memory usage. [11970](https://example.com/acme/logstore/pull/11897)
* **logstore** Memcached: Add mTLS support. [12318](https://example.com/acme/logstore/pull/12318)
* **logstore** Detect name of service emitting logs and add it as a label. [12392](https://example.com/acme/logstore/pull/12392)
* **logstore** LogQL: Introduces pattern match filter operators. [12398](https://example.com/acme/logstore/pull/12398)
* **logstore**: Helm: Use `/ingester/shutdown` for `preStop` hook in write pods. [11490](https://example.com/acme/logstore/pull/11490)
* **logstore** Upgrade thanos objstore, kit and other modules [10366](https://example.com/acme/logstore/pull/10366)
* **logstore** Upgrade thanos `objstore` [10451](https://example.com/acme/logstore/pull/10451)
* **logstore** Upgrade prometheus to v0.47.1 and kit [10814](https://example.com/acme/logstore/pull/10814)
* **logstore** introduce a backoff wait on subquery retries. [10959](https://example.com/acme/logstore/pull/10959)
* **logstore** Ensure all lifecycler cfgs ref a valid IPv6 addr and port combination [11121](https://example.com/acme/logstore/pull/11121)
* **logstore** Ensure the frontend uses a valid IPv6 addr and port combination [10650](https://example.com/acme/logstore/pull/10650)
* **logstore** Deprecate and flip `-legacy-read-mode` flag to `false` by default. [11665](https://example.com/acme/logstore/pull/11665)
* **logstore** BREAKING CHANGE: refactor how we do defaults for runtime overrides [12448](https://example.com/acme/logstore/pull/12448/files)
* **promtail**: structured_metadata: enable structured_metadata convert labels [10752](https://example.com/acme/logstore/pull/10752)
* **promtail**: chore(promtail): Improve default configuration that is shipped with rpm/deb packages to avoid possible high CPU utilisation if there are lots of files inside `/var/log`. [11511](https://example.com/acme/logstore/pull/11511)
* **promtail**: Lambda-Promtail: Add support for WAF logs in S3 [10416](https://example.com/acme/logstore/pull/10416)
* **promtail**: users can now define `additional_fields` in cloudflare configuration. [10301](https://example.com/acme/logstore/pull/10301)
* **promtail**: Lambda-Promtail: Add support for dropping labels passed via env var [10755](https://example.com/acme/logstore/pull/10755)

### Bug Fixes

* All lifecycler cfgs ref a valid IPv6 addr and port combination ([#11121](https://example.com/acme/logstore/issues/11121)) ([6385b19](https://example.com/acme/logstore/commit/6385b195739bd7d4e9706faddd0de663d8e5331a))
* **deps:** update github.com/c2h5oh/datasize digest to 859f65c (main) ([#10820](https://example.com/acme/logstore/issues/10820)) ([c66ffd1](https://example.com/acme/logstore/commit/c66ffd125cd89f5845a75a1751186fa46d003f70))
* **deps:** update github.com/docker/go-plugins-helpers digest to 6eecb7b (main) ([#10826](https://example.com/acme/logstore/issues/10826)) ([fb9c496](https://example.com/acme/logstore/commit/fb9c496b21be62f56866ae0f92440085e7860a2a))
* **deps:** update example.com/acme/gomemcache digest to 6947259 (main) ([#10836](https://example.com/acme/logstore/issues/10836)) ([2327789](https://example.com/acme/logstore/commit/2327789b5506d0ccc00d931195da17a2d47bf236))
* **deps:** update example.com/acme/logstore/pkg/push digest to 583aa28 (main) ([#10842](https://example.com/acme/logstore/issues/10842)) ([02d9418](https://example.com/acme/logstore/commit/02d9418270f4e615c1f78b0def635da7c0572ca4))
* **deps:** update example.com/acme/logstore/pkg/push digest to cfc4f0e (main) ([#10946](https://example.com/acme/logstore/issues/10946)) ([d27c4d2](https://example.com/acme/logstore/commit/d27c4d297dc6cce93ada98f16b962380ec933c6a))
* **deps:** update example.com/acme/logstore/pkg/push digest to e523809 (main) ([#11107](https://example.com/acme/logstore/issues/11107)) ([09cb9ae](https://example.com/acme/logstore/commit/09cb9ae76f4aef7dea477961c0c5424d7243bf2a))
* **deps:** update github.com/joncrlsn/dque digest to c2ef48c (main) ([#10947](https://example.com/acme/logstore/issues/10947)) ([1fe4885](https://example.com/acme/logstore/commit/1fe48858ae15b33646eedb85b05d6773a8bc5020))
* **deps:** update module google.golang.org/grpc [security] (main) ([#11031](https://example.com/acme/logstore/issues/11031)) ([0695424](https://example.com/acme/logstore/commit/0695424f7dd62435df3a9981276b40f3c5ef5641))
* **helm:** bump nginx-unprivilege to fix CVE ([#10754](https://example.com/acme/logstore/issues/10754)) ([dbf7dd4](https://example.com/acme/logstore/commit/dbf7dd4bac112a538a59907a8c6092504e7f4a91))
* **promtail:** correctly parse list of drop stage sources from YAML ([#10848](https://example.com/acme/logstore/issues/10848)) ([f51ee84](https://example.com/acme/logstore/commit/f51ee849b03c5f6b79f3e93cb7fd7811636bede2))
* **promtail:** prevent panic due to duplicate metric registration after reloaded ([#10798](https://example.com/acme/logstore/issues/10798)) ([47e2c58](https://example.com/acme/logstore/commit/47e2c5884f443667e64764f3fc3948f8f11abbb8))
* **logstore:** respect query matcher in ingester when getting label values ([#10375](https://example.com/acme/logstore/issues/10375)) ([85e2e52](https://example.com/acme/logstore/commit/85e2e52279ecac6dc111d5c113c54d6054d2c922))
* **helm:** Sidecar configuration for Backend ([#10603](https://example.com/acme/logstore/issues/10603)) ([c29ba97](https://example.com/acme/logstore/commit/c29ba973a0b5b7b59613d210b741d5a547ea0e83))
* **tools/lambda-promtail:** Do not evaluate empty string for drop_labels ([#11074](https://example.com/acme/logstore/issues/11074)) ([94169a0](https://example.com/acme/logstore/commit/94169a0e6b5bf96426ad21e40f9583b721f35d6c))
* **lambda-promtail** Fix panic in lambda-promtail due to mishandling of empty DROP_LABELS env var. [11074](https://example.com/acme/logstore/pull/11074)
* **logstore** Generate tsdb_shipper storage_config even if using_boltdb_shipper is false [11195](https://example.com/acme/logstore/pull/11195)
* **promtail**: Fix Promtail excludepath not evaluated on newly added files. [9831](https://example.com/acme/logstore/pull/9831)
* **logstore** Do not reflect label names in request metrics' "route" label. [11551](https://example.com/acme/logstore/pull/11551)
* **logstore** Fix duplicate logs from docker containers. [11563](https://example.com/acme/logstore/pull/11563)
* **logstore** Ruler: Fixed a panic that can be caused by concurrent read-write access of tenant configs when there are a large amount of rules. [11601](https://example.com/acme/logstore/pull/11601)
* **logstore** Fixed regression adding newlines to HTTP error response bodies which may break client integrations. [11606](https://example.com/acme/logstore/pull/11606)
* **logstore** Log results cache: compose empty response based on the request being served to avoid returning incorrect limit or direction. [11657](https://example.com/acme/logstore/pull/11657)
* **logstore** Fix semantics of label parsing logic of metrics and logs queries. Both only parse the first label if multiple extractions into the same label are requested. [11587](https://example.com/acme/logstore/pull/11587)
* **logstore** Background Cache: Fixes a bug that is causing the background queue size to be incremented twice for each enqueued item. [11776](https://example.com/acme/logstore/pull/11776)
* **logstore**: Parsing: String array elements were not being parsed correctly in JSON processing [11921](https://example.com/acme/logstore/pull/11921)


## [2.9.10](https://example.com/acme/logstore/compare/v2.9.9...v2.9.10) (2024-08-09)


### Bug Fixes

* Update dependencies versions to remove CVE ([#13835](https://example.com/acme/logstore/pull/13835)) ([567bef2](https://example.com/acme/logstore/commit/567bef286376663407c54f5da07fa00963ba5485))


## [2.9.9](https://example.com/acme/logstore/compare/v2.9.8...v2.9.9) (2024-07-04)

### All Changes

#### Logstore

##### Fixes

* [12925](https://example.com/acme/logstore/pull/12925) **grobinson-acme** Ingester: Add ingester_chunks_flush_failures_total
* [13140](https://example.com/acme/logstore/pull/13140) **grobinson-acme** Ingester: Add backoff to flush op

## [2.9.8](https://example.com/acme/logstore/compare/v2.9.7...v2.9.8) (2024-05-03)

### All Changes

#### Logstore

##### Fixes

* update module golang.org/x/net to v0.23.0 [security] (release-2.9.x) ([#12865](https://example.com/acme/logstore/issues/12865)) ([94e0029](https://example.com/acme/logstore/commit/94e00299ec9b36ad97c147641566b6922268c54e))

## [2.9.7](https://example.com/acme/logstore/compare/v2.9.6...v2.9.7) (2024-04-10)

### Bug Fixes

* Bump go to 1.21.9 and build image to 0.33.1 ([#12542](https://example.com/acme/logstore/issues/12542)) ([efc4d2f](https://example.com/acme/logstore/commit/efc4d2f009e04ecb1db58a637b89b33aa234de34))

## [2.9.6](https://example.com/acme/logstore/compare/v2.9.5...v2.9.6) (2024-03-21)

### Bug Fixes

* promtail failures connecting to local logstore installation [release-2.9.x]  ([#12184](https://example.com/acme/logstore/issues/12184)) ([8585e35](https://example.com/acme/logstore/commit/8585e3537375c0deb11462d7256f5da23228f5e1))
* **release-2.9.x:** frontend: Use `net.JoinHostPort` to support IPv6 addresses ([#10650](https://example.com/acme/logstore/issues/10650)) ([#11870](https://example.com/acme/logstore/issues/11870)) ([7def3b4](https://example.com/acme/logstore/commit/7def3b4e774252e13ba154ca13f72816a84da7dd))
* update google.golang.org/protobuf to v1.33.0 ([#12269](https://example.com/acme/logstore/issues/12269)) ([#12287](https://example.com/acme/logstore/issues/12287)) ([3186520](https://example.com/acme/logstore/commit/318652035059fdaa40405f263fc9e37b4d38b157))

## [2.9.5](https://example.com/acme/logstore/compare/v2.9.4...v2.9.5) (2024-02-28)

##### Changes

* [10677](https://example.com/acme/logstore/pull/10677) **chaudum** Remove deprecated `stream_lag_labels` setting from both the `options` and `client` configuration sections.
* [10689](https://example.com/acme/logstore/pull/10689) **dylanguedes**: Ingester: Make jitter to be 20% of flush check period instead of 1%.
* [11420](https://example.com/acme/logstore/pull/11420) **zry98**: Show a clearer reason in "disable watchConfig" log message when server is disabled.

##### Fixes

* [10708](https://example.com/acme/logstore/pull/10708) **joshuapare**: Fix UDP receiver on syslog transport
* [10631](https://example.com/acme/logstore/pull/10631) **thampiotr**: Fix race condition in cleaning up metrics when stopping to tail files.
* [10798](https://example.com/acme/logstore/pull/10798) **hainenber**: Fix agent panicking after reloaded due to duplicate metric collector registration.
* [10848](https://example.com/acme/logstore/pull/10848) **rgroothuijsen**: Correctly parse list of drop stage sources from YAML.

#### LogCLI

* [11852](https://example.com/acme/logstore/pull/11852) **MichelHollands**: feat: update logcli so it tries to load the latest version of the schemaconfig

#### Mixins

* [11087](https://example.com/acme/logstore/pull/11087) **JoaoBraveCoding**: Adds structured metadata panels for ingested data
* [11637](https://example.com/acme/logstore/pull/11637) **JoaoBraveCoding**: Add route to write Distributor Latency dashboard

#### Fixes

#### FluentD

#### Jsonnet

* [11312](https://example.com/acme/logstore/pull/11312) **sentoz**: Logstore ksonnet: Do not generate configMap for consul if you are using memberlist

* [11020](https://example.com/acme/logstore/pull/11020) **ashwanthgoli**: Logstore ksonnet: Do not generate table-manager manifests if shipper store is in-use.

* [10784](https://example.com/acme/logstore/pull/10894) **slim-bean** Update index gateway client to use a headless service.

* [10542](https://example.com/acme/logstore/pull/10542) **chaudum**: Remove legacy deployment mode for ingester (Deployment, without WAL) and instead always run them as StatefulSet.

## [2.8.11](https://example.com/acme/logstore/compare/v2.8.10...v2.8.11) (2024-03-22)

### Bug Fixes

* update google.golang.org/protobuf to v1.33.0 ([#12276](https://example.com/acme/logstore/issues/12276)) ([3c05724](https://example.com/acme/logstore/commit/3c05724ac9d7ea9b6048c6e67cd13dc55fa72782))

## [2.8.10](https://example.com/acme/logstore/compare/v2.8.9...v2.8.10) (2024-02-28)

### Bug Fixes

* image tag from env and pin release to v1.11.5 ([#12073](https://example.com/acme/logstore/issues/12073)) ([8e11cd7](https://example.com/acme/logstore/commit/8e11cd7a8222a64d60bff30a41e399ddbda3372e))

## [2.8.9](https://example.com/acme/logstore/compare/v2.8.8...v2.8.9) (2024-02-23)

### Bug Fixes

* bump alpine base image and go to fix CVEs ([#12026](https://example.com/acme/logstore/issues/12026)) ([196650e](https://example.com/acme/logstore/commit/196650e4c119249016df85a50a2cced521cbe9be))

## 2.9.2 (2023-10-16)

### All Changes

##### Security

* [10879](https://example.com/acme/logstore/pull/10879) **DylanGuedes**: Upgrade golang.org/x/net to v0.17.0 to patch CVE-2023-39325 / CVE-2023-44487
* [10871](https://example.com/acme/logstore/pull/10871) **ashwanthgoli**: Upgrade go to v1.21.3 and grpc-go to v1.56.3 to patch CVE-2023-39325 / CVE-2023-44487

## 2.9.1 (2023-09-14)

### All Changes

#### Logstore

##### Security

* [10573](https://example.com/acme/logstore/pull/10573) **DylanGuedes**: Bump Docker base images to Alpine version 3.18.3 to mitigate CVE-2022-48174

##### Fixes

* [10585](https://example.com/acme/logstore/pull/10585) **ashwanthgoli** / **chaudum**: Fix bug in index object client that could result in not showing all ingested logs in query results.
* [10314](https://example.com/acme/logstore/pull/10314) **bboreham**: Fix race conditions in indexshipper.

## 2.9.0 (2023-09-06)

### All Changes

##### Security

* [10188](https://example.com/acme/logstore/pull/10188) **shantanualsi**: Bump alpine version from 3.16.5 -> 3.16.7

#### Logstore

##### Enhancements

* [10101](https://example.com/acme/logstore/pull/10101) **owen-d**: Sharding optimizations and fix bug on `<aggr> by|without ()` groupings which removed the grouping while downstreaming
* [10324](https://example.com/acme/logstore/pull/10324) **ashwanthgoli**: Deprecate ingester.unordered-writes and a few unused configs(log.use-buffered, log.use-sync, frontend.forward-headers-list)
* [10322](https://example.com/acme/logstore/pull/10322) **chaudum**: Deprecate misleading setting `-ruler.evaluation-delay-duration`.
* [10295](https://example.com/acme/logstore/pull/10295) **changhyuni**: Storage: remove signatureversionv2 from s3.
* [10109](https://example.com/acme/logstore/pull/10109) **vardhaman-surana**: Ruler: add limit parameter in rulegroup
* [10187](https://example.com/acme/logstore/pull/10187) **roelarents**: Add connection-string option for Azure Blob Storage.
* [9621](https://example.com/acme/logstore/pull/9621) **DylanGuedes**: Introduce TSDB postings cache.
* [10010](https://example.com/acme/logstore/pull/10010) **rasta-rocket**: feat(promtail): retrieve BotTags field from cloudflare
* [9995](https://example.com/acme/logstore/pull/9995) **chaudum**: Add jitter to the flush interval to prevent multiple ingesters to flush at the same time.
* [9797](https://example.com/acme/logstore/pull/9797) **chaudum**: Add new `logstore_index_gateway_requests_total` counter metric to observe per-tenant RPS
* [9710](https://example.com/acme/logstore/pull/9710) **chaudum**: Add shuffle sharding to index gateway
* [9573](https://example.com/acme/logstore/pull/9573) **CCOLLOT**: Lambda-Promtail: Add support for AWS CloudFront log ingestion.
* [9497](https://example.com/acme/logstore/pull/9497) **CCOLLOT**: Lambda-Promtail: Add support for AWS CloudTrail log ingestion.
* [8886](https://example.com/acme/logstore/pull/8886) **MichelHollands**: Add new logql template function `unixToTime`
* [8067](https://example.com/acme/logstore/pull/9497) **CCOLLOT**: Lambda-Promtail: Add support for AWS CloudTrail log ingestion.
* [9515](https://example.com/acme/logstore/pull/9515) **MichelHollands**: Fix String() on vector aggregation LogQL expressions that contain `without ()`.
* [8067](https://example.com/acme/logstore/pull/8067) **DylanGuedes**: Distributor: Add auto-forget unhealthy members support.
* [9175](https://example.com/acme/logstore/pull/9175) **MichelHollands**: Ingester: update the `prepare_shutdown` endpoint so it supports GET and DELETE and stores the state on disk.
* [8953](https://example.com/acme/logstore/pull/8953) **dannykopping**: Querier: block queries by hash.
* [8851](https://example.com/acme/logstore/pull/8851) **jeschkies**: Introduce limit to require a set of labels for selecting streams.
* [9016](https://example.com/acme/logstore/pull/9016) **kavirajk**: Change response type of `format_query` handler to `application/json`
* [8972](https://example.com/acme/logstore/pull/8972) **salvacorts** Index stat requests are now cached in the results cache.
* [9177](https://example.com/acme/logstore/pull/9177) **salvacorts** Index stat cache can be enabled or disabled with the new `cache_index_stats_results` flag. Disabled by default.
* [9096](https://example.com/acme/logstore/pull/9096) **salvacorts**: Compute proportional TSDB index stats for chunks that doesn't fit fully in the queried time range.
* [8939](https://example.com/acme/logstore/pull/8939) **Suruthi-G-K**: Logstore: Add support for trusted profile authentication in COS client.
* [8852](https://example.com/acme/logstore/pull/8852) **wtchangdm**: Logstore: Add `route_randomly` to Redis options.
* [8848](https://example.com/acme/logstore/pull/8848) **dannykopping**: Ruler: add configurable rule evaluation jitter.
* [8826](https://example.com/acme/logstore/pull/8826) **amankrsingh2000**: Logstore: Add support for IBM cloud object storage as storage client.
* [8752](https://example.com/acme/logstore/pull/8752) **chaudum**: Add query fairness control across actors within a tenant to scheduler, which can be enabled by passing the `X-Logstore-Actor-Path` header to the HTTP request of the query.
* [8786](https://example.com/acme/logstore/pull/8786) **DylanGuedes**: Ingester: add new /ingester/prepare_shutdown endpoint.
* [8744](https://example.com/acme/logstore/pull/8744) **dannykopping**: Ruler: remote rule evaluation.
* [8670](https://example.com/acme/logstore/pull/8670) **salvacorts** Introduce two new limits to refuse log and metric queries that would read too much data.
* [8918](https://example.com/acme/logstore/pull/8918) **salvacorts** Introduce limit to require at least a number label matchers on metric and log queries.
* [8909](https://example.com/acme/logstore/pull/8909) **salvacorts** Requests to `/logstore/api/v1/index/stats` are split in 24h intervals.
* [8732](https://example.com/acme/logstore/pull/8732) **abaguas**: azure: respect retry config before cancelling the context
* [9206](https://example.com/acme/logstore/pull/9206) **dannykopping**: Ruler: log rule evaluation detail.
* [9184](https://example.com/acme/logstore/pull/9184) **periklis**: Bump kit to introduce IPv6 support for memberlist
* [9357](https://example.com/acme/logstore/pull/9357) **Indransh**: Add HTTP API to change the log level at runtime
* [9431](https://example.com/acme/logstore/pull/9431) **dannykopping**: Add more buckets to `logstore_memcache_request_duration_seconds` metric; latencies can increase if using memcached with NVMe
* [8684](https://example.com/acme/logstore/pull/8684) **oleksii-boiko-ua**: Helm: Add hpa templates for read, write and backend components.
* [9535](https://example.com/acme/logstore/pull/9535) **salvacorts** Index stats cache can be configured independently of the results cache. If it's not configured, but it's enabled, it will use the results cache configuration.
* [9626](https://example.com/acme/logstore/pull/9626) **ashwanthgoli** logfmt: add --strict flag to enable strict parsing, perform nostrict parsing by default
* [9672](https://example.com/acme/logstore/pull/9672) **zeitlinger**: Add `alignLeft` and `alignRight` line formatting functions.
* [9693](https://example.com/acme/logstore/pull/9693) **salvacorts** Add `keep` stage to LogQL.
* [7447](https://example.com/acme/logstore/pull/7447) **ashwanthgoli** compactor: multi-store support.
* [7754](https://example.com/acme/logstore/pull/7754) **ashwanthgoli** index-shipper: add support for multiple stores.
* [9813](https://example.com/acme/logstore/pull/9813) **jeschkies**: Enable Protobuf encoding via content negotiation between querier and query frontend.
* [10281](https://example.com/acme/logstore/pull/10281) **dannykopping**: Track effectiveness of hedged requests.
* [10341](https://example.com/acme/logstore/pull/10341) **ashwanthgoli** Deprecate older index types and non-object stores - `aws-dynamo, gcp, gcp-columnkey, bigtable, bigtable-hashed, cassandra, grpc`
* [10344](https://example.com/acme/logstore/pull/10344) **ashwanthgoli**  Compactor: deprecate `-boltdb.shipper.compactor.` prefix in favor of `-compactor.`.
* [10073](https://example.com/acme/logstore/pull/10073) **sandeepsukhani,salvacorts,vlad-diachenko** Support attaching structured metadata to log lines.
* [11151](https://example.com/acme/logstore/pull/11151) **ashwanthgoli**: Removes already deprecated configs: `ruler.evaluation-delay-duration`, `boltdb.shipper.compactor.deletion-mode`, `validation.enforce-metric-name` and flags with prefix `-boltdb.shipper.compactor.*`.

##### Fixes

* [10026](https://example.com/acme/logstore/pull/10026) **aminesnow**: Add support for Alibaba Cloud as storage backend for the ruler.
* [10065](https://example.com/acme/logstore/pull/10065) **fgouteroux**: Fix the syntax error message when parsing expression rule.
* [8979](https://example.com/acme/logstore/pull/8979) **slim-bean**: Fix the case where a logs query with start time == end time was returning logs when none should be returned.
* [9099](https://example.com/acme/logstore/pull/9099) **salvacorts**: Fix the estimated size of chunks when writing a new TSDB file during compaction.
* [9130](https://example.com/acme/logstore/pull/9130) **salvacorts**: Pass LogQL engine options down to the _split by range_, _sharding_, and _query size limiter_ middlewares.
* [9252](https://example.com/acme/logstore/pull/9252) **jeschkies**: Use un-escaped regex literal for string matching.
* [9176](https://example.com/acme/logstore/pull/9176) **DylanGuedes**: Fix incorrect association of per-stream rate limit when sharding is enabled.
* [9463](https://example.com/acme/logstore/pull/9463) **Totalus**: Fix OpenStack Swift client object listing to fetch all the objects properly.
* [9495](https://example.com/acme/logstore/pull/9495) **thampiotr**: Promtail: Fix potential goroutine leak in file tailer.
* [9650](https://example.com/acme/logstore/pull/9650) **ashwanthgoli**: Config: ensure storage config defaults apply to named stores.
* [9757](https://example.com/acme/logstore/pull/9757) **sandeepsukhani**: Frontend Caching: Fix a bug in negative logs results cache causing Logstore to unexpectedly send empty/incorrect results.
* [9754](https://example.com/acme/logstore/pull/9754) **ashwanthgoli**: Fixes an issue with indexes becoming unqueriable if the index prefix is different from the one configured in the latest period config.
* [9763](https://example.com/acme/logstore/pull/9763) **ssncferreira**: Fix the logic of the `offset` operator for downstream queries on instant query splitting of (range) vector aggregation expressions containing an offset.
* [9773](https://example.com/acme/logstore/pull/9773) **ssncferreira**: Fix instant query summary statistic's `splits` corresponding to the number of subqueries a query is split into based on `split_queries_by_interval`.
* [9949](https://example.com/acme/logstore/pull/9949) **masslessparticle**: Fix pipelines to clear caches when tailing to avoid resource exhaustion.
* [9936](https://example.com/acme/logstore/pull/9936) **masslessparticle**: Fix the way query stages are reordered when `unpack` is present.
* [10309](https://example.com/acme/logstore/pull/10309) **akhilanarayanan**: Fix race condition in series index store.
* [10221](https://example.com/acme/logstore/pull/10221) **periklis**: Allow using the forget button when access via the internal server

##### Changes

* [9857](https://example.com/acme/logstore/pull/9857) **DylanGuedes**: Stop emitting spans for every `AWS.S3` or `Azure.Blob` call.
* [9212](https://example.com/acme/logstore/pull/9212) **trevorwhitney**: Rename UsageReport to Analytics. The only external impact of this change is a change in the `-list-targets` output.

#### Promtail

##### Enhancements

* [8474](https://example.com/acme/logstore/pull/8787) **andriikushch**: Promtail: Add a new target for the Azure Event Hubs
* [8874](https://example.com/acme/logstore/pull/8874) **rfratto**: Promtail: Support expoential backoff when polling unchanged files for logs.
* [9508](https://example.com/acme/logstore/pull/9508) **farodin91**: Promtail: improve behavior of partial lines.
* [9986](https://example.com/acme/logstore/pull/9986) **vlad-diachenko**: Promtail: Add `structured_metadata` stage to attach metadata to each log line.

##### Fixes

* [8987](https://example.com/acme/logstore/pull/8987) **darxriggs**: Promtail: Fix file descriptor leak.
* [9863](https://example.com/acme/logstore/pull/9863) **ashwanthgoli**: Promtail: Apply defaults to HTTP client config. This ensures follow_redirects is set to true.
* [9915](https://example.com/acme/logstore/pull/9915) **frittentheke**: Promtail: Update acme/tail to address issue in retry logic

#### LogCLI

##### Fixes

* [9597](https://example.com/acme/logstore/pull/9597) **vlad-diachenko**: Set TSDB shipper mode to ReadOnly and disabled indexGatewayClient during local query run and changed index downloading timeout from `5s` to `1m`.
* [8566](https://example.com/acme/logstore/pull/8566) **ndrpnt**: Allow queries to start with negative filters (`!=` and `!~`) when omitting stream selector with `--stdin` flag

#### Mixins

#### Enhancements

#### Fixes

* [9684](https://example.com/acme/logstore/pull/9684) **thampiotr**: Mixins: Fix promtail cluster template not finding all clusters.
* [8995](https://example.com/acme/logstore/pull/8995) **dannykopping**: Mixins: Fix Jsonnet `RUNTIME ERROR` that occurs when you try to use the mixins with `use_boltdb_shipper: false`.

#### FluentD

##### Enhancements

* [LOG-4012](https://issues.redhat.com/browse/LOG-4012) **jcantril**: fluent-plugin-grapha-logstore: Add config to support tls: ciphers, min_versio

#### Jsonnet

* [9790](https://example.com/acme/logstore/pull/9790) **manohar-koukuntla**: Add TSDB equivalent of `use_boltdb_shipper` flag to be able to configure `tsdb_shipper` section.
* [8855](https://example.com/acme/logstore/pull/8855) **JoaoBraveCoding**: Add gRPC port to logstore compactor mixin
* [8880](https://example.com/acme/logstore/pull/8880) **JoaoBraveCoding**: Normalize headless service name for query-frontend/scheduler
* [9978](https://example.com/acme/logstore/pull/9978) ****vlad-diachenko****: replaced deprecated `policy.v1beta1` with `policy.v1`.

## 2.8.6 (2023-10-17)

#### Logstore

##### Security

* [10887](https://example.com/acme/logstore/pull/10887) upgrade go-grpc to v1.56.3 and golang.org/x/net to v0.17.0 to patch CVE-2023-39325 / CVE-2023-44487
* [10889](https://example.com/acme/logstore/pull/10889) upgrade go to v1.20.10 to patch CVE-2023-39325 / CVE-2023-44487

## 2.8.5 (2023-09-14)

#### Logstore

##### Security

* [10573](https://example.com/acme/logstore/pull/10573) **DylanGuedes**: Bump Docker base images to Alpine version 3.18.3 to mitigate CVE-2022-48174

## 2.8.3 (2023-07-21)

#### Logstore

##### Security

* [9913](https://example.com/acme/logstore/pull/9913) **MichelHollands**: Upgrade go version to 1.20.6

##### Enhancements

* [9604](https://example.com/acme/logstore/pull/9604) **dannykopping**: Querier: configurable writeback queue bytes size

##### Fixes

* [9471](https://example.com/acme/logstore/pull/9471) **sandeepsukhani**: query-scheduler: fix query distribution in SSD mode.
* [9629](https://example.com/acme/logstore/pull/9629) **periklis**: Fix duplicate label values from ingester streams.

#### Promtail

##### Fixes

* [9155](https://example.com/acme/logstore/pull/9155) **farodin91**: Promtail: Break on iterate journal failure.
* [8988](https://example.com/acme/logstore/pull/8988) **darxriggs**: Promtail: Prevent logging errors on normal shutdown.

## 2.8.2 (2023-05-03)

#### Logstore

##### Security

* [9370](https://example.com/acme/logstore/pull/9370) **dannykopping**: upgrade to go1.20.4

#### Promtail

##### Enhancements

* [8994](https://example.com/acme/logstore/pull/8994) **DylanGuedes**: Promtail: Add new `decompression` configuration to customize the decompressor behavior.

## 2.8.1 (2023-04-24)

#### Logstore

##### Fixes

* [9156](https://example.com/acme/logstore/pull/9156) **ashwanthgoli**: Expiration: do not drop index if period is a zero value.
* [8971](https://example.com/acme/logstore/pull/8971) **dannykopping**: Stats: fix `Cache.Chunk.BytesSent` statistic and logstore_chunk_fetcher_fetched_size_bytes metric with correct chunk size.
* [9185](https://example.com/acme/logstore/pull/9185) **dannykopping**: Prevent redis client from incorrectly choosing cluster mode with local address.

##### Changes

* [9106](https://example.com/acme/logstore/pull/9106) **trevorwhitney**: Update go to 1.20.3.

##### Build

* [9264](https://example.com/acme/logstore/pull/9264) **trevorwhitney**: Update build and other docker image to alpine 3.16.5.

#### Promtail

##### Fixes

* [9095](https://example.com/acme/logstore/pull/9095) **JordanRushing** Fix journald support in amd64 binary build.

## 2.8.0 (2023-04-04)

#### Logstore

##### Enhancements

* [8824](https://example.com/acme/logstore/pull/8824) **periklis**: Expose optional label matcher for label values handler
* [8727](https://example.com/acme/logstore/pull/8727) **cstyan** **jeschkies**: Propagate per-request limit header to querier.
* [8682](https://example.com/acme/logstore/pull/8682) **dannykopping**: Add fetched chunk size distribution metric `logstore_chunk_fetcher_fetched_size_bytes`.
* [8532](https://example.com/acme/logstore/pull/8532) **justcompile**: Adds Storage Class option to S3 objects
* [7951](https://example.com/acme/logstore/pull/7951) **MichelHollands**: Add a count template function to line_format and label_format.
* [7380](https://example.com/acme/logstore/pull/7380) **liguozhong**: metrics query: range vector support streaming agg when no overlap.
* [7731](https://example.com/acme/logstore/pull/7731) **bitkill**: Add healthchecks to the docker-compose example.
* [7759](https://example.com/acme/logstore/pull/7759) **kavirajk**: Improve error message for loading config with ENV variables.
* [7785](https://example.com/acme/logstore/pull/7785) **dannykopping**: Add query blocker for queries and rules.
* [7817](https://example.com/acme/logstore/pull/7817) **kavirajk**: fix(memcached): panic on send on closed channel.
* [7916](https://example.com/acme/logstore/pull/7916) **ssncferreira**: Add `doc-generator` tool to generate configuration flags documentation.
* [7964](https://example.com/acme/logstore/pull/7964) **slim-bean**: Add a `since` query parameter to allow querying based on relative time.
* [7989](https://example.com/acme/logstore/pull/7989) **liguozhong**: logql support `sort` and `sort_desc`.
* [7997](https://example.com/acme/logstore/pull/7997) **kavirajk**: fix(promtail): Fix cri tags extra new lines when joining partial lines
* [7975](https://example.com/acme/logstore/pull/7975) **adityacs**: Support drop labels in logql
* [7946](https://example.com/acme/logstore/pull/7946) **ashwanthgoli** config: Add support for named stores
* [8027](https://example.com/acme/logstore/pull/8027) **kavirajk**: chore(promtail): Make `batchwait` and `batchsize` config explicit with yaml tags
* [7978](https://example.com/acme/logstore/pull/7978) **chaudum**: Shut down query frontend gracefully to allow inflight requests to complete.
* [8047](https://example.com/acme/logstore/pull/8047) **bboreham**: Dashboards: add k8s resource requests to CPU and memory panels.
* [8061](https://example.com/acme/logstore/pull/8061) **kavirajk**: Remove circle from Logstore OSS
* [8092](https://example.com/acme/logstore/pull/8092) **dannykopping**: add rule-based sharding to ruler.
* [8131](https://example.com/acme/logstore/pull/8131) **jeschkies**: Compile Promtail ARM and ARM64 with journald support.
* [8212](https://example.com/acme/logstore/pull/8212) **kavirajk**: ingester: Add `ingester_memory_streams_labels_bytes metric` for more visibility of size of metadata of in-memory streams.
* [8271](https://example.com/acme/logstore/pull/8271) **kavirajk**: logql: Support urlencode and urldecode template functions
* [8259](https://example.com/acme/logstore/pull/8259) **mar4uk**: Extract push.proto from the logproto package to the separate module.
* [7906](https://example.com/acme/logstore/pull/7906) **kavirajk**: Add API endpoint that formats LogQL expressions and support new `fmt` subcommand in `logcli` to format LogQL query.
* [6675](https://example.com/acme/logstore/pull/6675) **btaani**: Add logfmt expression parser for selective extraction of labels from logfmt formatted logs
* [8474](https://example.com/acme/logstore/pull/8474) **farodin91**: Add support for short-lived S3 session tokens
* [8774](https://example.com/acme/logstore/pull/8774) **slim-bean**: Add new logql template functions `bytes`, `duration`, `unixEpochMillis`, `unixEpochNanos`, `toDateInZone`, `b64Enc`, and `b64Dec`

##### Fixes

* [7784](https://example.com/acme/logstore/pull/7784) **isodude**: Fix default values of connect addresses for compactor and querier workers to work with IPv6.
* [7880](https://example.com/acme/logstore/pull/7880) **sandeepsukhani**: consider range and offset in queries while looking for schema config for query sharding.
* [7937](https://example.com/acme/logstore/pull/7937) **ssncferreira**: Deprecate CLI flag `-ruler.wal-cleaer.period` and replace it with `-ruler.wal-cleaner.period`.
* [7966](https://example.com/acme/logstore/pull/7966) **sandeepsukhani**: Fix query-frontend request load balancing when using k8s service.
* [8251](https://example.com/acme/logstore/pull/8251) **sandeepsukhani** index-store: fix indexing of chunks overlapping multiple schemas.
* [8151](https://example.com/acme/logstore/pull/8151) **sandeepsukhani** fix log deletion with line filters.
* [8448](https://example.com/acme/logstore/pull/8448) **chaudum**: Fix bug in LogQL parser that caused certain queries that contain a vector expression to fail.
* [8775](https://example.com/acme/logstore/pull/8755) **sandeepsukhani**: index-gateway: fix failure in initializing index gateway when boltdb-shipper is not being used.
* [8448](https://example.com/acme/logstore/pull/8665) **sandeepsukhani**: deletion: fix issue in processing delete requests with tsdb index
* [8753](https://example.com/acme/logstore/pull/8753) **slim-bean** A zero value for retention_period will now disable retention.
* [8959](https://example.com/acme/logstore/pull/8959) **periklis**: Align common instance_addr with memberlist advertise_addr

##### Changes

* [8315](https://example.com/acme/logstore/pull/8315) **thepalbi** Relicense and export `pkg/ingester` WAL code to be used in Promtail's WAL.
* [8761](https://example.com/acme/logstore/pull/8761) **slim-bean** Remove "subqueries" from the metrics.go log line and instead provide `splits` and `shards`
* [8887](https://example.com/acme/logstore/issues/8887) **3deep5me** Helm: Removed support for PodDisruptionBudget in policy/v1alpha1 and upgraded it to policy/v1.

##### Build

#### Promtail

##### Enhancements

* [8231](https://example.com/acme/logstore/pull/8231) **CCOLLOT**: Lambda-promtail: add support for AWS SQS message ingestion.
* [7619](https://example.com/acme/logstore/pull/7619) **cadrake**: Add ability to pass query params to heroku drain targets for relabelling.
* [7973](https://example.com/acme/logstore/pull/7973) **chodges15**: Add configuration to drop rate limited batches in Logstore client and new metric label for drop reason.
* [8153](https://example.com/acme/logstore/pull/8153) **kavirajk**: promtail: Add `max-line-size` limit to drop on client side
* [8096](https://example.com/acme/logstore/pull/8096) **kavirajk**: doc(promtail): Doc about how log rotate works with promtail
* [8233](https://example.com/acme/logstore/pull/8233) **nicoche**: promtail: Add `max-line-size-truncate` limit to truncate too long lines on client side
* [7462](https://example.com/acme/logstore/pull/7462) **MarNicGit**: Allow excluding event message from Windows Event Log entries.
* [7597](https://example.com/acme/logstore/pull/7597) **redbaron**: allow ratelimiting by label
* [3493](https://example.com/acme/logstore/pull/3493) **adityacs** Support geoip stage.
* [8382](https://example.com/acme/logstore/pull/8382) **kelnage**: Promtail: Add event log message stage

##### Fixes

* [8231](https://example.com/acme/logstore/pull/8231) **CCOLLOT**: Lambda-promtail: fix flushing behavior of batches, leading to a significant increase in performance.

##### Changes

#### LogCLI

##### Enhancement

* [8413](https://example.com/acme/logstore/pull/8413) **chaudum**: Try to load tenant-specific `schemaconfig-{orgID}.yaml` when using `--remote-schema` argument and fallback to global `schemaconfig.yaml`.
* [8537](https://example.com/acme/logstore/pull/8537) **jeschkies**: Allow fetching all entries with `--limit=0`.

#### Fluent Bit

#### Logstore Canary

##### Enhancements

* [8024](https://example.com/acme/logstore/pull/8024) **jijotj**: Support passing logstore address as environment variable

#### Jsonnet

* [7923](https://example.com/acme/logstore/pull/7923) **manohar-koukuntla**: Add zone aware ingesters in jsonnet deployment

##### Fixes

* [8247](https://example.com/acme/logstore/pull/8247) **Whyeasy** fix usage of cluster label within Mixin.

#### Build

* [7938](https://example.com/acme/logstore/pull/7938) **ssncferreira**: Add DroneCI pipeline step to validate configuration flags documentation generation.

### Notes

### Dependencies

## 2.7.6 (2023-07-24)

#### Logstore

##### Fixes

* [10028](https://example.com/acme/logstore/pull/10028) **MichelHollands**: Use go 1.20.6
* [9185](https://example.com/acme/logstore/pull/9185) **dannykopping**: Prevent redis client from incorrectly choosing cluster mode with local address.
* [8824](https://example.com/acme/logstore/pull/8824) **periklis**: Expose optional label matcher for label values handler

## 2.7.5 (2023-03-28)

#### Logstore

##### Fixes

* [7924](https://example.com/acme/logstore/pull/7924) **jeschkies**: Flush buffered logger on exit

## 2.7.4 (2023-02-24)

#### Logstore

##### Fixes

* [8531](https://example.com/acme/logstore/pull/8531) **garrettlish**: logql: fix panics when cloning a special query
* [8120](https://example.com/acme/logstore/pull/8120) **ashwanthgoli**: fix panic on hitting /scheduler/ring when ring is disabled.
* [7988](https://example.com/acme/logstore/pull/7988) **ashwanthgoli**: store: write overlapping chunks to multiple stores.
* [7925](https://example.com/acme/logstore/pull/7925) **sandeepsukhani**: Fix bugs in logs results caching causing query-frontend to return logs outside of query window.

##### Build

* [8575](https://example.com/acme/logstore/pull/8575) **MichelHollands**: Update build image to go 1.20.1 and alpine 3.16.4.
* [8583](https://example.com/acme/logstore/pull/8583) **MichelHollands**: Use 0.28.1 build image and update go and alpine versions.

#### Promtail

##### Enhancements

##### Fixes

* [8497](https://example.com/acme/logstore/pull/8497) **kavirajk**: Fix `cri` tags treating different streams as the same
* [7771](https://example.com/acme/logstore/pull/7771) **GeorgeTsilias**: Handle nil error on target Details() call.
* [7461](https://example.com/acme/logstore/pull/7461) **MarNicGit**: Promtail: Fix collecting userdata field from Windows Event Log

## 2.7.3 (2023-02-01)

#### Logstore

##### Fixes

* [8340](https://example.com/acme/logstore/pull/8340) **MasslessParticle** Fix bug in compactor that caused panics when `startTime` and `endTime` of a delete request are equal.

#### Build

* [8232](https://example.com/acme/logstore/pull/8232) **TaehyunHwang** Fix build issue that caused `--version` to show wrong version for Logstore and Promtail binaries.

## 2.7.2 (2023-01-25)

#### Logstore

##### Fixes

* [7926](https://example.com/acme/logstore/pull/7926) **MichelHollands**: Fix bug in validation of `pattern` and `regexp` parsers where missing or empty parameters caused panics.
* [7720](https://example.com/acme/logstore/pull/7720) **sandeepsukhani**: Fix bugs in processing delete requests with line filters.
* [7708](https://example.com/acme/logstore/pull/7708) **DylanGuedes**: Fix bug in multi-tenant querying.

### Notes

This release was created from a branch starting at commit `706c22e9e40b0156031f214b63dc6ed4e210abc1` but it may also contain backported changes from main.

Check the history of the branch `release-2.7.x`.

### Dependencies

* Go version: 1.19.5

## 2.7.1 (2022-12-09)

#### Logstore

##### Enhancements

* [6360](https://example.com/acme/logstore/pull/6360) **liguozhong**: Hide error message when context timeout occurs in `s3.getObject`
* [7602](https://example.com/acme/logstore/pull/7602) **vmax**: Add decolorize filter to easily parse colored logs.
* [7804](https://example.com/acme/logstore/pull/7804) **sandeepsukhani**: Use grpc for communicating with compactor for query time filtering of data requested for deletion.
* [7684](https://example.com/acme/logstore/pull/7684) **kavirajk**: Add missing `embedded-cache` config under `cache_config` reference documentation.

##### Fixes

* [7453](https://example.com/acme/logstore/pull/7453) **periklis**: Add single compactor http client for delete and gennumber clients

##### Changes

* [7877](https://example.com/acme/logstore/pull/7877)A **trevorwhitney**: Due to a known bug with experimental new delete mode feature, the default delete mode has been changed to `filter-only`.

#### Promtail

##### Enhancements

* [7602](https://example.com/acme/logstore/pull/7602) **vmax**: Add decolorize stage to Promtail to easily parse colored logs.

##### Fixes

##### Changes

* [7587](https://example.com/acme/logstore/pull/7587) **mar4uk**: Add go build tag `promtail_journal_enabled` to include/exclude Promtail journald code from binary.

## 2.7.0

#### Logstore

##### Enhancements

* [7436](https://example.com/acme/logstore/pull/7436) **periklis**: Expose ring and memberlist handlers through internal server listener
* [7227](https://example.com/acme/logstore/pull/7227) **Red-GV**: Add ability to configure tls minimum version and cipher suites
* [7179](https://example.com/acme/logstore/pull/7179) **vlad-diachenko**: Add ability to use Azure Service Principals credentials to authenticate to Azure Blob Storage.
* [7063](https://example.com/acme/logstore/pull/7063) **kavirajk**: Add additional `push` mode to Logstore canary that can directly push logs to given Logstore URL.
* [7069](https://example.com/acme/logstore/pull/7069) **periklis**: Add support for custom internal server listener for readiness probes.
* [7023](https://example.com/acme/logstore/pull/7023) **liguozhong**: logql engine support exec `vector(0)` grammar.
* [6983](https://example.com/acme/logstore/pull/6983) **slim-bean**: `__timestamp__` and `__line__` are now available in the logql `label_format` query stage.
* [6821](https://example.com/acme/logstore/pull/6821) **kavirajk**: Introduce new cache type `embedded-cache` which is an in-process cache system that runs logstore without the need for an external cache (like memcached, redis, etc). It can be run in two modes `distributed: false` (default, and same as old `fifocache`) and `distributed: true` which runs cache in distributed fashion sharding keys across peers if Logstore is run in microservices or SSD mode.
* [6691](https://example.com/acme/logstore/pull/6691) **dannykopping**: Update production-ready Logstore cluster in docker-compose
* [6317](https://example.com/acme/logstore/pull/6317) **dannykoping**: General: add cache usage statistics
* [6444](https://example.com/acme/logstore/pull/6444) **aminesnow** Add TLS config to query frontend.
* [6179](https://example.com/acme/logstore/pull/6179) **chaudum**: Add new HTTP endpoint to delete ingester ring token file and shutdown process gracefully
* [5997](https://example.com/acme/logstore/pull/5997) **simonswine**: Querier: parallize label queries to both stores.
* [5406](https://example.com/acme/logstore/pull/5406) **ctovena**: Revise the configuration parameters that configure the usage report to acme.com.
* [7264](https://example.com/acme/logstore/pull/7264) **bboreham**: Chunks: decode varints directly from byte buffer, for speed.
* [7263](https://example.com/acme/logstore/pull/7263) **bboreham**: Dependencies: klauspost/compress package to v1.15.11; improves performance.
* [7270](https://example.com/acme/logstore/pull/7270) **wilfriedroset**: Add support for `username` to redis cache configuration.
* [6952](https://example.com/acme/logstore/pull/6952) **DylanGuedes**: Experimental: Introduce a new feature named stream sharding.

##### Fixes

* [7426](https://example.com/acme/logstore/pull/7426) **periklis**: Add missing compactor delete client tls client config
* [7238](https://example.com/acme/logstore/pull/7328) **periklis**: Fix internal server bootstrap for query frontend
* [7288](https://example.com/acme/logstore/pull/7288) **ssncferreira**: Fix query mapping in AST mapper `rangemapper` to support the new `VectorExpr` expression.
* [7040](https://example.com/acme/logstore/pull/7040) **bakunowski**: Remove duplicated `logstore_boltdb_shipper` prefix from `tables_upload_operation_total` metric.
* [6937](https://example.com/acme/logstore/pull/6937) **ssncferreira**: Fix topk and bottomk expressions with parameter <= 0.
* [6780](https://example.com/acme/logstore/pull/6780) **periklis**:  Attach the panic recovery handler on all HTTP handlers
* [6358](https://example.com/acme/logstore/pull/6358) **taharah**: Fixes sigv4 authentication for the Ruler's remote write configuration by allowing both a global and per tenant configuration.
* [6375](https://example.com/acme/logstore/pull/6375) **dannykopping**: Fix bug that prevented users from using the `json` parser after a `line_format` pipeline stage.
* [6505](https://example.com/acme/logstore/pull/6375) **dmitri-lerko** Fixes `failed to receive pubsub messages` error with promtail GCPLog client.
* [6372](https://example.com/acme/logstore/pull/6372) **splitice**: Add support for numbers in JSON fields.

##### Changes

* [6726](https://example.com/acme/logstore/pull/6726) **kavirajk**: upgrades go from 1.17.9 -> 1.18.4
* [6415](https://example.com/acme/logstore/pull/6415) **salvacorts**: Evenly spread queriers across kubernetes nodes.
* [6349](https://example.com/acme/logstore/pull/6349) **simonswine**: Update the default HTTP listen port from 80 to 3100. Make sure to configure the port explicitly if you are using port 80.
* [6835](https://example.com/acme/logstore/pull/6835) **DylanGuedes**: Add new per-tenant query timeout configuration and remove engine query timeout.
* [7212](https://example.com/acme/logstore/pull/7212) **Juneezee**: Replaces deprecated `io/ioutil` with `io` and `os`.
* [7292](https://example.com/acme/logstore/pull/7292) **jmherbst**: Add string conversion to value based drops to more intuitively match numeric fields. String conversion failure will result in no lines being dropped.
* [7361](https://example.com/acme/logstore/pull/7361) **szczepad**: Renames metric `logstore_log_messages_total` to `logstore_internal_log_messages_total`
* [7416](https://example.com/acme/logstore/pull/7416) **mstrzele**: Use the stable `HorizontalPodAutoscaler` v2, if possible, when installing using Helm
* [7510](https://example.com/acme/logstore/pull/7510) **slim-bean**: Limited queries (queries without filter expressions) will now be split and sharded.
* [5400](https://example.com/acme/logstore/pull/5400) **BenoitKnecht**: promtail/server: Disable profiling by default

#### Promtail

* [7470](https://example.com/acme/logstore/pull/7470) **Jack-King**: Add configuration for adding custom HTTP headers to push requests

##### Enhancements

* [7593](https://example.com/acme/logstore/pull/7593) **chodges15**: Promtail: Add tenant label to client drop metrics and logs
* [7101](https://example.com/acme/logstore/pull/7101) **liguozhong**: Promtail: Add support for max stream limit.
* [7247](https://example.com/acme/logstore/pull/7247) **liguozhong**: Add config reload endpoint / signal to promtail.
* [6708](https://example.com/acme/logstore/pull/6708) **DylanGuedes**: Add compressed files support to Promtail.
* [5977](https://example.com/acme/logstore/pull/5977) **juissi-t** lambda-promtail: Add support for Kinesis data stream events
* [6828](https://example.com/acme/logstore/pull/6828) **alexandre1984rj** Add the BotScore and BotScoreSrc fields once the Cloudflare API returns those two fields on the list of all available log fields.
* [6656](https://example.com/acme/logstore/pull/6656) **carlospeon**: Allow promtail to add matches to the journal reader
* [7401](https://example.com/acme/logstore/pull/7401) **thepalbi**: Add timeout to GCP Logs push target
* [7414](https://example.com/acme/logstore/pull/7414) **thepalbi**: Add basic tracing support

##### Fixes

* [7394](https://example.com/acme/logstore/pull/7394) **liguozhong**: Fix issue with the Cloudflare target that caused it to stop working after it received an error in the logpull request as explained in issue <https://example.com/acme/logstore/issues/6150>
* [6766](https://example.com/acme/logstore/pull/6766) **kavirajk**: fix(logql): Make `LabelSampleExtractor` ignore processing the line if it doesn't contain that specific label. Fixes unwrap behavior explained in the issue <https://example.com/acme/logstore/issues/6713>
* [7016](https://example.com/acme/logstore/pull/7016) **chodges15**: Fix issue with dropping logs when a file based SD target's labels are updated

##### Changes

* **quodlibetor**: Change Docker target discovery log level from `Error` to `Info`

#### Logcli

* [7325](https://example.com/acme/logstore/pull/7325) **dbirks**: Document setting up command completion
* [8518](https://example.com/acme/logstore/pull/8518) **SN9NV**: Add parallel flags

#### Fluent Bit

#### Logstore Canary

* [7398](https://example.com/acme/logstore/pull/7398) **verejoel**: Allow insecure TLS connections

#### Jsonnet

* [6189](https://example.com/acme/logstore/pull/6189) **irizzant**: Add creation of a `ServiceMonitor` object for Prometheus scraping through configuration parameter `create_service_monitor`. Simplify mixin usage by adding (<https://github.com/prometheus-operator/kube-prometheus>) library.
* [6662](https://example.com/acme/logstore/pull/6662) **Whyeasy**: Fixes memberlist error when using a stateful ruler.

### Notes

This release was created from a branch starting at commit `706c22e9e40b0156031f214b63dc6ed4e210abc1` but it may also contain backported changes from main.

Check the history of the branch `release-2.7.x`.

### Dependencies

* Go Version:     FIXME

# 2.6.1 (2022/07/18)

### All Changes

* [6658](https://example.com/acme/logstore/pull/6658) Updated the versions of [kit](https://example.com/acme/kit) and [memberlist](https://example.com/acme/memberlist) to allow configuring cluster labels for memberlist. Cluster labels prevent mixing the members between two consistent hash rings of separate applications that are run in the same Kubernetes cluster.
* [6681](https://example.com/acme/logstore/pull/6681) Fixed an HTTP connection leak between the querier and the compactor when the log entry deletion feature is enabled.
* [6583](https://example.com/acme/logstore/pull/6583) Fixed noisy error messages when the log entry deletion feature is disabled for a tenant.

# 2.6.0 (2022/07/08)

### All Changes

Here is the list with the changes that were produced since the previous release.

#### Logstore

##### Enhancements

* [5662](https://example.com/acme/logstore/pull/5662) **ssncferreira** **chaudum** Improve performance of instant queries by splitting range into multiple subqueries that are executed in parallel.
* [5848](https://example.com/acme/logstore/pull/5848) **arcosx**: Add Baidu AI Cloud as a storage backend choice.
* [6410](https://example.com/acme/logstore/pull/6410) **MichelHollands**: Add support for per tenant delete API access enabling.
* [5879](https://example.com/acme/logstore/pull/5879) **MichelHollands**: Remove lines matching delete request expression when using "filter-and-delete" deletion mode.
* [5984](https://example.com/acme/logstore/pull/5984) **dannykopping** and **salvacorts**: Improve query performance by preventing unnecessary querying of ingesters when the query data is old enough to be in object storage.
* [5971](https://example.com/acme/logstore/pull/5971) **kavirajk**: Extend the `metrics.go` recording of statistics about metadata queries to include labels and series queries.
* [6136](https://example.com/acme/logstore/pull/6136) **periklis**: Add support for alertmanager header authorization.
* [6163](https://example.com/acme/logstore/pull/6163) **jburnham**: LogQL: Add a `default` sprig template function in LogQL label/line formatter.

##### Fixes

* [6152](https://example.com/acme/logstore/pull/6152) **slim-bean**: Fixes unbounded ingester memory growth when live tailing under specific circumstances.
* [5685](https://example.com/acme/logstore/pull/5685) **chaudum**: Fix bug in push request parser that allowed users to send arbitrary non-string data as "log line".
* [5799](https://example.com/acme/logstore/pull/5799) **cyriltovena** Fix deduping issues when multiple entries with the same timestamp exist. !hide or not hide (bugfix Logstore)
* [5888](https://example.com/acme/logstore/pull/5888) **Papawy** Fix common configuration block net interface name when overwritten by ring common configuration.

##### Changes

* [6361](https://example.com/acme/logstore/pull/6361) **chaudum**: Sum values in unwrapped rate aggregation instead of treating them as counter.
* [6412](https://example.com/acme/logstore/pull/6412) **chaudum**: Add new unwrapped range aggregation `rate_counter()` to LogQL
* [6042](https://example.com/acme/logstore/pull/6042) **slim-bean**: Add a new configuration to allow fudging of ingested timestamps to guarantee sort order of duplicate timestamps at query time.
* [6120](https://example.com/acme/logstore/pull/6120) **KMiller-Acme**: Rename configuration parameter fudge_duplicate_timestamp to be increment_duplicate_timestamp.
* [5777](https://example.com/acme/logstore/pull/5777) **tatchiuleung**: storage: make Azure blobID chunk delimiter configurable
* [5650](https://example.com/acme/logstore/pull/5650) **cyriltovena**: Remove more chunkstore and schema version below v9
* [5643](https://example.com/acme/logstore/pull/5643) **simonswine**: Introduce a ChunkRef type as part of logproto
* [6435](https://example.com/acme/logstore/pull/6435) **MichelHollands**: Remove the `whole-stream-deletion` mode.
* [5899](https://example.com/acme/logstore/pull/5899) **simonswine**: Update go image to 1.17.9.

#### Promtail

##### Enhancements

* [6105](https://example.com/acme/logstore/pull/6105) **rutgerke** Export metrics for the Promtail journal target.
* [5943](https://example.com/acme/logstore/pull/5943) **tpaschalis**: Add configuration support for excluding configuration files when instantiating Promtail.
* [5790](https://example.com/acme/logstore/pull/5790) **chaudum**: Add UDP support for Promtail's syslog target.
* [6102](https://example.com/acme/logstore/pull/6102) **timchenko-a**: Add multi-tenancy support to lambda-promtail.
* [6099](https://example.com/acme/logstore/pull/6099) **cstyan**: Drop lines with malformed JSON in Promtail JSON pipeline stage.
* [5715](https://example.com/acme/logstore/pull/5715) **chaudum**: Allow promtail to push RFC5424 formatted syslog messages
* [6395](https://example.com/acme/logstore/pull/6395) **DylanGuedes**: Add encoding support

##### Fixes

* [6034](https://example.com/acme/logstore/pull/6034) **DylanGuedes**: Promtail: Fix symlink tailing behavior.

##### Changes

* [6371](https://example.com/acme/logstore/pull/6371) **witalisoft**: BREAKING: Support more complex match based on multiple extracted data fields in drop stage
* [5686](https://example.com/acme/logstore/pull/5686) **ssncferreira**: Move promtail StreamLagLabels config to upper level config.Config
* [5839](https://example.com/acme/logstore/pull/5839) **marctc**: Add ActiveTargets method to promtail
* [5661](https://example.com/acme/logstore/pull/5661) **masslessparticle**: Invalidate caches on deletes

#### Fluent Bit

* [5711](https://example.com/acme/logstore/pull/5711) **MichelHollands**: Update fluent-bit output name

#### Logstore Canary

* [6310](https://example.com/acme/logstore/pull/6310) **chodges15**: Add support for client-side TLS certs in logstore-canary for Logstore connection

### Notes

This release was created from a branch starting at commit `1794a766134f07b54386b1a431b58e1d44e6d7f7` but it may also contain backported changes from main.

Check the history of the branch `release-2.6.x`.

### Dependencies

* Go Version:     1.17.9

# 2.5.0 (2022/04/07)

Release notes for 2.5.0 can be found on the [release notes page](https://acme.com/docs/logstore/latest/release-notes/v2-5/)

### All Changes

Here is a list of all significant changes, in the past we have included all changes
but with over 500 PR's merged since the last release we decided to curate the list
to include only the most relevant.

#### Logstore

##### Enhancements

* [5542](https://example.com/acme/logstore/pull/5542) **bboreham**: regexp filter: use modified package with optimisations
* [5318](https://example.com/acme/logstore/pull/5318) **jeschkies**: Speed up `EntrySortIterator` by 20%.
* [5317](https://example.com/acme/logstore/pull/5317) **owen-d**: Logql/parallel binop
* [5315](https://example.com/acme/logstore/pull/5315) **bboreham**: filters: use faster regexp package
* [5311](https://example.com/acme/logstore/pull/5311) **vlad-diachenko**: Removed redundant memory allocations in parsers
* [5291](https://example.com/acme/logstore/pull/5291) **owen-d**: less opaque chunk keys on fs with v12
* [5275](https://example.com/acme/logstore/pull/5275) **SasSwart**: Parse duration expressions in accordance with promql
* [5249](https://example.com/acme/logstore/pull/5249) **3JIou-home**: Push: add deflate compression in post requests
* [5160](https://example.com/acme/logstore/pull/5160) **sandeepsukhani**: add objects list caching for boltdb-shipper index store to reduce object storage list api calls
* [5148](https://example.com/acme/logstore/pull/5148) **chaudum**: Auto-expire old items from FIFO cache
* [5093](https://example.com/acme/logstore/pull/5093) **liguozhong**: [enhancement] querier : Add "query_memory_only" to make logstore have option to rely only on memory availability.
* [5078](https://example.com/acme/logstore/pull/5078) **ssncferreira**: Logstore: Implement custom /config handler (#4785)
* [5054](https://example.com/acme/logstore/pull/5054) **JordanRushing**: new v12 schema optimized to better handle S3 prefix rate limits
* [5013](https://example.com/acme/logstore/pull/5013) **liguozhong**: [new feature] logql: extrapolate unwrapped rate function
* [4947](https://example.com/acme/logstore/pull/4947) **siavashs**: Support Redis Cluster Configuration Endpoint
* [4938](https://example.com/acme/logstore/pull/4938) **DylanGuedes**: Add distributor ring page
* [4879](https://example.com/acme/logstore/pull/4879) **cyriltovena**: LogQL: add **line** function to | line_format template
* [4858](https://example.com/acme/logstore/pull/4858) **sandy2008**: feat(): add ManagedIdentity in Azure Blob Storage

## Main

* [5789](https://example.com/acme/logstore/pull/5789) **bboreham**: Production config: add dot to some DNS address to reduce lookups.
* [5780](https://example.com/acme/logstore/pull/5780) **simonswine**: Update alpine image to 3.15.4.
* [5715](https://example.com/acme/logstore/pull/5715) **chaudum** Add option to push RFC5424 syslog messages from Promtail in syslog scrape target.
* [5696](https://example.com/acme/logstore/pull/5696) **paullryan** don't block scraping of new logs from cloudflare within promtail if an error is received from cloudflare about too early logs.
* [5685](https://example.com/acme/logstore/pull/5625) **chaudum** Fix bug in push request parser that allowed users to send arbitrary non-string data as "log line".
* [5707](https://example.com/acme/logstore/pull/5707) **franzwong** Promtail: Rename config name limit_config to limits_config.
* [5626](https://example.com/acme/logstore/pull/5626) **jeschkies** Apply query limits to multi-tenant queries by choosing the most restrictive limit from the set of tenant limits.
* [5622](https://example.com/acme/logstore/pull/5622) **chaudum**: Fix bug in query splitter that caused `interval` query parameter to be ignored and therefore returning more logs than expected.
* [5521](https://example.com/acme/logstore/pull/5521) **cstyan**: Move stream lag configuration to top level clients config struct and refactor stream lag metric, this resolves a bug with duplicate metric collection when a single Promtail binary is running multiple Promtail clients.
* [5568](https://example.com/acme/logstore/pull/5568) **afayngelerindbx**: Fix canary panics due to concurrent execution of `confirmMissing`
* [5552](https://example.com/acme/logstore/pull/5552) **jiachengxu**: Logstore mixin: add `DiskSpaceUtilizationPanel`
* [5541](https://example.com/acme/logstore/pull/5541) **bboreham**: Queries: reject very deeply nested regexps which could crash Logstore.
* [5536](https://example.com/acme/logstore/pull/5536) **jiachengxu**: Logstore mixin: make labelsSelector in logstore chunks dashboards configurable
* [5535](https://example.com/acme/logstore/pull/5535) **jiachengxu**: Logstore mixins: use labels selector for logstore chunks dashboard
* [5507](https://example.com/acme/logstore/pull/5507) **MichelHollands**: Remove extra param in call for inflightRequests metric.
* [5481](https://example.com/acme/logstore/pull/5481) **MichelHollands**: Add a DeletionMode config variable to specify the delete mode and validate match parameters.
* [5356](https://example.com/acme/logstore/pull/5356) **jbschami**: Enhance lambda-promtail to support adding extra labels from an environment variable value
* [5409](https://example.com/acme/logstore/pull/5409) **ldb**: Enable best effort parsing for Syslog messages
* [5392](https://example.com/acme/logstore/pull/5392) **MichelHollands**: Etcd credentials are parsed as secrets instead of plain text now.
* [5361](https://example.com/acme/logstore/pull/5361) **ctovena**: Add usage report to acme.com.
* [5354](https://example.com/acme/logstore/pull/5354) **tlinhart**: Add support for ARM64 to lambda-promtail drone build job.
* [5289](https://example.com/acme/logstore/pull/5289) **ctovena**: Fix deduplication bug in queries when mutating labels.
* [5302](https://example.com/acme/logstore/pull/5302) **MasslessParticle** Update azure blobstore client to use new sdk.
* [5243](https://example.com/acme/logstore/pull/5290) **ssncferreira**: Update Promtail to support duration string formats.
* [5266](https://example.com/acme/logstore/pull/5266) **jeschkies**: Write Promtail position file atomically on Unix.
* [5280](https://example.com/acme/logstore/pull/5280) **jeschkies**: Fix Docker target connection loss.
* [5243](https://example.com/acme/logstore/pull/5243) **owen-d**: moves `querier.split-queries-by-interval` to limits code only.
* [5139](https://example.com/acme/logstore/pull/5139) **DylanGuedes**: Drop support for legacy configuration rules format.
* [5262](https://example.com/acme/logstore/pull/5262) **MichelHollands**: Remove the labelFilter field
* [4911](https://example.com/acme/logstore/pull/4911) **jeschkies**: Support Docker service discovery in Promtail.
* [5107](https://example.com/acme/logstore/pull/5107) **chaudum** Fix bug in fluentd plugin that caused log lines containing non UTF-8 characters to be dropped.
* [5148](https://example.com/acme/logstore/pull/5148) **chaudum** Add periodic task to prune old expired items from the FIFO cache to free up memory.
* [5187](https://example.com/acme/logstore/pull/5187) **aknuds1** Rename metric `corestore_experimental_features_in_use_total` to `logstore_experimental_features_in_use_total` and metric `log_messages_total` to `logstore_log_messages_total`.
* [5170](https://example.com/acme/logstore/pull/5170) **chaudum** Fix deadlock in Promtail caused when targets got removed from a target group by the discovery manager.
* [5163](https://example.com/acme/logstore/pull/5163) **chaudum** Fix regression in fluentd plugin introduced with #5107 that caused `NoMethodError` when parsing non-string values of log lines.
* [5144](https://example.com/acme/logstore/pull/5144) **dannykopping** Ruler: fix remote write basic auth credentials.
* [5091](https://example.com/acme/logstore/pull/5091) **owen-d**: Changes `ingester.concurrent-flushes` default to 32
* [5031](https://example.com/acme/logstore/pull/5031) **liguozhong**: Promtail: Add global read rate limiting.
* [4879](https://example.com/acme/logstore/pull/4879) **cyriltovena**: LogQL: add **line** function to | line_format template.
* [5081](https://example.com/acme/logstore/pull/5081) **SasSwart**: Add the option to configure memory ballast for Logstore
* [5085](https://example.com/acme/logstore/pull/5085) **aknuds1**: Upgrade Corestore to [e0807c4eb487](https://github.com/corestoreproject/corestore/compare/4e9fc3a2b5ab..e0807c4eb487) and Prometheus to [692a54649ed7](https://github.com/prometheus/prometheus/compare/2a3d62ac8456..692a54649ed7)
* [5067](https://example.com/acme/logstore/pull/5057) **cstyan**: Add a metric to Azure Blob Storage client to track total egress bytes
* [5065](https://example.com/acme/logstore/pull/5065) **AndreZiviani**: lambda-promtail: Add ability to ingest logs from S3
* [4950](https://example.com/acme/logstore/pull/4950) **DylanGuedes**: Implement common instance addr/net interface
* [4949](https://example.com/acme/logstore/pull/4949) **ssncferreira**: Add query `queueTime` metric to statistics and metrics.go
* [4938](https://example.com/acme/logstore/pull/4938) **DylanGuedes**: Implement ring status page for the distributor
* [5023](https://example.com/acme/logstore/pull/5023) **ssncferreira**: Move `querier.split-queries-by-interval` to a per-tenant configuration
* [4993](https://example.com/acme/logstore/pull/4926) **thejosephstevens**: Fix parent of wal and wal_cleaner in logstore ruler config docs
* [4933](https://example.com/acme/logstore/pull/4933) **jeschkies**: Support matchers in series label values query.
* [4926](https://example.com/acme/logstore/pull/4926) **thejosephstevens**: Fix comment in Logstore module loading for accuracy
* [4920](https://example.com/acme/logstore/pull/4920) **chaudum**: Add `-list-targets` command line flag to list all available run targets
* [4860](https://example.com/acme/logstore/pull/4860) **cyriltovena**: Add rate limiting and metrics to hedging
* [4865](https://example.com/acme/logstore/pull/4865) **taisho6339**: Fix duplicate registry.MustRegister call in Promtail Kafka
* [4845](https://example.com/acme/logstore/pull/4845) **chaudum** Return error responses consistently as JSON
* [4826](https://example.com/acme/logstore/pull/4826) **cyriltovena**: Adds the ability to hedge storage requests.
* [4785](https://example.com/acme/logstore/pull/4785) **DylanGuedes**: Logstore: Print current config by calling /config
* [4775](https://example.com/acme/logstore/pull/4775) **jeschkies**: Make `*` and `+` non-greedy to double regex filter speed.
* [4769](https://example.com/acme/logstore/pull/4769) **cyriltovena**: Improve LogQL format stages requireLabel
* [4731](https://example.com/acme/logstore/pull/4731) **cyriltovena**: Improve heap iterators.
* [4394](https://example.com/acme/logstore/pull/4394) **cyriltovena**: Improve case insensitive search to avoid allocations.

##### Fixes

* [5768](https://example.com/acme/logstore/pull/5768) **slim-bean**: Logstore: Increase flush_op_timeout default from 10s to 10m
* [5761](https://example.com/acme/logstore/pull/5761) **slim-bean**: Promtil: Fix a panic when using the logstore push api target.
* [5622](https://example.com/acme/logstore/pull/5622) **chaudum**: Preserve interval parameter when splitting queries by time
* [5541](https://example.com/acme/logstore/pull/5541) **bboreham**: Queries: update package to reject very deeply nested regexps which could crash Logstore
* [5527](https://example.com/acme/logstore/pull/5527) **liguozhong**: [bugfix] fix nil pointer
* [5474](https://example.com/acme/logstore/pull/5474) **cyriltovena**: Disable sharding of count/avg when labels are mutated
* [5472](https://example.com/acme/logstore/pull/5472) **MasslessParticle**: Fix potential deadlock in the table manager
* [5444](https://example.com/acme/logstore/pull/5444) **cyriltovena**: Do not insert missing point when sharding
* [5425](https://example.com/acme/logstore/pull/5425) **cyriltovena**: Do not use WaitGroup context for StepEvaluator
* [5423](https://example.com/acme/logstore/pull/5423) **cyriltovena**: Correctly sets hash value for headblock iterator
* [5418](https://example.com/acme/logstore/pull/5418) **RangerCD**: Fix two remote_timeout configs in ingester_client block
* [5413](https://example.com/acme/logstore/pull/5413) **MasslessParticle**: Fix a deadlock in the Azure Blob client
* [5399](https://example.com/acme/logstore/pull/5399) **MasslessParticle**: Fix Azure issue where 404 not recognized
* [5362](https://example.com/acme/logstore/pull/5362) **gotjosh**: Ruler: Rule group not found API message
* [5342](https://example.com/acme/logstore/pull/5342) **sandeepsukhani**: Fix apply retention issue
* [5334](https://example.com/acme/logstore/pull/5334) **kavirajk**: Makes `tailer.droppedStreams` slice bounded.
* [5324](https://example.com/acme/logstore/pull/5324) **owen-d**: Release entryBufferPool once
* [5303](https://example.com/acme/logstore/pull/5303) **owen-d**: Better logic for when to shard wrt disabled lookback
* [5298](https://example.com/acme/logstore/pull/5298) **sandeepsukhani**: fix a panic in index-gateway caused by double closing of a channel
* [5297](https://example.com/acme/logstore/pull/5297) **vlad-diachenko**: Changed logic of handling RPC error with code Cancelled
* [5289](https://example.com/acme/logstore/pull/5289) **cyriltovena**: Fixes log deduplication when mutating Labels using LogQL
* [5261](https://example.com/acme/logstore/pull/5261) **sandeepsukhani**: use default retention period to check user index may have expired chunks when user does not have custom retention
* [5234](https://example.com/acme/logstore/pull/5234) **RangerCD**: Ignore missing stream while querying from ingester
* [5168](https://example.com/acme/logstore/pull/5168) **kavirajk**: Add `nil` check for Ruler BasicAuth config.
* [5144](https://example.com/acme/logstore/pull/5144) **dannykopping**: Ruler: Fix remote write basic auth credentials
* [5113](https://example.com/acme/logstore/pull/5113) **kavirajk**: Fix cancel issue between Query Frontend and Query Schdeduler
* [5080](https://example.com/acme/logstore/pull/5080) **kavirajk**: Handle `context` cancellation in some of the `querier` downstream requests
* [5075](https://example.com/acme/logstore/pull/5075) **cyriltovena**: Fixes a possible cancellation issue in the frontend
* [5063](https://example.com/acme/logstore/pull/5063) **cyriltovena**: Fix deadlock in disconnecting querier
* [5060](https://example.com/acme/logstore/pull/5060) **cyriltovena**: Fix race conditions in frontend_scheduler_worker.
* [5006](https://example.com/acme/logstore/pull/5006) **sandeepsukhani**: fix splitting of queries when step is larger than split interval
* [4904](https://example.com/acme/logstore/pull/4904) **bboreham**: ingester: use consistent set of instances to avoid panic
* [4902](https://example.com/acme/logstore/pull/4902) **cyriltovena**: Fixes 500 when query is outside of max_query_lookback
* [4828](https://example.com/acme/logstore/pull/4828) **chaudum**: Set correct `Content-Type` header in query response
* [4761](https://example.com/acme/logstore/pull/4761) **slim-bean**: Logstore: Set querier worker max concurrent regardless of run configuration.
* [4741](https://example.com/acme/logstore/pull/4741) **sandeepsukhani**: index cleanup fixes while applying retention

##### Changes

* [5544](https://example.com/acme/logstore/pull/5544) **ssncferreira**: Update vectorAggEvaluator to fail for expressions without grouping
* [5543](https://example.com/acme/logstore/pull/5543) **cyriltovena**: update logstore go version to 1.17.8
* [5450](https://example.com/acme/logstore/pull/5450) **BenoitKnecht**: pkg/ruler/base: Add external_labels option
* [5522](https://example.com/acme/logstore/pull/5522) **liguozhong**: chunk backend: Integrate Alibaba Cloud oss
* [5484](https://example.com/acme/logstore/pull/5484) **sandeepsukhani**: Add support for per user index query readiness with limits overrides
* [5719](https://example.com/acme/logstore/pull/5719) **kovaxur**: Logstore can use both basic-auth and tenant-id
* [5358](https://example.com/acme/logstore/pull/5358) **DylanGuedes**: Add `RingMode` support to `IndexGateway`
* [5435](https://example.com/acme/logstore/pull/5435) **slim-bean**: set match_max_concurrent true by default
* [5361](https://example.com/acme/logstore/pull/5361) **cyriltovena**: Add usage report into Logstore.
* [5243](https://example.com/acme/logstore/pull/5243) **owen-d**: Refactor/remove global splitby
* [5229](https://example.com/acme/logstore/pull/5229) **chaudum**: Return early if push payload does not contain data
* [5217](https://example.com/acme/logstore/pull/5217) **sandeepsukhani**: step align start and end time of the original query while splitting it
* [5204](https://example.com/acme/logstore/pull/5204) **trevorwhitney**: Default max_outstanding_per_tenant to 2048
* [5181](https://example.com/acme/logstore/pull/5181) **sandeepsukhani**: align metric queries by step and other queries by split interval
* [5178](https://example.com/acme/logstore/pull/5178) **liguozhong**: Handle `context` cancellation in some of the `querier` store.index-cache-read.
* [5172](https://example.com/acme/logstore/pull/5172) **cyriltovena**: Avoid splitting large range vector aggregation.
* [5125](https://example.com/acme/logstore/pull/5125) **sasagarw**: Remove split-queries-by-interval validation
* [5091](https://example.com/acme/logstore/pull/5091) **owen-d**: better defaults for flush queue parallelism
* [5083](https://example.com/acme/logstore/pull/5083) **liguozhong**: [enhancement] querier cache: WriteBackCache should be off query path
* [5081](https://example.com/acme/logstore/pull/5081) **SasSwart**: Add the option to configure memory ballast for Logstore
* [5077](https://example.com/acme/logstore/pull/5077) **trevorwhitney**: improve default config values
* [5067](https://example.com/acme/logstore/pull/5067) **cstyan**: Add an egress bytes total metric to the azure client.
* [5026](https://example.com/acme/logstore/pull/5026) **sandeepsukhani**: compactor changes for building per user index files in boltdb shipper
* [5023](https://example.com/acme/logstore/pull/5023) **ssncferreira**: Move querier.split-queries-by-interval to a per-tenant configuration
* [5022](https://example.com/acme/logstore/pull/5022) **owen-d**: adds instrumentation to azure object client
* [4942](https://example.com/acme/logstore/pull/4942) **cyriltovena**: Allow to disable http2 for GCS.
* [4891](https://example.com/acme/logstore/pull/4891) **liguozhong**: [optimization] cache prometheus : fix "logstore_cache_request_duration_seconds_bucket" ‘status_code’ label always equals "200"
* [4737](https://example.com/acme/logstore/pull/4737) **owen-d**: ensures components with required SRV lookups use the correct port
* [4736](https://example.com/acme/logstore/pull/4736) **sandeepsukhani**: allow applying retention at different interval than compaction with a config
* [4656](https://example.com/acme/logstore/pull/4656) **ssncferreira**: Fix kit/ring metric with 'corestore_' prefix

#### Promtail

##### Enhancements

* [5359](https://example.com/acme/logstore/pull/5359) **JBSchami**: Lambda-promtail: Enhance lambda-promtail to support adding extra labels from an environment variable value
* [5290](https://example.com/acme/logstore/pull/5290) **ssncferreira**: Update promtail to support duration string formats
* [5051](https://example.com/acme/logstore/pull/5051) **liguozhong**: [new] promtail pipeline:  Promtail Rate Limit stage #5048
* [5031](https://example.com/acme/logstore/pull/5031) **liguozhong**: [new] promtail: add readline rate limit
* [4911](https://example.com/acme/logstore/pull/4911) **jeschkies**: Provide Docker target and discovery in Promtail.
* [4813](https://example.com/acme/logstore/pull/4813) **cyriltovena**: Promtail pull cloudflare logs
* [4744](https://example.com/acme/logstore/pull/4744) **cyriltovena**: Add GELF support for Promtail.
* [4663](https://example.com/acme/logstore/pull/4663) **taisho6339**: Add SASL&mTLS authentication support for Kafka in Promtail

##### Fixes

* [5497](https://example.com/acme/logstore/pull/5497) **MasslessParticle**: Fix orphaned metrics in the file tailer
* [5409](https://example.com/acme/logstore/pull/5409) **ldb**: promtail/targets/syslog: Enable best effort parsing for Syslog messages
* [5246](https://example.com/acme/logstore/pull/5246) **rsteneteg**: Promtail: skip glob search if filetarget path is an existing file and not a directory
* [5238](https://example.com/acme/logstore/pull/5238) **littlepangdi**: Promtail: fix TargetManager.run() not exit after stop is called
* [4874](https://example.com/acme/logstore/pull/4874) **Alan01252**: Promtail: Fix replace missing adjacent capture groups
* [4832](https://example.com/acme/logstore/pull/4832) **taisho6339**: Use http prefix path correctly in promtail
* [4716](https://example.com/acme/logstore/pull/4716) **cyriltovena**: Fixes Promtail User-Agent.
* [5698](https://example.com/acme/logstore/pull/5698) **paullryan**: Promtail: Fix retry/stop when erroring for out of cloudflare retention range (e.g. over 168 hours old)

##### Changes

* [5377](https://example.com/acme/logstore/pull/5377) **slim-bean**: Promtail: Remove promtail_log_entries_bytes_bucket histogram
* [5266](https://example.com/acme/logstore/pull/5266) **jeschkies**: Write Promtail position file atomically.
* [4794](https://example.com/acme/logstore/pull/4794) **taisho6339**: Aggregate inotify watcher to file target manager
* [4745](https://example.com/acme/logstore/pull/4745) **taisho6339**: Expose Kafka message key in labels

#### Logcli

* [5477](https://example.com/acme/logstore/pull/5477) **atomic77**: logcli: Remove port from TLS server name when provided in --addr
* [4667](https://example.com/acme/logstore/pull/4667) **jeschkies**: Package logcli as rpm and deb.
* [4606](https://example.com/acme/logstore/pull/4606) **kavirajk**: Execute Logstore queries on raw log data piped to stdin

#### Lambda-Promtail

* [5065](https://example.com/acme/logstore/pull/5065) **AndreZiviani**: lambda-promtail: Add ability to ingest logs from S3
* [7632](https://example.com/acme/logstore/pull/7632) **changhyuni**: lambda-promtail: Add kinesis data stream to use in terraform

#### Fluent Bit

* [5223](https://example.com/acme/logstore/pull/5223) **cyriltovena**: fluent-bit: Attempt to unmarshal nested json.

#### FluentD

* [6240](https://example.com/acme/logstore/pull/6240) **taharah**: Add the feature flag `include_thread_label` to allow the `fluentd_thread` label included when using multiple threads for flushing to be configurable
* [5107](https://example.com/acme/logstore/pull/5107) **chaudum**: fluentd: Fix bug that caused lines to be dropped when containing non utf-8 characters
* [5163](https://example.com/acme/logstore/pull/5163) **chaudum**: Fix encoding error in fluentd client

### Notes

This release was created from a branch starting at commit 614912181e6f3988b2b22791053278cfb64e169c but it may also contain backported changes from main.

Check the history of the branch `release-2.5.x`.

### Dependencies

* Go Version:     1.17.8

# 2.4.1 (2021/11/07)

Release notes for 2.4.1 can be found on the [release notes page](https://acme.com/docs/logstore/latest/release-notes/v2-4/)

### All Changes

* [4687](https://example.com/acme/logstore/pull/4687) **owen-d**: overrides checks for nil tenant limits on AllByUserID
* [4683](https://example.com/acme/logstore/pull/4683) **owen-d**: Adds replication_factor doc to common config
* [4681](https://example.com/acme/logstore/pull/4681) **slim-bean**: Logstore: check new Read target when initializing boltdb-shipper store

# 2.4.0 (2021/11/05)

Release notes for 2.4.0 can be found on the [release notes page](https://acme.com/docs/logstore/latest/release-notes/v2-4/)

### All Changes

Here is a list of all changes included in 2.4.0.

#### Logstore

* [4649](https://example.com/acme/logstore/pull/4649) **cstyan**: Instrument s3 client DeleteObject requests.
* [4643](https://example.com/acme/logstore/pull/4643) **trevorwhitney**: compactor depends on memberlist for memberlist ring option
* [4642](https://example.com/acme/logstore/pull/4642) **slim-bean**: Logstore: fix handling of tail requests when using target `all` or `read`
* [4641](https://example.com/acme/logstore/pull/4641) **ssncferreira**: Migration to kit/ring
* [4638](https://example.com/acme/logstore/pull/4638) **DylanGuedes**: Logstore: Revert distributor defaulting to `inmemory`
* [4635](https://example.com/acme/logstore/pull/4635) **owen-d**: dont try to use the scheduler ring when a downstream url is configured
* [4630](https://example.com/acme/logstore/pull/4630) **chaudum**: Allow HTTP POST requests on ring pages
* [4627](https://example.com/acme/logstore/pull/4627) **slim-bean**: Logstore: Explicitly define allowed HTTP methods on HTTP endpoints
* [4625](https://example.com/acme/logstore/pull/4625) **sandeepsukhani**: Logs deletion fixes
* [4617](https://example.com/acme/logstore/pull/4617) **trevorwhitney**: Add common ring configuration
* [4615](https://example.com/acme/logstore/pull/4615) **owen-d**: uses ring.Write instead of ring.WriteNoExtend for compactor ring checks
* [4614](https://example.com/acme/logstore/pull/4614) **slim-bean**: Logstore: query scheduler should send shutdown to frontends when ReplicationSet changes
* [4608](https://example.com/acme/logstore/pull/4608) **trevorwhitney**: default ingester final sleep to 0 unless otherwise specified
* [4607](https://example.com/acme/logstore/pull/4607) **owen-d**: improves scheduler & compactor ringwatcher checks
* [4603](https://example.com/acme/logstore/pull/4603) **garrettlish**: add date time sprig template functions in logql label/line formatter
* [4598](https://example.com/acme/logstore/pull/4598) **kavirajk**: Fix `ip` matcher lexer to differentiate filter from identifier
* [4596](https://example.com/acme/logstore/pull/4596) **owen-d**: Ignore validity window during wal replay
* [4595](https://example.com/acme/logstore/pull/4595) **owen-d**: Cleans up redundant setting of stream.unorderedWrites=true during replay
* [4594](https://example.com/acme/logstore/pull/4594) **owen-d**: Enable unordered_writes by default
* [4593](https://example.com/acme/logstore/pull/4593) **taisho6339**: Respect gRPC context error when handling errors
* [4592](https://example.com/acme/logstore/pull/4592) **owen-d**: introduces "entry too far behind" instrumentation for unordered writes
* [4589](https://example.com/acme/logstore/pull/4589) **owen-d**: replaces fallthrough statement in InitFrontend
* [4586](https://example.com/acme/logstore/pull/4586) **dannykopping**: Configuring query-frontend interface names with loopback device
* [4585](https://example.com/acme/logstore/pull/4585) **sandeepsukhani**: set wal dir to /logstore/wal in docker config
* [4577](https://example.com/acme/logstore/pull/4577) **taisho6339**: Respect shard number in series api
* [4574](https://example.com/acme/logstore/pull/4574) **slim-bean**: Logstore: Add a ring to the compactor used to control concurrency when not running standalone
* [4573](https://example.com/acme/logstore/pull/4573) **sandeepsukhani**: validate default limits config with other configs at startup
* [4570](https://example.com/acme/logstore/pull/4570) **DylanGuedes**: Logstore: Append loopback to ingester net interface default list
* [4569](https://example.com/acme/logstore/pull/4569) **DylanGuedes**: Config: Change default RejectOldSamplesMaxAge from 14d to 7d
* [4563](https://example.com/acme/logstore/pull/4563) **cyriltovena**: Fixes the Series function to handle properly sharding.
* [4554](https://example.com/acme/logstore/pull/4554) **cyriltovena**: Fixes a panic in the labels API when no parameters are supplied.
* [4550](https://example.com/acme/logstore/pull/4550) **cyriltovena**: Fixes an edge case in the batch chunk iterator.
* [4546](https://example.com/acme/logstore/pull/4546) **slim-bean**: Logstore: Apply the ingester ring config to all other rings (distributor, ruler, query-scheduler)
* [4545](https://example.com/acme/logstore/pull/4545) **trevorwhitney**: Fix race condition in Query Scheduler ring with frontend/worker
* [4543](https://example.com/acme/logstore/pull/4543) **trevorwhitney**: Change a few default config values and improve application of common storage config
* [4542](https://example.com/acme/logstore/pull/4542) **owen-d**: only exports tenant limits which differ from defaults and export defa…
* [4531](https://example.com/acme/logstore/pull/4531) **JordanRushing**: Add quick nil check in TenantLimits for runtime_config
* [4529](https://example.com/acme/logstore/pull/4529) **owen-d**: correctly sets subservicesWatcher on scheduler
* [4525](https://example.com/acme/logstore/pull/4525) **owen-d**: Safely checks read ring for potentially nil scheduler
* [4524](https://example.com/acme/logstore/pull/4524) **dannykopping**: Clarify error message when no valid target scrape config is defined for `promtail` job
* [4520](https://example.com/acme/logstore/pull/4520) **JordanRushing**: Introduce `overrides-exporter` module to Logstore
* [4519](https://example.com/acme/logstore/pull/4519) **DylanGuedes**: Logstore: Enable FIFO cache by default
* [4518](https://example.com/acme/logstore/pull/4518) **slim-bean**: Logstore: Fix bug where items are returned to a sync.Pool incorrectly
* [4510](https://example.com/acme/logstore/pull/4510) **lingpeng0314**: add group_{left,right} to LogQL
* [4508](https://example.com/acme/logstore/pull/4508) **trevorwhitney**: Apply better defaults when boltdb shipper is being used
* [4498](https://example.com/acme/logstore/pull/4498) **trevorwhitney**: Feature: add virtual read and write targets
* [4487](https://example.com/acme/logstore/pull/4487) **cstyan**: Update go.mod to go 1.17
* [4484](https://example.com/acme/logstore/pull/4484) **dannykopping**: Replacing go-kit/kit/log with go-kit/log
* [4482](https://example.com/acme/logstore/pull/4482) **owen-d**: always expose logstore_build_info
* [4479](https://example.com/acme/logstore/pull/4479) **owen-d**: restores for state at seconds(now-forDuration)
* [4478](https://example.com/acme/logstore/pull/4478) **replay**: Update corestore to newer version
* [4473](https://example.com/acme/logstore/pull/4473) **trevorwhitney**: Configuration: add a common config section for object storage
* [4457](https://example.com/acme/logstore/pull/4457) **kavirajk**: Fix return values of Matrix and Vector during query range in QueryShardingMiddleware
* [4453](https://example.com/acme/logstore/pull/4453) **liguozhong**: [querier] s3: add getObject retry
* [4446](https://example.com/acme/logstore/pull/4446) **garrettlish**: make LogQL syntax scope from private to public
* [4443](https://example.com/acme/logstore/pull/4443) **DylanGuedes**: Logstore: Change how push API checks for contentType
* [4440](https://example.com/acme/logstore/pull/4440) **DylanGuedes**: Logstore: Override distributor's default ring KV store
* [4437](https://example.com/acme/logstore/pull/4437) **dannykopping**: Ruler: Do not clear remote-write HTTP client config
* [4436](https://example.com/acme/logstore/pull/4436) **JordanRushing**: Add metric prefix changes for chunk store and runtime config to upgrading.md
* [4435](https://example.com/acme/logstore/pull/4435) **trevorwhitney**: Change default values for two GRPC setting we have to set so the queriers can connect to a frontend or scheduler
* [4433](https://example.com/acme/logstore/pull/4433) **trevorwhitney**: Add more tests around config parsing changes from common config PR
* [4432](https://example.com/acme/logstore/pull/4432) **owen-d**: tests checkpoints immediately and gives more of a time buffer
* [4431](https://example.com/acme/logstore/pull/4431) **dannykopping**: Ruler: Overwrite instead of merge remote-write headers
* [4429](https://example.com/acme/logstore/pull/4429) **dannykopping**: Ruler: Refactoring remote-write config overrides
* [4424](https://example.com/acme/logstore/pull/4424) **slim-bean**: Logstore: Add a ring to the query scheduler to allow discovery via the ring as an alternative to DNS
* [4421](https://example.com/acme/logstore/pull/4421) **owen-d**: Safe per tenant overrides loading
* [4415](https://example.com/acme/logstore/pull/4415) **DylanGuedes**: Logstore: Change default limits to common values
* [4413](https://example.com/acme/logstore/pull/4413) **trevorwhitney**: add compactor working dir to auto-configured file paths
* [4411](https://example.com/acme/logstore/pull/4411) **slim-bean**: Logstore: Bug: frontend waiting on results which would never come
* [4400](https://example.com/acme/logstore/pull/4400) **trevorwhitney**: auto-apply memberlist ring config when join_members provided
* [4391](https://example.com/acme/logstore/pull/4391) **garrettlish**: add on and ignoring clauses in binOpExpr
* [4388](https://example.com/acme/logstore/pull/4388) **trevorwhitney**: default chunk target size to ~1MB~ 1.5MB
* [4367](https://example.com/acme/logstore/pull/4367) **owen-d**: removes deprecated duplicate per stream rate limit fields
* [4364](https://example.com/acme/logstore/pull/4364) **dannykopping**: Ruler: improve control over marshaling relabel.Config
* [4354](https://example.com/acme/logstore/pull/4354) **dannykopping**: Ruler: adding `pkg/metrics` from agent
* [4349](https://example.com/acme/logstore/pull/4349) **JordanRushing**: Add recovery middleware to Ingester; re-add recovery middleware to Querier when not running in standalone mode
* [4348](https://example.com/acme/logstore/pull/4348) **trevorwhitney**: allow ingester and distributor to run on same instance
* [4347](https://example.com/acme/logstore/pull/4347) **slim-bean**: Logstore: Common Config
* [4344](https://example.com/acme/logstore/pull/4344) **dannykopping**: Ruler: per-tenant WAL
* [4327](https://example.com/acme/logstore/pull/4327) **aknuds1**: Chore: Use kit/limiter
* [4322](https://example.com/acme/logstore/pull/4322) **owen-d**: Hotfix #4308 into k62
* [4321](https://example.com/acme/logstore/pull/4321) **owen-d**: Hotfix #4308 into k61
* [4313](https://example.com/acme/logstore/pull/4313) **aknuds1**: Chore: Use middleware package from kit
* [4312](https://example.com/acme/logstore/pull/4312) **aknuds1**: Chore: Use kit/grpcclient
* [4308](https://example.com/acme/logstore/pull/4308) **cyriltovena**: Fixes the pattern parser validation.
* [4304](https://example.com/acme/logstore/pull/4304) **aknuds1**: Chore: Reformat Go files
* [4302](https://example.com/acme/logstore/pull/4302) **cyriltovena**: Fixes a bug in the block cache code.
* [4301](https://example.com/acme/logstore/pull/4301) **trevorwhitney**: Feature: allow querier and query frontend targets to run on same process
* [4295](https://example.com/acme/logstore/pull/4295) **aknuds1**: Chore: Upgrade kit
* [4289](https://example.com/acme/logstore/pull/4289) **kavirajk**: Add custom UnmarshalJSON for bytesize type
* [4282](https://example.com/acme/logstore/pull/4282) **chaudum**: Chore: Update Corestore and use kv package from acme/kit
* [4276](https://example.com/acme/logstore/pull/4276) **chaudum**: Export MemberlistKV field on Logstore struct
* [4272](https://example.com/acme/logstore/pull/4272) **taisho6339**: Add count to 'logstore_ingester_memory_chunks' when recovery from wal
* [4265](https://example.com/acme/logstore/pull/4265) **owen-d**: remove empty streams after wal replay
* [4255](https://example.com/acme/logstore/pull/4255) **owen-d**: replaces old corestore_chunk_store prefix with logstore_chunk_store
* [4253](https://example.com/acme/logstore/pull/4253) **JordanRushing**: Change prefix for `runtimeconfig` metrics from `corestore_` to `logstore_`
* [4251](https://example.com/acme/logstore/pull/4251) **dannykopping**: Runtime config: do not validate nil limits
* [4246](https://example.com/acme/logstore/pull/4246) **JordanRushing**:     Add missing `Inc()` to correctly increment the `dropStage.dropCount` metric on valid dropped log line; update related docs
* [4240](https://example.com/acme/logstore/pull/4240) **bboreham**: Simplify Distributor.push
* [4238](https://example.com/acme/logstore/pull/4238) **liguozhong**: [fix] distributor: fix goroutine leak
* [4236](https://example.com/acme/logstore/pull/4236) **owen-d**: better per stream rate limits configuration options
* [4228](https://example.com/acme/logstore/pull/4228) **owen-d**: bumps per stream default rate limits
* [4227](https://example.com/acme/logstore/pull/4227) **aknuds1**: Chore: Use runtimeconfig from kit
* [4225](https://example.com/acme/logstore/pull/4225) **aknuds1**: Flagext: Use flagext package from kit
* [4213](https://example.com/acme/logstore/pull/4213) **owen-d**: Refactor per stream rate limit
* [4212](https://example.com/acme/logstore/pull/4212) **owen-d**: WAL replay discard metrics
* [4211](https://example.com/acme/logstore/pull/4211) **BenoitKnecht**: pkg/storage/chunk/aws: Add s3.http.ca-file option
* [4207](https://example.com/acme/logstore/pull/4207) **cstyan**: Improve error message for stream rate limit.
* [4196](https://example.com/acme/logstore/pull/4196) **56quarters**: Chore: Use services and modules from acme/kit
* [4193](https://example.com/acme/logstore/pull/4193) **owen-d**: adds logstore_ingester_wal_replay_active metric and records this more acc…
* [4192](https://example.com/acme/logstore/pull/4192) **owen-d**: Cleanup/unordered writes ingester config
* [4191](https://example.com/acme/logstore/pull/4191) **cstyan**: [ingester/stream]: Add a byte stream rate limit.
* [4188](https://example.com/acme/logstore/pull/4188) **aknuds1**: Chore: Upgrade to latest Corestore
* [4185](https://example.com/acme/logstore/pull/4185) **sandeepsukhani**: Canary: allow setting tenant id for querying logs from logstore
* [4181](https://example.com/acme/logstore/pull/4181) **owen-d**: initiate grpc health check always
* [4176](https://example.com/acme/logstore/pull/4176) **sokoide**: Authc/z: Enable grpc_client_config to allow mTLS
* [4172](https://example.com/acme/logstore/pull/4172) **sandeepsukhani**: Retention speedup
* [4160](https://example.com/acme/logstore/pull/4160) **owen-d**: safely close nonOverlapping iterators
* [4155](https://example.com/acme/logstore/pull/4155) **owen-d**: Auth followup - Remove unused
* [4153](https://example.com/acme/logstore/pull/4153) **owen-d**: uses more fleshed out corestore auth utility & adds new auth-ignored routes
* [4149](https://example.com/acme/logstore/pull/4149) **owen-d**: add unordered writes to local config
* [4141](https://example.com/acme/logstore/pull/4141) **dannykopping**: Ruler: write meaningful logs when remote-write is disabled or is misconfigured
* [4135](https://example.com/acme/logstore/pull/4135) **slim-bean**: Build: Fix build version info
* [4132](https://example.com/acme/logstore/pull/4132) **owen-d**: Promote/ruler api
* [4130](https://example.com/acme/logstore/pull/4130) **owen-d**: Tenant/unordered
* [4128](https://example.com/acme/logstore/pull/4128) **sandeepsukhani**: add a storage client for boltdb-shipper which would do all the object key management for storage operations
* [4126](https://example.com/acme/logstore/pull/4126) **cstyan**: Allow for logstore-canary to generate a percentage of out of order log lines
* [4114](https://example.com/acme/logstore/pull/4114) **owen-d**: Stream iterators account for unordered data
* [4111](https://example.com/acme/logstore/pull/4111) **owen-d**: ingester.index-shards config
* [4107](https://example.com/acme/logstore/pull/4107) **sandeepsukhani**: fix finding tables which would have out of retention data
* [4104](https://example.com/acme/logstore/pull/4104) **owen-d**: Discard/ooo
* [4071](https://example.com/acme/logstore/pull/4071) **jeschkies**: Support frontend V2 with query scheduler.

#### Promtail

* [4599](https://example.com/acme/logstore/pull/4599) **rsteneteg**: [Promtail] resolve issue with promtail not scraping target if only path changed in a simpler way that dont need mutex to sync threads
* [4588](https://example.com/acme/logstore/pull/4588) **owen-d**: regenerates assets from current vfsgen dependency
* [4568](https://example.com/acme/logstore/pull/4568) **cyriltovena**: Promtail Kafka target
* [4567](https://example.com/acme/logstore/pull/4567) **cyriltovena**: Refactor client configs in Promtail.
* [4556](https://example.com/acme/logstore/pull/4556) **james-callahan**: promtail: no need for GCP promtail_instance label now that logstore supports out-of-order writes
* [4516](https://example.com/acme/logstore/pull/4516) **lizzzcai**: promtail: update promtail base image to debian:bullseye-slim
* [4507](https://example.com/acme/logstore/pull/4507) **dannykopping**: Promtail: allow for customisable stream lag labels
* [4495](https://example.com/acme/logstore/pull/4495) **sankalp-r**: Promtail: add static labels in stage
* [4461](https://example.com/acme/logstore/pull/4461) **rsteneteg**: Promtail: fix filetarget to not be stuck if no files was detected on startup
* [4346](https://example.com/acme/logstore/pull/4346) **sandeepsukhani**: add logfmt promtail stage to be able to extract data from logfmt formatted log
* [4336](https://example.com/acme/logstore/pull/4336) **ldb**: clients/promtail: Add ndjson and plaintext formats to logstore_push
* [4235](https://example.com/acme/logstore/pull/4235) **kavirajk**: Add metrics for gcplog scrape.
* [3907](https://example.com/acme/logstore/pull/3907) **johanfleury**: promtail: add support for TLS/mTLS in syslog receiver

#### Logcli

* [4303](https://example.com/acme/logstore/pull/4303) **cyriltovena**: Allow to run local boltdb queries with logcli.
* [4242](https://example.com/acme/logstore/pull/4242) **chaudum**: cli: Register configuration option `store.max-look-back-period` as CLI argument
* [4203](https://example.com/acme/logstore/pull/4203) **invidian**: cmd/logcli: add --follow flag as an alias for --tail

#### Build

* [4639](https://example.com/acme/logstore/pull/4639) **slim-bean**: Build: simplify how protos are built
* [4609](https://example.com/acme/logstore/pull/4609) **slim-bean**: Build: Update CODEOWNERS to put Karen back in charge of the docs
* [4541](https://example.com/acme/logstore/pull/4541) **cstyan**: Fix drone ECR publish.
* [4481](https://example.com/acme/logstore/pull/4481) **cstyan**: Update golang and logstore-build-image image versions.
* [4480](https://example.com/acme/logstore/pull/4480) **cstyan**: Add drone build job for lambda-promtail images.
* [4462](https://example.com/acme/logstore/pull/4462) **cstyan**: Update logstore-build-image to drone 1.4.0
* [4373](https://example.com/acme/logstore/pull/4373) **jeschkies**: Instruct how to sign `drone.yml`.
* [4358](https://example.com/acme/logstore/pull/4358) **JordanRushing**: Add DroneCI pipeline stage to validate logstore example configs; create example configuration files
* [4353](https://example.com/acme/logstore/pull/4353) **dannykopping**: CI: Fixing linter deprecations
* [4286](https://example.com/acme/logstore/pull/4286) **slim-bean**: Build: Tweak stalebot message
* [4252](https://example.com/acme/logstore/pull/4252) **slim-bean**: Build: update stalebot message to be more descriptive and friendlier
* [4226](https://example.com/acme/logstore/pull/4226) **aknuds1**: Makefile: Add format target
* [4220](https://example.com/acme/logstore/pull/4220) **slim-bean**: Build: Add github action backport workflow
* [4189](https://example.com/acme/logstore/pull/4189) **mathew-fleisch**: Makefile: Add darwin/arm64 build to release binaries

#### Project

* [4535](https://example.com/acme/logstore/pull/4535) **carlpett**: Fix branch reference in PR template
* [4604](https://example.com/acme/logstore/pull/4604) **kavirajk**: Update PR template to include `changelog` update in the checklist
* [4494](https://example.com/acme/logstore/pull/4494) **cstyan**: Add a a parameter to keep/drop the stream label from cloudwatch.
* [4315](https://example.com/acme/logstore/pull/4315) **cstyan**: Rewrite lambda-promtail to use subscription filters.

#### Dashboards

* [4634](https://example.com/acme/logstore/pull/4634) **cyriltovena**: Fixes the operational dashboard using an old metric.
* [4618](https://example.com/acme/logstore/pull/4618) **cstyan**: logstore-mixin: fix label selectors + logs dashboard
* [4575](https://example.com/acme/logstore/pull/4575) **dannykopping**: Adding recording rules dashboard
* [4441](https://example.com/acme/logstore/pull/4441) **owen-d**: Revert "logstore-mixin: use centralized configuration for dashboard matchers / selectors"
* [4438](https://example.com/acme/logstore/pull/4438) **dannykopping**: Dashboards: adding "logs" into regex
* [4423](https://example.com/acme/logstore/pull/4423) **cstyan**: Add tag/link fix to operational dashboard and promtail mixin dashboard.
* [4401](https://example.com/acme/logstore/pull/4401) **cstyan**: Minor dashboard fixes

#### Docker-driver

* [4396](https://example.com/acme/logstore/pull/4396) **owen-d**: Removes docker driver empty log line message
* [4190](https://example.com/acme/logstore/pull/4190) **jeschkies**: Document known Docker driver issues.

#### FluentD

* [4261](https://example.com/acme/logstore/pull/4261) **MrWong99**: FluentD output plugin: Remove an unused variable when processing chunks

#### Docs

* [4646](https://example.com/acme/logstore/pull/4646) **KMiller-Acme**: Docs: revise modes of operation section
* [4631](https://example.com/acme/logstore/pull/4631) **kavirajk**: Add changelog and upgrade guide for #4556
* [4616](https://example.com/acme/logstore/pull/4616) **owen-d**: index-gw sts doc fix. closes #4583
* [4612](https://example.com/acme/logstore/pull/4612) **surdaft**: Docs: Fix typo in docs
* [4611](https://example.com/acme/logstore/pull/4611) **KMiller-Acme**: Docs: revise incendiary language added in PR 4507
* [4601](https://example.com/acme/logstore/pull/4601) **mustafacansevinc**: docs: fix promtail docs links in logstore installation page
* [4597](https://example.com/acme/logstore/pull/4597) **owen-d**: a few doc fixes in preparation for 2.4
* [4590](https://example.com/acme/logstore/pull/4590) **owen-d**: improves grouping docs examples
* [4579](https://example.com/acme/logstore/pull/4579) **DylanGuedes**: Docs: Modify modes of operation image
* [4576](https://example.com/acme/logstore/pull/4576) **DylanGuedes**: Rename hybrid mode to simple scalable mode
* [4566](https://example.com/acme/logstore/pull/4566) **dannykopping**: Documenting recording rules per-tenant WAL
* [4565](https://example.com/acme/logstore/pull/4565) **DylanGuedes**: Docs: Add virtual targets docs
* [4559](https://example.com/acme/logstore/pull/4559) **chri2547**: docs: Update curl POST  example in docs
* [4548](https://example.com/acme/logstore/pull/4548) **cstyan**: Improve lambda-promtail docs based on Owens review.
* [4540](https://example.com/acme/logstore/pull/4540) **JordanRushing**: Update CHANGELOG.md and /docs with info on new `overrides-exporter` module for Logstore
* [4539](https://example.com/acme/logstore/pull/4539) **cstyan**: Modify lambda-promtail docs based on rewrite.
* [4527](https://example.com/acme/logstore/pull/4527) **yangkb09**: Docs: add missing quote to log_queries.md
* [4521](https://example.com/acme/logstore/pull/4521) **owen-d**: brings storage architecture up to date
* [4499](https://example.com/acme/logstore/pull/4499) **vdm**: Docs: Remove ListObjects S3 permission
* [4493](https://example.com/acme/logstore/pull/4493) **DylanGuedes**: Docs: Move rule storages configs to their own sections
* [4486](https://example.com/acme/logstore/pull/4486) **KMiller-Acme**: Docs: correct the page parameter in the Acme Cloud advertisement
* [4485](https://example.com/acme/logstore/pull/4485) **DylanGuedes**: Document the common config section
* [4422](https://example.com/acme/logstore/pull/4422) **KMiller-Acme**: Docs: revise wording of Acme Cloud advertisement
* [4417](https://example.com/acme/logstore/pull/4417) **KMiller-Acme**: Docs: remove empty section "Generic placeholders"
* [4416](https://example.com/acme/logstore/pull/4416) **KMiller-Acme**: Docs: correctly represent product name
* [4403](https://example.com/acme/logstore/pull/4403) **KMiller-Acme**: Docs: introduce a fundamentals section
* [4399](https://example.com/acme/logstore/pull/4399) **KMiller-Acme**: Docs: prominently advertise free Acme Cloud availability
* [4374](https://example.com/acme/logstore/pull/4374) **KMiller-Acme**: Docs: clarify distinction between single binary and microservices.
* [4363](https://example.com/acme/logstore/pull/4363) **KMiller-Acme**: Docs: Remove wording like "As of version 1.6, you can..."
* [4361](https://example.com/acme/logstore/pull/4361) **JasonGiedymin**: fix(docs): spelling mistake
* [4357](https://example.com/acme/logstore/pull/4357) **carehart**: Correct typo
* [4345](https://example.com/acme/logstore/pull/4345) **pr0PM**: Deduplicating the compactor docs
* [4342](https://example.com/acme/logstore/pull/4342) **KMiller-Acme**: Docs: Organize and edit the LogQL section
* [4324](https://example.com/acme/logstore/pull/4324) **lingenavd**: Docs: Update _index.md to add value boltdb-shipper for the key store
* [4320](https://example.com/acme/logstore/pull/4320) **KMiller-Acme**: Docs: improve spelling, grammar, and formatting.
* [4310](https://example.com/acme/logstore/pull/4310) **dannykopping**: Correcting documentation example for `/api/prom/query`
* [4309](https://example.com/acme/logstore/pull/4309) **GneyHabub**: Docs: Fix a link
* [4294](https://example.com/acme/logstore/pull/4294) **mr-karan**: docs:  (logs-deletion.md) URL Encode curl command
* [4293](https://example.com/acme/logstore/pull/4293) **Birdi7**: docs: fix link to Promtail documentation
* [4283](https://example.com/acme/logstore/pull/4283) **SeriousM**: Correct the indention for azure storage configuration
* [4277](https://example.com/acme/logstore/pull/4277) **ivanahuckova**: Update example for /series endpoint in _index.md
* [4247](https://example.com/acme/logstore/pull/4247) **KMiller-Acme**: Docs: inject newlines for configuration section readability
* [4245](https://example.com/acme/logstore/pull/4245) **KMiller-Acme**: Docs: revise max_query_lookback knob definition
* [4244](https://example.com/acme/logstore/pull/4244) **JordanRushing**: Update limits_config docs to include querier.max_query_lookback flag
* [4237](https://example.com/acme/logstore/pull/4237) **KMiller-Acme**: Docs: first draft, Logstore accepts out-of-order writes
* [4231](https://example.com/acme/logstore/pull/4231) **Aletor93**: doc: fix typo on logstore-external-labels for docker client labels
* [4222](https://example.com/acme/logstore/pull/4222) **KMiller-Acme**: Docs: minor improvements to Logstore Canary docs
* [4208](https://example.com/acme/logstore/pull/4208) **cstyan**: Update tanka installation docs to refer to tanka section about `jb`
* [4206](https://example.com/acme/logstore/pull/4206) **jeschkies**: Link Kubernetes service discovery configuration.
* [4199](https://example.com/acme/logstore/pull/4199) **owen-d**: fixes typo
* [4184](https://example.com/acme/logstore/pull/4184) **mcdeck**: Update docker.md
* [4175](https://example.com/acme/logstore/pull/4175) **KMiller-Acme**: Docs: correct path to Promtail configuration file
* [4163](https://example.com/acme/logstore/pull/4163) **smuth4**: Docs: Update docker install to work out of the box
* [4152](https://example.com/acme/logstore/pull/4152) **charles-woshicai**: Docs: example about using azure storage account as storage
* [4147](https://example.com/acme/logstore/pull/4147) **KMiller-Acme**: Docs: fluentd client phrasing and formatting
* [4145](https://example.com/acme/logstore/pull/4145) **KMiller-Acme**: Docs: improve LogQL section
* [4134](https://example.com/acme/logstore/pull/4134) **KMiller-Acme**: Docs: revise section header (out of order writes)
* [4131](https://example.com/acme/logstore/pull/4131) **owen-d**: updates unordered writes config docs
* [4125](https://example.com/acme/logstore/pull/4125) **owen-d**: Initial out of order docs
* [4122](https://example.com/acme/logstore/pull/4122) **yasharne**: update boltdb-shipper index period
* [4120](https://example.com/acme/logstore/pull/4120) **vitaliyf**: Docs: Fix broken "Upgrading" link
* [4113](https://example.com/acme/logstore/pull/4113) **KMiller-Acme**: Docs: Fix typos and grammar. Inject newlines for readability.
* [4112](https://example.com/acme/logstore/pull/4112) **slim-bean**: Docs: updated changelog and references to 2.3
* [4100](https://example.com/acme/logstore/pull/4100) **jeschkies**: Document operation with the query scheduler.
* [4088](https://example.com/acme/logstore/pull/4088) **KMiller-Acme**: Update Logstore README with better links and descriptions
* [3880](https://example.com/acme/logstore/pull/3880) **timothydlister**: Update fluent-plugin-logstore documentation URLs

#### Jsonnet

* [4629](https://example.com/acme/logstore/pull/4629) **owen-d**: Default wal to enabled in jsonnet lib
* [4624](https://example.com/acme/logstore/pull/4624) **chaudum**: Disable chunk transfers in jsonnet lib
* [4530](https://example.com/acme/logstore/pull/4530) **owen-d**: Jsonnet/overrides exporter
* [4496](https://example.com/acme/logstore/pull/4496) **jeschkies**: Use different metrics for `PromtailFileLagging`.
* [4405](https://example.com/acme/logstore/pull/4405) **jdbaldry**: fix: Correct acme-token creation command
* [4279](https://example.com/acme/logstore/pull/4279) **kevinschoonover**: logstore-mixin: use centralized configuration for dashboard matchers / selectors
* [4259](https://example.com/acme/logstore/pull/4259) **eamonryan**: Jsonnet: Update license path argument name
* [4217](https://example.com/acme/logstore/pull/4217) **Duologic**: fix(rules): upstream recording rule switched to sum_irate
* [4182](https://example.com/acme/logstore/pull/4182) **owen-d**: fine tune grpc configs jsonnet
* [4180](https://example.com/acme/logstore/pull/4180) **owen-d**: corrects query scheduler image
* [4165](https://example.com/acme/logstore/pull/4165) **jdbaldry**: Jsonnet: Add Acme Enterprise Logs library
* [4154](https://example.com/acme/logstore/pull/4154) **owen-d**: updates scheduler libsonnet
* [4102](https://example.com/acme/logstore/pull/4102) **jeschkies**: Define ksonnet lib for query scheduler.

### Notes

This release was created from a branch starting at commit e95d193acf1633a6ec33a328b8a4a3d844e8e5f9 but it may also contain backported changes from main.

Check the history of the branch `release-2.4`.

### Dependencies

* Go Version:     1.17.2
* Corestore Version: 3f329a21cad432325268717eecf2b77c8d95150f

# 2.3.0 (2021/08/06)

Release notes for 2.3.0 can be found on the [release notes page](https://acme.com/docs/logstore/latest/release-notes/v2-3/)

### All Changes

#### Logstore

* [4048](https://example.com/acme/logstore/pull/4048) **dannykopping**: Ruler: implementing write relabelling on recording rule samples
* [4091](https://example.com/acme/logstore/pull/4091) **cyriltovena**: Fixes instant queries in the frontend.
* [4087](https://example.com/acme/logstore/pull/4087) **cyriltovena**: Fixes unaligned shards between ingesters and storage.
* [4047](https://example.com/acme/logstore/pull/4047) **cyriltovena**: Add min_sharding_lookback limits to the frontends
* [4027](https://example.com/acme/logstore/pull/4027) **jdbaldry**: fix: Restore /config endpoint and correct handlerFunc for buildinfo
* [4020](https://example.com/acme/logstore/pull/4020) **simonswine**: Restrict path segments in TenantIDs (CVE-2021-36156 CVE-2021-36157)
* [4019](https://example.com/acme/logstore/pull/4019) **cyriltovena**: Improve decoding of JSON responses.
* [4018](https://example.com/acme/logstore/pull/4018) **sandeepsukhani**: Compactor improvements
* [4017](https://example.com/acme/logstore/pull/4017) **aknuds1**: Chore: Upgrade Prometheus and Corestore
* [3996](https://example.com/acme/logstore/pull/3996) **owen-d**: fixes a badly referenced variable name in StepEvaluator code
* [3995](https://example.com/acme/logstore/pull/3995) **owen-d**: Headblock interop
* [3992](https://example.com/acme/logstore/pull/3992) **MichelHollands**: Update Corestore version
* [3991](https://example.com/acme/logstore/pull/3991) **periklis**: Add LogQL AST walker
* [3990](https://example.com/acme/logstore/pull/3990) **cyriltovena**: Intern label keys for LogQL parser.
* [3986](https://example.com/acme/logstore/pull/3986) **kavirajk**: Ip matcher for LogQL
* [3984](https://example.com/acme/logstore/pull/3984) **jeschkies**: Filter instant queries and shard them.
* [3983](https://example.com/acme/logstore/pull/3983) **cyriltovena**: Reject labels with invalid runes when using implicit extraction parser.
* [3981](https://example.com/acme/logstore/pull/3981) **owen-d**: fixes chunk size method in facade
* [3979](https://example.com/acme/logstore/pull/3979) **MichelHollands**: Add a chunk filterer field to the config
* [3977](https://example.com/acme/logstore/pull/3977) **sandeepsukhani**: add a metric for counting number of failures in opening existing active index files
* [3976](https://example.com/acme/logstore/pull/3976) **sandeepsukhani**: fix flaky retention tests
* [3974](https://example.com/acme/logstore/pull/3974) **owen-d**: WAL Replay counter
* [3973](https://example.com/acme/logstore/pull/3973) **56quarters**: Use the Corestore wrapper for getting tenant ID from a context
* [3972](https://example.com/acme/logstore/pull/3972) **jeschkies**: Return build info under `/logstore/api/v1/status/buildinfo`.
* [3970](https://example.com/acme/logstore/pull/3970) **sandeepsukhani**: log name of the file failed to open during startup by ingester
* [3969](https://example.com/acme/logstore/pull/3969) **sandeepsukhani**: add some tests in compactor and fix a bug in IntervalHasExpiredChunks check in retention with tests
* [3968](https://example.com/acme/logstore/pull/3968) **cyriltovena**: Improve head chunk allocations when reading samples.
* [3967](https://example.com/acme/logstore/pull/3967) **sandeepsukhani**: fix a panic in compactor when retention is not enabled
* [3966](https://example.com/acme/logstore/pull/3966) **sandeepsukhani**: fix panic in compactor when retention is not enabled
* [3957](https://example.com/acme/logstore/pull/3957) **owen-d**: Unordered head block
* [3949](https://example.com/acme/logstore/pull/3949) **cyriltovena**: Allow no overrides config for tenants.
* [3946](https://example.com/acme/logstore/pull/3946) **cyriltovena**: Improve marker file current time metrics.
* [3934](https://example.com/acme/logstore/pull/3934) **sandeepsukhani**: optimize table retetion
* [3932](https://example.com/acme/logstore/pull/3932) **Timbus**: Parser: Allow literal control chars in logfmt decoder
* [3929](https://example.com/acme/logstore/pull/3929) **sandeepsukhani**: remove boltdb files from ingesters on startup which do not have a index bucket
* [3928](https://example.com/acme/logstore/pull/3928) **dannykopping**: Querier/Ingester: Fixing json expression parser bug
* [3919](https://example.com/acme/logstore/pull/3919) **github-vincent-miszczak**: Add ingester.autoforget-unhealthy-timeout opt-in feature
* [3888](https://example.com/acme/logstore/pull/3888) **kavirajk**: Make `overrides` configmap names and mount path as variables.
* [3871](https://example.com/acme/logstore/pull/3871) **kavirajk**: Add explict syntax for using `pattern` parser
* [3865](https://example.com/acme/logstore/pull/3865) **sandeepsukhani**: feat: index-gateway for boltdb-shipper index store
* [3856](https://example.com/acme/logstore/pull/3856) **cyriltovena**: Shards Series API.
* [3852](https://example.com/acme/logstore/pull/3852) **cyriltovena**: Shard ingester queries.
* [3849](https://example.com/acme/logstore/pull/3849) **cyriltovena**: Logs ingester and store queries boundaries.
* [3840](https://example.com/acme/logstore/pull/3840) **cyriltovena**: Add retention label to logstore_distributor_bytes_received_total metrics
* [3837](https://example.com/acme/logstore/pull/3837) **cyriltovena**: LogQL: Pattern Parser
* [3835](https://example.com/acme/logstore/pull/3835) **sesky4**: lz4: update lz4 version to v4.1.7 to avoid possibly panic
* [3833](https://example.com/acme/logstore/pull/3833) **cyriltovena**: Fixes a flaky retention test.
* [3827](https://example.com/acme/logstore/pull/3827) **sandeepsukhani**: Logs deletion fixes
* [3816](https://example.com/acme/logstore/pull/3816) **dannykopping**: Extracting queue interface
* [3807](https://example.com/acme/logstore/pull/3807) **dannykopping**: Logstore: allow for multiple targets
* [3797](https://example.com/acme/logstore/pull/3797) **dannykopping**: Exposing remote writer for use in integration tests
* [3792](https://example.com/acme/logstore/pull/3792) **MichelHollands**: Add a QueryFrontendTripperware module
* [3785](https://example.com/acme/logstore/pull/3785) **sandeepsukhani**: just log a warning when a store type other than boltdb-shipper is detected when custom retention is enabled
* [3772](https://example.com/acme/logstore/pull/3772) **sandeepsukhani**: initialize retention and deletion components only when they are enabled
* [3771](https://example.com/acme/logstore/pull/3771) **sandeepsukhani**: revendor corestore to latest master
* [3769](https://example.com/acme/logstore/pull/3769) **sandeepsukhani**: reduce allocs in delete requests manager by reusing slice for tracing non-deleted intervals for chunks
* [3766](https://example.com/acme/logstore/pull/3766) **dannykopping**: Ruler: Recording Rules
* [3763](https://example.com/acme/logstore/pull/3763) **cyriltovena**: Fixes parser labels hint for grouping.
* [3762](https://example.com/acme/logstore/pull/3762) **cyriltovena**: Improve mark file processing.
* [3758](https://example.com/acme/logstore/pull/3758) **owen-d**: exposes logstore codec
* [3746](https://example.com/acme/logstore/pull/3746) **sandeepsukhani**: Boltdb shipper deletion fixes
* [3743](https://example.com/acme/logstore/pull/3743) **cyriltovena**: Replace satori.uuid with gofrs/uuid
* [3736](https://example.com/acme/logstore/pull/3736) **cyriltovena**: Add fromJson to the template stage.
* [3733](https://example.com/acme/logstore/pull/3733) **cyriltovena**: Fixes a goroutine leak in the store when doing cancellation.
* [3706](https://example.com/acme/logstore/pull/3706) **cyriltovena**: Improve retention mark files.
* [3700](https://example.com/acme/logstore/pull/3700) **slim-bean**: Logstore: Add a flag for queriers to run standalone and only query store
* [3693](https://example.com/acme/logstore/pull/3693) **cyriltovena**: Removes file sync syscall for compaction.
* [3688](https://example.com/acme/logstore/pull/3688) **sandeepsukhani**: Logs deletion support for boltdb-shipper store
* [3687](https://example.com/acme/logstore/pull/3687) **cyriltovena**: Use model.Duration for easy yaml/json marshalling.
* [3686](https://example.com/acme/logstore/pull/3686) **cyriltovena**: Fixes a panic with the frontend when use with downstream URL.
* [3677](https://example.com/acme/logstore/pull/3677) **cyriltovena**: Deprecate max_look_back_period in the chunk storage.
* [3673](https://example.com/acme/logstore/pull/3673) **cyriltovena**: Pass in the now value to the retention.
* [3672](https://example.com/acme/logstore/pull/3672) **cyriltovena**: Use pgzip in the compactor.
* [3665](https://example.com/acme/logstore/pull/3665) **cyriltovena**: Trigger compaction prior retention.
* [3664](https://example.com/acme/logstore/pull/3664) **owen-d**: revendor compatibility: various prom+k8s+corestore
* [3643](https://example.com/acme/logstore/pull/3643) **cyriltovena**: Rejects push requests with  streams without labels.
* [3642](https://example.com/acme/logstore/pull/3642) **cyriltovena**: Custom Retention
* [3641](https://example.com/acme/logstore/pull/3641) **owen-d**: removes naming collision
* [3632](https://example.com/acme/logstore/pull/3632) **kavirajk**: replace `time.Duration` -> `model.Duration` for `Limits`.
* [3628](https://example.com/acme/logstore/pull/3628) **kavirajk**: Add json struct tags to limits.
* [3627](https://example.com/acme/logstore/pull/3627) **MichelHollands**: Update corestore to 1.8
* [3623](https://example.com/acme/logstore/pull/3623) **slim-bean**: Logstore/Promtail: Client Refactor
* [3619](https://example.com/acme/logstore/pull/3619) **liguozhong**: [ui] add '/config' page
* [3618](https://example.com/acme/logstore/pull/3618) **MichelHollands**: Add interceptor override and make ingester and cfg public
* [3605](https://example.com/acme/logstore/pull/3605) **sandeepsukhani**: cleanup boltdb files failing to open during loading tables which are possibly corrupt
* [3603](https://example.com/acme/logstore/pull/3603) **cyriltovena**: Adds chunk filter hook for ingesters.
* [3602](https://example.com/acme/logstore/pull/3602) **MichelHollands**: Loli: Make the store field public
* [3595](https://example.com/acme/logstore/pull/3595) **owen-d**: locks trailers during iteration
* [3594](https://example.com/acme/logstore/pull/3594) **owen-d**: adds distributor replication factor metric
* [3573](https://example.com/acme/logstore/pull/3573) **cyriltovena**: Fixes a race when using specific tenant and multi-client.
* [3569](https://example.com/acme/logstore/pull/3569) **cyriltovena**: Add a chunk filter hook in the store.
* [3566](https://example.com/acme/logstore/pull/3566) **cyriltovena**: Properly release the ticker in Logstore client.
* [3564](https://example.com/acme/logstore/pull/3564) **cyriltovena**: Improve matchers validations.
* [3563](https://example.com/acme/logstore/pull/3563) **sandeepsukhani**: ignore download of missing boltdb files possibly removed during compaction
* [3562](https://example.com/acme/logstore/pull/3562) **cyriltovena**: Fixes a test from #3216.
* [3553](https://example.com/acme/logstore/pull/3553) **cyriltovena**: Add a target to reproduce fuzz testcase
* [3550](https://example.com/acme/logstore/pull/3550) **cyriltovena**: Fixes a bug in MatrixStepper when sharding queries.
* [3549](https://example.com/acme/logstore/pull/3549) **MichelHollands**: LBAC changes
* [3544](https://example.com/acme/logstore/pull/3544) **alrs**: single import of jsoniter in logql subpackages
* [3540](https://example.com/acme/logstore/pull/3540) **cyriltovena**: Support for single step metric query.
* [3532](https://example.com/acme/logstore/pull/3532) **MichelHollands**: Logstore: Update corestore version and fix resulting changes
* [3530](https://example.com/acme/logstore/pull/3530) **sandeepsukhani**: split series api queries by day in query-frontend
* [3517](https://example.com/acme/logstore/pull/3517) **cyriltovena**: Fixes a race introduced by #3434.
* [3515](https://example.com/acme/logstore/pull/3515) **cyriltovena**: Add sprig text/template functions to template stage.
* [3509](https://example.com/acme/logstore/pull/3509) **sandeepsukhani**: fix live tailing of logs from Logstore
* [3572](https://example.com/acme/logstore/pull/3572) **slim-bean**: Logstore: Distributor log message bodySize should always reflect the compressed size
* [3496](https://example.com/acme/logstore/pull/3496) **owen-d**: reduce replay flush threshold
* [3491](https://example.com/acme/logstore/pull/3491) **sandeepsukhani**: make prefix for keys of objects created by boltdb-shipper configurable
* [3487](https://example.com/acme/logstore/pull/3487) **cyriltovena**: Set the byte slice cap correctly when unsafely converting string.
* [3471](https://example.com/acme/logstore/pull/3471) **cyriltovena**: Set a max size for the logql parser to 5k.
* [3470](https://example.com/acme/logstore/pull/3470) **cyriltovena**: Fixes Issue 28593: logstore:fuzz_parse_expr: Timeout in fuzz_parse_expr.
* [3469](https://example.com/acme/logstore/pull/3469) **cyriltovena**: Fixes out-of-memory fuzzing issue.
* [3466](https://example.com/acme/logstore/pull/3466) **pracucci**: Upgrade Corestore
* [3455](https://example.com/acme/logstore/pull/3455) **garrettlish**: Implement offset modifier for range vector aggregation in LogQL
* [3434](https://example.com/acme/logstore/pull/3434) **adityacs**: support math functions in line_format and label_format
* [3216](https://example.com/acme/logstore/pull/3216) **sandeepsukhani**: check for stream selectors to have atleast one equality matcher
* [3050](https://example.com/acme/logstore/pull/3050) **cyriltovena**: first_over_time and last_over_time

#### Docs

* [4031](https://example.com/acme/logstore/pull/4031) **KMiller-Acme**: Docs: add weights to YAML metadata to order the LogQL subsections
* [4029](https://example.com/acme/logstore/pull/4029) **bearice**: Docs: Update S3 permissions list
* [4026](https://example.com/acme/logstore/pull/4026) **KMiller-Acme**: Docs: correct fluentbit config value for DqueSync
* [4024](https://example.com/acme/logstore/pull/4024) **KMiller-Acme**: Docs: fix bad links
* [4016](https://example.com/acme/logstore/pull/4016) **lizzzcai**: <docs>:fix typo in remote debugging docs
* [4012](https://example.com/acme/logstore/pull/4012) **KMiller-Acme**: Revise portions of the docs LogQL section
* [3998](https://example.com/acme/logstore/pull/3998) **owen-d**: Fixes regexReplaceAll docs
* [3980](https://example.com/acme/logstore/pull/3980) **KMiller-Acme**: Docs: Revise/update the overview section.
* [3965](https://example.com/acme/logstore/pull/3965) **mamil**: fix typos
* [3962](https://example.com/acme/logstore/pull/3962) **KMiller-Acme**: Docs: added new target (docs-next) to the docs' Makefile.
* [3956](https://example.com/acme/logstore/pull/3956) **sandeepsukhani**: add config and documentation about index-gateway
* [3938](https://example.com/acme/logstore/pull/3938) **seiffert**: Doc: List 'compactor' as valid value for target option
* [3936](https://example.com/acme/logstore/pull/3936) **lukahartwig**: Fix typo
* [3921](https://example.com/acme/logstore/pull/3921) **KMiller-Acme**: Docs: revise the LogCLI subsection
* [3911](https://example.com/acme/logstore/pull/3911) **KMiller-Acme**: Docs: Make identification of experimental items consistent and obvious
* [3910](https://example.com/acme/logstore/pull/3910) **KMiller-Acme**: Docs: add structure for a release notes section
* [3909](https://example.com/acme/logstore/pull/3909) **kavirajk**: Sync `main` branch docs to `next` folder
* [3899](https://example.com/acme/logstore/pull/3899) **KMiller-Acme**: Docs: correct “ and ” with " and same with single quote mark.
* [3897](https://example.com/acme/logstore/pull/3897) **kavirajk**: Update steps to release versioned docs
* [3882](https://example.com/acme/logstore/pull/3882) **KMiller-Acme**: Docs: improve section on building from source
* [3876](https://example.com/acme/logstore/pull/3876) **ivanahuckova**: Documentation: Unify spelling of backtick in documentation
* [3873](https://example.com/acme/logstore/pull/3873) **KMiller-Acme**: Docs: remove duplicated arch info from the overview section
* [3875](https://example.com/acme/logstore/pull/3875) **kavirajk**: Add missing `-querier.max-concurrent` config in the doc
* [3868](https://example.com/acme/logstore/pull/3868) **sanadhis**: docs: http_path_prefix as correct item of server_config
* [3860](https://example.com/acme/logstore/pull/3860) **KMiller-Acme**: Docs: Correct capitalization and formatting of "Promtail"
* [3851](https://example.com/acme/logstore/pull/3851) **dannykopping**: Ruler: documentation for recording rules
* [3846](https://example.com/acme/logstore/pull/3846) **crockk**: Docs: Minor syntax tweaks for consistency
* [3843](https://example.com/acme/logstore/pull/3843) **azuwis**: multiline: Add regex stage example and note
* [3829](https://example.com/acme/logstore/pull/3829) **arempter**: Add oauth2 docs options for promtail client
* [3828](https://example.com/acme/logstore/pull/3828) **julienduchesne**: Fix broken link in `Windows Event Log` scraping docs
* [3826](https://example.com/acme/logstore/pull/3826) **sandeepsukhani**: docs for logs deletion feature
* [3824](https://example.com/acme/logstore/pull/3824) **KMiller-Acme**: Docs: add and order missing design docs
* [3823](https://example.com/acme/logstore/pull/3823) **KMiller-Acme**: Docs: updates
* [3815](https://example.com/acme/logstore/pull/3815) **paketb0te**: Docs: fixed typo in "Logstore compared to other log systems" (levels -> labels)
* [3810](https://example.com/acme/logstore/pull/3810) **alegmal**: documentation:  corrected double "the the" in index.md
* [3799](https://example.com/acme/logstore/pull/3799) **bt909**: docs: Add memached_client parameter "addresses" list
* [3798](https://example.com/acme/logstore/pull/3798) **bt909**: docs: Change redis configuration value for enabling TLS to correct syntax
* [3790](https://example.com/acme/logstore/pull/3790) **KMiller-Acme**: Docs: remove unnecessary lists of sections
* [3775](https://example.com/acme/logstore/pull/3775) **cyriltovena**: Retention doc
* [3764](https://example.com/acme/logstore/pull/3764) **slim-bean**: Docs: fix makefile
* [3757](https://example.com/acme/logstore/pull/3757) **fionaliao**: [docs] Remove unnecessary backtick from example
* [3756](https://example.com/acme/logstore/pull/3756) **fredrikekre**: [docs] add LogstoreLogger.jl to unofficial clients.
* [3723](https://example.com/acme/logstore/pull/3723) **oddlittlebird**: Docs: Update _index.md
* [3720](https://example.com/acme/logstore/pull/3720) **fredrikekre**: [docs/clients] fix header for "Unofficial clients" and add a reference to said section.
* [3715](https://example.com/acme/logstore/pull/3715) **jaddqiu**: Update troubleshooting.md
* [3714](https://example.com/acme/logstore/pull/3714) **kavirajk**: Fluent-bit git repo link fix
* [3713](https://example.com/acme/logstore/pull/3713) **cyriltovena**: Add a target to find dead links in our documentation.
* [3690](https://example.com/acme/logstore/pull/3690) **atxviking**: API Documentation: Fix document links for /logstore/api/v1/push example
* [3655](https://example.com/acme/logstore/pull/3655) **trevorwhitney**: Documentation: add note about wildcard log patterns and log rotation
* [3648](https://example.com/acme/logstore/pull/3648) **Ruppsn**: Update labels.md in Logstore Docs
* [3647](https://example.com/acme/logstore/pull/3647) **3Xpl0it3r**: fix the promtail-default-config download link in doc
* [3644](https://example.com/acme/logstore/pull/3644) **periklis**: Add Red Hat to adopters
* [3633](https://example.com/acme/logstore/pull/3633) **osg-acme**: Fix wget link.
* [3596](https://example.com/acme/logstore/pull/3596) **timazet**: documentation: typo correction
* [3578](https://example.com/acme/logstore/pull/3578) **liguozhong**: [doc] mtric -> metric
* [3576](https://example.com/acme/logstore/pull/3576) **sergeykranga**: Promtail documentation: fix template example for regexReplaceAll function
* [3568](https://example.com/acme/logstore/pull/3568) **MichelHollands**: docs: some small docs fixes
* [3559](https://example.com/acme/logstore/pull/3559) **klausenbusk**: Doc: Remove removed --ingester.recover-from-wal option and fix out-of-date defaults
* [3555](https://example.com/acme/logstore/pull/3555) **samjewell**: LogQL Docs: Remove key-value pair missing from logfmt output
* [3552](https://example.com/acme/logstore/pull/3552) **lkokila**: Update README.md
* [3551](https://example.com/acme/logstore/pull/3551) **cyriltovena**: Fixes doc w/r/t grpc compression.
* [3542](https://example.com/acme/logstore/pull/3542) **kavirajk**: Remove memberlist config from ring config.
* [3529](https://example.com/acme/logstore/pull/3529) **Whyeasy**: Added docs for GCP internal labels.
* [3525](https://example.com/acme/logstore/pull/3525) **robbymilo**: docs: add title to Lambda Promtail
* [3516](https://example.com/acme/logstore/pull/3516) **cyriltovena**: Fixes broken link in the documentation.
* [3513](https://example.com/acme/logstore/pull/3513) **owen-d**: fixes broken link
* [3543](https://example.com/acme/logstore/pull/3543) **owen-d**: compactor docs
* [3526](https://example.com/acme/logstore/pull/3526) **wardbekker**: Added Architecture Diagram
* [3518](https://example.com/acme/logstore/pull/3518) **wardbekker**: fix spelling in doc
* [3503](https://example.com/acme/logstore/pull/3503) **cyriltovena**: Update README.md
* [3484](https://example.com/acme/logstore/pull/3484) **thomasrockhu**: Add Codecov badge to README
* [3478](https://example.com/acme/logstore/pull/3478) **chancez**: docs/upgrading: Fix typo
* [3477](https://example.com/acme/logstore/pull/3477) **slim-bean**: Jsonnet/Docs: update for 2.2 release
* [3472](https://example.com/acme/logstore/pull/3472) **aronisstav**: Docs: Fix markdown for promtail's output stage
* [3464](https://example.com/acme/logstore/pull/3464) **camilleryr**: Documentation: Update boltdb-shipper.md to fix typo
* [3442](https://example.com/acme/logstore/pull/3442) **owen-d**: adds deprecation notice for chunk transfers
* [3430](https://example.com/acme/logstore/pull/3430) **kavirajk**: doc(gcplog): Add note on scraping multiple GCP projects

#### Promtail

* [4011](https://example.com/acme/logstore/pull/4011) **dannykopping**: Promtail: adding pipeline stage inspector
* [4006](https://example.com/acme/logstore/pull/4006) **dannykopping**: Promtail: output timestamp with nanosecond precision in dry-run mode
* [3971](https://example.com/acme/logstore/pull/3971) **cyriltovena**: Fixes negative gauge in Promtail.
* [3834](https://example.com/acme/logstore/pull/3834) **trevorwhitney**: Promtail: add consul agent service discovery
* [3711](https://example.com/acme/logstore/pull/3711) **3Xpl0it3r**: add debug information for extracted data
* [3683](https://example.com/acme/logstore/pull/3683) **kbudde**: promtail: added timezone to logger in dry-run mode #3679"
* [3654](https://example.com/acme/logstore/pull/3654) **cyriltovena**: Adds the ability to provide a tripperware to Promtail client.
* [3587](https://example.com/acme/logstore/pull/3587) **rsteneteg**: Promtail: Remove non-ready filemanager targets
* [3501](https://example.com/acme/logstore/pull/3501) **kavirajk**: Add unique promtail_instance id to labels for gcptarget
* [3457](https://example.com/acme/logstore/pull/3457) **nmiculinic**: Promtail: Added path information to deleted tailed file
* [3400](https://example.com/acme/logstore/pull/3400) **adityacs**: support max_message_length configuration for syslog parser

#### Logcli

* [3879](https://example.com/acme/logstore/pull/3879) **vyzigold**: logcli: Add retries to unsuccessful log queries
* [3749](https://example.com/acme/logstore/pull/3749) **dbluxo**: logcli: add support for bearer token authentication
* [3739](https://example.com/acme/logstore/pull/3739) **rsteneteg**: correct logcli instant query timestamp param name
* [3678](https://example.com/acme/logstore/pull/3678) **cyriltovena**: Add the ability to wrap the roundtripper of the logcli client.

#### Build

* [4034](https://example.com/acme/logstore/pull/4034) **aknuds1**: logstore-build-image: Fix building
* [4028](https://example.com/acme/logstore/pull/4028) **aknuds1**: logstore-build-image: Upgrade golangci-lint and Go
* [4007](https://example.com/acme/logstore/pull/4007) **dannykopping**: Adding @acme/logstore-team as default CODEOWNERS
* [3997](https://example.com/acme/logstore/pull/3997) **owen-d**: aligns rule path in docker img with bundled config. closes #3952
* [3950](https://example.com/acme/logstore/pull/3950) **julienduchesne**: Sign drone.yml file
* [3944](https://example.com/acme/logstore/pull/3944) **jeschkies**: Lint script files.
* [3941](https://example.com/acme/logstore/pull/3941) **cyriltovena**: Development Docker Compose Setup
* [3935](https://example.com/acme/logstore/pull/3935) **ecraven**: Makefile: Only set PROMTAIL_CGO if CGO_ENABLED is not 0.
* [3832](https://example.com/acme/logstore/pull/3832) **julienduchesne**: Add step to identify windows Drone runner
* [3731](https://example.com/acme/logstore/pull/3731) **cyriltovena**: Fix website branch to trigger update.
* [3708](https://example.com/acme/logstore/pull/3708) **julienduchesne**: Deploy logstore with Drone plugin instead of CircleCI job
* [3703](https://example.com/acme/logstore/pull/3703) **darkn3rd**: Update docker.md for 2.2.1
* [3625](https://example.com/acme/logstore/pull/3625) **slim-bean**: Build: Update CI for branch rename to main
* [3624](https://example.com/acme/logstore/pull/3624) **slim-bean**: Build: Fix drone dependencies on manifest step
* [3615](https://example.com/acme/logstore/pull/3615) **slim-bean**: Remove codecov
* [3481](https://example.com/acme/logstore/pull/3481) **slim-bean**: Update Go and Alpine versions

#### Jsonnet

* [4030](https://example.com/acme/logstore/pull/4030) **cyriltovena**: Improve the sweep lag panel in the retention dashboard.
* [3917](https://example.com/acme/logstore/pull/3917) **jvrplmlmn**: refactor(production/ksonnet): Remove kausal from the root element
* [3893](https://example.com/acme/logstore/pull/3893) **sandeepsukhani**: update uid of logstore-deletion dashboard
* [3891](https://example.com/acme/logstore/pull/3891) **sandeepsukhani**: add index-gateway to reads and reads-resources dashboards
* [3877](https://example.com/acme/logstore/pull/3877) **sandeepsukhani**: Fix jsonnet for index-gateway
* [3854](https://example.com/acme/logstore/pull/3854) **cyriltovena**: Fixes Logstore reads dashboard.
* [3848](https://example.com/acme/logstore/pull/3848) **kavirajk**: Add explicit `main` to pull logstore and promtail to install it via Tanka
* [3794](https://example.com/acme/logstore/pull/3794) **sandeepsukhani**: add a dashboard for log deletion requests in logstore
* [3697](https://example.com/acme/logstore/pull/3697) **owen-d**: better operational dashboard annotations via diff logger
* [3658](https://example.com/acme/logstore/pull/3658) **cyriltovena**: Add a dashboard for retention to the logstore-mixin.
* [3601](https://example.com/acme/logstore/pull/3601) **owen-d**: Dashboard/fix operational vars
* [3584](https://example.com/acme/logstore/pull/3584) **sandeepsukhani**: add logstore resource usage dashboard for read and write path

#### Project

* [3963](https://example.com/acme/logstore/pull/3963) **rfratto**: Remove Robert Fratto from list of team members
* [3926](https://example.com/acme/logstore/pull/3926) **cyriltovena**: Add Danny Kopping to the Logstore Team.
* [3732](https://example.com/acme/logstore/pull/3732) **dannykopping**: Issue Templates: Improve wording and add warnings
* [3722](https://example.com/acme/logstore/pull/3722) **oddlittlebird**: Update CODEOWNERS
* [3951](https://example.com/acme/logstore/pull/3951) **owen-d**: update sizing calc
* [3931](https://example.com/acme/logstore/pull/3931) **owen-d**: Hackathon/cluster
* [3920](https://example.com/acme/logstore/pull/3920) **owen-d**: adds replication &  deduping into cost
* [3630](https://example.com/acme/logstore/pull/3630) **slim-bean**: Re-license to AGPLv3

#### Docker-driver

* [3814](https://example.com/acme/logstore/pull/3814) **kavirajk**: Update the docker-driver doc about default labels
* [3727](https://example.com/acme/logstore/pull/3727) **3Xpl0it3r**: docker-driver: remove duplicated code
* [3709](https://example.com/acme/logstore/pull/3709) **cyriltovena**: Fixes docker driver that would panic when closed.

### Notes

This release was created from revision 8012362674568379a3871ff8c4a2bfd1ddba7ad1 (Which was PR 3460)

### Dependencies

* Go Version:     1.16.2
* Corestore Version: 485474c9afb2614fb89af3f48803c37d016bbaed

## 2.2.1 (2021/04/05)

2.2.1 fixes several important bugs, it is recommended everyone running 2.2.0 upgrade to 2.2.1

2.2.1 also adds the `labelallow` pipeline stage in Promtail which lets an allowlist be created for what labels will be sent by Promtail to Logstore.

* [3468](https://example.com/acme/logstore/pull/3468) **adityacs**: Support labelallow stage in Promtail
* [3502](https://example.com/acme/logstore/pull/3502) **cyriltovena**: Fixes a bug where unpack would mutate log line.
* [3540](https://example.com/acme/logstore/pull/3540) **cyriltovena**: Support for single step metric query.
* [3550](https://example.com/acme/logstore/pull/3550) **cyriltovena**: Fixes a bug in MatrixStepper when sharding queries.
* [3566](https://example.com/acme/logstore/pull/3566) **cyriltovena**: Properly release the ticker in Logstore client.
* [3573](https://example.com/acme/logstore/pull/3573) **cyriltovena**: Fixes a race when using specific tenant and multi-client.

## 2.2.0 (2021/03/10)

With over 200 PR's 2.2 includes significant features, performance improvements, and bug fixes!

The most upvoted issue for Logstore was closed in this release! [Issue 74](https://example.com/acme/logstore/issues/74) requesting support for handling multi-line logs in Promtail was implemented in [PR 3024](https://example.com/acme/logstore/pull/3024). Thanks @jeschkies!

Other exciting news for Promtail, [PR 3246](https://example.com/acme/logstore/pull/3246) by @cyriltovena introduces support for reading Windows Events!

Switching to Logstore, @owen-d has added a write ahead log to Logstore! [PR 2981](https://example.com/acme/logstore/pull/2981) was the first of many as we have spent the last several months using and abusing our write ahead logs to flush out any bugs!

A significant number of the PR's in this release have gone to improving the features introduced in Logstore 2.0. @cyriltovena overhauled the JSON parser in [PR 3163](https://example.com/acme/logstore/pull/3163) (and a few other PR's), to provide both a faster and smarter parsing to only extract JSON content which is used in the query output.  The newest Logstore squad member @dannykopping fine tuned the JSON parser options in [PR 3280](https://example.com/acme/logstore/pull/3280) allowing you to specific individual JSON elements, including support now for accessing elements in an array.  Many, many other additional improvements have been made, as well as several fixes to the new LogQL features added some months ago, this upgrade should have everyone seeing improvements in their queries.

@cyriltovena also set his PPROF skills loose on the Logstore write path which resulted in about 8x less memory usage on our distributors and a much more stable memory usage when ingesters are flushing a lot of chunks at the same time.

There are many other noteworthy additions and fixes, too many to list, but we should call out one more feature all you Google Cloud Platform users might be excited about: in [PR 3083](https://example.com/acme/logstore/pull/3083) @kavirajk added support to Promtail for listening on Google Pub/Sub topics, letting you setup log sinks for your GCP logs to be ingested by Promtail and sent to Logstore!

Thanks to everyone for another exciting Logstore release!!

Please read the [Upgrade Guide](https://example.com/acme/logstore/blob/master/docs/sources/setup/upgrade/_index.md#220) before upgrading for a smooth experience.

TL;DR Logstore 2.2 changes the internal chunk format which limits what versions you can downgrade to, a bug in how many queries were allowed to be scheduled per tenant was fixed which might affect your `max_query_parallelism` and `max_outstanding_per_tenant` settings, and we fixed a mistake related `scrape_configs` which do not have a `pipeline_stages` defined. If you have any Promtail `scrape_configs` which do not specify `pipeline_stages` you should go read the upgrade notes!

### All Changes

#### Logstore

* [3460](https://example.com/acme/logstore/pull/3460) **slim-bean**: Logstore: Per Tenant Runtime Configs
* [3459](https://example.com/acme/logstore/pull/3459) **cyriltovena**: Fixes split interval for metrics queries.
* [3432](https://example.com/acme/logstore/pull/3432) **slim-bean**: Logstore: change ReadStringAsSlice to ReadString so it doesn't parse quotes inside the packed _entry
* [3429](https://example.com/acme/logstore/pull/3429) **cyriltovena**: Improve the parser hint tests.
* [3426](https://example.com/acme/logstore/pull/3426) **cyriltovena**: Only unpack entry if the key `_entry` exist.
* [3424](https://example.com/acme/logstore/pull/3424) **cyriltovena**: Add fgprof to Logstore and Promtail.
* [3423](https://example.com/acme/logstore/pull/3423) **cyriltovena**: Add limit and line_returned in the query log.
* [3420](https://example.com/acme/logstore/pull/3420) **cyriltovena**: Introduce a unpack parser.
* [3417](https://example.com/acme/logstore/pull/3417) **cyriltovena**: Fixes a race in the scheduling limits.
* [3416](https://example.com/acme/logstore/pull/3416) **ukolovda**: Distributor: append several tests for HTTP parser.
* [3411](https://example.com/acme/logstore/pull/3411) **slim-bean**: Logstore: fix alignment of atomic 64 bit to work with 32 bit OS
* [3409](https://example.com/acme/logstore/pull/3409) **gotjosh**: Instrumentation: Add histogram for request duration on gRPC client to Ingesters
* [3408](https://example.com/acme/logstore/pull/3408) **jgehrcke**: distributor: fix snappy-compressed protobuf POST request handling (#3407)
* [3388](https://example.com/acme/logstore/pull/3388) **owen-d**: prevents duplicate log lines from being replayed. closes #3378
* [3383](https://example.com/acme/logstore/pull/3383) **cyriltovena**: Fixes head chunk iterator direction.
* [3380](https://example.com/acme/logstore/pull/3380) **slim-bean**: Logstore: Fix parser hint for extracted labels which collide with stream labels
* [3372](https://example.com/acme/logstore/pull/3372) **cyriltovena**: Fixes a panic with whitespace key.
* [3350](https://example.com/acme/logstore/pull/3350) **cyriltovena**: Fixes ingester stats.
* [3348](https://example.com/acme/logstore/pull/3348) **cyriltovena**: Fixes 500 in the querier when returning multiple errors.
* [3347](https://example.com/acme/logstore/pull/3347) **cyriltovena**: Fixes a tight loop in the Engine with LogQL parser.
* [3344](https://example.com/acme/logstore/pull/3344) **cyriltovena**: Fixes some 500 returned by querier when storage cancellation happens.
* [3342](https://example.com/acme/logstore/pull/3342) **cyriltovena**: Bound parallelism frontend
* [3340](https://example.com/acme/logstore/pull/3340) **owen-d**: Adds some flushing instrumentation/logs
* [3339](https://example.com/acme/logstore/pull/3339) **owen-d**: adds Start() method to WAL interface to delay checkpointing until aft…
* [3338](https://example.com/acme/logstore/pull/3338) **sandeepsukhani**: dedupe index on all the queries for a table instead of query batches
* [3326](https://example.com/acme/logstore/pull/3326) **owen-d**: removes wal recover flag
* [3307](https://example.com/acme/logstore/pull/3307) **slim-bean**: Logstore: fix validation error and metrics
* [3306](https://example.com/acme/logstore/pull/3306) **cyriltovena**: Add finalizer to zstd.
* [3300](https://example.com/acme/logstore/pull/3300) **sandeepsukhani**: increase db retain period in ingesters to cover index cache validity period as well
* [3299](https://example.com/acme/logstore/pull/3299) **owen-d**: Logql/absent label optimization
* [3295](https://example.com/acme/logstore/pull/3295) **jtlisi**: chore: update corestore to latest and fix refs
* [3291](https://example.com/acme/logstore/pull/3291) **ukolovda**: Distributor: Logstore API can receive gzipped JSON
* [3280](https://example.com/acme/logstore/pull/3280) **dannykopping**: LogQL: Simple JSON expressions
* [3279](https://example.com/acme/logstore/pull/3279) **cyriltovena**: Fixes logfmt parser hints.
* [3278](https://example.com/acme/logstore/pull/3278) **owen-d**: Testware/ rate-unwrap-multi
* [3274](https://example.com/acme/logstore/pull/3274) **liguozhong**: [ingester_query] change var "clients" to "reps"
* [3267](https://example.com/acme/logstore/pull/3267) **jeschkies**: Update vendored Corestore to 0976147451ee
* [3263](https://example.com/acme/logstore/pull/3263) **cyriltovena**: Fix a bug with  metric queries and label_format.
* [3261](https://example.com/acme/logstore/pull/3261) **sandeepsukhani**: fix broken json logs push path
* [3256](https://example.com/acme/logstore/pull/3256) **jtlisi**: update vendored corestore and add new replace overrides
* [3251](https://example.com/acme/logstore/pull/3251) **cyriltovena**: Ensure we have parentheses for bin ops.
* [3249](https://example.com/acme/logstore/pull/3249) **cyriltovena**: Fixes a bug where slice of Entries where not zeroed
* [3241](https://example.com/acme/logstore/pull/3241) **cyriltovena**: Allocate entries array with correct size  while decoding WAL entries.
* [3237](https://example.com/acme/logstore/pull/3237) **cyriltovena**: Fixes unmarshalling of tailing responses.
* [3236](https://example.com/acme/logstore/pull/3236) **slim-bean**: Logstore: Log a crude lag metric for how far behind a client is.
* [3234](https://example.com/acme/logstore/pull/3234) **cyriltovena**: Fixes previous commit not using the new sized body.
* [3233](https://example.com/acme/logstore/pull/3233) **cyriltovena**: Re-introduce <https://example.com/acme/logstore/pull/3178>.
* [3228](https://example.com/acme/logstore/pull/3228) **MichelHollands**: Add config endpoint
* [3218](https://example.com/acme/logstore/pull/3218) **owen-d**: WAL backpressure
* [3217](https://example.com/acme/logstore/pull/3217) **cyriltovena**: Rename checkpoint proto package to avoid conflict with corestore.
* [3215](https://example.com/acme/logstore/pull/3215) **cyriltovena**: Corestore update pre 1.7
* [3211](https://example.com/acme/logstore/pull/3211) **cyriltovena**: Fixes tail api marshalling for v1.
* [3210](https://example.com/acme/logstore/pull/3210) **cyriltovena**: Reverts flush buffer pooling.
* [3201](https://example.com/acme/logstore/pull/3201) **sandeepsukhani**: limit query range in async store for ingesters when query-ingesters-within flag is set
* [3200](https://example.com/acme/logstore/pull/3200) **cyriltovena**: Improve ingester flush memory usage.
* [3195](https://example.com/acme/logstore/pull/3195) **owen-d**: Ignore flushed chunks during checkpointing
* [3194](https://example.com/acme/logstore/pull/3194) **cyriltovena**: Fixes unwrap expressions from  last optimization.
* [3193](https://example.com/acme/logstore/pull/3193) **cyriltovena**: Improve checkpoint series iterator.
* [3188](https://example.com/acme/logstore/pull/3188) **cyriltovena**: Improves checkpointerWriter memory usage
* [3180](https://example.com/acme/logstore/pull/3180) **owen-d**: moves boltdb flags to config file
* [3178](https://example.com/acme/logstore/pull/3178) **cyriltovena**: Logs PushRequest data.
* [3163](https://example.com/acme/logstore/pull/3163) **cyriltovena**: Uses custom json-iter decoder for log entries.
* [3159](https://example.com/acme/logstore/pull/3159) **MichelHollands**: Make httpAuthMiddleware field public
* [3153](https://example.com/acme/logstore/pull/3153) **cyriltovena**: Improve wal entries encoding.
* [3152](https://example.com/acme/logstore/pull/3152) **AllenzhLi**: update github.com/gorilla/websocket to fixes a potential denial-of-service (DoS) vector
* [3146](https://example.com/acme/logstore/pull/3146) **owen-d**: More semantically correct flush shutdown
* [3143](https://example.com/acme/logstore/pull/3143) **cyriltovena**: Fixes absent_over_time to work with all log selector.
* [3141](https://example.com/acme/logstore/pull/3141) **owen-d**: Swaps mutex for atomic in ingester's OnceSwitch
* [3137](https://example.com/acme/logstore/pull/3137) **owen-d**: label_format no longer shardable and introduces the Shardable() metho…
* [3136](https://example.com/acme/logstore/pull/3136) **owen-d**: Don't fail writes due to full WAL disk
* [3134](https://example.com/acme/logstore/pull/3134) **cyriltovena**: Improve distributors validation and apply in-place filtering.
* [3132](https://example.com/acme/logstore/pull/3132) **owen-d**: Integrates label replace into sharding code
* [3131](https://example.com/acme/logstore/pull/3131) **MichelHollands**: Update corestore to 1 6
* [3126](https://example.com/acme/logstore/pull/3126) **dannykopping**: Implementing line comments
* [3122](https://example.com/acme/logstore/pull/3122) **owen-d**: Self documenting pipeline process interface
* [3117](https://example.com/acme/logstore/pull/3117) **owen-d**: Wal/recover corruption
* [3114](https://example.com/acme/logstore/pull/3114) **owen-d**: Disables the stream limiter until wal has recovered
* [3092](https://example.com/acme/logstore/pull/3092) **liguozhong**: lru cache logql.ParseLabels
* [3090](https://example.com/acme/logstore/pull/3090) **cyriltovena**: Improve tailer matching by using the index.
* [3087](https://example.com/acme/logstore/pull/3087) **MichelHollands**: feature: make server publicly available
* [3080](https://example.com/acme/logstore/pull/3080) **cyriltovena**: Improve JSON parser and add labels parser hints.
* [3077](https://example.com/acme/logstore/pull/3077) **MichelHollands**: Make the moduleManager field public
* [3065](https://example.com/acme/logstore/pull/3065) **cyriltovena**: Optimizes SampleExpr to remove unnecessary line_format.
* [3064](https://example.com/acme/logstore/pull/3064) **cyriltovena**: Add zstd and flate compressions algorithms.
* [3053](https://example.com/acme/logstore/pull/3053) **cyriltovena**: Add absent_over_time
* [3048](https://example.com/acme/logstore/pull/3048) **cyriltovena**: Support rate for unwrapped expressions.
* [3047](https://example.com/acme/logstore/pull/3047) **cyriltovena**: Add function label_replace.
* [3030](https://example.com/acme/logstore/pull/3030) **cyriltovena**: Allows by/without to be empty and available for max/min_over_time
* [3025](https://example.com/acme/logstore/pull/3025) **cyriltovena**: Fixes a swallowed context err in the batch storage.
* [3013](https://example.com/acme/logstore/pull/3013) **owen-d**: headblock checkpointing up to v3
* [3008](https://example.com/acme/logstore/pull/3008) **cyriltovena**: Fixes the ruler storage with  the boltdb store.
* [3000](https://example.com/acme/logstore/pull/3000) **owen-d**: Introduces per stream chunks mutex
* [2981](https://example.com/acme/logstore/pull/2981) **owen-d**: Adds WAL support (experimental)
* [2960](https://example.com/acme/logstore/pull/2960) **sandeepsukhani**: fix table deletion in table client for boltdb-shipper

#### Promtail

* [3422](https://example.com/acme/logstore/pull/3422) **kavirajk**: Modify script to accept inclusion and exclustion filters as variables
* [3404](https://example.com/acme/logstore/pull/3404) **dannykopping**: Remove default docker pipeline stage
* [3401](https://example.com/acme/logstore/pull/3401) **slim-bean**: Promtail: Add pack stage
* [3381](https://example.com/acme/logstore/pull/3381) **adityacs**: fix nested captured groups indexing in replace stage
* [3332](https://example.com/acme/logstore/pull/3332) **cyriltovena**: Embed timezone data in Promtail.
* [3304](https://example.com/acme/logstore/pull/3304) **kavirajk**: Use project-id from the variables. Remove hardcoding
* [3303](https://example.com/acme/logstore/pull/3303) **cyriltovena**: Increase the windows bookmark buffer.
* [3302](https://example.com/acme/logstore/pull/3302) **cyriltovena**: Fixes races in multiline stage and promtail.
* [3298](https://example.com/acme/logstore/pull/3298) **gregorybrzeski**: Promtail: fix typo in config variable name - BookmarkPath
* [3285](https://example.com/acme/logstore/pull/3285) **kavirajk**: Make incoming labels from gcp into Logstore internal labels.
* [3284](https://example.com/acme/logstore/pull/3284) **kavirajk**: Avoid putting all the GCP labels into logstore labels
* [3246](https://example.com/acme/logstore/pull/3246) **cyriltovena**: Windows events
* [3224](https://example.com/acme/logstore/pull/3224) **veltmanj**: Fix(pkg/promtail)  CVE-2020-11022 JQuery vulnerability
* [3207](https://example.com/acme/logstore/pull/3207) **cyriltovena**: Fixes panic when using multiple clients
* [3191](https://example.com/acme/logstore/pull/3191) **rfratto**: promtail: pass registerer to gcplog
* [3175](https://example.com/acme/logstore/pull/3175) **rfratto**: Promtail: pass a prometheus registerer to promtail components
* [3083](https://example.com/acme/logstore/pull/3083) **kavirajk**: Gcplog targetmanager
* [3024](https://example.com/acme/logstore/pull/3024) **jeschkies**: Collapse multiline logs based on a start line.
* [3015](https://example.com/acme/logstore/pull/3015) **cyriltovena**: Add more information about why a tailer would stop.
* [2996](https://example.com/acme/logstore/pull/2996) **cyriltovena**: Asynchronous Promtail stages
* [2898](https://example.com/acme/logstore/pull/2898) **kavirajk**: fix(docker-driver): Propagate promtail's `client.Stop` properly

#### Logcli

* [3325](https://example.com/acme/logstore/pull/3325) **cyriltovena**: Fixes step encoding in logcli.
* [3271](https://example.com/acme/logstore/pull/3271) **chancez**: Refactor logcli Client interface to use time objects for LiveTailQueryConn
* [3270](https://example.com/acme/logstore/pull/3270) **chancez**: logcli: Fix handling of logcli query using --since/--from and --tail
* [3229](https://example.com/acme/logstore/pull/3229) **dethi**: logcli: support --include-label when not using --tail

#### Jsonnet

* [3447](https://example.com/acme/logstore/pull/3447) **owen-d**: Use better memory metric on operational dashboard
* [3439](https://example.com/acme/logstore/pull/3439) **owen-d**: simplifies jsonnet sharding
* [3357](https://example.com/acme/logstore/pull/3357) **owen-d**: Libsonnet/better sharding parallelism defaults
* [3356](https://example.com/acme/logstore/pull/3356) **owen-d**: removes sharding queue math after global concurrency PR
* [3329](https://example.com/acme/logstore/pull/3329) **sandeepsukhani**: fix config for statefulset rulers when using boltdb-shipper
* [3297](https://example.com/acme/logstore/pull/3297) **owen-d**: adds stateful ruler clause for boltdb shipper jsonnet
* [3254](https://example.com/acme/logstore/pull/3254) **hairyhenderson**: ksonnet: Remove invalid hostname from default promtail configuration
* [3181](https://example.com/acme/logstore/pull/3181) **owen-d**: remaining sts use parallel mgmt policy
* [3179](https://example.com/acme/logstore/pull/3179) **owen-d**: Ruler statefulsets
* [3156](https://example.com/acme/logstore/pull/3156) **slim-bean**: Jsonnet: Changes ingester PVC from 5Gi to 10Gi
* [3139](https://example.com/acme/logstore/pull/3139) **owen-d**: Less confusing jsonnet error message when using boltdb shipper defaults.
* [3079](https://example.com/acme/logstore/pull/3079) **rajatvig**: Fix ingester PVC data declaration to use configured value
* [3074](https://example.com/acme/logstore/pull/3074) **c0ffeec0der**: Ksonnet: Assign appropriate pvc size and class to compactor and ingester
* [3062](https://example.com/acme/logstore/pull/3062) **cyriltovena**: Remove regexes in the operational dashboard.
* [3014](https://example.com/acme/logstore/pull/3014) **owen-d**: logstore wal libsonnet
* [3010](https://example.com/acme/logstore/pull/3010) **cyriltovena**: Fixes promtail mixin dashboard.

#### fluentd

* [3358](https://example.com/acme/logstore/pull/3358) **fpob**: Fix fluentd plugin when kubernetes labels were missing

#### fluent bit

* [3240](https://example.com/acme/logstore/pull/3240) **sbaier1**: fix fluent-bit output plugin generating invalid JSON

#### Docker Logging Driver

* [3331](https://example.com/acme/logstore/pull/3331) **cyriltovena**: Add pprof endpoint to docker-driver.
* [3225](https://example.com/acme/logstore/pull/3225) **Le0tk0k**: (fix: cmd/docker-driver): Insert a space in the error message

#### Docs

* [5934](https://example.com/acme/logstore/pull/5934) **johgsc**: Docs: revise modes of operation section
* [3437](https://example.com/acme/logstore/pull/3437) **caleb15**: docs: add note about regex
* [3421](https://example.com/acme/logstore/pull/3421) **kavirajk**: doc(gcplog): Advanced log export filter example
* [3419](https://example.com/acme/logstore/pull/3419) **suitupalex**: docs: promtail: Fix typo w/ windows_events hyperlink.
* [3418](https://example.com/acme/logstore/pull/3418) **dannykopping**: Adding upgrade documentation for promtail pipeline_stages change
* [3385](https://example.com/acme/logstore/pull/3385) **paaacman**: Documentation: enable environment variable in configuration
* [3373](https://example.com/acme/logstore/pull/3373) **StMarian**: Documentation: Fix configuration description
* [3371](https://example.com/acme/logstore/pull/3371) **owen-d**: Distributor overview docs
* [3370](https://example.com/acme/logstore/pull/3370) **tkowalcz**: documentation: Add Tjahzi to the list of unofficial clients
* [3352](https://example.com/acme/logstore/pull/3352) **kavirajk**: Remove extra space between broken link
* [3351](https://example.com/acme/logstore/pull/3351) **kavirajk**: Add some operation details to gcplog doc
* [3316](https://example.com/acme/logstore/pull/3316) **kavirajk**: docs(fix): Make best practices docs look better
* [3292](https://example.com/acme/logstore/pull/3292) **wapmorgan**: Patch 2 - fix link to another documentation files
* [3265](https://example.com/acme/logstore/pull/3265) **sandeepsukhani**: Boltdb shipper doc fixes
* [3239](https://example.com/acme/logstore/pull/3239) **owen-d**: updates tanka installation docs
* [3235](https://example.com/acme/logstore/pull/3235) **scoof**: docs: point to latest release for docker installation
* [3220](https://example.com/acme/logstore/pull/3220) **liguozhong**: [doc] fix err. "logstore_frontend" is invalid
* [3212](https://example.com/acme/logstore/pull/3212) **nvtkaszpir**: Fix: Update docs for logcli
* [3190](https://example.com/acme/logstore/pull/3190) **kavirajk**: doc(gcplog): Fix titles for Cloud provisioning for GCP logs
* [3165](https://example.com/acme/logstore/pull/3165) **liguozhong**: [doc] fix:querier do not handle "/flush" api
* [3164](https://example.com/acme/logstore/pull/3164) **owen-d**: updates alerting docs post 2.0
* [3162](https://example.com/acme/logstore/pull/3162) **huikang**: Doc: Add missing wal in configuration
* [3148](https://example.com/acme/logstore/pull/3148) **huikang**: Doc: add missing type supported by table manager
* [3147](https://example.com/acme/logstore/pull/3147) **marionxue**: Markdown Code highlighting
* [3138](https://example.com/acme/logstore/pull/3138) **jeschkies**: Give another example for multiline.
* [3128](https://example.com/acme/logstore/pull/3128) **cyriltovena**: Fixes LogQL documentation links.
* [3124](https://example.com/acme/logstore/pull/3124) **wujie1993**: fix time duration unit
* [3123](https://example.com/acme/logstore/pull/3123) **scoren-gl**: Update getting-in-touch.md
* [3115](https://example.com/acme/logstore/pull/3115) **valmack**: Docs: Include instruction to enable variable expansion
* [3109](https://example.com/acme/logstore/pull/3109) **nileshcs**: Documentation fix for downstream_url
* [3102](https://example.com/acme/logstore/pull/3102) **slim-bean**: Docs: Changelog: fix indentation and add helm repo url
* [3094](https://example.com/acme/logstore/pull/3094) **benjaminhuo**: Fix storage guide links
* [3088](https://example.com/acme/logstore/pull/3088) **cyriltovena**: Small fixes for the documentation.
* [3084](https://example.com/acme/logstore/pull/3084) **ilpianista**: Update reference to old helm chart repo
* [3078](https://example.com/acme/logstore/pull/3078) **kavirajk**: mention the use of `config.expand-env` flag in the doc.
* [3049](https://example.com/acme/logstore/pull/3049) **vitalets**: [Docs] Clarify docker-driver configuration options
* [3039](https://example.com/acme/logstore/pull/3039) **jdbaldry**: doc: logql formatting fixes
* [3035](https://example.com/acme/logstore/pull/3035) **unguiculus**: Fix multiline docs
* [3033](https://example.com/acme/logstore/pull/3033) **RichiH**: docs: Create ADOPTERS.md
* [3032](https://example.com/acme/logstore/pull/3032) **oddlittlebird**: Docs: Update _index.md
* [3029](https://example.com/acme/logstore/pull/3029) **jeschkies**: Correct `multiline` documentation.
* [3027](https://example.com/acme/logstore/pull/3027) **nop33**: Fix docs header inconsistency
* [3026](https://example.com/acme/logstore/pull/3026) **owen-d**: wal docs
* [3017](https://example.com/acme/logstore/pull/3017) **jdbaldry**: doc: Cleanup formatting
* [3009](https://example.com/acme/logstore/pull/3009) **jdbaldry**: doc: Fix query-frontend typo
* [3002](https://example.com/acme/logstore/pull/3002) **keyolk**: Fix typo
* [2991](https://example.com/acme/logstore/pull/2991) **jontg**: Documentation:  Add a missing field to the extended config s3 example
* [2956](https://example.com/acme/logstore/pull/2956) **owen-d**: Updates chunkenc doc for V3

#### Build

* [3412](https://example.com/acme/logstore/pull/3412) **rfratto**: Remove unneeded prune-ci-tags drone job
* [3390](https://example.com/acme/logstore/pull/3390) **wardbekker**: Don't auto-include pod labels as logstore labels as a sane default
* [3321](https://example.com/acme/logstore/pull/3321) **owen-d**: defaults promtail to 2.1.0 in install script
* [3277](https://example.com/acme/logstore/pull/3277) **kavirajk**: Add step to version Logstore docs during public release process.
* [3243](https://example.com/acme/logstore/pull/3243) **chancez**: dist: Build promtail for windows/386 to support 32 bit windows hosts
* [3206](https://example.com/acme/logstore/pull/3206) **kavirajk**: Terraform script to automate GCP provisioning for gcplog
* [3149](https://example.com/acme/logstore/pull/3149) **jlosito**: Allow dependabot to keep github actions up-to-date
* [3072](https://example.com/acme/logstore/pull/3072) **slim-bean**: Helm: Disable CI
* [3031](https://example.com/acme/logstore/pull/3031) **AdamKorcz**: Testing: Introduced continuous fuzzing
* [3006](https://example.com/acme/logstore/pull/3006) **huikang**: Fix the docker image version in compose deployment

#### Tooling

* [3377](https://example.com/acme/logstore/pull/3377) **slim-bean**: Tooling: Update chunks-inspect to understand the new chunk format as well as new compression algorithms
* [3151](https://example.com/acme/logstore/pull/3151) **slim-bean**: Logstore migrate-tool

### Notes

This release was created from revision 8012362674568379a3871ff8c4a2bfd1ddba7ad1 (Which was PR 3460)

### Dependencies

* Go Version:     1.15.3
* Corestore Version: 7dac81171c665be071bd167becd1f55528a9db32

## 2.1.0 (2020/12/23)

Happy Holidays from the Logstore team! Please enjoy a new Logstore release to welcome in the New Year!

2.1.0 Contains a number of fixes, performance improvements and enhancements to the 2.0.0 release!

### Notable changes

#### Helm users read this

The Helm charts have moved!

* [2720](https://example.com/acme/logstore/pull/2720) **torstenwalter**: Deprecate Charts as they have been moved

This was done to consolidate Acme's helm charts for all Acme projects in one place: <https://example.com/acme/helm-charts/>

**From now moving forward, please use the new Helm repo url: <https://acme.github.io/helm-charts>**

The charts in the Logstore repo will soon be removed so please update your Helm repo to the new URL and submit your PR's over there as well

Special thanks to @torstenwalter, @unguiculus, and @scottrigby for their initiative and amazing work to make this happen!

Also go check out the microservices helm chart contributed by @unguiculus in the new repo!

#### Fluent bit plugin users read this

Fluent bit officially supports Logstore as an output plugin now! WoooHOOO!

However this created a naming conflict with our existing output plugin (the new native output uses the name `logstore`) so we have renamed our plugin.

* [2974](https://example.com/acme/logstore/pull/2974) **hedss**: fluent-bit: Rename Fluent Bit plugin output name.

In time our plan is to deprecate and eliminate our output plugin in favor of the native Logstore support. However until then you can continue using the plugin with the following change:

Old:

```
[Output]
    Name logstore
```

New:

```
[Output]
    Name acme-logstore
```

#### Fixes

A lot of work went into 2.0 with a lot of new code and rewrites to existing, this introduced and uncovered some bugs which are fixed in 2.1:

* [2807](https://example.com/acme/logstore/pull/2807) **cyriltovena**: Fix error swallowed in the frontend.
* [2805](https://example.com/acme/logstore/pull/2805) **cyriltovena**: Improve pipeline stages ast errors.
* [2824](https://example.com/acme/logstore/pull/2824) **owen-d**: Fix/validate compactor config
* [2830](https://example.com/acme/logstore/pull/2830) **sandeepsukhani**: fix panic in ingester when not running with boltdb shipper while queriers does
* [2850](https://example.com/acme/logstore/pull/2850) **owen-d**: Only applies entry limits to non-SampleExprs.
* [2855](https://example.com/acme/logstore/pull/2855) **sandeepsukhani**: fix query intervals when running boltdb-shipper in single binary
* [2895](https://example.com/acme/logstore/pull/2895) **shokada**: Fix error 'Unexpected: ("$", "$") while parsing field definition'
* [2902](https://example.com/acme/logstore/pull/2902) **cyriltovena**: Fixes metric query issue with no grouping.
* [2901](https://example.com/acme/logstore/pull/2901) **cyriltovena**: Fixes a panic with the logql.NoopPipeline.
* [2913](https://example.com/acme/logstore/pull/2913) **cyriltovena**: Fixes logql.QueryType.
* [2917](https://example.com/acme/logstore/pull/2917) **cyriltovena**: Fixes race condition in tailer since logql v2.
* [2960](https://example.com/acme/logstore/pull/2960) **sandeepsukhani**: fix table deletion in table client for boltdb-shipper

#### Enhancements

A number of performance and resource improvements have been made as well!

* [2911](https://example.com/acme/logstore/pull/2911) **sandeepsukhani**: Boltdb shipper query readiness
* [2875](https://example.com/acme/logstore/pull/2875) **cyriltovena**: Labels computation LogQLv2
* [2927](https://example.com/acme/logstore/pull/2927) **cyriltovena**: Improve logql parser allocations.
* [2926](https://example.com/acme/logstore/pull/2926) **cyriltovena**: Cache label strings in ingester to improve memory usage.
* [2931](https://example.com/acme/logstore/pull/2931) **cyriltovena**: Only append tailed entries if needed.
* [2973](https://example.com/acme/logstore/pull/2973) **cyriltovena**: Avoid parsing labels when tailer is sending from a stream.
* [2959](https://example.com/acme/logstore/pull/2959) **cyriltovena**: Improve tailer matcher function.
* [2876](https://example.com/acme/logstore/pull/2876) **jkellerer**: LogQL: Add unwrap bytes() conversion function

#### Notable mentions

Thanks to @timbyr for adding an often requested feature, the ability to support environment variable expansion in config files!

* [2837](https://example.com/acme/logstore/pull/2837) **timbyr**: Configuration: Support environment expansion in configuration

Thanks to @huikang for adding a new docker-compose file for running Logstore as microservices!

* [2740](https://example.com/acme/logstore/pull/2740) **huikang**: Deploy: add docker-compose cluster deployment file

### All Changes

#### Logstore

* [2988](https://example.com/acme/logstore/pull/2988) **slim-bean**: Logstore: handle faults when opening boltdb files
* [2984](https://example.com/acme/logstore/pull/2984) **owen-d**: adds the ability to read chunkFormatV3 while writing v2
* [2983](https://example.com/acme/logstore/pull/2983) **slim-bean**: Logstore: recover from panic opening boltdb files
* [2975](https://example.com/acme/logstore/pull/2975) **cyriltovena**: Fixes vector grouping injection.
* [2972](https://example.com/acme/logstore/pull/2972) **cyriltovena**: Add ProcessString to Pipeline.
* [2962](https://example.com/acme/logstore/pull/2962) **cyriltovena**: Implement io.WriteTo by chunks.
* [2951](https://example.com/acme/logstore/pull/2951) **owen-d**: bumps rules-action ref to logqlv2+ version
* [2946](https://example.com/acme/logstore/pull/2946) **cyriltovena**: Fixes the Stringer of the byte label operator.
* [2945](https://example.com/acme/logstore/pull/2945) **cyriltovena**: Fixes iota unexpected behaviour with bytes for chunk encoding.
* [2941](https://example.com/acme/logstore/pull/2941) **jeschkies**: Test label filter for bytes.
* [2934](https://example.com/acme/logstore/pull/2934) **owen-d**: chunk schema v3
* [2930](https://example.com/acme/logstore/pull/2930) **cyriltovena**: Fixes all in one grpc registrations.
* [2929](https://example.com/acme/logstore/pull/2929) **cyriltovena**: Cleanup labels parsing.
* [2922](https://example.com/acme/logstore/pull/2922) **codewithcheese**: Distributor registers logproto.Pusher service to receive logs via GRPC
* [2918](https://example.com/acme/logstore/pull/2918) **owen-d**: Includes delete routes for ruler namespaces
* [2903](https://example.com/acme/logstore/pull/2903) **cyriltovena**: Limit series for metric queries.
* [2892](https://example.com/acme/logstore/pull/2892) **cyriltovena**: Improve the chunksize test.
* [2891](https://example.com/acme/logstore/pull/2891) **sandeepsukhani**: fix flaky load tables test for boltdb-shipper uploads table-manager
* [2836](https://example.com/acme/logstore/pull/2836) **andir**: tests: fix quoting issues in test output when building with Go 1.15
* [2831](https://example.com/acme/logstore/pull/2831) **sandeepsukhani**: fix flaky tests in boltdb-shipper
* [2822](https://example.com/acme/logstore/pull/2822) **cyriltovena**: LogQL: Improve template format
* [2794](https://example.com/acme/logstore/pull/2794) **sandeepsukhani**: Revendor corestore to latest master
* [2764](https://example.com/acme/logstore/pull/2764) **owen-d**: WAL/marshalable chunks
* [2751](https://example.com/acme/logstore/pull/2751) **jeschkies**: Logging: Log throughput and total bytes human readable.

#### Helm

* [2986](https://example.com/acme/logstore/pull/2986) **cyriltovena**: Move CI to helm3.
* [2967](https://example.com/acme/logstore/pull/2967) **czunker**: Remove `helm init`
* [2965](https://example.com/acme/logstore/pull/2965) **czunker**: [Helm Chart Logstore] Add needed k8s objects for alerting config
* [2940](https://example.com/acme/logstore/pull/2940) **slim-bean**: Helm: Update logstash to new chart and newer version
* [2835](https://example.com/acme/logstore/pull/2835) **tracyde**: Iss2734
* [2789](https://example.com/acme/logstore/pull/2789) **bewiwi**: Allows service targetPort modificaion
* [2651](https://example.com/acme/logstore/pull/2651) **scottrigby**: helm chart: Fix broken logo

#### Jsonnet

* [2976](https://example.com/acme/logstore/pull/2976) **beorn7**: Improve promtail alerts to retain the namespace label
* [2961](https://example.com/acme/logstore/pull/2961) **sandeepsukhani**: add missing ingester query routes in logstore reads and operational dashboard
* [2899](https://example.com/acme/logstore/pull/2899) **halcyondude**: gateway: fix regression in tanka jsonnet
* [2873](https://example.com/acme/logstore/pull/2873) **Duologic**: fix(logstore-mixin): refer to super.annotations
* [2852](https://example.com/acme/logstore/pull/2852) **chancez**: production/ksonnet: Add config_hash annotation to gateway deployment based on gateway configmap
* [2820](https://example.com/acme/logstore/pull/2820) **owen-d**: fixes promtail libsonnet tag. closes #2818
* [2718](https://example.com/acme/logstore/pull/2718) **halcyondude**: parameterize PVC storage class (ingester, querier, compactor)

#### Docs

* [2969](https://example.com/acme/logstore/pull/2969) **simonswine**: Add community forum to README.md
* [2968](https://example.com/acme/logstore/pull/2968) **yuichi10**: logcli: Fix logcli logql document URL
* [2942](https://example.com/acme/logstore/pull/2942) **hedss**: Docs: Corrects Fluent Bit documentation link to build the plugin.
* [2933](https://example.com/acme/logstore/pull/2933) **oddlittlebird**: Update CODEOWNERS
* [2909](https://example.com/acme/logstore/pull/2909) **fredr**: Docs: Add max_cache_freshness_per_query to limit_config
* [2890](https://example.com/acme/logstore/pull/2890) **dfang**: Fix typo
* [2888](https://example.com/acme/logstore/pull/2888) **oddlittlebird**: Update CODEOWNERS
* [2879](https://example.com/acme/logstore/pull/2879) **zhanghjster**: documentation: add tail_proxy_url option to query_frontend_config section
* [2869](https://example.com/acme/logstore/pull/2869) **nehaev**: documentation: Add logstore4j to the list of unofficial clients
* [2853](https://example.com/acme/logstore/pull/2853) **RangerCD**: Fix typos in promtail
* [2848](https://example.com/acme/logstore/pull/2848) **dminca**: documentation: fix broken link in Best Practices section
* [2833](https://example.com/acme/logstore/pull/2833) **siavashs**: Docs: -querier.split-queries-by-day deprecation
* [2819](https://example.com/acme/logstore/pull/2819) **owen-d**: updates docs with delete permissions notice
* [2817](https://example.com/acme/logstore/pull/2817) **scoof**: Documentation: Add S3 IAM policy to be able to run Compactor
* [2811](https://example.com/acme/logstore/pull/2811) **slim-bean**: Docs: improve the helm upgrade section
* [2810](https://example.com/acme/logstore/pull/2810) **hedss**: CHANGELOG: Update update document links to point to the right place.
* [2704](https://example.com/acme/logstore/pull/2704) **owen-d**: WAL design doc
* [2636](https://example.com/acme/logstore/pull/2636) **LTek-online**: promtail documentation: changing the headers of the configuration docu to reflect configuration code

#### Promtail

* [2957](https://example.com/acme/logstore/pull/2957) **slim-bean**: Promtail: Update debian image and use a newer libsystemd
* [2928](https://example.com/acme/logstore/pull/2928) **cyriltovena**: Skip journald bad message.
* [2914](https://example.com/acme/logstore/pull/2914) **chancez**: promtail: Add support for using syslog message timestamp
* [2910](https://example.com/acme/logstore/pull/2910) **rfratto**: Expose underlying promtail client

#### Logcli

* [2948](https://example.com/acme/logstore/pull/2948) **tomwilkie**: Add a few more instructions to logcli --help.

#### Build

* [2877](https://example.com/acme/logstore/pull/2877) **cyriltovena**: Update to go 1.15
* [2814](https://example.com/acme/logstore/pull/2814) **torkelo**: Stats: Adding metrics collector GitHub action

#### Fluentd

* [2825](https://example.com/acme/logstore/pull/2825) **cyriltovena**: Bump fluentd plugin
* [2434](https://example.com/acme/logstore/pull/2434) **andsens**: fluent-plugin: Improve escaping in key_value format

### Notes

This release was created from revision ae9c4b82ec4a5d21267da50d6a1a8170e0ef82ff (Which was PR 2960) and the following PR's were cherry-picked

* [2984](https://example.com/acme/logstore/pull/2984) **owen-d**: adds the ability to read chunkFormatV3 while writing v2
* [2974](https://example.com/acme/logstore/pull/2974) **hedss**: fluent-bit: Rename Fluent Bit plugin output name.

### Dependencies

* Go Version:     1.15.3
* Corestore Version: 85942c5703cf22b64cecfd291e7e7c42d1b8c30c

## 2.0.1 (2020/12/10)

2.0.1 is a special release, it only exists to add the v3 support to Logstore's chunk format.

**There is no reason to upgrade from 2.0.0 to 2.0.1**

This chunk version is internal to Logstore and not configurable, and in a future version v3 will become the default (Likely 2.2.0).

We are creating this to enable users to roll back from a future release which was writing v3 chunks, back as far as 2.0.0 and still be able to read chunks.

This is mostly a safety measure to help if someone upgrades from 2.0.0 and skips versions to a future version which is writing v3 chunks and they encounter an issue which they would like to roll back. They would be able to then roll back to 2.0.1 and still read v3 chunks.

It should be noted this does not help anyone upgrading from a version older than 2.0.0, that is you should at least upgrade to 2.0.0 before going to a newer version if you are on a version older than 2.0.0.

## 2.0.0 (2020/10/26)

2.0.0 is here!!

We are extremely excited about the new features in 2.0.0, unlocking a whole new world of observability of our logs.

Thanks again for the many incredible contributions and improvements from the wonderful Logstore community, we are very excited for the future!

### Important Notes

**Please Note** There are several changes in this release which require your attention!

* Anyone using a docker image please go read the [upgrade guide](https://example.com/acme/logstore/blob/master/docs/sources/setup/upgrade/_index.md#200)!! There is one important consideration around a potentially breaking schema change depending on your configuration.
* MAJOR changes have been made to the boltdb-shipper index, breaking changes are not expected but extra precautions are highly recommended, more details in the [upgrade guide](https://example.com/acme/logstore/blob/master/docs/sources/setup/upgrade/_index.md#200).
* The long deprecated `entry_parser` config in Promtail has been removed, use [pipeline_stages](https://acme.com/docs/logstore/latest/clients/promtail/configuration/#pipeline_stages) instead.

Check the [upgrade guide](https://example.com/acme/logstore/blob/master/docs/sources/setup/upgrade/_index.md#200) for detailed information on all these changes.

### 2.0

There are too many PR's to list individually for the major improvements which we thought justified a 2.0 but here is the high level:

* Significant enhancements to the [LogQL query language](https://acme.com/docs/logstore/latest/logql/)!
* [Parse](https://acme.com/docs/logstore/latest/logql/#parser-expression) your logs to extract labels at query time.
* [Filter](https://acme.com/docs/logstore/latest/logql/#label-filter-expression) on query time extracted labels.
* [Format](https://acme.com/docs/logstore/latest/logql/#line-format-expression) your log lines any way you please!
* [Graph](https://acme.com/docs/logstore/latest/logql/#unwrapped-range-aggregations) the contents of your log lines as metrics, including support for many more of your favorite PromQL functions.
* Generate prometheus [alerts directly from your logs](https://acme.com/docs/logstore/latest/alerting/)!
* Create alerts using the same prometheus alert rule syntax and let Logstore send alerts directly to your Prometheus Alertmanager!
* [boltdb-shipper](https://acme.com/docs/logstore/latest/operations/storage/boltdb-shipper/) is now production ready!
* This is it! Now Logstore only needs a single object store (S3,GCS,Filesystem...) to store all the data, no more Cassandra, DynamoDB or Bigtable!

We are extremely excited about these new features, expect some talks, webinars, and blogs where we explain all this new functionality in detail.

### Notable mention

This is a small change but very helpful!

* [2737](https://example.com/acme/logstore/pull/2737) **dlemel8**: cmd/logstore: add "verify-config" flag

Thank you @dlemel8 for this PR! Now you can start Logstore with `-verify-config` to make sure your config is valid and Logstore will exit with a status code 0 if it is!

### All Changes

#### Logstore

* [2804](https://example.com/acme/logstore/pull/2804) **slim-bean**: Logstore: log any chunk fetch failure
* [2803](https://example.com/acme/logstore/pull/2803) **slim-bean**: Update local and docker default config files to use boltdb-shipper with a few other config changes
* [2796](https://example.com/acme/logstore/pull/2796) **cyriltovena**: Fixes a bug that would add **error** label incorrectly.
* [2793](https://example.com/acme/logstore/pull/2793) **cyriltovena**: Improve the way we reverse iterator for backward queries.
* [2790](https://example.com/acme/logstore/pull/2790) **sandeepsukhani**: Boltdb shipper metrics changes
* [2788](https://example.com/acme/logstore/pull/2788) **sandeepsukhani**: add a metric in compactor to record timestamp of last successful run
* [2786](https://example.com/acme/logstore/pull/2786) **cyriltovena**: Logqlv2 pushes groups down to edge
* [2778](https://example.com/acme/logstore/pull/2778) **cyriltovena**: Logqv2 optimization
* [2774](https://example.com/acme/logstore/pull/2774) **cyriltovena**: Handle panic in the store goroutine.
* [2773](https://example.com/acme/logstore/pull/2773) **cyriltovena**: Fixes race conditions in the batch iterator.
* [2770](https://example.com/acme/logstore/pull/2770) **sandeepsukhani**: Boltdb shipper query performance improvements
* [2769](https://example.com/acme/logstore/pull/2769) **cyriltovena**: LogQL: Labels and Metrics Extraction
* [2768](https://example.com/acme/logstore/pull/2768) **cyriltovena**: Fixes all lint errors.
* [2761](https://example.com/acme/logstore/pull/2761) **owen-d**: Service discovery refactor
* [2755](https://example.com/acme/logstore/pull/2755) **owen-d**: Revendor Corestore
* [2752](https://example.com/acme/logstore/pull/2752) **kavirajk**: fix: Remove depricated `entry_parser` from scrapeconfig
* [2741](https://example.com/acme/logstore/pull/2741) **owen-d**: better tenant logging in ruler memstore
* [2737](https://example.com/acme/logstore/pull/2737) **dlemel8**: cmd/logstore: add "verify-config" flag
* [2735](https://example.com/acme/logstore/pull/2735) **cyriltovena**: Fixes the frontend logs to include org_id.
* [2732](https://example.com/acme/logstore/pull/2732) **sandeepsukhani**: set timestamp in instant query done by canaries
* [2726](https://example.com/acme/logstore/pull/2726) **dvrkps**: hack: clean getStore
* [2711](https://example.com/acme/logstore/pull/2711) **owen-d**: removes r/w pools from block/chunk types
* [2709](https://example.com/acme/logstore/pull/2709) **cyriltovena**: Bypass sharding middleware when a query can't be sharded.
* [2671](https://example.com/acme/logstore/pull/2671) **alrs**: pkg/querier: fix dropped error
* [2665](https://example.com/acme/logstore/pull/2665) **cnbailian**: Logstore: Querier APIs respond JSON Content-Type
* [2663](https://example.com/acme/logstore/pull/2663) **owen-d**: improves numeric literal stringer impl
* [2662](https://example.com/acme/logstore/pull/2662) **owen-d**: exposes rule group validation fn
* [2661](https://example.com/acme/logstore/pull/2661) **owen-d**: Enable local rules backend & disallow configdb.
* [2656](https://example.com/acme/logstore/pull/2656) **sandeepsukhani**: run multiple queries per table at once with boltdb-shipper
* [2655](https://example.com/acme/logstore/pull/2655) **sandeepsukhani**: fix store query bug when running logstore in single binary mode with boltdb-shipper
* [2650](https://example.com/acme/logstore/pull/2650) **owen-d**: Adds prometheus ruler routes
* [2647](https://example.com/acme/logstore/pull/2647) **arl**: pkg/chunkenc: fix test using string(int) conversion
* [2645](https://example.com/acme/logstore/pull/2645) **arl**: Tests: fix issue 2356: distributor_test.go fails when the system has no interface name in [eth0, en0, lo0]
* [2642](https://example.com/acme/logstore/pull/2642) **sandeepsukhani**: fix an issue with building logstore
* [2640](https://example.com/acme/logstore/pull/2640) **sandeepsukhani**: improvements for boltdb-shipper compactor
* [2637](https://example.com/acme/logstore/pull/2637) **owen-d**: Ruler docs + single binary inclusion
* [2627](https://example.com/acme/logstore/pull/2627) **sandeepsukhani**: revendor corestore to latest master
* [2620](https://example.com/acme/logstore/pull/2620) **alrs**: pkg/storage/stores/shipper/uploads: fix test error
* [2614](https://example.com/acme/logstore/pull/2614) **cyriltovena**: Improve lz4 compression
* [2613](https://example.com/acme/logstore/pull/2613) **sandeepsukhani**: fix a panic when trying to stop boltdb-shipper multiple times using sync.once
* [2610](https://example.com/acme/logstore/pull/2610) **slim-bean**: Logstore: Fix query-frontend ready handler
* [2601](https://example.com/acme/logstore/pull/2601) **sandeepsukhani**: rpc for querying ingesters to get chunk ids from its store
* [2589](https://example.com/acme/logstore/pull/2589) **owen-d**: Ruler/logstore rule validator
* [2582](https://example.com/acme/logstore/pull/2582) **yeya24**: Add _total suffix to ruler counter metrics
* [2580](https://example.com/acme/logstore/pull/2580) **owen-d**: strict rule unmarshaling
* [2578](https://example.com/acme/logstore/pull/2578) **owen-d**: exports grouploader
* [2576](https://example.com/acme/logstore/pull/2576) **owen-d**: Better rule loading
* [2574](https://example.com/acme/logstore/pull/2574) **sandeepsukhani**: fix closing of compressed file from boltdb-shipper compactor
* [2572](https://example.com/acme/logstore/pull/2572) **adityacs**: Validate max_query_length in Labels API
* [2564](https://example.com/acme/logstore/pull/2564) **owen-d**: Error on no schema configs
* [2559](https://example.com/acme/logstore/pull/2559) **sandeepsukhani**: fix dir setup based on which mode it is running
* [2558](https://example.com/acme/logstore/pull/2558) **sandeepsukhani**: cleanup boltdb files in queriers during startup/shutdown
* [2552](https://example.com/acme/logstore/pull/2552) **owen-d**: fixes batch metrics help text & corrects bucketing
* [2550](https://example.com/acme/logstore/pull/2550) **sandeepsukhani**: fix a flaky test in boltdb shipper
* [2548](https://example.com/acme/logstore/pull/2548) **sandeepsukhani**: add some metrics for monitoring compactor
* [2546](https://example.com/acme/logstore/pull/2546) **sandeepsukhani**: register boltdb shipper compactor cli flags
* [2543](https://example.com/acme/logstore/pull/2543) **sandeepsukhani**: revendor corestore to latest master
* [2534](https://example.com/acme/logstore/pull/2534) **owen-d**: Consistent chunk metrics
* [2530](https://example.com/acme/logstore/pull/2530) **sandeepsukhani**: minor fixes and improvements for boltdb shipper
* [2526](https://example.com/acme/logstore/pull/2526) **sandeepsukhani**: compactor for compacting boltdb files uploaded by shipper
* [2510](https://example.com/acme/logstore/pull/2510) **owen-d**: adds batch based metrics
* [2507](https://example.com/acme/logstore/pull/2507) **sandeepsukhani**: compress boltdb files to gzip while uploading from shipper
* [2458](https://example.com/acme/logstore/pull/2458) **owen-d**: Feature/ruler (take 2)
* [2487](https://example.com/acme/logstore/pull/2487) **sandeepsukhani**: upload boltdb files from shipper only when they are not expected to be modified or during shutdown

#### Docs

* [2797](https://example.com/acme/logstore/pull/2797) **cyriltovena**: Logqlv2 docs
* [2772](https://example.com/acme/logstore/pull/2772) **DesistDaydream**: reapir Retention Example Configuration
* [2762](https://example.com/acme/logstore/pull/2762) **PabloCastellano**: fix: typo in upgrade.md
* [2750](https://example.com/acme/logstore/pull/2750) **owen-d**: fixes path in prom rules api docs
* [2733](https://example.com/acme/logstore/pull/2733) **owen-d**: Removes wrong capitalizations
* [2728](https://example.com/acme/logstore/pull/2728) **vishesh92**: Docs: Update docs for redis
* [2725](https://example.com/acme/logstore/pull/2725) **dvrkps**: fix some misspells
* [2724](https://example.com/acme/logstore/pull/2724) **MadhavJivrajani**: DOCS: change format of unordered lists in technical docs
* [2716](https://example.com/acme/logstore/pull/2716) **huikang**: Doc: fixing parameter name in configuration
* [2705](https://example.com/acme/logstore/pull/2705) **owen-d**: shows corestoretool lint command for logstore in alerting docs
* [2702](https://example.com/acme/logstore/pull/2702) **huikang**: Doc: fix broken links in production/README.md
* [2699](https://example.com/acme/logstore/pull/2699) **sandangel**: docs: use repetitive numbering
* [2698](https://example.com/acme/logstore/pull/2698) **bemasher**: Doc: Vague link text.
* [2697](https://example.com/acme/logstore/pull/2697) **owen-d**: updates alerting docs with new corestore tool logstore linting support
* [2692](https://example.com/acme/logstore/pull/2692) **philnichol**: Docs: Corrected incorrect instances of (setup|set up)
* [2691](https://example.com/acme/logstore/pull/2691) **UniqueTokens**: Update metrics.md
* [2689](https://example.com/acme/logstore/pull/2689) **pgassmann**: docker plugin documentation update
* [2686](https://example.com/acme/logstore/pull/2686) **demon**: docs: Fix link to code of conduct
* [2657](https://example.com/acme/logstore/pull/2657) **owen-d**: fixes ruler docs & includes ruler configs in cmd/configs + docker img
* [2622](https://example.com/acme/logstore/pull/2622) **sandeepsukhani**: add compactor details and other boltdb-shipper doc improvments
* [2621](https://example.com/acme/logstore/pull/2621) **cyriltovena**: Fixes links in aws tutorials.
* [2606](https://example.com/acme/logstore/pull/2606) **cyriltovena**: More template stage examples.
* [2605](https://example.com/acme/logstore/pull/2605) **Decad**: Update docs to use raw link
* [2600](https://example.com/acme/logstore/pull/2600) **slim-bean**: Docs: Fix broken links on generated site
* [2597](https://example.com/acme/logstore/pull/2597) **nek-00-ken**: Fixup: url to access promtail config sample
* [2595](https://example.com/acme/logstore/pull/2595) **sh0rez**: docs: fix broken links
* [2594](https://example.com/acme/logstore/pull/2594) **wardbekker**: Update README.md
* [2592](https://example.com/acme/logstore/pull/2592) **owen-d**: fixes some doc links
* [2591](https://example.com/acme/logstore/pull/2591) **woodsaj**: Docs: fix links in installation docs
* [2586](https://example.com/acme/logstore/pull/2586) **ms42Q**: Doc fixes: remove typos and long sentence
* [2579](https://example.com/acme/logstore/pull/2579) **oddlittlebird**: Update CODEOWNERS
* [2566](https://example.com/acme/logstore/pull/2566) **owen-d**: Website doc link fixes
* [2528](https://example.com/acme/logstore/pull/2528) **owen-d**: Update tanka.md with steps for using k8s-alpha lib
* [2512](https://example.com/acme/logstore/pull/2512) **palemtnrider**: Documentation: Fixes  install and getting-started links in the readme
* [2508](https://example.com/acme/logstore/pull/2508) **owen-d**: memberlist correct yaml path. closes #2499
* [2506](https://example.com/acme/logstore/pull/2506) **ferdikurniawan**: Docs: fix dead link
* [2505](https://example.com/acme/logstore/pull/2505) **sh0rez**: doc: close code block
* [2501](https://example.com/acme/logstore/pull/2501) **tivvit**: fix incorrect upgrade link
* [2500](https://example.com/acme/logstore/pull/2500) **oddlittlebird**: Docs: Update README.md

#### Helm

* [2746](https://example.com/acme/logstore/pull/2746) **marcosartori**: helm/fluentbit K8S-Logging.Exclude &  and Mem_Buf_Limit toggle
* [2742](https://example.com/acme/logstore/pull/2742) **steven-sheehy**: Fix linting errors and use of deprecated repositories
* [2659](https://example.com/acme/logstore/pull/2659) **rskrishnar**: [Promtail] enables configuring psp in helm chart
* [2554](https://example.com/acme/logstore/pull/2554) **alexandre-allard-scality**: production/helm: add support for PV selector in Logstore statefulset

#### FluentD

* [2739](https://example.com/acme/logstore/pull/2739) **jgehrcke**: FluentD logstore plugin: add support for bearer_token_file parameter

#### Fluent Bit

* [2568](https://example.com/acme/logstore/pull/2568) **zjj2wry**: fluent-bit plugin support TLS

#### Promtail

* [2723](https://example.com/acme/logstore/pull/2723) **carlpett**: Promtail: Add counter promtail_batch_retries_total
* [2717](https://example.com/acme/logstore/pull/2717) **slim-bean**: Promtail: Fix deadlock on tailer shutdown.
* [2710](https://example.com/acme/logstore/pull/2710) **slim-bean**: Promtail: (and also fluent-bit) change the max batch size to 1MB
* [2708](https://example.com/acme/logstore/pull/2708) **Falco20019**: Promtail: Fix timestamp parser for short year format
* [2658](https://example.com/acme/logstore/pull/2658) **slim-bean**: Promtail: do not mark the position if the file is removed
* [2618](https://example.com/acme/logstore/pull/2618) **slim-bean**: Promtail: Add a stream lagging metric
* [2615](https://example.com/acme/logstore/pull/2615) **aminjam**: Add fallback_formats for timestamp stage
* [2603](https://example.com/acme/logstore/pull/2603) **rfratto**: Expose UserAgent and fix User-Agent version source
* [2575](https://example.com/acme/logstore/pull/2575) **unguiculus**: Promtail: Fix docker-compose.yaml
* [2571](https://example.com/acme/logstore/pull/2571) **rsteneteg**: Promtail: adding pipeline stage for dropping labels
* [2570](https://example.com/acme/logstore/pull/2570) **slim-bean**: Promtail: Fix concurrent map iteration when using stdin
* [2565](https://example.com/acme/logstore/pull/2565) **carlpett**: Add a counter for empty syslog messages
* [2542](https://example.com/acme/logstore/pull/2542) **slim-bean**: Promtail: implement shutdown for the no-op server
* [2532](https://example.com/acme/logstore/pull/2532) **slim-bean**: Promtail: Restart the tailer if we fail to read and upate current position

#### Ksonnet

* [2719](https://example.com/acme/logstore/pull/2719) **halcyondude**: nit: fix formatting for ksonnet/logstore
* [2677](https://example.com/acme/logstore/pull/2677) **sandeepsukhani**: fix jsonnet for memcached-writes when using boltdb-shipper
* [2617](https://example.com/acme/logstore/pull/2617) **periklis**: Add config options for logstore dashboards
* [2612](https://example.com/acme/logstore/pull/2612) **fredr**: Dashboard: typo in Logstore Operational dashboard
* [2599](https://example.com/acme/logstore/pull/2599) **sandeepsukhani**: fix closing bracket in dashboards from logstore-mixin
* [2584](https://example.com/acme/logstore/pull/2584) **sandeepsukhani**: Read, Write and operational dashboard improvements
* [2560](https://example.com/acme/logstore/pull/2560) **owen-d**: Jsonnet/ruler
* [2547](https://example.com/acme/logstore/pull/2547) **sandeepsukhani**: jsonnet for running logstore using boltdb-shipper
* [2525](https://example.com/acme/logstore/pull/2525) **Duologic**: fix(ksonnet): don't depend on specific k8s version
* [2521](https://example.com/acme/logstore/pull/2521) **charandas**: fix: broken links in Tanka documentation
* [2503](https://example.com/acme/logstore/pull/2503) **owen-d**: Ksonnet docs
* [2494](https://example.com/acme/logstore/pull/2494) **primeroz**: Jsonnet Promtail: Change function for mounting configmap in promtail daemonset

#### Logstash

* [2607](https://example.com/acme/logstore/pull/2607) **adityacs**: Logstash cpu usage fix

#### Build

* [2602](https://example.com/acme/logstore/pull/2602) **sandeepsukhani**: add support for building querytee
* [2561](https://example.com/acme/logstore/pull/2561) **tharun208**: Added logcli docker image
* [2549](https://example.com/acme/logstore/pull/2549) **simnv**: Ignore .exe files build for Windows
* [2527](https://example.com/acme/logstore/pull/2527) **owen-d**: Update docker-compose.yaml to use 1.6.0

#### Docker Logging Driver

* [2459](https://example.com/acme/logstore/pull/2459) **RaitoBezarius**: Docker logging driver: Add a keymod for the extra attributes from the Docker logging driver

### Dependencies

* Go Version:     1.14.2
* Corestore Version: 85942c5703cf22b64cecfd291e7e7c42d1b8c30c

## 1.6.1 (2020-08-24)

This is a small release and only contains two fixes for Promtail:

* [2542](https://example.com/acme/logstore/pull/2542) **slim-bean**: Promtail: implement shutdown for the no-op server
* [2532](https://example.com/acme/logstore/pull/2532) **slim-bean**: Promtail: Restart the tailer if we fail to read and upate current position

The first only applies if you are running Promtail with both `--stdin` and `--server.disabled=true` flags.

The second is a minor rework to how Promtail handles a very specific error when attempting to read the size of a file and failing to do so.

Upgrading Promtail from 1.6.0 to 1.6.1 is only necessary if you have logs full of `msg="error getting tail position and/or size"`,
the code changed in this release has been unchanged for a long time and we suspect very few people are seeing this issue.

No changes to any other components (Logstore, Logcli, etc) are included in this release.

## 1.6.0 (2020-08-13)

It's the second thursday of the eighth month of the year which means it's time for another Logstore Release!!

Before we highlight important features and changes, congratulations to [@adityacs](https://github.com/adityacs), who is the newest member of the Logstore team!
Aditya has been regularly contributing to the Logstore project for the past year, with each contribution better than the last.
Many of the items on the following list were thanks to his hard work. Thank you, Aditya, and welcome to the team!

I think we might have set a new record with 189 PR's in this release!

### Important Notes

**Please Note** There are several changes in this release which might require your attention!

* The NET_BIND_SERVICE capability was removed from the Logstore process in the docker image, it's no longer possible to run Logstore with the supplied image on a port less than 1024
* If you run microservices, there is an important rollout sequence to prevent query errors.
* Scrape configs have changed for Promtail in both Helm and Ksonnet affecting two labels: `instance` -> `pod` and `container_name` -> `container`.
* Almost all of the Logstore Canary metrics were renamed.
* A few command line flags where changed (although they are likely not commonly used)
* If you use ksonnet and run on GCS and Bigtable you may see an error in your config as a default value was removed.
* If you are using boltdb-shipper, you will likekly need to add a new schema_config entry.

Check the [upgrade guide](https://example.com/acme/logstore/blob/master/docs/sources/operations/upgrade.md#160) for detailed information on all these changes.

### Notable Features and Fixes

#### Query language enhancements

* [2150](https://example.com/acme/logstore/pull/2150) introduces `bytes_rate`, which calculates the per second byte rate of a log stream, and `bytes_over_time`, which returns the byte size of a log stream.
* [2182](https://example.com/acme/logstore/pull/2182) introduces a long list of comparison operators, which will let you write queries like `count_over_time({foo="bar"}[1m]) > 10`. Check out the PR for a more detailed description.

#### Logstore performance improvements

* [2216](https://example.com/acme/logstore/pull/2216), [2218](https://example.com/acme/logstore/pull/2218), and [2219](https://example.com/acme/logstore/pull/2219) all improve how memory is allocated and reused for queries.
* [2239](https://example.com/acme/logstore/pull/2239) is a huge improvement for certain cases in which a query covers a large number of streams that all overlap in time. Overlapping data is now internally cached while Logstore works to sort all the streams into the proper time order.
* [2293](https://example.com/acme/logstore/pull/2293) was a big refactor to how Logstore internally processes log queries vs. metric queries, creating separate code paths to further optimize metric queries. Metric query performance is now 2 to 10 times faster.

If you are using the query-frontend:

* [2441](https://example.com/acme/logstore/pull/2441) improves how label queries can be split and queried in parallel
* [2123](https://example.com/acme/logstore/pull/2123) allows queries to the `series` API to be split by time and parallelized; and last but most significant
* [1927](https://example.com/acme/logstore/pull/1927) allows for a much larger range of queries to be sharded and performed in parallel. Query sharding is a topic in itself, but as a rough summary, this type of sharding is not time dependent and leverages how data is already stored by Logstore to be able to split queries up into 16 separate pieces to be queried at the same time.

#### Promtail

* [2296](https://example.com/acme/logstore/pull/2296) allows Promtail to expose the Logstore Push API. With this, you can push from any client to Promtail as if it were Logstore, and Promtail can then forward those logs to another Promtail or to Logstore. There are some good use cases for this with the Logstore Docker Logging Driver; if you want an easier way to configure pipelines or expose metrics collection, point your Docker drivers at a Promtail instance.
* [2282](https://example.com/acme/logstore/pull/2282) contains an example Amazon Lambda where you can use a fan-in approach and ingestion timestamping in Promtail to work around `out of order` issues with multiple Lambdas processing the same log stream. This is one way to get logs from a high-cardinality source without adding a high-cardinality label.
* [2060](https://example.com/acme/logstore/pull/2060) introduces the `Replace` stage, which lets you find and replace or remove text inside a log line. Combined with [2422](https://example.com/acme/logstore/pull/2422) and [2480](https://example.com/acme/logstore/pull/2480), you can now find and replace sensitive data in a log line like a password or email address and replace it with ****, or hash the value to prevent readability, while still being able to trace the value through your logs. Last on the list of pipeline additions,
* [2496](https://example.com/acme/logstore/pull/2496) adds a `Drop` pipeline stage, which lets you drop log lines based on several criteria options including regex matching content, line length, or the age of the log line. The last two are useful to prevent sending to Logstore logs that you know would be rejected based on configured limits in the Logstore server.

#### Logstash output plugin

* [1822](https://example.com/acme/logstore/pull/1822) added a Logstash output plugin for Logstore. If you have an existing Logstash install, you can now use this plugin to send your logs to Logstore to make it easier to try out, or use Logstore alongside an existing logging installation.

#### Logstore Canary

* [2344](https://example.com/acme/logstore/pull/2344) improved the canaries capabilities for checking for data integrity, including spot checking for logs over a longer time window and running metric queries to verify count_over_time accuracy.

#### Logcli

* [2470](https://example.com/acme/logstore/pull/2470) allows you to color code your log lines based on their stream labels for a nice visual indicator of streams.
* [2497](https://example.com/acme/logstore/pull/2497) expands on the series API query to Logstore with the`--analyze-labels` flag, which can show you a detailed breakdown of your label key and value combinations. This is very useful for finding improper label usage in Logstore or labels with high cardinality.
* [2482](https://example.com/acme/logstore/pull/2482), in which LogCLI will automatically batch requests to Logstore to allow making queries with a `--limit=` far larger than the server side limit defined in Logstore. LogCLI will dispatch the request in a series of queries configured by the `--batch=` parameter (which defaults to 1000) until the requested limit is reached!

#### Misc

* [2453](https://example.com/acme/logstore/pull/2453) improves the error messages when a query times out, as `Context Deadline Exceeded` wasn’t the most intuitive.
* [2336](https://example.com/acme/logstore/pull/2336) provides two new flags that will print the entire Logstore config object at startup. Be warned there are a lot of config options, and many won’t apply to your setup (such as storage configs you aren’t using), but this can be a really useful tool when troubleshooting. Sticking with the theme of best for last,
* [2224](https://example.com/acme/logstore/pull/2224) and [2288](https://example.com/acme/logstore/pull/2288) improve support for running Logstore with a shared Ring using memberlist while not requiring Consul or Etcd. We need to follow up soon with some better documentation or a blog post on this!

### Dependencies

* Go Version:     1.14.2
* Corestore Version: 7014ff11ed70d9d59ad29d0a95e73999c436c47c

### All Changes

#### Logstore

* [2484](https://example.com/acme/logstore/pull/2484) **slim-bean**: Logstore: fix batch iterator error when all chunks overlap and chunk time ranges are greater than query time range
* [2483](https://example.com/acme/logstore/pull/2483) **sandeepsukhani**: download boltdb files parallelly during reads
* [2472](https://example.com/acme/logstore/pull/2472) **owen-d**: series endpoint uses normal splits
* [2466](https://example.com/acme/logstore/pull/2466) **owen-d**: BatchIter edge cases
* [2463](https://example.com/acme/logstore/pull/2463) **sandeepsukhani**: revendor corestore to latest master
* [2457](https://example.com/acme/logstore/pull/2457) **adityacs**: Fix panic in cassandra storage while registering metrics
* [2453](https://example.com/acme/logstore/pull/2453) **slim-bean**: Logstore: Improve error messages on query timeout or cancel
* [2450](https://example.com/acme/logstore/pull/2450) **adityacs**: Fixes panic in runtime_config
* [2449](https://example.com/acme/logstore/pull/2449) **jvrplmlmn**: Replace usage of sync/atomic with uber-go/atomic
* [2441](https://example.com/acme/logstore/pull/2441) **cyriltovena**: Split label names queries in the frontend.
* [2427](https://example.com/acme/logstore/pull/2427) **owen-d**: Revendor corestore
* [2392](https://example.com/acme/logstore/pull/2392) **owen-d**: avoid mutating config while parsing -config.file
* [2346](https://example.com/acme/logstore/pull/2346) **cyriltovena**: Fixes LogQL grouping
* [2336](https://example.com/acme/logstore/pull/2336) **slim-bean**: Logstore: add -print-config-stderr flag to dump logstore's runtime config to stderr
* [2330](https://example.com/acme/logstore/pull/2330) **slim-bean**: Logstore: Use a new context to update the ring state after a failed chunk transfer
* [2328](https://example.com/acme/logstore/pull/2328) **slim-bean**: Logstore: Transfer one chunk at a time per series during chunk transfers
* [2327](https://example.com/acme/logstore/pull/2327) **adityacs**: Fix data race in ingester
* [2323](https://example.com/acme/logstore/pull/2323) **cyriltovena**: Improve object key parsing for boltdb shipper.
* [2306](https://example.com/acme/logstore/pull/2306) **cyriltovena**: Fixes buffered iterator skipping very long lines.
* [2302](https://example.com/acme/logstore/pull/2302) **cyriltovena**: Improve entry deduplication.
* [2294](https://example.com/acme/logstore/pull/2294) **cyriltovena**: Remove NET_BIND_SERVICE capability requirement.
* [2293](https://example.com/acme/logstore/pull/2293) **cyriltovena**: Improve metric queries by computing samples at the edges.
* [2288](https://example.com/acme/logstore/pull/2288) **periklis**: Add support for memberlist dns-based discovery
* [2268](https://example.com/acme/logstore/pull/2268) **owen-d**: lock fix for flaky test
* [2266](https://example.com/acme/logstore/pull/2266) **cyriltovena**: Update to latest corestore.
* [2264](https://example.com/acme/logstore/pull/2264) **adityacs**: Fix ingester results for series query
* [2261](https://example.com/acme/logstore/pull/2261) **sandeepsukhani**: create smaller unique files from boltdb shipper and other code refactorings
* [2254](https://example.com/acme/logstore/pull/2254) **slim-bean**: Logstore: Series API will return all series with no match or empty matcher
* [2252](https://example.com/acme/logstore/pull/2252) **owen-d**: avoids further time splitting in querysharding mware
* [2250](https://example.com/acme/logstore/pull/2250) **slim-bean**: Logstore: Remove redundant log warning
* [2249](https://example.com/acme/logstore/pull/2249) **owen-d**: avoids recording stats in the sharded engine
* [2248](https://example.com/acme/logstore/pull/2248) **cyriltovena**: Add performance profile flags for logcli.
* [2239](https://example.com/acme/logstore/pull/2239) **cyriltovena**: Cache overlapping blocks
* [2224](https://example.com/acme/logstore/pull/2224) **periklis**: Replace memberlist service in favor of corestore provided service
* [2223](https://example.com/acme/logstore/pull/2223) **adityacs**: Add Error method for step evaluators
* [2219](https://example.com/acme/logstore/pull/2219) **cyriltovena**: Reuse slice for the range vector allocations.
* [2218](https://example.com/acme/logstore/pull/2218) **cyriltovena**: Reuse buffer for hash computation in the engine.
* [2216](https://example.com/acme/logstore/pull/2216) **cyriltovena**: Improve point allocations for each steps in the logql engine.
* [2211](https://example.com/acme/logstore/pull/2211) **sandeepsukhani**: query tee proxy with support for comparison of responses
* [2206](https://example.com/acme/logstore/pull/2206) **sandeepsukhani**: disable index dedupe when rf > 1 and current or upcoming index type is boltdb-shipper
* [2204](https://example.com/acme/logstore/pull/2204) **owen-d**: bumps corestore & fixes conflicts
* [2191](https://example.com/acme/logstore/pull/2191) **periklis**: Add flag to disable tracing activation
* [2189](https://example.com/acme/logstore/pull/2189) **owen-d**: Fix vector-scalar comparisons
* [2182](https://example.com/acme/logstore/pull/2182) **owen-d**: Logql comparison ops
* [2178](https://example.com/acme/logstore/pull/2178) **cyriltovena**: Fixes path prefix in the querier.
* [2166](https://example.com/acme/logstore/pull/2166) **sandeepsukhani**: enforce requirment for periodic config for index tables to be 24h when using boltdb shipper
* [2161](https://example.com/acme/logstore/pull/2161) **cyriltovena**: Fix error message for max tail connections.
* [2156](https://example.com/acme/logstore/pull/2156) **sandeepsukhani**: boltdb shipper download failure handling and some refactorings
* [2150](https://example.com/acme/logstore/pull/2150) **cyriltovena**: Bytes aggregations
* [2136](https://example.com/acme/logstore/pull/2136) **cyriltovena**: Fixes Iterator boundaries
* [2123](https://example.com/acme/logstore/pull/2123) **adityacs**: Fix Series API slowness
* [1927](https://example.com/acme/logstore/pull/1927) **owen-d**: Feature/querysharding ii
* [2032](https://example.com/acme/logstore/pull/2032) **tivvit**: Added support for tail to query frontend

#### Promtail

* [2496](https://example.com/acme/logstore/pull/2496) **slim-bean**: Promtail: Drop stage
* [2475](https://example.com/acme/logstore/pull/2475) **slim-bean**: Promtail: force the log level on any Logstore Push API target servers to match Promtail's log level.
* [2474](https://example.com/acme/logstore/pull/2474) **slim-bean**: Promtail: use --client.external-labels for all clients
* [2471](https://example.com/acme/logstore/pull/2471) **owen-d**: Fix/promtail yaml config
* [2464](https://example.com/acme/logstore/pull/2464) **slim-bean**: Promtail: Bug: logstore push api, clone labels before handling
* [2438](https://example.com/acme/logstore/pull/2438) **rfratto**: pkg/promtail: propagate a logger rather than using util.Logger globally
* [2432](https://example.com/acme/logstore/pull/2432) **pyr0hu**: Promtail: Allow empty replace values for replace stage
* [2422](https://example.com/acme/logstore/pull/2422) **wardbekker**: Template: Added a sha256 template function for obfuscating / anonymize PII data in e.g. the replace stage
* [2414](https://example.com/acme/logstore/pull/2414) **rfratto**: Add RegisterFlagsWithPrefix to config structs
* [2386](https://example.com/acme/logstore/pull/2386) **cyriltovena**: Add regex function to promtail template stage.
* [2345](https://example.com/acme/logstore/pull/2345) **adityacs**: Refactor Promtail target manager code
* [2301](https://example.com/acme/logstore/pull/2301) **flixr**: Promtail: support unix timestamps with fractional seconds
* [2296](https://example.com/acme/logstore/pull/2296) **slim-bean**: Promtail: Logstore Push API
* [2282](https://example.com/acme/logstore/pull/2282) **owen-d**: Lambda-Promtail
* [2242](https://example.com/acme/logstore/pull/2242) **carlpett**: Set user agent on outgoing http requests
* [2196](https://example.com/acme/logstore/pull/2196) **cyriltovena**: Adds default -config.file for the promtail docker images.
* [2127](https://example.com/acme/logstore/pull/2127) **bastjan**: Update go-syslog to accept non-UTF8 encoding in syslog message
* [2111](https://example.com/acme/logstore/pull/2111) **adityacs**: Fix Promtail journal seeking known position
* [2105](https://example.com/acme/logstore/pull/2105) **fatpat**: promtail: Add Entry variable to template
* [1118](https://example.com/acme/logstore/pull/1118) **shuttie**: promtail: fix high CPU usage on large kubernetes clusters.
* [2060](https://example.com/acme/logstore/pull/2060) **adityacs**: Feature: Replace stage in pipeline
* [2087](https://example.com/acme/logstore/pull/2087) **adityacs**: Set JournalTarget Priority value to keyword

#### Logcli

* [2497](https://example.com/acme/logstore/pull/2497) **slim-bean**: logcli: adds --analyize-labels to logcli series command and changes how labels are provided to the command
* [2482](https://example.com/acme/logstore/pull/2482) **slim-bean**: Logcli: automatically batch requests
* [2470](https://example.com/acme/logstore/pull/2470) **adityacs**: colored labels output for logcli
* [2235](https://example.com/acme/logstore/pull/2235) **pstibrany**: logcli: Remove single newline from the raw line before printing.
* [2126](https://example.com/acme/logstore/pull/2126) **cyriltovena**: Validate local storage config for the logcli
* [2083](https://example.com/acme/logstore/pull/2083) **adityacs**: Support querying labels on time range in logcli

#### Docs

* [2473](https://example.com/acme/logstore/pull/2473) **owen-d**: fixes lambda-promtail relative doc link
* [2454](https://example.com/acme/logstore/pull/2454) **oddlittlebird**: Create CODEOWNERS
* [2439](https://example.com/acme/logstore/pull/2439) **till**: Docs: updated "Upgrading" for docker driver
* [2437](https://example.com/acme/logstore/pull/2437) **wardbekker**: DOCS: clarified globbing behaviour of **path** of the doublestar library
* [2431](https://example.com/acme/logstore/pull/2431) **endu**: fix dead link
* [2425](https://example.com/acme/logstore/pull/2425) **RichiH**: Change conduct contact email address
* [2420](https://example.com/acme/logstore/pull/2420) **petuhovskiy**: Fix docker driver doc
* [2418](https://example.com/acme/logstore/pull/2418) **cyriltovena**: Add logstash to clients page with FrontMatter.
* [2402](https://example.com/acme/logstore/pull/2402) **cyriltovena**: More fixes for the website.
* [2400](https://example.com/acme/logstore/pull/2400) **tontongg**: Fix URL to LogQL documentation
* [2398](https://example.com/acme/logstore/pull/2398) **robbymilo**: Docs - update links, readme
* [2397](https://example.com/acme/logstore/pull/2397) **coderanger**: 📝 Note that entry_parser is deprecated.
* [2396](https://example.com/acme/logstore/pull/2396) **dnsmichi**: Docs: Fix Fluentd title (visible in menu)
* [2391](https://example.com/acme/logstore/pull/2391) **cyriltovena**: Update fluentd docs and fixes links for the website.
* [2390](https://example.com/acme/logstore/pull/2390) **cyriltovena**: Fluent bit docs
* [2389](https://example.com/acme/logstore/pull/2389) **cyriltovena**: Docker driver doc
* [2385](https://example.com/acme/logstore/pull/2385) **abowloflrf**: Update logo link in README.md
* [2378](https://example.com/acme/logstore/pull/2378) **robbymilo**: Sync docs to website
* [2360](https://example.com/acme/logstore/pull/2360) **owen-d**: Makes timestamp parsing docs clearer
* [2358](https://example.com/acme/logstore/pull/2358) **rille111**: Documentation: Add example for having separate pvc for logstore, using helm
* [2357](https://example.com/acme/logstore/pull/2357) **owen-d**: Storage backend examples
* [2338](https://example.com/acme/logstore/pull/2338) **cyriltovena**: Add a complete tutorial on how to ship logs from AWS EKS.
* [2335](https://example.com/acme/logstore/pull/2335) **cyriltovena**: Improve documentation of the metric stage.
* [2331](https://example.com/acme/logstore/pull/2331) **cyriltovena**: Add a tutorial to forward AWS ECS logs to Logstore.
* [2321](https://example.com/acme/logstore/pull/2321) **cyriltovena**: Tutorial to run Promtail on AWS EC2
* [2318](https://example.com/acme/logstore/pull/2318) **adityacs**: Configuration documentation improvements
* [2317](https://example.com/acme/logstore/pull/2317) **owen-d**: remove DynamoDB chunk store doc
* [2308](https://example.com/acme/logstore/pull/2308) **wardbekker**: Added a link to the replace parsing stage
* [2305](https://example.com/acme/logstore/pull/2305) **rafaelpissolatto**: Fix schema_config store value
* [2285](https://example.com/acme/logstore/pull/2285) **adityacs**: Fix local.md doc
* [2284](https://example.com/acme/logstore/pull/2284) **owen-d**: Update local.md
* [2279](https://example.com/acme/logstore/pull/2279) **Fra-nk**: Documentation: Refine LogQL documentation
* [2273](https://example.com/acme/logstore/pull/2273) **RichiH**: Fix typo
* [2247](https://example.com/acme/logstore/pull/2247) **carlpett**: docs: Fix missing quotes
* [2233](https://example.com/acme/logstore/pull/2233) **vyzigold**: docs: Add readmes to individual helm charts
* [2220](https://example.com/acme/logstore/pull/2220) **oddlittlebird**: Docs: Local install edits
* [2217](https://example.com/acme/logstore/pull/2217) **fredr**: docs: BoltDB typo
* [2215](https://example.com/acme/logstore/pull/2215) **fredr**: docs: Correct logstore address for docker-compose
* [2172](https://example.com/acme/logstore/pull/2172) **cyriltovena**: Update old link for pipeline stages.
* [2163](https://example.com/acme/logstore/pull/2163) **slim-bean**: docs: fix an error in the example log line and byte counter metrics
* [2160](https://example.com/acme/logstore/pull/2160) **slim-bean**: Fix some errors in the upgrade guide to 1.5.0 and add some missing notes discovered by users.
* [2152](https://example.com/acme/logstore/pull/2152) **eamonryan**: Fix typo in promtail ClusterRole
* [2139](https://example.com/acme/logstore/pull/2139) **adityacs**: Fix configuration docs
* [2137](https://example.com/acme/logstore/pull/2137) **RichiH**: Propose new governance
* [2136](https://example.com/acme/logstore/pull/2136) **cyriltovena**: Fixes Iterator boundaries
* [2125](https://example.com/acme/logstore/pull/2125) **theMercedes**: Update logql.md
* [2112](https://example.com/acme/logstore/pull/2112) **nileshcs**: Documentation: Outdated fluentd image name, UID details, link update
* [2092](https://example.com/acme/logstore/pull/2092) **i-takizawa**: docs: make <placeholders> visible

#### Build

* [2467](https://example.com/acme/logstore/pull/2467) **slim-bean**: Update Logstore build image

#### Ksonnet

* [2460](https://example.com/acme/logstore/pull/2460) **Duologic**: refactor: use $.core.v1.envVar
* [2452](https://example.com/acme/logstore/pull/2452) **slim-bean**: ksonnet: Reduce querier parallelism to a more sane default value and remove the default setting for storage_backend
* [2377](https://example.com/acme/logstore/pull/2377) **Duologic**: refactor: moved jaeger-agent-mixin
* [2373](https://example.com/acme/logstore/pull/2373) **slim-bean**: Ksonnet: Add a Pod Disruption Budget to Logstore Ingesters
* [2185](https://example.com/acme/logstore/pull/2185) **cyriltovena**: Refactor mixin routes and add series API.
* [2162](https://example.com/acme/logstore/pull/2162) **slim-bean**: ksonnet: Fix up datasources and variables in Logstore Operational
* [2091](https://example.com/acme/logstore/pull/2091) **beorn7**: Keep scrape config in line with the new Prometheus scrape config

#### Docker logging driver

* [2435](https://example.com/acme/logstore/pull/2435) **cyriltovena**: Add more precisions on the docker driver installed on the daemon.
* [2343](https://example.com/acme/logstore/pull/2343) **jdfalk**: logstore-docker-driver: Change "ignoring empty line" to debug logging
* [2295](https://example.com/acme/logstore/pull/2295) **cyriltovena**: Remove mount in the docker driver.
* [2199](https://example.com/acme/logstore/pull/2199) **cyriltovena**: Docker driver relabeling
* [2116](https://example.com/acme/logstore/pull/2116) **cyriltovena**: Allows to change the log driver mode and buffer size.

#### Logstash output plugin

* [2415](https://example.com/acme/logstore/pull/2415) **cyriltovena**: Set service values via --set for logstash.
* [2410](https://example.com/acme/logstore/pull/2410) **adityacs**: logstash code refactor and doc improvements
* [1822](https://example.com/acme/logstore/pull/1822) **adityacs**: Logstore Logstash Plugin

#### Logstore canary

* [2413](https://example.com/acme/logstore/pull/2413) **slim-bean**: Logstore-Canary: Backoff retries on query failures, add histograms for query performance.
* [2369](https://example.com/acme/logstore/pull/2369) **slim-bean**: Logstore Canary: One more round of improvements to query for missing websocket entries up to max-wait
* [2350](https://example.com/acme/logstore/pull/2350) **slim-bean**: Canary tweaks
* [2344](https://example.com/acme/logstore/pull/2344) **slim-bean**: Logstore-Canary: Add query spot checking and metric count checking
* [2259](https://example.com/acme/logstore/pull/2259) **ombre8**: Canary: make stream configurable

#### Fluentd

* [2407](https://example.com/acme/logstore/pull/2407) **cyriltovena**: bump fluentd version to release a new gem.
* [2399](https://example.com/acme/logstore/pull/2399) **tarokkk**: fluentd: Make fluentd version requirements permissive
* [2179](https://example.com/acme/logstore/pull/2179) **takanabe**: Improve fluentd plugin development experience
* [2171](https://example.com/acme/logstore/pull/2171) **takanabe**: Add server TLS certificate verification

#### Fluent Bit

* [2375](https://example.com/acme/logstore/pull/2375) **cyriltovena**: Fixes the fluentbit batchwait  backward compatiblity.
* [2367](https://example.com/acme/logstore/pull/2367) **dojci**: fluent-bit: Add more logstore client configuration options
* [2365](https://example.com/acme/logstore/pull/2365) **dojci**: fluent-bit: Fix fluent-bit exit callback when buffering is enabled
* [2290](https://example.com/acme/logstore/pull/2290) **cyriltovena**: Fixes the lint issue merged to master.
* [2286](https://example.com/acme/logstore/pull/2286) **adityacs**: Fix fluent-bit newline and tab characters
* [2142](https://example.com/acme/logstore/pull/2142) **dojci**: Add FIFO queue persistent buffering for fluent bit output plugin
* [2089](https://example.com/acme/logstore/pull/2089) **FrederikNS**: Allow configuring more options for output configuration

#### Helm

* [2406](https://example.com/acme/logstore/pull/2406) **steven-sheehy**: Helm: Fix regression in chart name
* [2379](https://example.com/acme/logstore/pull/2379) **StevenReitsma**: production/helm: Add emptyDir volume type to promtail PSP
* [2366](https://example.com/acme/logstore/pull/2366) **StevenReitsma**: production/helm: Add projected and downwardAPI volume types to PodSecurityPolicy (#2355)
* [2258](https://example.com/acme/logstore/pull/2258) **Synehan**: helm: add annotations to service monitor
* [2241](https://example.com/acme/logstore/pull/2241) **chauffer**: Kubernetes manifests: Remove namespace from cluster-wide resources
* [2238](https://example.com/acme/logstore/pull/2238) **vhrosales**: helm: Add loadBalancerIP option to logstore chart
* [2205](https://example.com/acme/logstore/pull/2205) **joschi36**: BUG: add missing namespace in ingress object
* [2197](https://example.com/acme/logstore/pull/2197) **cyriltovena**: Render logstore datasources even if Acme is disabled.
* [2141](https://example.com/acme/logstore/pull/2141) **cyriltovena**: Adds the ability to have a pull secrets for Promtail.
* [2099](https://example.com/acme/logstore/pull/2099) **allout58**: helm/logstore-stack: Support Prometheus on a sub-path in Acme config
* [2086](https://example.com/acme/logstore/pull/2086) **osela**: helm/logstore-stack: render logstore datasource only if acme is enabled
* [2091](https://example.com/acme/logstore/pull/2091) **beorn7**: Keep scrape config in line with the new Prometheus scrape config

#### Build

* [2371](https://example.com/acme/logstore/pull/2371) **cyriltovena**: Fixes helm publish that needs now to add repo.
* [2341](https://example.com/acme/logstore/pull/2341) **slim-bean**: Build: Fix CI helm test
* [2309](https://example.com/acme/logstore/pull/2309) **cyriltovena**: Test again arm32 on internal ci.
* [2307](https://example.com/acme/logstore/pull/2307) **cyriltovena**: Removes arm32 for now as we're migrating the CI.
* [2287](https://example.com/acme/logstore/pull/2287) **wardbekker**: Change the Acme image to latest
* [2212](https://example.com/acme/logstore/pull/2212) **roidelapluie**: Remove unhelpful/problematic term in circleci.yml

## 1.5.0 (2020-05-20)

It's been a busy month and a half since 1.4.0 was released, and a lot of new improvements have been added to Logstore since!

Be prepared for some configuration changes that may cause some bumps when upgrading,
we apologize for this but are always striving to reach the right compromise of code simplicity and user/operating experience.

In this case we opted to keep a simplified configuration inline with Corestore rather than a more complicated and error prone internal config mapping or difficult to implement support for multiple config names for the same feature.

This does result in breaking config changes for some configurations, however, these will fail fast and with the [list of diffs](https://corestoremetrics.io/docs/changelog/#config-file-breaking-changes) from the Corestore project should be quick to fix.

### Important Notes

**Be prepared for breaking config changes.**  Logstore 1.5.0 vendors corestore [v1.0.1-0.20200430170006-3462eb63f324](https://github.com/corestoreproject/corestore/commit/3462eb63f324c649bbaa122933bc591b710f4e48),
there were substantial breaking config changes in Corestore 1.0 which standardized config options, and fixed typos.

**The Logstore docker image user has changed to no longer be root**

Check the [upgrade guide](https://example.com/acme/logstore/blob/master/docs/sources/operations/upgrade.md#150) for more detailed information on these changes.

### Notable Features and Fixes

There are quite a few we want to mention listed in order they were merged (mostly)

* [1837](https://example.com/acme/logstore/pull/1837) **sandeepsukhani**: flush boltdb to object store

This is perhaps the most exciting feature of 1.5.0, the first steps in removing a dependency on a separate index store!  This feature is still very new and experimental, however, we want this to be the future for Logstore.  Only requiring just an object store.

If you want to test this new feature, and help us find any bugs, check out the [docs](docs/operations/storage/boltdb-shipper.md) to learn more and get started.

* [2073](https://example.com/acme/logstore/pull/2073) **slim-bean**: Logstore: Allow configuring query_store_max_look_back_period when running a filesystem store and boltdb-shipper

This is even more experimental than the previous feature mentioned however also pretty exciting for Logstore users who use the filesystem storage. We can leverage changes made in [1837](https://example.com/acme/logstore/pull/1837) to now allow Logstore to run in a clustered mode with individual filesystem stores!

Please check out the last section in the [filesystem docs](docs/operations/storage/filesystem.md) for more details on how this works and how to use it!

* [2095](https://example.com/acme/logstore/pull/2095) **cyriltovena**: Adds backtick for the quoted string token lexer.

This will come as a big win to anyone who is writing complicated reqular expressions in either their Label matchers or Filter Expressions.  Starting now you can use the backtick to encapsulate your regex **and not have to do any escaping of special characters!!**

Examples:

```
{name="cassandra"} |~ `error=\w+`
{name!~`mysql-\d+`}
```

* [2055](https://example.com/acme/logstore/pull/2055) **aknuds1**: Chore: Fix spelling of per second in code

This is technically a breaking change for anyone who wrote code to processes the new statistics output in the query result added in 1.4.0, we apologize to anyone in this situation but if we don't fix this kind of error now it will be there forever.
And at the same time we didn't feel it was appropriate to make any major api revision changes for such a new feature and simple change.  We are always trying to use our best judgement in cases like this.

* [2031](https://example.com/acme/logstore/pull/2031) **cyriltovena**: Improve protobuf serialization

Thanks @cyriltovena for another big performance improvement in Logstore, this time around protbuf's!

* [2021](https://example.com/acme/logstore/pull/2021) **slim-bean**: Logstore: refactor validation and improve error messages
* [2012](https://example.com/acme/logstore/pull/2012) **slim-bean**: Logstore: Improve logging and add metrics to streams dropped by stream limit

These two changes standardize the metrics used to report when a tenant hits a limit, now all discarded samples should be reported under `logstore_discarded_samples_total` and you no longer need to also reference `corestore_discarded_samples_total`.
Additionally error messages were improved to help clients take better action when hitting limits.

* [1970](https://example.com/acme/logstore/pull/1970) **cyriltovena**: Allow to aggregate binary operations.

Another nice improvement to the query language which allows queries like this to work now:

```
sum by (job) (count_over_time({namespace="tns"}[5m] |= "level=error") / count_over_time({namespace="tns"}[5m]))
```

* [1713](https://example.com/acme/logstore/pull/1713) **adityacs**: Log error message for invalid checksum

In the event something went wrong with a stored chunk, rather than fail the query we ignore the chunk and return the rest.

* [2066](https://example.com/acme/logstore/pull/2066) **slim-bean**: Promtail: metrics stage can also count line bytes

This is a nice extension to a previous feature which let you add a metric to count log lines per stream, you can now count log bytes per stream.

Check out [this example](docs/clients/promtail/configuration.md#counter) to configure this in your promtail pipelines.

* [1935](https://example.com/acme/logstore/pull/1935) **cyriltovena**: Support stdin target via flag instead of automatic detection.

Third times a charm!  With 1.4.0 we allowed sending logs directly to promtail via stdin, with 1.4.1 we released a patch for this feature which wasn't detecting stdin correctly on some operating systems.
Unfortunately after a few more bug reports it seems this change caused some more undesired side effects so we decided to not try to autodetect stdin at all, instead now you must pass the `--stdin` flag if you want Promtail to listen for logs on stdin.

* [2076](https://example.com/acme/logstore/pull/2076) **cyriltovena**: Allows to pass inlined pipeline stages to the docker driver.
* [1906](https://example.com/acme/logstore/pull/1906) **cyriltovena**: Add no-file and keep-file log option for docker driver.

The docker logging driver received a couple very nice updates, it's always been challenging to configure pipeline stages for the docker driver, with the first PR there are now a few easier ways to do this!
In the second PR we added config options to control keeping any log files on the host when using the docker logging driver, allowing you to run with no disk access if you would like, as well as allowing you to control keeping log files available after container restarts.

* [1864](https://example.com/acme/logstore/pull/1864) **cyriltovena**: Sign helm package with GPG.

We now GPG sign helm packages!

### All Changes

#### Logstore

* [2097](https://example.com/acme/logstore/pull/2097) **owen-d**: simplifies/updates some of our configuration examples
* [2095](https://example.com/acme/logstore/pull/2095) **cyriltovena**: Adds backtick for the quoted string token lexer.
* [2093](https://example.com/acme/logstore/pull/2093) **cyriltovena**: Fixes unit in stats request log.
* [2088](https://example.com/acme/logstore/pull/2088) **slim-bean**: Logstore: allow no encoding/compression on chunks
* [2078](https://example.com/acme/logstore/pull/2078) **owen-d**: removes yolostring
* [2073](https://example.com/acme/logstore/pull/2073) **slim-bean**: Logstore: Allow configuring query_store_max_look_back_period when running a filesystem store and boltdb-shipper
* [2064](https://example.com/acme/logstore/pull/2064) **cyriltovena**: Reverse entry iterator pool
* [2059](https://example.com/acme/logstore/pull/2059) **cyriltovena**: Recover from panic in http and grpc handlers.
* [2058](https://example.com/acme/logstore/pull/2058) **cyriltovena**: Fix a bug in range vector skipping data.
* [2055](https://example.com/acme/logstore/pull/2055) **aknuds1**: Chore: Fix spelling of per second in code
* [2046](https://example.com/acme/logstore/pull/2046) **gouthamve**: Fix bug in logql parsing that leads to crash.
* [2050](https://example.com/acme/logstore/pull/2050) **aknuds1**: Chore: Correct typo "per seconds"
* [2034](https://example.com/acme/logstore/pull/2034) **sandeepsukhani**: some metrics for measuring performance and failures in boltdb shipper
* [2031](https://example.com/acme/logstore/pull/2031) **cyriltovena**: Improve protobuf serialization
* [2030](https://example.com/acme/logstore/pull/2030) **adityacs**: Update logstore to corestore master
* [2023](https://example.com/acme/logstore/pull/2023) **cyriltovena**: Support post requests in the frontend queryrange handler.
* [2021](https://example.com/acme/logstore/pull/2021) **slim-bean**: Logstore: refactor validation and improve error messages
* [2019](https://example.com/acme/logstore/pull/2019) **slim-bean**: make `logstore_ingester_memory_streams` Gauge per tenant.
* [2012](https://example.com/acme/logstore/pull/2012) **slim-bean**: Logstore: Improve logging and add metrics to streams dropped by stream limit
* [2010](https://example.com/acme/logstore/pull/2010) **cyriltovena**: Update lz4 library to latest to ensure deterministic output.
* [2001](https://example.com/acme/logstore/pull/2001) **sandeepsukhani**: table client for boltdb shipper to enforce retention
* [1995](https://example.com/acme/logstore/pull/1995) **sandeepsukhani**: make boltdb shipper singleton and some other minor refactoring
* [1987](https://example.com/acme/logstore/pull/1987) **slim-bean**: Logstore: Add a missing method to facade which is called by the metrics storage client in corestore
* [1982](https://example.com/acme/logstore/pull/1982) **cyriltovena**: Update corestore to latest.
* [1977](https://example.com/acme/logstore/pull/1977) **cyriltovena**: Ensure trace propagation in our logs.
* [1976](https://example.com/acme/logstore/pull/1976) **slim-bean**: incorporate some better defaults into table-manager configs
* [1975](https://example.com/acme/logstore/pull/1975) **slim-bean**: Update corestore vendoring to latest master
* [1970](https://example.com/acme/logstore/pull/1970) **cyriltovena**: Allow to aggregate binary operations.
* [1965](https://example.com/acme/logstore/pull/1965) **slim-bean**: Logstore: Adds an `interval` paramater to query_range queries allowing a sampling of events to be returned based on the provided interval
* [1964](https://example.com/acme/logstore/pull/1964) **owen-d**: chunk bounds metric now records 8h range in 1h increments
* [1963](https://example.com/acme/logstore/pull/1963) **cyriltovena**: Improve the local config to work locally and inside docker.
* [1961](https://example.com/acme/logstore/pull/1961) **jpmcb**: [Bug] Workaround for broken etcd gomod import
* [1958](https://example.com/acme/logstore/pull/1958) **owen-d**: chunk lifespan histogram
* [1956](https://example.com/acme/logstore/pull/1956) **sandeepsukhani**: update corestore to latest master
* [1953](https://example.com/acme/logstore/pull/1953) **jpmcb**: Go mod: explicit golang.org/x/net replace
* [1950](https://example.com/acme/logstore/pull/1950) **cyriltovena**: Fixes case handling in regex simplification.
* [1949](https://example.com/acme/logstore/pull/1949) **SerialVelocity**: [Logstore]: Cleanup dockerfile
* [1946](https://example.com/acme/logstore/pull/1946) **slim-bean**: Logstore Update the cut block size counter when creating a memchunk from byte slice
* [1939](https://example.com/acme/logstore/pull/1939) **owen-d**: adds config validation, similar to corestore
* [1916](https://example.com/acme/logstore/pull/1916) **cyriltovena**: Add cap_net_bind_service linux capabilities to Logstore.
* [1914](https://example.com/acme/logstore/pull/1914) **owen-d**: only fetches one chunk per series in /series
* [1875](https://example.com/acme/logstore/pull/1875) **owen-d**: support `match[]` encoding
* [1869](https://example.com/acme/logstore/pull/1869) **pstibrany**: Update Corestore to latest master
* [1846](https://example.com/acme/logstore/pull/1846) **owen-d**: Sharding optimizations I: AST mapping
* [1838](https://example.com/acme/logstore/pull/1838) **cyriltovena**: Move default port for Logstore to 3100 everywhere.
* [1837](https://example.com/acme/logstore/pull/1837) **sandeepsukhani**: flush boltdb to object store
* [1834](https://example.com/acme/logstore/pull/1834) **Mario-Hofstaetter**: Logstore/Change local storage directory to /logstore/ and fix permissions (#1833)
* [1819](https://example.com/acme/logstore/pull/1819) **cyriltovena**: Adds a counter for total flushed chunks per reason.
* [1816](https://example.com/acme/logstore/pull/1816) **sdojjy**: logstore can not be started with logstore-local-config.yaml
* [1810](https://example.com/acme/logstore/pull/1810) **cyriltovena**: Optimize empty filter queries.
* [1809](https://example.com/acme/logstore/pull/1809) **cyriltovena**: Test stats memchunk
* [1804](https://example.com/acme/logstore/pull/1804) **pstibrany**: Convert Logstore modules to services
* [1799](https://example.com/acme/logstore/pull/1799) **pstibrany**: logstore: update Corestore to master
* [1798](https://example.com/acme/logstore/pull/1798) **adityacs**: Support configurable maximum of the limits parameter
* [1713](https://example.com/acme/logstore/pull/1713) **adityacs**: Log error message for invalid checksum
* [1706](https://example.com/acme/logstore/pull/1706) **cyriltovena**: Non-root user docker image for Logstore.

#### Logcli

* [2027](https://example.com/acme/logstore/pull/2027) **pstibrany**: logcli: Query needs to be stored into url.RawQuery, and not url.Path
* [2000](https://example.com/acme/logstore/pull/2000) **cyriltovena**: Improve URL building in the logcli to strip trailing /.
* [1922](https://example.com/acme/logstore/pull/1922) **bavarianbidi**: logcli: org-id/tls-skip-verify set via env var
* [1861](https://example.com/acme/logstore/pull/1861) **yeya24**: Support series API in logcli
* [1850](https://example.com/acme/logstore/pull/1850) **chrischdi**: BugFix: Fix logcli client to use OrgID in LiveTail
* [1814](https://example.com/acme/logstore/pull/1814) **cyriltovena**: Logcli remote storage.
* [1712](https://example.com/acme/logstore/pull/1712) **rfratto**: clarify logcli commands and output

#### Promtail

* [2069](https://example.com/acme/logstore/pull/2069) **slim-bean**: Promtail: log at debug level when nothing matches the specified path for a file target
* [2066](https://example.com/acme/logstore/pull/2066) **slim-bean**: Promtail: metrics stage can also count line bytes
* [2049](https://example.com/acme/logstore/pull/2049) **adityacs**: Fix promtail client default values
* [2075](https://example.com/acme/logstore/pull/2075) **cyriltovena**: Fixes a panic in dry-run when using external labels.
* [2026](https://example.com/acme/logstore/pull/2026) **adityacs**: Targets not required in promtail config
* [2004](https://example.com/acme/logstore/pull/2004) **cyriltovena**: Adds config to disable HTTP and GRPC server in Promtail.
* [1935](https://example.com/acme/logstore/pull/1935) **cyriltovena**: Support stdin target via flag instead of automatic detection.
* [1920](https://example.com/acme/logstore/pull/1920) **alexanderGalushka**: feat: tms readiness check bypass implementation
* [1894](https://example.com/acme/logstore/pull/1894) **cyriltovena**: Fixes possible panic in json pipeline stage.
* [1865](https://example.com/acme/logstore/pull/1865) **adityacs**: Fix flaky promtail test
* [1815](https://example.com/acme/logstore/pull/1815) **adityacs**: Log error message when source does not exist in extracted values
* [1627](https://example.com/acme/logstore/pull/1627) **rfratto**: Proposal: Promtail Push API

#### Docker Driver

* [2076](https://example.com/acme/logstore/pull/2076) **cyriltovena**: Allows to pass inlined pipeline stages to the docker driver.
* [2054](https://example.com/acme/logstore/pull/2054) **bkmit**: Docker driver: Allow to provision external pipeline files to plugin
* [1906](https://example.com/acme/logstore/pull/1906) **cyriltovena**: Add no-file and keep-file log option for docker driver.
* [1903](https://example.com/acme/logstore/pull/1903) **cyriltovena**: Log docker driver config map.

#### Fluentd

* [2074](https://example.com/acme/logstore/pull/2074) **osela**: fluentd plugin: support placeholders in tenant field
* [2006](https://example.com/acme/logstore/pull/2006) **Skeen**: fluent-plugin-logstore: Restructuring and CI
* [1909](https://example.com/acme/logstore/pull/1909) **jgehrcke**: fluentd logstore plugin README: add note about labels
* [1853](https://example.com/acme/logstore/pull/1853) **wardbekker**: bump gem version
* [1811](https://example.com/acme/logstore/pull/1811) **JamesJJ**: Error handling: Show data stream at "debug" level, not "warn"

#### Fluent Bit

* [2040](https://example.com/acme/logstore/pull/2040) **avii-ridge**: Add extraOutputs variable to support multiple outputs for fluent-bit
* [1915](https://example.com/acme/logstore/pull/1915) **DirtyCajunRice**: Fix fluent-bit metrics
* [1890](https://example.com/acme/logstore/pull/1890) **dottedmag**: fluentbit: JSON encoding: avoid base64 encoding of []byte inside other slices
* [1791](https://example.com/acme/logstore/pull/1791) **cyriltovena**: Improve fluentbit logfmt.

#### Ksonnet

* [1980](https://example.com/acme/logstore/pull/1980) **cyriltovena**: Log slow query from the frontend by default in ksonnet.

##### Mixins

* [2080](https://example.com/acme/logstore/pull/2080) **beorn7**: mixin: Accept suffixes to pod name in instance labels
* [2044](https://example.com/acme/logstore/pull/2044) **slim-bean**: Dashboards: fixes the cpu usage graphs
* [2043](https://example.com/acme/logstore/pull/2043) **joe-elliott**: Swapped to container restarts over terminated reasons
* [2041](https://example.com/acme/logstore/pull/2041) **slim-bean**: Dashboard: Logstore Operational improvements
* [1934](https://example.com/acme/logstore/pull/1934) **tomwilkie**: Put logstore-mixin and promtail-mixin dashboards in a folder.
* [1913](https://example.com/acme/logstore/pull/1913) **tomwilkie**: s/dashboards/acmeDashboards.

#### Helm

* [2038](https://example.com/acme/logstore/pull/2038) **oke-py**: Docs: update Logstore Helm Chart document to support Helm 3
* [2015](https://example.com/acme/logstore/pull/2015) **etashsingh**: Change image tag from 1.4.1 to 1.4.0 in Helm chart
* [1981](https://example.com/acme/logstore/pull/1981) **sshah90**: added extraCommandlineArgs in values file
* [1967](https://example.com/acme/logstore/pull/1967) **rdxmb**: helm chart: add missing line feed
* [1898](https://example.com/acme/logstore/pull/1898) **stefanandres**: [helm logstore/promtail] make UpdateStrategy configurable
* [1871](https://example.com/acme/logstore/pull/1871) **stefanandres**: [helm logstore/promtail] Add systemd-journald example with extraMount, extraVolumeMount
* [1864](https://example.com/acme/logstore/pull/1864) **cyriltovena**: Sign helm package with GPG.
* [1825](https://example.com/acme/logstore/pull/1825) **polar3130**: Helm/logstore-stack: refresh default acme.image.tag to 6.7.0
* [1817](https://example.com/acme/logstore/pull/1817) **bclermont**: Helm chart: Prevent prometheus to scrape both services

#### Logstore Canary

* [1891](https://example.com/acme/logstore/pull/1891) **joe-elliott**: Addition of a `/suspend` endpoint to Logstore Canary

#### Docs

* [2056](https://example.com/acme/logstore/pull/2056) **cyriltovena**: Update api.md
* [2014](https://example.com/acme/logstore/pull/2014) **jsoref**: Spelling
* [1999](https://example.com/acme/logstore/pull/1999) **oddlittlebird**: Docs: Added labels content
* [1974](https://example.com/acme/logstore/pull/1974) **rfratto**: fix stores for chunk and index in documentation for period_config
* [1966](https://example.com/acme/logstore/pull/1966) **oddlittlebird**: Docs: Update docker.md
* [1951](https://example.com/acme/logstore/pull/1951) **cstyan**: Move build from source instructions to root readme.
* [1945](https://example.com/acme/logstore/pull/1945) **FlorianLudwig**: docs: version pin the docker image in docker-compose
* [1925](https://example.com/acme/logstore/pull/1925) **wardbekker**: Clarified that the api push path needs to be specified.
* [1905](https://example.com/acme/logstore/pull/1905) **sshah90**: updating typo for end time parameter in api docs
* [1888](https://example.com/acme/logstore/pull/1888) **slim-bean**: docs: cleaning up the comments for the cache_config, default_validity option
* [1887](https://example.com/acme/logstore/pull/1887) **slim-bean**: docs: Adding a config change in release 1.4 upgrade doc, updating readme with new doc links
* [1881](https://example.com/acme/logstore/pull/1881) **cyriltovena**: Add precision about the range notation for LogQL.
* [1879](https://example.com/acme/logstore/pull/1879) **slim-bean**: docs: update promtail docs for backoff
* [1873](https://example.com/acme/logstore/pull/1873) **owen-d**: documents frontend worker
* [1870](https://example.com/acme/logstore/pull/1870) **ushuz**: Docs: Keep plugin install command example in one line
* [1856](https://example.com/acme/logstore/pull/1856) **slim-bean**: docs: tweak the doc section of the readme a little
* [1852](https://example.com/acme/logstore/pull/1852) **slim-bean**: docs: clean up schema recommendations
* [1843](https://example.com/acme/logstore/pull/1843) **vishesh92**: Docs: Update configuration docs for redis

#### Build

* [2042](https://example.com/acme/logstore/pull/2042) **rfratto**: Fix drone
* [2009](https://example.com/acme/logstore/pull/2009) **cyriltovena**: Adds :delegated flags to speed up build experience on MacOS.
* [1942](https://example.com/acme/logstore/pull/1942) **owen-d**: delete tag script filters by prefix instead of substring
* [1918](https://example.com/acme/logstore/pull/1918) **slim-bean**: build: This Dockerfile is a remnant from a long time ago, not needed.
* [1911](https://example.com/acme/logstore/pull/1911) **slim-bean**: build: push images for `k` branches
* [1849](https://example.com/acme/logstore/pull/1849) **cyriltovena**: Pin helm version in circle-ci helm testing workflow.

## 1.4.1 (2020-04-06)

We realized after the release last week that piping data into promtail was not working on Linux or Windows, this should fix this issue for both platforms:

* [1893](https://example.com/acme/logstore/pull/1893) **cyriltovena**: Removes file size check for pipe, not provided by linux.

Also thanks to @dottedmag for providing this fix for Fluent Bit!

* [1890](https://example.com/acme/logstore/pull/1890) **dottedmag**: fluentbit: JSON encoding: avoid base64 encoding of []byte inside other slices

## 1.4.0 (2020-04-01)

Over 130 PR's merged for this release, from 40 different contributors!!  We continue to be humbled and thankful for the growing community of contributors and users of Logstore.  Thank you all so much.

### Important Notes

**Really, this is important**

Before we get into new features, version 1.4.0 brings with it the first (that we are aware of) upgrade dependency.

We have created a dedicated page for upgrading Logstore in the [operations section of the docs](https://example.com/acme/logstore/blob/master/docs/sources/operations/upgrade.md#140)

The docker image tag naming was changed, the starting in 1.4.0 docker images no longer have the `v` prefix: `acme/logstore:1.4.0`

Also you should be aware we are now pruning old `master-xxxxx` docker images from docker hub, currently anything older than 90 days is removed.  **We will never remove released versions of Logstore**

### Notable Features

* [1661](https://example.com/acme/logstore/pull/1661) **cyriltovena**: Frontend & Querier query statistics instrumentation.

The API now returns a plethora of stats into the work Logstore performed to execute your query, eventually this will be displayed in some form in Acme to help users better understand how "expensive" their queries are.  Our goal here initially was to better instrument the recent work done in v1.3.0 on query parallelization and to better understand the performance of each part of Logstore.  In the future we are looking at additional ideas to provide feedback to users to tailor their queries for better performance.

* [1652](https://example.com/acme/logstore/pull/1652) **cyriltovena**: --dry-run Promtail.
* [1649](https://example.com/acme/logstore/pull/1649) **cyriltovena**: Pipe data to Promtail

This is a long overdue addition to Promtail which can help setup and debug pipelines, with these new features you can do this to feed a single log line into Promtail:

```bash
echo -n 'level=debug msg="test log (200)"' | cmd/promtail/promtail -config.file=cmd/promtail/promtail-local-config.yaml --dry-run -log.level=debug 2>&1 | sed 's/^.*stage/stage/g'
```

`-log.level=debug 2>&1 | sed 's/^.*stage/stage/g` are added to enable debug output, direct the output to stdout, and a sed filter to remove some noise from the log lines.

The `stdin` functionality also works without `--dry-run` allowing you to feed any logs into Promtail via `stdin` and send them to Logstore

* [1677](https://example.com/acme/logstore/pull/1677) **owen-d**: Literal Expressions in LogQL
* [1662](https://example.com/acme/logstore/pull/1662) **owen-d**: Binary operators in LogQL

These two extensions to LogQL now let you execute queries like this:

    * `sum(rate({app="foo"}[5m])) * 2`
    * `sum(rate({app="foo"}[5m]))/1e6`

* [1678](https://example.com/acme/logstore/pull/1678) **slim-bean**: promtail: metrics pipeline count all log lines

Now you can get per-stream line counts as a metric from promtail, useful for seeing which applications log the most

```yaml
- metrics:
    line_count_total:
      config:
        action: inc
        match_all: true
      description: A running counter of all lines with their corresponding
        labels
      type: Counter
```

* [1558](https://example.com/acme/logstore/pull/1558) **owen-d**: ingester.max-chunk-age
* [1572](https://example.com/acme/logstore/pull/1572) **owen-d**: Feature/query ingesters within

These two configs let you set the max time a chunk can stay in memory in Logstore, this is useful to keep memory usage down as well as limit potential loss of data if ingesters crash.  Combine this with the `query_ingesters_within` config and you can have your queriers skip asking the ingesters for data which you know won't still be in memory (older than max_chunk_age).

**NOTE** Do not set the `max_chunk_age` too small, the default of 1h is probably a good point for most people.  Logstore does not perform well when you flush many small chunks (such as when your logs have too much cardinality), setting this lower than 1h risks flushing too many small chunks.

* [1581](https://example.com/acme/logstore/pull/1581) **slim-bean**: Add sleep to canary reconnect on error

This isn't a feature but it's an important fix, this is the second time our canaries have tried to DDOS our Logstore clusters so you should update to prevent them from trying to attack you.  Aggressive little things these canaries...

* [1840](https://example.com/acme/logstore/pull/1840) **slim-bean**: promtail: Retry 429 rate limit errors from Logstore, increase default retry limits
* [1845](https://example.com/acme/logstore/pull/1845) **wardbekker**: throw exceptions on HTTPTooManyRequests and HTTPServerError so Fluentd will retry

These two PR's change how 429 HTTP Response codes are handled (Rate Limiting), previously these responses were dropped, now they will be retried for these clients

    * Promtail
    * Docker logging driver
    * Fluent Bit
    * Fluentd

This pushes the failure to send logs to two places. First is the retry limits. The defaults in promtail (and thus also the Docker logging driver and Fluent Bit, which share the same underlying code) will retry 429s (and 500s) on an exponential backoff for up to about 8.5 mins on the default configurations. (This can be changed; see the [config docs](https://example.com/acme/logstore/blob/v1.4.0/docs/clients/promtail/configuration.md#client_config) for more info.)

The second place would be the log file itself. At some point, most log files roll based on size or time. Promtail makes an attempt to read a rolled log file but will only try once. If you are very sensitive to lost logs, give yourself really big log files with size-based rolling rules and increase those retry timeouts. This should protect you from Logstore server outages or network issues.

### All Changes

There are many other important fixes and improvements to Logstore, way too many to call out in individual detail, so take a look!

#### Logstore

* [1810](https://example.com/acme/logstore/pull/1810) **cyriltovena**: Optimize empty filter queries.
* [1809](https://example.com/acme/logstore/pull/1809) **cyriltovena**: Test stats memchunk
* [1807](https://example.com/acme/logstore/pull/1807) **pracucci**: Enable global limits by default in production mixin
* [1802](https://example.com/acme/logstore/pull/1802) **cyriltovena**: Add a test for duplicates count in the heap iterator and fixes it.
* [1799](https://example.com/acme/logstore/pull/1799) **pstibrany**: logstore: update Corestore to master
* [1797](https://example.com/acme/logstore/pull/1797) **cyriltovena**: Use ingester client GRPC call options from config.
* [1794](https://example.com/acme/logstore/pull/1794) **pstibrany**: logstore: Convert module names to string
* [1793](https://example.com/acme/logstore/pull/1793) **johncming**: pkg/chunkenc: fix leak of pool.
* [1789](https://example.com/acme/logstore/pull/1789) **adityacs**: Fix logstore exit on jaeger agent not being present
* [1787](https://example.com/acme/logstore/pull/1787) **cyriltovena**: Regexp simplification
* [1785](https://example.com/acme/logstore/pull/1785) **pstibrany**: Update Corestore to master
* [1758](https://example.com/acme/logstore/pull/1758) **cyriltovena**: Query range should not support date where start == end.
* [1750](https://example.com/acme/logstore/pull/1750) **talham7391**: Clearer error response from push endpoint when labels are malformed
* [1746](https://example.com/acme/logstore/pull/1746) **cyriltovena**: Update corestore vendoring to include frontend status code improvement.
* [1745](https://example.com/acme/logstore/pull/1745) **cyriltovena**: Refactor querier http error handling.
* [1736](https://example.com/acme/logstore/pull/1736) **adityacs**: Add /ready endpoint to table-manager
* [1733](https://example.com/acme/logstore/pull/1733) **cyriltovena**: This logs queries with latency tag when  recording stats.
* [1730](https://example.com/acme/logstore/pull/1730) **adityacs**: Fix nil pointer dereference in ingester client
* [1719](https://example.com/acme/logstore/pull/1719) **cyriltovena**: Expose QueryType function.
* [1718](https://example.com/acme/logstore/pull/1718) **cyriltovena**: Better logql metric status code.
* [1708](https://example.com/acme/logstore/pull/1708) **cyriltovena**: Increase discarded samples when line is too long.
* [1704](https://example.com/acme/logstore/pull/1704) **owen-d**: api support for scalars
* [1686](https://example.com/acme/logstore/pull/1686) **owen-d**: max line lengths (component + tenant overrides)
* [1684](https://example.com/acme/logstore/pull/1684) **cyriltovena**: Ensure status codes are set correctly in the frontend.
* [1677](https://example.com/acme/logstore/pull/1677) **owen-d**: Literal Expressions in LogQL
* [1662](https://example.com/acme/logstore/pull/1662) **owen-d**: Binary operators in LogQL
* [1661](https://example.com/acme/logstore/pull/1661) **cyriltovena**: Frontend & Querier query statistics instrumentation.
* [1651](https://example.com/acme/logstore/pull/1651) **owen-d**: removes duplicate logRangeExprExt grammar
* [1636](https://example.com/acme/logstore/pull/1636) **cyriltovena**: Fixes stats summary computation.
* [1630](https://example.com/acme/logstore/pull/1630) **owen-d**: adds stringer methods for all ast expr types
* [1626](https://example.com/acme/logstore/pull/1626) **owen-d**: compiler guarantees for logql exprs
* [1616](https://example.com/acme/logstore/pull/1616) **owen-d**: cache key cant be reused when an interval changes
* [1615](https://example.com/acme/logstore/pull/1615) **cyriltovena**: Add statistics to query_range and instant_query API.
* [1612](https://example.com/acme/logstore/pull/1612) **owen-d**: bumps corestore to 0.6.1 master
* [1605](https://example.com/acme/logstore/pull/1605) **owen-d**: Decouple logql engine/AST from execution context
* [1582](https://example.com/acme/logstore/pull/1582) **slim-bean**: Change new stats names
* [1579](https://example.com/acme/logstore/pull/1579) **rfratto**: Disable transfers in logstore-local-config.yaml
* [1572](https://example.com/acme/logstore/pull/1572) **owen-d**: Feature/query ingesters within
* [1677](https://example.com/acme/logstore/pull/1677) **owen-d**: Introduces numeric literals in LogQL
* [1569](https://example.com/acme/logstore/pull/1569) **owen-d**: refactors splitby to not require buffered channels
* [1567](https://example.com/acme/logstore/pull/1567) **owen-d**: adds span metadata for split queries
* [1565](https://example.com/acme/logstore/pull/1565) **owen-d**: Feature/per tenant splitby
* [1562](https://example.com/acme/logstore/pull/1562) **sandeepsukhani**: limit for concurrent tail requests
* [1558](https://example.com/acme/logstore/pull/1558) **owen-d**: ingester.max-chunk-age
* [1484](https://example.com/acme/logstore/pull/1484) **pstibrany**: logstore: use new runtimeconfig package from Corestore

#### Promtail

* [1840](https://example.com/acme/logstore/pull/1840) **slim-bean**: promtail: Retry 429 rate limit errors from Logstore, increase default retry limits
* [1775](https://example.com/acme/logstore/pull/1775) **slim-bean**: promtail: remove the read lines counter when the log file stops being tailed
* [1770](https://example.com/acme/logstore/pull/1770) **adityacs**: Fix single job with multiple service discovery elements
* [1765](https://example.com/acme/logstore/pull/1765) **adityacs**: Fix error in templating when extracted key has nil value
* [1743](https://example.com/acme/logstore/pull/1743) **dtennander**: Promtail: Ignore dropped entries in subsequent metric-stages in pipelines.
* [1687](https://example.com/acme/logstore/pull/1687) **adityacs**: Fix panic in labels debug message
* [1683](https://example.com/acme/logstore/pull/1683) **slim-bean**: promtail: auto-prune stale metrics
* [1678](https://example.com/acme/logstore/pull/1678) **slim-bean**: promtail: metrics pipeline count all log lines
* [1666](https://example.com/acme/logstore/pull/1666) **adityacs**: Support entire extracted value map in template pipeline stage
* [1664](https://example.com/acme/logstore/pull/1664) **adityacs**: Support custom prefix name in metrics stage
* [1660](https://example.com/acme/logstore/pull/1660) **rfratto**: pkg/promtail/positions: handle empty positions file
* [1652](https://example.com/acme/logstore/pull/1652) **cyriltovena**: --dry-run Promtail.
* [1649](https://example.com/acme/logstore/pull/1649) **cyriltovena**: Pipe data to Promtail
* [1602](https://example.com/acme/logstore/pull/1602) **slim-bean**: Improve promtail configuration docs

#### Helm

* [1731](https://example.com/acme/logstore/pull/1731) **billimek**: [promtail helm chart] - Expand promtail syslog svc to support values
* [1688](https://example.com/acme/logstore/pull/1688) **fredgate**: Logstore stack helm chart can deploy datasources without Acme
* [1632](https://example.com/acme/logstore/pull/1632) **lukipro**: Added support for imagePullSecrets in Logstore Helm chart
* [1620](https://example.com/acme/logstore/pull/1620) **rsteneteg**: [promtail helm chart] option to set fs.inotify.max_user_instances with init container
* [1617](https://example.com/acme/logstore/pull/1617) **billimek**: [promtail helm chart] Enable support for syslog service
* [1590](https://example.com/acme/logstore/pull/1590) **polar3130**: Helm/logstore-stack: refresh default acme.image.tag to 6.6.0
* [1587](https://example.com/acme/logstore/pull/1587) **polar3130**: Helm/logstore-stack: add template for the service name to connect to logstore
* [1585](https://example.com/acme/logstore/pull/1585) **monotek**: [logstore helm chart] added ingress
* [1553](https://example.com/acme/logstore/pull/1553) **got-root**: helm: Allow setting 'loadBalancerSourceRanges' for the logstore service
* [1529](https://example.com/acme/logstore/pull/1529) **tourea**: Promtail Helm Chart: Add support for passing environment variables

#### Jsonnet

* [1776](https://example.com/acme/logstore/pull/1776) **Eraac**: fix typo: Not a binary operator: =
* [1767](https://example.com/acme/logstore/pull/1767) **joe-elliott**: Dashboard Cleanup
* [1766](https://example.com/acme/logstore/pull/1766) **joe-elliott**: Move dashboards out into their own json files
* [1757](https://example.com/acme/logstore/pull/1757) **slim-bean**: promtail-mixin: Allow choosing promtail name
* [1756](https://example.com/acme/logstore/pull/1756) **sh0rez**: fix(ksonnet): named parameters for containerPort
* [1749](https://example.com/acme/logstore/pull/1749) **slim-bean**: Increasing the threshold for a file lag and reducing the severity to warning
* [1748](https://example.com/acme/logstore/pull/1748) **slim-bean**: jsonnet: Breakout promtail mixin.
* [1739](https://example.com/acme/logstore/pull/1739) **cyriltovena**: Fixes frontend args in libsonnet.
* [1735](https://example.com/acme/logstore/pull/1735) **cyriltovena**: Allow to configure global limits via the jsonnet deployment.
* [1705](https://example.com/acme/logstore/pull/1705) **cyriltovena**: Add overrides file for our jsonnet library.
* [1699](https://example.com/acme/logstore/pull/1699) **pracucci**: Increased production distributors memory request and limit
* [1689](https://example.com/acme/logstore/pull/1689) **shokada**: Add headers for WebSocket
* [1665](https://example.com/acme/logstore/pull/1665) **cyriltovena**: Query frontend service should be headless.
* [1613](https://example.com/acme/logstore/pull/1613) **cyriltovena**: Fixes config change in the result cache

#### Fluent Bit

* [1791](https://example.com/acme/logstore/pull/1791) **cyriltovena**: Improve fluentbit logfmt.
* [1717](https://example.com/acme/logstore/pull/1717) **adityacs**: Fluent-bit: Fix panic error when AutoKubernetesLabels is true

#### Fluentd

* [1811](https://example.com/acme/logstore/pull/1811) **JamesJJ**: Error handling: Show data stream at "debug" level, not "warn"
* [1728](https://example.com/acme/logstore/pull/1728) **irake99**: docs: fix outdated link to fluentd
* [1703](https://example.com/acme/logstore/pull/1703) **Skeen**:  fluent-plugin-acme-logstore: Update fluentd base image to current images (edge)
* [1656](https://example.com/acme/logstore/pull/1656) **takanabe**: Convert second(Integer class) to nanosecond precision
* [1646](https://example.com/acme/logstore/pull/1646) **takanabe**: Fix rubocop violation for fluentd/fluent-plugin-logstore
* [1603](https://example.com/acme/logstore/pull/1603) **tarokkk**: fluentd-plugin: add URI validation

#### Docs

* [1781](https://example.com/acme/logstore/pull/1781) **candlerb**: Docs: Recommended schema is now v11
* [1771](https://example.com/acme/logstore/pull/1771) **rfratto**: change slack url to slack.acme.com and use https
* [1738](https://example.com/acme/logstore/pull/1738) **jgehrcke**: docs: observability.md: clarify lines vs. entries
* [1707](https://example.com/acme/logstore/pull/1707) **dangoodman**: Fix regex in pipeline-example.yml
* [1697](https://example.com/acme/logstore/pull/1697) **oke-py**: fix promtail/templates/NOTES.txt to show correctly port-forward command
* [1675](https://example.com/acme/logstore/pull/1675) **owen-d**: maintainer links & usernames
* [1673](https://example.com/acme/logstore/pull/1673) **cyriltovena**: Add Owen to the maintainer team.
* [1671](https://example.com/acme/logstore/pull/1671) **shokada**: Update tanka.md so that promtail.yml is the correct format
* [1648](https://example.com/acme/logstore/pull/1648) **ShotaKitazawa**: logstore-canary: fix indent of DaemonSet manifest written in .md file
* [1642](https://example.com/acme/logstore/pull/1642) **slim-bean**: Improve systemd field docs
* [1641](https://example.com/acme/logstore/pull/1641) **pastatopf**: Correct syntax of rate example
* [1634](https://example.com/acme/logstore/pull/1634) **takanabe**: Unite docs for fluentd plugin
* [1619](https://example.com/acme/logstore/pull/1619) **shaikatz**: PeriodConfig documentation fix dynamodb -> aws-dynamo
* [1611](https://example.com/acme/logstore/pull/1611) **owen-d**: logstore frontend docs additions
* [1609](https://example.com/acme/logstore/pull/1609) **Lusitaniae**: Fix wget syntax in documentation
* [1608](https://example.com/acme/logstore/pull/1608) **PabloCastellano**: Documentation: Recommend using the latest schema version (v11)
* [1601](https://example.com/acme/logstore/pull/1601) **rfratto**: Clarify regex escaping rules
* [1598](https://example.com/acme/logstore/pull/1598) **cyriltovena**: Update tanka.md doc.
* [1586](https://example.com/acme/logstore/pull/1586) **MrSaints**: Fix typo in changelog for 1.3.0
* [1504](https://example.com/acme/logstore/pull/1504) **hsraju**: Updated configuration.md

#### Logcli

* [1808](https://example.com/acme/logstore/pull/1808) **slim-bean**: logcli: log the full stats and send to stderr instead of stdout
* [1682](https://example.com/acme/logstore/pull/1682) **adityacs**: BugFix: Fix logcli --quiet parameter parsing issue
* [1644](https://example.com/acme/logstore/pull/1644) **cyriltovena**: This improves the log output for statistics in the logcli.
* [1638](https://example.com/acme/logstore/pull/1638) **owen-d**: adds query stats and org id options in logcli
* [1573](https://example.com/acme/logstore/pull/1573) **cyriltovena**: Improve logql query statistics collection.

#### Logstore Canary

* [1653](https://example.com/acme/logstore/pull/1653) **slim-bean**: Canary needs its logo
* [1581](https://example.com/acme/logstore/pull/1581) **slim-bean**: Add sleep to canary reconnect on error

#### Build

* [1780](https://example.com/acme/logstore/pull/1780) **slim-bean**: build: Update the CD deploy task name
* [1762](https://example.com/acme/logstore/pull/1762) **dgzlopes**: Bump testify to 1.5.1
* [1742](https://example.com/acme/logstore/pull/1742) **slim-bean**: build: fix deploy on tagged build
* [1741](https://example.com/acme/logstore/pull/1741) **slim-bean**: add darwin and freebsd binaries to release output
* [1740](https://example.com/acme/logstore/pull/1740) **rfratto**: Fix 32-bit Promtail ARM docker builds from Drone
* [1710](https://example.com/acme/logstore/pull/1710) **adityacs**: Add goimport local-prefixes configuration to .golangci.yml
* [1647](https://example.com/acme/logstore/pull/1647) **mattmendick**: Attempting to add `informational` only feedback for codecov
* [1640](https://example.com/acme/logstore/pull/1640) **rfratto**: ci: print error messages when an API request fails
* [1639](https://example.com/acme/logstore/pull/1639) **rfratto**: ci: prune docker tags prefixed with "master-" older than 90 days
* [1637](https://example.com/acme/logstore/pull/1637) **rfratto**: ci: pin plugins/manifest image tag
* [1633](https://example.com/acme/logstore/pull/1633) **rfratto**: ci: make manifest publishing run in serial
* [1629](https://example.com/acme/logstore/pull/1629) **slim-bean**: Ignore markdown files in codecoverage
* [1628](https://example.com/acme/logstore/pull/1628) **rfratto**: Exempt proposals from stale bot
* [1614](https://example.com/acme/logstore/pull/1614) **mattmendick**: Codecov: Update config to add informational flag
* [1600](https://example.com/acme/logstore/pull/1600) **mattmendick**: Codecov circleci test [WIP]

#### Tooling

* [1577](https://example.com/acme/logstore/pull/1577) **pstibrany**: Move chunks-inspect tool to Logstore repo

## 1.3.0 (2020-01-16)

### What's New?? ###

With 1.3.0 we are excited to announce several improvements focusing on performance!

First and most significant is the Query Frontend:

* [1442](https://example.com/acme/logstore/pull/1442) **cyriltovena**: Logstore Query Frontend

The query frontend allows for sharding queries by time and dispatching them in parallel to multiple queriers, giving true horizontal scaling ability for queries.  Take a look at the [jsonnet changes](https://example.com/acme/logstore/pull/1442/files?file-filters%5B%5D=.libsonnet) to see how we are deploying this in our production setup.  Keep an eye out for a blog post with more information on how the frontend works and more information on this exciting new feature.

In our quest to improve query performance, we discovered that gzip, while good for compression ratio, is not the best for speed.  So we introduced the ability to select from several different compression algorithms:

* [1411](https://example.com/acme/logstore/pull/1411) **cyriltovena**: Adds configurable compression algorithms for chunks

We are currently testing out LZ4 and snappy, LZ4 seemed like a good fit however we found that it didn't always compress the same data to the same output which was causing some troubles for another important improvement:

* [1438](https://example.com/acme/logstore/pull/1438) **pstibrany**: pkg/ingester: added sync period flags

Extending on the work done by @bboreham on Corestore, @pstibrany added a few new flags and code to synchronize chunks between ingesters, which reduces the number of chunks persisted to object stores and therefore also reduces the number of chunks loaded on queries and the amount of de-duplication work which needs to be done.

As mentioned above, LZ4 was in some cases compressing the same data with a different result which was interfering with this change, we are still investigating the cause of this issue (It may be in how we implemented something, or may be in the compression code itself).  For now we have switched to snappy which has seen a reduction in data written to the object store from almost 3x the source data (with a replication factor of 3) to about 1.5x, saving a lot of duplicated log storage!

Another valuable change related to chunks:

* [1406](https://example.com/acme/logstore/pull/1406) **slim-bean**: allow configuring a target chunk size in compressed bytes

With this change you can set a `chunk_target_size` and Logstore will attempt to fill a chunk to approx that size before flushing (previously a chunk size was a hard coded 10 blocks where the default block size is 262144 bytes).  Larger chunks are beneficial for a few reasons, mainly on reducing API calls to your object store when performing queries, but also in reducing overhead in a few places, especially when processing very high volume log streams.

Another big improvement is the introduction of accurate rate limiting when running microservices:

* [1486](https://example.com/acme/logstore/pull/1486) **pracucci**: Add ingestion rate global limit support

Previously the rate limit was applied at each distributor, however with traffic split over many distributors the limit would need to be adjusted accordingly.  This meant that scaling up distributors required changing the limit.  Now this information is communicated between distributors such that the limit should be applied accurately regardless of the number of distributors.

And last but not least on the notable changes list is a new feature for Promtail:

* [1275](https://example.com/acme/logstore/pull/1275) **bastjan**: pkg/promtail: IETF Syslog (RFC5424) Support

With this change Promtail can receive syslogs via TCP!  Thanks to @bastjan for all the hard work on this submission!

### Important things to note

* [1519](https://example.com/acme/logstore/pull/1519) Changes a core behavior in Logstore regarding logs with duplicate content AND duplicate timestamps, previously Logstore would store logs with duplicate timestamps and content, moving forward logs with duplicate content AND timestamps will be silently ignored.  Mainly this change is to prevent duplicates that appear when a batch is retried (the first entry in the list would be inserted again, now it will be ignored).  Logs with the same timestamp and different content will still be accepted.
* [1486](https://example.com/acme/logstore/pull/1486) Deprecated `-distributor.limiter-reload-period` flag / distributor's `limiter_reload_period` config option.

### All Changes

Once again we can't thank our community and contributors enough for the significant work that everyone is adding to Logstore, the entire list of changes is long!!

#### Logstore

* [1526](https://example.com/acme/logstore/pull/1526) **codesome**: Support <selector> <range> <filters> for aggregation
* [1522](https://example.com/acme/logstore/pull/1522) **cyriltovena**: Adds support for the old query string regexp in the frontend.
* [1519](https://example.com/acme/logstore/pull/1519) **rfratto**: pkg/chunkenc: ignore duplicate lines pushed to a stream
* [1511](https://example.com/acme/logstore/pull/1511) **sandlis**: querier: fix panic in tailer when max tail duration exceeds
* [1499](https://example.com/acme/logstore/pull/1499) **slim-bean**: Fix a panic in chunk prefetch
* [1495](https://example.com/acme/logstore/pull/1495) **slim-bean**: Prefetch chunks while processing
* [1496](https://example.com/acme/logstore/pull/1496) **cyriltovena**: Add duplicates info and remove timing informations.
* [1490](https://example.com/acme/logstore/pull/1490) **owen-d**: Fix/deadlock frontend queue
* [1489](https://example.com/acme/logstore/pull/1489) **owen-d**: unifies reverse iterators
* [1488](https://example.com/acme/logstore/pull/1488) **cyriltovena**: Fixes response json encoding and add regression tests.
* [1486](https://example.com/acme/logstore/pull/1486) **pracucci**: Add ingestion rate global limit support* [1493](https://example.com/acme/logstore/pull/1493) **pracucci**: Added max streams per user global limit
* [1480](https://example.com/acme/logstore/pull/1480) **cyriltovena**: Close iterator properly and check nil before releasing buffers.
* [1473](https://example.com/acme/logstore/pull/1473) **rfratto**: pkg/querier: don't query all ingesters
* [1470](https://example.com/acme/logstore/pull/1470) **cyriltovena**: Validates limit parameter.
* [1448](https://example.com/acme/logstore/pull/1448) **cyriltovena**: Improving storage benchmark
* [1445](https://example.com/acme/logstore/pull/1445) **cyriltovena**: Add decompression tracing instrumentation.
* [1442](https://example.com/acme/logstore/pull/1442) **cyriltovena**: Logstore Query Frontend
* [1438](https://example.com/acme/logstore/pull/1438) **pstibrany**: pkg/ingester: added sync period flags
* [1433](https://example.com/acme/logstore/pull/1433) **zendern**: Using strict parsing for yaml configs
* [1425](https://example.com/acme/logstore/pull/1425) **pstibrany**: pkg/ingester: Added possibility to disable transfers.
* [1423](https://example.com/acme/logstore/pull/1423) **pstibrany**: pkg/chunkenc: Fix BenchmarkRead to focus on reading chunks, not converting bytes to string
* [1421](https://example.com/acme/logstore/pull/1421) **pstibrany**: pkg/chunkenc: change default LZ4 buffer size to 64k.
* [1420](https://example.com/acme/logstore/pull/1420) **cyriltovena**: Sets the chunk encoding correctly when creating chunk from bytes.
* [1419](https://example.com/acme/logstore/pull/1419) **owen-d**: Enables Series API in logstore
* [1413](https://example.com/acme/logstore/pull/1413) **pstibrany**: RangeQuery benchmark optimizations
* [1411](https://example.com/acme/logstore/pull/1411) **cyriltovena**: Adds configurable compression algorithms for chunks
* [1409](https://example.com/acme/logstore/pull/1409) **slim-bean**: change the chunk size histogram to allow for bigger buckets
* [1408](https://example.com/acme/logstore/pull/1408) **slim-bean**: forgot to register the new metric for counting blocks per chunk
* [1406](https://example.com/acme/logstore/pull/1406) **slim-bean**: allow configuring a target chunk size in compressed bytes
* [1405](https://example.com/acme/logstore/pull/1405) **pstibrany**: Convert string to bytes once only when doing string filtering.
* [1396](https://example.com/acme/logstore/pull/1396) **pstibrany**: pkg/cfg: print help only when requested, and print it on stdout
* [1383](https://example.com/acme/logstore/pull/1383) **beornf**: Read websocket close in tail handler
* [1071](https://example.com/acme/logstore/pull/1071) **rfratto**: pkg/ingester: limit total number of errors a stream can return on push
* [1545](https://example.com/acme/logstore/pull/1545) **joe-elliott**: Critical n => m conversions
* [1541](https://example.com/acme/logstore/pull/1541) **owen-d**: legacy endpoint 400s metric queries

#### Promtail

* [1515](https://example.com/acme/logstore/pull/1515) **slim-bean**: Promtail: Improve position and size metrics
* [1485](https://example.com/acme/logstore/pull/1485) **p37ruh4**: Fileglob parsing fixes
* [1472](https://example.com/acme/logstore/pull/1472) **owen-d**: positions.ignore-corruptions
* [1453](https://example.com/acme/logstore/pull/1453) **chancez**: pkg/promtail: Initialize counters to 0 when creating client
* [1436](https://example.com/acme/logstore/pull/1436) **rfratto**: promtail: add support for passing through journal entries as JSON
* [1426](https://example.com/acme/logstore/pull/1426) **wphan**: Support microsecond timestamp format
* [1416](https://example.com/acme/logstore/pull/1416) **pstibrany**: pkg/promtail/client: missing URL in client returns error
* [1275](https://example.com/acme/logstore/pull/1275) **bastjan**: pkg/promtail: IETF Syslog (RFC5424) Support

#### Fluent Bit

* [1455](https://example.com/acme/logstore/pull/1455) **JensErat**: fluent-bit-plugin: re-enable failing JSON marshaller tests; pass error instead of logging and ignoring
* [1294](https://example.com/acme/logstore/pull/1294) **JensErat**: fluent-bit: multi-instance support
* [1514](https://example.com/acme/logstore/pull/1514) **shane-axiom**: fluent-plugin-acme-logstore: Add `fluentd_thread` label when `flush_thread_count` > 1

#### Fluentd

* [1500](https://example.com/acme/logstore/pull/1500) **cyriltovena**: Bump fluentd plugin to 1.2.6.
* [1475](https://example.com/acme/logstore/pull/1475) **Horkyze**: fluentd-plugin: call gsub for strings only

#### Docker Driver

* [1414](https://example.com/acme/logstore/pull/1414) **cyriltovena**: Adds tenant-id for docker driver.

#### Logcli

* [1492](https://example.com/acme/logstore/pull/1492) **sandlis**: logcli: replaced ACME_*with LOGSTORE_* in logcli env vars, set default server url for logcli to localhost

#### Helm

* [1534](https://example.com/acme/logstore/pull/1534) **olivierboudet**: helm : fix fluent-bit parser configuration syntax
* [1506](https://example.com/acme/logstore/pull/1506) **terjesannum**: helm: add podsecuritypolicy for fluent-bit
* [1431](https://example.com/acme/logstore/pull/1431) **eugene100**: Helm: fix issue with config.clients
* [1430](https://example.com/acme/logstore/pull/1430) **olivierboudet**: helm : allow to define custom parsers to use with fluentbit.io/parser annotation
* [1418](https://example.com/acme/logstore/pull/1418) **evalsocket**: Helm chart url added in helm.md
* [1336](https://example.com/acme/logstore/pull/1336) **terjesannum**: helm: support adding init containers to the logstore pod
* [1530](https://example.com/acme/logstore/pull/1530) **WeiBanjo**: Allow extra command line args for external labels like hostname

#### Jsonnet

* [1518](https://example.com/acme/logstore/pull/1518) **benjaminhuo**: Fix error 'Field does not exist: jaeger_mixin' in tk show
* [1501](https://example.com/acme/logstore/pull/1501) **anarcher**: jsonnet: fix common/defaultPorts parameters
* [1497](https://example.com/acme/logstore/pull/1497) **cyriltovena**: Update Logstore mixin to include frontend QPS and latency.
* [1478](https://example.com/acme/logstore/pull/1478) **cyriltovena**: Fixes the typo in the result cache config of the Logstore ksonnet lib.
* [1543](https://example.com/acme/logstore/pull/1543) **sh0rez**: fix(ksonnet): use apps/v1

#### Docs

* [1531](https://example.com/acme/logstore/pull/1531) **fitzoh**: Documentation: Add note on using Logstore with Amazon ECS
* [1521](https://example.com/acme/logstore/pull/1521) **rfratto**: docs: Document timestamp ordering rules
* [1516](https://example.com/acme/logstore/pull/1516) **rfratto**: Link to release docs in README.md, not master docs
* [1508](https://example.com/acme/logstore/pull/1508) **cyriltovena**: Fixes bad json in Logstore API documentation.
* [1505](https://example.com/acme/logstore/pull/1505) **sandlis**: doc: fix sample yaml in docs for installing promtail to k8s
* [1481](https://example.com/acme/logstore/pull/1481) **terjesannum**: docs: fix broken promtail link
* [1474](https://example.com/acme/logstore/pull/1474) **Eraac**: <doc>: information about max_look_back_period
* [1471](https://example.com/acme/logstore/pull/1471) **cyriltovena**: Update README.md
* [1466](https://example.com/acme/logstore/pull/1466) **Eraac**: <documentation>: Update IAM requirement
* [1441](https://example.com/acme/logstore/pull/1441) **vtereso**: <Docs>: README spelling fix
* [1437](https://example.com/acme/logstore/pull/1437) **daixiang0**: fix all misspell
* [1432](https://example.com/acme/logstore/pull/1432) **joe-elliott**: Removed unsupported encodings from docs
* [1399](https://example.com/acme/logstore/pull/1399) **vishesh92**: Docs: Add configuration docs for redis
* [1394](https://example.com/acme/logstore/pull/1394) **chancez**: Documentation: Fix example AWS storage configuration
* [1227](https://example.com/acme/logstore/pull/1227) **daixiang0**: Add docker install doc
* [1560](https://example.com/acme/logstore/pull/1560) **robshep**: Promtail Docs: Update output.md
* [1546](https://example.com/acme/logstore/pull/1546) **mattmendick**: Removing third-party link
* [1539](https://example.com/acme/logstore/pull/1539) **j18e**: docs: fix syntax error in pipeline example

#### Build

* [1494](https://example.com/acme/logstore/pull/1494) **pracucci**: Fixed TOUCH_PROTOS in all DroneCI pipelines
* [1479](https://example.com/acme/logstore/pull/1479) **owen-d**: TOUCH_PROTOS build arg for dockerfile
* [1476](https://example.com/acme/logstore/pull/1476) **owen-d**: initiates docker daemon for circle windows builds
* [1469](https://example.com/acme/logstore/pull/1469) **rfratto**: Makefile: re-enable journal scraping on ARM

#### New Members

* [1415](https://example.com/acme/logstore/pull/1415) **cyriltovena**: Add Joe as member of the team.

# 1.2.0 (2019-12-09)

One week has passed since the last Logstore release, and it's time for a new one!

## Notable Changes

We have continued our work making our API Prometheus-compatible. The key
changes centered around API compatibility are:

* [1370](https://example.com/acme/logstore/pull/1370) **slim-bean**: Change `/logstore/api/v1/label` to `logstore/api/v1/labels`
* [1381](https://example.com/acme/logstore/pull/1381) **owen-d**: application/x-www-form-urlencoded support

Meanwhile, @pstibrany has done great work ensuring that Logstore handles hash
collisions properly:

* [1247](https://example.com/acme/logstore/pull/1247) **pstibrany**: pkg/ingester: handle labels mapping to the same fast fingerprint.

## Other Changes

:heart: All PR's are important to us, thanks everyone for continuing to help support and improve Logstore! :heart:

### Features

* [1372](https://example.com/acme/logstore/pull/1372) **cyriltovena**: Let Logstore start when using the debug image.
* [1300](https://example.com/acme/logstore/pull/1300) **pstibrany**: pkg/ingester: check that ingester is in LEAVING state when transferring chunks and claiming tokens. Required when using memberlist client.

### Bug Fixes/Improvements

* [1376](https://example.com/acme/logstore/pull/1376) **jstaffans**: Fluentd: guard against nil values when sanitizing labels
* [1371](https://example.com/acme/logstore/pull/1371) **cyriltovena**: Logql benchmark and performance improvement.
* [1363](https://example.com/acme/logstore/pull/1363) **cyriltovena**: Fixes fluentd new push path API.
* [1353](https://example.com/acme/logstore/pull/1353) **pstibrany**: docs: Fix grpc_listen_host and http_listen_host.
* [1350](https://example.com/acme/logstore/pull/1350) **Eraac**: documentation: iam requirement for autoscaling

# 1.1.0 (2019-12-04)

It's been a busy 2 weeks since the 1.0.0 release and quite a few important PR's have been merged to Logstore.

The most significant:

* [1322](https://example.com/acme/logstore/pull/1322) **rfratto**: Fix v1 label API to be Prometheus-compatible

Some might call this a **breaking change**, we are instead calling it a bug fix as our goal was to be prometheus compatible and we were not :smiley:

**But please be aware if you are using the `/logstore/api/v1/label` or `/logstore/api/v1/label/<name>/values` the JSON result will be different in 1.1.0**

Old result:

```json
{
  "values": [
    "label1",
    "label2",
    "labeln"
  ]
}
```

New result:

```json
{
  "status": "success",
  "data": [
    "label1",
    "label2",
    "labeln"
  ]
}
```

**ALSO IMPORTANT**

* [1160](https://example.com/acme/logstore/pull/1160) **daixiang0**: replace gzip with zip

Binaries will now be zipped instead of gzipped as many people voiced their opinion that zip is likely to be installed on more systems by default.

**If you had existing automation to download and install binaries this will have to be updated to use zip instead of gzip**

## Notable Fixes and Improvements

* Broken version info in startup log message:

    [1095](https://example.com/acme/logstore/pull/1095) **pstibrany**: Makefile changes to allow easy builds with or without vendoring. Also fixes version bug for both cases.

* The hashing algorithm used to calculate the hash for a stream was creating hash collisions in some instances.
**Please Note** this is just one part of the fix and is only in Promtail, the second part for Logstore can be tracked [in PR1247](https://example.com/acme/logstore/pull/1247) which didn't quite make the cut for 1.1.0 and will be in 1.2.0:

    [1254](https://example.com/acme/logstore/pull/1254) **pstibrany**: pkg/promtail/client: Handle fingerprint hash collisions

* Thank you @putrasattvika for finding and fixing an important bug where logs were some logs were missed in a query shortly after a flush!

    [1299](https://example.com/acme/logstore/pull/1299) **putrasattvika**: storage: fix missing logs with batched chunk iterator

* Thank you @danieldabate for helping to again improve our API to be more Prometheus compatible:

    [1355](https://example.com/acme/logstore/pull/1355) **danieldabate**: HTTP API: Support duration and float formats for step parameter

* LogQL will support duration formats that are not typically handled by Go like [1d] or [1w]

    [1357](https://example.com/acme/logstore/pull/1357) **cyriltovena**: Supports same duration format in LogQL as Prometheus

## Everything Else

:heart: All PR's are important to us, thanks everyone for continuing to help support and improve Logstore! :heart:

* [1349](https://example.com/acme/logstore/pull/1349) **Eraac**: documentation: using parsable value in example
* [1343](https://example.com/acme/logstore/pull/1343) **dgzlopes**: doc(configuration): Fix duration format.
* [1342](https://example.com/acme/logstore/pull/1342) **whothey**: Makefile: add debug symbols to logstore and promtail debug builds
* [1341](https://example.com/acme/logstore/pull/1341) **adamjohnson01**: Update logstore helm chart to support service account annotations
* [1340](https://example.com/acme/logstore/pull/1340) **adamjohnson01**: Pull in corestore changes to support IAM roles for EKS
* [1339](https://example.com/acme/logstore/pull/1339) **cyriltovena**: Update gem version.
* [1333](https://example.com/acme/logstore/pull/1333) **daixiang0**: fix broken link
* [1328](https://example.com/acme/logstore/pull/1328) **cyriltovena**: Fixes linter warning from the yacc file.
* [1326](https://example.com/acme/logstore/pull/1326) **dawidmalina**: Wrong api endpoint in fluent-plugin-acme-logstore
* [1320](https://example.com/acme/logstore/pull/1320) **roidelapluie**: Metrics: use Namespace everywhere when declaring metrics
* [1318](https://example.com/acme/logstore/pull/1318) **roidelapluie**: Use tenant as label name for discarded_samples metrics
* [1317](https://example.com/acme/logstore/pull/1317) **roidelapluie**: Expose discarded bytes metric
* [1316](https://example.com/acme/logstore/pull/1316) **slim-bean**: Removing old file needed for dep (no longer needed)
* [1312](https://example.com/acme/logstore/pull/1312) **ekeih**: Docs: Add missing ) in LogQL example
* [1311](https://example.com/acme/logstore/pull/1311) **pstibrany**: Include positions filename in the error when YAML unmarshal fails.
* [1310](https://example.com/acme/logstore/pull/1310) **JensErat**: fluent-bit: sorted JSON and properly convert []byte to string
* [1304](https://example.com/acme/logstore/pull/1304) **pstibrany**: promtail: write positions to new file first, move to target location afterwards
* [1303](https://example.com/acme/logstore/pull/1303) **zhangjianweibj**: <https://example.com/acme/logstore/issues/1302>
* [1298](https://example.com/acme/logstore/pull/1298) **rfratto**: pkg/promtail: remove journal target forced path
* [1279](https://example.com/acme/logstore/pull/1279) **rfratto**: Fix logstore_discarded_samples_total metric
* [1278](https://example.com/acme/logstore/pull/1278) **rfratto**: docs: update limits_config to new structure from #948
* [1276](https://example.com/acme/logstore/pull/1276) **roidelapluie**: Update fluentbit README.md based on my experience
* [1274](https://example.com/acme/logstore/pull/1274) **sh0rez**: chore(ci): drone-cli
* [1273](https://example.com/acme/logstore/pull/1273) **JensErat**: fluent-bit: tenant ID configuration
* [1266](https://example.com/acme/logstore/pull/1266) **polar3130**: add description about tenant stage
* [1262](https://example.com/acme/logstore/pull/1262) **Eraac**: documentation: iam requirement for autoscaling
* [1261](https://example.com/acme/logstore/pull/1261) **rfratto**: Document systemd journal scraping
* [1249](https://example.com/acme/logstore/pull/1249) **cyriltovena**: Move to jsoniter instead of default json package
* [1223](https://example.com/acme/logstore/pull/1223) **jgehrcke**: authentication.md: replace "user" with "tenant"
* [1204](https://example.com/acme/logstore/pull/1204) **allanhung**: fluent-bit-plugin: Auto add Kubernetes labels to Logstore labels

# 1.0.0 (2019-11-19)

:tada: Nearly a year since Logstore was announced at KubeCon in Seattle 2018 we are very excited to announce the 1.0.0 release of Logstore! :tada:

A lot has happened since the announcement, the project just recently passed 1000 commits by 138 contributors over 700+ PR's accumulating over 7700 GitHub stars!

Internally at Acme Labs we have been using Logstore to monitor all of our infrastructure and ingest around 1.5TB/10 billion log lines a day. Since the v0.2.0 release we have found Logstore to be reliable and stable in our environments.

We are comfortable with the state of the project in our production environments and think it's time to promote Logstore to a non-beta release to communicate to everyone that they should feel comfortable using Logstore in their production environments too.

## API Stability

With the 1.0.0 release our intent is to try to follow Semver rules regarding stability with some aspects of Logstore, focusing mainly on the operating experience of Logstore as an application.  That is to say we are not planning any major changes to the HTTP API, and anything breaking would likely be accompanied by a major release with backwards compatibility support.

We are currently NOT planning on maintaining Go API stability with this release, if you are importing Logstore as a library you should be prepared for any kind of change, including breaking, even in minor or bugfix releases.

Logstore is still a young and active project and there might be some breaking config changes in non-major releases, rest assured this will be clearly communicated and backwards or overlapping compatibility will be provided if possible.

## Changes

There were not as many changes in this release as the last, mainly we wanted to make sure Logstore was mostly stable before 1.0.0.  The most notable change is the inclusion of the V11 schema in PR's [1201](https://example.com/acme/logstore/pull/1201) and [1280](https://example.com/acme/logstore/pull/1280).  The V11 schema adds some more data to the index to improve label queries over large amounts of time and series.  Currently we have not updated the Helm or Ksonnet to use the new schema, this will come soon with more details on how it works.

The full list of changes:

* [1280](https://example.com/acme/logstore/pull/1280) **owen-d**: Fix duplicate labels (update corestore)
* [1260](https://example.com/acme/logstore/pull/1260) **rfratto**: pkg/logstore: unmarshal module name from YAML
* [1257](https://example.com/acme/logstore/pull/1257) **rfratto**: helm: update default terminationGracePeriodSeconds to 4800
* [1251](https://example.com/acme/logstore/pull/1251) **obitech**: docs: Fix promtail releases download link
* [1248](https://example.com/acme/logstore/pull/1248) **rfratto**: docs: slightly modify language in community Logstore packages section
* [1242](https://example.com/acme/logstore/pull/1242) **tarokkk**: fluentd: Suppress unread configuration warning
* [1239](https://example.com/acme/logstore/pull/1239) **pracucci**: Move ReservedLabelTenantID out from a dedicated file
* [1238](https://example.com/acme/logstore/pull/1238) **oke-py**: helm: logstore-stack supports k8s 1.16
* [1237](https://example.com/acme/logstore/pull/1237) **joe-elliott**: Rollback google.golang.org/api to 0.8.0
* [1235](https://example.com/acme/logstore/pull/1235) **woodsaj**: ci: update triggers to use new deployment_tools location
* [1234](https://example.com/acme/logstore/pull/1234) **rfratto**: Standardize schema used in `match` stage
* [1233](https://example.com/acme/logstore/pull/1233) **wapmorgan**: Update docker-driver Dockerfile: add tzdb
* [1232](https://example.com/acme/logstore/pull/1232) **rfratto**: Fix drone deploy job
* [1231](https://example.com/acme/logstore/pull/1231) **joe-elliott**: Removed references to Logstore free tier
* [1226](https://example.com/acme/logstore/pull/1226) **clickyotomy**: Update dependencies to use weaveworks/common upstream
* [1221](https://example.com/acme/logstore/pull/1221) **slim-bean**: use regex label matcher to not alert on any tail route latencies
* [1219](https://example.com/acme/logstore/pull/1219) **MightySCollins**: docs: Updated Kubernetes docs links in Helm charts
* [1218](https://example.com/acme/logstore/pull/1218) **slim-bean**: update dashboards to include the new /logstore/api/v1/* endpoints
* [1217](https://example.com/acme/logstore/pull/1217) **slim-bean**: sum the bad words by name and level
* [1216](https://example.com/acme/logstore/pull/1216) **joe-elliott**: Remove rules that reference no longer existing metrics
* [1215](https://example.com/acme/logstore/pull/1215) **Eraac**: typo url
* [1214](https://example.com/acme/logstore/pull/1214) **takanabe**: Correct wrong document paths about querying
* [1213](https://example.com/acme/logstore/pull/1213) **slim-bean**: Fix docker latest and master tags
* [1212](https://example.com/acme/logstore/pull/1212) **joe-elliott**: Update logstore operational
* [1206](https://example.com/acme/logstore/pull/1206) **sandlis**: ksonnet: fix replication always set to 3 in ksonnet
* [1203](https://example.com/acme/logstore/pull/1203) **joe-elliott**: Chunk iterator performance improvement
* [1202](https://example.com/acme/logstore/pull/1202) **beorn7**: Simplify regexp's
* [1201](https://example.com/acme/logstore/pull/1201) **cyriltovena**: Update corestore to bring v11 schema
* [1189](https://example.com/acme/logstore/pull/1189) **putrasattvika**: fluent-plugin: Add client certificate verification
* [1186](https://example.com/acme/logstore/pull/1186) **tarokkk**: fluentd: Refactor label_keys and and add extract_kubernetes_labels configuration

# 0.4.0 (2019-10-24)

A **huge** thanks to the **36 contributors** who submitted **148 PR's** since 0.3.0!

## Notable Changes

* With PR [654](https://example.com/acme/logstore/pull/654) @cyriltovena added a really exciting new capability to Logstore, a Prometheus compatible API with support for running metric style queries against your logs! [Take a look at how to write metric queries for logs](https://example.com/acme/logstore/blob/master/docs/logql.md#counting-logs)
    > PLEASE NOTE: To use metric style queries in the current Acme release 6.4.x you will need to add Logstore as a Prometheus datasource in addition to having it as a Log datasource and you will have to select the correct source for querying logs vs metrics, coming soon Acme will support both logs and metric queries directly to the Logstore datasource!
* PR [1022](https://example.com/acme/logstore/pull/1022) (and a few others) @joe-elliott added a new set of HTTP endpoints in conjunction with the work @cyriltovena to create a Prometheus compatible API as well as improve how labels/timestamps are handled
    > IMPORTANT: The new `/api/v1/*` endpoints contain breaking changes on the query paths (push path is unchanged) Eventually the `/api/prom/*` endpoints will be removed
* PR [847](https://example.com/acme/logstore/pull/847) owes a big thanks to @cosmo0920 for contributing his Fluent Bit go plugin, now logstore has Fluent Bit plugin support!!

* PR [982](https://example.com/acme/logstore/pull/982) was a couple weeks of painstaking work by @rfratto for a much needed improvement to Logstore's docs! [Check them out!](https://example.com/acme/logstore/tree/master/docs)

* PR [980](https://example.com/acme/logstore/pull/980) by @sh0rez improved how flags and config file's are loaded to honor a more traditional order of precedence:
    1. Defaults
    2. Config file
    3. User-supplied flag values (command line arguments)
    > PLEASE NOTE: This is potentially a breaking change if you were passing command line arguments that also existed in a config file in which case the order they are given priority now has changed!

* PR [1062](https://example.com/acme/logstore/pull/1062) and [1089](https://example.com/acme/logstore/pull/1089) have moved Logstore from Dep to Go Modules and to Go 1.13

## Logstore

### Features/Improvements/Changes

* **Logstore** [1171](https://example.com/acme/logstore/pull/1171) **cyriltovena**: Moves request parsing into the loghttp package
* **Logstore** [1145](https://example.com/acme/logstore/pull/1145) **joe-elliott**: Update `/logstore/api/v1/push` to use the v1 json format
* **Logstore** [1128](https://example.com/acme/logstore/pull/1128) **sandlis**: bigtable-backup: list backups just before starting deletion of wanted backups
* **Logstore** [1100](https://example.com/acme/logstore/pull/1100) **sandlis**: logging: removed some noise in logs from live-tailing
* **Logstore/build** [1089](https://example.com/acme/logstore/pull/1089) **joe-elliott**: Go 1.13
* **Logstore** [1088](https://example.com/acme/logstore/pull/1088) **pstibrany**: Updated corestore to latest master.
* **Logstore** [1085](https://example.com/acme/logstore/pull/1085) **pracucci**: Do not retry chunks transferring on shutdown in the local dev env
* **Logstore** [1084](https://example.com/acme/logstore/pull/1084) **pracucci**: Skip ingester tailer filtering if no filter is set
* **Logstore/build**[1062](https://example.com/acme/logstore/pull/1062) **joe-elliott**: dep => go mod
* **Logstore** [1049](https://example.com/acme/logstore/pull/1049) **joe-elliott**: Update logstore push path
* **Logstore** [1044](https://example.com/acme/logstore/pull/1044) **joe-elliott**: Fixed broken logql request filtering
* **Logstore/tools** [1043](https://example.com/acme/logstore/pull/1043) **sandlis**: bigtable-backup: use latest bigtable backup docker image with fix for list backups
* **Logstore** [1030](https://example.com/acme/logstore/pull/1030) **polar3130**: fix typo in error messages
* **Logstore/tools** [1028](https://example.com/acme/logstore/pull/1028) **sandlis**: bigtable-backup: verify backups to work on latest list of backups
* **Logstore** [1022](https://example.com/acme/logstore/pull/1022) **joe-elliott**: Logstore HTTP/JSON Model Layer
* **Logstore** [1016](https://example.com/acme/logstore/pull/1016) **slim-bean**: Revert "Updated stream json objects to be more parse friendly (#1010)"
* **Logstore** [1010](https://example.com/acme/logstore/pull/1010) **joe-elliott**: Updated stream json objects to be more parse friendly
* **Logstore** [1009](https://example.com/acme/logstore/pull/1009) **cyriltovena**: Make Logstore HTTP API more compatible with Prometheus
* **Logstore** [1008](https://example.com/acme/logstore/pull/1008) **wardbekker**: Improved Ingester out-of-order error for faster troubleshooting
* **Logstore** [1001](https://example.com/acme/logstore/pull/1001) **slim-bean**: Update new API paths
* **Logstore** [998](https://example.com/acme/logstore/pull/998) **sandlis**: Change unit of duration params to hours to align it with duration config at other places in Logstore
* **Logstore** [980](https://example.com/acme/logstore/pull/980) **sh0rez**: feat: configuration source precedence
* **Logstore** [948](https://example.com/acme/logstore/pull/948) **sandlis**: limits: limits implementation for logstore
* **Logstore** [947](https://example.com/acme/logstore/pull/947) **sandlis**: added a variable for storing periodic table duration as an int to be …
* **Logstore** [938](https://example.com/acme/logstore/pull/938) **sandlis**: vendoring: update corestore to latest master
* **Logstore/tools** [930](https://example.com/acme/logstore/pull/930) **sandlis**: fix incrementing of bigtable_backup_job_backups_created metric
* **Logstore/tools** [920](https://example.com/acme/logstore/pull/920) **sandlis**: bigtable-backup tool fix
* **Logstore/tools** [895](https://example.com/acme/logstore/pull/895) **sandlis**: bigtable-backup-tool: Improvements
* **Logstore** [755](https://example.com/acme/logstore/pull/755) **sandlis**: Use grpc client config from corestore for Ingester to get more control
* **Logstore** [654](https://example.com/acme/logstore/pull/654) **cyriltovena**: LogQL: Vector and Range Vector Aggregation.

### Bug Fixes

* **Logstore** [1114](https://example.com/acme/logstore/pull/1114) **rfratto**: pkg/ingester: prevent shutdowns from processing during joining handoff
* **Logstore** [1097](https://example.com/acme/logstore/pull/1097) **joe-elliott**: Reverted cloud.google.com/go to 0.44.1
* **Logstore** [986](https://example.com/acme/logstore/pull/986) **pracucci**: Fix panic in tailer due to race condition between send() and close()
* **Logstore** [975](https://example.com/acme/logstore/pull/975) **sh0rez**: fix(distributor): parseError BadRequest
* **Logstore** [944](https://example.com/acme/logstore/pull/944) **rfratto**: pkg/querier: fix concurrent access to querier tail clients

## Promtail

### Features/Improvements/Changes

* **Promtail/pipeline** [1179](https://example.com/acme/logstore/pull/1179) **pracucci**: promtail: fix handling of JMESPath expression returning nil while parsing JSON
* **Promtail/pipeline** [1123](https://example.com/acme/logstore/pull/1123) **pracucci**: promtail: added action_on_failure support to timestamp stage
* **Promtail/pipeline** [1122](https://example.com/acme/logstore/pull/1122) **pracucci**: promtail: initialize extracted map with initial labels
* **Promtail/pipeline** [1112](https://example.com/acme/logstore/pull/1112) **cyriltovena**: Add logql filter to match stages and drop capability
* **Promtail/journal** [1109](https://example.com/acme/logstore/pull/1109) **rfratto**: Clarify journal warning
* **Promtail** [1083](https://example.com/acme/logstore/pull/1083) **pracucci**: Increased promtail's backoff settings in prod and improved doc
* **Promtail** [1026](https://example.com/acme/logstore/pull/1026) **erwinvaneyk**: promtail: fix externalURL and path prefix issues
* **Promtail** [976](https://example.com/acme/logstore/pull/976) **slim-bean**: Wrap debug log statements in conditionals to save allocations
* **Promtail** [973](https://example.com/acme/logstore/pull/973) **ctrox**: tests: Set default value for BatchWait as ticker does not accept 0
* **Promtail** [969](https://example.com/acme/logstore/pull/969) **ctrox**: promtail: Use ticker instead of timer for batch wait
* **Promtail** [952](https://example.com/acme/logstore/pull/952) **pracucci**: promtail: add metrics on sent and dropped log entries
* **Promtail** [934](https://example.com/acme/logstore/pull/934) **pracucci**: promtail: do not send the last batch - to ingester - if empty
* **Promtail** [921](https://example.com/acme/logstore/pull/921) **rfratto**: promtail: add "max_age" field to configure cutoff for journal reading
* **Promtail** [883](https://example.com/acme/logstore/pull/883) **adityacs**: Add pipeline unit testing to promtail

### Bugfixes

* **Promtail** [1194](https://example.com/acme/logstore/pull/1194) **slim-bean**: Improve how we record file size metric to avoid a race in our file lagging alert
* **Promtail/journal** [1072](https://example.com/acme/logstore/pull/1072) **rfratto**: build: enable journal in promtail linux release build

## Docs

* **Docs** [1176](https://example.com/acme/logstore/pull/1176) **rfratto**: docs: add example and documentation about using JMESPath literals
* **Docs** [1139](https://example.com/acme/logstore/pull/1139) **joe-elliott**: Moved client docs and add serilog example
* **Docs** [1132](https://example.com/acme/logstore/pull/1132) **kailwallin**: FixedTypo.Update README.md
* **Docs** [1130](https://example.com/acme/logstore/pull/1130) **pracucci**: docs: fix Promtail / Logstore capitalization
* **Docs** [1129](https://example.com/acme/logstore/pull/1129) **pracucci**: docs: clarified the relation between retention period and table period
* **Docs** [1124](https://example.com/acme/logstore/pull/1124) **geowa4**: Client recommendations documentation tweaks
* **Docs** [1106](https://example.com/acme/logstore/pull/1106) **cyriltovena**: Add fluent-bit missing link in the main documentation page.
* **Docs** [1099](https://example.com/acme/logstore/pull/1099) **pracucci**: docs: improve table manager documentation
* **Docs** [1094](https://example.com/acme/logstore/pull/1094) **rfratto**: docs: update stages README with the docker and cri stages
* **Docs** [1091](https://example.com/acme/logstore/pull/1091) **daixiang0**: docs(stage): add docker and cri
* **Docs** [1077](https://example.com/acme/logstore/pull/1077) **daixiang0**: doc(fluent-bit): add missing namespace
* **Docs** [1073](https://example.com/acme/logstore/pull/1073) **flouthoc**: Re Fix Docs: PR <https://example.com/acme/logstore/pull/1053> got erased due to force push.
* **Docs** [1069](https://example.com/acme/logstore/pull/1069) **daixiang0**: doc: unify GOPATH
* **Docs** [1068](https://example.com/acme/logstore/pull/1068) **daixiang0**: doc: skip jb init when using Tanka
* **Docs** [1067](https://example.com/acme/logstore/pull/1067) **rfratto**: Fix broken links to docs in README.md
* **Docs** [1064](https://example.com/acme/logstore/pull/1064) **jonaskello**: Fix spelling of HTTP header
* **Docs** [1063](https://example.com/acme/logstore/pull/1063) **rfratto**: docs: fix deprecated warning in api.md
* **Docs** [1060](https://example.com/acme/logstore/pull/1060) **rfratto**: Add Drone CI badge to README.md
* **Docs** [1053](https://example.com/acme/logstore/pull/1053) **flouthoc**: Fix Docs: Change Imagepull policy to IfNotpresent / Add logstore-canary b…
* **Docs** [1048](https://example.com/acme/logstore/pull/1048) **wassan128**: Logstore: Fix README link
* **Docs** [1042](https://example.com/acme/logstore/pull/1042) **daixiang0**: doc(ksonnet): include ksonnet-lib
* **Docs** [1039](https://example.com/acme/logstore/pull/1039) **sh0rez**: doc(production): replace ksonnet with Tanka
* **Docs** [1036](https://example.com/acme/logstore/pull/1036) **sh0rez**: feat: -version flag
* **Docs** [1025](https://example.com/acme/logstore/pull/1025) **oddlittlebird**: Update CONTRIBUTING.md
* **Docs** [1024](https://example.com/acme/logstore/pull/1024) **oddlittlebird**: Update README.md
* **Docs** [1014](https://example.com/acme/logstore/pull/1014) **polar3130**: Fix a link to correct doc and fix a typo
* **Docs** [1006](https://example.com/acme/logstore/pull/1006) **slim-bean**: fixing lots of broken links and a few typos
* **Docs** [1005](https://example.com/acme/logstore/pull/1005) **SmilingNavern**: Fix links to correct doc
* **Docs** [1004](https://example.com/acme/logstore/pull/1004) **rfratto**: docs: fix example with pulling systemd logs
* **Docs** [1003](https://example.com/acme/logstore/pull/1003) **oddlittlebird**: Logstore: Update README.md
* **Docs** [984](https://example.com/acme/logstore/pull/984) **tomgs**: Changing "Usage" link in main readme after docs change
* **Docs** [983](https://example.com/acme/logstore/pull/983) **daixiang0**: update positions.yaml location reference
* **Docs** [982](https://example.com/acme/logstore/pull/982) **rfratto**: Documentation Rewrite
* **Docs** [961](https://example.com/acme/logstore/pull/961) **worr**: doc: Add permissions that IAM roles for Logstore need
* **Docs** [933](https://example.com/acme/logstore/pull/933) **pracucci**: doc: move promtail doc into dedicated subfolder
* **Docs** [924](https://example.com/acme/logstore/pull/924) **pracucci**: doc: promtail known failure modes
* **Docs** [910](https://example.com/acme/logstore/pull/910) **slim-bean**: docs(build): Update docs around releasing and fix bug in version updating script
* **Docs** [850](https://example.com/acme/logstore/pull/850) **sh0rez**: docs: general documentation rework

## Build

* **Build** [1157](https://example.com/acme/logstore/pull/1157) **daixiang0**: Update golint
* **Build** [1133](https://example.com/acme/logstore/pull/1133) **daixiang0**: bump up golangci to 1.20
* **Build** [1121](https://example.com/acme/logstore/pull/1121) **pracucci**: Publish logstore-canary binaries on release
* **Build** [1054](https://example.com/acme/logstore/pull/1054) **pstibrany**: Fix dep check warnings by running dep ensure
* **Build/release** [1018](https://example.com/acme/logstore/pull/1018) **slim-bean**: updating the image version for logstore-canary and adding the version increment to the release_prepare script
* **Build/CI** [997](https://example.com/acme/logstore/pull/997) **slim-bean**: full circle
* **Build/CI** [996](https://example.com/acme/logstore/pull/996) **rfratto**: ci/drone: fix deploy command by escaping double quotes in JSON body
* **Build/CI** [995](https://example.com/acme/logstore/pull/995) **slim-bean**: use the logstore-build-image for calling circle
* **Build/CI** [994](https://example.com/acme/logstore/pull/994) **slim-bean**: Also need bash for the deploy step from drone
* **Build/CI** [993](https://example.com/acme/logstore/pull/993) **slim-bean**: Add make to the alpine image used for calling the circle deploy task from drone.
* **Build/CI** [992](https://example.com/acme/logstore/pull/992) **sh0rez**: chore(packaging): fix GOPATH being overwritten
* **Build/CI** [991](https://example.com/acme/logstore/pull/991) **sh0rez**: chore(packaging): deploy from drone
* **Build/CI** [990](https://example.com/acme/logstore/pull/990) **sh0rez**: chore(ci/cd): breaking the circle
* **Build** [989](https://example.com/acme/logstore/pull/989) **sh0rez**: chore(packaging): simplify tagging
* **Build** [981](https://example.com/acme/logstore/pull/981) **sh0rez**: chore(packaging): logstore windows/amd64
* **Build** [958](https://example.com/acme/logstore/pull/958) **daixiang0**: sync release pkgs name with release note
* **Build/CI** [914](https://example.com/acme/logstore/pull/914) **rfratto**: ci: update apt-get before installing deps for rootless step
* **Build** [911](https://example.com/acme/logstore/pull/911) **daixiang0**: optimize image tag script

## Deployment

* **Ksonnet** [1023](https://example.com/acme/logstore/pull/1023) **slim-bean**: make promtail daemonset name configurable
* **Ksonnet** [1021](https://example.com/acme/logstore/pull/1021) **rfratto**: ksonnet: update memcached and memcached-exporter images
* **Ksonnet** [1020](https://example.com/acme/logstore/pull/1020) **rfratto**: ksonnet: use consistent hashing in memcached client configs
* **Ksonnet** [1017](https://example.com/acme/logstore/pull/1017) **slim-bean**: make promtail configmap name configurable
* **Ksonnet** [946](https://example.com/acme/logstore/pull/946) **rfratto**: ksonnet: remove prefix from kvstore.consul settings in logstore config
* **Ksonnet** [926](https://example.com/acme/logstore/pull/926) **slim-bean**: feat(promtail): Make cluster role configurable
<!-- -->
* **Helm** [1174](https://example.com/acme/logstore/pull/1174) **rally25rs**: logstore-stack: Add release name to prometheus service name.
* **Helm** [1152](https://example.com/acme/logstore/pull/1152) **nicr9**: docs(helm): fix broken link to acme datasource
* **Helm** [1134](https://example.com/acme/logstore/pull/1134) **minhdanh**: Helm chart: Allow additional scrape_configs to be added
* **Helm** [1111](https://example.com/acme/logstore/pull/1111) **ekarlso**: helm: Add support for passing arbitrary secrets
* **Helm** [1110](https://example.com/acme/logstore/pull/1110) **marcosnils**: Bump acme image in logstore helm chart
* **Helm** [1104](https://example.com/acme/logstore/pull/1104) **marcosnils**: <Examples>: Deploy prometheus from helm chart
* **Helm** [1058](https://example.com/acme/logstore/pull/1058) **polar3130**: Helm: Remove default value of storageClassName in logstore/logstore helm chart
* **Helm** [1056](https://example.com/acme/logstore/pull/1056) **polar3130**: Helm: Fix the reference error of logstore/logstore helm chart
* **Helm** [967](https://example.com/acme/logstore/pull/967) **makocchi-git**: helm chart: Add missing operator to promtail
* **Helm** [937](https://example.com/acme/logstore/pull/937) **minhdanh**: helm chart: Add support for additional labels and scrapeTimeout for serviceMonitors
* **Helm** [909](https://example.com/acme/logstore/pull/909) **angelbarrera92**: Feature: Add extra containers to logstore helm chart
* **Helm** [855](https://example.com/acme/logstore/pull/855) **ikeeip**: set helm chart appVersion while release
* **Helm** [675](https://example.com/acme/logstore/pull/675) **cyriltovena**: Helm default ingester config

## Logstore Canary

* **Logstore-canary** [1137](https://example.com/acme/logstore/pull/1137) **slim-bean**: Add some additional logging to the canary on queries
* **Logstore-canary** [1131](https://example.com/acme/logstore/pull/1131) **rfratto**: pkg/canary: use default HTTP client when reading from Logstore

## Logcli

* **Logcli** [1168](https://example.com/acme/logstore/pull/1168) **sh0rez**: feat(cli): order flags by categories
* **Logcli** [1115](https://example.com/acme/logstore/pull/1115) **pracucci**: logcli: introduced QueryStringBuilder utility to clean up query string encoding
* **Logcli** [1103](https://example.com/acme/logstore/pull/1103) **pracucci**: logcli: added --step support to query command
* **Logcli** [987](https://example.com/acme/logstore/pull/987) **joe-elliott**: Logcli: Add Support for New Query Path

## Tooling

* **Dashboards** [1188](https://example.com/acme/logstore/pull/1188) **joe-elliott**: Adding Operational dashboards
* **Dashboards** [1143](https://example.com/acme/logstore/pull/1143) **joe-elliott**: Improved compression ratio histogram
* **Dashboards** [1126](https://example.com/acme/logstore/pull/1126) **joe-elliott**: Fix Logstore Chunks Dashboard
* **Tools** [1108](https://example.com/acme/logstore/pull/1108) **joe-elliott**: Updated push path to current prod

## Plugins

* **DockerDriver** [972](https://example.com/acme/logstore/pull/972) **cyriltovena**: Add stream label to docker driver
* **DockerDriver** [971](https://example.com/acme/logstore/pull/971) **cyriltovena**: Allow to pass max-size and max-file to the docker driver
* **DockerDriver** [970](https://example.com/acme/logstore/pull/970) **mindfl**: docker-driver compose labels support
<!-- -->
* **Fluentd** [928](https://example.com/acme/logstore/pull/928) **candlerb**: fluent-plugin-acme-logstore: Escape double-quotes in labels, and suppress labels with value nil
<!-- -->
* **Fluent Bit** [1155](https://example.com/acme/logstore/pull/1155) **cyriltovena**: rollback fluent-bit push path until we release 0.4
* **Fluent Bit** [1096](https://example.com/acme/logstore/pull/1096) **JensErat**: fluent-bit: edge case tests
* **Fluent Bit** [847](https://example.com/acme/logstore/pull/847) **cosmo0920**: fluent-bit shared object go plugin

## Misc

Logstore is now using a Bot to help keep issues and PR's pruned based on age/relevancy.  Please don't hesitate to comment on an issue or PR that you think was closed by the stale-bot which you think should remain open!!

* **Github** [965](https://example.com/acme/logstore/pull/965) **rfratto**: Change label used to keep issues from being marked as stale to keepalive
* **Github** [964](https://example.com/acme/logstore/pull/964) **rfratto**: Add probot-stale configuration to close stale issues.

# 0.3.0 (2019-08-16)

### Features/Enhancements

* **Logstore** [877](https://example.com/acme/logstore/pull/877) **pracucci**: logstore: Improve Tailer loop
* **Logstore** [870](https://example.com/acme/logstore/pull/870) **sandlis**: bigtable-backup: update docker image for bigtable-backup tool
* **Logstore** [862](https://example.com/acme/logstore/pull/862) **sandlis**: live-tailing: preload all the historic entries before query context is cancelled
* **Logstore** [858](https://example.com/acme/logstore/pull/858) **pracucci**: logstore: removed unused TestGZIPCompression
* **Logstore** [854](https://example.com/acme/logstore/pull/854) **adityacs**: Readiness probe for querier
* **Logstore** [851](https://example.com/acme/logstore/pull/851) **cyriltovena**: Add readiness probe to distributor deployment.
* **Logstore** [894](https://example.com/acme/logstore/pull/894) **rfratto**: ksonnet: update ingester config to transfer chunks on rollout
<!-- -->
* **Build** [901](https://example.com/acme/logstore/pull/901) **sh0rez**: chore(packaging): set tag length to 7
* **Build** [900](https://example.com/acme/logstore/pull/900) **sh0rez**: chore(ci/cd): fix acmesaur credentials and CircleCI image build
* **Build** [891](https://example.com/acme/logstore/pull/891) **sh0rez**: chore(ci/cd): build containers using drone.io
* **Build** [888](https://example.com/acme/logstore/pull/888) **rfratto**: Makefile: disable building promtail with systemd support on non-amd64 platforms
* **Build** [887](https://example.com/acme/logstore/pull/887) **slim-bean**: chore(packaging): Dockerfile make avoid containers
* **Build** [886](https://example.com/acme/logstore/pull/886) **sh0rez**: chore(packaging): wrong executable format
* **Build** [855](https://example.com/acme/logstore/pull/855) **ikeeip**: set helm chart appVersion while release
<!-- -->
* **Promtail** [856](https://example.com/acme/logstore/pull/856) **martinbaillie**: promtail: Add ServiceMonitor and headless Service
* **Promtail** [809](https://example.com/acme/logstore/pull/809) **rfratto**: Makefile: build promtail with CGO_ENABLED if GOHOSTOS=GOOS=linux
* **Promtail** [730](https://example.com/acme/logstore/pull/730) **rfratto**: promtail: Add systemd journal support

> 809, 730 NOTE: Systemd journal support is currently limited to amd64 images, arm support should come in the future when the transition to building the arm image and binaries is done natively via an arm container
<!-- -->
* **Docs** [896](https://example.com/acme/logstore/pull/896) **dalance**: docs: fix link format
* **Docs** [876](https://example.com/acme/logstore/pull/876) **BouchaaraAdil**: update Docs: update Retention section on Operations doc file
* **Docs** [864](https://example.com/acme/logstore/pull/864) **temal-**: docs: Replace old values in operations.md
* **Docs** [853](https://example.com/acme/logstore/pull/853) **cyriltovena**: Add governance documentation
<!-- -->
* **Deployment** [874](https://example.com/acme/logstore/pull/874) **slim-bean**: make our ksonnet a little more modular by parameterizing the chunk and index stores
* **Deployment** [857](https://example.com/acme/logstore/pull/857) **slim-bean**: Reorder relabeling rules to prevent pod label from overwriting config define labels

> 857 POSSIBLY BREAKING: If you relied on a custom pod label to overwrite one of the labels configured by the other sections of the scrape config: `job`, `namespace`, `instance`, `container_name` and/or `__path__`, this will no longer happen, the custom pod labels are now loaded first and will be overwritten by any of these listed labels.

### Fixes

* **Logstore** [897](https://example.com/acme/logstore/pull/897) **pracucci**: Fix panic in tailer when an ingester is removed from the ring while tailing
* **Logstore** [880](https://example.com/acme/logstore/pull/880) **cyriltovena**: fix a bug where nil line buffer would be put back
* **Logstore** [859](https://example.com/acme/logstore/pull/859) **pracucci**: logstore: Fixed out of order entries allowed in a chunk on edge case
<!-- -->
* **Promtail** [893](https://example.com/acme/logstore/pull/893) **rfratto**: pkg/promtail/positions: remove executable bit from positions file
<!-- -->
* **Deployment** [867](https://example.com/acme/logstore/pull/867) **slim-bean**: Update read dashboard to include only query and label query routes
* **Deployment** [865](https://example.com/acme/logstore/pull/865) **sandlis**: fix broken jsonnet for querier
<!-- -->
* **Canary** [889](https://example.com/acme/logstore/pull/889) **slim-bean**: fix(canary): Fix Flaky Tests
<!-- -->
* **Pipeline** [869](https://example.com/acme/logstore/pull/869) **jojohappy**: Pipeline: Fixed labels process test with same objects
<!-- -->
* **Logcli** [863](https://example.com/acme/logstore/pull/863) **adityacs**: Fix Nolabels parse metrics

# 0.2.0 (2019-08-02)

There were over 100 PR's merged since 0.1.0 was released, here's a highlight:

### Features / Enhancements

* **Logstore**:  [521](https://example.com/acme/logstore/pull/521) Query label values and names are now fetched from the store.
* **Logstore**:  [541](https://example.com/acme/logstore/pull/541) Improvements in live tailing of logs.
* **Logstore**: [713](https://example.com/acme/logstore/pull/713) Storage memory improvement.
* **Logstore**: [764](https://example.com/acme/logstore/pull/764) Tailing can fetch previous logs for context.
* **Logstore**: [782](https://example.com/acme/logstore/pull/782) Performance improvement: Query storage by iterating through chunks in batches.
* **Logstore**: [788](https://example.com/acme/logstore/pull/788) Querier timeouts.
* **Logstore**: [794](https://example.com/acme/logstore/pull/794) Support ingester chunk transfer on shutdown.
* **Logstore**: [729](https://example.com/acme/logstore/pull/729) Bigtable backup tool support.
<!-- -->
* **Pipeline**: [738](https://example.com/acme/logstore/pull/738) Added a template stage for manipulating label values.
* **Pipeline**: [732](https://example.com/acme/logstore/pull/732) Support for Unix timestamps.
* **Pipeline**: [760](https://example.com/acme/logstore/pull/760) Support timestamps without year.
<!-- -->
* **Helm**:  [641](https://example.com/acme/logstore/pull/641) Helm integration testing.
* **Helm**: [824](https://example.com/acme/logstore/pull/824) Add service monitor.
* **Helm**: [830](https://example.com/acme/logstore/pull/830) Customize namespace.
<!-- -->
* **Docker-Plugin**: [663](https://example.com/acme/logstore/pull/663) Created a Docker logging driver plugin.
<!-- -->
* **Fluent-Plugin**: [669](https://example.com/acme/logstore/pull/669) Ability to specify keys to remove.
* **Fluent-Plugin**: [709](https://example.com/acme/logstore/pull/709) Multi-worker support.
* **Fluent-Plugin**: [792](https://example.com/acme/logstore/pull/792) Add prometheus for metrics and update gems.
<!-- -->
* **Build**: [668](https://example.com/acme/logstore/pull/668),[762](https://example.com/acme/logstore/pull/762) Build multiple architecture containers.
<!-- -->
* **Logstore-Canary**: [772](https://example.com/acme/logstore/pull/772) Moved into Logstore project.

### Bugfixes

There were many fixes, here are a few of the most important:

* **Promtail**: [650](https://example.com/acme/logstore/pull/650) Build on windows.
* **Fluent-Plugin**: [667](https://example.com/acme/logstore/pull/667) Rename fluent plugin.
* **Docker-Plugin**: [813](https://example.com/acme/logstore/pull/813) Fix panic for newer docker version (18.09.7+).

# 0.1.0 (2019-06-03)

First (beta) Release!
