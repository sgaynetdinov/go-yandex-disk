package yandexdisk

import (
	"net/url"
)

func (client *Client) CreateFolder(name string) (link *Link, err error) {
	params := url.Values{}
	params.Add("path", name)

	err = client.put(&link, "/v1/disk/resources", &params)
	return
}
