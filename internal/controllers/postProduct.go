package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type postProduct struct {
	Name string `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required,gt=0"`
	Description string `json:"description" binding:"omitempty,max=250"`
}

func PostProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context){
		var payload postProduct

		if e := c.ShouldBindJSON(&payload); e != nil {
			c.JSON(http.StatusBadRequest, e.Error())
			return
		}
		c.String(http.StatusOK, "post product")
	}
}