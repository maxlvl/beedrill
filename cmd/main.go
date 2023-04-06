package main

import (
	"fmt"
	"github.com/maxlvl/gocust/config"
	"github.com/maxlvl/gocust/internal/loadtester"
	"github.com/maxlvl/gocust/web/server"
	yaml "gopkg.in/yaml.v3"
	"io/ioutil"
)

func main() {
	configData, err := ioutil.ReadFile("../config.yaml")
	if err != nil {
		fmt.Sprintf("Error reading config: %s\n", err)
		return
	}

	var cfg config.Config
	err = yaml.Unmarshal(configData, &cfg)
	if err != nil {
		fmt.Sprintf("Error unmarshalling config: %s\n", err)
		return
	}

	ltConfig := loadtester.LoadTesterConfig{
		Concurrency:      cfg.Concurrency,
		TestDuration:     cfg.TestDuration,
		HTTPClientConfig: cfg.HTTPClientConfig,
	}

	lt := loadtester.NewLoadTester(ltConfig)

	srv := web.NewServer(":8080", lt, cfg)
	err = srv.Start()
	if err != nil {
		fmt.Sprintf("Error starting webserver : %s\n", err)
		return
	}
}
