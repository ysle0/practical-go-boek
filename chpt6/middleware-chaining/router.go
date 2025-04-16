package main

import (
	"net/http"

	"github.com/ysle0/chpt6/handlerWrap"
)

func Route(mx *http.ServeMux, shared *HandlerShared) {
	mx.Handle("/", &handlerWrap.Wrapper[HandlerShared]{
		Shared:  shared,
		Handler: Index,
	})

	mx.Handle("/panic", &handlerWrap.Wrapper[HandlerShared]{
		Shared:  shared,
		Handler: Panic,
	})
}
