package main

import "fmt"

type Producto interface {
	Precio() float64
}

type Pequeno struct {
	precio float64
}

func (p Pequeno) Precio() float64 {
	return p.precio
}

type Mediano struct {
	precio float64
}

func (m Mediano) Precio() float64 {
	return m.precio * 3/100 + m.precio
}

type Grande struct {
	precio float64
}

func (g Grande) Precio() float64 {
	return g.precio *6 /100 + g.precio + 2500
}

func CrearProducto(tipo string, precio float64) Producto {
	switch tipo {
	case "Pequeno":
		return Pequeno{precio: precio}
	case "Mediano":
		return Mediano{precio: precio}
	case "Grande":
		return Grande{precio: precio}
	default:
		panic("Tipo de producto no válido")

	}


}

func main() {
producto1 := CrearProducto("Pequeno", 100000)
fmt.Println(producto1.Precio())
producto2 := CrearProducto("Mediano", 100000)
fmt.Println(producto2.Precio())
producto3 := CrearProducto("Grande", 100000)
fmt.Println(producto3.Precio())

}
