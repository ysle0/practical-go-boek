package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestFetchBadRemoteResourceV2(t *testing.T) {
	shutdownServer := make(chan struct{})
	ts := newBadTestHttpServerV2(shutdownServer)
	defer ts.Close()
	defer func() {
		shutdownServer <- struct{}{}
	}()

	client := newHttpClient(200 * time.Millisecond)
	_, err := fetchRemoteResource(client, ts.URL)
	if err == nil {
		t.Fatal("expected error, got none")
	}

	if !strings.Contains(err.Error(), "Client.Timeout exceeded while awaiting headers") {
		t.Fatalf("expected error to contain: Client.Timeout exceeded while awaiting headers, Got: %v",
			err.Error())
	}
}

func newBadTestHttpServerV2(shutdownServer <-chan struct{}) *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				<-shutdownServer
				fmt.Fprint(w, "Hello World")
			}))

}
