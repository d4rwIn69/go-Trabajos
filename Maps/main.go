package main

import "fmt"

//Estructura map[tipodeclave]tipodevalor

func main() {
	visitas := make(map[string]int)
	visitas["Inicio"] = 10
	visitas["Noticias"] = 20
	visitas["Deportes"] = 25
	visitas["Hogar"] = 80
	fmt.Println(visitas)
	fmt.Println("Las Noticias:", visitas["Noticias"])
	fmt.Println("Las Noticias de hogar son:", visitas["Hogar"])
	visitas["Hogar"] = 20
	fmt.Println("Actualizado Hogar.....\nLas Noticias de hogar son:", visitas["Hogar"])
	delete(visitas, "Hogar")
	fmt.Println(visitas)

	for key, value := range visitas {
		fmt.Println(key, ":", value)

	}
	visitas["Galeria"] = 20
	fmt.Println("Las noticias de visitas de Galeria son:", visitas["Galeria"])
	visitas["Inicio"] = 200
	fmt.Println("Las noticias de visitas de Inicio son:", visitas["Inicio"])
	fmt.Println("La suma de todas las visitas es:", Suma(visitas))

}
func Suma(visitas map[string]int) int {
	sum := 0
	for _, value := range visitas {
		sum += value
	}
	return sum
}
