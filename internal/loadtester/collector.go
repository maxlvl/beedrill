package loadtester

import (
  "sync"
  "github.com/maxlvl/gocust/internal/result"
)


type Collector struct {
  Results sync.Map
}

func NewCollector() *Collector {
  collector := &Collector{}
  return collector
}

func (c *Collector) Collect(rs result.Result) {
  scenarioResults, _ := c.Results.LoadOrStore(rs.Scenario, []result.Result{})
  c.Results.Store(rs.Scenario, append(scenarioResults.([]result.Result), rs))
}

