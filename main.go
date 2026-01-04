package main

import (
	"database/sql"
	"log"

	"github.com/afan104/products-api/internal/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// GET /products 200
// GET /products/:guid 200
// POST /products 201 (new product)
// PUT /products/:guid 200 (updated product)
// DELETE /products/:guid 204



func main(){
	var router = gin.Default()
	var address = ":3000"
	var db *sql.DB
	var e error
	if db, e = sql.Open("sqlite3", "./data.db"); e != nil {
		log.Fatalf("Error: %v", e)
	}
	defer db.Close()

	if e := db.Ping(); e != nil {
		log.Fatalf("Error: %v", e)
	}

	router.GET("/products", controllers.GetProducts(db))
	router.GET("/products/:guid", controllers.GetProduct(db))
	router.POST("/products", controllers.PostProduct(db))
	router.PUT("/products/:guid", controllers.PutProduct(db))
	router.DELETE("/products/:guid", controllers.DeleteProduct(db))

	log.Fatalln(router.Run(address))
}