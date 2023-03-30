package loadtester

import (
  "encoding/json"
  "net/http"
	"github.com/maxlvl/gocust/internal/client"
)

type SimpleScenario struct {
  URL     string
}

func (s *SimpleScenario) Execute(httpClient *http.Client) {
  client.Get(httpClient, s.URL)
}


type ComplexScenario struct {
  GetURL       string
  PostURL      string
  Payload      interface{}
}

func (c *ComplexScenario) Execute(httpClient *http.Client) {
  httpClient.Get(c.GetURL)
  data, _ := json.Marshal(c.Payload)
  client.Post(httpClient, c.PostURL, string(data))
}
