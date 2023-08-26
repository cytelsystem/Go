// ./internal/productos/productos.go
package productos

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

var Products []Product

// Agregar funciones de manejo de productos (create, getAll, etc.)
// ...

func CreateProduct(c *gin.Context) {
	var newProduct Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Products = append(Products, newProduct)
	c.JSON(http.StatusCreated, newProduct)
}

func GetAllProducts(c *gin.Context) {
	c.JSON(http.StatusOK, Products)
}

func GetProductByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, product := range Products {
		if product.ID == id {
			c.JSON(http.StatusOK, product)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}

func BuyProduct(c *gin.Context) {
	codeValue := c.Query("code_value")
	quantityStr := c.Query("quantity")
	quantity, _ := strconv.Atoi(quantityStr)

	var selectedProduct Product
	for _, product := range Products {
		if product.CodeValue == codeValue {
			selectedProduct = product
			break
		}
	}

	if selectedProduct.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	totalPrice := float64(quantity) * selectedProduct.Price
	purchaseDetail := struct {
		Name     string  `json:"name"`
		Quantity int     `json:"quantity"`
		Total    float64 `json:"total"`
	}{
		Name:     selectedProduct.Name,
		Quantity: quantity,
		Total:    totalPrice,
	}

	c.JSON(http.StatusOK, purchaseDetail)
}

func SearchByQuantity(c *gin.Context) {
	minQuantityStr := c.Query("min")
	maxQuantityStr := c.Query("max")

	minQuantity, _ := strconv.Atoi(minQuantityStr)
	maxQuantity, _ := strconv.Atoi(maxQuantityStr)

	var filteredProducts []Product
	for _, product := range Products {
		if product.Quantity >= minQuantity && product.Quantity <= maxQuantity {
			filteredProducts = append(filteredProducts, product)
		}
	}

	c.JSON(http.StatusOK, filteredProducts)
}
