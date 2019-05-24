package yandexdisk

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUrlUpload(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
  			"href": "https://uploader1d.dst.yandex.net:443/upload-target/...",
  			"method": "GET",
  			"templated": false
		}`))
	}))
	defer ts.Close()

	client := NewClient("YOUR_TOKEN")
	client.api_url = ts.URL
	link, err := client.getUrlUpload("/")

	if err != nil {
		t.Error("Error is not nil")
	}

	if link.Href != "https://uploader1d.dst.yandex.net:443/upload-target/..." {
		t.Error("Invalid Href")
	}

	if link.Method != "GET" {
		t.Error("Invalid Method")
	}

	if link.Templated != false {
		t.Error("Invalid Templated")
	}
}

func TestGetUrlUploadWithError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusRequestEntityTooLarge)
		w.Write([]byte(`{
			"description": "resource already exists",
			"error": "PlatformResourceAlreadyExists"
		}`))
	}))
	defer ts.Close()

	client := NewClient("YOUR_TOKEN")
	client.api_url = ts.URL
	link, err := client.getUrlUpload("/")

	if err == nil {
		t.Error("Error is nil")
	}

	if link != nil {
		t.Error("Link not nil")
	}
}