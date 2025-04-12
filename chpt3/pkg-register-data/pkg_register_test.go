package pkg_register_data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"strings"
	"testing"
)

type Form struct {
	Value map[string][]string
	File  map[string][]*FileHeader
}

type FileHeader struct {
	Filename string
	Header   textproto.MIMEHeader
	Size     int64
}

func TestRegisterPkg(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(postHandlePkgData))
	defer ts.Close()

	data := strings.NewReader("data")
	p := pkgData{
		Name:     "myPackage",
		Version:  "0.1",
		Filename: "myFile-0.1.tar.gz",
		Bytes:    data,
	}

	res, err := registerPkgData(ts.URL, p)
	if err != nil {
		t.Fatal(err)
	}

	if res.ID != fmt.Sprintf("%s-%s", p.Name, p.Version) {
		t.Errorf(
			"expected package ID: %s-%s, actual: %s",
			p.Name, p.Version, res.ID)
	}

	if res.Filename != p.Filename {
		t.Errorf("expected filename: %s, actual: %s",
			p.Filename, res.Filename)

		if res.Size != 4 {
			t.Errorf("expected size: 4, actual: %d", res.Size)
		}
	}

}

func postHandlePkgData(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		d := pkgRegisterResult{}
		err := r.ParseMultipartForm(5000)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		mForm := r.MultipartForm
		f := mForm.File["filedata"][0]
		d.ID = fmt.Sprintf(
			"%s-%s",
			mForm.Value["name"][0],
			mForm.Value["version"][0],
		)
		d.Filename = f.Filename
		d.Size = f.Size
		data, err := json.Marshal(d)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, string(data))

	default:
		http.Error(w, "Invalid HTTP request method",
			http.StatusMethodNotAllowed)
	}
}
