package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func CalcularValorTotal(p Producto) float64 {
	return p.Precio * float64(p.Cantidad)
}

func TieneStock(p Producto) bool {
	return p.Cantidad > 0
}

func MostrarInformacion(p Producto) {
	fmt.Println("Nombre: ", p.Nombre)
	fmt.Println("Precio: ", p.Precio)
	fmt.Println("Cantidad: ", p.Cantidad)

	total := CalcularValorTotal(p)
	fmt.Println("Valor Total: ", total)

	if TieneStock(p) {
		fmt.Println("Estado de stock: Disponible")
	} else {
		fmt.Println("Estado de stock: Sin stock")
	}
	fmt.Println("---------------------------------")
}

func main() {
	p1 := Producto{Nombre: "Teclado", Precio: 25.50, Cantidad: 10}
	p2 := Producto{Nombre: "Mouse", Precio: 15.00, Cantidad: 0}
	p3 := Producto{Nombre: "Monitor", Precio: 145.99, Cantidad: 5}

	fmt.Println("======= REPORTE GENERAL DE PRODUCTOS =======")
	MostrarInformacion(p1)
	MostrarInformacion(p2)
	MostrarInformacion(p3)
}
