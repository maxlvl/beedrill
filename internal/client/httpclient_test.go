package client


import (
  "bytes"
  "io/ioutil"
  "net/http"
  "net/http/httptest"
  "testing"
  "time"

  "github.com/stretchr/testify/assert"
)

func TestHTTPClient(t *testing.T) {
  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
  }))
  defer ts.Close()

  config := client.HTTPClientConfig{
    Timeout: 5 * time.Second,
  }

  httpClient := client.NewHTTPClient(config)

  resp, err := httpClient.Get(ts.URL)
  assert.NoError(t, err)
  assert.Equal(t, http.StatusOK, resp.StatusCode)

  body, err := ioutil.ReadAll(resp.Body)
  resp.Body.Close()
  assert.NoError(t, err)
  assert.Equal("OK", string(body))
}
