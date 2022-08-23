package idgenerator

import (
	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/idgenerator/slug"
	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/idgenerator/uuid"
)

func NewUUIDGenerator() *uuid.Generator {
	return uuid.NewGenerator()
}

func NewSlugGenerator() *slug.Generator {
	return slug.NewGenerator()
}
