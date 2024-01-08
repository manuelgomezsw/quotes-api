package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"quotes-api/internal/rest"
)

func main() {
	router := gin.Default()

	router.Use(CORSMiddleware())

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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", os.Getenv("CORS_ORIGIN"))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
