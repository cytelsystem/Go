package producto

import (
	"net/http"
	"strconv"
	"github.com/cytelsystem/Go/tree/practicaHJM/internal/domain/producto"
	"github.com/gin-gonic/gin"
)

type Controlador struct {
	service producto.Service
}

func NewControladorProducto(service producto.Service) *Controlador {
	return &Controlador{
		service: service,
	}
}

//**************************Metodos***************************//

// Post
func (c *Controlador) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
			// Decodifica los datos del producto desde la solicitud JSON
			var requestProducto producto.RequestProducto
			if err := ctx.BindJSON(&requestProducto); err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{
							"Mensaje": "Datos de creación inválidos",
					})
					return
			}

			// Llama al método correspondiente en el servicio para crear el nuevo producto.
			productoCreado, err := c.service.Create(ctx, &requestProducto)
			if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{
							"Mensaje": "No se pudo crear el producto",
					})
					return
			}

			ctx.JSON(http.StatusCreated, gin.H{
					"data": productoCreado,
			})
	}
}


// Get
func (c *Controlador) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productos, err := c.service.GetAll(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Mensaje": "Intercal Server Error",
			})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": productos,
		})

	}

}

// PorID
func (c *Controlador) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Mensaje": "ID de producto inválido",
			})
			return
		}

		producto, err := c.service.GetByID(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"Mensaje": "Producto no encontrado",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": producto,
		})
	}
}

// Update
func (c *Controlador) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Mensaje": "ID de producto inválido",
			})
			return
		}

		// Decodifica los datos del producto actualizado desde la solicitud JSON
		var productoActualizado producto.Producto
		if err := ctx.BindJSON(&productoActualizado); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Mensaje": "Datos de actualización inválidos",
			})
			return
		}

		// Llama al método correspondiente en el servicio para actualizar el producto
		producto, err := c.service.Update(ctx, id, &productoActualizado)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Mensaje": "No se pudo actualizar el producto",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": producto,
		})
	}
}

// Delete
func (c *Controlador) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Mensaje": "id invalido",
			})
			return
		}

		err = c.service.Delete(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Mensaje": "No se pudo eliminar el producto",
			})
			return
		}
		ctx.JSON(http.StatusNoContent, gin.H{
			"Mensaje": "Producto Eliminado",
		})
	}

}

// Patch
func (c *Controlador) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Mensaje": "ID de producto inválido",
			})
			return
		}

		// Decodifica los campos de actualización desde la solicitud JSON
		var campos map[string]interface{}
		if err := ctx.BindJSON(&campos); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Mensaje": "Datos de actualización inválidos",
			})
			return
		}

		// Llama al método correspondiente en el servicio para aplicar la actualización de campos.
		producto, err := c.service.Patch(ctx, id, campos)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Mensaje": "No se pudo actualizar el producto",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": producto,
		})
	}
}
