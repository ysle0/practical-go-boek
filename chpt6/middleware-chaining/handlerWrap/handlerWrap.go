package handlerWrap

import (
	"log"
	"net/http"
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
