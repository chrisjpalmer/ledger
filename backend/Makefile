check-pull-request:
	dagger call check-pull-request --src=$${PWD}

run-local:
	go run ./cmd/backend

run:
	dagger call ledger --src=$${PWD} --postgres=tcp://localhost:5432 up

integration-local:
	go test -count=1 ./test

integration:
	dagger call integration --src=$${PWD} --ledger=tcp://localhost:8080

postgres-migrate:
	dagger call postgres-migrate --src=$${PWD} up
.PHONY: postgres-migrate

postgres:
	dagger call postgres up
.PHONY: postgres

psql:
	dagger call psql --svc=tcp://localhost:5432

migrate:
	dagger call migrate --src=$${PWD} --svc=tcp://localhost:5432
.PHONY: migrate

openapi-generate:
	rm -rf ./internal/api
	mkdir -p ./internal/api
	dagger call openapi-generate --src=. export --path=$${PWD}/internal/api
.PHONY: openapi-generate