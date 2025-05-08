{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  packages = with pkgs; [
    go
    templ
    air
    git
    nodejs
    tailwindcss
	gnumake
  ];

  shellHook = ''
    echo "✅ Goth stack dev shell ready (Go + Templ + TailwindCSS)"
  '';
}

