package metrics

import (
  "fmt"
  "time"
  "github.com/maxlvl/gocust/internal/loadtester"
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
  resultMap := make(map[string][]loadtester.Result)

  for result := range r.Collector.Results{
    resultMap[result.Scenario] = append(resultMap[result.Scenario], result)
  }

  fmt.Println("\nLoad Testing Results:\n")
  for scenarioName, scenarioResults := range resultMap {
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
  }
}
