package middleware

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// LoggingMiddleware logs the details of each request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Create a response writer to capture the status code
		rr := &responseRecorder{w, http.StatusOK}

		// Call the next handler
		next.ServeHTTP(rr, r)

		duration := time.Since(startTime)

		logFields := logrus.Fields{
			"client_ip":  r.RemoteAddr,
			"method":     r.Method,
			"url":        r.URL.String(),
			"status":     rr.statusCode,
			"duration":   duration.Seconds(),
			"user_agent": r.UserAgent(),
			"headers":    r.Header,
		}

		logEntry := logrus.WithFields(logFields)
		if rr.statusCode >= 400 {
			logEntry.Error("Handled request")
		} else {
			logEntry.Info("Handled request")
		}
	})
}

// responseRecorder is a custom response writer that captures the status code
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code
func (rr *responseRecorder) WriteHeader(statusCode int) {
	rr.statusCode = statusCode
	rr.ResponseWriter.WriteHeader(statusCode)
}
