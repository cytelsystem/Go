// main.go
package main

import (
	"github.com/cytelsystem/Go/internal/productos"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Agregar productos de prueba
	productos.Products = append(productos.Products,
		productos.Product{ID: 1, Name: "Cheese - St.Andre", Quantity: 60, CodeValue: "S73191A", IsPublished: true, Expiration: "12/04/2022", Price: 50.15},
		productos.Product{ID: 2, Name: "Example Product 2", Quantity: 150, CodeValue: "C12345B", IsPublished: true, Expiration: "01/01/2023", Price: 25.99},
		productos.Product{ID: 3, Name: "Sample Item", Quantity: 400, CodeValue: "D98765X", IsPublished: false, Expiration: "05/15/2023", Price: 10.5},
	)

	r.GET("/products", productos.GetAllProducts)
	r.GET("/products/:id", productos.GetProductByID)
	r.GET("/productparams", productos.GetProductParams)
	r.GET("/buy", productos.BuyProduct)
	r.GET("/searchbyquantity", productos.SearchByQuantity)

	r.Run(":8080")
}
