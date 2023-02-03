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
	tx := r.db.Begin()
	tx.Create(&Raffle{
		Name:            rm.Name,
		Description:     rm.Description,
		Slug:            rm.Slug,
		Status:          string(rm.Status),
		ImageURL:        rm.ImageURL,
		UnitPrice:       rm.UnitPrice,
		UserLimit:       rm.UserLimit,
		Quantity:        rm.Quantity,
		PrizeDrawNumber: rm.PrizeDrawNumber,
	})

	for _, v := range rm.Numbers {
		r.db.Create(&RaffleNumbers{
			Number: v.Number,
			Slug:   rm.Slug,
			Status: string(v.Status),
		})
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (r *RaffleRepo) GetAll(ctx context.Context) ([]raffle.Raffle, error) {
	var models []Raffle
	raffleResults := make([]raffle.Raffle, 0, len(models))

	if err := r.db.Where(&Raffle{Status: string(raffle.Open)}).Find(&models).Error; err != nil {
		return nil, err
	}

	for i := range models {
		raffleResults = append(raffleResults, ToRaffle(&models[i]))
	}

	return raffleResults, nil
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

func ToRaffle(r *Raffle) raffle.Raffle {
	return raffle.Raffle{
		Name:            r.Name,
		Description:     r.Description,
		Slug:            r.Slug,
		Status:          raffle.Status(r.Status),
		ImageURL:        r.ImageURL,
		UnitPrice:       r.UnitPrice,
		Quantity:        r.Quantity,
		UserLimit:       r.UserLimit,
		PrizeDrawNumber: r.PrizeDrawNumber,
	}
}
