package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/chrisjpalmer/ledger/backend/integration/testreq"
	"github.com/stretchr/testify/require"
)

// Tests intentionally not generated from open api spec
// to allow behavior to be pinned between spec changes.

func TestAddIncome(t *testing.T) {
	rs, err := testreq.Post(t,
		fmt.Sprintf("/month/%d/income", 1),
		nil,
		map[string]any{
			"date":     "2025-01-01",
			"amount":   50.0,
			"name":     "salary",
			"received": true,
		},
	)
	require.NoError(t, err, "request error")

	require.Equal(t, http.StatusOK, rs.StatusCode)

	defer rs.Body.Close()
	dec := json.NewDecoder(rs.Body)

	var data map[string]any
	err = dec.Decode(&data)
	require.NoError(t, err, "error while decoding response")

	require.Contains(t, data, "id")
	require.NotEmpty(t, data["id"])
}
