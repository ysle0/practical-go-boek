package main

import (
	"testing"

	"github.com/ysle0/shared"
)

func TestFetchRemoteResource(t *testing.T) {
	ts := shared.StartTestHttpServer()
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
