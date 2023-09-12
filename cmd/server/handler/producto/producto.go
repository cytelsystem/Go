package producto

import (
	"net/http"

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

func (c *Controlador) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productos, err := c.service.GetAll(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Mensaje": "Intercal Server Error",
			})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data" : productos, 
		})

	}

}
