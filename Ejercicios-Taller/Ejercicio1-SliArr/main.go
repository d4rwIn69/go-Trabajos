package main

import "fmt"

func main() {

	notas := [6][4]float64{
		{8.5, 9.0, 7.8, 9.2}, 
		{7.0, 6.5, 8.0, 7.2},  
		{9.5, 8.8, 10.0, 9.1}, 
		{6.0, 7.5, 6.8, 7.0},  
		{8.8, 8.2, 7.9, 8.5}, 
		{5.5, 6.2, 6.0, 5.8},  
	}

	var sumaGeneral float64 = 0

	fmt.Println("============ REPORTE DE NOTAS POR ESTUDIANTE ============")

	for i := 0; i < 6; i++ {
		materiasEstudiante := notas[i][:]

		var sumaEstudiante float64 = 0
		notaAlta := materiasEstudiante[0]
		notaBaja := materiasEstudiante[0]

		for _, nota := range materiasEstudiante {
			sumaEstudiante += nota

			if nota > notaAlta {
				notaAlta = nota
			}
			if nota < notaBaja {
				notaBaja = nota
			}
		}

		promedioEstudiante := sumaEstudiante / 4.0
		sumaGeneral += promedioEstudiante

		fmt.Printf("Estudiante %d -> Promedio: %.2f | Nota Alta: %.2f | Nota Baja: %.2f\n", i+1, promedioEstudiante, notaAlta, notaBaja)
	}

	promedioGeneral := sumaGeneral / 6.0
	fmt.Printf("PROMEDIO GENERAL DE LA CLASE: %.2f\n", promedioGeneral)
}
