package tries

func Mapslices() map[string]interface{} {
  items := []map[string]interface{}{}
  append(items, map[string]interface{}{"hola": "Patatas", "nada": 1})
  append(items, map[string]interface{}{"hola": 2, "menos": "cachopo"})
  return items
}
