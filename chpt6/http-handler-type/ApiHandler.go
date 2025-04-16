package main

import (
	"io"
	"net/http"
)

func ApiHandler(w http.ResponseWriter, r *http.Request, shared *HandlerShared) {
	shared.logger.Println("Handling api request")
	io.WriteString(w, "Hello, world!\n")
}
