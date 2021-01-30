package yandexdisk

import (
	"bufio"
	"net/http"
	"net/http/httptest"
	"os"
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
	client.apiURL = ts.URL
	link, err := client.getUrlUpload("test.txt", false)

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

	if req.URL.RawQuery != "overwrite=false&path=disk%3A%2Ftest.txt" {
		t.Error("Invalid", req.URL.RawQuery)
	}

	if req.URL.Path != "/v1/disk/resources/upload" {
		t.Error("Invalid url path", req.URL.Path)
	}
}

func TestGetUrlUploadOverwrite(t *testing.T) {
	var req *http.Request

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
		req = r
	}))
	defer ts.Close()

	client := NewClient("YOUR_TOKEN")
	client.apiURL = ts.URL
	_, err := client.getUrlUpload("test.txt", true)

	if err != nil {
		t.Error("Error is not nil")
	}

	if req.URL.RawQuery != "overwrite=true&path=disk%3A%2Ftest.txt" {
		t.Error("Invalid", req.URL.RawQuery)
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
	client.apiURL = ts.URL
	link, err := client.getUrlUpload("/", false)

	if err == nil {
		t.Error("Error is nil")
	}

	if link != nil {
		t.Error("Link not nil")
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

	file, _ := os.Open("testdata/upload_file.txt")
	err := client.uploadFile(ts.URL, bufio.NewReader(file))

	if err != nil {
		t.Error("Invalid not nil")
	}

	if req.Method != http.MethodPut {
		t.Error("Method not PUT")
	}
}
