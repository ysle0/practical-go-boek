package data_downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr,
			"must specify a HTTP URL to get data from")
		os.Exit(1)
	}

	body, err := fetchRemoteResource(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stdout, "%#v\n", err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, "responsed:")
	fmt.Fprintf(os.Stdout, "%s\n", body)
}

func fetchRemoteResource(url string) ([]byte, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return io.ReadAll(r.Body)
}
