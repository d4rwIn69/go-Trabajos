package main

import "fmt"

type Persona struct {
	nombre string
}

func (p Persona) GetNombre() string {
	return p.nombre
}

func (p *Persona) SetNombre(nombre string) {
	p.nombre = nombre
}

func main() {

	persona := Persona{}
	persona.SetNombre("Cayetano")
	fmt.Println("Su nombre es: ", persona.GetNombre())

	persona.SetNombre("Adrian")
	fmt.Println("Ahora su nombre es: ", persona.GetNombre())


}
