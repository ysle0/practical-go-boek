package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiHandler(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	b := new(bytes.Buffer)
	shared := NewHandlerShared(b)

	Index(w, r, &shared)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading response body: %v\n", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf(
			"expected status: OK, actual: %v\n",
			resp.StatusCode,
		)
	}

	expectedBody := "/Index"
	if string(body) != expectedBody {
		t.Errorf(
			"expected body: %v, actual: %v\n",
			expectedBody,
			string(body),
		)
	}
}
