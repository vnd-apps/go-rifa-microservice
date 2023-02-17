package token

import (
	"context"
	"errors"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

type Token struct {
	Username string
	jwt.StandardClaims
}

const (
	BearerLength  = 2
	publicKeysURL = "https://cognito-idp.us-east-1.amazonaws.com/us-east-1_FFPBQIE9C/.well-known/jwks.json"
)

func ExtractToken(bearerToken string) string {
	if len(strings.Split(bearerToken, " ")) == BearerLength {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}

func Valid(bearerToken string) error {
	tokenString := ExtractToken(bearerToken)

	publicKeySet, err := jwk.Fetch(context.TODO(), publicKeysURL)
	if err != nil {
		return err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
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
		return errInvalidToken
	}

	// Check if token is valid
	if !token.Valid {
		return errInvalidToken
	}

	return nil
}

func Claims(bearerToken string) (*Token, error) {
	tokenString := ExtractToken(bearerToken)

	publicKeySet, err := jwk.Fetch(context.TODO(), publicKeysURL)
	if err != nil {
		return nil, err
	}

	cognitoClaims := &Token{}

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

	return cognitoClaims, nil
}

var (
	errSignIn       = errors.New("unexpected signing method")
	errKID          = errors.New("kid header not found")
	errKeyNotFound  = errors.New("key not found")
	errTokenKey     = errors.New("failed to create token key")
	errInvalidToken = errors.New("token is invalid or expired")
)
