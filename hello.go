package main

import (
  "fmt"
  "github.com/ivanuria/somethingstupid/greetings"
  "github.com/ivanuria/somethingstupid/tries"
)

func main() {
  message := greetings.Hello("GladOS")
  fmt.Println(message)

  me := &tries.Persona{Nombre:"Iván", Edad:36}
  fmt.Println(me.MayorDeEdad())

  nephew := &tries.Persona{"Ariel", 1}
  fmt.Println(nephew.MayorDeEdad())

  items, daitems := tries.Mapslices()
  fmt.Println(items)
  items[0]["daughter"] = map[string]interface{}{"Lord": false}
  fmt.Println(items)
  fmt.Println(daitems)

  uno, dos, tres := tries.Types()
  fmt.Println(uno, dos, tres)
}
