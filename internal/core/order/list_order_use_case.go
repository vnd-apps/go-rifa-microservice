package order

import (
	"context"

	"github.com/evmartinelli/go-rifa-microservice/internal/core/shared"
)

type ListOrderUseCase struct {
	repo Repo
	user shared.Auth
}

func NewListOrderUseCase(r Repo, s shared.Auth) *ListOrderUseCase {
	return &ListOrderUseCase{
		repo: r,
		user: s,
	}
}

func (uc *ListOrderUseCase) Run(ctx context.Context, token string) ([]Order, error) {
	claims, err := uc.user.Claims(token)
	if err != nil {
		return nil, err
	}

	order, err := uc.repo.GetUserOrders(ctx, claims.Username)
	if err != nil {
		return nil, err
	}

	return order, nil
}
