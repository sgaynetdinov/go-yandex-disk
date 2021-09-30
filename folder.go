package yandexdisk

import (
	"net/url"
	"path/filepath"
	"strings"
)

const PATH_SEPARATOR = string(filepath.Separator)

func (client *Client) Mkdir(path string) error {
	params := url.Values{}
	params.Add("path", path)

	err := client.put(&link{}, "/v1/disk/resources", &params)
	return err
}

func pathList(fullpath string) (pathItems []string) {
	var accPath string
	folderItems := strings.Split(fullpath, PATH_SEPARATOR)
	for _, folder := range folderItems {
		if folder == "" {
			continue
		}
		accPath = filepath.Join(accPath, folder)
		pathItems = append(pathItems, accPath)
	}

	return
}

func (client *Client) MkdirAll(path string) error {
	if exist, _ := client.IsExistsFolder(path); exist {
		return nil
	}

	for _, folder := range pathList(path) {
		exist, _ := client.IsExistsFolder(folder)

		if exist {
			continue
		}

		err := client.Mkdir(folder)
		if err != nil {
			return err
		}
	}
	return nil
}

func (client *Client) IsExistsFolder(path string) (bool, error) {
	params := url.Values{}
	params.Add("path", path)

	var emptyResponse struct{}
	err := client.get(&emptyResponse, "/v1/disk/resources", &params)

	if err == nil {
		return true, nil
	}

	if err.Error() == ErrResourceNotFound.Error() {
		return false, nil
	}

	return false, err
}
