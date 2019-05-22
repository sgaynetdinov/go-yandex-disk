package yandexdisk

type yaError struct {
	Description string `json:"description"`
	Err         string `json:"error"`
}

func (e *yaError) Error() string {
	return e.Description + " - " + e.Err
}
