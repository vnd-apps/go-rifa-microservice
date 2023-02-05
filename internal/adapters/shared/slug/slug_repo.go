package slug

import (
	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/shared/slug/postgres"
	pg "github.com/evmartinelli/go-rifa-microservice/pkg/postgres"
)

func NewPostgresSlugRepo(config *pg.Database) *postgres.SlugRepo {
	return postgres.NewSlugRepo(config)
}
