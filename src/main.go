package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	crearArchivo("./prueba.txt")
	// readFile("prueba.txt")

}

func readFile(name string) {
	file, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("Error al leer el archivo")
		return
	}

	fmt.Println(file)
}

func crearArchivo(name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("No se pudo crear el archivo")
		return
	}

	defer file.Close()
	texto := "Hola"

	_, err = file.WriteString(texto)
	if err != nil {
		fmt.Println("no se pudo adicionar el texto")
	}

	log.Println("Archivo creado correctamente", file, err)
}
