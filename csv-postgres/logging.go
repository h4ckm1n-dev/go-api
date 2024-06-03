package main

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func setupLogging() {
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)

	// Add a hook to include timestamps in every log entry
	log.AddHook(&logrusHook{})
}

// logrusHook is a custom hook to add fields to every log entry
type logrusHook struct{}

func (hook *logrusHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *logrusHook) Fire(entry *logrus.Entry) error {
	entry.Data["timestamp"] = time.Now().Format(time.RFC3339)
	// Add more global fields if needed
	return nil
}

// LogExecutionTime logs the execution time of a function
func LogExecutionTime(start time.Time, functionName string) {
	elapsed := time.Since(start)
	log.WithFields(logrus.Fields{
		"function": functionName,
		"elapsed":  elapsed,
	}).Info("Execution time")
}
