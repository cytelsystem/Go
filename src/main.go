package main

import (
	"fmt"
)

type Producto struct {
	nombre   string
	precio   float64
	cantidad int
}

type Servicio struct {
	nombre            string
	precio            float64
	minutosTrabajados int
}

type Mantenimiento struct {
	nombre string
	precio float64
}

func sumarProductos(productos []Producto, ch chan<- float64) {
	total := 0.0
	for _, p := range productos {
		total += p.precio * float64(p.cantidad)
	}
	ch <- total
}

func sumarServicios(servicios []Servicio, ch chan<- float64) {
	total := 0.0
	for _, s := range servicios {
		mediaHoraTrabajada := float64(s.minutosTrabajados) / 30.0
		total += s.precio * mediaHoraTrabajada
	}
	ch <- total
}

func sumarMantenimientos(mantenimientos []Mantenimiento, ch chan<- float64) {
	total := 0.0
	for _, m := range mantenimientos {
		total += m.precio
	}
	ch <- total
}

func main() {
	productos := []Producto{
		{"Producto1", 10.0, 3},
		{"Producto2", 15.0, 2},
	}

	servicios := []Servicio{
		{"Servicio1", 5.0, 45},
		{"Servicio2", 8.0, 20},
	}

	mantenimientos := []Mantenimiento{
		{"Mantenimiento1", 50.0},
		{"Mantenimiento2", 30.0},
	}

	totalCh := make(chan float64)

	go sumarProductos(productos, totalCh)
	go sumarServicios(servicios, totalCh)
	go sumarMantenimientos(mantenimientos, totalCh)

	total := 0.0
	for i := 0; i < 3; i++ {
		total += <-totalCh
	}

	fmt.Printf("Monto final: %.2f\n", total)
}
