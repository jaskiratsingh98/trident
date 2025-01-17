// Copyright 2023 NetApp, Inc. All Rights Reserved.

package acp

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/netapp/trident/logging"
)

func TestMain(m *testing.M) {
	// Disable any standard log output
	logging.InitLogOutput(io.Discard)
	os.Exit(m.Run())
}

// newHttpTestServer sets up a httptest server with a supplied handler and pattern.
func newHttpTestServer(pattern string, handler http.HandlerFunc) *httptest.Server {
	router := http.NewServeMux()
	router.HandleFunc(pattern, handler)
	return httptest.NewServer(router)
}

func TestNewAPI(t *testing.T) {
	tests := map[string]struct {
		baseURL string
		timeout time.Duration
	}{
		"WithMisconfiguredBaseURL": {
			baseURL: "://127.0.0.1:8100",
			timeout: httpClientTimeout,
		},
		// ΗΤΤΡ here is defined using unicode: Η = capital eta; Τ = capital tau; Ρ=capital rho
		"WithInvalidCharactersInBaseURL": {
			baseURL: "ΗΤΤΡ://127.0.0.1:8100",
			timeout: httpClientTimeout,
		},
		"WithNegativeTimeoutSpecified": {
			baseURL: "http://127.0.0.1.8100",
			timeout: time.Duration(-1),
		},
		"WithZeroTimeoutSpecified": {
			baseURL: "http://127.0.0.1.8100",
			timeout: time.Duration(0),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			c, err := NewAPI(test.baseURL, test.timeout)
			assert.Nil(t, c, "expected nil client")
			assert.Error(t, err, "expected error")
		})
	}
}

func TestClient_GetVersion(t *testing.T) {
	ctx := context.Background()
	t.Run("WithNoServerRunning", func(t *testing.T) {
		server := newHttpTestServer(versionEndpoint, func(w http.ResponseWriter, request *http.Request) {})
		// Close the server immediately.
		server.Close()

		client := &Client{baseURL: server.URL, httpClient: *server.Client()}

		version, err := client.GetVersion(ctx)
		assert.Nil(t, version, "expected version to be nil")
		assert.Error(t, err, "expected error")
	})

	t.Run("WithInternalServerError", func(t *testing.T) {
		handler := func(w http.ResponseWriter, request *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte{})
		}
		server := newHttpTestServer(versionEndpoint, handler)
		defer server.Close()

		client := &Client{baseURL: server.URL, httpClient: *server.Client()}

		version, err := client.GetVersion(ctx)
		assert.Nil(t, version, "expected version to be nil")
		assert.Error(t, err, "expected error")
	})

	t.Run("WithIncorrectResponseType", func(t *testing.T) {
		handler := func(w http.ResponseWriter, request *http.Request) {
			if _, err := w.Write(nil); err != nil {
				t.Fatalf("failed to write response for GetVersion")
			}
		}
		server := newHttpTestServer(versionEndpoint, handler)
		defer server.Close()

		client := &Client{baseURL: server.URL, httpClient: *server.Client()}

		version, err := client.GetVersion(ctx)
		assert.Nil(t, version, "expected version to be nil")
		assert.Error(t, err, "expected error")
	})

	t.Run("WithErrorInResponse", func(t *testing.T) {
		expectedVersion := "23.10.0-custom+unknown"
		handler := func(w http.ResponseWriter, request *http.Request) {
			v := &getVersionResponse{Version: expectedVersion, Error: "non-empty error"}
			bytes, err := json.Marshal(v)
			if err != nil {
				t.Fatalf("failed to create fake response for GetVersion")
			}

			if _, err = w.Write(bytes); err != nil {
				t.Fatalf("failed to write response for GetVersion")
			}
		}
		server := newHttpTestServer(versionEndpoint, handler)
		defer server.Close()

		client := &Client{baseURL: server.URL, httpClient: *server.Client()}

		version, err := client.GetVersion(ctx)
		assert.Nil(t, version, "expected version to be nil")
		assert.Error(t, err, "expected error")
	})

	t.Run("WithCorrectResponseType", func(t *testing.T) {
		expectedVersion := "23.10.0-custom+unknown"
		handler := func(w http.ResponseWriter, request *http.Request) {
			bytes, err := json.Marshal(&getVersionResponse{Version: expectedVersion})
			if err != nil {
				t.Fatalf("failed to create fake response for GetVersion")
			}

			if _, err = w.Write(bytes); err != nil {
				t.Fatalf("failed to write response for GetVersion")
			}
		}
		server := newHttpTestServer(versionEndpoint, handler)
		defer server.Close()

		client := &Client{baseURL: server.URL, httpClient: http.Client{Timeout: httpClientTimeout}}

		version, err := client.GetVersion(ctx)
		assert.NotNil(t, version, "expected version to exist")
		assert.Equal(t, version.String(), expectedVersion, "expected equivalent versions")
		assert.NoError(t, err, "unexpected error")
	})
}

