package metrics

import "github.com/maxlvl/gocust/internal/loadtester"

type Collector struct {
  Results chan loadtester.Result
}

func NewCollector() *Collector {
  return &Collector{
    Results: make(chan loadtester.Result),
  }
}

func (c *Collector) Collect(result loadtester.Result) {
  c.Results <- result
}

func (c *Collector) Close() {
  reporter := metrics.NewReporter(c)
  reporter.Report()
  close(c.Results)
}
