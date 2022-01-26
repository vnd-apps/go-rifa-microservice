// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/evmartinelli/go-rifa-microservice/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Translation -.
	Raffle interface {
		Create(context.Context, entity.Raffle) error
		GetAvaliableRaffle(context.Context) ([]entity.Raffle, error)
	}

	// TranslationRepo -.
	RaffleRepo interface {
		Create(context.Context, entity.Raffle) error
		GetAvaliableRaffle(context.Context) ([]entity.Raffle, error)
	}
)
