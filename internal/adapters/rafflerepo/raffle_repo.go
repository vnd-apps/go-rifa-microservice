package rafflerepo

import (
	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/rafflerepo/dynamodb"
	"github.com/evmartinelli/go-rifa-microservice/pkg/mongodb"
)

func NewDynamoDBRaffleRepo(db *mongodb.MongoDB) *dynamodb.RaffleRepo {
	return dynamodb.NewRaffleRepo(db)
}
