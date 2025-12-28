package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteProduct(c *gin.Context){
	c.String(http.StatusOK, "delete product")
}