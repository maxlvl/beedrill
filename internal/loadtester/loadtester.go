package loadtester

import (
  "sync"
  "fmt"
  "time"
	"net/http"
  "math/rand"
  "github.com/maxlvl/gocust/internal/client"
)

type LoadTesterConfig struct {
  Concurrency       int
  TestDuration      time.Duration
  HTTPClientConfig  client.HTTPClientConfig 
}

type LoadTester struct {
  config        LoadTesterConfig  
  httpClient    *http.Client
}

func NewLoadTester(config LoadTesterConfig) *LoadTester {
  return &LoadTester{
    config: config,
    httpClient: client.NewHTTPClient(config.HTTPClientConfig),
  }
}

func (lt *LoadTester) Run(scenarios []Scenario) {
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
        scenarios[scenarioIndex].Execute(lt.httpClient)
      }
    }()
  }

  waitGroup.Wait()
}

type Scenario interface {
  Execute(httpClient *http.Client)
}
