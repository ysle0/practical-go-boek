package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/decode", decodeHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}

type logLine struct {
	UserIP string `json:"user_ip"`
	Event  string `json:"event"`
}

func decodeHandler(w http.ResponseWriter, r *http.Request) {
	dc := json.NewDecoder(r.Body)
	dc.DisallowUnknownFields()

	var e *json.UnmarshalTypeError
	for {
		var l logLine
		err := dc.Decode(&l)

		if err == io.EOF {
			break
		}
		if errors.As(err, &e) {
			log.Println(e)
			continue
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Printf("%v\n", l)
	}
	fmt.Fprintf(w, "OK - JSON stream decoded")
}
