package tries

var (
  dict = map[string]interface
  items = []dict
)

func Mapslices() slice {
  items.append(dict{"hola": "Patatas", "nada": 1})
  items.append(dict{"hola": 2, "menos": "cachopo"})
  return items
}
