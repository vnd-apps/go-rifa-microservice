package rafflerepo

import (
	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/rafflerepo/dynamodb"
	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/rafflerepo/postgres"
	repo "github.com/evmartinelli/go-rifa-microservice/internal/core/raffle"
	db "github.com/evmartinelli/go-rifa-microservice/pkg/dynamodb"
	pg "github.com/evmartinelli/go-rifa-microservice/pkg/postgres"
)

func NewDynamoDBRaffleRepo(config *db.DynamoConfig) repo.Repo {
	return dynamodb.NewRaffleRepo(config)
}

func NewPostgresRaffleRepo(config *pg.Database) repo.Repo {
	return postgres.NewRaffleRepo(config)
}
