package dynamodb

import (
	"context"

	"github.com/evmartinelli/go-rifa-microservice/internal/core/raffle"
	db "github.com/evmartinelli/go-rifa-microservice/pkg/dynamodb"
	"github.com/google/uuid"
)

type RaffleRepo struct {
	db *db.DBConfig
}

func NewRaffleRepo(mdb *db.DBConfig) *RaffleRepo {
	return &RaffleRepo{mdb}
}

func (r *RaffleRepo) Create(ctx context.Context, rm raffle.Raffle) error {
	rm.ID = uuid.New().String()

	_, err := r.db.Save(rm)
	if err != nil {
		return err
	}

	return nil
}

func (r *RaffleRepo) GetAvailableRaffle(ctx context.Context) ([]raffle.Raffle, error) {
	results := []raffle.Raffle{}

	err := r.db.FindByGsi("", "user-index", "user", &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}
