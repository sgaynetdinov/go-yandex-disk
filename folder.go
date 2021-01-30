package yandexdisk

import (
	"net/url"
)

func (client *Client) CreateFolder(name string) error {
	params := url.Values{}
	params.Add("path", name)

	err := client.put(&link{}, "/v1/disk/resources", &params)
	return err
}

func (client *Client) IsExistsFolder(name string) (bool, error) {
	params := url.Values{}
	params.Add("path", name)

	var emptyResponse struct{}
	err := client.get(&emptyResponse, "/v1/disk/resources", &params)

	if err == nil {
		return true, nil
	}

	if err.Error() == "Resource not found. - DiskNotFoundError" {
		return false, nil
	}

	return false, err
}
