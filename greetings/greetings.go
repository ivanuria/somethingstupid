package greetings

import (
  "fmt"
)

func Hello(name string) string {
  //Returns a String greetings to given name
  message := fmt.Sprintf("Hola, %v, peazo de cretine.", name)
  return message
}
