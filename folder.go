package yandexdisk

import (
	"net/url"
	"strings"
)

func (client *Client) CreateFolder(name string) error {
	if !strings.HasPrefix(name, "/") {
		name = "/" + name
	}

	params := url.Values{}
	params.Add("path", name)

	err := client.put(&Link{}, "/v1/disk/resources", &params)
	return err
}
