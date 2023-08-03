//***************************************interfaces******************************************//

// Una empresa de redes sociales requiere implementar una estructura usuarios con
// funciones que vayan agregando información a la misma. Para optimizar y ahorrar memoria
// requieren que la estructura usuarios ocupe el mismo lugar en memoria para el main del
// programa y para las funciones. La estructura debe tener los campos: nombre, apellido,
// edad, correo y contraseña. Y deben implementarse las funciones:
// ● cambiarNombre: permite cambiar el nombre y apellido.
// ● cambiarEdad: permite cambiar la edad.
// ● cambiarCorreo: permite cambiar el correo.
// ● cambiarContraseña: permite cambiar la contraseña.

package main

import (
	"encoding/json"
	"fmt"
)

// Definición de la estructura Usuario
type Usuario struct {
	Nombre     string `json:"Nombre"`
	Apellido   string `json:"Apellido"`
	Edad       int    `json:"Edad"`
	Correo     string `json:"Email"`
	Contraseña string `json:"Constraseña"`
}

// Interfaz para las funciones de cambio de datos del usuario
type Modificable interface {
	cambiarNombre(nombre, apellido string)
	cambiarEdad(edad int)
	cambiarCorreo(correo string)
	cambiarContraseña(contraseña string)
}

// Función para cambiar el nombre y apellido del usuario
func (u *Usuario) cambiarNombre(nombre, apellido string) {
	u.Nombre = nombre
	u.Apellido = apellido
}

// Función para cambiar la edad del usuario
func (u *Usuario) cambiarEdad(edad int) {
	u.Edad = edad
}

// Función para cambiar el correo del usuario
func (u *Usuario) cambiarCorreo(correo string) {
	u.Correo = correo
}

// Función para cambiar la contraseña del usuario
func (u *Usuario) cambiarContraseña(contraseña string) {
	u.Contraseña = contraseña
}

func convertirJson(usuario Usuario){
	jsonUsuario, err := json.MarshalIndent(usuario, "", "    ")

	if err != nil {
		fmt.Println("Error is:", err)
	}
	fmt.Println(string(jsonUsuario))
}

func main() {
	// Crear una instancia de Usuario
	usuario := Usuario{
		Nombre:     "Juan",
		Apellido:   "Pérez",
		Edad:       30,
		Correo:     "juan@example.com",
		Contraseña: "secreto123",
	}

	// Imprimir los datos originales del usuario
	fmt.Println("\nDatos originales del usuario")
	convertirJson(usuario)

	// Realizar cambios en los datos del usuario usando las funciones
	var modificableUser Modificable = &usuario
	modificableUser.cambiarNombre("Luis", "Gómez")
	modificableUser.cambiarEdad(25)
	modificableUser.cambiarCorreo("luis@example.com")
	modificableUser.cambiarContraseña("nuevaContraseña")

	// Imprimir los datos actualizados del usuario
	fmt.Println("\nDatos actualizados del usuario:")
	convertirJson(usuario)

}

//*******************************************************************************************************//
