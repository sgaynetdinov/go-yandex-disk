name: Golang

on: push

jobs:
  test:
    runs-on: cedrickring/golang-action@1.3.0
    steps:
      - name: Test
      - run: go build && go test -v -failfast
