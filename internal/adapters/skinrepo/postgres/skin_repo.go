package postgres

import (
	"context"

	"github.com/evmartinelli/go-rifa-microservice/internal/core/skin"
	pg "github.com/evmartinelli/go-rifa-microservice/pkg/postgres"
)

type PlayerSkinRepo struct {
	db *pg.Database
}

func NewPlayerSkinRepo(db *pg.Database) *PlayerSkinRepo {
	return &PlayerSkinRepo{db}
}

func (r *PlayerSkinRepo) Create(ctx context.Context, sk skin.Skin) error {
	return nil
}
