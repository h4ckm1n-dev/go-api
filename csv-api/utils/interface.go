package utils

import "csv-api/models"

// CSVReader defines the interface for reading and converting CSV data
type CSVReader interface {
	ReadCSV(filePath string) ([][]string, error)
	ConvertToRecords(csvData [][]string) ([]models.Record, error)
}
