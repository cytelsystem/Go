package main

import (
	"net/http"

	"github.com/cytelsystem/Go/tree/practicaHJM/internal/domain/producto"
	handlerProducto "github.com/cytelsystem/Go/tree/practicaHJM/cmd/server/handler/producto"
	"github.com/gin-gonic/gin"
)

const (
	port = ":8080"
)

func main() {

	router := gin.Default()

	repository := producto.NewRepository()
	service := producto.NewService(repository)
	controlador := handlerProducto.NewControladorProducto(service)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"mensage": "pong",
		})
	})

	router.GET("/productos", controlador.GetAll())
	router.GET("/productos/:id", controlador.GetByID())
	router.PUT("/productos/:id", controlador.Update())
	router.DELETE("/productos/:id", controlador.Delete())
	router.PATCH("/productos/:id", controlador.Patch())
	router.POST("/productos", controlador.Create())

	if err := router.Run(port) ; err != nil {
		panic(err)
	}

}
