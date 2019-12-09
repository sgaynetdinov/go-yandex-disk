package main

import (
	"net/http"

	yandexdisk "github.com/sgaynetdinov/go-yandex-disk"
)

func main() {
	resp, err := http.Get("https://bit.ly/lazy_gopher")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	client := yandexdisk.NewClient("YOUR_TOKEN")
	client.UploadFile("lazy_gopher.png", false, resp.Body)
}
