package yandexdisk

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var resourceJSON = []byte(`{
  "type": "media",
  "path": "/music/2pac/Changes.mp3",
  "name": "Changes.mp3",
  "created": "1998-10-13T00:00:00+00:00",
  "modified": "1998-10-13T00:00:00+00:00"
}`)

var resourceOptionalFieldJSON = []byte(`{
  "type": "media",
  "path": "/music/2pac/Changes.mp3",
  "name": "Changes.mp3",
  "created": "1998-10-13T00:00:00+00:00",
  "modified": "1998-10-13T00:00:00+00:00",
  "md5": "100500"
}`)

func TestResource(t *testing.T) {
	var resource Resource

	if err := json.Unmarshal(resourceJSON, &resource); err != nil {
		t.Fatal("Unmarshal json")
	}

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

	json.Unmarshal(resourceOptionalFieldJSON, &resource)

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

func TestResourceGot(t *testing.T) {
	var req *http.Request
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(resourceOptionalFieldJSON)
		req = r
	}))
	defer ts.Close()

	client := NewClient("YOUR_TOKEN")
	client.apiURL = ts.URL

	resource, err := client.Stat("/music/2pac/Changes.mp3")

	if err != nil {
		t.Error("Error is not nil")
	}

	if resource.Name != "Changes.mp3" {
		t.Error("Error Name")
	}

	if req.Method != http.MethodGet {
		t.Error("Invalid method http")
	}

	if req.URL.RawQuery != "path=disk%3A%2Fmusic%2F2pac%2FChanges.mp3" {
		t.Error("Invalid", req.URL.RawQuery)
	}

	if req.URL.Path != "/v1/disk/resources" {
		t.Error("Invalid url")
	}
}
