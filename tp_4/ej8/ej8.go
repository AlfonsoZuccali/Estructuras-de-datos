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

	//caso 3, el nodo tiene dos hijos
	if observador.derecha != nil && observador.izquierda != nil {

		//puntero que apunta al nodo que sera el reemplazo
		var aux *Nodo
		var auxPadre *Nodo

		//nos situamos en el subarbol derecho del que va a ser eliminado
		auxPadre = observador
		aux = observador.derecha

		//encontramos el menor del subarbol derecho
		for aux.izquierda != nil {
			auxPadre = aux
			aux = aux.izquierda
		}

		//acomodamos el valor del nodo
		observador.valor = aux.valor

		if observador == arbol.raiz {
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
