package main

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
	objetivo := arbol.raiz

	//si el arbol esta vacio return false
	if arbol.EsVacio() {
		return false
	} else {
		//se frena el for loop cuando el valor es encontrado o se llega a una hoja
		for objetivo.valor != input && (objetivo.izquierda != nil || objetivo.derecha != nil) {

			//si el valor del nodo es mayor al input vamos hacia la izquierda
			if objetivo.valor > input && objetivo.izquierda != nil {
				objetivo = objetivo.izquierda

				// si el valor del nodo es menor al input vamos hacia la derecha
			} else if objetivo.valor < input && objetivo.derecha != nil {
				objetivo = objetivo.derecha
			}
		}
	}

	//devuelve true o false en caso de encontrar o no el valor
	if objetivo.valor == input {
		return true
	} else {
		return false
	}
}

func (arbol *ABB) eliminar(input int) {

	if arbol.EsVacio() {
		return
	}

	objetivo := arbol.raiz
	padre := (*Nodo)(nil)

	//se frena el for loop cuando el valor es encontrado
	//o se llega a un nil
	for objetivo != nil && objetivo.valor != input {
		padre = objetivo
		if objetivo.valor < input {
			objetivo = objetivo.derecha
		} else {
			objetivo = objetivo.izquierda
		}
	}

	// si el valor no existe
	if objetivo == nil {
		return
	}

	/*
		tres casos posibles
		1. sin hijos
		2. con un hijo
		3. con dos hijos
	*/

	//caso 1, el nodo es una hoja
	if objetivo.derecha == nil && objetivo.izquierda == nil {
		// si el nodo a eliminar es la raiz
		if arbol.raiz.valor == input {
			arbol.raiz = nil
			return
		}
		//chequeamos si el nodo a eliminar es el izquierdo o el derecho y luego lo eliminamos
		if padre.izquierda == objetivo {
			padre.izquierda = nil
			return
		}
		if padre.derecha == objetivo {
			padre.derecha = nil
			return
		}
	}

	//caso 2, el nodo tiene un hijo
	//es verdadero solo si uno es nil y el otro no, es decir, solo tiene un hijo.
	if (objetivo.derecha == nil) != (objetivo.izquierda == nil) {
		var hijo *Nodo

		//verificamos cual es el nodo a conectar con el padre del eliminado
		// y lo llamamos hijo
		if objetivo.derecha != nil {
			hijo = objetivo.derecha
		}
		if objetivo.izquierda != nil {
			hijo = objetivo.izquierda
		}

		// si el valor a eliminar es la raiz, se elimina,
		// si no se chequea si es el hijo izquierdo o derecho
		if objetivo == arbol.raiz {
			arbol.raiz = hijo
		} else if padre.derecha == objetivo {
			padre.derecha = hijo
		} else if padre.izquierda == objetivo {
			padre.izquierda = hijo
		}
	}

	//caso 3, el nodo tiene dos hijos
	if objetivo.derecha != nil && objetivo.izquierda != nil {

		//puntero que apunta al nodo que sera el reemplazo
		var aux *Nodo
		var auxPadre *Nodo

		//nos situamos en el subarbol derecho del que va a ser eliminado
		auxPadre = objetivo
		aux = objetivo.derecha

		//encontramos el menor del subarbol derecho
		for aux.izquierda != nil {
			auxPadre = aux
			aux = aux.izquierda
		}

		//acomodamos el valor del nodo
		objetivo.valor = aux.valor

		if objetivo == arbol.raiz {
			aux = arbol.raiz
		}

		//eliminamos el nodo
		//aca verificamos si el auxiliar es sucesor
		// derecho directo del nodo a eliminar
		if auxPadre.izquierda == aux {
			auxPadre.izquierda = aux.derecha
		} else {
			auxPadre.derecha = aux.derecha
		}
	}
}

func main() {
	arbol := ABB{raiz: nil}
	arbol.insertar(50)
	arbol.insertar(48)
	arbol.insertar(17)
	arbol.insertar(68)
	arbol.insertar(99)
	arbol.insertar(75)
	arbol.insertar(5)
	arbol.insertar(51)
	arbol.insertar(76)

	arbol.buscar(5)

	arbol.eliminar(68)
	arbol.eliminar(75)
	arbol.eliminar(17)
	arbol.eliminar(5)
	arbol.eliminar(50)
}
