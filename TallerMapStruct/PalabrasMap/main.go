package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Ingrese una frase o texto: ")
	lector := bufio.NewReader(os.Stdin)
	texto, _ := lector.ReadString('\n')
	texto = strings.TrimSpace(texto)

	palabras := strings.Fields(texto)

	conteo := make(map[string]int)

	for _, palabra := range palabras {
		conteo[palabra]++
	}

	fmt.Println("\n======= REPORTE DE PALABRAS =======")
	for key, value := range conteo {
		fmt.Println(key, ":", value, "veces")
	}

	maxRepeticiones := -1
	palabraMasRepetida := ""

	for key, value := range conteo {
		if value > maxRepeticiones {
			maxRepeticiones = value
			palabraMasRepetida = key
		}
	}

	fmt.Println("\n=========================================")
	fmt.Println("La palabra que más se repite es:", palabraMasRepetida, "con", maxRepeticiones, "repeticiones")
	fmt.Println("=========================================")
}
