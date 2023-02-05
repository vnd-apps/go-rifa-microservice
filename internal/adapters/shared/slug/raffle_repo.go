package rafflerepo

import (
	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/shared/slug/postgres"
	pg "github.com/evmartinelli/go-rifa-microservice/pkg/postgres"
)

func NewPostgresRaffleRepo(config *pg.Database) *postgres.SlugRepo {
	return postgres.NewSlugRepo(config)
}
