package main

import (
	"io"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, shared *HandlerShared) {
	shared.l.Println("hello! /Index")
	io.WriteString(w, "/Index")
}

func Panic(w http.ResponseWriter, r *http.Request, shared *HandlerShared) {
	panic("panic")
}
