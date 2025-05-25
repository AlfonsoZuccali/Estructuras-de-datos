package main

//Nodo
type Nodo struct {
	izquierda *Nodo
	valor     int
	derecha   *Nodo
	altura    int
}

//arbol binario de busqueda
type ABB struct {
	raiz *Nodo
}

//funcion que devuelve el mayor de dos valores
func maximo(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

//definicion de altura de un nodo
func altura(nodo *Nodo) int {
	if nodo == nil {
		return 0
	}
	return nodo.altura
}

//verifica si un arbol esta vacio
func (arbol *ABB) EsVacio() bool {
	return arbol.raiz == nil
}

//metodo que devuelve el nodo con el menor valor del arbol
func menor_nodo(nodo *Nodo) *Nodo {
	if nodo == nil {
		return nil
	} else if nodo.izquierda == nil {
		return nodo
	}
	return menor_nodo(nodo.izquierda)
}

//funcion que actualiza la altura despues de eliminar o agregar nodos al arbol
func actualizarAltura(nodo *Nodo) {
	if nodo == nil {
		return
	} else {
		//la altura se actualiza a la suma de la altura del nodo actual mas la altura mas grande de sus hijos
		nodo.altura = 1 + (maximo(altura(nodo.izquierda), altura(nodo.derecha)))
		return
	}
}

//saca el factor de balanceo del nodo dado
func factor_balanceo(nodo *Nodo) int {
	return altura(nodo.izquierda) - altura(nodo.derecha)
}

//funcion de rotacion derecha
func rotacion_derecha(a *Nodo) *Nodo {
	//establecemos el subarbol de referencia
	b := a.izquierda
	c := b.derecha
	//cambiamos de lugar los nodos
	b.derecha = a
	a.izquierda = c

	//actualizamos la altura solo a y b ya que c se mantiene a la misma altura
	actualizarAltura(a)
	actualizarAltura(b)
	return b
}

//funcion de rotacion izquierda
func rotacion_izquierda(a *Nodo) *Nodo {
	//establecemos el subarbol de referencia
	b := a.derecha
	c := b.izquierda
	//cambiamos de lugar los nodos
	b.izquierda = a
	a.derecha = c

	//actualizamos la altura solo a y b ya que c se mantiene a la misma altura
	actualizarAltura(a)
	actualizarAltura(b)
	return b
}

//funcion de balanceo que devuelve la nueva raiz
func balancear(a *Nodo) *Nodo {
	actualizarAltura(a)

	//desbalanceo a la derecha
	if factor_balanceo(a) < -1 {
		if a.derecha != nil && factor_balanceo(a.derecha) > 0 {
			a.derecha = rotacion_izquierda(a.derecha)
		}
		a = rotacion_izquierda(a)

		//desbalanceo a la izquierda
	} else if factor_balanceo(a) > 1 {
		if a.izquierda != nil && factor_balanceo(a.izquierda) < 0 {
			a.izquierda = rotacion_derecha(a.izquierda)
		}
		a = rotacion_derecha(a)
	}

	return a
}

func insertar_recursiva(nodo *Nodo, valor int) *Nodo {
	//caso base
	if nodo == nil {
		return &Nodo{valor: valor, altura: 1}
	}
	// si el valor insertado es menor al del nodo del arbol
	if valor < nodo.valor {
		nodo.izquierda = insertar_recursiva(nodo.izquierda, valor)
	}

	//si el valor insertado es mayor al del nodo del arbol
	if valor > nodo.valor {
		nodo.derecha = insertar_recursiva(nodo.derecha, valor)
	}

	return balancear(nodo)
}

//metodo para insertar valores en un arbol
func (arbol *ABB) insertar(input int) {
	//llamada al insertar recursivamente
	arbol.raiz = insertar_recursiva(arbol.raiz, input)
}

//metodo para buscar un elemento en el ABB
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

func eliminar_recursiva(nodo *Nodo, valor int) *Nodo {
	//caso base
	if nodo == nil {
		return nil
	}

	//busqueda recursiva del objetivo
	if valor < nodo.valor {
		nodo.izquierda = eliminar_recursiva(nodo.izquierda, valor)
	} else if valor > nodo.valor {
		nodo.derecha = eliminar_recursiva(nodo.derecha, valor)

		// si no es mayor ni menor, es igual, por ende encontramos el nodo a eliminar
	} else {
		//Caso sin hijos
		if nodo.derecha == nil && nodo.izquierda == nil {
			return nil

			//Caso con un hijo
		} else if (nodo.izquierda == nil) != (nodo.derecha == nil) {
			if nodo.izquierda == nil {
				return nodo.derecha
			} else {
				return nodo.izquierda
			}

			//caso con dos hijos
		} else if nodo.izquierda != nil && nodo.derecha != nil {
			aux := menor_nodo(nodo.derecha)
			nodo.valor = aux.valor
			nodo.derecha = eliminar_recursiva(nodo.derecha, aux.valor)
		}
	}
	return balancear(nodo)
}

//metodo que elimina valores en un arbol
func (arbol *ABB) eliminar(input int) {
	arbol.raiz = eliminar_recursiva(arbol.raiz, input)
}

func main() {

	//test de funciones
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
