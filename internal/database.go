package internal

import (
	"fmt"
	"time"

	"github.com/cristian0193/golang-service-template/internal/utils"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB     *gorm.DB
	logger *zap.SugaredLogger
}

func New(logger *zap.SugaredLogger) *Database {
	return &Database{logger: logger}
}

func (db *Database) Open() error {
	host, err := utils.GetString("DB_HOST")
	if err != nil {
		return err
	}
	user, err := utils.GetString("DB_USER")
	if err != nil {
		return err
	}
	pass, err := utils.GetString("DB_PASSWORD")
	if err != nil {
		return err
	}
	database, err := utils.GetString("DB_DATABASE")
	if err != nil {
		return err
	}
	port, err := utils.GetString("DB_PORT")
	if err != nil {
		return err
	}

	if db.DB == nil {
		conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			host, user, pass, database, port)
		openDB, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
		if err != nil {
			return err
		}
		sqlDB, err := openDB.DB()
		if err != nil {
			return err
		}
		sqlDB.SetConnMaxLifetime(5 * time.Minute)
		sqlDB.SetConnMaxIdleTime(10)
		sqlDB.SetMaxOpenConns(10)

		db.DB = openDB
		db.logger.Info("connection established with the database")
	}
	return nil
}
