package pkg_register

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type pkgRegisterResult struct {
	ID string `json:"id"`
}

type pkgData struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func registerPkgData(url string, data pkgData) (pkgRegisterResult, error) {
	p := pkgRegisterResult{}
	b, err := json.Marshal(data)
	if err != nil {
		return p, err
	}

	rdr := bytes.NewReader(b)
	r, err := http.Post(url, "application/json", rdr)
	if err != nil {
		return p, err
	}
	defer r.Body.Close()

	resp, err := io.ReadAll(r.Body)
	if err != nil {
		return p, err
	}
	if r.StatusCode != http.StatusOK {
		return p, errors.New(string(resp))
	}
	err = json.Unmarshal(resp, &p)
	return p, err
}
