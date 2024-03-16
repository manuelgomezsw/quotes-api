package search

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quotes-api/internal/domain/quotes/search/services"
	"strconv"
)

func GetQuoteByID(c *gin.Context) {
	quoteID, err := strconv.ParseInt(c.Param("quote_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing quote_id",
			"error":   err.Error(),
		})
		return
	}

	quote, err := services.GetQuoteByID(quoteID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting quote",
			"error":   err.Error(),
		})
		return
	}

	if quote.QuoteID == 0 {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, quote)
}

func GetQuotesByKeyword(c *gin.Context) {
	if c.Param("keyword") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "keyword is required",
		})
		return
	}

	quotes, err := services.GetQuotesByKeyword(c.Param("keyword"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting quotes",
			"error":   err.Error(),
		})
		return
	}

	if quotes == nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, quotes)
}

func GetQuotesByAuthor(c *gin.Context) {
	if c.Param("author") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "author is required",
		})
		return
	}

	quotes, err := services.GetQuotesByAuthor(c.Param("author"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting quotes",
			"error":   err.Error(),
		})
		return
	}

	if quotes == nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, quotes)
}

func GetQuotesByWork(c *gin.Context) {
	if c.Param("work") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "work is required",
		})
		return
	}

	quotes, err := services.GetQuotesByWork(c.Param("work"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting quotes",
			"error":   err.Error(),
		})
		return
	}

	if quotes == nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, quotes)
}
