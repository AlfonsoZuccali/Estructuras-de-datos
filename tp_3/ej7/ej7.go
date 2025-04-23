package main

import (
	"fmt"
)

type Nodo struct {
	valor     int
	siguiente *Nodo
	anterior  *Nodo
}

type DLL struct {
	head   *Nodo
	tail   *Nodo
	length int
}

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
		nuevoNodo.siguiente = lista.tail
		lista.tail.anterior = nuevoNodo
		lista.tail = nuevoNodo
	}
	//Incrementa el tamaño de la lista
	lista.length++
}

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
		nuevoNodo.siguiente = lista.head
		lista.head.anterior = nuevoNodo
		lista.head = nuevoNodo
	}
	//Incrementa el tamaño de la lista
	lista.length++
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
}
