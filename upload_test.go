package yandexdisk

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUrlUpload(t *testing.T) {
	var req *http.Request

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
  			"href": "https://uploader1d.dst.yandex.net:443/upload-target/...",
  			"method": "GET",
  			"templated": false
		}`))
		req = r
	}))
	defer ts.Close()

	client := NewClient("YOUR_TOKEN")
	client.api_url = ts.URL
	link, err := client.getUrlUpload("test.txt")

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

	if req.URL.RawQuery != "path=test.txt" {
		t.Error("Invalid", req.URL.RawQuery)
	}

	if req.URL.Path != "/v1/disk/resources/upload" {
		t.Error("Invalid url path", req.URL.Path)
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

func TestUploadFileIfNotFound(t *testing.T) {
	client := NewClient("YOUR_TOKEN")

	err := client.uploadFile("url", "testdata/upload.txt")

	if err == nil {
		t.Error("Error not nil")
	}
}

func TestUploadRequest(t *testing.T) {
	var req *http.Request

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		req = r
	}))
	defer ts.Close()
	client := NewClient("YOUR_TOKEN")

	err := client.uploadFile(ts.URL, "testdata/upload_file.txt")

	if err != nil {
		t.Error("Invalid not nil")
	}

	if req.Method != http.MethodPut {
		t.Error("Method not PUT")
	}
}
