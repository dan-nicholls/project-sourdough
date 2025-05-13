# 🍞 Project Sourdough

**Project Sourdough** is a small web app I decided to create while i was busy
persuing a new hobby, baking sourdough. I had several requests come in and
wanted to give my friends and family a space where they could make a request /
order easily.

Built with:

- [Go](https://go.dev/) — the core backend framework
- [Templ](https://templ.guide/) — HTML templating components for Go
- [Tailwind CSS](https://tailwindcss.com/) — styling
- [HTMX](https://htmx.org/) — frontend interactivity
- [SQLite3](https://sqlite.org/) — DB

## Requirements

- Go 1.21+
- PostgreSQL (installed and running)
- `templ` CLI (`go install github.com/a-h/templ/cmd/templ@latest`)
- `tailwindcss` (`npm install --save-dev tailwindcss`)
- `Make` - Available in `build-essentials` (ubuntu) or `GNUMake` (nixpkgs)
- (optional) `air` (`go install github.com/cosmtrek/air@latest`)

OR

- `nix` - To make use of `nix-shell` for a reproducable dev environment

```bash
nix-shell #In the same directory
```

## Development

To start the dev environment with hot-reloading:

```bash
make dev
```

This will:

- Regenerate `*_templ.go` files using templ
- Rebuild CSS using `tailwindcss`
- Run the GO server (built in `tmp/server` by default)
- Detect changes to `.go` and `.templ` files

The standard server will be accessable at `http://localhost:3000`, but
hot-reloading (rebuild on server changes & force browser refresh) will be
available at `http://localhost:3030`.
