package loadtester

import (
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
	"github.com/maxlvl/gocust/internal/client"
	"github.com/maxlvl/gocust/scenarios"

)

func TestLoadTesterIntegration(t *testing.T) {
  var requestCounter int32

  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    atomic.AddInt32(&requestCounter, 1)
    w.WriteHeader(http.StatusOK)
  }))
  
  defer ts.Close()

  config := LoadTesterConfig{
    Concurrency:        2,
    TestDuration:       500 * time.Millisecond,
    HTTPClientConfig:   client.HTTPClientConfig{
      Timeout: 3 * time.Second,
    },
  }

  lt := NewLoadTester(config)

  scenarios := []scenarios.Scenario{&scenarios.SimpleScenario{URL: ts.URL}}

  lt.Run(scenarios)
  assert.Greater(t, int(atomic.LoadInt32(&requestCounter)), 0, "load tester should've incremented the counter at least once")
}
