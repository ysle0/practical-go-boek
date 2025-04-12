package pkg_register_data

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type pkgData struct {
	Name     string
	Version  string
	Filename string
	Bytes    io.Reader
}

type pkgRegisterResult struct {
	ID       string `json:"id"`
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
}

func registerPkgData(url string, data pkgData) (pkgRegisterResult, error) {
	payload, contentType, err := newMultipartMessage(data)
	if err != nil {
		return pkgRegisterResult{}, err
	}
	rdr := bytes.NewReader(payload)
	r, err := http.Post(url, contentType, rdr)
	if err != nil {
		return pkgRegisterResult{}, err
	}
	defer r.Body.Close()

	resp, err := io.ReadAll(r.Body)
	if err != nil {
		return pkgRegisterResult{}, err
	}

	var p pkgRegisterResult
	err = json.Unmarshal(resp, &p)
	return p, err
}
