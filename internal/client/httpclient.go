package client

import (
	"net/http"
	"io/ioutil"
	"time"
  "log"
)

type HTTPClientConfig struct {
	Timeout time.Duration
}


func NewHTTPClient(config HTTPClientConfig) *http.Client {
	return &http.Client{
		Timeout: config.Timeout,
	}
}

func (c *http.Client) Get(url string) (string, error) {
  resp, err := c.Get(url)
  if err != nil {
    log.Fatal("Error GET url: %s", err)
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal("Error reading response body on get request url: %s", err)
  }

  return string(body)
}

func (c *http.Client) Post(url string, body string) (string, error) {
	resp, err := c.Post(url, "application/json", body)
	if err != nil {
    log.Fatal("Error POST url: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
    log.Fatal("Error reading response body on POST url: %s", err)
	}

	return string(body), nil
}
