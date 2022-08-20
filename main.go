package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"

	"gorm.io/driver/postgres"
  	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func setupRouter(db *gorm.DB) *gin.Engine {
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
		Books := Books{}
		Books.Title = "test1"
		Books.Price = 200
		db.Create(&Books)

		// var json BookRequest
		// if err := context.ShouldBindJSON(&json); err != nil {
		// 	context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }
		// context.JSON(http.StatusOK,
		// 	gin.H{
		// 		"title": json.Title, "price": json.Price, "published": json.Published})
	})

	return router
}

func initDB() *gorm.DB {
	dsn := "host=localhost user=gin password=gin dbname=gin port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Info("[error]: fail connect database: %v\n", err)
	}

	log.Info("success connect database")
	return database
}

func main() {
	db := initDB()
	router := setupRouter(db)
	router.Run(":8080")
}

type Books struct {
	Id      int `json:"id"`
	Title     string `json:"title"`
	Price     int    `json:"price"`
}
