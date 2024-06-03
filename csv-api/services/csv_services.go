package services

import (
	"csv-api/models"
	"csv-api/utils"
	"fmt"
	"strconv"
	"time"

	"github.com/patrickmn/go-cache"
)

// CSVService represents the service for managing CSV data
type CSVService struct {
	csvReader utils.CSVReader
	cache     *cache.Cache
}

// NewCSVService creates a new CSV service
func NewCSVService(csvReader utils.CSVReader, cache *cache.Cache) *CSVService {
	return &CSVService{
		csvReader: csvReader,
		cache:     cache,
	}
}

// FetchAllData fetches all data from the CSV file
func (s *CSVService) FetchAllData() ([]models.Record, error) {
	if data, found := s.cache.Get("allData"); found {
		return data.([]models.Record), nil
	}

	csvData, err := s.csvReader.ReadCSV("data/data.csv")
	if err != nil {
		return nil, err
	}

	records, err := s.csvReader.ConvertToRecords(csvData)
	if err != nil {
		return nil, err
	}

	s.cache.Set("allData", records, 5*time.Minute)
	return records, nil
}

// FetchDataByID fetches data by ID from the CSV file
func (s *CSVService) FetchDataByID(id string) (models.Record, error) {
	if data, found := s.cache.Get("allData"); found {
		for _, record := range data.([]models.Record) {
			if strconv.Itoa(record.ID) == id {
				return record, nil
			}
		}
		return models.Record{}, fmt.Errorf("record not found")
	}

	csvData, err := s.csvReader.ReadCSV("data/data.csv")
	if err != nil {
		return models.Record{}, err
	}

	records, err := s.csvReader.ConvertToRecords(csvData)
	if err != nil {
		return models.Record{}, err
	}

	s.cache.Set("allData", records, 5*time.Minute)

	for _, record := range records {
		if strconv.Itoa(record.ID) == id {
			return record, nil
		}
	}
	return models.Record{}, fmt.Errorf("record not found")
}
