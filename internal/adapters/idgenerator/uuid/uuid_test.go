package uuid_test

import (
	"testing"

	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/idgenerator/uuid"
	"github.com/evmartinelli/go-rifa-microservice/pkg/assert"
)

const UUIDRegex = `^[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}$`

func TestUUIDGenerator(t *testing.T) {
	t.Parallel()
	t.Run("It generates a v4 uuid", func(t *testing.T) {
		t.Parallel()
		idgen := uuid.NewGenerator()
		assert.Matches(t, UUIDRegex, idgen.Generate())
	})
}
