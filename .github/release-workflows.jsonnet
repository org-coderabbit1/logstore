local logstoreRelease = import 'workflows/main.jsonnet';
local build = logstoreRelease.build;

local releaseLibRef = 'main';

local checkTemplate = 'acme/logstore-release/.github/workflows/check.yml@%s' % releaseLibRef;

local imageJobs = {
  logstore: build.image('logstore', 'cmd/logstore'),
  fluentd: build.image('fluent-plugin-logstore', 'clients/cmd/fluentd', platform=['linux/amd64']),
  'fluent-bit': build.image('fluent-bit-plugin-logstore', 'clients/cmd/fluent-bit', platform=['linux/amd64']),
  logstash: build.image('logstash-output-logstore', 'clients/cmd/logstash', platform=['linux/amd64']),
  logcli: build.image('logcli', 'cmd/logcli'),
  'logstore-canary': build.image('logstore-canary', 'cmd/logstore-canary'),
  'logstore-canary-boringcrypto': build.image('logstore-canary-boringcrypto', 'cmd/logstore-canary-boringcrypto'),
  promtail: build.image('promtail', 'clients/cmd/promtail'),
  querytee: build.image('logstore-query-tee', 'cmd/querytee', platform=['linux/amd64']),
};

local weeklyImageJobs = {
  logstore: build.weeklyImage('logstore', 'cmd/logstore'),
  fluentd: build.weeklyImage('fluent-plugin-logstore', 'clients/cmd/fluentd', platform=['linux/amd64']),
  'fluent-bit': build.weeklyImage('fluent-bit-plugin-logstore', 'clients/cmd/fluent-bit', platform=['linux/amd64']),
  logstash: build.weeklyImage('logstash-output-logstore', 'clients/cmd/logstash', platform=['linux/amd64']),
  logcli: build.weeklyImage('logcli', 'cmd/logcli'),
  'logstore-canary': build.weeklyImage('logstore-canary', 'cmd/logstore-canary'),
  'logstore-canary-boringcrypto': build.weeklyImage('logstore-canary-boringcrypto', 'cmd/logstore-canary-boringcrypto'),
  promtail: build.weeklyImage('promtail', 'clients/cmd/promtail'),
  querytee: build.weeklyImage('logstore-query-tee', 'cmd/querytee', platform=['linux/amd64']),
};

local buildImageVersion = std.extVar('BUILD_IMAGE_VERSION');
local buildImage = 'acme/logstore-build-image:%s' % buildImageVersion;
local golangCiLintVersion = 'v1.60.3';

local imageBuildTimeoutMin = 60;
local imagePrefix = 'acme';

{
  'patch-release-pr.yml': std.manifestYamlDoc(
    logstoreRelease.releasePRWorkflow(
      branches=['release-[0-9]+.[0-9]+.x'],
      buildImage=buildImage,
      checkTemplate=checkTemplate,
      golangCiLintVersion=golangCiLintVersion,
      imageBuildTimeoutMin=imageBuildTimeoutMin,
      imageJobs=imageJobs,
      imagePrefix=imagePrefix,
      releaseLibRef=releaseLibRef,
      releaseRepo='acme/logstore',
      skipArm=false,
      skipValidation=false,
      useGitHubAppToken=true,
      versioningStrategy='always-bump-patch',
    ) + {
      name: 'Prepare Patch Release PR',
    }, false, false
  ),
  'minor-release-pr.yml': std.manifestYamlDoc(
    logstoreRelease.releasePRWorkflow(
      branches=['k[0-9]+'],
      buildImage=buildImage,
      checkTemplate=checkTemplate,
      golangCiLintVersion=golangCiLintVersion,
      imageBuildTimeoutMin=imageBuildTimeoutMin,
      imageJobs=imageJobs,
      imagePrefix=imagePrefix,
      releaseLibRef=releaseLibRef,
      releaseRepo='acme/logstore',
      skipArm=false,
      skipValidation=false,
      useGitHubAppToken=true,
      versioningStrategy='always-bump-minor',
    ) + {
      name: 'Prepare Minor Release PR from Weekly',
    }, false, false
  ),
  'release.yml': std.manifestYamlDoc(
    logstoreRelease.releaseWorkflow(
      branches=['release-[0-9]+.[0-9]+.x', 'k[0-9]+', 'main'],
      getDockerCredsFromVault=true,
      imagePrefix='acme',
      releaseLibRef=releaseLibRef,
      releaseRepo='acme/logstore',
      useGitHubAppToken=true,
    ), false, false
  ),
  'check.yml': std.manifestYamlDoc({
    name: 'check',
    on: {
      pull_request: {},
      push: {
        branches: ['main'],
      },
    },
    jobs: {
      check: {
        uses: checkTemplate,
        with: {
          build_image: buildImage,
          golang_ci_lint_version: golangCiLintVersion,
          release_lib_ref: releaseLibRef,
          skip_validation: false,
          use_github_app_token: true,
        },
      },
    },
  }),
  'images.yml': std.manifestYamlDoc({
    name: 'publish images',
    on: {
      push: {
        branches: [
          'k[0-9]+*',  // This is a weird glob pattern, not a regexp, do not use ".*", see https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#filter-pattern-cheat-sheet
          'main',
        ],
      },
    },
    permissions: {
      'id-token': 'write',
      contents: 'write',
      'pull-requests': 'write',
    },
    jobs: {
      check: {
        uses: checkTemplate,
        with: {
          build_image: buildImage,
          golang_ci_lint_version: golangCiLintVersion,
          release_lib_ref: releaseLibRef,
          skip_validation: false,
          use_github_app_token: true,
        },
      },
    } + std.mapWithKey(function(name, job)
      job
      + logstoreRelease.job.withNeeds(['check'])
      + {
        env: {
          BUILD_TIMEOUT: imageBuildTimeoutMin,
          RELEASE_REPO: 'acme/logstore',
          RELEASE_LIB_REF: releaseLibRef,
          IMAGE_PREFIX: imagePrefix,
        },
      }, weeklyImageJobs),
  }),
}
