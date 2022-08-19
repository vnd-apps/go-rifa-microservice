// Package app configures and runs application.
package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/evmartinelli/go-rifa-microservice/config"
	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/rafflerepo"
	v1 "github.com/evmartinelli/go-rifa-microservice/internal/controller/http/v1"
	"github.com/evmartinelli/go-rifa-microservice/internal/core/raffle"
	"github.com/evmartinelli/go-rifa-microservice/pkg/httpserver"
	"github.com/evmartinelli/go-rifa-microservice/pkg/logger"
	"github.com/evmartinelli/go-rifa-microservice/pkg/mongodb"
)

// Run creates objects via constructors.
func Run() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	l := logger.New(cfg.Log.Level)

	// Mongo Repository
	mdb, err := mongodb.New(cfg.MDB.URL, cfg.MDB.Database)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}

	// Repo
	raffleRepo := rafflerepo.NewDynamoDBRaffleRepo(mdb)

	// UseCases
	generateRaffleUseCase := raffle.NewGenerateRaffleUseCase(raffleRepo)
	listRaffleUseCase := raffle.NewListRaffleUseCase(raffleRepo)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, *generateRaffleUseCase, *listRaffleUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
