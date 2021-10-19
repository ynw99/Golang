package main

import (
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

	// Repository Layer
	bookRepository := book.NewRepository(db)

	// Service Layer
	bookService := book.NewService(bookRepository)

	// bookRequest := book.BookRequest{
	// 	Title: "Komik Gundam",
	// 	Price: "200000",
	// }

	// bookService.Create(bookRequest)

	// Book Handler
	bookHandler := handler.NewBookHandler(bookService)

	// FindAll
	// books, err := bookRepository.FindAll()

	// for _, book := range books {
	// 	fmt.Println("Title:", book.Title)
	// }

	// FindByID
	// book, err := bookRepository.FindByID(2)

	// fmt.Println(book.Title)

	// Struct method create
	// book := book.Book{
	// 	Title:       "$100 Startup",
	// 	Description: "Good Book",
	// 	Price:       95000,
	// 	Rating:      4,
	// 	Discount:    0,
	// }

	// bookRepository.Create(book)

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
	// // var book book.Book

	// // Slice of books
	// var books []book.Book

	// // First row
	// // err = db.First(&book).Error

	// // First row - primary key
	// // err = db.First(&book, 2).Error

	// // Last row
	// // err = db.Last(&book).Error

	// // Show query first
	// // err = db.Debug().First(&book).Error

	// // Show query last
	// // err = db.Debug().Last(&book).Error

	// // Retrieve many data objects
	// // err = db.Debug().Find(&books).Error

	// // Retrieve many data by string conditions
	// err = db.Debug().Where("rating <> ?", 4).Find(&books).Error

	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("ERROR FINDING BOOK RECORD")
	// 	fmt.Println("==========================")
	// }

	// // fmt.Println("Title:", book.Title)
	// // fmt.Println("Price:", book.Price)
	// // fmt.Println("Description:", book.Description)
	// // fmt.Println("Book Object:", book)

	// // Iterate through books slice
	// for _, b := range books {
	// 	fmt.Println("Title:", b.Title)
	// 	fmt.Println("Objek:", b)
	// }

	// // Update
	// var book book.Book

	// err = db.Debug().Where("id = ?", 1).First(&book).Error

	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("ERROR FINDING BOOK RECORD")
	// 	fmt.Println("==========================")
	// }

	// book.Title = "Man Tiger (Revised Edition)"

	// err = db.Save(&book).Error

	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("ERROR UPDATING BOOK RECORD")
	// 	fmt.Println("==========================")
	// }

	// // Delete
	// var book book.Book

	// err = db.Debug().Where("id = ?", 1).First(&book).Error

	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("ERROR FINDING BOOK RECORD")
	// 	fmt.Println("==========================")
	// }

	// err = db.Delete(&book).Error

	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("ERROR DELETING BOOK RECORD")
	// 	fmt.Println("==========================")
	// }

	router := gin.Default()
	v1 := router.Group("/v1")

	v1.GET("/", bookHandler.RootHandler)
	v1.GET("/hello", bookHandler.HelloHandler)
	v1.GET("/books/:id/:title", bookHandler.BooksHandler)
	v1.GET("/query", bookHandler.QueryHandler)
	v1.POST("/books", bookHandler.PostBooksHandler)

	router.Run(":8080")

	/*
		1. main
		2. handler
		3. service
		4. repository
		5. db
		6. mysql
	*/
}
