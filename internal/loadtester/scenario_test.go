package loadtester

import (
  "testing"
  "time"
  "net/http"
  "net/http/httptest"
  "github.com/maxlvl/gocust/internal/client"
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
