package dynamodb

import (
	"context"

	"github.com/evmartinelli/go-rifa-microservice/internal/core/skin"
	db "github.com/evmartinelli/go-rifa-microservice/pkg/dynamodb"
)

type PlayerSkinRepo struct {
	db *db.DynamoConfig
}

func NewPlayerSkinRepo(mdb *db.DynamoConfig) *PlayerSkinRepo {
	return &PlayerSkinRepo{mdb}
}

func (r *PlayerSkinRepo) Create(ctx context.Context, sk skin.Skin) error {
	_, err := r.db.Save(sk)
	if err != nil {
		return err
	}

	return nil
}
