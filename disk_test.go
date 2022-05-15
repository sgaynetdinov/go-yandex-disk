package yandexdisk

import (
	"net/http"
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

func TestDiskInfo(t *testing.T) {
	req, client := makeServer(diskJSON, http.StatusOK)
	disk, err := client.DiskInfo()

	if err != nil {
		t.Error("Error is not nil", err)
	}

	if req.URL.Path != "/v1/disk" {
		t.Error(req.URL.Path)
	}

	if req.URL.RawQuery != "" {
		t.Error(req.URL.RawQuery)
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
	req, client := makeServer(
		[]byte(`{"description": "resource already exists", "error": "PlatformResourceAlreadyExists"}`),
		http.StatusInternalServerError,
	)

	got, err := client.DiskInfo()

	if got != nil {
		t.Error()
	}

	if err == nil {
		t.Error("Error is nil")
	}

	if err.Error() != "resource already exists - PlatformResourceAlreadyExists" {
		t.Error("Invalid error message")
	}

	if req.URL.Path != "/v1/disk" {
		t.Error(req.URL.Path)
	}

	if req.URL.RawQuery != "" {
		t.Error(req.URL.RawQuery)
	}
}
