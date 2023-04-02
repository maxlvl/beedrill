package loadtester

import (
  "sync"
  "fmt"
  "time"
	"net/http"
  "math/rand"
  "github.com/maxlvl/gocust/internal/client"
  "log"
)

type Result struct {
	Scenario    string
	Success     bool
	Latency     time.Duration
	StartTime   time.Time
	EndTime     time.Time
	StatusCode  int
	Error       error
}

type LoadTesterConfig struct {
  Concurrency       int
  TestDuration      time.Duration
  HTTPClientConfig  client.HTTPClientConfig 
}

type LoadTester struct {
  config        LoadTesterConfig  
  httpClient    *http.Client
  collector     *Collector
}

func NewLoadTester(config LoadTesterConfig) *LoadTester {
  return &LoadTester{
    config: config,
    httpClient: client.NewHTTPClient(config.HTTPClientConfig),
    collector: NewCollector(),
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
        result, err := scenarios[scenarioIndex].Execute(lt.httpClient)
        if err != nil {
          log.Fatal("Error running Scenario: %s", err)
        }
        lt.collector.Collect(*result)
      }
    }()
  }
  waitGroup.Wait()
  fmt.Printf("Waited until waitgroup was done")
  lt.collector.Close()
  reporter := NewReporter(lt.collector)
  reporter.Report()
}

