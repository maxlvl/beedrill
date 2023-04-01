package loadtester

type Collector struct {
  Results chan Result
}

func NewCollector() *Collector {
  return &Collector{
    Results: make(chan Result),
  }
}

func (c *Collector) Collect(result *Result) {
  c.Results <- result
}

func (c *Collector) Close() {
  reporter := NewReporter(c)
  reporter.Report()
  close(c.Results)
}
