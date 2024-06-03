package main

import (
	"fmt"
	"strings"
)

func generateCreateTableQuery(tableName string, headers []string) string {
	columns := []string{}
	for _, header := range headers {
		columns = append(columns, fmt.Sprintf("%s TEXT", header))
	}
	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", tableName, strings.Join(columns, ", "))
}

func generateInsertQuery(tableName string, headers []string) string {
	columns := strings.Join(headers, ", ")
	placeholders := []string{}
	for i := range headers {
		placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
	}
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", tableName, columns, strings.Join(placeholders, ", "))
}

func convertRecordToInterface(record []string) []interface{} {
	values := make([]interface{}, len(record))
	for i, v := range record {
		values[i] = v
	}
	return values
}
