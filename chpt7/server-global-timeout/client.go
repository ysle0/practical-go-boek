package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{}
	logr, logw := io.Pipe()
	go writeSlowly(logw, 1*time.Second)
	r, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"http://localhost:8080/api/users",
		logr,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("sending request to server\n")

	resp, err := client.Do(r)
	if err != nil {
		log.Fatalf("error sending request: %v\n", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(data))
}

func writeSlowly(pw *io.PipeWriter, delay time.Duration) {
	defer pw.Close()

	for i := range 10 {
		fmt.Fprintf(pw, fmt.Sprintf("#%d Hello \n", i))
		time.Sleep(delay)

	}
}
