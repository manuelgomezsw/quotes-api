package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quotes-api/api/domain"
	"quotes-api/api/repository"
	"strconv"
)

func GetQuotesByTag(c *gin.Context) {
	setCORSHeaders(c)

	quotes, errorGettingQuotes := repository.GetQuotesByTag(c, c.Query("tag"))
	if errorGettingQuotes != nil {
		apiErr := domain.NewInternalServerApiError(errorGettingQuotes.Error(), errorGettingQuotes)
		c.IndentedJSON(apiErr.Status(), apiErr)
		return
	}

	c.IndentedJSON(http.StatusOK, quotes)
}

func GetQuotes(c *gin.Context) {
	setCORSHeaders(c)

	limitQuery, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limitQuery = 0
	}

	quotes, errorGettingQuotes := repository.GetLatestQuotes(limitQuery)
	if errorGettingQuotes != nil {
		apiErr := domain.NewInternalServerApiError(errorGettingQuotes.Error(), errorGettingQuotes)
		c.IndentedJSON(apiErr.Status(), apiErr)
		return
	}

	c.IndentedJSON(http.StatusOK, quotes)
}

func setCORSHeaders(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
}
