package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stdout,
			"must specify a HTTP URL to get data from")
		os.Exit(1)
	}
	client := newHttpClientWithTimeout(15 * time.Second)
	body, err := fetchRemoteResource(client, os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stdout, "%v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "%s\n", body)
}

func newHttpClientWithTimeout(d time.Duration) *http.Client {
	return &http.Client{
		Timeout:       d,
		CheckRedirect: redirectPolicyFn,
	}
}

func redirectPolicyFn(r *http.Request, via []*http.Request) error {
	if len(via) >= 1 {
		return errors.New(
			fmt.Sprintf("Attempted redirect to %s", r.URL))
	}
	return nil
}

func fetchRemoteResource(c *http.Client, url string) ([]byte, error) {
	r, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return io.ReadAll(r.Body)

}
