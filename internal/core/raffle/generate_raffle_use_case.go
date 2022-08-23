package raffle

import (
	"context"
)

type GenerateRaffleUseCase struct {
	repo Repo
}

func NewGenerateRaffleUseCase(r Repo) *GenerateRaffleUseCase {
	return &GenerateRaffleUseCase{
		repo: r,
	}
}

func (uc *GenerateRaffleUseCase) Run(ctx context.Context, model Raffle) error {
	return uc.repo.Create(ctx, model)
}
