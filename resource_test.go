package yandexdisk

import (
	"encoding/json"
	"testing"
)

var resource_json = []byte(`{
  "type": "media",
  "path": "/music/2pac/Changes.mp3",
  "name": "Changes.mp3",
  "created": "1998-10-13T00:00:00+00:00",
  "modified": "1998-10-13T00:00:00+00:00"
}`)

var resource_optional_field_json = []byte(`{
  "type": "media",
  "path": "/music/2pac/Changes.mp3",
  "name": "Changes.mp3",
  "created": "1998-10-13T00:00:00+00:00",
  "modified": "1998-10-13T00:00:00+00:00",
  "md5": "100500"
}`)

func TestResource(t *testing.T) {
	var resource Resource

	json.Unmarshal(resource_json, &resource)

	if resource.Type != "media" {
		t.Error("Invalid Type")
	}

	if resource.Path != "/music/2pac/Changes.mp3" {
		t.Error("Invalid Path")
	}

	if resource.Name != "Changes.mp3" {
		t.Error("Invalid Name")
	}

	if resource.Created != "1998-10-13T00:00:00+00:00" {
		t.Error("Invalid Created")
	}

	if resource.Modified != "1998-10-13T00:00:00+00:00" {
		t.Error("Invalid Modified")
	}

	if resource.Md5 != "" {
		t.Error("Invalid Md5")
	}

	if resource.Sha256 != "" {
		t.Error("Invalid Sha256")
	}
}

func TestResourceOptionalField(t *testing.T) {
	var resource Resource

	json.Unmarshal(resource_optional_field_json, &resource)

	if resource.Type != "media" {
		t.Error("Invalid Type")
	}

	if resource.Path != "/music/2pac/Changes.mp3" {
		t.Error("Invalid Path")
	}

	if resource.Name != "Changes.mp3" {
		t.Error("Invalid Name")
	}

	if resource.Created != "1998-10-13T00:00:00+00:00" {
		t.Error("Invalid Created")
	}

	if resource.Modified != "1998-10-13T00:00:00+00:00" {
		t.Error("Invalid Modified")
	}

	if resource.Md5 != "100500" {
		t.Error("Invalid Md5")
	}
}
