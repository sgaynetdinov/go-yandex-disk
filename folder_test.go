package yandexdisk

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMkdir(t *testing.T) {
	var req *http.Request

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{
  			"href": "https://cloud-api.yandex.net/v1/disk/resources?path", 
  			"method": "GET",
  			"templated": false
		}`)) req = r
	}))
	defer ts.Close()

	client := NewClient("YOUR_TOKEN")
	client.apiURL = ts.URL
	err := client.Mkdir("/Music/2pac")

	if err != nil {
		t.Error("Error is not nil")
	}

	if req.Method != http.MethodPut {
		t.Error("Invalid method")
	}

	if req.URL.RawQuery != "path=disk%3A%2FMusic%2F2pac" {
		t.Error("Invalid", req.URL.RawQuery)
	}

	if req.URL.Path != "/v1/disk/resources" {
		t.Error("Invalid url path", req.URL.Path)
	}
}

func TestMkdirAddStartSlash(t *testing.T) {
	var req *http.Request

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		req = r
	}))
	defer ts.Close()

	client := NewClient("YOUR_TOKEN")
	client.apiURL = ts.URL
	client.Mkdir("Music/2pac")

	if req.URL.RawQuery != "path=disk%3A%2FMusic%2F2pac" {
		t.Error("Invalid", req.URL.RawQuery)
	}
}

func TestMkdirError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer ts.Close()

	client := NewClient("YOUR_TOKEN")
	client.apiURL = ts.URL
	err := client.Mkdir("Music/2pac")

	if err == nil {
		t.Error()
	}
}

func TestIsExistsFolder_1(t *testing.T) {
	var req *http.Request
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{
	        "message": "Не удалось найти запрошенный ресурс.",
			"description": "Resource not found.",
		    "error": "DiskNotFoundError"
		}`))
		req = r
	}))
	defer ts.Close()

	client := NewClient("YOUR_TOKEN")
	client.apiURL = ts.URL
	isExists, err := client.IsExistsFolder("Music/2pac")

	if err != nil {
		t.Error(err)
	}

	if isExists != false {
		t.Error()
	}

	if req.URL.RawQuery != "path=disk%3A%2FMusic%2F2pac" {
		t.Error("Invalid", req.URL.RawQuery)
	}
}

func TestIsExistsFolder_2(t *testing.T) {
	var req *http.Request
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
		req = r
	}))
	defer ts.Close()

	client := NewClient("YOUR_TOKEN")
	client.apiURL = ts.URL
	isExists, err := client.IsExistsFolder("Music/2pac")

	if err != nil {
		t.Error()
	}

	if isExists != true {
		t.Error()
	}

	if req.URL.RawQuery != "path=disk%3A%2FMusic%2F2pac" {
		t.Error("Invalid", req.URL.RawQuery)
	}
}

func TestIsExistsFolder_3(t *testing.T) {
	var req *http.Request
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{
	        "message": "",
			"description": "",
		    "error": "Disk"
		}`))
		req = r
	}))
	defer ts.Close()

	client := NewClient("YOUR_TOKEN")
	client.apiURL = ts.URL
	isExists, err := client.IsExistsFolder("Music/2pac")

	if err == nil {
		t.Error()
	}

	if isExists != false {
		t.Error()
	}

	if req.URL.RawQuery != "path=disk%3A%2FMusic%2F2pac" {
		t.Error("Invalid", req.URL.RawQuery)
	}
}
