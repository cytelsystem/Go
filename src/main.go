package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	fileName = "./data.csv"
)

func main() {

	// crearArchivo(fileName)
	readFile(fileName)

}

func readFile(name string) {
	file, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("Error al leer el archivo")
		return
	}

	data := strings.Split(string(file), ";")

	var total float64

	for i := 0; i < len(data)-1; i++ {
		var line = strings.Split(string(data[i]), ",")

		if i != 0 {
			precio, err := strconv.ParseFloat(line[1],64)
			if err != nil {
				log.Println("No se pudo parsear el precio")
			}
			cantidad, err2 := strconv.ParseFloat(line[2],64)
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

func crearArchivo(name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("No se pudo crear el archivo")
		return
	}

	defer file.Close()
	texto := "ID,Precio,Cantidad;0001,10,2;0002,15,1;0003,1,5;"

	_, err = file.WriteString(texto)
	if err != nil {
		fmt.Println("no se pudo adicionar el texto")
	}

	log.Println("Archivo creado correctamente", file, err)
}
