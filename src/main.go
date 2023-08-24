package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var items []Item
var currentID int

func main() {
	router := gin.Default()

	router.GET("/items", GetItems)
	router.GET("/items/:id", GetItem)
	router.POST("/items", CreateItem)
	router.PUT("/items/:id", UpdateItem)
	router.DELETE("/items/:id", DeleteItem)

	router.Run(":8080")
}

func GetItems(c *gin.Context) {
	c.JSON(http.StatusOK, items)
}

func GetItem(c *gin.Context) {
	id := c.Param("id")
	for _, item := range items {
		if id == string(item.ID) {
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}

func CreateItem(c *gin.Context) {
	var newItem Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	currentID++
	newItem.ID = currentID
	items = append(items, newItem)
	c.JSON(http.StatusCreated, newItem)
}

func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var updatedItem Item
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, item := range items {
		if id == string(item.ID) {
			updatedItem.ID = item.ID
			items[i] = updatedItem
			c.JSON(http.StatusOK, updatedItem)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}

func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	for i, item := range items {
		if id == string(item.ID) {
			items = append(items[:i], items[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}
