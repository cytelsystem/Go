package main

import "fmt"

type Autor struct {
	Nombre   string
	Apellido string
}

func (a Autor) nombreAutor() string {
	return fmt.Sprintf("%s %s", a.Nombre, a.Apellido)
}

type Libro struct {
	Titulo      string
	Descripcion string
	Autor       Autor
}

func (l Libro) informacion() {
	fmt.Println("Titulo:", l.Titulo)
	fmt.Println("Descripcion:", l.Descripcion)
	fmt.Println("Autor:", l.Autor.nombreAutor())
}

func main() {

	nombreAutor := Autor{
		Nombre:   "hector",
		Apellido: "Moreno",
	}

	libro := Libro{
		Titulo:      "como Apender go",
		Descripcion: "clase de dh como aprender go",
		Autor:       nombreAutor,
	}

	libro.informacion()

}
