// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/evmartinelli/go-rifa-microservice/internal/core/raffle"
)

//go:generate mockgen -source=raffle_interfaces.go -destination=./mocks_raffle_test.go -package=usecase_test

type (
	// Raffle -.
	Raffle interface {
		Create(context.Context, raffle.Raffle) error
		GetAvailableRaffle(context.Context) ([]raffle.Raffle, error)
	}

	// RaffleRepo -.
	RaffleRepo interface {
		Create(context.Context, raffle.Raffle) error
		GetAvailableRaffle(context.Context) ([]raffle.Raffle, error)
	}
)
