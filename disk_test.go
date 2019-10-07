package yandexdisk

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var diskJSON = []byte(`{
  "trash_size": 4631577437,
  "total_space": 319975063552,
  "used_space": 26157681270,
  "system_folders":
  {
    "applications": "disk:/Приложения",
    "downloads": "disk:/Загрузки/"
  }
}`)

func TestDisk(t *testing.T) {
	var disk Disk

	json.Unmarshal(diskJSON, &disk)

	if disk.TrashSize != 4631577437 {
		t.Error("Disk.TrashSize")
	}

	if disk.TotalSpace != 319975063552 {
		t.Error("Disk.TotalSpace")
	}

	if disk.UsedSpace != 26157681270 {
		t.Error("Disk.UsedSpace")
	}
}

func TestDiskInfo(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(diskJSON)
	}))
	defer ts.Close()

	client := NewClient("YOUR_TOKEN")
	client.apiURL = ts.URL

	disk, err := client.DiskInfo()

	if err != nil {
		t.Error("Error is not nil")
	}

	if disk.TrashSize != 4631577437 {
		t.Error("Disk.TrashSize")
	}

	if disk.TotalSpace != 319975063552 {
		t.Error("Disk.TotalSpace")
	}

	if disk.UsedSpace != 26157681270 {
		t.Error("Disk.UsedSpace")
	}
}

func TestDiskInfoIfStatusNot200(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{
  			"description": "resource already exists",
  			"error": "PlatformResourceAlreadyExists"
		}`))
	}))
	defer ts.Close()

	client := NewClient("YOUR_TOKEN")
	client.apiURL = ts.URL

	_, err := client.DiskInfo()

	if err == nil {
		t.Error("Error is nil")
	}

	if err.Error() != "resource already exists - PlatformResourceAlreadyExists" {
		t.Error("Invalid error message")
	}

}
