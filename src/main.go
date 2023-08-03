//***************************************interfaces******************************************//

package main

import (
    "fmt"
)

// Definimos una estructura para representar a un usuario
type Usuario struct {
    ID     int
    Nombre string
    Email  string
}

// Definimos una interfaz para el CRUD de usuarios
type CRUDUsuarios interface {
    CrearUsuario(usuario Usuario)
    LeerUsuario(id int) Usuario
    ActualizarUsuario(id int, usuario Usuario)
    EliminarUsuario(id int)
}

// Implementamos la interfaz CRUDUsuarios para una lista de usuarios en memoria
type ListaUsuarios struct {
    usuarios []Usuario
}

// Método para crear un nuevo usuario y agregarlo a la lista
func (lu *ListaUsuarios) CrearUsuario(usuario Usuario) {
    lu.usuarios = append(lu.usuarios, usuario)
}

// Método para leer un usuario específico por su ID
func (lu *ListaUsuarios) LeerUsuario(id int) Usuario {
    for _, usuario := range lu.usuarios {
        if usuario.ID == id {
            return usuario
        }
    }
    return Usuario{}
}

// Método para actualizar la información de un usuario existente
func (lu *ListaUsuarios) ActualizarUsuario(id int, usuario Usuario) {
    for i, u := range lu.usuarios {
        if u.ID == id {
            lu.usuarios[i] = usuario
            break
        }
    }
}

// Método para eliminar un usuario de la lista por su ID
func (lu *ListaUsuarios) EliminarUsuario(id int) {
    for i, usuario := range lu.usuarios {
        if usuario.ID == id {
            // Eliminamos el usuario de la lista utilizando "append" y "copy"
            lu.usuarios = append(lu.usuarios[:i], lu.usuarios[i+1:]...)
            break
        }
    }
}

func main() {
    // Creamos una nueva lista de usuarios
    listaUsuarios := ListaUsuarios{}

    // Creamos usuarios de ejemplo y los agregamos a la lista
    usuario1 := Usuario{ID: 1, Nombre: "Juan", Email: "juan@example.com"}
    usuario2 := Usuario{ID: 2, Nombre: "María", Email: "maria@example.com"}
    listaUsuarios.CrearUsuario(usuario1)
    listaUsuarios.CrearUsuario(usuario2)

    // Mostramos la lista de usuarios antes de hacer cambios
    fmt.Println("Lista de usuarios:")
    for _, usuario := range listaUsuarios.usuarios {
        fmt.Printf("ID: %d, Nombre: %s, Email: %s\n", usuario.ID, usuario.Nombre, usuario.Email)
    }

    // Actualizamos el nombre del usuario con ID 1
    usuarioActualizado := Usuario{ID: 1, Nombre: "Juan Carlos", Email: "juancarlos@example.com"}
    listaUsuarios.ActualizarUsuario(1, usuarioActualizado)

    // Mostramos la lista de usuarios después de la actualización
    fmt.Println("\nLista de usuarios después de la actualización:")
    for _, usuario := range listaUsuarios.usuarios {
        fmt.Printf("ID: %d, Nombre: %s, Email: %s\n", usuario.ID, usuario.Nombre, usuario.Email)
    }

    // Eliminamos el usuario con ID 2
    listaUsuarios.EliminarUsuario(2)

    // Mostramos la lista de usuarios después de la eliminación
    fmt.Println("\nLista de usuarios después de la eliminación:")
    for _, usuario := range listaUsuarios.usuarios {
        fmt.Printf("ID: %d, Nombre: %s, Email: %s\n", usuario.ID, usuario.Nombre, usuario.Email)
    }
}


//*******************************************************************************************************//
