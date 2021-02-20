package yandexdisk

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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

func (client *Client) do(method string, path string, params *url.Values) (*[]byte, error) {
	url := client.apiURL

	if path != "" {
		url += path
	}

	if params != nil && params.Get("path") != "" {
		name := params.Get("path")
		if !strings.HasPrefix(name, "/") {
			name = "/" + name
		}

		params.Set("path", "disk:"+name)
	}

	if params != nil {
		url += "?" + params.Encode()
	}

	request, err := http.NewRequest(method, url, nil)
	request.Header = *client.header
	if err != nil {
		return nil, err
	}

	response, err := client.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	text, _ := ioutil.ReadAll(response.Body)

	if !json.Valid(text) {
		return nil, &yaError{
			Err: "JSON invalid",
		}
	}

	statusCode := response.StatusCode
	if (statusCode != http.StatusOK) && (statusCode != http.StatusCreated) {
		var errya yaError
		if err = json.Unmarshal(text, &errya); err != nil {
			return nil, &yaError{
				Description: "json.Unmarshal",
				Err:         err.Error(),
			}
		}
		return nil, &errya
	}

	return &text, nil
}

func (client *Client) get(v interface{}, path string, params *url.Values) error {
	text, err := client.do(http.MethodGet, path, params)

	if err != nil {
		return err
	}

	if err = json.Unmarshal(*text, v); err != nil {
		return err
	}
	return nil
}

func (client *Client) put(v interface{}, path string, params *url.Values) error {
	text, err := client.do(http.MethodPut, path, params)

	if err != nil {
		return err
	}

	if err = json.Unmarshal(*text, v); err != nil {
		return err
	}
	return nil
}
