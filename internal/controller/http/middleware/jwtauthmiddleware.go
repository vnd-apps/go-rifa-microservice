package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	token "github.com/evmartinelli/go-rifa-microservice/internal/adapters/shared/token"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		err := token.NewAuth().CheckIsValid(authHeader)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()

			return
		}

		c.Next()
	}
}
