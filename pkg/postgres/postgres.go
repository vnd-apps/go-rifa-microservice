package lib

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/evmartinelli/go-rifa-microservice/pkg/logger"
)

const (
	maxOpenConns    = 60
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)

type Database struct {
	*gorm.DB
}

func NewDatabase(url string, log logger.Logger) Database {
	log.Info("opening db connection")

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{Logger: log.GetGormLogger()})
	if err != nil {
		log.Info("Url: ", url)
		log.Panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Info("DB: ", sqlDB)
		log.Panic(err)
	}

	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetConnMaxIdleTime(connMaxIdleTime)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(connMaxLifetime)

	return Database{
		DB: db,
	}
}
