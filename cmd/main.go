package main

import (
	"fmt"
	"github.com/maxlvl/gocust/internal/client"
	"github.com/maxlvl/gocust/internal/loadtester"
	"github.com/maxlvl/gocust/scenarios"
	"github.com/maxlvl/gocust/web/server"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"time"
)

type Config struct {
	Concurrency      int                     `yaml:"concurrency"`
	TestDuration     time.Duration           `yaml:"test_duration"`
	HTTPClientConfig client.HTTPClientConfig `yaml:"http_client_config"`
	Scenarios        []ScenarioConfig        `yaml:"scenarios"`
}

type ScenarioConfig struct {
	Type    string            `yaml:"type"`
	URL     string            `yaml:"url,omitempty"`
	GetURL  string            `yaml:"get_url,omitempty"`
	PostURL string            `yaml:"post_url,omitempty"`
	Payload map[string]string `yaml:"payload,omitempty"`
}

func main() {
	configData, err := ioutil.ReadFile("../config.yaml")
	if err != nil {
		fmt.Sprintf("Error reading config: %s\n", err)
		return
	}

	var config Config
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		fmt.Sprintf("Error unmarshalling config: %s\n", err)
		return
	}

	ltConfig := loadtester.LoadTesterConfig{
		Concurrency:      config.Concurrency,
		TestDuration:     config.TestDuration,
		HTTPClientConfig: config.HTTPClientConfig,
	}

	lt := loadtester.NewLoadTester(ltConfig)

	srv := web.NewServer(":8080", lt)
	err = srv.Start()
	if err != nil {
		fmt.Sprintf("Error starting webserver : %s\n", err)
		return
	}

	var scenarios_array []scenarios.Scenario
	for _, s := range config.Scenarios {
		switch s.Type {
		case "simple":
			scenarios_array = append(scenarios_array, &scenarios.SimpleScenario{URL: s.URL})
		case "complex":
			scenarios_array = append(
				scenarios_array,
				&scenarios.ComplexScenario{
					GetURL:  s.GetURL,
					PostURL: s.PostURL,
					Payload: s.Payload,
				})
		}
	}
	lt.Run(scenarios_array)
}
