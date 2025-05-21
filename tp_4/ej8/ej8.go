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
	return arbol.raiz == nil
}

func (arbol *ABB) insertar(input int) {
	//creamos el nuevo nodo con el valor de input
	nuevoNodo := &Nodo{valor: input, izquierda: nil, derecha: nil}

	//creamos el nodo auxiliar
	auxNodo := arbol.raiz

	//creamos un nodo padre
	padre := (*Nodo)(nil)

	//si el arbol esta vacio
	if arbol.EsVacio() {
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

	//devuelve true o false en caso de encontrar o no el valor
	if observador.valor == input {
		return true
	} else {
		return false
	}
}

func (arbol *ABB) eliminar(input int) {

	if arbol.EsVacio() {
		return
	}
	observador := arbol.raiz
	padre := (*Nodo)(nil)

	/*
		tres casos posibles
		1. sin hijos
		2. con un hijo
		3. con dos hijos
	*/

	//se frena el for loop cuando el valor es encontrado o se llega a un nil
	for observador != nil && observador.valor != input {
		padre = observador
		if observador.valor < input {
			observador = observador.derecha
		} else {
			observador = observador.izquierda
		}
	}

	// si el valor no existe
	if observador == nil {
		return
	}

	//caso 1, el nodo es una hoja
	if observador.derecha == nil && observador.izquierda == nil {
		// si el nodo a eliminar es la raiz
		if arbol.raiz.valor == input {
			arbol.raiz = nil
			return
		}
		//chequeamos si el nodo a eliminar es el izquierdo o el derecho y luego lo eliminamos
		if padre.izquierda == observador {
			padre.izquierda = nil
			return
		}
		if padre.derecha == observador {
			padre.derecha = nil
			return
		}
	}

	//caso 2, el nodo tiene un hijo
	//es verdadero solo si uno es nil y el otro no, es decir, solo tiene un hijo.
	if (observador.derecha == nil) != (observador.izquierda == nil) {
		var hijo *Nodo

		//verificamos cual es el nodo a conectar con el padre del eliminado
		// y lo llamamos hijo
		if observador.derecha != nil {
			hijo = observador.derecha
		}
		if observador.izquierda != nil {
			hijo = observador.izquierda
		}

		// si el valor a eliminar es la raiz, se elimina,
		// si no se chequea si es el hijo izquierdo o derecho
		if observador == arbol.raiz {
			arbol.raiz = hijo
		} else if padre.derecha == observador {
			padre.derecha = hijo
		} else if padre.izquierda == observador {
			padre.izquierda = hijo
		}
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

	arbol.eliminar(3)
	arbol.eliminar(12)
	arbol.eliminar(10)
}
