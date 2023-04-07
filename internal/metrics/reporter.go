package metrics

import (
	"fmt"
	"github.com/maxlvl/beedrill/internal/result"
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

func (r *Reporter) Report() bool {
	fmt.Println("\nLoad Testing Results:")
	reporter_success := true
	r.Collector.Results.Range(func(key, value interface{}) bool {
		scenarioName := key.(string)
		scenarioResults := value.([]result.Result)
		successCount := 0
		totalLatency := time.Duration(0)

		for _, result := range scenarioResults {
			if result.Success {
				successCount++
			}
			if result.Latency < 0 {
				fmt.Printf("Error encountered: result.Latency is below zero\n")
				reporter_success = false
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
	return reporter_success
}
