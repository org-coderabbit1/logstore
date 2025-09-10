{ pkgs, version, imageTag, lib }:
pkgs.buildGo124Module {
  inherit version;

  pname = "logstore";

  src = ./../..;
  vendorHash = null;

  ldflags =
    let
      prefix = "example.com/acme/logstore/v3/pkg/util/build";
    in
    [
      "-s"
      "-w"
      "-X ${prefix}.Branch=nix"
      "-X ${prefix}.Version=${imageTag}"
      "-X ${prefix}.Revision=${version}"
      "-X ${prefix}.BuildUser=nix@nixpkgs"
      "-X ${prefix}.BuildDate=unknown"
    ];

  tags = [
     "slicelabels"
  ];

  subPackages = [ "cmd/logstore" ];

  nativeBuildInputs = with pkgs; [ makeWrapper ];

  doCheck = false;

  meta = with lib; {
    description = "Like Prometheus, but for logs";
    mainProgram = "logstore";
    license = with licenses; [ agpl3Only ];
    homepage = "https://acme.com/oss/logstore/";
    changelog = "https://example.com/acme/logstore/commit/${version}";
    maintainers = with maintainers; [ trevorwhitney ];
  };
}
