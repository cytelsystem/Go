//***************************************Punteros******************************************//

package main
import "fmt"
// La función incrementar recibe un puntero de tipo entero
func Incrementar(v *int) {
 // Desreferenciamos la variable v con el operador “*”
 // para obtener su valor e incrementarlo en 1
 *v++ //el * es convierta de nuevo al dato que puedo sumar
}
func main() {
 var v int = 50
 // La función Incrementar recibe un puntero
 // utilizamos el operador de dirección “&”
 // para obtener la dirección de memoria de v
 Incrementar(&v) //el & es utilice la dirección de memoria
 fmt.Println("El valor de v ahora vale: ", v)
}
//*******************************************************************************************************//
