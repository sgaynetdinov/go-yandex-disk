package yandexdisk

type Resource struct {
	Type     string `json:"type"`
	Path     string `json: "path"`
	Name     string `json: "name"`
	Created  string `json: "created"`
	Modified string `json: "modified"`
	Md5      string `json: "md5,omitetype"`
	Sha256   string `json: "sha256,omitetype"`
}
