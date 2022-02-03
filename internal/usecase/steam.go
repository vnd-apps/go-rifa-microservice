package usecase

import (
	"context"
	"fmt"

	"github.com/evmartinelli/go-rifa-microservice/internal/entity"
)

// SteamUseCase -.
type SteamUseCase struct {
	webAPI SteamWebAPI
}

// New -.
func NewSteam(w SteamWebAPI) *SteamUseCase {
	return &SteamUseCase{
		webAPI: w,
	}
}

// Player Itens - getting inventory player.
func (uc *SteamUseCase) GetPlayerInventory(ctx context.Context, id string) (entity.Skin, error) {
	skin, err := uc.webAPI.PlayerItens(id)
	if err != nil {
		return entity.Skin{}, fmt.Errorf("TranslationWebAPI - Translate - trans.Translate: %w", err)
	}

	return skin, nil
}
