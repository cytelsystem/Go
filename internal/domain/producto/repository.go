package producto

import (
	"context"
	"errors"
	"time"
)

// Errores
var (
	ErrEmptyList = errors.New("la lista de productos esta vacia")
	ErrNotFound  = errors.New("producto no encontrado")
)

// Base de datos en memoria.
var (
	productos = []Producto{
		{
			ID:          1,
			Name:        "Banana",
			CodeValue:   "AABBCCC",
			Quantity:    10,
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       10.0,
		},
		{
			ID:          2,
			Name:        "Manzana",
			CodeValue:   "AABBDDD",
			Quantity:    5,
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       5.0,
		},
		{
			ID:          3,
			Name:        "Pera",
			CodeValue:   "AAZZZCCC",
			Quantity:    8,
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       8.0,
		},
	}
)

type Repository interface {
	GetAll(ctx context.Context) ([]Producto, error)
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db []Producto
}

func NewRepository() Repository {
	return &repository{
		db: productos,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]Producto, error) {
	if len(r.db) < 1 {
		return []Producto{}, ErrEmptyList
	}

	return r.db, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {

	for key, producto := range r.db {
		if producto.ID == id {
			r.db = append(r.db[:key], r.db[key+1:]...)
			return nil
		}

	}

	return ErrNotFound
}
