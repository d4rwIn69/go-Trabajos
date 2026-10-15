package main

import "fmt"

func main() {
	votosActividades := make(map[string]int)

	votosActividades["Deportes"] = 0
	votosActividades["Videojuegos"] = 0
	votosActividades["Cine"] = 0
	votosActividades["Música"] = 0

	fmt.Println("=== REGISTRO DE VOTOS (5 ESTUDIANTES) ===")

	for i := 1; i <= 5; i++ {
		var seleccion int
		fmt.Println("\nVoto número: ", i)
		fmt.Println("1. Deportes")
		fmt.Println("2. Videojuegos")
		fmt.Println("3. Cine")
		fmt.Println("4. Música")
		fmt.Print("Seleccione un número (1-4): ")
		fmt.Scanln(&seleccion)

		switch seleccion {
		case 1:
			votosActividades["Deportes"]++
		case 2:
			votosActividades["Videojuegos"]++
		case 3:
			votosActividades["Cine"]++
		case 4:
			votosActividades["Música"]++
		default:
			fmt.Println("Opción no válida. Voto perdido.")
		}
	}

	fmt.Println("\n======= RESULTADOS DE LA VOTACIÓN =======")
	
	for key, value := range votosActividades {
		fmt.Println(key, ":", value, "votos")
	}

	fmt.Println("\nLa actividad con mayor número de votos es:", DeterminarGanador(votosActividades))
}

func DeterminarGanador(votacion map[string]int) string {
	mayorVotos := -1
	actividadGanadora := ""

	for key, value := range votacion {
		if value > mayorVotos {
			mayorVotos = value
			actividadGanadora = key
		}
	}
	
	return actividadGanadora
}
