package router

import (
	"github.com/gin-gonic/gin"
	quotesDaily "quotes-api/internal/infraestructure/controller/quotes/daily"
	quotesRegistry "quotes-api/internal/infraestructure/controller/quotes/registry"
	quotesSearch "quotes-api/internal/infraestructure/controller/quotes/search"
	reviewsRegistry "quotes-api/internal/infraestructure/controller/reviews/registry"
	reviewsSearch "quotes-api/internal/infraestructure/controller/reviews/search"
	wordsRegistry "quotes-api/internal/infraestructure/controller/words/registry"
	wordsSearch "quotes-api/internal/infraestructure/controller/words/search"
)

func mapURLs(router *gin.Engine) {
	quotesUrls(router)
	wordsUrls(router)
	reviewsUrls(router)
}

func quotesUrls(router *gin.Engine) {
	// Registry quotes
	router.POST("/quotes", quotesRegistry.Create)
	router.PUT("/quotes/:quote_id", quotesRegistry.Update)
	router.DELETE("/quotes/:quote_id", quotesRegistry.Delete)

	// Search quotes
	router.GET("/quotes/:quote_id", quotesSearch.GetQuoteByID)
	router.GET("/quotes/author/:author", quotesSearch.GetQuotesByAuthor)
	router.GET("/quotes/work/:work", quotesSearch.GetQuotesByWork)
	router.GET("/quotes/keyword/:keyword", quotesSearch.GetQuotesByKeyword)

	// Daily job quotes
	router.POST("/quotes/daily", quotesDaily.SendDailyQuote)
}

func wordsUrls(router *gin.Engine) {
	// Registry words
	router.POST("/words", wordsRegistry.Create)
	router.PUT("/words/:word_id", wordsRegistry.Update)
	router.DELETE("/words/:word_id", wordsRegistry.Delete)

	// Search words
	router.GET("/words/:word_id", wordsSearch.GetByID)
	router.GET("/words/keyword/:keyword", wordsSearch.GetByKeyword)
}

func reviewsUrls(router *gin.Engine) {
	// Registry reviews
	router.POST("/reviews", reviewsRegistry.Create)
	router.PUT("/reviews/:review_id", reviewsRegistry.Update)
	router.DELETE("/reviews/:review_id", reviewsRegistry.Delete)

	// Search reviews
	router.GET("/reviews/:review_id", reviewsSearch.GetByID)
}
