package client

import (
  "bytes"
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

func Get(httpclient *http.Client, url string) (string, error) {
  resp, err := httpclient.Get(url)
  if err != nil {
    log.Fatalf("Error GET url: %s", err)
    return "", err
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatalf("Error reading response body on get request url: %s", err)
    return "", err
  }

  return string(body), nil
}

func Post(httpclient *http.Client, url string, payload string) (string, error) {
  payloadBuffer := bytes.NewBufferString(payload)
	resp, err := httpclient.Post(url, "application/json", payloadBuffer)
	if err != nil {
    log.Fatalf("Error POST url: %s", err)
    return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
    log.Fatalf("Error reading response body on POST url: %s", err)
    return "", err
	}

	return string(body), nil
}
