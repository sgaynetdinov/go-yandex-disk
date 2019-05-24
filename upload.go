package yandexdisk

func (client *Client) getUrlUpload(path string) (link *Link, err error) {
	err = client.get(&link)
	return
}
