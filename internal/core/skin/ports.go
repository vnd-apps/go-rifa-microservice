// Package usecase implements application business logic. Each logic group in own file.
package skin

import (
	"context"
)

//go:generate mockgen -source=ports.go -destination=./mock_skin_test.go -package=skin_test

type (
	SteamWebAPI interface {
		PlayerItens(string) (Skin, error)
	}
	PlayerSkinRepo interface {
		Create(context.Context, Skin) error
	}
)
