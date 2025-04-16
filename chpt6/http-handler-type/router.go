package main

import "net/http"

func Route(mux *http.ServeMux, shared *HandlerShared) {
	mux.Handle(
		"/health",
		&sharedHandlerWrapper{shared: shared, handler: HealthCheckHandler},
	)
	mux.Handle(
		"/api",
		&sharedHandlerWrapper{shared: shared, handler: ApiHandler},
	)
}
