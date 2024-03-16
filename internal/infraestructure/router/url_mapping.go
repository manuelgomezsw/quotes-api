package router

import (
	"github.com/gin-gonic/gin"
	quotesDaily "quotes-api/internal/infraestructure/controller/quotes/daily"
	quotesRegistry "quotes-api/internal/infraestructure/controller/quotes/registry"
	quotesSearch "quotes-api/internal/infraestructure/controller/quotes/search"
	wordsRegistry "quotes-api/internal/infraestructure/controller/words/registry"
	wordsSearch "quotes-api/internal/infraestructure/controller/words/search"
)

func mapURLs(router *gin.Engine) {
	quotesUrls(router)
	wordsUrls(router)
}

func quotesUrls(router *gin.Engine) {
	// Registry quotes
	router.POST("/quote", quotesRegistry.Create)
	router.PUT("/quote/:quote_id", quotesRegistry.Update)
	router.DELETE("/quote/:quote_id", quotesRegistry.Delete)

	// Search quotes
	router.GET("/quote/:quote_id", quotesSearch.GetQuoteByID)
	router.GET("/quote/author/:author", quotesSearch.GetQuotesByAuthor)
	router.GET("/quote/work/:work", quotesSearch.GetQuotesByWork)
	router.GET("/quote/keyword/:keyword", quotesSearch.GetQuotesByKeyword)

	// Daily job quotes
	router.POST("/quote/daily", quotesDaily.SendDailyQuote)
}

func wordsUrls(router *gin.Engine) {
	// Registry words
	router.POST("/word", wordsRegistry.Create)
	router.PUT("/word/:word_id", wordsRegistry.Update)
	router.DELETE("/word/:word_id", wordsRegistry.Delete)

	// Search words
	router.GET("/word/:word_id", wordsSearch.GetByID)
	router.GET("/word/keyword/:keyword", wordsSearch.GetByKeyword)
}
