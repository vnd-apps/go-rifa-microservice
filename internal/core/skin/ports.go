// Package usecase implements application business logic. Each logic group in own file.
package skin

import (
	"context"
)

//go:generate mockgen -source=steam_interfaces.go -destination=./mocks_steam_test.go -package=usecase_test

type (
	// Steam -.
	Steam interface {
		GetPlayerInventory(context.Context, string) (Steam, error)
	}
	// SteamWebApi -.
	SteamWebAPI interface {
		PlayerItens(string) (Steam, error)
	}
	// RaffleRepo -.
	PlayerSkinRepo interface {
		Create(context.Context, Steam) error
	}
)
