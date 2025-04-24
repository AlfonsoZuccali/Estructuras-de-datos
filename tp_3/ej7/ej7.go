package main

import (
	"fmt"
)

// Definimos el nodo
type Nodo struct {
	valor int
	//usamos siguiente y anterior, siendo siguiente, el que se acerca
	//a la cola, mientras que el anterior se acerca a la cabeza
	siguiente *Nodo
	anterior  *Nodo
}

// Definimos la Double linked list
type DLL struct {
	head   *Nodo
	tail   *Nodo
	length int
}

// Metodo para agregar a partir de la cola
func (lista *DLL) AddToTail(elemento int) {
	//creamos el nuevo nodo
	nuevoNodo := &Nodo{valor: elemento}

	//si la lista esta vacia
	if lista.head == nil && lista.tail == nil {

		//el nuevo nodo es el head y tail de la lista
		lista.head = nuevoNodo
		lista.tail = nuevoNodo

	} else {
		//se agrega el nuevo nodo al final de la lista y
		//acomodamos los punteros
		lista.tail.siguiente = nuevoNodo
		nuevoNodo.anterior = lista.tail
		lista.tail = nuevoNodo
	}
	//Incrementa el tamaño de la lista
	lista.length++
}

// Metodo para agregar a partir de la cabeza
func (lista *DLL) AddToHead(elemento int) {
	//creamos el nuevo nodo
	nuevoNodo := &Nodo{valor: elemento}

	//si la lista esta vacia
	if lista.head == nil && lista.tail == nil {

		//el nuevo nodo es el head y tail de la lista
		lista.head = nuevoNodo
		lista.tail = nuevoNodo

	} else {
		//se agrega el nuevo nodo al final de la lista y
		//acomodamos los punteros

		lista.head.anterior = nuevoNodo
		nuevoNodo.siguiente = lista.head
		lista.head = nuevoNodo

	}
	//Incrementa el tamaño de la lista
	lista.length++
}

// Funcion para imprimir numeros de Head a Tail
func (lista *DLL) Print() {

	//creamos un nodo observador auxiliar
	aux := lista.head

	//si no estamos en el ultimo nodo
	for i := 0; i < lista.length; i++ {
		if aux != nil {
			fmt.Println(aux.valor)
			aux = aux.siguiente
		} else {
			return
		}
	}
}

func main() {
	p := &DLL{}
	p.AddToTail(1)
	p.AddToTail(2)
	p.AddToTail(3)
	p.AddToTail(6)

	fmt.Println("Head:", p.head.valor)
	fmt.Println("Tail:", p.tail.valor)
	fmt.Println("Size:", p.length)

	p.AddToHead(-1)
	p.AddToHead(-2)
	p.AddToHead(-3)
	p.AddToHead(-6)
	fmt.Println("Head:", p.head.valor)
	fmt.Println("Tail:", p.tail.valor)
	fmt.Println("Size:", p.length)
	p.Print()
}
