package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context){
	c.String(http.StatusOK, "get products")
}