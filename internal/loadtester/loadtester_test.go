package loadtester

import (
	"github.com/maxlvl/beedrill/internal/client"
	"github.com/maxlvl/beedrill/internal/result"
	"github.com/maxlvl/beedrill/scenarios"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

type MockScenario struct{}

func (ms *MockScenario) Execute(httpclient *http.Client) (*result.Result, error) {
	time.Sleep(100 * time.Millisecond)
	return &result.Result{}, nil
}

func TestLoadTester_Run(t *testing.T) {
	config := LoadTesterConfig{
		Concurrency:  2,
		TestDuration: 500 * time.Millisecond,
		HTTPClientConfig: client.HTTPClientConfig{
			Timeout: 3 * time.Second,
		},
	}

	lt := NewLoadTester(config)

	startTime := time.Now()
	scenarios := []scenarios.Scenario{&MockScenario{}}
	lt.Run(scenarios)
	elapsed := time.Since(startTime)

	assert.GreaterOrEqual(t, int64(elapsed), int64(500*time.Millisecond))
	assert.LessOrEqual(t, int64(elapsed), int64(600*time.Millisecond))

}
