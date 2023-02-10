package claim

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

type CognitoClaims struct {
	Username string
	jwt.StandardClaims
}

func NewClaims() *CognitoClaims {
	return &CognitoClaims{}
}

func (c *CognitoClaims) GetUsername(tokenString string) (*string, error) {
	publicKeysURL := "https://cognito-idp.us-east-1.amazonaws.com/us-east-1_FFPBQIE9C/.well-known/jwks.json"

	publicKeySet, err := jwk.Fetch(context.TODO(), publicKeysURL)
	if err != nil {
		return nil, err
	}

	cognitoClaims := &CognitoClaims{}

	token, err := jwt.ParseWithClaims(tokenString, cognitoClaims, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodRSA)

		if !ok {
			return nil, errSignIn
		}

		// Get "kid" value from token header
		// "kid" is shorthand for Key ID
		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, errKID
		}

		// "kid" must be present in the public keys set
		keys, ok := publicKeySet.LookupKeyID(kid)
		if !ok {
			return nil, errKeyNotFound
		}

		// In our case, we are returning only one key = keys[0]
		// Return token key as []byte{string} type
		var tokenKey interface{}
		if err = keys.Raw(&tokenKey); err != nil {
			return nil, errTokenKey
		}

		return tokenKey, nil
	})
	if err != nil {
		// This place can throw expiration error
		return nil, errInvalidToken
	}

	// Check if token is valid
	if !token.Valid {
		return nil, errInvalidToken
	}

	return &cognitoClaims.Username, nil
}

var (
	errSignIn       = errors.New("unexpected signing method")
	errKID          = errors.New("kid header not found")
	errKeyNotFound  = errors.New("key not found")
	errTokenKey     = errors.New("failed to create token key")
	errInvalidToken = errors.New("token is invalid or expired")
)
