package main

import (
	"net/http"
)

func main() {
	shared := NewHandlerShared()
	mux := http.NewServeMux()

	Route(mux, &shared)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		shared.logger.Fatal(err)
	}
}
