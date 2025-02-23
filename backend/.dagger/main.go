package main

import (
	"context"
	"dagger/backend/internal/dagger"
)

const AlpineVersion = "alpine:3.20"
const GolangVersion = "1.23.1-alpine"

// CheckPullRequest - runs checks to validate ledger
func (m *Backend) CheckPullRequest(ctx context.Context, src *dagger.Directory) (string, error) {

	// start postgres
	postgres, err := m.PostgresMigrate(ctx, src)
	if err != nil {
		return "", err
	}
	defer postgres.Stop(ctx)

	// start ledger
	ledger := m.Ledger(ctx, src, postgres)
	ledger, err = ledger.Start(ctx)
	if err != nil {
		return "", err
	}
	defer ledger.Stop(ctx)

	// run integration
	return m.Integration(ctx, src, ledger)
}

// Ledger - runs the ledger application
func (m *Backend) Ledger(ctx context.Context, src *dagger.Directory, postgres *dagger.Service) *dagger.Service {

	ledger := m.BuildLedger(src)

	env := map[string]string{
		"APP_LOGLEVEL":          "DEBUG",
		"APP_SERVER_PORT":       "8080",
		"APP_POSTGRES_DATABASE": "postgres",
		"APP_POSTGRES_HOST":     "database",
		"APP_POSTGRES_PASSWORD": "password",
		"APP_POSTGRES_PORT":     "5432",
		"APP_POSTGRES_USER":     "postgres",
	}
	ledger = withEnvVars(ledger, env)

	ledger = ledger.WithServiceBinding("database", postgres)

	return ledger.AsService()
}

// BuildLedger - builds ledger application
func (m *Backend) BuildLedger(src *dagger.Directory) *dagger.Container {

	src = cleanSource(src)

	ledger := dag.Go(dagger.GoOpts{Version: GolangVersion}).
		WithSource(src).
		Build(dagger.GoWithSourceBuildOpts{Pkg: "./cmd/backend"})

	return dag.Container().
		From(AlpineVersion).
		WithWorkdir("/app").
		WithFile("ledger", ledger).
		WithEntrypoint([]string{"/app/ledger"}).
		WithExposedPort(8080)
}

// Integration - runs the integration tests
func (m *Backend) Integration(ctx context.Context, src *dagger.Directory, ledger *dagger.Service) (string, error) {

	src = cleanSource(src)

	return dag.Go(dagger.GoOpts{Version: GolangVersion}).
		WithSource(src).
		WithServiceBinding("ledger", ledger).
		WithEnvVariable("LEDGER_URL", "http://ledger:8080").
		Exec([]string{"go", "test", "-count=1", "./integration"}). // pass -count=1 to bust test caching
		Stdout(ctx)
}

// cleanSource excludes files from the source directory
// which are not meant to be in builds or images
func cleanSource(src *dagger.Directory) *dagger.Directory {
	return src.WithoutFile(".env")
}

func withEnvVars(ctn *dagger.Container, env map[string]string) *dagger.Container {
	for k, v := range env {
		ctn = ctn.WithEnvVariable(k, v)
	}
	return ctn
}
