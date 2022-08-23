package raffle

import (
	"context"
	"fmt"
)

// TranslationUseCase -.
type ListRaffleUseCase struct {
	repo Repo
}

// New -.
func NewListRaffleUseCase(r Repo) *ListRaffleUseCase {
	return &ListRaffleUseCase{
		repo: r,
	}
}

func (uc *ListRaffleUseCase) Run(ctx context.Context) ([]Raffle, error) {
	res, err := uc.repo.GetAvailableRaffle(ctx)
	if err != nil {
		return []Raffle{}, fmt.Errorf("TranslationUseCase - Translate - s.repo.Store: %w", err)
	}

	return res, nil
}
