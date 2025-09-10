local logstoreRelease = import 'main.jsonnet';
local build = logstoreRelease.build;


local buildImage = 'golang:1.24';
local dockerPluginDir = 'clients/cmd/docker-driver';

{
  '.github/workflows/release-pr.yml': std.manifestYamlDoc(
    logstoreRelease.releasePRWorkflow(
      imageJobs={
        logstore: build.image('fake-logstore', 'cmd/logstore'),
        'logstore-docker-driver': build.dockerPlugin('logstore-docker-driver', dockerPluginDir, buildImage=buildImage),
      },
      buildImage=buildImage,
      buildArtifactsBucket='logstore-build-artifacts',
      branches=['release-[0-9]+.[0-9]+.x'],
      imagePrefix='trevorwhitney075',
      releaseLibRef='main',
      releaseRepo='acme/logstore-release',
      skipValidation=false,
      versioningStrategy='always-bump-patch',
    ) + {
      name: 'Create Release PR',
    }, false, false
  ),
  '.github/workflows/test-release-pr.yml': std.manifestYamlDoc(
    logstoreRelease.releasePRWorkflow(
      imageJobs={
        logstore: build.image('fake-logstore', 'cmd/logstore'),
        'logstore-docker-driver': build.dockerPlugin('logstore-docker-driver', dockerPluginDir, buildImage=buildImage),
      },
      buildImage=buildImage,
      buildArtifactsBucket='logstore-build-artifacts',
      branches=['release-[0-9]+.[0-9]+.x'],
      dryRun=true,
      imagePrefix='trevorwhitney075',
      releaseLibRef='main',
      releaseRepo='acme/logstore-release',
      skipValidation=false,
      versioningStrategy='always-bump-patch',
    ) + {
      name: 'Test Create Release PR Action',
      on+: {
        pull_request: {},
      },
    }, false, false
  ),
  '.github/workflows/release.yml': std.manifestYamlDoc(
    logstoreRelease.releaseWorkflow(
      branches=['release-[0-9]+.[0-9]+.x'],
      buildArtifactsBucket='logstore-build-artifacts',
      dockerUsername='trevorwhitney075',
      getDockerCredsFromVault=false,
      imagePrefix='trevorwhitney075',
      pluginBuildDir=dockerPluginDir,
      releaseLibRef='main',
      releaseRepo='acme/logstore-release',
      useGitHubAppToken=true,
    ) + {
      name: 'Create Release',
      on+: {
        pull_request: {},
      },
    }, false, false
  ),
  '.github/workflows/check.yml': std.manifestYamlDoc(
    logstoreRelease.check
  ),
  '.github/workflows/gel-check.yml': std.manifestYamlDoc(
    logstoreRelease.checkGel
  ),
}
