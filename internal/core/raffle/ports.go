// Package usecase implements application business logic. Each logic group in own file.
package raffle

import (
	"context"
)

//go:generate mockgen -source=ports.go -destination=../mock/raffle/mock_raffle.go

type Repo interface {
	Create(ctx context.Context, raffle *Raffle) error
	GetAll(ctx context.Context) ([]Raffle, error)
	GetByID(ctx context.Context, id string) (Raffle, error)
	GetProduct(ctx context.Context, id string) (Raffle, error)
	UpdateItems(ctx context.Context, itens []Numbers) error
	UpdateItem(ctx context.Context, slug string, numbers int) error
}
