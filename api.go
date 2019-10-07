package yandexdisk

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	apiURL     string
	header     *http.Header
	httpClient *http.Client
}

func NewClient(token string) *Client {
	header := make(http.Header)
	header.Add("Authorization", "OAuth "+token)
	header.Add("Accept", "application/json")
	header.Add("Content-Type", "application/json")

	return &Client{
		apiURL:     "https://cloud-api.yandex.net:443",
		header:     &header,
		httpClient: new(http.Client),
	}
}

func (client *Client) do(method string, path string) (*[]byte, error) {
	request, err := http.NewRequest(method, path, nil)
	request.Header = *client.header
	if err != nil {
		panic(err)
	}

	response, err := client.httpClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	text, _ := ioutil.ReadAll(response.Body)

	statusCode := response.StatusCode
	if (statusCode != http.StatusOK) && (statusCode != http.StatusCreated) {
		var err yaError
		json.Unmarshal(text, &err)
		return nil, &err
	}

	return &text, nil
}

func (client *Client) get(v interface{}, path string, params *url.Values) error {
	var url string

	url = client.apiURL

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

func (client *Client) put(v interface{}, path string, params *url.Values) error {
	url := client.apiURL

	if path != "" {
		url += path
	}

	if params != nil {
		url += "?" + params.Encode()
	}

	text, err := client.do(http.MethodPut, url)

	if err != nil {
		return err
	}

	json.Unmarshal(*text, v)
	return nil
}
