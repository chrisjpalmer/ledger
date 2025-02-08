openapi-generate:
	rm -rf ./internal/api
	mkdir -p ./internal/api
	dagger call openapi-generate --src=. export --path=$${PWD}/internal/api
.PHONY: openapi-generate