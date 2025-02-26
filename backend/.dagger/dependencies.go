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

type Backend struct{}

const PostgresVersion = "postgres:17.2-bookworm"
const AlpineGitVersion = "alpine/git:2.47.2"

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
		WithExec([]string{
			"generate",
			"-i", "/local/spec.yaml",
			"-g", "go-server",
			"-p", "outputAsLibrary=true",
			"-o", "/out",
		}, dagger.ContainerWithExecOpts{UseEntrypoint: true}).
		Directory("/out").
		WithoutFile("README.md")
}

func (m *Backend) OpenapiDrift(ctx context.Context, src *dagger.Directory) (string, error) {
	gen := m.OpenapiGenerate(ctx, src)

	return dag.Container().
		From(AlpineGitVersion).
		WithWorkdir("/app").
		WithDirectory(".", src.Directory("./internal/api")).
		WithExec([]string{"git", "init"}).
		WithExec([]string{"git", "config", "--global", "user.email", "you@example.com"}).
		WithExec([]string{"git", "config", "--global", "user.name", "Your Name"}).
		WithExec([]string{"git", "add", "*"}).
		WithExec([]string{"git", "commit", "-m", "base"}).
		WithDirectory(".", gen).
		WithExec([]string{"sh", "-c", "if [[ $(git status --porcelain | wc -l) -gt 0 ]]; then git status --porcelain; exit 1; else exit 0; fi"}).
		Stdout(ctx)
}

// Psql - opens a shell to the running database
func (m *Backend) Psql(svc *dagger.Service) *dagger.Container {
	return dag.Container().
		From(PostgresVersion).
		WithEnvVariable("PGPASSWORD", "password").
		WithServiceBinding("database", svc).
		Terminal(dagger.ContainerTerminalOpts{Cmd: []string{"psql", "-h", "database", "-U", "postgres", "-d", "postgres"}})
}
