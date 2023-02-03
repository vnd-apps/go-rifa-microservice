package customlogger

import (
	"bytes"
	"io"

	"github.com/gin-gonic/gin"

	"github.com/evmartinelli/go-rifa-microservice/pkg/logger"
)

func RequestLoggerMiddleware(l *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)

		body, err := io.ReadAll(tee)
		if err != nil {
			c.Next()
		}

		c.Request.Body = io.NopCloser(&buf)
		l.Info(string(body), c.Request.Header)
		c.Next()
	}
}
