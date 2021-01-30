package yandexdisk

import (
	"encoding/json"
	"testing"
)

func TestError(t *testing.T) {
	var e yaError

	error_json := []byte(`{
  "description": "resource already exists",
  "error": "PlatformResourceAlreadyExists"
}`)

	json.Unmarshal(error_json, &e)

	if e.Error() != "resource already exists - PlatformResourceAlreadyExists" {
		t.Error("Invalid error message")
	}
}
