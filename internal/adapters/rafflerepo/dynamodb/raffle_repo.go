package dynamodb

import (
	"context"
	"fmt"
	"strconv"

	"github.com/evmartinelli/go-rifa-microservice/internal/core/raffle"
	db "github.com/evmartinelli/go-rifa-microservice/pkg/dynamodb"
)

const (
	productType = "RAFFLE"
	SK          = "P#%v"
)

type RaffleRepo struct {
	db *db.DynamoConfig
}

func NewRaffleRepo(mdb *db.DynamoConfig) *RaffleRepo {
	return &RaffleRepo{mdb}
}

func RaffleToDynamo(r *raffle.Raffle) DynamoRaffle {
	return DynamoRaffle{
		PK:          productType,
		SK:          fmt.Sprintf(SK, r.ID),
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
		Slug:        r.Slug,
		Status:      string(r.Status),
		ImageURL:    r.ImageURL,
		UnitPrice:   r.UnitPrice,
		Quantity:    r.Quantity,
	}
}

func RaffleItemToDynamoItem(r *raffle.Variation) DynamoRaffleItem {
	return DynamoRaffleItem{
		PK:     r.ID,
		SK:     strconv.Itoa(r.Number),
		ID:     r.ID,
		Number: r.Number,
		Name:   r.Name,
		Status: string(r.Status),
	}
}

func DynamoToRaffle(dyn *DynamoRaffle) raffle.Raffle {
	return raffle.Raffle{
		ID:          dyn.ID,
		Name:        dyn.Name,
		Description: dyn.Description,
		Slug:        dyn.Slug,
		Status:      raffle.Status(dyn.Status),
		ImageURL:    dyn.ImageURL,
		UnitPrice:   dyn.UnitPrice,
		Quantity:    dyn.Quantity,
	}
}

func DynamoItemToRaffleItem(dyn *DynamoRaffleItem) raffle.Variation {
	return raffle.Variation{
		ID:     dyn.ID,
		Number: dyn.Number,
		Name:   dyn.Name,
		Status: raffle.ItemStatus(dyn.Status),
	}
}

func (r *RaffleRepo) Create(ctx context.Context, rm *raffle.Raffle) error {
	_, err := r.db.Save(RaffleToDynamo(rm))
	if err != nil {
		return err
	}

	raffleItems := make([]DynamoRaffleItem, 0, len(rm.Variation))
	for i := range rm.Variation {
		raffleItems = append(raffleItems, RaffleItemToDynamoItem(&rm.Variation[i]))
	}

	err = r.db.SaveMany(raffleItems)
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
