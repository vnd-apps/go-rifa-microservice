// Package usecase implements application business logic. Each logic group in own file.
package order

import (
	"context"
)

//go:generate mockgen -source=ports.go -destination=../mock/order/mock_order_test.go
type Repo interface {
	CreateOrder(ctx context.Context, r *Order) (Order, error)
	GetUserOrders(ctx context.Context, pid string) ([]Order, error)
}

type PixPayment interface {
	GeneratePix() (Pix, error)
}
