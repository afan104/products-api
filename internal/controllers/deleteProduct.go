package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context){
		c.String(http.StatusOK, "delete product")
	}
}