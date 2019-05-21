package yandexdisk

import (
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
