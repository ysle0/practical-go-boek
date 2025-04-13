package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
	"os"
	"time"
)

func main() {
	cx := context.Background()
	client := newHttpClientWithTimeout(10 * time.Second)

	req, err := newHttpGetRequestWithTrace(cx, os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	for {
		client.Do(req)
		fmt.Println("Request sent\n--------------------------------")
	}

}

func newHttpClientWithTimeout(d time.Duration) *http.Client {
	return &http.Client{
		Timeout: d,
		Transport: &http.Transport{
			// 풀 내에 유지할 최대 유휴 연결 개수.
			// 기본값은 0, 별도로 최대 연결 개수에 제한없음.
			MaxIdleConns: 10,
			// 호스트마다 최대 유휴 연결 개수.
			// 기본값은 DefaultMaxIdleConnsPerHost 이며, Go 1.18 에서는 2
			MaxIdleConnsPerHost: 2,
		},
	}
}

func newHttpGetRequestWithTrace(ctx context.Context, url string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	trace := &httptrace.ClientTrace{
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", dnsInfo)
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("GotConn Info: %+v\n", connInfo)
		},
	}

	cxTrace := httptrace.WithClientTrace(req.Context(), trace)
	req = req.WithContext(cxTrace)
	return req, nil
}
