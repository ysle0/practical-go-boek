package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	addr := ":8080"

	mux := http.NewServeMux()
	mux.HandleFunc("/api/user", handleUserAPI)

	s := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Printf("Listening on %s\n", addr)
	err := s.ListenAndServe()
	log.Fatal(err)
}

func handleUserAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		log.Println("starting to process request")
		defer r.Body.Close()

		data, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("error reading request body: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println(string(data))
		fmt.Fprintln(w, "Hello world!")
		log.Println("finished processing the request")
	}
}
