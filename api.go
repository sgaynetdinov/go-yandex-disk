package yandexdisk

import (
	"net/http"
)

type Client struct {
	api_url     string
	header      *http.Header
	http_client *http.Client
}

func NewClient(token string) *Client {
	header := make(http.Header)
	header.Add("Authorization", "OAuth "+token)
	header.Add("Accept", "application/json")
	header.Add("Content-Type", "application/json")

	return &Client{
		api_url:     "https://cloud-api.yandex.net/v1/disk/",
		header:      &header,
		http_client: new(http.Client),
	}
}
