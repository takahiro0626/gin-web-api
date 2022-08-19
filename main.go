package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func main() {
	router := setupRouter()
	router.Run(":8080")
}

type BookRequest struct {
	Title  string `json:"title"`
	Price  int    `json:"price"`
	Published int   `json:"published"`
}