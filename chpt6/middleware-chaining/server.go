package main

import (
	"io"
	"net/http"

	"github.com/ysle0/chpt6/handlerWrap"
	"github.com/ysle0/chpt6/middleware"
)

func main() {
	mx := http.NewServeMux()
	handlerShared := NewHandlerShared()

	route(mx, &handlerShared)

	mmx := middleware.Chain(mx,
		middleware.LogMiddleware,
		middleware.PanicMiddleware,
	)

	err := http.ListenAndServe(":8080", mmx)
	if err != nil {
		handlerShared.l.Fatal(err)
	}
}

func route(mx *http.ServeMux, shared *HandlerShared) {
	mx.Handle("/", &handlerWrap.Wrapper[HandlerShared]{
		Shared:  shared,
		Handler: Index,
	})
	mx.Handle("/panic", &handlerWrap.Wrapper[HandlerShared]{
		Shared:  shared,
		Handler: Panic,
	})
}

func Index(w http.ResponseWriter, r *http.Request, shared *HandlerShared) {
	shared.l.Println("hello! /Index")
	io.WriteString(w, "/Index")
}

func Panic(w http.ResponseWriter, r *http.Request, shared *HandlerShared) {
	panic("panic")
}
