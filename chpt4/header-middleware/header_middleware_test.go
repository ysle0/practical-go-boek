package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddHeadersMiddleware(t *testing.T) {
	headers := map[string]string{
		"X-Client-Id": "test-client",
		"X-Auth-Hash": "random$string",
	}

	client := newClient(headers)

	ts := newHttpServer()
	defer ts.Close()

	resp, err := client.Get(ts.URL)
	if err != nil {
		t.Fatalf("expected non-nil [AU: \"nil\"-JA] error, actual: %v\n", err)
	}

	for k, v := range headers {
		if resp.Header.Get(k) != headers[k] {
			t.Fatalf("expected header: %s:%s actual: %s:%s",
				k, v,
				k, headers[k],
			)
		}

	}

}

func newHttpServer() *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				for k, v := range r.Header {
					w.Header().Add(k, v[0])
				}
				w.Write([]byte("I am the request header echoing server"))
			},
		),
	)
}
