.PHONY: sqlc-generate
sqlc-generate:
	sqlc generate -f=./internal/store/sqlc.json

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: tailwind-build
tailwind-build:
	tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify

.PHONY: build
build:
	make sqlc-generate
	make tailwind-build
	make templ-generate
	go build -o tmp/main ./cmd

.PHONY: dev
dev:
	make build
	air