package client


import (
  "net/http"
  "net/http/httptest"
  "testing"
  "time"

  "github.com/stretchr/testify/assert"
)

func TestHTTPClientGet(t *testing.T) {
  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
  }))
  defer ts.Close()

  config := HTTPClientConfig{
    Timeout: 5 * time.Second,
  }

  httpClient := NewHTTPClient(config)

  body, err := Get(httpClient, ts.URL)
  assert.NoError(t, err)
  assert.Equal(t, "OK", body)
}

func TestHTTPClientPost(t *testing.T) {
  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
  }))
  defer ts.Close()

  config := HTTPClientConfig{
    Timeout: 5 * time.Second,
  }

  httpClient := NewHTTPClient(config)
  payload := "{'hello': 'world'}"

  body, err := Post(httpClient, ts.URL, payload)
  assert.NoError(t, err)
  assert.Equal(t, "OK", body)
}
