package v1

import (
	"context"

	"github.com/evmartinelli/go-rifa-microservice/internal/core/raffle"
	"github.com/evmartinelli/go-rifa-microservice/internal/core/skin"
)

type UseCases struct {
	GenerateRaffle  GenerateRaffleUseCase
	ListRaffle      ListRaffleUseCase
	PlayerInventory PlayerInventoryUseCase
}

type GenerateRaffleUseCase interface {
	Run(context.Context, raffle.Raffle) error
}

type ListRaffleUseCase interface {
	Run(ctx context.Context) ([]raffle.Raffle, error)
}

type PlayerInventoryUseCase interface {
	Run(ctx context.Context, id string) (skin.Skin, error)
}
