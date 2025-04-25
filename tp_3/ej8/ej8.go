package main

import (
	"fmt"
)

type Nodo struct {
	valor rune
	//usamos siguiente y anterior, siendo siguiente, el que se acerca
	//a la cola, mientras que el anterior se acerca a la cabeza
	siguiente *Nodo
}

// definimos la lista
type ListaSimple struct {
	head   *Nodo
	length int
}

// Metodo para agregar a partir de la cabeza
func (lista *ListaSimple) AddToHead(elemento rune) {
	//creamos el nuevo nodo
	nuevoNodo := &Nodo{valor: elemento}

	//si la lista esta vacia
	if lista.head == nil {

		//el nuevo nodo es el head de la lista
		lista.head = nuevoNodo

	} else {
		//se agrega el nuevo nodo al principio de la lista y
		//acomodamos los punteros
		nuevoNodo.siguiente = lista.head
		lista.head = nuevoNodo

	}
	//Incrementa el tamaño de la lista
	lista.length++
}

// devuelve un string a partir de una lista
func (lista *ListaSimple) ToString() string {

	//en esta variable se devuelve el string dado vuelta
	var string_invertida string

	//generamos un nodo observador
	aux := lista.head

	//vamos nodo por nodo en la lista
	for i := 0; i < lista.length; i++ {
		//concatenamos cada uno de los caracteres que esta
		//en cada nodo y los guardamos en la variable string_invertida
		string_invertida += string(aux.valor)

		//pasamos al siguiente nodo
		aux = aux.siguiente
	}

	return string_invertida
}

// se hace una lista con el string en orden
func StringToLista(entrada string, lista *ListaSimple) {

	for i := len(entrada) - 1; i >= 0; i-- {
		lista.AddToHead(rune(entrada[i]))
	}
}

// se hace una lista con el string invertido
func StringInvertidoToLista(entrada string, lista *ListaSimple) {

	for i := range entrada {
		lista.AddToHead(rune(entrada[i]))
	}
}

// compara dos strings y devuelve el bool de la operacion
func EsPalindromo(string1 string, string2 string) bool {
	return string1 == string2
}

func main() {
	string_normal := &ListaSimple{}
	string_invertido := &ListaSimple{}
	string := "anita lava la tina"

	StringToLista(string, string_normal)
	StringInvertidoToLista(string, string_invertido)

	fmt.Println("La string original es: ", string_normal.ToString())
	fmt.Println("La string invertida es: ", string_invertido.ToString())

	fmt.Printf("La palabra %s, ¿es palindromo?: %t\n", string, EsPalindromo(string, string_invertido.ToString()))

}
