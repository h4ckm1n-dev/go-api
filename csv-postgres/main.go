package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"io"
	"runtime"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	// Setup log formatting
	setupLogging()

	log.WithFields(logrus.Fields{
		"event": "start",
	}).Info("Starting the CSV to PostgreSQL importer")

	// Measure execution time
	defer LogExecutionTime(time.Now(), "main")

	// Read CSV from local folder
	folderPath := "/home/appuser/data/"
	csvFileName := "data.csv"
	csvFile, err := readCSVFromFolder(folderPath, csvFileName)
	if err != nil {
		log.WithFields(logrus.Fields{
			"event": "read_csv",
			"error": err,
		}).Fatal("Failed to read file from folder")
	}
	defer csvFile.Close()

	// Or read CSV from S3
	// bucket := os.Getenv("S3_BUCKET")
	// key := "path/to/your/csvfile.csv"
	// region := os.Getenv("AWS_REGION")
	// csvFile, err := readCSVFromS3(bucket, key, region)
	// if err != nil {
	//     log.WithFields(logrus.Fields{
	//         "event": "read_csv_s3",
	//         "error": err,
	//     }).Fatal("Failed to read file from S3")
	// }
	// defer csvFile.Close()

	// Read and parse CSV file
	reader := csv.NewReader(bufio.NewReader(csvFile))
	headers, err := reader.Read()
	if err != nil {
		log.WithFields(logrus.Fields{
			"event": "parse_csv",
			"error": err,
		}).Fatal("Failed to read CSV headers")
	}

	// Connect to PostgreSQL
	dbPool, err := connectToPostgres()
	if err != nil {
		log.WithFields(logrus.Fields{
			"event": "db_connect",
			"error": err,
		}).Fatal("Unable to connect to database")
	}
	defer dbPool.Close()

	log.WithFields(logrus.Fields{
		"event": "db_connect",
	}).Info("Connected to the database")

	// Create table if not exists
	tableName := "backend"
	createTableQuery := generateCreateTableQuery(tableName, headers)
	_, err = dbPool.Exec(context.Background(), createTableQuery)
	if err != nil {
		log.WithFields(logrus.Fields{
			"event": "create_table",
			"error": err,
		}).Fatal("Failed to create table")
	}

	log.WithFields(logrus.Fields{
		"event": "create_table",
		"table": tableName,
	}).Info("Table creation query executed")

	// Insert data into PostgreSQL using goroutines for concurrent processing
	insertQuery := generateInsertQuery(tableName, headers)
	records := make(chan []string)
	var wg sync.WaitGroup

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			insertRecords(dbPool, insertQuery, records)
		}()
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.WithFields(logrus.Fields{
				"event": "read_record",
				"error": err,
			}).Error("Error reading record")
			continue
		}
		records <- record
	}
	close(records)
	wg.Wait()

	log.WithFields(logrus.Fields{
		"event": "complete",
	}).Info("CSV data successfully inserted into PostgreSQL")
}
