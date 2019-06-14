package yandexdisk

import (
	"net/http"
	"net/url"
	"os"
)

func (client *Client) getUrlUpload(path string) (link *Link, err error) {
	params := url.Values{}
	params.Add("path", path)

	err = client.get(&link, "/v1/disk/resources/upload", &params)
	return
}

func (client *Client) uploadFile(urlUpload string, filepath string) (err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}

	client_http := &http.Client{}
	req, _ := http.NewRequest(http.MethodPut, urlUpload, file)
	client_http.Do(req)

	return
}

func (client *Client) UploadFile(path string, filepath string) (err error) {
	link, err := client.getUrlUpload(path)
	if err != nil {
		return
	}

	err = client.uploadFile(link.Href, filepath)
	if err != nil {
		return
	}

	return
}
