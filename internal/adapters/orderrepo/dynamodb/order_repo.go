package dynamodb

import (
	"context"
	"fmt"

	"github.com/evmartinelli/go-rifa-microservice/internal/core/order"
	db "github.com/evmartinelli/go-rifa-microservice/pkg/dynamodb"
)

const (
	PK = "P#%v#U#%v"
	SK = "O#%v"
)

type OrderRepo struct {
	db *db.DynamoConfig
}

func NewOrderRepo(mdb *db.DynamoConfig) *OrderRepo {
	return &OrderRepo{mdb}
}

func (r *OrderRepo) CreateOrder(ctx context.Context, rm *order.Request) (order.Order, error) {
	_, err := r.db.Save(rm)
	if err != nil {
		return order.Order{}, err
	}

	return order.Order{}, nil
}

func (r *OrderRepo) GetUserOrders(ctx context.Context, pid string) ([]order.Order, error) {
	results := []order.Order{}

	err := r.db.GetAll(fmt.Sprintf(PK, pid, pid), &results)
	if err != nil {
		return []order.Order{}, err
	}

	return results, nil
}
