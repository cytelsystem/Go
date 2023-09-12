package producto

import (
	"context"
	"errors"
	"time"
)

var (
	ErrEmptyList = errors.New("la lista de productos esta vacia")
	ErrNofFound = errors.New("Producto no encontrado")
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
	Delete(ctx context.Context, id int) (error)
}

type repository struct {
	db []Producto
}

func NewRepository() Repository {
	return &repository{
		db: productos,
	}

}

//**********Metodos de la interface**************************************
//Get
func (r *repository) GetAll(ctx context.Context) ([]Producto, error) {
	if len(r.db) < 1 {
		return []Producto{}, ErrEmptyList
	}
	return r.db, nil
}
//Delete
func (r *repository) Delete(ctx context.Context, id int) (error) {
	for key, producto := range r.db {
		if producto.ID == id {
			r.db = append(r.db[:key], r.db[key+1:]...)
			return nil
		}
	}
	return ErrNofFound
}