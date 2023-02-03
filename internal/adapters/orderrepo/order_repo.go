package orderrepo

import (
	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/orderrepo/dynamodb"
	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/orderrepo/postgres"
	db "github.com/evmartinelli/go-rifa-microservice/pkg/dynamodb"
	pg "github.com/evmartinelli/go-rifa-microservice/pkg/postgres"
)

func NewDynamoDBOrderRepo(config *db.DynamoConfig) *dynamodb.OrderRepo {
	return dynamodb.NewOrderRepo(config)
}

func NewPostgresOrderRepo(config *pg.Database) *postgres.OrderRepo {
	return postgres.NewOrderRepo(config)
}
