package main

//Nodo
type Nodo struct {
	valor     int
	izquierda *Nodo
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

func main() {
	arbol := ABB{raiz: nil}
	arbol.insertar(5)
	arbol.insertar(10)
	arbol.insertar(8)
	arbol.insertar(3)
	arbol.insertar(12)
	arbol.insertar(0)
}
