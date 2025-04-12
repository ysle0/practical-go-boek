package pkgquery

import (
	"encoding/json"
	"io"
	"net/http"
)

type pkgData struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func fetchPkgData(url string) ([]pkgData, error) {
	var pkgs []pkgData
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if r.Header.Get("Content-Type") != "application/json" {
		return pkgs, nil
	}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return pkgs, err
	}

	err = json.Unmarshal(data, &pkgs)
	return pkgs, err
}
