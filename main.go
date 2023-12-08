package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"quotes-api/internal/rest"
)

func main() {
	router := gin.Default()

	router.POST("/quote", rest.CreateQuote)
	router.PUT("/quote/:quote_id", rest.UpdateQuote)
	router.DELETE("/quote/:quote_id", rest.DeleteQuote)
	router.GET("/quote/:quote_id", rest.GetQuoteByID)
	router.GET("/quote/author/:author", rest.GetQuotesByAuthor)
	router.GET("/quote/work/:work", rest.GetQuotesByWork)
	router.GET("/quote/keyword/:keyword", rest.GetQuotesByKeyword)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
