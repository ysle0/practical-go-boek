package main

import (
	"net/http"
	"os"

	"github.com/ysle0/chpt6/middleware-chaining/middleware"
)

func main() {
	mx := http.NewServeMux()
	handlerShared := NewHandlerShared(os.Stdout)

	Route(mx, &handlerShared)

	mmx := middleware.Chain(mx,
		middleware.LogMiddleware,
		middleware.PanicMiddleware,
	)

	err := http.ListenAndServe(":8080", mmx)
	if err != nil {
		handlerShared.l.Fatal(err)
	}
}
