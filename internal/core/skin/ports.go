// Package usecase implements application business logic. Each logic group in own file.
package skin

import (
	"context"
)

//go:generate mockgen -source=ports.go -destination=../mock/skin/mock_skin.go

type (
	SteamWebAPI interface {
		PlayerItens(string) (Skin, error)
	}
	PlayerSkinRepo interface {
		Create(context.Context, Skin) error
	}
)
