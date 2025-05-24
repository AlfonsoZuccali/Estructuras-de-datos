# Guía completa de tu Árbol AVL en **Go**  
A continuación, encontrarás todas las funciones del archivo estructuradas en _Markdown_ con secciones claras y algunos **tips** para recordarlas mejor.

---

## 1. Estructuras principales

```go
// Definición del nodo
type Nodo struct {
    izquierda *Nodo
    valor     int
    derecha   *Nodo
    altura    int
}

// Definición del árbol
type ABB struct {
    raiz *Nodo
}
```

- **Nodo**: contiene el valor, referencias a los hijos izquierdo y derecho, además de la altura.  
- **ABB**: encapsula la raíz del árbol.

---

## 2. Función máxima y altura

```go
func maximo(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func altura(nodo *Nodo) int {
    if nodo == nil {
        return 0
    }
    return nodo.altura
}
```

- **`maximo`**: obtiene el mayor de dos números.  
- **`altura`**: retorna la altura de un nodo; si es `nil`, devuelve `0`.

---

## 3. Verificar si el árbol está vacío

```go
func (arbol *ABB) EsVacio() bool {
    return arbol.raiz == nil
}
```

- Retorna `true` si la raíz es `nil`.

---

## 4. Buscar el nodo con el menor valor

```go
func menor_nodo(nodo *Nodo) *Nodo {
    if nodo == nil {
        return nil
    } else if nodo.izquierda == nil {
        return nodo
    }
    return menor_nodo(nodo.izquierda)
}
```

- Desciende recursivamente por la izquierda hasta hallar el nodo más pequeño.  
- Ideal para usar al eliminar un nodo con dos hijos.

---

## 5. Actualizar la altura del nodo

```go
func actualizarAltura(nodo *Nodo) {
    if nodo == nil {
        return
    }
    nodo.altura = 1 + maximo(
        altura(nodo.izquierda),
        altura(nodo.derecha),
    )
}
```

- Al insertar o eliminar, recalcula la altura como `1 + max(hijoIzq, hijoDer)`.

---

## 6. Factor de balanceo

```go
func factor_balanceo(nodo *Nodo) int {
    return altura(nodo.izquierda) - altura(nodo.derecha)
}
```

- **Factor de balanceo (FB)** = altura subárbol izquierdo − altura subárbol derecho.  
- Para un **AVL**, FB debe estar en `{-1, 0, 1}`.

---

## 7. Rotaciones

### Rotación Derecha

```go
func rotacion_derecha(a *Nodo) *Nodo {
    b := a.izquierda
    c := b.derecha
    b.derecha = a
    a.izquierda = c

    actualizarAltura(a)
    actualizarAltura(b)
    return b
}
```

### Rotación Izquierda

```go
func rotacion_izquierda(a *Nodo) *Nodo {
    b := a.derecha
    c := b.izquierda
    b.izquierda = a
    a.derecha = c

    actualizarAltura(a)
    actualizarAltura(b)
    return b
}
```

- **Objetivo**: corregir desequilibrios girando subárboles.  
- **Rotación Derecha**: cuando subárbol izquierdo es más pesado.  
- **Rotación Izquierda**: cuando subárbol derecho es más pesado.

---

## 8. Balanceo del nodo

```go
func balancear(a *Nodo) *Nodo {
    actualizarAltura(a)

    // Desbalanceo a la derecha
    if factor_balanceo(a) < -1 {
        if a.derecha != nil && factor_balanceo(a.derecha) > 0 {
            a.derecha = rotacion_izquierda(a.derecha)
        }
        a = rotacion_izquierda(a)

    // Desbalanceo a la izquierda
    } else if factor_balanceo(a) > 1 {
        if a.izquierda != nil && factor_balanceo(a.izquierda) < 0 {
            a.izquierda = rotacion_derecha(a.izquierda)
        }
        a = rotacion_derecha(a)
    }

    return a
}
```

