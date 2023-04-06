package config

import (
	"github.com/maxlvl/gocust/internal/client"
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
