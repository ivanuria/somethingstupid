package main

import (
  "fmt"
  "github.com/ivanuria/somethingstupid/greetings"
)

func main() {
  message := greetings.Hello("GladOS")
  fmt.Println(message)
}
