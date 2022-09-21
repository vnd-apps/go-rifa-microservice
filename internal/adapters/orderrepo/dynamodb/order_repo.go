package dynamodb

import (
	"context"

	"github.com/evmartinelli/go-rifa-microservice/internal/core/order"
	db "github.com/evmartinelli/go-rifa-microservice/pkg/dynamodb"
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

func (r *OrderRepo) GetOrder(ctx context.Context) (order.Order, error) {
	results := order.Order{}

	err := r.db.Get("id", "id", &results)
	if err != nil {
		return order.Order{}, err
	}

	return results, nil
}
