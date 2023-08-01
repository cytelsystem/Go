package main

import "fmt"

// Product representa un producto con sus campos respectivos.
type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

// Products es un slice global de Product, inicializado con valores de ejemplo.
var Products = []Product{
	{ID: 1, Name: "Producto 1", Price: 10.99, Description: "Descripción del producto 1", Category: "Categoría 1"},
	{ID: 2, Name: "Producto 2", Price: 19.99, Description: "Descripción del producto 2", Category: "Categoría 2"},
	{ID: 3, Name: "Producto 3", Price: 5.99, Description: "Descripción del producto 3", Category: "Categoría 1"},
}

// Save agrega el producto actual al slice global Products.
func (p Product) Save() {
	Products = append(Products, p)
}

// GetAll imprime todos los productos guardados en el slice global Products.
func (p Product) GetAll() {
	fmt.Println("Lista de productos:")
	for _, product := range Products {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f, Description: %s, Category: %s\n",
			product.ID, product.Name, product.Price, product.Description, product.Category)
	}
}

// getById recibe un ID como parámetro y devuelve el producto correspondiente si se encuentra.
func getById(id int) (Product, error) {
	for _, product := range Products {
		if product.ID == id {
			return product, nil
		}
	}
	return Product{}, fmt.Errorf("Producto no encontrado con ID: %d", id)
}

func main() {
	// Crear un nuevo producto y guardarlo.
	newProduct := Product{
		ID:          4,
		Name:        "Nuevo Producto",
		Price:       15.50,
		Description: "Descripción del nuevo producto",
		Category:    "Categoría 2",
	}
	newProduct.Save()

	// Imprimir todos los productos.
	newProduct.GetAll()

	// Obtener un producto por ID y mostrarlo.
	idToSearch := 2
	product, err := getById(idToSearch)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Producto encontrado con ID %d:\n", idToSearch)
		fmt.Printf("ID: %d, Name: %s, Price: %.2f, Description: %s, Category: %s\n",
			product.ID, product.Name, product.Price, product.Description, product.Category)
	}
}
