package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFetchRemoteResource(t *testing.T) {
	ts := newTestHttpServer()
	defer ts.Close()

	client := newHttpClient(20 * time.Millisecond)
	data, err := fetchRemoteResource(client, ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	exp := "Hello World"
	actual := string(data)

	if exp != actual {
		t.Errorf("expected response: %s, actual: %s", exp, actual)
	}
}

func newTestHttpServer() *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				io.WriteString(w, "Hello World")
			}))
}
