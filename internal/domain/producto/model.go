package producto

import "time"

// Producto ...
type Producto struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Quantity    int       `json:"quantity"`
	CodeValue   string    `json:"code_value"`
	Expiration  time.Time `json:"expiration"`
	IsPublished bool      `json:"is_published"`
	Price       float64   `json:"price"`
}
