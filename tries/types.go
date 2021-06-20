package tries

import "fmt"

type dict = map[string]interface{}

/*
func Types() (dict, ent[0], ent[1]){
  return dict{"uno": false}, type(ent[0]){"Sad"}, type(ent[1]){1}
}
*/

func Types() {
  fmt.Println(dict{"hola": 23})
}
