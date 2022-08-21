package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"

	"gorm.io/driver/postgres"
  	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	router.GET("/books/:id", func(context *gin.Context) {
		bookId := context.Param("id")
		Books := Books{}
		id, err := strconv.Atoi(bookId)
		if err != nil {
		    log.Error(err)
			return
		}
		Books.Id = id
		db.First(&Books)
		context.JSON(http.StatusOK, gin.H{
			"status": gin.H{
				"code":   http.StatusOK,
				"status": "success",
			},
			"id": &Books.Id,
			"title": &Books.Title,
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
		var json Books

		if err := context.ShouldBindJSON(&json); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			Books := Books{}
			Books.Title = json.Title
			Books.Price = json.Price
			db.Create(&Books)
		}

		// var json BookRequest
		// if err := context.ShouldBindJSON(&json); err != nil {
		// 	context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }
		// context.JSON(http.StatusOK,
		// 	gin.H{
		// 		"title": json.Title, "price": json.Price, "published": json.Published})
	})

	router.DELETE("/books/:id", func(context *gin.Context) {
		bookId := context.Param("id")
		Books := Books{}
		id, err := strconv.Atoi(bookId)
		if err != nil {
		    log.Error(err)
			return
		}
		Books.Id = id
		db.First(&Books)
		db.Delete(&Books)
	})

	router.PUT("/books/:id", func(context *gin.Context) {
		var json Books

		bookId := context.Param("id")
		Books := Books{}
		id, err := strconv.Atoi(bookId)
		if err != nil {
		    log.Error(err)
			return
		}
		Books.Id = id
		db.First(&Books)

		if err := context.ShouldBindJSON(&json); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			Books.Title = json.Title
			Books.Price = json.Price
			db.Save(&Books)
		}
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
