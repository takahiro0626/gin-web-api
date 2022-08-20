package main

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/books/:id", func(context *gin.Context) {
		id := context.Param("id")
		context.JSON(http.StatusOK, gin.H{
			"status": gin.H{
				"code":   http.StatusOK,
				"status": "success",
			},
			"id": id,
		})
	})

	router.GET("/books", func(context *gin.Context) {
		price := context.Query("price")
		published := context.Query("published")
		context.JSON(http.StatusOK, gin.H{
			"status": gin.H{
				"code":   http.StatusOK,
				"status": "success",
			},
			"price":     price,
			"published": published,
		})
	})

	router.POST("/books", func(context *gin.Context) {
		var json BookRequest
		if err := context.ShouldBindJSON(&json); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK,
			gin.H{
				"title": json.Title, "price": json.Price, "published": json.Published})
	})

	return router
}

func initDB() *sql.DB {
	database, err := sql.Open("postgres", "user=gin password=gin host=localhost port=5432 dbname=gin sslmode=disable")
	if err != nil {
		log.Info("[error]: fail connect database: %v\n", err)
	}

	err = database.Ping()
	if err != nil {
		log.Info("[error]: fail connect database: %v\n", err)
	}

	log.Info("success connect database")
	return database
}

func main() {
	db := initDB()
	defer db.Close()
	router := setupRouter()
	router.Run(":8080")
}

type BookRequest struct {
	Title     string `json:"title"`
	Price     int    `json:"price"`
	Published int    `json:"published"`
}
