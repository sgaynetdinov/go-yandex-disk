package yandexdisk

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
)

const API_URL = "https://cloud-api.yandex.net:443"
const VERSION_API = "v1"

func JoinURL(apiURL string, paths ...string) string {
	u, _ := url.Parse(apiURL)

	for _, path := range paths {
		u.Path = filepath.Join(u.Path, path)
	}
	return u.String()
}

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
		apiURL:     JoinURL(API_URL, VERSION_API),
		header:     &header,
		httpClient: new(http.Client),
	}
}

func (client *Client) do(method string, path string, params *url.Values) (*[]byte, error) {
	apiURL := client.apiURL

	if path != "" {
		apiURL = JoinURL(apiURL, path)
	}

    if params != nil && params.Get("path") != "" && !strings.HasPrefix(params.Get("path"), "disk:") {
		name := params.Get("path")

		if !strings.HasPrefix(name, "/") {
			name = "/" + name
		}

		params.Set("path", "disk:"+name)
	}

	if params != nil {	
		apiURL += "?" + params.Encode()
	}

	request, err := http.NewRequest(method, apiURL, nil)
	if err != nil {
		return nil, err
	}

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
