package metrics

import (
	"github.com/maxlvl/beedrill/internal/result"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReporter(t *testing.T) {
	collector := NewCollector()

	r1 := result.Result{
		Scenario: "scenario1",
		Success:  true,
	}

	collector.Collect(r1)

	reporter := NewReporter(collector)
	reporter_result := reporter.Report()
	assert.True(t, reporter_result)
}
