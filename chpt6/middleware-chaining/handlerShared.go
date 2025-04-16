package main

import (
	"io"
	"log"
)

type HandlerShared struct {
	l *log.Logger
}

func NewHandlerShared(w io.Writer) HandlerShared {
	return HandlerShared{
		l: log.New(w, "", log.LstdFlags|log.Lshortfile|log.Ltime|log.Lmicroseconds),
	}
}

func (h HandlerShared) GetLogger() *log.Logger {
	return h.l
}
