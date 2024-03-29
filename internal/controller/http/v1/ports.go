package v1

import (
	"context"

	"github.com/evmartinelli/go-rifa-microservice/internal/core/order"
	"github.com/evmartinelli/go-rifa-microservice/internal/core/raffle"
	"github.com/evmartinelli/go-rifa-microservice/internal/core/skin"
)

type UseCases struct {
	GenerateRaffle           GenerateRaffleUseCase
	ListRaffle               ListRaffleUseCase
	GetRaffle                GetRaffleUseCase
	PlayerInventory          PlayerInventoryUseCase
	PlaceOrder               PlaceOrderUseCase
	ChangeRaffleNumberStatus ChangeRaffleNumberStatusUseCase
	ListOrder                ListUserOrdersUseCase
}

type ListUserOrdersUseCase interface {
	Run(ctx context.Context, token string) ([]order.Order, error)
}
type GenerateRaffleUseCase interface {
	Run(context.Context, *raffle.Raffle) error
}

type ListRaffleUseCase interface {
	Run(ctx context.Context) ([]raffle.Raffle, error)
}

type GetRaffleUseCase interface {
	Run(ctx context.Context, id string) (raffle.Raffle, error)
}

type PlayerInventoryUseCase interface {
	Run(ctx context.Context, id string) (skin.Skin, error)
}

type PlaceOrderUseCase interface {
	Run(ctx context.Context, model *order.Request, token string) (*order.Order, error)
}

type ChangeRaffleNumberStatusUseCase interface {
	Run(ctx context.Context, slug string, number int) error
}
