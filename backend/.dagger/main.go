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

// Returns a container that echoes whatever string argument is provided
func (m *Backend) ContainerEcho(stringArg string) *dagger.Container {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg})
}

// Returns lines that match a pattern in the files of the provided Directory
func (m *Backend) GrepDir(ctx context.Context, directoryArg *dagger.Directory, pattern string) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"grep", "-R", pattern, "."}).
		Stdout(ctx)
}

func (m *Backend) Postgres(ctx context.Context) *dagger.Service {
	return dag.Container().
		From("postgres:17.2-bookworm").
		WithEnvVariable("POSTGRES_PASSWORD", "password").
		WithExposedPort(5432).
		AsService()
}

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
