package main

import (
	"fmt"
	"os"
	"strings"
)

type Producto struct {
	codigo   int
	nombre   string
	precio   float32
	cantidad int
}

type Categoria struct {
	codigo    string
	nombre    string
	productos []Producto
}

func agregarProductoACategoria(c *Categoria, producto Producto) {
	c.productos = append(c.productos, producto)
}

func archivoCategorias(categorias ...Categoria) error {
	f, err := os.Create("categorias.csv")
	if err != nil {
		return err
	}
	defer f.Close()

	var datos string = "ID, categoria, productos;\n"
	for _, categoria := range categorias {
		productosStr := []string{}
		for _, producto := range categoria.productos {
			productosStr = append(productosStr, fmt.Sprintf("%v", producto))
		}
		datos += fmt.Sprintf("%v,%v,%v;\n", categoria.codigo, categoria.nombre, strings.Join(productosStr, ", "))
	}

	_, err = f.WriteString(datos)
	if err != nil {
		return err
	}

	data, err := os.ReadFile("./categorias.csv")
	if err != nil {
		return err
	}
	fmt.Println(strings.Split(string(data), ";\n"))
	return nil
}

func main() {
	cat1 := Categoria{"001", "Electro", nil}
	cat2 := Categoria{"002", "Pinturas", nil}
	cat3 := Categoria{"003", "Automor", nil}

	prod1 := Producto{1, "computador", 10.0, 2}
	prod2 := Producto{2, "celular", 1.0, 1}
	prod3 := Producto{3, "fiat", 5.0, 5}
	prod4 := Producto{4, "chevrolet", 10.0, 10}
	prod5 := Producto{5, "amarilla", 10.0, 10}

	agregarProductoACategoria(&cat1, prod1)
	agregarProductoACategoria(&cat1, prod2)
	agregarProductoACategoria(&cat3, prod3)
	agregarProductoACategoria(&cat3, prod4)
	agregarProductoACategoria(&cat2, prod5)

	// err := archivoCategorias(cat1, cat2, cat3)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	err := archivoCategorias(cat1, cat2, cat3)
	if err != nil {
		fmt.Println(err)
	}

	err = escribirEstimativos(cat1, cat2, cat3)
	if err != nil {
		fmt.Println(err)
	}
}

func escribirEstimativos(categorias ...Categoria) error {
	f, err := os.Create("estimaciones.csv")
	if err != nil {
		return err
	}
	defer f.Close()

	var estimativos string = "Categoría\tEstimativoPorCategoría\n"
	totalEstimativo := 0.0
	for _, categoria := range categorias {
		estimativoCategoria := 0.0
		for _, producto := range categoria.productos {
			estimativoCategoria += float64(producto.precio) * float64(producto.cantidad)
		}
		estimativos += fmt.Sprintf("%s\t%.2f\n", categoria.nombre, estimativoCategoria)
		totalEstimativo += estimativoCategoria
	}
	estimativos += fmt.Sprintf("TotalEstimativo\t%.2f\n", totalEstimativo)

	_, err = f.WriteString(estimativos)
	if err != nil {
		return err
	}

	return nil
}


