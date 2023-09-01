package producto

import (
	"context"
	"log"
)

type service struct {
	repository Repository
}

type Service interface {
	GetAll(ctx context.Context) ([]Producto, error)
	Delete(ctx context.Context, id int) error
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetAll(ctx context.Context) ([]Producto, error) {
	productos, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("log de error en service de producto", err.Error())
		return []Producto{}, ErrEmptyList
	}

	return productos, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("log de error borrado de producto", err.Error())
		return ErrNotFound
	}

	return nil
}
