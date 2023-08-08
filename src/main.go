//****************funciones multiretorno para manejo de errores en go********************************//
// Las funciones de múltiple retorno son muy útiles en Go, especialmente para el manejo de errores.

package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("No se puede dividir por cero")
	}
	return a / b, nil
}

func main() {
	num1, num2 := 10.0, 2.0
	result, err := divide(num1, num2)
	if err != nil {
		fmt.Println("Error al dividir:", err)
		return
	}
	fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, result)
}



//**********************************************************************//