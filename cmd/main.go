package main

import (
	"github.com/maxlvl/gocust/internal/client"
	"github.com/maxlvl/gocust/internal/loadtester"
	"time"
)

func main() {
	ltConfig := loadtester.LoadTesterConfig{
		Concurrency:  5,
		TestDuration: 10 * time.Second,
		HTTPClientConfig: client.HTTPClientConfig{
			Timeout: 3 * time.Second,
		},
	}

	lt := loadtester.NewLoadTester(ltConfig)

	scenarios := []loadtester.Scenario{
		&loadtester.SimpleScenario{URL: "http://example.com"},
		&loadtester.ComplexScenario{
			GetURL:  "http://example.com",
			PostURL: "http://example.com/post",
			Payload: map[string]string{
				"key": "value",
			},
		},
	}
	lt.Run(scenarios)
}
