package token_test

import (
	"testing"

	token "github.com/evmartinelli/go-rifa-microservice/internal/adapters/shared/token"
	"github.com/evmartinelli/go-rifa-microservice/pkg/assert"
)

func TestClaims(t *testing.T) {
	t.Parallel()
	t.Run("It return an username given a jwt token", func(t *testing.T) {
		t.Parallel()
		claims, err := token.NewAuth().Claims("Bearer randomtoken")
		if err != nil {
			assert.NotNil(t, err)
		}
		assert.NotNil(t, claims)
	})
}
