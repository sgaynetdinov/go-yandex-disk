package yandexdisk

import (
	"net/url"
)

func (client *Client) CreateFolder(name string) error {
	params := url.Values{}
	params.Add("path", name)

	err := client.put(&Link{}, "/v1/disk/resources", &params)
	return err
}
