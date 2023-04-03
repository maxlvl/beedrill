package result

import (
  "time"
)

type Result struct {
	Scenario    string
	Success     bool
	Latency     time.Duration
	StartTime   time.Time
	EndTime     time.Time
	StatusCode  int
	Error       error
}
