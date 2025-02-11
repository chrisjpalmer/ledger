# Ledger

A simple monthly ledger to help one manage their expenses

## Requirements

- GNU Make
- Docker
- Dagger
- Go (1.21 or higher)

## Getting Started

All commands assume you are in the `/backend` directory:

```sh
cd ./backend
```

Run each of these in a different shell:

```sh
make postgres

make run-local
```

## Contributing

### Open API

To add new resources to the spec, simply alter the open api spec in `/api`.
Then use the following command to regenerate the server stubs:

```sh
make openapi-generate
```