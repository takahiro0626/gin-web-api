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
			"price": price,
			"published": published,
		})
	})

	return router
}

func main() {
	router := setupRouter()
	router.Run(":8080")
}
