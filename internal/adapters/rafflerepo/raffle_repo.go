package rafflerepo

import (
	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/rafflerepo/dynamodb"
	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/rafflerepo/postgres"
	db "github.com/evmartinelli/go-rifa-microservice/pkg/dynamodb"
	pg "github.com/evmartinelli/go-rifa-microservice/pkg/postgres"
)

func NewDynamoDBRaffleRepo(config *db.DynamoConfig) *dynamodb.RaffleRepo {
	return dynamodb.NewRaffleRepo(config)
}

func NewPostgresRaffleRepo(config *pg.Database) *postgres.RaffleRepo {
	return postgres.NewRaffleRepo(config)
}
