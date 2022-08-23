package rafflerepo

import (
	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/rafflerepo/dynamodb"
	db "github.com/evmartinelli/go-rifa-microservice/pkg/dynamodb"
)

func NewDynamoDBRaffleRepo(db *db.DBConfig) *dynamodb.RaffleRepo {
	return dynamodb.NewRaffleRepo(db)
}
