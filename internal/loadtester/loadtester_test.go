package loadtester

import (
  "testing"
  "time"
  "net/http"
  "github.com/maxlvl/gocust/internal/client"
	"github.com/stretchr/testify/assert"
)

type MockScenario struct {}

func (ms *MockScenario) Execute(httpclient *http.Client) (*Result, error) {
  time.Sleep(100 * time.Millisecond)
  return &Result{}, nil
}

func TestLoadTester_Run(t *testing.T) {
  config := LoadTesterConfig{
    Concurrency:        2,
    TestDuration:       500 * time.Millisecond,
    HTTPClientConfig:   client.HTTPClientConfig{
      Timeout: 3 * time.Second,
    },
  }

  lt := NewLoadTester(config)

  startTime := time.Now()
  scenarios := []Scenario{&MockScenario{}}
  lt.Run(scenarios)
  elapsed := time.Since(startTime)

  assert.GreaterOrEqual(t, int64(elapsed), int64(500 * time.Millisecond))
  assert.LessOrEqual(t, int64(elapsed), int64(600 * time.Millisecond))

}
