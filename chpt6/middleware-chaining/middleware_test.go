package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ysle0/chpt6/middleware-chaining/middleware"
)

func TestMiddlewares(t *testing.T) {
	buf := new(bytes.Buffer)
	shared := NewHandlerShared(buf)
	mx := http.NewServeMux()
	Route(mx, &shared)

	mmx := middleware.Chain(mx,
		middleware.LogMiddleware,
		middleware.PanicMiddleware,
	)

	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	mmx.ServeHTTP(w, r)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading response body: %v\n", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status: OK, actual: %v\n", resp.StatusCode)
	}

	expectedBody := "/Index"
	if string(body) != expectedBody {
		t.Errorf("expected body: %v, actual: %v\n", expectedBody, string(body))
	}

	r = httptest.NewRequest("GET", "/panic", nil)
	w = httptest.NewRecorder()
	mmx.ServeHTTP(w, r)

	resp = w.Result()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading response body: %v\n", err)
	}

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected status: InternalServerError, actual: %v\n", resp.StatusCode)
	}

	expectedBody = "Unexpected server error"
	if string(body) != expectedBody {
		t.Errorf("expected body: %v, actual: %v\n", expectedBody, string(body))
	}
}
