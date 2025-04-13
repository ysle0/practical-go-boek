package client

import "net/http"

type AddHeadersMiddleware struct {
	http.RoundTripper
	headers map[string]string
}

func (h AddHeadersMiddleware) RoundTrip(r *http.Request) (*http.Response, error) {
	req := r.Clone(r.Context())
	for k, v := range h.headers {
		req.Header.Set(k, v)
	}
	return http.DefaultTransport.RoundTrip(req)
}

func newClient(headers map[string]string) *http.Client {
	h := AddHeadersMiddleware{
		headers: headers,
	}
	return &http.Client{
		Transport: h,
	}
}
