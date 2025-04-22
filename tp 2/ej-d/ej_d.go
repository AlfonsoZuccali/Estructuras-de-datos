package main

import (
	"fmt"
)

// iterativa
func mayor_elemento(array []int) int {
	var mayor int
	for i := range array {
		if i == 0 {
			mayor = array[i]
		} else if array[i] > mayor {
			mayor = array[i]
		}
	}
	return mayor
}

func mayor_recursiva(array []int, mayor int, indice int) int {
	if indice >= len(array) {
		return mayor
	} else if indice == 0 {
		mayor = array[indice]
	} else if array[indice] > mayor {
		mayor = array[indice]
	}
	return mayor_recursiva(array, mayor, indice+1)
}

func main() {
	var arreglo = []int{9, 27, 78, 37, 3, 91}
	var mayor int = mayor_elemento(arreglo)

	fmt.Printf("El mayor numero del array es: %v \n", mayor)

	mayor = mayor_recursiva(arreglo, 0, 0)
	fmt.Printf("El mayor numero del array es: %v \n", mayor)
}
