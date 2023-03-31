package loadtester

import (
  "sync"
  "fmt"
  "time"
	"net/http"
  "math/rand"
  "github.com/maxlvl/gocust/internal/client"
  "github.com/maxlvl/gocust/internal/metrics"
)

type LoadTesterConfig struct {
  Concurrency       int
  TestDuration      time.Duration
  HTTPClientConfig  client.HTTPClientConfig 
}

type LoadTester struct {
  config        LoadTesterConfig  
  httpClient    *http.Client
  collector     *metrics.Collector
}

func NewLoadTester(config LoadTesterConfig) *LoadTester {
  return &LoadTester{
    config: config,
    httpClient: client.NewHTTPClient(config.HTTPClientConfig),
    collector := metrics.NewCollector()
  }
}

func (lt *LoadTester) Run(scenarios []Scenario) {
  // no need for channels yet as each goroutine can run independently
  // will likely introduce channels once we need aggregated metrics for all goroutines
  var waitGroup sync.WaitGroup
  fmt.Printf("Adding %d concurrent workers...\n", lt.config.Concurrency)
  waitGroup.Add(lt.config.Concurrency)

  for i := 0; i < lt.config.Concurrency; i++ {
    go func() {
      defer waitGroup.Done()
      startTime := time.Now()
      fmt.Printf("Running random scenario for startTime %s\n", startTime.Format("2006-01-02 15:04:05"))


      for time.Since(startTime) < lt.config.TestDuration {
        scenarioIndex := rand.Intn(len(scenarios))
        result := scenarios[scenarioIndex].Execute(lt.httpClient)
        lt.collector.Collect(result)
      }
    }()
  }

  go func() {
    waitGroup.Wait()
    lt.collector.Close()
    reporter := metrics.NewReporter(lt.collector)
    reporter.Report()
  }()
}

type Scenario interface {
  Execute(httpClient *http.Client)
}
