package producto

import (
	"context"
	"log"
	"sync"
)

type service struct {
	repository Repository
	mutex      sync.Mutex
}
type Service interface {
	GetAll(ctx context.Context) ([]Producto, error)
	GetByID(ctx context.Context, id int) (*Producto, error)
	Update(ctx context.Context, id int, producto *Producto) (*Producto, error)
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
