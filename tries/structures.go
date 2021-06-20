package tries

type Persona struct {
  Nombre string
  Edad int
}

func (persona *Persona) MayorDeEdad() bool {
  if persona.Edad >= 18 {
    return true
  } else {
    return false
  }
}
