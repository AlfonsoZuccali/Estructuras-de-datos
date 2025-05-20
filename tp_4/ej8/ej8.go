package main

import "fmt"

//Nodo
type Nodo struct {
	izquierda *Nodo
	valor     int
	derecha   *Nodo
}

//arbol binario de busqueda
type ABB struct {
	raiz *Nodo
}

func (arbol *ABB) EsVacio() bool {
	if arbol.raiz == nil {
		return true
	}
	return false
}

func (arbol *ABB) insertar(input int) {
	//creamos el nuevo nodo con el valor de input
	nuevoNodo := &Nodo{valor: input, izquierda: nil, derecha: nil}

	//creamos el nodo auxiliar
	auxNodo := arbol.raiz

	//creamos un nodo padre
	padre := (*Nodo)(nil)

	//si el arbol esta vacio
	if arbol.EsVacio() == true {
		arbol.raiz = nuevoNodo
		//si no esta vacio
	} else {

		//recorremos el arbol
		for auxNodo != nil {
			padre = auxNodo
			if auxNodo.valor > input {
				auxNodo = auxNodo.izquierda
			} else {
				auxNodo = auxNodo.derecha
			}
		}
		//al final de este bloque auxNodo va a ser el nodo con el valor nil
		//mientras que el padre sera el antecesor al cual deberemos reemplazarle
		//los valores de derecha o izquierda con nuestro nuevoNodo

		//cambiamos los valores del padre para que apunte al nuevoNodo
		if padre.valor > input {
			padre.izquierda = nuevoNodo
		} else {
			padre.derecha = nuevoNodo
		}

	}
}

//funcion para buscar un elemento en el ABB
func (arbol *ABB) buscar(input int) bool {

	//creamos el nodo auxiliar
	observador := arbol.raiz

	//si el arbol esta vacio return false
	if arbol.EsVacio() {
		return false
	} else {
		//se frena el for loop cuando el valor es encontrado o se llega a una hoja
		for observador.valor != input && (observador.izquierda != nil || observador.derecha != nil) {

			//si el valor del nodo es mayor al input vamos hacia la izquierda
			if observador.valor > input && observador.izquierda != nil {
				observador = observador.izquierda

				// si el valor del nodo es menor al input vamos hacia la derecha
			} else if observador.valor < input && observador.derecha != nil {
				observador = observador.derecha
			}
		}
	}
	if observador.valor == input {
		return true
	} else {
		return false
	}
}

func main() {
	arbol := ABB{raiz: nil}
	arbol.insertar(5)
	arbol.insertar(10)
	arbol.insertar(8)
	arbol.insertar(3)
	arbol.insertar(12)
	arbol.insertar(0)

	fmt.Println("EL valor 2 esta en el arbol: ", arbol.buscar(2))
	fmt.Println("EL valor 0 esta en el arbol: ", arbol.buscar(0))
	fmt.Println("EL valor 12 esta en el arbol: ", arbol.buscar(12))
	fmt.Println("EL valor 3 esta en el arbol: ", arbol.buscar(3))
	fmt.Println("EL valor 5 esta en el arbol: ", arbol.buscar(5))
	fmt.Println("EL valor 1 esta en el arbol: ", arbol.buscar(1))
}
