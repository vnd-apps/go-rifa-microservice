package postgres

import (
	"context"

	entity "github.com/evmartinelli/go-rifa-microservice/internal/adapters/rafflerepo/postgres"
	postgres "github.com/evmartinelli/go-rifa-microservice/pkg/postgres"
)

type SlugRepo struct {
	db *postgres.Database
}

func NewSlugRepo(db *postgres.Database) *SlugRepo {
	return &SlugRepo{db}
}

func (r *SlugRepo) CheckExists(ctx context.Context, slug string) (bool, error) {
	var count int64

	if err := r.db.Where(&entity.Raffle{Slug: slug}).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}
