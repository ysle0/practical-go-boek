package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", longRunningProcessHandler)
	http.ListenAndServe(":8080", mux)
}

func longRunningProcess(logw *io.PipeWriter) {
	defer logw.Close()
	fmt.Println("got request streaming")

	for i := range 20 {
		fmt.Fprintf(logw,
			`{"id":%d,"user_ip":"172.121.19.21","event":"click_on_add_cart"`,
			i)
		fmt.Printf("[%d/20] processing ...\n", i+1)
		fmt.Fprintln(logw)
		time.Sleep(200 * time.Millisecond)
	}
}

func longRunningProcessHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	done := make(chan struct{})
	logr, logw := io.Pipe()
	go longRunningProcess(logw)
	go processStreamer(logr, w, done)
	<-done
}

func processStreamer(
	logr *io.PipeReader,
	w http.ResponseWriter,
	done chan struct{},
) {
	buf := make([]byte, 512)
	f, flushSupported := w.(http.Flusher)
	defer logr.Close()

	header := w.Header()
	header.Set("Content-Type", "text/plain")
	header.Set("X-Content-Type-Options", "nosniff")

	for {
		n, err := logr.Read(buf)
		if err == io.EOF {
			break
		}

		w.Write(buf[:n])
		if flushSupported {
			f.Flush()
		}
	}
	done <- struct{}{}
}
