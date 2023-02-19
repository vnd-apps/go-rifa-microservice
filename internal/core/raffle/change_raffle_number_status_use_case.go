package raffle

import (
	"context"
	"fmt"

	"github.com/evmartinelli/go-rifa-microservice/internal/core/shared"
)

type ChangeRaffleNumberStatusUseCase struct {
	repo Repo
	user shared.Auth
}

func NewChangeRaffleNumberStatusUseCase(r Repo, s shared.Auth) *ChangeRaffleNumberStatusUseCase {
	return &ChangeRaffleNumberStatusUseCase{
		repo: r,
		user: s,
	}
}

func (uc *ChangeRaffleNumberStatusUseCase) Run(ctx context.Context, slug string, number int) error {
	err := uc.repo.UpdateItem(ctx, slug, number)
	if err != nil {
		return fmt.Errorf("errorchangerafflestatus: %w", err)
	}

	return nil
}
