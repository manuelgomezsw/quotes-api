package router

import (
	"github.com/gin-gonic/gin"
	"quotes-api/internal/infraestructure/controller/daily"
	"quotes-api/internal/infraestructure/controller/registry"
	"quotes-api/internal/infraestructure/controller/search"
)

func mapURLs(router *gin.Engine) {
	registryURLs(router)
	searchURLs(router)
	dailyURLs(router)
}

func registryURLs(router *gin.Engine) {
	router.POST("/quote", registry.CreateQuote)
	router.PUT("/quote/:quote_id", registry.UpdateQuote)
	router.DELETE("/quote/:quote_id", registry.DeleteQuote)
}

func searchURLs(router *gin.Engine) {
	router.GET("/quote/:quote_id", search.GetQuoteByID)
	router.GET("/quote/author/:author", search.GetQuotesByAuthor)
	router.GET("/quote/work/:work", search.GetQuotesByWork)
	router.GET("/quote/keyword/:keyword", search.GetQuotesByKeyword)
}

func dailyURLs(router *gin.Engine) {
	router.POST("/quote/daily", daily.SendDailyQuote)
}
