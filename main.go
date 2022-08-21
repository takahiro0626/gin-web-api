package main

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
			"id":    &Books.Id,
			"title": &Books.Title,
		})
	})

	router.GET("/books", func(context *gin.Context) {
		title := context.Query("title")
		var books []Books
		db.Where("title LIKE '%" + title + "%'").Find(&books)
		context.JSON(http.StatusOK, gin.H{"books": books})
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

	router.POST("/signup", func(context *gin.Context) {
		var json User

		// user := User{}
		// db.Where("email = ?", json.Email).First(&user)
		// if user.Id != 0 {
		// 	err := errors.New("Same Email Registered")
		// 	log.Error(err)
		// }

		hash, err := PasswordEncrypt(json.Password)
		if err != nil {
			err := errors.New("Fail Password Crypt")
			log.Error(err)
		}

		if err := context.ShouldBindJSON(&json); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {

			create_user := User{}
			create_user.Mail = json.Mail
			create_user.Password = hash
			db.Save(&create_user)
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

func PasswordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func main() {
	db := initDB()
	router := setupRouter(db)
	router.Run(":8080")
}

type Books struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
}

type User struct {
	Id       int    `json:"id"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
}
