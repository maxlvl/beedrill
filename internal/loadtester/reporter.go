package loadtester

import (
  "fmt"
  "time"
)

type Reporter struct {
  Collector *Collector
}


func NewReporter(collector *Collector) *Reporter {
  return &Reporter{
    Collector: collector,
  }
}

func (r *Reporter) Report() {
  fmt.Println("\nLoad Testing Results:")
  r.Collector.Results.Range(func(key, value interface{}) bool {
    scenarioName := key.(string)
    scenarioResults := value.([]Result)
    successCount := 0
    totalLatency := time.Duration(0)

    for _, result := range scenarioResults {
      if result.Success {
        successCount++
      }
      totalLatency += result.Latency
    }

    averageLatency := totalLatency / time.Duration(len(scenarioResults))
    successRate := float64(successCount) / float64(len(scenarioResults)) * 100

    fmt.Printf("\nScenario: %s\n", scenarioName)
		fmt.Printf("  Total Requests: %d\n", len(scenarioResults))
		fmt.Printf("  Successful Requests: %d\n", successCount)
		fmt.Printf("  Success Rate: %.2f%%\n", successRate)
		fmt.Printf("  Average Latency: %v\n", averageLatency)
    return true
  })
}
