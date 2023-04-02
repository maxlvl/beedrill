package loadtester

import (
  "sync"
)


type Collector struct {
  Results sync.Map
}

func NewCollector() *Collector {
  collector := &Collector{}
  return collector
}

func (c *Collector) Collect(result Result) {
  scenarioResults, _ := c.Results.LoadOrStore(result.Scenario, []Result{})
  c.Results.Store(result.Scenario, append(scenarioResults.([]Result), result))
}

