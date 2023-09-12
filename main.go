package main

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/cytelsystem/Go/tree/practicaHJM/internal/producto"
	"github.com/gin-gonic/gin"
)

const (
	port = ":8080"
)

func main() {

	storage := producto.Storage{
		Productos: loadData(),
	}

	storage.PrintInfo()

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"mensage": "pong",
		})
	})

	router.GET("/productos/search", func(ctx *gin.Context) {
		precioQuery := ctx.Query("priceGt")
		user := ctx.Query("user")

		if precioQuery != "" {
			precio, err := strconv.ParseFloat(precioQuery, 64)
			if err != nil {
				ctx.JSON(400, gin.H{
					"mensaje": "Precio invalido",
				})
				return
			}
			data := storage.GetProductosMayorPrecio(precio)
			ctx.JSON(200, gin.H{
				"data": data,
			})

			return
		}
		nuevoContexto := addToContext(ctx, user)
		ctx.JSON(200, gin.H{
			"data": storage.GetAll(nuevoContexto),
		})
	})

	router.GET("/productparams", func(ctx *gin.Context) {
		ID, _ := strconv.Atoi(ctx.Query("id"))
		Name := ctx.Query("name")
		Quantity, _ := strconv.Atoi(ctx.Query("quantity"))
		CodeValue := ctx.Query("codeValue")
		IsPublished, _ := strconv.ParseBool(ctx.Query("isPublished"))
		ExpirationStr := ctx.Query("expiration")
		Price, _ := strconv.ParseFloat(ctx.Query("price"), 64)

		product := producto.Producto{
			ID:          ID,
			Name:        Name,
			Quantity:    Quantity,
			CodeValue:   CodeValue,
			IsPublished: IsPublished,
		}

		// Maneja la fecha de expiración (Expiration)
		expirationTime, err := time.Parse("2006-01-02", ExpirationStr)
		if err == nil {
				product.Expiration = expirationTime
		}

		product.Price = Price

    // Devuelve el producto en formato JSON
    ctx.JSON(http.StatusOK, product)
	})

	router.Run(port)
}

func loadData() []producto.Producto {
	productos := []producto.Producto{
		{
			ID:          1,
			Name:        "Banana",
			Quantity:    10,
			CodeValue:   "AABBCCC",
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       10.0,
		},
		{
			ID:          2,
			Name:        "Manzana",
			Quantity:    5,
			CodeValue:   "AABBDDD",
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       5.0,
		},
		{
			ID:          3,
			Name:        "Pera",
			Quantity:    8,
			CodeValue:   "AAAZZZCCC",
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       8.0,
		},
	}

	return productos

}

func addToContext(ctx context.Context, user string) context.Context {
	nuevoContexto := context.WithValue(ctx, "user", user)

	return nuevoContexto
}
