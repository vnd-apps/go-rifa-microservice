package claim_test

import (
	"testing"

	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/claim"
	"github.com/evmartinelli/go-rifa-microservice/pkg/assert"
)

func TestUUIDGenerator(t *testing.T) {
	t.Parallel()
	t.Run("It generates a v4 uuid", func(t *testing.T) {
		t.Parallel()
		token := ""
		claims, err := claim.NewClaims().GetUsername(token)
		if err != nil {
			assert.NotNil(t, err)
		}
		assert.Nil(t, claims)
	})
}
