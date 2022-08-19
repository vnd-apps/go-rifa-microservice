// Package usecase implements application business logic. Each logic group in own file.
package raffle

import (
	"context"
)

//go:generate mockgen -source=raffle_interfaces.go -destination=./mocks_raffle_test.go -package=usecase_test

type RaffleRepo interface {
	Create(context.Context, Raffle) error
	GetAvailableRaffle(context.Context) ([]Raffle, error)
}
