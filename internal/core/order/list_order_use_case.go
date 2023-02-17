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

func (uc *ListOrderUseCase) Run(ctx context.Context, token string) ([]Order, error) {
	order, err := uc.repo.GetUserOrders(ctx, token)
	if err != nil {
		return nil, err
	}

	return order, nil
}
