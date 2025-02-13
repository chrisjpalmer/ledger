run-local:
	go run ./cmd/backend

postgres-migrate:
	dagger call postgres-migrate --src=$${PWD} up
.PHONY: postgres-migrate

postgres:
	dagger call postgres up
.PHONY: postgres

migrate:
	dagger call migrate --src=$${PWD} --svc=tcp://localhost:5432
.PHONY: migrate

openapi-generate:
	rm -rf ./internal/api
	mkdir -p ./internal/api
	dagger call openapi-generate --src=. export --path=$${PWD}/internal/api
.PHONY: openapi-generate