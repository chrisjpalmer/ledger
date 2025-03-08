package testreq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// Patch - sends a patch request
func Patch(t *testing.T, path string, query map[string]string, body interface{}) (*http.Response, error) {
	t.Helper()

	return req(t, http.MethodPatch, path, query, body)
}

// Post - sends a post request
func Post(t *testing.T, path string, query map[string]string, body interface{}) (*http.Response, error) {
	t.Helper()

	return req(t, http.MethodPost, path, query, body)
}

// Delete - sends a delete request
func Delete(t *testing.T, path string, query map[string]string, body interface{}) (*http.Response, error) {
	t.Helper()

	return req(t, http.MethodDelete, path, query, body)
}

// Get - sends a get request
func Get(t *testing.T, path string, query map[string]string) (*http.Response, error) {
	t.Helper()

	return req(t, http.MethodGet, path, query, nil)
}

func req(t *testing.T, method string, path string, query map[string]string, body interface{}) (*http.Response, error) {
	t.Helper()

	// fetch url from URL or use default
	ledgerURL := os.Getenv("LEDGER_URL")
	if ledgerURL == "" {
		ledgerURL = "http://localhost:8080"
	}

	_u := fmt.Sprintf("%s%s", ledgerURL, path)

	// query
	if query != nil {
		var q url.Values
		for k, v := range query {
			q.Set(k, v)
		}
		_u = fmt.Sprintf("%s?%s", _u, q.Encode())
	}

	// url
	u, err := url.Parse(_u)
	require.NoError(t, err, "parsing request url")

	// req
	r := &http.Request{
		URL:    u,
		Method: method,
	}

	//body
	if body != nil {
		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		err = enc.Encode(body)
		require.NoError(t, err, "decoding request body")

		r.Body = io.NopCloser(&b)
	}

	// send
	rs, err := http.DefaultClient.Do(r)

	return rs, err
}
