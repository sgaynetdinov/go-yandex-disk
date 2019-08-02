package yandexdisk

import (
	"bufio"
	"net/http"
	"net/url"
)

func (client *Client) getUrlUpload(path string, overwrite bool) (link *link, err error) {
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

func (client *Client) uploadFile(urlUpload string, reader *bufio.Reader) (err error) {
	client_http := &http.Client{}
	req, _ := http.NewRequest(http.MethodPut, urlUpload, reader)
	client_http.Do(req)

	return
}

func (client *Client) UploadFile(path string, overwrite bool, reader *bufio.Reader) (err error) {
	link, err := client.getUrlUpload(path, overwrite)
	if err != nil {
		return
	}

	err = client.uploadFile(link.Href, reader)
	if err != nil {
		return
	}

	return
}
