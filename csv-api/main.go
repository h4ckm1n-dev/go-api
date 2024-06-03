package main

import (
	"csv-api/handlers"
	"csv-api/middleware"
	"csv-api/services"
	"csv-api/utils"
	"net/http"
	"os"
	"time"

	"github.com/didip/tollbooth/v6"
	"github.com/didip/tollbooth/v6/limiter"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8088"
	}

	c := cache.New(5*time.Minute, 10*time.Minute)
	csvService := services.NewCSVService(&utils.CSVUtils{}, c)

	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.SecurityHeadersMiddleware)

	lmt := tollbooth.NewLimiter(1, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})

	r.Handle("/data", tollbooth.LimitFuncHandler(lmt, handlers.GetData(csvService))).Methods("GET")
	r.Handle("/data/{id}", tollbooth.LimitFuncHandler(lmt, handlers.GetDataByID(csvService))).Methods("GET")

	logrus.WithFields(logrus.Fields{
		"port": port,
	}).Info("Server running")

	if err := http.ListenAndServe(":"+port, r); err != nil {
		logrus.WithError(err).Fatal("Server failed to start")
	}
}
