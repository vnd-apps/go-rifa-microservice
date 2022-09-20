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

func (r *OrderRepo) CreateOrder(ctx context.Context, rm *order.PlaceOrderRequest) (order.PlaceOrderResponse, error) {
	_, err := r.db.Save(rm)
	if err != nil {
		return order.PlaceOrderResponse{}, err
	}

	return order.PlaceOrderResponse{}, nil
}

func (r *OrderRepo) GetOrder(ctx context.Context) (order.PlaceOrderResponse, error) {
	results := order.PlaceOrderResponse{}

	err := r.db.Get("id", "id", &results)
	if err != nil {
		return order.PlaceOrderResponse{}, err
	}

	return results, nil
}
