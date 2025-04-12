package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestFetchBadRemoteResourceV1(t *testing.T) {
	ts := newBadTestHttpServerV1()
	defer ts.Close()

	client := newHttpClient(200 * time.Millisecond)
	_, err := fetchRemoteResource(client, ts.URL)
	if err == nil {
		t.Fatal("expected error, got none")
	}

	exp := "context deadline exceeded"
	if !strings.Contains(err.Error(), exp) {
		t.Fatalf("expected error to contain: "+exp+", Got: %v",
			err.Error())
	}
}

func newBadTestHttpServerV1() *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(60 * time.Second)
				fmt.Fprint(w, "Hello World")
			}))

}
