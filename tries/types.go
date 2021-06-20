package tries

import "fmt"

var ent = [2]interface{}{}

type dict = map[string]interface{}
type ent[0] string
type ent[1] int

/*
func Types() (dict, ent[0], ent[1]){
  return dict{"uno": false}, type(ent[0]){"Sad"}, type(ent[1]){1}
}
*/

func Types() {
  fmt.Println(ent[0])
}
