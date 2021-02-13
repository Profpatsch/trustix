let
  pkgs = import <nixpkgs> { overlays = import ../../nix/overlays.nix; };
  inherit (pkgs) poetry2nix;

  # Prevent the entirety of hydra to be in $PATH/runtime closure
  # We only want the evaluator
  hydra-eval-jobs = pkgs.runCommand "hydra-eval-jobs-${pkgs.hydra-unstable.version}" { } ''
    mkdir -p $out/bin
    cp -s ${pkgs.hydra-unstable}/bin/hydra-eval-jobs $out/bin/
  '';

in poetry2nix.mkPoetryApplication {
  projectDir = ./.;

  propagatedBuildInputs = [ hydra-eval-jobs ];

  # Don't propagate anything, hydra-eval-jobs is already wrapped in $PATH
  postFixup = "rm $out/nix-support/propagated-build-inputs";

  overrides = poetry2nix.overrides.withDefaults (
    import ./overrides.nix { inherit pkgs; }
  );
}