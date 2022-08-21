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
	"github.com/evmartinelli/go-rifa-microservice/internal/core/skin"
	"github.com/evmartinelli/go-rifa-microservice/pkg/httpserver"
	"github.com/evmartinelli/go-rifa-microservice/pkg/logger"
	"github.com/evmartinelli/go-rifa-microservice/pkg/mongodb"
)

type Context struct {
	cfg        *config.Config
	rafflerepo raffle.RaffleRepo
	logger     *logger.Logger
}

func NewContext() *Context {
	return &Context{}
}

func (c *Context) UseCases() *v1.UseCases {
	return &v1.UseCases{
		GenerateRaffle:  c.GenerateRaffleUseCase(),
		ListRaffle:      c.ListRaffleUseCase(),
		PlayerInventory: nil,
	}
}

func (c *Context) GenerateRaffleUseCase() *raffle.GenerateRaffleUseCase {
	return raffle.NewGenerateRaffleUseCase(c.PostRepo())
}

func (c *Context) ListRaffleUseCase() *raffle.ListRaffleUseCase {
	return raffle.NewListRaffleUseCase(c.PostRepo())
}

func (c *Context) PlayerInventoryUseCase() *skin.PlayerInventoryUseCase {
	return nil
}

func (c *Context) PostRepo() raffle.RaffleRepo {
	if c.rafflerepo == nil {
		c.rafflerepo = rafflerepo.NewDynamoDBRaffleRepo(c.DB())
	}

	return c.rafflerepo
}

func (c *Context) DB() *mongodb.MongoDB {
	mdb, err := mongodb.New(c.cfg.MDB.URL, c.cfg.MDB.Database)
	if err != nil {
		panic(fmt.Errorf("app - Run - postgres.New: %w", err))
	}

	return mdb
}

func (c *Context) Config() *config.Config {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	c.cfg = cfg

	return c.cfg
}

func (c *Context) HTTPServer() {
	handler := gin.New()
	v1.NewRouter(handler, c.Logger(), *c.UseCases())
	httpServer := httpserver.New(handler, httpserver.Port(c.cfg.HTTP.Port))
	c.signal(httpServer)
}

func (c *Context) signal(httpServer *httpserver.Server) {
	var err error

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		c.Logger().Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		c.Logger().Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	err = httpServer.Shutdown()
	if err != nil {
		c.Logger().Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}

func (c *Context) Logger() *logger.Logger {
	if c.logger == nil {
		c.logger = logger.New(c.cfg.Log.Level)
	}

	return c.logger
}
