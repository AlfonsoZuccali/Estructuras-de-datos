package main

import "fmt"

//definicion del Nodo
type Nodo struct {
	valor     int
	siguiente *Nodo
}

//la cola tiene un fondo, un tope y un tamaño
type Cola struct {
	tope  *Nodo
	fondo *Nodo
	size  int
}

//pushea los valores a la cola
func (c *Cola) push(elemento int) {
	//creamos el nuevo nodo
	nuevoNodo := &Nodo{valor: elemento}

	//si la cola esta vacia
	if c.tope == nil {
		//el nuevo nodo es el unico elemento y por ende
		// es el tope y el fondo al mismo tiempo
		c.tope = nuevoNodo
		c.fondo = nuevoNodo

		//si la cola no esta vacia
	} else {
		//el nuevo nodo es el nuevo fondo y debemos
		//conectarlo a la lista
		c.fondo.siguiente = nuevoNodo
		c.fondo = nuevoNodo
	}

	//aumenta el tamaño de la cola
	c.size++
}

//sacamos el frente de la cola
func (c *Cola) pop() int {

	//si la cola esta vacia
	if c.tope == nil {
		fmt.Println("La cola esta vacia, para sacar un valor, primero agrega valores a la cola")
		return -1
	}

	//guardamos el valor del nodo del frente
	//para poder retornarlo
	valor := c.tope.valor

	//hacemos que el frente sea el nodo siguiente
	//al frente anterior
	c.tope = c.tope.siguiente

	//disminuimos el tamaño de la cola
	c.size--

	//retornamos el valor almacenado en el frente
	return valor
}

func (c *Cola) getTop() int {

	//si la cola esta vacia
	if c.tope == nil {
		fmt.Println("La cola esta vacia")
		return -1
	}

	//retornamos el valor del nodo del frente
	return c.tope.valor
}

func (c *Cola) getBack() int {

	//si la cola esta vacia
	if c.fondo == nil {
		fmt.Println("La cola esta vacia")
		return -1
	}

	//retornamos el valor del nodo del fondo
	return c.fondo.valor
}

func (c *Cola) getSize() int {

	//retornamos el tamaño de la cola
	return c.size
}

func main() {

	// Creamos una pila vacía
	p := &Cola{}
	fmt.Println("Cola recién creada:")

	// Verifica comportamiento con cola vacía
	fmt.Println("Tamaño:", p.getSize())
	fmt.Println("Tope:", p.getTop())
	fmt.Println("Fondo:", p.getBack())

	//enviamos valores a la cola
	p.push(10)
	p.push(20)
	p.push(30)
	p.push(40)
	p.push(50)
	p.push(60)
	p.push(70)

	//verificamos comportamiento
	fmt.Println("Tamaño actual:", p.getSize())
	fmt.Println("Elemento en el tope:", p.getTop())
	fmt.Println("Elemento en el fondo:", p.getBack())

	p.pop()
	fmt.Println("Nuevo tope:", p.getTop())
	fmt.Println("Tamaño:", p.getSize())
	p.pop()
	p.pop()
	p.pop()

	fmt.Println("Tope actual:", p.getTop())
	fmt.Println("Tamaño final:", p.getSize())
}
