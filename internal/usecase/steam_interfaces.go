// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/evmartinelli/go-rifa-microservice/internal/entity"
)

//go:generate mockgen -source=steam_interfaces.go -destination=./mocks_steam_test.go -package=usecase_test

type (
	// Steam -.
	Steam interface {
		GetPlayerInventory(context.Context, string) (entity.Skin, error)
	}
	// SteamWebApi -.
	SteamWebAPI interface {
		PlayerItens(string) (entity.Skin, error)
	}
)
