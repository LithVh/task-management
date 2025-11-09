package database

import (
	"fmt"

	"task-management/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func PostgresConnect(config *config.Config) error {
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DBName,
	)

	logLevel := logger.Silent
	if config.Server.Env == "development" {
		logLevel = logger.Info
	}
	var err error
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return fmt.Errorf("failed to connect to DB - dbConnect: %v", err)
	}

	innerDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to reach inner DB - dbConnect:  %v", err)
	}

	err = innerDB.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping to DB server - dbConnect: %v", err)
	}

	return nil
}

func GetDB() *gorm.DB {
	return DB
}
