// A generated module for Backend functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/backend/internal/dagger"
)

const PostgresVersion = "postgres:17.2-bookworm"
const AlpineVersion = "alpine:3.20"
const GolangVersion = "1.23.1-alpine"

type Backend struct{}

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

func (m *Backend) Ledger(ctx context.Context, src *dagger.Directory, postgres *dagger.Service) *dagger.Service {

	env := map[string]string{
		"APP_LOGLEVEL":          "DEBUG",
		"APP_SERVER_PORT":       "8080",
		"APP_POSTGRES_DATABASE": "postgres",
		"APP_POSTGRES_HOST":     "database",
		"APP_POSTGRES_PASSWORD": "password",
		"APP_POSTGRES_PORT":     "5432",
		"APP_POSTGRES_USER":     "postgres",
	}

	src = src.WithoutFile(".env")

	ledger := m.BuildLedger(src).
		WithExposedPort(8080).
		WithServiceBinding("database", postgres)

	ledger = withEnvVars(ledger, env)

	return ledger.AsService()
}

func (m *Backend) Integration(ctx context.Context, src *dagger.Directory, ledger *dagger.Service) (string, error) {

	return dag.Go(dagger.GoOpts{Version: GolangVersion}).
		WithSource(src).
		WithServiceBinding("ledger", ledger).
		WithEnvVariable("LEDGER_URL", "http://ledger:8080").
		Exec([]string{"go", "test", "-count=1", "./test"}). // pass -count=1 to bust test caching
		Stdout(ctx)
}

// BuildLedger - builds ledger application
func (m *Backend) BuildLedger(src *dagger.Directory) *dagger.Container {
	ledger := dag.Go(dagger.GoOpts{Version: GolangVersion}).
		WithSource(src).
		Build(dagger.GoWithSourceBuildOpts{Pkg: "./cmd/backend"})

	return dag.Container().
		From(AlpineVersion).
		WithWorkdir("/app").
		WithFile("ledger", ledger).
		WithEntrypoint([]string{"/app/ledger"})
}

// PostgresMigrate - spins up a postgres database and runs the migrations against it.
func (m *Backend) PostgresMigrate(ctx context.Context, src *dagger.Directory) (*dagger.Service, error) {
	pg, err := m.Postgres().Start(ctx)
	if err != nil {
		return nil, err
	}

	_, err = m.Migrate(ctx, src, pg)
	if err != nil {
		return nil, err
	}

	return pg, nil
}

// Postgres - creates a new postgres database
func (m *Backend) Postgres() *dagger.Service {
	return dag.Container().
		From(PostgresVersion).
		WithEnvVariable("POSTGRES_PASSWORD", "password").
		WithExposedPort(5432).
		AsService()
}

func (m *Backend) Psql(svc *dagger.Service) *dagger.Container {
	return dag.Container().
		From(PostgresVersion).
		WithEnvVariable("PGPASSWORD", "password").
		WithServiceBinding("database", svc).
		Terminal(dagger.ContainerTerminalOpts{Cmd: []string{"psql", "-h", "database", "-U", "postgres", "-d", "postgres"}})
}

// Migrate - migrates a postgres database.
// `src` is the directory of the backend project.
// `svc` is the database endpoint.
func (m *Backend) Migrate(ctx context.Context, src *dagger.Directory, svc *dagger.Service) (string, error) {
	return dag.Container().
		From("flyway/flyway").
		WithMountedDirectory("/flyway/project", src).
		WithServiceBinding("db", svc).
		WithExec([]string{
			"-url=jdbc:postgresql://db:5432/postgres?user=postgres&password=password",
			"-workingDirectory=project",
			"migrate",
		}, dagger.ContainerWithExecOpts{UseEntrypoint: true}).
		Stdout(ctx)
}

// OpenapiGenerate - generates go server boilerplate code from the openapi spec.
// `src` is the directory of the backend project.
func (m *Backend) OpenapiGenerate(ctx context.Context, src *dagger.Directory) *dagger.Directory {
	return dag.Container().
		From("openapitools/openapi-generator-cli").
		WithDirectory("/local", src.Directory("api")).
		WithDirectory("/out", src.Directory("internal/api")).
		WithExec([]string{
			"generate",
			"-i", "/local/spec.yaml",
			"-g", "go-server",
			"-p", "outputAsLibrary=true",
			"-o", "/out",
		}, dagger.ContainerWithExecOpts{UseEntrypoint: true}).
		Directory("/out")
}

func withEnvVars(ctn *dagger.Container, env map[string]string) *dagger.Container {
	for k, v := range env {
		ctn = ctn.WithEnvVariable(k, v)
	}
	return ctn
}
