package utils

import (
	"csv-api/models"
	"encoding/csv"
	"os"
	"strconv"
)

// CSVUtils provides utility functions for working with CSV files
type CSVUtils struct{}

// ReadCSV reads a CSV file and returns records as a slice of string slices
func (u *CSVUtils) ReadCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	return reader.ReadAll()
}

// ConvertToRecords converts a slice of string slices to a slice of Record structs
func (u *CSVUtils) ConvertToRecords(csvData [][]string) ([]models.Record, error) {
	var records []models.Record
	for _, record := range csvData[1:] { // skip header
		id, _ := strconv.Atoi(record[0])
		quantity, _ := strconv.Atoi(record[5])
		price, _ := strconv.ParseFloat(record[6], 64)

		records = append(records, models.Record{
			ID:       id,
			Item:     record[1],
			Value:    record[2],
			Category: record[3],
			Date:     record[4],
			Quantity: quantity,
			Price:    price,
		})
	}
	return records, nil
}
