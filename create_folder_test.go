package yandexdisk

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateFolder(t *testing.T) {
	var req *http.Request

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{
  			"href": "https://cloud-api.yandex.net/v1/disk/resources?path", 
  			"method": "GET",
  			"templated": false
		}`))
		req = r
	}))
	defer ts.Close()

	client := NewClient("YOUR_TOKEN")
	client.api_url = ts.URL
	err := client.CreateFolder("/Music/2pac")

	if err != nil {
		t.Error("Error is not nil")
	}

	if req.Method != http.MethodPut {
		t.Error("Invalid method")
	}

	if req.URL.RawQuery != "path=%2FMusic%2F2pac" {
		t.Error("Invalid", req.URL.RawQuery)
	}

	if req.URL.Path != "/v1/disk/resources" {
		t.Error("Invalid url path", req.URL.Path)
	}
}

func TestCreateFolderAddStartSlash(t *testing.T) {
	var req *http.Request

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		req = r
	}))
	defer ts.Close()

	client := NewClient("YOUR_TOKEN")
	client.api_url = ts.URL
	client.CreateFolder("Music/2pac")

	if req.URL.RawQuery != "path=%2FMusic%2F2pac" {
		t.Error("Invalid", req.URL.RawQuery)
	}
}

func TestCreateFolderError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer ts.Close()

	client := NewClient("YOUR_TOKEN")
	client.api_url = ts.URL
	err := client.CreateFolder("Music/2pac")

	if err == nil {
		t.Error()
	}
}
