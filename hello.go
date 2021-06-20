package main

import (
  "fmt"
  "github.com/ivanuria/somethingstupid/greetings"
  "github.com/ivanuria/somethingstupid/tries"
)

func main() {
  message := greetings.Hello("GladOS")
  fmt.Println(message)
}
