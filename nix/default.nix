{ self, pkgs, lib }:
let
  # self.rev is only set on a clean git tree
  gitRevision = if (self ? rev) then self.rev else "dirty";
  shortGitRevsion = with lib;
    if (self ? rev) then
      (strings.concatStrings
        (lists.take 8 (strings.stringToCharacters gitRevision)))
    else
      "dirty";

  # the image tag script is hard coded to take only 7 characters
  imageTagVersion = with lib;
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

  meta = with lib; {
    homepage = "https://acme.com/oss/logstore/";
    changelog = "https://example.com/acme/logstore/commit/${shortGitRevsion}";
    maintainers = with maintainers; [ trevorwhitney ];

  };

  logstore-helm-test = pkgs.callPackage ../production/helm/logstore/src/helm-test {
    inherit pkgs;
    inherit (pkgs) lib buildGoModule dockerTools;
    rev = gitRevision;
  };
in
{
  inherit (logstore-helm-test) logstore-helm-test logstore-helm-test-docker;
} // rec {
  logstore = pkgs.callPackage ./packages/logstore.nix {
    inherit imageTag pkgs;
    version = shortGitRevsion;
  };

  logcli = logstore.overrideAttrs (oldAttrs: {
    pname = "logcli";

    subPackages = [ "cmd/logcli" ];

     tags = [
        "slicelabels"
     ];

    meta = with lib; {
      description = "LogCLI is a command line tool for interacting with Logstore.";
      mainProgram = "logcli";
      license = with licenses; [ agpl3Only ];
    } // meta;
  });

  logstore-canary = logstore.overrideAttrs (oldAttrs: {
    pname = "logstore-canary";

    subPackages = [ "cmd/logstore-canary" ];

     tags = [
        "slicelabels"
     ];

    meta = with lib; {
      description = "Logstore Canary is a canary for the Logstore project.";
      mainProgram = "logstore-canary";
      license = with licenses; [ agpl3Only ];
    } // meta;
  });

  promtail = logstore.overrideAttrs (oldAttrs: {
    pname = "promtail";

    buildInputs = with pkgs; lib.optionals stdenv.hostPlatform.isLinux [ systemd.dev ];

    tags = [
        "promtail_journal_enabled"
        "slicelabels"
    ];

    subPackages = [ "clients/cmd/promtail" ];

    preFixup = lib.optionalString pkgs.stdenv.hostPlatform.isLinux ''
      wrapProgram $out/bin/promtail \
        --prefix LD_LIBRARY_PATH : "${lib.getLib pkgs.systemd}/lib"
    '';

    meta = with lib; {
      description = "Client for sending logs to Logstore";
      mainProgram = "promtail";
      license = with licenses; [ asl20 ];
    } // meta;
  });
}
