package main

import (
	"io"
	"net/http"

	"github.com/ysle0/chpt6/handlerWrap"
)

func main() {
	mx := http.NewServeMux()
	handlerShared := NewHandlerShared()

	mx.Handle("/", &handlerWrap.Wrapper[HandlerShared]{
		Shared:     &handlerShared,
		Handler:    Index,
		UseMeasure: true,
	})

	err := http.ListenAndServe(":8080", mx)
	if err != nil {
		handlerShared.l.Fatal(err)
	}
}

func Index(w http.ResponseWriter, r *http.Request, shared *HandlerShared) {
	shared.l.Println("hello! /Index")
	io.WriteString(w, "/Index")
}
