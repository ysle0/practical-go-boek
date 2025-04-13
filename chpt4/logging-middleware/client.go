package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

type LoggingClient struct {
	log *log.Logger
}

// RoundTrip implements http.RoundTripper.
func (c LoggingClient) RoundTrip(r *http.Request) (*http.Response, error) {
	c.log.Printf("request [%s] %s(%s)\n",
		r.Method, r.URL, r.Proto)
	resp, err := http.DefaultTransport.RoundTrip(r)
	c.log.Printf("response [%s] %s(%s)\n",
		r.Method, r.URL, r.Proto)
	return resp, err
}

func newHttpClientWithTimeout(d time.Duration) *http.Client {
	return &http.Client{
		Timeout: d,
	}
}

func fetchRemoteResource(client *http.Client, url string) ([]byte, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
