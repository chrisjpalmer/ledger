package main

import (
	"context"
	"dagger/backend/internal/dagger"
	"strings"
)

const AlpineVersion = "alpine:3.20"

// CheckPullRequest - runs checks to validate ledger
func (m *Backend) CheckPullRequest(ctx context.Context, spec *dagger.File, src *dagger.Directory) (string, error) {

	// run open api drift
	drift, err := m.OpenapiDrift(ctx, spec, src)
	if err != nil {
		return "", err
	}

	return strings.Join([]string{drift}, "\n\n"), nil
}
