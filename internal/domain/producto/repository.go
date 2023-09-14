package producto

import (
	"context"
	"errors"
	"sync"
	"time"
)

var (
	ErrEmptyList = errors.New("la lista de productos esta vacia")
	ErrNofFound  = errors.New("Producto no encontrado")
)

var (
	productos = []Producto{
		{
			ID:          1,
			Name:        "Banana",
			Quantity:    10,
			CodeValue:   "AABBCCC",
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       10.0,
		},
		{
			ID:          2,
			Name:        "Manzana",
			Quantity:    5,
			CodeValue:   "AABBDDD",
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       5.0,
		},
		{
			ID:          3,
			Name:        "Pera",
			Quantity:    8,
			CodeValue:   "AAAZZZCCC",
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       8.0,
		},
	}
)

type Repository interface {
	GetAll(ctx context.Context) ([]Producto, error)
	GetByID(ctx context.Context, id int) (*Producto, error)
	Update(ctx context.Context, id int, producto *Producto) (*Producto, error)
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db    []Producto
	mutex sync.Mutex
}

func NewRepository() Repository {
	return &repository{
		db: productos,
	}

}

// **********Metodos de la interface**************************************
// Get
func (r *repository) GetAll(ctx context.Context) ([]Producto, error) {
	if len(r.db) < 1 {
		return []Producto{}, ErrEmptyList
	}
	return r.db, nil
}

// GetByID
func (r *repository) GetByID(ctx context.Context, id int) (*Producto, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Buscar el producto por su ID
	for _, p := range r.db {
		if p.ID == id {
			return &p, nil
		}
	}

	// Si no se encuentra el producto, devuelve un error
	return nil, ErrNofFound
}


//Update
func (r *repository) Update(ctx context.Context, id int, producto *Producto) (*Producto, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Buscar el producto por su ID
	for key, p := range r.db {
		if p.ID == id {
			// Actualizar los campos del producto
			r.db[key] = *producto
			return producto, nil
		}
	}

	// Si no se encuentra el producto, devuelve un error
	return nil, ErrNofFound
}

// Delete
func (r *repository) Delete(ctx context.Context, id int) error {
	for key, producto := range r.db {
		if producto.ID == id {
			r.db = append(r.db[:key], r.db[key+1:]...)
			return nil
		}
	}
	return ErrNofFound
}
