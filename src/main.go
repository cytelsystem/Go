package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	fileName = "./data.csv"
)

func main() {

	// Uso del envaironment para setear variables.
	os.Setenv("usuario", "supervisor")

	usuario := os.Getenv("usuario")

	if usuario == "admin" {
		readFile(fileName)
	} else {
		fmt.Println("Este usuario no es administrador")
		fmt.Println("Este usuario tiene privilegios de:", usuario)
	}

}

func readFile(name string) {

	file, err := os.ReadFile(name)

	if err != nil {
		log.Println("No se pudo leer el archivo")
	}

	data := strings.Split(string(file), ";")

	var total float64
	for i := 0; i < len(data)-1; i++ {
		var line = strings.Split(string(data[i]), ",")

		if i != 0 {
			precio, err := strconv.ParseFloat(line[1], 64)
			if err != nil {
				log.Println("No se pudo parsear el precio")
			}
			cantidad, err2 := strconv.ParseFloat(line[2], 64)
			if err2 != nil {
				log.Println("No se pudo parsear la cantidad")
			}

			totalProducto := precio * cantidad
			total += totalProducto
		}

		for i := 0; i < len(line); i++ {
			fmt.Printf("%s\t\t", line[i])
			if i == len(line)-1 {
				fmt.Print("\n")
			}
		}
	}

	fmt.Printf("\nTotal\t\t%.2f\n", total)

}