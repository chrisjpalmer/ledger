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

const (
	AlpineGitVersion           = "alpine/git:2.47.2"
	OpenapiGeneratorCLIVersion = "openapitools/openapi-generator-cli:v7.12.0"
)

// OpenapiGenerate - generates go server boilerplate code from the openapi spec.
// `src` is the directory of the backend project.
func (m *Backend) OpenapiGenerate(ctx context.Context, spec *dagger.File) *dagger.Directory {
	return dag.Container().
		From(OpenapiGeneratorCLIVersion).
		WithFile("/local/spec.yaml", spec).
		WithExec([]string{
			"generate",
			"-i", "/local/spec.yaml",
			"-g", "typescript-fetch",
			"-p", "outputAsLibrary=true",
			"-o", "/out",
		}, dagger.ContainerWithExecOpts{UseEntrypoint: true}).
		Directory("/out").
		WithoutFile("README.md")
}

func (m *Backend) OpenapiDrift(ctx context.Context, spec *dagger.File, src *dagger.Directory) (string, error) {
	gen := m.OpenapiGenerate(ctx, spec)

	return dag.Container().
		From(AlpineGitVersion).
		WithWorkdir("/app").
		WithDirectory(".", src.Directory("./src/lib/api")).
		WithExec([]string{"git", "init"}).
		WithExec([]string{"git", "config", "--global", "user.email", "you@example.com"}).
		WithExec([]string{"git", "config", "--global", "user.name", "Your Name"}).
		WithExec([]string{"git", "add", "*"}).
		WithExec([]string{"git", "commit", "-m", "base"}).
		WithDirectory(".", gen).
		WithExec([]string{"sh", "-c", "if [[ $(git status --porcelain | wc -l) -gt 0 ]]; then git status --porcelain; exit 1; else exit 0; fi"}).
		Stdout(ctx)
}
