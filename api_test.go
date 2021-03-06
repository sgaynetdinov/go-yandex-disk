package yandexdisk

import (
	"net/http"
	"net/url"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient("YOUR_TOKEN")

	if client.apiURL != "https://cloud-api.yandex.net:443" {
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
	req, ts := makeServer(diskJSON, http.StatusOK)
	defer ts.Close()

	client := NewClient("YOUR_TOKEN")
	client.apiURL = ts.URL
	_, err := client.do(http.MethodGet, "", &url.Values{})

	if err != nil {
		t.Error("Error is not nil")
	}

	if req.Header.Get("Authorization") != "OAuth YOUR_TOKEN" {
		t.Error("Invalid oauth token")
	}

	if req.Header.Get("Accept") != "application/json" {
		t.Error("Invalid Accept")
	}

	if req.Header.Get("Content-Type") != "application/json" {
		t.Error("Invalid Content-Type")
	}
}

func TestDoIfFail(t *testing.T) {
	cases := []struct {
		name         string
		responseBody []byte
		statusCode   int
		expected     string
	}{
		{"Status code equal 500", []byte(`{"description": "resource already exists", "error": "PlatformResourceAlreadyExists"}`), http.StatusInternalServerError, "resource already exists - PlatformResourceAlreadyExists"},
		{"Invalid JSON", []byte(`{{"description": "resource already exists", "error": "PlatformResourceAlreadyExists"}`), http.StatusInternalServerError, " - JSON invalid"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, ts := makeServer(tc.responseBody, tc.statusCode)
			defer ts.Close()

			client := NewClient("YOUR_TOKEN")
			client.apiURL = ts.URL
			_, err := client.do(http.MethodGet, "", &url.Values{})

			if err == nil {
				t.Error("Error is nil")
			}

			if err.Error() != tc.expected {
				t.Error("Invalid error message")
			}
		})
	}
}
