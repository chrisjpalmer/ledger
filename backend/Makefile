check-pull-request:
	dagger call check-pull-request --src=$${PWD}
.PHONY: check-pull-request

run-local:
	go run ./cmd/backend
.PHONY: run-local

run:
	dagger call ledger --src=$${PWD} --postgres=tcp://localhost:5432 up
.PHONY: run

integration-local:
	go test -count=1 ./integration
.PHONY: integration-local

integration:
	dagger call integration --src=$${PWD} --ledger=tcp://localhost:8080
.PHONY: integration

postgres-test-local:
	go test -count=1 ./internal/postgres
.PHONY: postgres-test-local

postgres-test:
	dagger call postgres-test --src=$${PWD}
.PHONY: postgres-test

postgres-migrate:
	dagger call postgres-migrate --src=$${PWD} up
.PHONY: postgres-migrate

postgres:
	dagger call postgres up
.PHONY: postgres

migrate:
	dagger call migrate --src=$${PWD} --svc=tcp://localhost:5432
.PHONY: migrate

openapi-drift:
	dagger call openapi-drift --src=$${PWD}
.PHONY: openapi-drift

openapi-generate:
	rm -rf ./internal/api
	mkdir -p ./internal/api
	dagger call openapi-generate --src=$${PWD} export --path=$${PWD}/internal/api

.PHONY: openapi-generate

psql:
	dagger call psql --svc=tcp://localhost:5432
.PHONY: psql
