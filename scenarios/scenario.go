package scenarios

import (
	"encoding/json"
	"github.com/maxlvl/gocust/internal/client"
	"github.com/maxlvl/gocust/internal/result"
	"net/http"
	"time"
)

type Scenario interface {
	Execute(httpClient *http.Client) (*result.Result, error)
}

type SimpleScenario struct {
	URL string
}

func (s *SimpleScenario) Execute(httpClient *http.Client) (*result.Result, error) {
	startTime := time.Now()
	result := &result.Result{
		Scenario:  "SimpleScenario",
		StartTime: startTime,
	}

	resp, err := client.Get(httpClient, s.URL)
	if err != nil {
		result.Error = err
		result.Success = false
		return result, err
	}
	defer resp.Body.Close()

	result.EndTime = time.Now()
	result.Latency = result.EndTime.Sub(startTime)
	result.StatusCode = resp.StatusCode
	result.Success = resp.StatusCode == 200 && resp.StatusCode < 300

	return result, nil
}

type ComplexScenario struct {
	GetURL  string
	PostURL string
	Payload interface{}
}

func (c *ComplexScenario) Execute(httpClient *http.Client) (*result.Result, error) {
	startTime := time.Now()

	result := &result.Result{
		Scenario:  "ComplexScenario",
		StartTime: startTime,
	}

	httpClient.Get(c.GetURL)
	data, _ := json.Marshal(c.Payload)
	resp, err := client.Post(httpClient, c.PostURL, string(data))
	if err != nil {
		result.Error = err
		result.Success = false
		return result, err
	}
	defer resp.Body.Close()

	result.EndTime = time.Now()
	result.Latency = result.EndTime.Sub(startTime)
	result.StatusCode = resp.StatusCode
	result.Success = resp.StatusCode == 200 && resp.StatusCode < 300

	return result, nil
}
