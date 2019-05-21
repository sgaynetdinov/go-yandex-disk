package yandexdisk

import (
	"encoding/json"
	"io/ioutil"
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

func (client *Client) get(v interface{}) {
	request, err := http.NewRequest("GET", client.api_url, nil)
	if err != nil {
		panic(err)
	}

	response, err := client.http_client.Do(request)
	if err != nil {
		panic(err)
	}

	text, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()

	json.Unmarshal(text, v)
}
