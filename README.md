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

## Documentation
- API Yandex.Disk: https://yandex.ru/dev/disk/api/concepts/about-docpage/
