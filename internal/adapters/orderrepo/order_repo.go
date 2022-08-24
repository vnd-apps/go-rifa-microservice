package orderrepo

import (
	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/orderrepo/dynamodb"
	db "github.com/evmartinelli/go-rifa-microservice/pkg/dynamodb"
)

func NewDynamoDBOrderRepo(config *db.DynamoConfig) *dynamodb.OrderRepo {
	return dynamodb.NewOrderRepo(config)
}
