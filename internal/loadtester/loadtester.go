package loadtester

import (
	"fmt"
	"github.com/maxlvl/beedrill/internal/client"
	"github.com/maxlvl/beedrill/internal/metrics"
	"github.com/maxlvl/beedrill/scenarios"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type LoadTesterConfig struct {
	Concurrency      int
	TestDuration     time.Duration
	HTTPClientConfig client.HTTPClientConfig
}

type LoadTester struct {
	config     LoadTesterConfig
	httpClient *http.Client
	collector  *metrics.Collector
}

func NewLoadTester(config LoadTesterConfig) *LoadTester {
	return &LoadTester{
		config:     config,
		httpClient: client.NewHTTPClient(config.HTTPClientConfig),
		collector:  metrics.NewCollector(),
	}
}

func (lt *LoadTester) Run(scenarios []scenarios.Scenario) {
	// no need for channels yet as each goroutine can run independently
	// will likely introduce channels once we need aggregated metrics for all goroutines
	var waitGroup sync.WaitGroup
	fmt.Printf("Adding %d concurrent workers...\n", lt.config.Concurrency)
	waitGroup.Add(lt.config.Concurrency)

	fmt.Printf("Test duration: %s\n", lt.config.TestDuration) // Print TestDuration

	for i := 0; i < lt.config.Concurrency; i++ {
		go func() {
			defer waitGroup.Done()
			startTime := time.Now()

			for time.Since(startTime) < lt.config.TestDuration {
				scenarioIndex := rand.Intn(len(scenarios))
				result, err := scenarios[scenarioIndex].Execute(lt.httpClient)
				if err != nil {
					fmt.Printf("Error running Scenario: %s", err)
				}
				lt.collector.Collect(*result)
			}
		}()
	}

	waitGroup.Wait()
	reporter := metrics.NewReporter(lt.collector)
	reporter.Report()
}
