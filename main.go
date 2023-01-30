package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"quotes-api/api/controller"
)

func main() {
	router := gin.Default()
	router.GET("/", controller.GetQuotes)
	router.GET("/search", controller.GetQuotesByTag)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
