// Package app configures and runs application.
package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/evmartinelli/go-rifa-microservice/config"
	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/idgenerator"
	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/orderrepo"
	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/pix/fake"
	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/rafflerepo"
	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/skinrepo"
	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/steam"
	v1 "github.com/evmartinelli/go-rifa-microservice/internal/controller/http/v1"
	"github.com/evmartinelli/go-rifa-microservice/internal/core/order"
	"github.com/evmartinelli/go-rifa-microservice/internal/core/raffle"
	"github.com/evmartinelli/go-rifa-microservice/internal/core/shared"
	"github.com/evmartinelli/go-rifa-microservice/internal/core/skin"
	db "github.com/evmartinelli/go-rifa-microservice/pkg/dynamodb"
	"github.com/evmartinelli/go-rifa-microservice/pkg/httpserver"
	"github.com/evmartinelli/go-rifa-microservice/pkg/logger"
)

const (
	DBTableConfigName string = "raffle_table"
	DBTableConfigPK   string = "PK"
	DBTableConfigSK   string = "SK"
)

type Context struct {
	cfg        *config.Config
	rafflerepo raffle.Repo
	logger     *logger.Logger
}

func NewContext() *Context {
	return &Context{}
}

func (c *Context) UseCases() *v1.UseCases {
	return &v1.UseCases{
		GenerateRaffle:  c.GenerateRaffleUseCase(),
		ListRaffle:      c.ListRaffleUseCase(),
		GetRaffle:       c.GetRaffleUseCase(),
		PlayerInventory: c.PlayerInventoryUseCase(),
		PlaceOrder:      c.PlaceOrderUseCase(),
	}
}

func (c *Context) GenerateRaffleUseCase() *raffle.GenerateRaffleUseCase {
	return raffle.NewGenerateRaffleUseCase(c.PostRepo(), c.UUIDGenerator(), c.SlugGenrator())
}

func (c *Context) ListRaffleUseCase() *raffle.ListRaffleUseCase {
	return raffle.NewListRaffleUseCase(c.PostRepo())
}

func (c *Context) GetRaffleUseCase() *raffle.GetRaffleUseCase {
	return raffle.NewGetRaffleUseCase(c.PostRepo())
}

func (c *Context) PlayerInventoryUseCase() *skin.PlayerInventoryUseCase {
	return skin.NewPlayerInventoryUseCase(c.PlayerSkinRepo(), c.SteamWebAPI())
}

func (c *Context) PlaceOrderUseCase() *order.PlaceOrderUseCase {
	return order.NewPlaceOrderUseCase(c.OrderRepo(), c.PixPayment(), c.UUIDGenerator())
}

func (c *Context) PostRepo() raffle.Repo {
	if c.rafflerepo == nil {
		c.rafflerepo = rafflerepo.NewDynamoDBRaffleRepo(c.DB())
	}

	return c.rafflerepo
}

func (c *Context) OrderRepo() order.Repo {
	return orderrepo.NewDynamoDBOrderRepo(c.DB())
}

func (c *Context) PlayerSkinRepo() skin.PlayerSkinRepo {
	return skinrepo.NewPlayerSkinRepo(c.DB())
}

func (c *Context) SteamWebAPI() skin.SteamWebAPI {
	return steam.NewSteamAPI()
}

func (c *Context) PixPayment() order.PixPayment {
	return fake.NewFakePixPayment()
}

func (c *Context) DB() *db.DynamoConfig {
	return db.NewDynamoDB(DBTableConfigName, DBTableConfigPK, DBTableConfigSK)
}

func (c *Context) Config() *config.Config {
	cfg, err := config.NewConfig()
	if err != nil {
		c.Logger().Fatal("Config error: %s", err)
	}

	c.cfg = cfg

	return c.cfg
}

func (c *Context) UUIDGenerator() shared.UUIDGenerator {
	return idgenerator.NewUUIDGenerator()
}

func (c *Context) SlugGenrator() shared.SlugGenerator {
	return idgenerator.NewSlugGenerator()
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
		c.Logger().Error("app - Run - httpServer.Notify: %w", err)
	}

	err = httpServer.Shutdown()
	if err != nil {
		c.Logger().Error("app - Run - httpServer.Shutdown: %w", err)
	}
}

func (c *Context) Logger() *logger.Logger {
	if c.logger == nil {
		c.logger = logger.New(c.cfg.Log.Level)
	}

	return c.logger
}
