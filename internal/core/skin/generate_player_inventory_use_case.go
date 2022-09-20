package skin

import (
	"context"
	"fmt"
)

type PlayerInventoryUseCase struct {
	repo   PlayerSkinRepo
	webAPI SteamWebAPI
}

// New -.
func NewPlayerInventoryUseCase(r PlayerSkinRepo, w SteamWebAPI) *PlayerInventoryUseCase {
	return &PlayerInventoryUseCase{
		repo:   r,
		webAPI: w,
	}
}

func (uc *PlayerInventoryUseCase) Run(ctx context.Context, id string) (Skin, error) {
	skin, err := uc.webAPI.PlayerItens(id)
	if err != nil {
		return Skin{}, fmt.Errorf("SteamAPI - Error Get Player Items: %w", err)
	}

	err = uc.repo.Create(ctx, skin)
	if err != nil {
		return Skin{}, fmt.Errorf("SteamRepo - Error Creating Itens: %w", err)
	}

	return skin, nil
}
