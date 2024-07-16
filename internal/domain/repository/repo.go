package repository

import (
	"context"
	core "proj/internal/domain/service"
)

type Repository struct {
	//db
}

func (r *Repository) QueryByID(ctx context.Context, orderID string) (core.Order, error) {
	// fmt.Println("Retrieving order")
	// out, err := s.repo.aaa
	// return out, nil
	return core.Order{}, nil
}
