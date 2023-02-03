package postgres

import (
	"context"

	"github.com/evmartinelli/go-rifa-microservice/internal/core/order"
	postgres "github.com/evmartinelli/go-rifa-microservice/pkg/postgres"
)

type OrderRepo struct {
	db *postgres.Database
}

func NewOrderRepo(db *postgres.Database) *OrderRepo {
	return &OrderRepo{db}
}

func (r *OrderRepo) CreateOrder(ctx context.Context, rm *order.Order) (order.Order, error) {
	return order.Order{}, nil
}

func (r *OrderRepo) GetUserOrders(ctx context.Context, pid string) ([]order.Order, error) {
	return nil, nil
}
