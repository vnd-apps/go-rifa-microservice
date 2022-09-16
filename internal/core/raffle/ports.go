// Package usecase implements application business logic. Each logic group in own file.
package raffle

import (
	"context"
)

//go:generate mockgen -source=ports.go -destination=./mock_raffle_repo_test.go -package=raffle_test

type Repo interface {
	Create(context.Context, *Raffle) error
	GetAvailableRaffle(context.Context) ([]Raffle, error)
}
