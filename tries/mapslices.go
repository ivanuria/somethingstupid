package tries

func Mapslices() ([]map[string]interface{}, *[]map[string]interface{}) {
  items := []map[string]interface{}{}
  daitems := &items
  items = append(items, map[string]interface{}{"hola": "Patatas", "nada": 1})
  items = append(items, map[string]interface{}{"hola": 2, "menos": "cachopo"})
  return items, daitems
}
