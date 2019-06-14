workflow "Golang Test" {
  resolves = ["Test"]
  on = "push"
}

action "Test" {
  uses = "cedrickring/golang-action@1.3.0"
  
  args = "go build && go test -v -failfast"
}
