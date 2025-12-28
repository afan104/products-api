package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "get products")
	}
}