package producto

import (
	"context"
	"log"
	"sync"
	"time"
)

type service struct {
	repository Repository
	mutex      sync.Mutex
}
type Service interface {
	GetAll(ctx context.Context) ([]Producto, error)
	GetByID(ctx context.Context, id int) (*Producto, error)
	Update(ctx context.Context, id int, producto *Producto) (*Producto, error)
	Patch(ctx context.Context, id int, campos map[string]interface{}) (*Producto, error)
	Delete(ctx context.Context, id int) error
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

// ***************Metodos interface service***************************//
// Get
func (s *service) GetAll(ctx context.Context) ([]Producto, error) {
	productos, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("log de error en service de producto", err.Error())
		return []Producto{}, ErrEmptyList
	}

	return productos, nil
}

// GetByID
func (s *service) GetByID(ctx context.Context, id int) (*Producto, error) {
	// Bloquear el acceso concurrente al servicio.
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Llama al método correspondiente en el repositorio para obtener el producto por su ID.
	producto, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("Error al obtener producto en el servicio:", err.Error())
		return nil, err
	}

	return producto, nil
}

// Update
func (s *service) Update(ctx context.Context, id int, producto *Producto) (*Producto, error) {
	// Bloquear el acceso concurrente al servicio.
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Llama al método correspondiente en el repositorio para actualizar el producto por su ID.
	productoActualizado, err := s.repository.Update(ctx, id, producto)
	if err != nil {
		log.Println("Error al actualizar producto en el servicio:", err.Error())
		return nil, err
	}

	return productoActualizado, nil
}

// Delete
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("log de error borrado de producto", err.Error())
		return ErrNofFound
	}
	return nil
}

// Patch
func (s *service) Patch(ctx context.Context, id int, campos map[string]interface{}) (*Producto, error) {
	// Obtén el producto por su ID utilizando el método GetByID
	producto, err := s.GetByID(ctx, id)
	if err != nil {
		log.Println("Error al obtener producto en el servicio:", err.Error())
		return nil, err
	}

	// Actualiza los campos del producto con los valores proporcionados en el mapa "campos"
	for campo, valor := range campos {
		switch campo {
		case "name":
			producto.Name = valor.(string)
		case "quantity":
			producto.Quantity = int(valor.(float64))
		case "code_value":
			producto.CodeValue = valor.(string)
		case "is_published":
			producto.IsPublished = valor.(bool)
		case "expiration":
			expirationStr := valor.(string)
			expirationTime, err := time.Parse("2006-01-02 00:00:00", expirationStr)
			if err != nil {
				log.Println("Fech no tiene el formato adecuado")
			}
			producto.Expiration = expirationTime
		case "price":
			producto.Price = valor.(float64)
		}
	}

	// Llama al método correspondiente en el servicio para actualizar el producto con los nuevos datos.
	productoActualizado, err := s.Update(ctx, id, producto)
	if err != nil {
		log.Println("Error al actualizar producto en el servicio:", err.Error())
		return nil, err
	}

	return productoActualizado, nil
}
