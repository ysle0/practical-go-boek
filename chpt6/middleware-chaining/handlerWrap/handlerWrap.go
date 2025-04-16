package handlerWrap

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type LoggerConstraint interface {
	GetLogger() *log.Logger
}

type Wrapper[T LoggerConstraint] struct {
	Shared  *T
	Handler func(w http.ResponseWriter, r *http.Request, shared *T)
}

func (h Wrapper[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Handler(w, r, h.Shared)
}

func LoggingMiddleware(h http.Handler) http.Handler {
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

func PanicMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rval := recover(); rval != nil {
					log.Println("panic: ", rval)
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprintf(w, "Unexpected server error")
				}
			}()
			h.ServeHTTP(w, r)
		},
	)
}
