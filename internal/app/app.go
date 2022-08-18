// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/evmartinelli/go-rifa-microservice/config"
	v1 "github.com/evmartinelli/go-rifa-microservice/internal/controller/http/v1"
	"github.com/evmartinelli/go-rifa-microservice/internal/usecase"
	"github.com/evmartinelli/go-rifa-microservice/internal/usecase/repo/mongodbrepo"
	"github.com/evmartinelli/go-rifa-microservice/internal/usecase/webapi"
	"github.com/evmartinelli/go-rifa-microservice/pkg/httpserver"
	"github.com/evmartinelli/go-rifa-microservice/pkg/logger"
	"github.com/evmartinelli/go-rifa-microservice/pkg/mongodb"
	"github.com/evmartinelli/go-rifa-microservice/pkg/postgres"
	"github.com/evmartinelli/go-rifa-microservice/pkg/rabbitmq/rmq_pub/pub"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Mongo Repository
	mdb, err := mongodb.New(cfg.MDB.URL, cfg.MDB.Database)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}

	// Raffle Use Case
	raffleUseCase := usecase.NewRaffleUseCase(
		mongodbrepo.NewRaffle(mdb),
	)

	// Steam Use Case
	rmqPub, err := pub.New(cfg.RMQ.URL, cfg.RMQ.PubExchange)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - rmqServer - server.New: %w", err))
	}

	steamUseCase := usecase.NewSteam(
		mongodbrepo.NewPlayerSkin(mdb),
		webapi.NewSteamAPI(rmqPub),
	)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, raffleUseCase, steamUseCase)
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
