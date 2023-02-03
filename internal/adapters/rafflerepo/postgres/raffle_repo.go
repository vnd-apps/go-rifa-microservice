package postgres

import (
	"context"

	"github.com/evmartinelli/go-rifa-microservice/internal/core/raffle"
	postgres "github.com/evmartinelli/go-rifa-microservice/pkg/postgres"
)

type RaffleRepo struct {
	db *postgres.Database
}

func NewRaffleRepo(db *postgres.Database) *RaffleRepo {
	return &RaffleRepo{db}
}

func (r *RaffleRepo) Create(ctx context.Context, rm *raffle.Raffle) error {
	panic("not implemented yet")
}

func (r *RaffleRepo) GetAll(ctx context.Context) ([]raffle.Raffle, error) {
	panic("not implemented yet")
}

func (r *RaffleRepo) GetByID(ctx context.Context, id string) (raffle.Raffle, error) {
	panic("not implemented yet")
}

func (r *RaffleRepo) GetProduct(ctx context.Context, id string) (raffle.Raffle, error) {
	panic("not implemented yet")
}

func (r *RaffleRepo) UpdateItems(ctx context.Context, itens []raffle.Numbers) error {
	panic("not implemented yet")
}
