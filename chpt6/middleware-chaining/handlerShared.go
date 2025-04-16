package main

import (
	"log"
	"os"
)

type HandlerShared struct {
	l *log.Logger
}

func NewHandlerShared() HandlerShared {
	return HandlerShared{
		l: log.New(
			os.Stdout,
			"",
			log.LstdFlags|log.Lshortfile|log.Ltime|log.Lmicroseconds,
		),
	}
}

func (h HandlerShared) GetLogger() *log.Logger {
	return h.l
}
