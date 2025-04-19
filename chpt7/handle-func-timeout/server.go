package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	timeoutDuration := 10 * time.Second
	hTimeout := http.TimeoutHandler(
		http.HandlerFunc(handleUserAPI),
		timeoutDuration,
		"I ran out of time\n",
	)

	mux := http.NewServeMux()
	mux.Handle("/api/user", hTimeout)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handleUserAPI(w http.ResponseWriter, r *http.Request) {
	time.Sleep(15 * time.Second)
	w.WriteHeader(http.StatusOK)

	log.Println(
		"before continuing, I will check if the timeout has already expired",
	)

	if r.Context().Err() != nil {
		err := r.Context().Err()
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			log.Printf("context deadline exceeded: %v\n", err)
		case errors.Is(err, context.Canceled):
			log.Printf("context canceled: %v\n", err)
		default:
			log.Printf("aborting further processing: %v\n", err)
		}
		return
	}

	fmt.Fprint(w, "I'm a slow response\n")
	log.Println("I finished processing the request")
}
