package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		listenAddr = ":8080"
	}

	timeoutDuration := 1 * time.Second

	handleUserWithTimeout := http.TimeoutHandler(
		http.HandlerFunc(handleUserAPI),
		timeoutDuration,
		"I ran out of time\n",
	)

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", handlePing)
	mux.Handle("/api/user", handleUserWithTimeout)

	log.Fatal(http.ListenAndServe(listenAddr, mux))
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	log.Println("ping")
	fmt.Fprint(w, "pong")
}

func doSomeWork() {
	time.Sleep(2 * time.Second)
}

func handleUserAPI(w http.ResponseWriter, r *http.Request) {
	log.Println("starting processing the request")
	doSomeWork()
	req, err := http.NewRequestWithContext(
		r.Context(),
		http.MethodGet,
		"http://localhost:8080/ping",
		nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	client := http.Client{}
	log.Println("Outgoing HTTP request")

	resp, err := client.Do(req)
	if err != nil {
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			log.Printf("context deadline exceeded: %v\n", err)
		case errors.Is(err, context.Canceled):
			log.Printf("context canceled: %v\n", err)
		default:
			log.Printf("error making request: %v\n", err)
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	fmt.Fprint(w, string(data))
	log.Println("finished processing the request")
}
