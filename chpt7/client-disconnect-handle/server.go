package main

import (
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

	timeoutDuration := 30 * time.Second

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", handlePing)
	mux.Handle("/api/user",
		http.TimeoutHandler(
			http.HandlerFunc(handleUserAPI),
			timeoutDuration,
			"I ran out of time\n"),
	)

	log.Fatal(http.ListenAndServe(listenAddr, mux))
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	log.Println("ping: Got a request, processing it")
	time.Sleep(3 * time.Second)
	fmt.Fprint(w, "pong")
}

func doHeavyWork([]byte) {
	time.Sleep(5 * time.Second)
}

func handleUserAPI(w http.ResponseWriter, r *http.Request) {
	done := make(chan bool)
	defer close(done)

	log.Println("starting to process request")

	req, err := http.NewRequestWithContext(r.Context(),
		http.MethodGet,
		"http://localhost:8080/ping",
		nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	log.Println("Outgoing HTTP request")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("error making request: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	log.Println("Processing the response I got")

	go func() {
		doHeavyWork(data)
		done <- true
	}()

	select {
	case <-done:
		log.Println("doHeavyWork() done: Continuing request processing")
	case <-r.Context().Done():
		log.Printf("Aborting further request processing: %v\n", r.Context().Err())
		return
	}

	fmt.Fprint(w, string(data))
	log.Println("finished processing the request")
}
