package router

import (
	"github.com/gin-gonic/gin"
	"quotes-api/internal/infraestructure/controller/misc"
	"quotes-api/internal/infraestructure/controller/quotes"
	"quotes-api/internal/infraestructure/controller/reviews"
	"quotes-api/internal/infraestructure/controller/words"
)

func mapURLs(router *gin.Engine) {
	quotesUrls(router)
	wordsUrls(router)
	reviewsUrls(router)
	miscUrls(router)
}

func quotesUrls(router *gin.Engine) {
	// Registry quotes
	router.POST("/quotes", quotes.Create)
	router.PUT("/quotes/:quote_id", quotes.Update)
	router.DELETE("/quotes/:quote_id", quotes.Delete)

	// Search quotes
	router.GET("/quotes/:quote_id", quotes.GetQuoteByID)
	router.GET("/quotes/author/:author", quotes.GetQuotesByAuthor)
	router.GET("/quotes/work/:work", quotes.GetQuotesByWork)
	router.GET("/quotes/keyword/:keyword", quotes.GetQuotesByKeyword)
	router.GET("/quotes/topics", quotes.GetTopics)

	// Daily job quotes
	router.POST("/quotes/daily", quotes.SendDailyQuote)
}

func wordsUrls(router *gin.Engine) {
	// Registry words
	router.POST("/words", words.Create)
	router.PUT("/words/:word_id", words.Update)
	router.DELETE("/words/:word_id", words.Delete)

	// Search words
	router.GET("/words/:word_id", words.GetByID)
	router.GET("/words/keyword/:keyword", words.GetByKeyword)
}

func reviewsUrls(router *gin.Engine) {
	// Registry reviews
	router.POST("/reviews", reviews.Create)
	router.PUT("/reviews/:review_id", reviews.Update)
	router.DELETE("/reviews/:review_id", reviews.Delete)

	// Search reviews
	router.GET("/reviews/:review_id", reviews.GetByID)
	router.GET("/reviews/title/:title", reviews.GetByTitle)
}

func miscUrls(router *gin.Engine) {
	router.GET("/authors", misc.GetAuthors)
	router.GET("/works", misc.GetWorks)
}
