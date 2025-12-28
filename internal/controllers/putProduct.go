package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PutProduct(c *gin.Context){
	c.String(http.StatusOK, "put product")
}