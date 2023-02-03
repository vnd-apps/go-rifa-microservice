// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs.
	_ "github.com/evmartinelli/go-rifa-microservice/docs"
	customlogger "github.com/evmartinelli/go-rifa-microservice/internal/controller/http/middleware"
	"github.com/evmartinelli/go-rifa-microservice/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @Title       Go Rifa Microservice
// @description Microservice for Rifa
// @Version     1.0
// @Host        localhost:8080
// @BasePath    /v1
// End.
func NewRouter(handler *gin.Engine, l *logger.Logger, u *UseCases) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	handler.Use(customlogger.RequestLoggerMiddleware(l))

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	h := handler.Group("/v1")
	{
		newRaffleRoutes(h, l, u)
		newOrderRoutes(h, l, u)
		newSteamRoutes(h, l, u)
	}
}
