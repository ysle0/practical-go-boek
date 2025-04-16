package middleware

import (
	"fmt"
	"log"
	"net/http"
)

func PanicMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rval := recover(); rval != nil {
					log.Println("panic: ", rval)
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprintf(w, "Unexpected server error")
				}
			}()
			h.ServeHTTP(w, r)
		},
	)
}
