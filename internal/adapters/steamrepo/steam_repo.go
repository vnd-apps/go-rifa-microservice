package steamrepo

import (
	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/steamrepo/dynamodb"
	"github.com/evmartinelli/go-rifa-microservice/pkg/mongodb"
)

func NewSteamRepo(db *mongodb.MongoDB) *dynamodb.PlayerSkinRepo {
	return dynamodb.NewSteamRepo(db)
}
