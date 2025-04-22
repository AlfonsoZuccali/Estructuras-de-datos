package main

import (
	"fmt"
)

// iterativa
func invertir_array(array []int, length int) {

	var invertido = [5]int{}
	var indice int = 0
	for i := length - 1; i >= 0; i-- {
		invertido[indice] = array[i]
		indice++
	}
	fmt.Println(invertido)
}

func main() {
	var array = []int{1, 2, 3, 4, 5}
	invertir_array(array, 5)
}
