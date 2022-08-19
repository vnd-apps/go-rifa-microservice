package skinrepo

import (
	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/skinrepo/dynamodb"
	"github.com/evmartinelli/go-rifa-microservice/pkg/mongodb"
)

func NewPlayerSkinRepo(db *mongodb.MongoDB) *dynamodb.PlayerSkinRepo {
	return dynamodb.NewPlayerSkinRepo(db)
}
