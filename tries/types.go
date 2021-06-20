package tries

var {
  ent = [2]type{string, int}
}

type (
  dict = map[string]interface{}
  uno = ent[0]
  dos = ent[1]
)

func Types() (dict, uno, dos){
  return dict{"uno": false}, uno{"Sad"}, dos{1}
}
