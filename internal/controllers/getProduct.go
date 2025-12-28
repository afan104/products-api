package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context){
	c.String(http.StatusOK, "get product")
}