> **TIP**:  
> - Primero se actualiza la altura.  
> - Revisa el factor de balanceo:  
>   - Si FB < -1 → subárbol derecho está muy largo.  
>   - Si FB > 1 → subárbol izquierdo está muy largo.  
> - Aplica rotaciones simples o dobles según corresponda.

---

## 9. Inserción recursiva

```go
func insertar_recursiva(nodo *Nodo, valor int) *Nodo {
    if nodo == nil {
        return &Nodo{valor: valor, altura: 1}
    }
    if valor < nodo.valor {
        nodo.izquierda = insertar_recursiva(nodo.izquierda, valor)
    } else if valor > nodo.valor {
        nodo.derecha = insertar_recursiva(nodo.derecha, valor)
    }
    return balancear(nodo)
}
```

```go
func (arbol *ABB) insertar(input int) {
    arbol.raiz = insertar_recursiva(arbol.raiz, input)
}
```

> **TIP**:  
> - Recorres recursivamente hasta encontrar la posición correcta.  
> - Al retornar en la recursión, llamas a `balancear`.  
> - El método público actualiza `arbol.raiz` con el nuevo nodo raíz balanceado.

---

## 10. Búsqueda

```go
func (arbol *ABB) buscar(input int) bool {
    objetivo := arbol.raiz

    if arbol.EsVacio() {
        return false
    }
    for objetivo.valor != input && (objetivo.izquierda != nil || objetivo.derecha != nil) {
        if objetivo.valor > input && objetivo.izquierda != nil {
            objetivo = objetivo.izquierda
        } else if objetivo.valor < input && objetivo.derecha != nil {
            objetivo = objetivo.derecha
        }
    }
    return (objetivo.valor == input)
}
```

- Baja iterativamente por la izquierda o derecha según el valor buscado.  
- Devuelve `true` si encuentra el valor, `false` en caso contrario.

---

## 11. Eliminación recursiva

```go
func eliminar_recursiva(nodo *Nodo, valor int) *Nodo {
    if nodo == nil {
        return nil
    }

    if valor < nodo.valor {
        nodo.izquierda = eliminar_recursiva(nodo.izquierda, valor)
    } else if valor > nodo.valor {
        nodo.derecha = eliminar_recursiva(nodo.derecha, valor)
    } else {
        // Caso sin hijos
        if nodo.derecha == nil && nodo.izquierda == nil {
            return nil
        } 
        // Caso con un solo hijo
        else if (nodo.izquierda == nil) != (nodo.derecha == nil) {
            if nodo.izquierda == nil {
                return nodo.derecha
            } else {
                return nodo.izquierda
            }
        } 
        // Caso con dos hijos
        else if nodo.izquierda != nil && nodo.derecha != nil {
            aux := menor_nodo(nodo.derecha)
            nodo.valor = aux.valor
            nodo.derecha = eliminar_recursiva(nodo.derecha, aux.valor)
        }
    }
    return balancear(nodo)
}

func (arbol *ABB) eliminar(input int) {
    arbol.raiz = eliminar_recursiva(arbol.raiz, input)
}
```

> **TIP**:  
> - Si el nodo **no existe** → retorna `nil`.  
> - **Caso sin hijos**: retornas `nil`, eliminando el nodo.  
> - **Caso un hijo**: retornas el hijo no nulo.  
> - **Caso dos hijos**: copias el valor del sucesor (menor del subárbol derecho) y eliminas el sucesor.  
> - Finalmente, **balanceas** el nodo y lo retornas.

---

## 12. Función Principal

```go
func main() {
    arbol := ABB{raiz: nil}
    arbol.insertar(50)
    arbol.insertar(48)
    // ...
    arbol.eliminar(68)
    arbol.eliminar(75)
    // ...
}
```

- Inserta, busca y elimina valores para **testear** el árbol AVL.

---

# ¡Listo!  
Con este formato, puedes repasar cada sección de tu AVL en **Go** de manera más ordenada y recordar los **tips** clave para sus algoritmos de balanceo, inserción y eliminación. ¡Éxitos en tu examen!