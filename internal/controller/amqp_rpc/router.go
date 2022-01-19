package amqprpc

import (
	"github.com/evmartinelli/go-rifa-microservice/internal/usecase"
	"github.com/evmartinelli/go-rifa-microservice/pkg/rabbitmq/rmq_rpc/server"
)

// NewRouter -.
func NewRouter(t usecase.Translation) map[string]server.CallHandler {
	routes := make(map[string]server.CallHandler)
	{
		newTranslationRoutes(routes, t)
	}

	return routes
}
