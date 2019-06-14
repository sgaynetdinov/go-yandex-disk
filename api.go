package yandexdisk

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
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
		api_url:     "https://cloud-api.yandex.net:443",
		header:      &header,
		http_client: new(http.Client),
	}
}

func (client *Client) do(method string, path string) (*[]byte, error) {
	request, err := http.NewRequest(method, path, nil)
	request.Header = *client.header
	if err != nil {
		panic(err)
	}

	response, err := client.http_client.Do(request)
	if err != nil {
		panic(err)
	}

	text, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()

	if response.StatusCode != http.StatusOK {
		var err yaError
		json.Unmarshal(text, &err)
		return nil, &err
	}

	return &text, nil
}

func (client *Client) get(v interface{}, path string, params *url.Values) error {
	var url string

	url = client.api_url

	if path != "" {
		url += path
	}

	if params != nil {
		url += "?" + params.Encode()
	}

	text, err := client.do(http.MethodGet, url)

	if err != nil {
		return err
	}

	json.Unmarshal(*text, v)
	return nil
}
