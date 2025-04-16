package main

import (
	"net/http"
)

type sharedHandlerWrapper struct {
	shared  *HandlerShared
	handler func(w http.ResponseWriter, r *http.Request, shared *HandlerShared)
}

func (a sharedHandlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.handler(w, r, a.shared)
}
