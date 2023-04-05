package metrics

import (
	"github.com/maxlvl/gocust/internal/result"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollector(t *testing.T) {
	collector := NewCollector()

	r1 := result.Result{
		Scenario: "scenario1",
		Success:  true,
	}

	collector.Collect(r1)

	scenarioResults, ok := collector.Results.Load("scenario1")
	assert.True(t, ok, "Scenario results should be found")
	assert.Len(t, scenarioResults, 1, "Scenario results should have 1 result")
	assert.Equal(t, r1, scenarioResults.([]result.Result)[0], "First result should match r1")

	r2 := result.Result{
		Scenario: "scenario1",
		Success:  false,
	}

	collector.Collect(r2)

	scenarioResults, ok = collector.Results.Load("scenario1")
	assert.True(t, ok, "Scenario results should be found")
	assert.Len(t, scenarioResults, 2, "Scenario results should have 2 results")
	assert.Equal(t, r2, scenarioResults.([]result.Result)[1], "Second result should match r2")
}
