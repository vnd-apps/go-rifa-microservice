package order

import (
	"context"
)

type ListOrderUseCase struct {
	repo Repo
}

func NewListOrderUseCase(r Repo) *ListOrderUseCase {
	return &ListOrderUseCase{
		repo: r,
	}
}

func (uc *ListOrderUseCase) Run(ctx context.Context) ([]Order, error) {
	return nil, nil
}
