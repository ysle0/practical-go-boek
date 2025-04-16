package middleware

import (
	"log"
	"net/http"
	"time"
)

func LogMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			h.ServeHTTP(w, r)
			endTime := time.Since(startTime).Seconds()
			log.Printf(
				"path=%s, method=%s, duration=%f",
				r.URL.Path, r.Method, endTime)
		},
	)
}
