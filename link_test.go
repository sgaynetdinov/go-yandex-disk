package yandexdisk

import (
	"encoding/json"
	"testing"
)

var linkJSON = []byte(`{
  "href": "https://uploader1d.dst.yandex.net:443/upload-target/...",
  "method": "PUT",
  "templated": false
}`)

func TestLink(t *testing.T) {
	var link link

	json.Unmarshal(linkJSON, &link)

	if link.Href != "https://uploader1d.dst.yandex.net:443/upload-target/..." {
		t.Error("Invalid Href")
	}

	if link.Method != "PUT" {
		t.Error("Invalid Method")
	}

	if link.Templated != false {
		t.Error("Invalid Templated")
	}
}
