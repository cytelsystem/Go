// *********************Calcular salario**************************************//
// **salarioMensual
// **cantidadHoras
// **salarioPorHora
package main

import "fmt"

func calcularSalario(cantidaHoras int, categoria string) float64{
	var salarioPorHora float64
	var salarioMensual float64

	switch categoria{
	case "A":
		salarioPorHora = 3000.0
	case "B":
		salarioPorHora = 1500.0
	case "C":
		salarioPorHora = 1000.0
	default:
		fmt.Println("no existe esa categoria")
		return 0
	}

	salarioMensual = float64(cantidaHoras) * salarioPorHora / 60.0

	if categoria == "A" {
		salarioMensual += salarioMensual * 0.5
	} else if categoria == "B" {
		salarioMensual += salarioMensual * 0.2
	}
	return salarioMensual
}

func main() {

 valorTotal := calcularSalario(2400, "B")
	fmt.Printf("El salario mensual es: $%.2f\n", valorTotal)

}

//***************************************************************************************************//
