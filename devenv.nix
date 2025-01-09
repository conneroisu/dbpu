{
  pkgs,
  # lib,
  # config,
  inputs,
  ...
}: let
  unstable-pkgs = import inputs.nixpkgs-unstable {inherit (pkgs) system;};
in {
  name = "dbpu";
  devcontainer.enable = true;
  languages = {
    nix.enable = true;
    go = {
      enable = true;
      package = unstable-pkgs.go;
    };
  };

  git-hooks = {
    hooks = {
      golangci-lint.enable = true;
    };
  };

  packages = with pkgs; [
    sqldiff
    pprof
    podman
    revive
    esbuild
    golangci-lint
    golangci-lint-langserver
    gomarkdoc
    gotests
    gotools
    templ
    sqlc
    flyctl
    air
    wireguard-tools
  ];

  scripts = {
    generate.exec = ''
      go generate -v ./...
    '';
    tests.exec = ''
      go test -v -short ./...
    '';
    unit-tests.exec = ''
      go test -v ./...
    '';
    lint.exec = ''
      golangci-lint run
    '';
    dx.exec = ''
      $EDITOR $(git rev-parse --show-toplevel)/devenv.nix
    '';
  };

  enterShell = ''
    git status
  '';
  enterTest = ''
    echo "Running tests"
    git --version | grep --color=auto "${pkgs.git.version}"
  '';

  cachix.enable = true;
}
