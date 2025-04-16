package main

import (
	"io"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request, shared *HandlerShared) {
	if r.Method != http.MethodGet {
		http.Error(w,
			"Method not allowed",
			http.StatusMethodNotAllowed)
		return
	}

	shared.logger.Println("Handling healthcheck request")
	io.WriteString(w, "OK")
}
