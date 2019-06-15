package yandexdisk

import (
	"net/http"
	"net/url"
	"os"
)

func (client *Client) getUrlUpload(path string, overwrite bool) (link *Link, err error) {
	params := url.Values{}
	params.Add("path", path)
	if overwrite {
		params.Add("overwrite", "true")
	} else {
		params.Add("overwrite", "false")
	}

	err = client.get(&link, "/v1/disk/resources/upload", &params)
	return
}

func (client *Client) uploadFile(urlUpload string, file *os.File) (err error) {
	client_http := &http.Client{}
	req, _ := http.NewRequest(http.MethodPut, urlUpload, file)
	client_http.Do(req)

	return
}

func (client *Client) UploadFile(path string, overwrite bool, file *os.File) (err error) {
	link, err := client.getUrlUpload(path, overwrite)
	if err != nil {
		return
	}

	err = client.uploadFile(link.Href, file)
	if err != nil {
		return
	}

	return
}
