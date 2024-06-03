package handlers

import (
	"context"
	"fmt"
	"go-postgres-api/db"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTableData(c *gin.Context) {
	tableName := c.Param("name")

	if tableName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Table name is required"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, err := db.GetDB().Query(ctx, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	columnNames := rows.FieldDescriptions()
	results := make([]map[string]interface{}, 0)
	for rows.Next() {
		values := make([]interface{}, len(columnNames))
		valuePtrs := make([]interface{}, len(columnNames))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		rowMap := make(map[string]interface{})
		for i, col := range columnNames {
			rowMap[string(col.Name)] = values[i]
		}

		results = append(results, rowMap)
	}

	c.JSON(http.StatusOK, results)
}
