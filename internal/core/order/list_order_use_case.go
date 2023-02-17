package order

import (
	"context"

	auth "github.com/evmartinelli/go-rifa-microservice/internal/adapters/shared/token"
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
	claims, err := auth.Claims(token)
	if err != nil {
		return nil, err
	}

	order, err := uc.repo.GetUserOrders(ctx, claims.Username)
	if err != nil {
		return nil, err
	}

	return order, nil
}
