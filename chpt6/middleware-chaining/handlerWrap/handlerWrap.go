package handlerWrap

import (
	"log"
	"net/http"
	"time"
)

// Define an interface for the required logger method(s).
type LoggerConstraint interface {
	GetLogger() *log.Logger
}

// Use the new LoggerConstraint for the type parameter T.
type Wrapper[T LoggerConstraint] struct {
	Shared     *T
	UseMeasure bool
	Handler    func(w http.ResponseWriter, r *http.Request, shared *T)
}

func (h Wrapper[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var startTime time.Time
	if h.UseMeasure {
		startTime = time.Now()
	}
	h.Handler(w, r, h.Shared)
	if h.UseMeasure {
		endTime := time.Since(startTime).Seconds()

		// Call Printf directly on the dereferenced pointer *T,
		// which is guaranteed to have the Printf method due to the constraint.
		// Add a nil check for safety.
		if h.Shared != nil {
			(*h.Shared).GetLogger().Printf(
				"path=%s, method=%s, duration=%f",
				r.URL.Path, r.Method, endTime)
		}
	}
}
