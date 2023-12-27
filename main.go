package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

	router.Use(CORS)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}

func CORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}
