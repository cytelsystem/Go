package main

type Producto interface {
	Precio() float64
}

type Pequeno struct{
	Precio float64
}

func (p Pequeno) Precio() float64 {
	return p.Precio
}

type Mediano struct{
	Precio float64
}
type Grande struct{
	Precio float64
}



func main() {



}
