//***************************************Utilizacion de etiquetas******************************************//

package main

import (
	"encoding/json"
	"fmt"
)

// Definición de la estructura Autor con campos Nombre y Apellido.
type Autor struct {
	Nombre   string `json:"Nombre"`
	Apellido string `json:"Apellido"`
}

// Método para obtener el nombre completo del autor.
func (a Autor) nombreAutor() string {
	return fmt.Sprintf("%s %s", a.Nombre, a.Apellido)
}

// Definición de la estructura Libro con campos Titulo, Descripcion y Autor (que es de tipo Autor).
type Libro struct {
	Titulo      string `json:"Titulo"`
	Descripcion string `json:"Descripcion"`
	Autor       Autor  `json:"Autor"`
}

// Método para imprimir la información del libro (título, descripción y autor).
func (l Libro) informacion() {
	fmt.Println(l.Titulo)
	fmt.Println(l.Descripcion)
	fmt.Println(l.Autor.nombreAutor())
}

func main() {

	// Creación de un nuevo Autor con el nombre "Hector" y el apellido "Moreno".
	nombreAutor := Autor{
		Nombre:   "Hector",
		Apellido: "Moreno",
	}

	// Creación de un nuevo Libro con título, descripción y autor previamente definidos.
	libro := Libro{
		Titulo:      "Como Aprender Go",
		Descripcion: "Clase de DH: Cómo aprender Go",
		Autor:       nombreAutor,
	}

	// Llamada al método informacion() para imprimir la información del libro.
	libro.informacion()

	// Extra. Conversión a JSON + Tags.
	{
		// Se convierte la estructura libro a su representación JSON con formato indentado.
		jsonEmpleado, err := json.MarshalIndent(libro, "", "    ")

		if err != nil {
			fmt.Println("Error is:", err)
		}

		// Se imprime el JSON resultante.
		fmt.Println(string(jsonEmpleado))
	}

}

//*******************************************************************************************************//
