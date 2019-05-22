package yandexdisk

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient("YOUR_TOKEN")

	if client.api_url != "https://cloud-api.yandex.net/v1/disk/" {
		t.Error("Invalid api url")
	}

	if client.header.Get("Authorization") != "OAuth YOUR_TOKEN" {
		t.Error("Invalid oauth token")
	}

	if client.header.Get("Accept") != "application/json" {
		t.Error("Invalid Accept")
	}

	if client.header.Get("Content-Type") != "application/json" {
		t.Error("Invalid Content-Type")
	}
}

func TestDo(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(disk_json))
	}))
	defer ts.Close()

	client := NewClient("YOUR_TOKEN")
	response, _, err := client.do("GET", ts.URL)

	if err != nil {
		t.Error("Error is not nil")
	}

	if response.Request.Header.Get("Authorization") != "OAuth YOUR_TOKEN" {
		t.Error("Invalid oauth token")
	}

	if response.Request.Header.Get("Accept") != "application/json" {
		t.Error("Invalid Accept")
	}

	if response.Request.Header.Get("Content-Type") != "application/json" {
		t.Error("Invalid Content-Type")
	}
}

func TestDoIfStatusCodeNot200(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{
  			"description": "resource already exists",
  			"error": "PlatformResourceAlreadyExists"
		}`))
	}))
	defer ts.Close()

	client := NewClient("YOUR_TOKEN")
	_, _, err := client.do("GET", ts.URL)

	if err == nil {
		t.Error("Error is nil")
	}

	if err.Error() != "resource already exists - PlatformResourceAlreadyExists" {
		t.Error("Invalid error message")
	}
}
