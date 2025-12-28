package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostProduct(c *gin.Context){
	c.String(http.StatusOK, "post product")
}