package yandexdisk

import (
	"encoding/json"
	"testing"
)

var disk_json = []byte(`{
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

	json.Unmarshal(disk_json, &disk)

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
