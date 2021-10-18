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

	// db.AutoMigrate(&book.Book{})

	// CRUD

	// Create
	/*
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
	*/

	// Read
	// var book book.Book

	// Slice of books
	var books []book.Book

	// First row
	// err = db.First(&book).Error

	// First row - primary key
	// err = db.First(&book, 2).Error

	// Last row
	// err = db.Last(&book).Error

	// Show query first
	// err = db.Debug().First(&book).Error

	// Show query last
	// err = db.Debug().Last(&book).Error

	// Retrieve many data objects
	// err = db.Debug().Find(&books).Error

	// Retrieve many data by string conditions
	err = db.Debug().Where("rating <> ?", 4).Find(&books).Error

	if err != nil {
		fmt.Println("==========================")
		fmt.Println("ERROR FINDING BOOK RECORD")
		fmt.Println("==========================")
	}

	// fmt.Println("Title:", book.Title)
	// fmt.Println("Price:", book.Price)
	// fmt.Println("Description:", book.Description)
	// fmt.Println("Book Object:", book)

	// Iterate through books slice
	for _, b := range books {
		fmt.Println("Title:", b.Title)
		fmt.Println("Objek:", b)
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
