package client

import (
  "fmt"
  "bytes"
	"net/http"
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

func Get(httpclient *http.Client, url string) (*http.Response, error) {
  resp, err := httpclient.Get(url)
  fmt.Printf("Sending GET request to %s\n", url)
  if err != nil {
    log.Fatalf("Error GET url: %s", err)
    return nil, err
  }

  return resp, nil
}

func Post(httpclient *http.Client, url string, payload string) (*http.Response, error) {
  payloadBuffer := bytes.NewBufferString(payload)
  fmt.Printf("Sending POST request to %s\n", url)
	resp, err := httpclient.Post(url, "application/json", payloadBuffer)
	if err != nil {
    log.Fatalf("Error POST url: %s", err)
    return nil, err
	}
	return resp, nil
}
