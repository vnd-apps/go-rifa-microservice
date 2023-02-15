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
	order, err := uc.repo.GetUserOrders(ctx, "123")
	if err != nil {
		return nil, err
	}

	return order, nil
}
