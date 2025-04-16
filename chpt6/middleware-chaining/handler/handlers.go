package handler

import (
	"io"
	"log"
	"net/http"
)

type HandlerConstraint interface {
	GetLogger() *log.Logger
}

func Index[T HandlerConstraint](w http.ResponseWriter, r *http.Request, shared T) {
	shared.GetLogger().Println("hello! /Index")
	io.WriteString(w, "/Index")
}

func Panic[T HandlerConstraint](w http.ResponseWriter, r *http.Request, shared T) {
	panic("panic")
}
