package yandexdisk

import (
	"net/http"
	"net/http/httptest"
)

func makeServer(body []byte, status int) (*http.Request, *Client) {
	var req http.Request

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write(body)
		req = *r
	}

	ts := httptest.NewUnstartedServer(http.HandlerFunc(handler))
	ts.Start()
	ts.URL = JoinURL(ts.URL, VERSION_API)

	client := NewClient("YOUR_TOKEN")
	client.apiURL = ts.URL

	return &req, client
}
