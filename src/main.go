package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

var producto []Producto //Aqui estan los productos a mostrar

type Producto struct {
	Id              int     `json:"Id"`
	Nombre          string  `json:"Nombre"`
	Precio          float64 `json:"Precio"`
	Stock           int     `json:"Stock"`
	Codigo          string  `json:"Codigo"`
	Publicado       bool    `json:"Publicado"`
	FechaDeCreacion string  `json:"FechaDeCreacion"`
}

func main() {
	router := gin.Default()
	router.GET("/productos", todosLosProductos)

	router.Run(":8080")

}

func readFile(name string) {
	file, err1 := os.ReadFile(name)
	if err1 != nil {
		fmt.Println("Error al leer el archivo")
		return
	}

	err2 := json.Unmarshal(file, &producto)
	if err2 != nil {
		fmt.Println("Error:", err2)
		return
	}

}

func todosLosProductos(c *gin.Context) {
	readFile("./productos.json")
	c.JSON(http.StatusOK, producto)
}
