package yandexdisk

import (
	"net/url"
)

type Resource struct {
	Type     string `json:"type"`
	Path     string `json:"path"`
	Name     string `json:"name"`
	Created  string `json:"created"`
	Modified string `json:"modified"`
	Md5      string `json:"md5,omitempty"`
	Sha256   string `json:"sha256,omitempty"`
	Size     int64  `json:"size,omitempty"`
}

func (client *Client) Stat(path string) (resource *Resource, err error) {
	params := url.Values{}
	params.Add("path", path)
	err = client.get(&resource, "disk/resources", &params)
	return
}
