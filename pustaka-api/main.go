package main

import (
	"fmt"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"gorm.io/driver/mysql"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	dsn := "admin:password@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection Error")
	}

	db.AutoMigrate(&book.Book{})

	// CRUD
	book := book.Book{}
	book.Title = "Atomic Habits"
	book.Price = 120000
	book.Discount = 15
	book.Rating = 4
	book.Description = "Buku self development tentang membangun kebiasaan baik dan menghilangkan kebiasaan buruk"

	err = db.Create(&book).Error
	if err != nil {
		fmt.Println("==========================")
		fmt.Println("ERROR CREATING BOOK RECORD")
		fmt.Println("==========================")
	}

	router := gin.Default()
	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run(":8080")
}
