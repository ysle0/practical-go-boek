package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		listenAddr = ":8080"
	}

	waitForShutdownDone := make(chan struct{})
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/users", handleUserAPI)

	s := http.Server{
		Addr:    listenAddr,
		Handler: mux,
	}

	go gracefulShutdown(ctx, &s, waitForShutdownDone)

	err := s.ListenAndServe()
	log.Printf("waiting for shutdown done..")

	<-waitForShutdownDone

	log.Fatal(err)
}

func handleUserAPI(w http.ResponseWriter, r *http.Request) {
	log.Println("starting to process request")
	defer func() {
		log.Println("done processing request")
		r.Body.Close()
	}()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading request body: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println(string(data))
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

func gracefulShutdown(
	ctx context.Context,
	s *http.Server,
	waitForShutdownDone chan<- struct{},
) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigCh

	log.Printf("recv signal: %v, Server shutting down", sig)
	if err := s.Shutdown(ctx); err != nil {
		log.Printf("error during shutdown: %v\n", err)
	}
	waitForShutdownDone <- struct{}{}
}
