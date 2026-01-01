package database

import (
	"fmt"
	"log"

	"github.com/Listantiyo/pos-system/internal/config"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	DB		*gorm.DB
	Redis 	*redis.Client
}

func ConnectDB(cfg *config.Config) (*Database, error) {
	// Postgres connection
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBName,
		cfg.DBPort,
		cfg.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgresSQL database: %w", err)
	}

	// Resdis connection
	redisClient := redis.NewClient(&redis.Options{
		Addr: cfg.RedisHost + ":" + cfg.RedisPort,
		Password: cfg.RedisPass,
		DB: cfg.RedisDB,
	})

	log.Println("✅ Connected to Redis and PostgresSQL successfully.")
	return &Database{
		DB: db,
		Redis: redisClient,
	}, nil
}
