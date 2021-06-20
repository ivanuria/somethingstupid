package main

import (
  "fmt"
  "github.com/ivanuria/somethingstupid/greetings"
  "github.com/ivanuria/somethingstupid/tries"
)

func main() {
  message := greetings.Hello("GladOS")
  fmt.Println(message)

  me := &tries.Persona{Nombre:="Iván", Edad:=36}
  fmt.Println(me.MayorDeEdad())
}
