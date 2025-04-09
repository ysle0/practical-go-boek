package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchRemoteResource(t *testing.T) {
	ts := startTestHttpServer()
	defer ts.Close()
	expected := "Hello World"
	actual, err := fetchRemoteResource(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	if expected != string(actual) {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func startTestHttpServer() *httptest.Server {
	t := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello World")
		}))
	return t
}
