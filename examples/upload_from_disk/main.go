package main

import (
	"os"

	yandexdisk "github.com/sgaynetdinov/go-yandex-disk"
)

func main() {
	file, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	client := yandexdisk.NewClient("YOUR_TOKEN")
	client.UploadFile("main.go", false, file)
}
