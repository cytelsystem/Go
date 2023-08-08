// ***************************************************************//
// Error() metodo error de la interface error
// fmt.Errorf() del paquete fmt
// Libreria errors permite estas 4 opciones New(), is(), As() y Unwrap(). EJM: errors.New()
// **cadena de errores
//	error_1 := fmtErrorf("primer error")
//	error_2 := fmtErrorf("segundo error %w", error_1)  //incluir el %w para llamar el otro error

// Usamos errors.New() para crear un nuevo error simple y lo asignamos a la variable err.
// Usamos errors.Is() para verificar si el error err es igual al error que creamos previamente. Si es igual, mostramos un mensaje diciendo que son del mismo tipo.
// Creamos un error que envuelve al error original utilizando fmt.Errorf(). Usamos la cadena "%w" para indicar que queremos envolver un error en este lugar.
// Utilizamos errors.Is() nuevamente para verificar si el error envuelto contiene el error original. Si es cierto, mostramos un mensaje.
// Utilizamos errors.As() para intentar extraer un error relacionado. Si es exitoso, mostramos el mensaje de error relacionado.
// Utilizamos errors.Unwrap() para obtener el error relacionado sin el envoltorio. Si existe, mostramos el mensaje del error relacionado.
// Hemos reemplazado el término "subyacente" por "relacionado" para hacer la explicación más clara y comprensible. Este código muestra cómo trabajar con errores en Go utilizando las funciones proporcionadas por la librería errors

package main

import (
	"errors"
	"fmt"
)

func main() {
	// Creamos un nuevo error simple utilizando la función New()
	err := errors.New("Hubo un problema en la conexión")

	// Verificamos si el error es igual al error específico que creamos antes
	if errors.Is(err, errors.New("Hubo un problema en la conexión")) {
		fmt.Println("El error es del mismo tipo")
	}

	// Creamos un nuevo error que envuelve el error original
	errorEnvuelto := fmt.Errorf("Ocurrió un problema: %w", err)

	// Verificamos si el error envuelto contiene el error original
	if errors.Is(errorEnvuelto, err) {
		fmt.Println("El error envuelto contiene el error original")
	}

	// Intentamos extraer el error relacionado utilizando As()
	var errComun error
	if errors.As(errorEnvuelto, &errComun) {
		fmt.Println("Error relacionado:", errComun.Error())
	}

	// Desempaquetamos el error para obtener el error relacionado
	errRelacionado := errors.Unwrap(errorEnvuelto)
	if errRelacionado != nil {
		fmt.Println("Error relacionado desempaquetado:", errRelacionado.Error())
	}
}

//***************************************************************//
