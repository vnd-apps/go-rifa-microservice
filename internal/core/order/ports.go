// Package usecase implements application business logic. Each logic group in own file.
package order

import (
	"context"
)

//go:generate mockgen -source=raffle_interfaces.go -destination=./mocks_raffle_test.go -package=usecase_test

type Repo interface {
	CreateOrder(ctx context.Context, rm *Request) (Order, error)
	GetOrder(ctx context.Context) (Order, error)
}

type PixPayment interface {
	GeneratePix() (Pix, error)
}
