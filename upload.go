package yandexdisk

import (
	"io"
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

func (client *Client) uploadFile(urlUpload string, reader io.Reader) (err error) {
	req, err := http.NewRequest(http.MethodPut, urlUpload, reader)
	if err != nil {
		return
	}

	clientHTTP := &http.Client{}
	res, err := clientHTTP.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	return
}

func (client *Client) UploadFile(path string, overwrite bool, reader io.Reader) (err error) {
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