func TestClient_Entitled(t *testing.T) {
	ctx := context.Background()
	feature := FeatureSnapshotMirrorUpdate
	t.Run("WithNoServerRunning", func(t *testing.T) {
		server := newHttpTestServer("/any", func(w http.ResponseWriter, request *http.Request) {})
		// Close the server immediately.
		server.Close()

		client := &Client{baseURL: server.URL, httpClient: *server.Client()}

		allowed, err := client.Entitled(ctx, feature)
		assert.False(t, allowed, "expected allowed to be false")
		assert.Error(t, err, "expected error")
	})

	t.Run("WithInternalServerError", func(t *testing.T) {
		handler := func(w http.ResponseWriter, request *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte{})
		}
		server := newHttpTestServer(entitledEndpoint, handler)
		defer server.Close()

		client := &Client{baseURL: server.URL, httpClient: *server.Client()}

		allowed, err := client.Entitled(ctx, feature)
		assert.False(t, allowed, "expected allowed to be false")
		assert.Error(t, err, "expected error")
	})

	t.Run("WithAllowedResponse", func(t *testing.T) {
		handler := func(w http.ResponseWriter, request *http.Request) {
			w.WriteHeader(200)
		}
		server := newHttpTestServer(entitledEndpoint, handler)
		defer server.Close()

		client := &Client{baseURL: server.URL, httpClient: http.Client{Timeout: httpClientTimeout}}

		allowed, err := client.Entitled(ctx, feature)
		assert.True(t, allowed, "expected allowed to be true")
		assert.NoError(t, err, "unexpected error")
	})
	// TODO: Add unit tests for inspecting the response body when the Entitled API changes.
}

func TestClient_newRequest(t *testing.T) {
	tests := map[string]struct {
		ctx         context.Context
		method, url string
		data        []byte
	}{
		"WithNilContext": {
			method: "GET",
			url:    versionEndpoint,
		},
		"WithNonNilBody": {
			method: "GET",
			url:    versionEndpoint,
			data:   []byte{'t', 'r', 'i', 'd', 'e', 'n', 't'},
		},
		"WithEmptyMethod": {
			method: "",
			url:    versionEndpoint,
		},
		"WithEmptyURL": {
			method: "GET",
			url:    "",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			client, err := NewAPI("http://127.0.0.1:8100", httpClientTimeout)
			if err != nil {
				t.Fatalf("failed to create ACP API client for tests")
			}

			req, err := client.newRequest(test.ctx, test.method, test.url, test.data)
			assert.NotNil(t, req)
			assert.NoError(t, err)

			var body []byte
			if test.data != nil {
				body, err = io.ReadAll(req.Body)
				if err != nil {
					t.Fatalf("failed to parse fake request body; %v", err)
				}
			}
			assert.Equal(t, test.data, body)
		})
	}
}

func TestClient_invokeAPI(t *testing.T) {
	ctx := context.Background()
	t.Run("WithNoServerRunning", func(t *testing.T) {
		server := newHttpTestServer(versionEndpoint, func(w http.ResponseWriter, request *http.Request) {})
		// Close the server immediately.
		server.Close()

		client := &Client{baseURL: server.URL, httpClient: *server.Client()}
		req, err := client.newRequest(ctx, "GET", server.URL+versionEndpoint, nil)
		if err != nil {
			t.Fatalf("failed to create fake test request; %v", err)
		}

		res, body, err := client.invokeAPI(req)
		assert.Nil(t, res, "expected res to be nil")
		assert.Nil(t, body, "expected body to be nil")
		assert.Error(t, err, "expected error")
	})

	t.Run("WithInternalServerError", func(t *testing.T) {
		handler := func(w http.ResponseWriter, request *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte{})
		}
		server := newHttpTestServer(versionEndpoint, handler)
		defer server.Close()

		client := &Client{baseURL: server.URL, httpClient: *server.Client()}
		req, err := client.newRequest(ctx, "GET", server.URL+versionEndpoint, nil)
		if err != nil {
			t.Fatalf("failed to create fake test request; %v", err)
		}

		res, body, err := client.invokeAPI(req)
		assert.NotNil(t, res, "expected res to not be nil")
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "expected internal server error")
		assert.NotNil(t, body, "expected body to not be nil")
		assert.NoError(t, err, "expected no error")
	})
}
