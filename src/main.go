//****************************Panic con defer y recover********************************//
// La declaración defer en Go se utiliza para posponer la ejecución de una función o expresión hasta que la función que la contiene haya finalizado
// recover() es una función incorporada en Go que se utiliza para manejar panics y evitar que un programa se cierre abruptamente debido a un error.

package main
import "fmt"

func isPair(num int) {

	// Dentro de la función defer, estás usando recover() para capturar cualquier valor de panic que se haya producido en el bloque de código anterior. Si se produce un panic, recover() detendrá el flujo normal del programa y devolverá el valor del panic.

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	if (num % 2) != 0 {
		panic("no es un número par")
	}else {
		fmt.Println(num, " es un número par")
	}
}

func main() {
	isPair(3)

}

//**********************************************************************//