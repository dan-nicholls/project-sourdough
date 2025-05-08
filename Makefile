DIR ?= tmp/

.PHONY: build dev

build:
	@echo "Building to $(DIR)"
	tailwindcss -i internal/web/static/input.css -o static/css/tw.css
	templ generate ./internal/web/templates
	go build -o ${DIR}/server  cmd/server/main.go

dev:
	air
