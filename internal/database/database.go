package database

import (
	"context"
	"fmt"
	"time"

	"task-management/internal/config"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// var DB *gorm.DB

func PostgresConnect(config *config.Config) (*gorm.DB, error) {
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
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB - PostgresConnect: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	innerDB, err := db.WithContext(ctx).DB()
	if err != nil {
		return nil, fmt.Errorf("failed to reach inner DB - PostgresConnect:  %v", err)
	}

	err = innerDB.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping to DB server - PostgresConnect: %v", err)
	}

	return db, nil
}

// func GetDB() *gorm.DB {
// 	return DB
// }

func RedisConnect(config *config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host + ":" + config.Server.Port,
		Password: config.Redis.Password,
		DB:       config.Redis.DBName,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to ping to DB server - RedisConnect: %v", err)
	}

	return rdb, nil
}
