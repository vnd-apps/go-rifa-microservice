package dynamodb

import (
	"context"

	"github.com/google/uuid"

	"github.com/evmartinelli/go-rifa-microservice/internal/core/raffle"
	db "github.com/evmartinelli/go-rifa-microservice/pkg/dynamodb"
)

type RaffleRepo struct {
	db *db.DynamoConfig
}

func NewRaffleRepo(mdb *db.DynamoConfig) *RaffleRepo {
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

	err := r.db.FindAll(&results)
	if err != nil {
		return nil, err
	}

	return results, nil
}
