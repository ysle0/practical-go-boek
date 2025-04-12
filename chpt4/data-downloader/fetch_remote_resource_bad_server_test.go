package data_downloader

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func newBadTestHttpServer() *httptest.Server {
	handler := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(60 * time.Second)
		io.WriteString(w, "Hello World")
	}
	return httptest.NewServer(http.HandlerFunc(handler))
}

func TestFetchBadRemoteResource(t *testing.T) {
	ts := newBadTestHttpServer()
	defer ts.Close()

	data, err := fetchRemoteResource(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	exp := "Hello World"
	actual := string(data)

	if actual != exp {
		t.Errorf("expected response: %s, actual: %s",
			exp, actual)
	}
}
