package pkg_register_data

import (
	"bytes"
	"io"
	"mime/multipart"
)

func newMultipartMessage(data pkgData) ([]byte, string, error) {
	body := new(bytes.Buffer)
	mw := multipart.NewWriter(body)
	defer mw.Close()

	part, err := mw.CreateFormField("name")
	if err != nil {
		return nil, "", err
	}
	part.Write([]byte(data.Name))

	part, err = mw.CreateFormField("version")
	if err != nil {
		return nil, "", err
	}
	part.Write([]byte(data.Version))

	part, err = mw.CreateFormFile("filedata", data.Filename)

	_, err = io.Copy(part, data.Bytes)
	if err != nil {
		return nil, "", err
	}

	return body.Bytes(), mw.FormDataContentType(), nil
}
