package raffle

import (
	"context"
	"fmt"
)

type ChangeRaffleNumberStatusUseCase struct {
	repo Repo
}

func NewChangeRaffleNumberStatusUseCase(r Repo) *ChangeRaffleNumberStatusUseCase {
	return &ChangeRaffleNumberStatusUseCase{
		repo: r,
	}
}

func (uc *ChangeRaffleNumberStatusUseCase) Run(ctx context.Context, slug string, number int) error {
	err := uc.repo.UpdateItem(ctx, slug, number)
	if err != nil {
		return fmt.Errorf("errorchangerafflestatus: %w", err)
	}

	return nil
}
