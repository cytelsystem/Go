//***************************************interfaces******************************************//

package main

import "fmt"

type Animal interface {
	Sonido() string
}

type Perro struct{}

func (p Perro) Sonido() string {
	return "Guau guau"
}

func main() {
	var a Animal = Perro{}

	// Comprobamos si el valor de interfaz a implementa la interfaz Animal
	if perro, ok := a.(Animal); ok {
			fmt.Println("El valor de interfaz a es un Perro y hace:", perro.Sonido())
	} else {
			fmt.Println("La aserción falló.")
	}
}


//*******************************************************************************************************//
