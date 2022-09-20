package raffle

import (
	"context"
	"fmt"
)

type GetRaffleUseCase struct {
	repo Repo
}

func NewGetRaffleUseCase(r Repo) *GetRaffleUseCase {
	return &GetRaffleUseCase{
		repo: r,
	}
}

func (uc *GetRaffleUseCase) Run(ctx context.Context, id string) (Raffle, error) {
	res, err := uc.repo.GetProduct(ctx, id)
	if err != nil {
		return Raffle{}, fmt.Errorf("errorgetbyid: %w", err)
	}

	return res, nil
}
