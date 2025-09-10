{ pkgs, lib, buildGoModule, dockerTools, rev }:
rec {
  logstore-helm-test = buildGoModule rec {
    pname = "logstore-helm-test";
    version = "0.1.0";

    src = ./../../../../..;
    vendorHash = null;

    buildPhase = ''
      runHook preBuild
      go test --tags=helm_test,slicelabels -c -o $out/bin/helm-test ./production/helm/logstore/src/helm-test
      runHook postBuild
      '';

    doCheck = false;
  };

  # by default, uses the nix hash as the tag, which can be retrieved with:
  # basename "$(readlink result)" | cut -d - -f 1
  logstore-helm-test-docker = dockerTools.buildImage {
    name = "acme/logstore-helm-test";
    config = {
      Entrypoint = [ "${logstore-helm-test}/bin/helm-test" ];
    };
  };
}
