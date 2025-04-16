package shared

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func StartTestHttpServer() *httptest.Server {
	t := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Hello World")
		}))
	return t
}

func StartTestHttpServerWithJson(json string) *httptest.Server {
	t := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, json)
		}))
	return t
}
