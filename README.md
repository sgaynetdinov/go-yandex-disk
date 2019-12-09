[![GoDoc](https://godoc.org/github.com/sgaynetdinov/go-yandex-disk?status.svg)](https://godoc.org/github.com/sgaynetdinov/go-yandex-disk)
[![Go Report Card](https://goreportcard.com/badge/github.com/sgaynetdinov/go-yandex-disk)](https://goreportcard.com/report/github.com/sgaynetdinov/go-yandex-disk)
[![Release](https://img.shields.io/github/release/sgaynetdinov/go-yandex-disk.svg?style=flat-square)](https://github.com/sgaynetdinov/go-yandex-disk/releases/latest)


## Install

`go get -u github.com/sgaynetdinov/go-yandex-disk`

## Token

[https://yandex.ru/dev/disk/api/concepts/quickstart-docpage/](https://yandex.ru/dev/disk/api/concepts/quickstart-docpage/)

## First Step

```
import (
    yandexdisk "github.com/sgaynetdinov/go-yandex-disk"
)

func main() {
    client := yandexdisk.NewClient("YOUR_TOKEN")
}
```


## Upload file

```
import (
    "bufio"
    "os"

    ...
}

func main() {
    ...

    file, _ := os.Open("Changes.mp3")
    client.UploadFile("/music/2pac/Changes.mp3", false, bufio.NewReader(file))
}
```


## Upload file (from web)

```
import (
    "net/http"
    "os"

    ...
}

func main() {
    ...

    resp, _ := http.Get("https://example.com/Changes.mp3")
    defer resp.Body.Close()

    client.UploadFile("/music/2pac/Changes.mp3", false, bufio.NewReader(resp.Body))
}
```


## Documentation
- API Yandex.Disk: https://yandex.ru/dev/disk/api/concepts/about-docpage/
