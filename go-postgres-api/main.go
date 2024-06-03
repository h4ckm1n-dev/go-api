package main

import (
	"context"
	"go-postgres-api/config"
	"go-postgres-api/db"
	"go-postgres-api/handlers"
	"go-postgres-api/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file, continuing with environment variables")
	}

	cfg := config.LoadConfig()

	db.InitDB(cfg)
	defer db.CloseDB()

	// Initialize logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Can't initialize zap logger: %v", err)
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			log.Printf("Error syncing logger: %v", err)
		}
	}()

	r := gin.New()

	// Use zap logger middleware
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	// Custom middleware
	r.Use(middleware.RequestLogger(logger))
	r.Use(middleware.SecurityHeaders())

	// Routes
	r.GET("/tables", handlers.GetTables)
	r.GET("/table/:name", handlers.GetTableData)

	srv := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: r,
	}

	// Graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to run server", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown", zap.Error(err))
	}

	logger.Info("Server exiting")
}
