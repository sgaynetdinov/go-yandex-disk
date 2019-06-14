workflow "Golang Test" {
  resolves = ["Golang Action"]
  on = "push"
}

action "test" {
  uses = "cedrickring/golang-action@1.3.0"
  
  args = "go build && go test"
}
