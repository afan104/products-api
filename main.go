package main

import (
	"log"

	"github.com/afan104/products-api/internal/controllers"
	"github.com/gin-gonic/gin"
)

// GET /products 200
// GET /products/:guid 200
// POST /products 201 (new product)
// PUT /products/:guid 200 (updated product)
// DELETE /products/:guid 204



func main(){
	var router = gin.Default()
	var address = ":3000"

	router.GET("/products", controllers.GetProducts)
	router.GET("/products/:guid", controllers.GetProduct)
	router.POST("/products", controllers.PostProduct)
	router.PUT("/products/:guid", controllers.PutProduct)
	router.DELETE("/products/:guid", controllers.DeleteProduct)

	log.Fatalln(router.Run(address))
}