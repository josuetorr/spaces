SERVER_APP_NAME = server

.PHONY: tailwind-watch
tailwind-watch:
	pnpm run tw:watch

.PHONY: tailwind-build
tailwind-build:
	pnpm run tw:build

.PHONY: templ-watch
templ-watch:
	templ generate --watch

.PHONY: templ-generate
templ-generate:
	templ generate
	
.PHONY: dev
dev:
	go build -o ./tmp/main ./cmd/server/main.go && air

.PHONY: build
build-server:
	make tailwind-build
	make templ-generate
	go build -ldflags "-X main.Environment=production" -o ./bin/$(SERVER_APP_NAME) ./cmd/main.go

