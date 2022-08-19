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
func NewSteam(r PlayerSkinRepo, w SteamWebAPI) *PlayerInventoryUseCase {
	return &PlayerInventoryUseCase{
		repo:   r,
		webAPI: w,
	}
}

func (uc *PlayerInventoryUseCase) GetPlayerInventory(ctx context.Context, id string) (Skin, error) {
	skin, err := uc.webAPI.PlayerItens(id)
	if err != nil {
		return Skin{}, fmt.Errorf("TranslationWebAPI - Translate - trans.Translate: %w", err)
	}

	err = uc.repo.Create(ctx, skin)
	if err != nil {
		return Skin{}, fmt.Errorf("TranslationWebAPI - Translate - trans.Translate: %w", err)
	}

	return skin, nil
}
