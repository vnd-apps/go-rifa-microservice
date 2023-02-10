package claim_test

import (
	"testing"

	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/shared/claim"
	"github.com/evmartinelli/go-rifa-microservice/pkg/assert"
)

func TestGetUserNameToken(t *testing.T) {
	t.Parallel()
	t.Run("It return an username given a jwt token", func(t *testing.T) {
		t.Parallel()
		token := ""
		claims, err := claim.NewClaims().GetUsername(token)
		if err != nil {
			assert.NotNil(t, err)
		}
		assert.NotNil(t, claims)
	})
}
