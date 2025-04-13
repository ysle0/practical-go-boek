package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stdout,
			"must specify a HTTP URL to get data from\n")
		os.Exit(1)
	}

	transport := LoggingClient{}
	l := log.New(os.Stdout, "header_middleware", log.LstdFlags)
	transport.log = l

	client := newHttpClientWithTimeout(15 * time.Second)
	client.Transport = transport

	body, err := fetchRemoteResource(client, os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stdout, "error fetching resource: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "response size of bytes: %v\n", len(body))
}
