package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Nombre    string `json:"nombre"`
	Apellidos string `json:"apellidos"`
	Edad      string `json:"edad"`
	Profesion string `json:"profesion"`
}

func main() {
	jsonData := `{"nombre":"Hector","apellidos":"Moreno","edad":"47 años","profesion":"programador"}`

	var person Person
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Nombre:", person.Nombre)
	fmt.Println("Apellidos:", person.Apellidos)
	fmt.Println("Edad:", person.Edad)
	fmt.Println("Profesión:", person.Profesion)
}
