// Package usecase implements application business logic. Each logic group in own file.
package order

import (
	"context"
)

//go:generate mockgen -source=ports.go -destination=../mock/order/mock_order.go
type Repo interface {
	CreateOrder(ctx context.Context, r *Order) error
	GetUserOrders(ctx context.Context, userID string) ([]Order, error)
}

type PixPayment interface {
	GeneratePix() (Pix, error)
}
