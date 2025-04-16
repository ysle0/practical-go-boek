package main

import (
	"log"
	"os"
)

func NewHandlerShared() HandlerShared {
	return HandlerShared{
		logger: log.New(
			os.Stdout,
			"[AppCfg]",
			log.Lmsgprefix|log.Ldate|log.Ltime|log.Lshortfile|log.Lmicroseconds,
		),
	}
}
