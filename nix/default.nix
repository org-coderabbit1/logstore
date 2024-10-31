{ self }:
{
  overlay = final: prev:
    let
      # self.rev is only set on a clean git tree
      gitRevision = if (self ? rev) then self.rev else "dirty";
      shortGitRevsion = with prev.lib;
        if (self ? rev) then
          (strings.concatStrings
            (lists.take 8 (strings.stringToCharacters gitRevision)))
        else
          "dirty";

      # the image tag script is hard coded to take only 7 characters
      imageTagVersion = with prev.lib;
        if (self ? rev) then
          (strings.concatStrings
            (lists.take 8 (strings.stringToCharacters gitRevision)))
        else
          "dirty";

      imageTag =
        if (self ? rev) then
          "${imageTagVersion}"
        else
          "${imageTagVersion}-WIP";

      logstore-helm-test = prev.callPackage ../production/helm/logstore/src/helm-test {
        inherit (prev) pkgs lib buildGoModule dockerTools;
        rev = gitRevision;
      };
    in
    {
      inherit (logstore-helm-test) logstore-helm-test logstore-helm-test-docker;
    } // rec {
      logstore = prev.callPackage ./packages/logstore.nix {
        inherit imageTag;
        version = shortGitRevsion;
        pkgs = prev;
      };

      logcli = logstore.overrideAttrs (oldAttrs: {
        pname = "logcli";

        buildPhase = ''
          export GOCACHE=$TMPDIR/go-cache
          make clean logcli
        '';

        installPhase = ''
          mkdir -p $out/bin
          install -m755 cmd/logcli/logcli $out/bin/logcli
        '';
      });

      logstore-canary = logstore.overrideAttrs (oldAttrs: {
        pname = "logstore-canary";

        buildPhase = ''
          export GOCACHE=$TMPDIR/go-cache
          make clean logstore-canary
        '';

        installPhase = ''
          mkdir -p $out/bin
          install -m755 cmd/logstore-canary/logstore-canary $out/bin/logstore-canary
        '';
      });

      promtail = logstore.overrideAttrs (oldAttrs: {
        pname = "promtail";

        buildInputs =
          let
            inherit (oldAttrs) buildInputs;
          in
          if prev.stdenv.hostPlatform.isLinux then
            buildInputs ++ (with prev; [ systemd ])
          else buildInputs;

        buildPhase = ''
          export GOCACHE=$TMPDIR/go-cache
          make clean promtail
        '';

        installPhase = ''
          mkdir -p $out/bin
          install -m755 clients/cmd/promtail/promtail $out/bin/promtail
        '';
      });
    };
}
