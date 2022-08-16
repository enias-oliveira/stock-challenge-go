{ pkgs ? import <nixpkgs> { } }:

with pkgs;

mkShell {
  buildInputs = [
    docker
    docker-compose
    go_1_18
    gopls
    gomodifytags
    gotests
    gore
    emacs28Packages.guru-mode
    emacs28Packages.go-guru
  ];

}
