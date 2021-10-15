package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "Yusuf Nur Wahid",
			"bio":  "A Software Engineer",
		})
	})

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"title":    "Hello World!",
			"subtitle": "Belajar Golang bareng Agung Setiawan",
		})
	})

	router.Run(":8888")
}
