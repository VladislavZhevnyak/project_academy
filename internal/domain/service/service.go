package service

import (
	"context"
	"fmt"
)

type Service struct {
	repo Storer
}

type Order struct {
	ID   string
	Name string
}

type Storer interface {
	QueryByID(context.Context, string) (Order, error)
}

func NewService(store Storer) *Service {
	return &Service{
		repo: store,
	}
}

func (s *Service) ByID(ctx context.Context, orderID string) (Order, error) {
	fmt.Println("Retrieving order")
	out, err := s.repo.QueryByID(ctx, orderID)
	if err != nil {
		return Order{}, fmt.Errorf("service error: %w", &err)
	}
	return out, nil
}
