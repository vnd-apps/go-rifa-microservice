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
		tx.Create(&RaffleNumbers{
			Number: v.Number,
			Slug:   rm.Slug,
			Status: string(v.Status),
		})
	}

	err := tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func (r *RaffleRepo) GetAll(ctx context.Context) ([]raffle.Raffle, error) {
	var models []Raffle
	raffleResults := make([]raffle.Raffle, 0, len(models))

	err := r.db.Where(&Raffle{Status: string(raffle.Open)}).Find(&models).Error
	if err != nil {
		return nil, err
	}

	for i := range models {
		raffleResults = append(raffleResults, ToRaffle(&models[i]))
	}

	return raffleResults, nil
}

func (r *RaffleRepo) GetByID(ctx context.Context, id string) (raffle.Raffle, error) {
	var models Raffle

	err := r.db.Where(&Raffle{Slug: id}).First(&models).Error
	if err != nil {
		return raffle.Raffle{}, err
	}

	return ToRaffle(&models), nil
}

func (r *RaffleRepo) GetProduct(ctx context.Context, id string) (raffle.Raffle, error) {
	var models Raffle

	var itemModels []RaffleNumbers

	tx := r.db.Begin()

	err := tx.Where(&Raffle{Slug: id}).First(&models).Error
	if err != nil {
		return raffle.Raffle{}, err
	}

	result := ToRaffle(&models)

	err = tx.Where(&RaffleNumbers{Slug: id}).Find(&itemModels).Error
	if err != nil {
		return raffle.Raffle{}, err
	}

	for i := range itemModels {
		result.Numbers = append(result.Numbers, ToNumbers(&itemModels[i]))
	}

	return result, nil
}

func (r *RaffleRepo) UpdateItems(ctx context.Context, itens []raffle.Numbers) error {
	tx := r.db.Begin()
	for _, v := range itens {
		if err := tx.Where(&RaffleNumbers{Slug: v.Slug, Number: v.Number}).Updates(&RaffleNumbers{Status: string(raffle.Ordered)}).Error; err != nil {
			return err
		}
	}

	err := tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func (r *RaffleRepo) UpdateItem(ctx context.Context, slug string, numbers int) error {
	err := r.db.Where(&RaffleNumbers{Slug: slug, Number: numbers}).
		Updates(
			&RaffleNumbers{
				Status:         string(raffle.Pending),
				ReservedUserID: "1234",
				ReservedAt:     r.db.NowFunc(),
			}).Error
	if err != nil {
		return err
	}

	return nil
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

func ToNumbers(n *RaffleNumbers) raffle.Numbers {
	return raffle.Numbers{
		Slug:   n.Slug,
		Number: n.Number,
		Status: raffle.ItemStatus(n.Status),
	}
}
