package yandexdisk

import (
	"net/http"
	"net/http/httptest"
)

func makeServer(body []byte, status int) (*http.Request, *httptest.Server) {
	var req http.Request

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write(body)
		req = *r
	}

	ts := httptest.NewServer(http.HandlerFunc(handler))

	return &req, ts
}
