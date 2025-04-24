package main

type Nodo struct {
	caracter rune
	//usamos siguiente y anterior, siendo siguiente, el que se acerca
	//a la cola, mientras que el anterior se acerca a la cabeza
	siguiente *Nodo
	anterior  *Nodo
}

type ListaSimple struct {
}

func main() {

}
