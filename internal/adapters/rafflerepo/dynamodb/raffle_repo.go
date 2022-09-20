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
	GSI1PK      = "GSI1PK"
	GSI1PKIndex = GSI1PK + "-index"
)

type RaffleRepo struct {
	db *db.DynamoConfig
}

func NewRaffleRepo(mdb *db.DynamoConfig) *RaffleRepo {
	return &RaffleRepo{mdb}
}

func RaffleToDynamo(r *raffle.Raffle) DynamoRaffle {
	return DynamoRaffle{
		PK:           r.ID,
		SK:           fmt.Sprintf(SK, r.ID),
		GSI1PK:       productType,
		ID:           r.ID,
		Name:         r.Name,
		Description:  r.Description,
		Slug:         r.Slug,
		Status:       string(r.Status),
		ImageURL:     r.ImageURL,
		UnitPrice:    r.UnitPrice,
		Quantity:     r.Quantity,
		UserLimit:    r.UserLimit,
		SortedNumber: r.SortedNumber,
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
		ID:           dyn.ID,
		Name:         dyn.Name,
		Description:  dyn.Description,
		Slug:         dyn.Slug,
		Status:       raffle.Status(dyn.Status),
		ImageURL:     dyn.ImageURL,
		UnitPrice:    dyn.UnitPrice,
		Quantity:     dyn.Quantity,
		UserLimit:    dyn.UserLimit,
		SortedNumber: dyn.SortedNumber,
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

func (r *RaffleRepo) GetAll(ctx context.Context) ([]raffle.Raffle, error) {
	results := []DynamoRaffle{}

	err := r.db.FindByGsi(productType, GSI1PKIndex, GSI1PK, &results)
	if err != nil {
		return nil, err
	}

	raffleResults := make([]raffle.Raffle, 0, len(results))

	for i := range results {
		raffleResults = append(raffleResults, DynamoToRaffle(&results[i]))
	}

	return raffleResults, nil
}

func (r *RaffleRepo) GetByID(ctx context.Context, id string) (raffle.Raffle, error) {
	result := DynamoRaffle{}

	err := r.db.Get(id, "P#"+id, &result)
	if err != nil {
		return raffle.Raffle{}, err
	}

	return DynamoToRaffle(&result), nil
}

func (r *RaffleRepo) GetProduct(ctx context.Context, id string) (raffle.Raffle, error) {
	raffleDynamoResult := DynamoRaffle{}
	raffleDynamoItemResult := []DynamoRaffleItem{}
	raffleResult := raffle.Raffle{}

	err := r.db.GetProduct(id, &raffleDynamoResult, &raffleDynamoItemResult)
	if err != nil {
		return raffleResult, err
	}

	raffleResult = DynamoToRaffle(&raffleDynamoResult)
	for i := range raffleDynamoItemResult {
		raffleResult.Variation = append(raffleResult.Variation, DynamoItemToRaffleItem(&raffleDynamoItemResult[i]))
	}

	return raffleResult, nil
}
