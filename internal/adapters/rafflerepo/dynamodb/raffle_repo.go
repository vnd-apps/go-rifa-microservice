package dynamodb

import (
	"context"
	"fmt"
	"strconv"

	"github.com/evmartinelli/go-rifa-microservice/internal/core/raffle"
	db "github.com/evmartinelli/go-rifa-microservice/pkg/dynamodb"
)

const (
	productType     = "RAFFLE"
	productItemType = productType + "ITEM"
	SK              = "P#%v"
	GSI1PK          = "GSI1PK"
	GSI1PKIndex     = GSI1PK + "-index"
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
		SortedNumber: r.PrizeDrawNumber,
		ItemType:     productType,
	}
}

func RaffleItemToDynamoItem(r *raffle.Numbers) DynamoRaffleItem {
	return DynamoRaffleItem{
		PK:       r.ID,
		SK:       strconv.Itoa(r.Number),
		ID:       r.ID,
		Number:   r.Number,
		Name:     r.Slug,
		Status:   string(r.Status),
		ItemType: productItemType,
	}
}

func DynamoToRaffle(dyn *DynamoRecRaffle) raffle.Raffle {
	return raffle.Raffle{
		ID:              dyn.ID,
		Name:            dyn.Name,
		Description:     dyn.Description,
		Slug:            dyn.Slug,
		Status:          raffle.Status(dyn.Status),
		ImageURL:        dyn.ImageURL,
		UnitPrice:       dyn.UnitPrice,
		Quantity:        dyn.Quantity,
		UserLimit:       dyn.UserLimit,
		PrizeDrawNumber: dyn.SortedNumber,
	}
}

func DynamoItemToRaffleItem(dyn *DynamoRecRaffle) raffle.Numbers {
	return raffle.Numbers{
		ID:     dyn.ID,
		Number: dyn.Number,
		Slug:   dyn.Slug,
		Status: raffle.ItemStatus(dyn.Status),
	}
}

func (r *RaffleRepo) Create(ctx context.Context, rm *raffle.Raffle) error {
	_, err := r.db.Save(RaffleToDynamo(rm))
	if err != nil {
		return err
	}

	raffleItems := make([]DynamoRaffleItem, 0, len(rm.Numbers))
	for i := range rm.Numbers {
		raffleItems = append(raffleItems, RaffleItemToDynamoItem(&rm.Numbers[i]))
	}

	err = r.db.SaveMany(raffleItems)
	if err != nil {
		return err
	}

	return nil
}

func (r *RaffleRepo) GetAll(ctx context.Context) ([]raffle.Raffle, error) {
	results := []DynamoRecRaffle{}

	err := r.db.QueryByGSI(productType, GSI1PKIndex, GSI1PK, &results)
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
	result := DynamoRecRaffle{}

	err := r.db.Get(id, fmt.Sprintf(SK, id), &result)
	if err != nil {
		return raffle.Raffle{}, err
	}

	return DynamoToRaffle(&result), nil
}

func (r *RaffleRepo) GetProduct(ctx context.Context, id string) (raffle.Raffle, error) {
	dynamoRaffleRec := []DynamoRecRaffle{}

	err := r.db.Query(id, &dynamoRaffleRec)
	if err != nil {
		return raffle.Raffle{}, err
	}

	return DynamoToProduct(dynamoRaffleRec), nil
}

func (r *RaffleRepo) UpdateItems(ctx context.Context, items []raffle.Numbers) error {
	itemsSK := make([]string, len(items))

	for i := range items {
		dynRaffleItem := RaffleItemToDynamoItem(&items[i])
		itemsSK[i] = dynRaffleItem.SK
	}

	err := r.db.UpdateMany(items[0].ID, string(raffle.Pending), itemsSK)
	if err != nil {
		return err
	}

	return nil
}

func DynamoToProduct(dynamoRaffleRec []DynamoRecRaffle) raffle.Raffle {
	var raffleResult raffle.Raffle

	var raffleItems []raffle.Numbers

	for i := range dynamoRaffleRec {
		if dynamoRaffleRec[i].ItemType == productType {
			raffleResult = DynamoToRaffle(&dynamoRaffleRec[i])
		} else if dynamoRaffleRec[i].ItemType == productItemType {
			raffleItems = append(raffleItems, DynamoItemToRaffleItem(&dynamoRaffleRec[i]))
		}
	}

	raffleResult.Numbers = raffleItems

	return raffleResult
}
