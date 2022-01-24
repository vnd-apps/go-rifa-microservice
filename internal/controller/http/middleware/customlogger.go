package customlogger

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/evmartinelli/go-rifa-microservice/pkg/logger"
	"github.com/gin-gonic/gin"
)

func RequestLoggerMiddleware(l logger.Interface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		body, _ := ioutil.ReadAll(tee)
		c.Request.Body = ioutil.NopCloser(&buf)
		l.Info(string(body), c.Request.Header)
		c.Next()
	}
}
