package handlers

import (
	"csv-api/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// ErrorResponse represents the structure of error responses
type ErrorResponse struct {
	Message string `json:"message"`
}

// SuccessResponse represents the structure of success responses
type SuccessResponse struct {
	Data interface{} `json:"data"`
}

// respondJSON sends a JSON response with the given status code and payload
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		logrus.WithError(err).Error("Error encoding response")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// GetData handles the request to fetch all data
func GetData(service *services.CSVService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logrus.WithFields(logrus.Fields{
			"client_ip": r.RemoteAddr,
			"method":    r.Method,
			"url":       r.URL.String(),
		}).Info("Received request to fetch all data")

		data, err := service.FetchAllData()
		if err != nil {
			logrus.WithError(err).Error("Error fetching data")
			respondJSON(w, http.StatusInternalServerError, ErrorResponse{Message: "Internal Server Error"})
			return
		}

		respondJSON(w, http.StatusOK, SuccessResponse{Data: data})
	}
}

// GetDataByID handles the request to fetch data by ID
func GetDataByID(service *services.CSVService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		logrus.WithFields(logrus.Fields{
			"client_ip": r.RemoteAddr,
			"method":    r.Method,
			"url":       r.URL.String(),
			"id":        id,
		}).Info("Received request to fetch data for ID")

		record, err := service.FetchDataByID(id)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"client_ip": r.RemoteAddr,
				"id":        id,
			}).WithError(err).Error("Error fetching data for ID")
			respondJSON(w, http.StatusNotFound, ErrorResponse{Message: "Record not found"})
			return
		}

		respondJSON(w, http.StatusOK, SuccessResponse{Data: record})
	}
}
