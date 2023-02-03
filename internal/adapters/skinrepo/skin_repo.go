package skinrepo

import (
	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/skinrepo/postgres"
	db "github.com/evmartinelli/go-rifa-microservice/pkg/postgres"
)

func NewPlayerSkinRepo(config *db.Database) *postgres.PlayerSkinRepo {
	return postgres.NewPlayerSkinRepo(config)
}
