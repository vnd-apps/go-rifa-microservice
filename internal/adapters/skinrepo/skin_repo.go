package skinrepo

import (
	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/skinrepo/dynamodb"
	db "github.com/evmartinelli/go-rifa-microservice/pkg/dynamodb"
)

func NewPlayerSkinRepo(config *db.DynamoConfig) *dynamodb.PlayerSkinRepo {
	return dynamodb.NewPlayerSkinRepo(config)
}
