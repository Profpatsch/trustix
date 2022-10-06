{ flakeInputs ? import ./flake-fetch.nix
, system ? builtins.currentSystem
}:

let
  inherit (flakeInputs) nixpkgs gomod2nix npmlock2nix gitignore nix-eval-jobs;
in

import nixpkgs {
  inherit system;
  config.allowAliases = false;
  overlays = [
    (import "${gomod2nix}/overlay.nix")

    (final: prev: (import "${gitignore}" { inherit (final) lib; }))

    (final: prev: {
      # Prevent the entirety of hydra to be in $PATH/runtime closure
      # We only want the evaluator
      hydra-eval-jobs = prev.runCommand "hydra-eval-jobs-${prev.hydra_unstable.version}" { } ''
        mkdir -p $out/bin
        cp -s ${prev.hydra_unstable}/bin/hydra-eval-jobs $out/bin/
      '';
    })

    (final: prev: {
      npmlock2nix = import npmlock2nix { pkgs = final; };
    })

    (final: prev: {
      nix-eval-jobs = final.callPackage nix-eval-jobs { };
    })

    (final: prev:
      let
        inherit (prev) lib;
        dirNames = lib.attrNames (lib.filterAttrs (pkgDir: type: type == "directory" && builtins.pathExists (./packages + "/${pkgDir}/default.nix")) (builtins.readDir ./packages));
      in
      builtins.listToAttrs (map
        (pkgDir: {
          value = final.callPackage (./packages + "/${pkgDir}") { };
          name = pkgDir;
        })
        dirNames)
    )
  ];
}
