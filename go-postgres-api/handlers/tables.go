package handlers

import (
	"context"
	"go-postgres-api/db"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTables(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	rows, err := db.GetDB().Query(ctx, "SELECT tablename FROM pg_tables WHERE schemaname='public'")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tables = append(tables, table)
	}

	c.JSON(http.StatusOK, tables)
}
