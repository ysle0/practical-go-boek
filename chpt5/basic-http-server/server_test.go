package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	testcases := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "index",
			path:     "/api",
			expected: "Hello, world",
		},
		{
			name:     "healthcheck",
			path:     "/health",
			expected: "OK",
		},
	}

	mx := http.NewServeMux()
	setupHandlers(mx)

	ts := httptest.NewServer(mx)
	defer ts.Close()

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Get(ts.URL + tc.path)
			body, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				log.Fatal(err)
			}

			bodyStr := string(body)
			if bodyStr != tc.expected {
				t.Errorf(
					"expected response body: %s, actual: %s",
					tc.expected, bodyStr)
			}
		})
	}
}
