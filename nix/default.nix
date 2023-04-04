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

      logstore = prev.callPackage ./logstore.nix {
        inherit imageTag;
        version = shortGitRevsion;
        pkgs = prev;
      };

      faillint = prev.callPackage ./faillint.nix {
        inherit (prev) lib buildGoModule fetchFromGitHub;
      };

      chart-releaser = prev.callPackage ./chart-releaser.nix {
        inherit (prev) pkgs lib buildGoModule fetchFromGitHub;
      };
    };
}
