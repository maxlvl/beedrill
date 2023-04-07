package scenarios

import (
	"github.com/maxlvl/beedrill/internal/client"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSimpleScenario_Execute(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	httpClient := client.NewHTTPClient(client.HTTPClientConfig{
		Timeout: 3 * time.Second,
	})

	scenario := &SimpleScenario{URL: ts.URL}
	// no assserts, just check that it runs without errors since it doesn't return anything
	scenario.Execute(httpClient)
}

func TestComplexScenario_Execute(t *testing.T) {
	var getRequestReceived, postRequestReceived bool

	GetTs := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		getRequestReceived = true
	}))
	defer GetTs.Close()

	PostTs := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		postRequestReceived = true
	}))
	defer PostTs.Close()

	httpClient := client.NewHTTPClient(client.HTTPClientConfig{
		Timeout: 3 * time.Second,
	})

	scenario := &ComplexScenario{
		GetURL:  GetTs.URL,
		PostURL: PostTs.URL,
		Payload: map[string]string{
			"key": "value",
		},
	}

	scenario.Execute(httpClient)

	assert.True(t, getRequestReceived, "getRequestReceived should be true")
	assert.True(t, postRequestReceived, "postRequestReceived should be true")
}
