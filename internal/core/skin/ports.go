// Package usecase implements application business logic. Each logic group in own file.
package skin

import (
	"context"
)

//go:generate mockgen -source=steam_interfaces.go -destination=./mocks_steam_test.go -package=usecase_test

type (
	SteamWebAPI interface {
		PlayerItens(string) (Skin, error)
	}
	PlayerSkinRepo interface {
		Create(context.Context, Skin) error
	}
)
