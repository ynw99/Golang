package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BookInput struct {
	Title string      `json:"judul" binding:"required"`
	Price json.Number `json:"price" binding:"required,number"`
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Yusuf Nur Wahid",
		"bio":  "A Software Engineer",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":    "Hello World!",
		"subtitle": "Belajar Golang bareng Agung Setiawan",
	})
}

func booksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")
	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

func postBooksHandler(c *gin.Context) {
	var bookInput BookInput

	err := c.ShouldBindJSON(&bookInput)

	errorMessages := []string{}
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on filed %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"judul": bookInput.Title,
		"price": bookInput.Price,
	})
}

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")

	v1.GET("/", rootHandler)
	v1.GET("/hello", helloHandler)
	v1.GET("/books/:id/:title", booksHandler)
	v1.GET("/query", queryHandler)
	v1.POST("/books", postBooksHandler)

	router.Run(":8080")
}
