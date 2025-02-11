run-local:
	go run ./cmd/backend

postgres:
	dagger call postgres up
.PHONY: postgres

openapi-generate:
	rm -rf ./internal/api
	mkdir -p ./internal/api
	dagger call openapi-generate --src=. export --path=$${PWD}/internal/api
.PHONY: openapi-generate