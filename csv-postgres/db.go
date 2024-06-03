package main

import (
	"context"
	"fmt"
	"os"
	"runtime"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

func connectToPostgres() (*pgxpool.Pool, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	config.MaxConns = int32(runtime.NumCPU() * 2)
	return pgxpool.ConnectConfig(context.Background(), config)
}

func insertRecords(dbPool *pgxpool.Pool, insertQuery string, records <-chan []string) {
	for record := range records {
		_, err := dbPool.Exec(context.Background(), insertQuery, convertRecordToInterface(record)...)
		if err != nil {
			log.WithFields(logrus.Fields{
				"event": "insert_record",
				"error": err,
			}).Error("Failed to insert record")
		}
	}
}
