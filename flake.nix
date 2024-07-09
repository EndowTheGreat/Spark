{
  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };
  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        packages.default = pkgs.stdenv.mkDerivation {
          pname = "Spark";
          version = "0.1.0";
          src = pkgs.lib.cleanSource ./.;
          buildInputs = with pkgs; [ go gopls gotools go-tools ];
          buildPhase = ''
            export HOME=$TMPDIR
            go build -o spark cmd/spark/main.go
          '';
          installPhase = ''
            mkdir -p $out/bin
            cp spark $out/bin/
          '';
        };
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [ go gopls gotools go-tools ];
          shellHook = ''
            go run ./...
          '';
        };
      });
}
