package yandexdisk

type Disk struct {
	TrashSize  uint64 `json:"trash_size"`
	TotalSpace uint64 `json:"total_space"`
	UsedSpace  uint64 `json:"used_space"`
}

func (client *Client) DiskInfo() (disk *Disk, err error) {
	err = client.get(&disk, "/v1/disk", nil)
	if err != nil {
		return nil, err
	}

	return
}
