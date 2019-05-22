package yandexdisk

import (
	"encoding/json"
	"testing"
)

var error_json = []byte(`{
  "description": "resource already exists",
  "error": "PlatformResourceAlreadyExists"
}`)

func TestError(t *testing.T) {
	var e yaError
	json.Unmarshal(error_json, &e)

	if e.Error() != "resource already exists - PlatformResourceAlreadyExists" {
		t.Error("Invalid error message")
	}
}
