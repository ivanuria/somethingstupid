package main

import (
  "fmt"
  somethingstupid.greetings "github.com/ivanuria/somethingstupid/greetings"
)

func main() {
  message := greetings.Hello("GladOS")
  fmt.Println(message)
}
