package main

import (
	"net/http"

	"github.com/ysle0/chpt6/middleware-chaining/handler"
	"github.com/ysle0/chpt6/middleware-chaining/handlerWrap"
)

func Route(mx *http.ServeMux, shared *HandlerShared) {
	mx.Handle("/", &handlerWrap.Wrapper[HandlerShared]{
		Shared:  shared,
		Handler: handler.Index[*HandlerShared],
	})

	mx.Handle("/panic", &handlerWrap.Wrapper[HandlerShared]{
		Shared:  shared,
		Handler: handler.Panic[*HandlerShared],
	})
}
