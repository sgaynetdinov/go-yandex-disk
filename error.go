package yandexdisk

import "errors"

var ErrResourceNotFound = errors.New("Resource not found. - DiskNotFoundError")

type yaError struct {
	Description string `json:"description"`
	Err         string `json:"error"`
}

func (e *yaError) Error() string {
	return e.Description + " - " + e.Err
}
