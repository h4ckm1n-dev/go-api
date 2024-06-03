package db

import (
	"context"
	"go-postgres-api/config"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

var pool *pgxpool.Pool

func InitDB(cfg config.Config) {
	config, err := pgxpool.ParseConfig(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Unable to parse DATABASE_URL: %v\n", err)
	}

	// Set maximum connections
	config.MaxConns = 10

	pool, err = pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
}

func GetDB() *pgxpool.Pool {
	return pool
}

func CloseDB() {
	pool.Close()
}
