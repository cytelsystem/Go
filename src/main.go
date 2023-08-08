package main

import (
	"fmt"
	"os"
)

func readDataFromFile(filename string) (string, error) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	content, err := os.ReadFile(filename)
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}

	return string(content), nil
}

func main() {
	filename := "customers.txt"
	data, err := readDataFromFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)
	fmt.Println("Ejecución finalizada")
}


