package producto

import (
	"context"
	"errors"
	"log"
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
	Patch(ctx context.Context, id int, campos map[string]interface{}) (*Producto, error)
	Create(ctx context.Context, producto *RequestProducto) (*Producto, error)
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

//Post

func (r *repository) Create(ctx context.Context, requestProducto *RequestProducto) (*Producto, error) {
	nuevoProducto := &Producto{
			Name:        requestProducto.Name,
			Quantity:    requestProducto.Quantity,
			CodeValue:   requestProducto.CodeValue,
			Expiration:  requestProducto.Expiration,
			IsPublished: requestProducto.IsPublished,
			Price:       requestProducto.Price,
	}

	// Genera un nuevo ID para el producto (puedes usar alguna lógica para asignar un ID único).
	nuevoID := len(r.db) + 1
	nuevoProducto.ID = nuevoID

	// Agrega el nuevo producto a la lista.
	r.db = append(r.db, *nuevoProducto)

	return nuevoProducto, nil
}



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

// Update
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

// Patch
func (r *repository) Patch(ctx context.Context, id int, campos map[string]interface{}) (*Producto, error) {
	// Obtén el producto por su ID utilizando el método GetByID
	producto, err := r.GetByID(ctx, id)
	if err != nil {
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
			expirationTime, err := time.Parse("2006-01-02", expirationStr)
			if err != nil {
				log.Println("Fech no tiene el formato adecuado")
			}
			producto.Expiration = expirationTime

		case "price":
			producto.Price = valor.(float64)
		}
	}

	// Llama al método correspondiente en el repositorio para aplicar la actualización.
	_, err = r.Update(ctx, id, producto)
	if err != nil {
		return nil, err
	}

	return producto, nil
}
