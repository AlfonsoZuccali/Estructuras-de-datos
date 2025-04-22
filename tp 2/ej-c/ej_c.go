package main

import (
	"fmt"
	"strings"
)

// iterativa
func contar_vocales(cadena string) (int, int) {
	var vocales string = "aeiouAEIOU"
	var num_vocales int = 0
	var num_consonantes int = 0

	for _, i := range cadena {
		if strings.ContainsRune(vocales, i) {
			num_vocales++
		} else if (i >= 'A' && i <= 'Z') || (i >= 'a' && i <= 'z') {
			num_consonantes++
		}
	}
	return num_consonantes, num_vocales
}

func contar_recursiva(cadena string, num_vocales int, num_consonantes int, indice int) (int, int) {

	if indice >= len(cadena) { //ponemos >= porque el indice maximo es length-1
		return num_vocales, num_consonantes
	}
	vocales := "aeiouAEIOU"
	caracter := cadena[indice]
	if strings.ContainsRune(vocales, rune(caracter)) {
		num_vocales++
	} else if (caracter >= 'A' && caracter <= 'Z') || (caracter >= 'a' && caracter <= 'z') {
		num_consonantes++
	}
	return contar_recursiva(cadena, num_vocales, num_consonantes, indice+1)
}

func main() {

	s, s1 := "hola soy totooo", "aeiou"
	c, v := contar_vocales(s)
	c1, v1 := contar_vocales(s1)
	//iterativa
	fmt.Printf("En la palabra \"%v\", la cantidad de vocales son %v y la cantidad de consonantes son %v \n", s, v, c)
	fmt.Printf("En la palabra \"%v\", la cantidad de vocales son %v y la cantidad de consonantes son %v \n", s1, v1, c1)

	//recursiva
	v, c = contar_recursiva(s, 0, 0, 0)
	v1, c1 = contar_recursiva(s1, 0, 0, 0)
	fmt.Printf("En la palabra \"%v\", la cantidad de vocales son %v y la cantidad de consonantes son %v \n", s, v, c)
	fmt.Printf("En la palabra \"%v\", la cantidad de vocales son %v y la cantidad de consonantes son %v \n", s1, v1, c1)

}
