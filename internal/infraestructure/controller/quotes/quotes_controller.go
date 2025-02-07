package quotes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quotes-api/internal/domain/quotes"
	"quotes-api/internal/domain/quotes/service"
	"strconv"
)

func Create(c *gin.Context) {
	var newQuote quotes.Quote
	if err := c.ShouldBindJSON(&newQuote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing body",
			"error":   err.Error(),
		})
		return
	}

	if err := service.CreateQuoteService(&newQuote); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error posting quote",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, newQuote)
}

func Update(c *gin.Context) {
	var quote quotes.Quote
	if err := c.ShouldBindJSON(&quote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing body",
			"error":   err.Error(),
		})
		return
	}

	quoteID, err := strconv.ParseInt(c.Param("quote_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing quote_id",
			"error":   err.Error(),
		})
		return
	}

	if err := service.UpdateQuoteService(quoteID, &quote); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating quote",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, quote)
}

func Delete(c *gin.Context) {
	quoteID, err := strconv.ParseInt(c.Param("quote_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing quote_id",
			"error":   err.Error(),
		})
		return
	}

	if err := service.DeleteQuoteService(quoteID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting quote",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func GetQuoteByID(c *gin.Context) {
	quoteID, err := strconv.ParseInt(c.Param("quote_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing quote_id",
			"error":   err.Error(),
		})
		return
	}

	quote, err := service.GetQuoteByID(quoteID)
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

	quotesByKeyword, err := service.GetQuotesByKeyword(c.Param("keyword"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting quotes",
			"error":   err.Error(),
		})
		return
	}

	if quotesByKeyword == nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, quotesByKeyword)
}

func GetQuotesByAuthor(c *gin.Context) {
	if c.Param("author") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "author is required",
		})
		return
	}

	quotesByAuthor, err := service.GetQuotesByAuthor(c.Param("author"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting quotes",
			"error":   err.Error(),
		})
		return
	}

	if quotesByAuthor == nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, quotesByAuthor)
}

func GetQuotesByWork(c *gin.Context) {
	if c.Param("work") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "work is required",
		})
		return
	}

	quotesByWork, err := service.GetQuotesByWork(c.Param("work"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting quotes",
			"error":   err.Error(),
		})
		return
	}

	if quotesByWork == nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, quotesByWork)
}

func GetTopics(c *gin.Context) {
	quotesByTopic, err := service.GetTopics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting quotes",
			"error":   err.Error(),
		})
		return
	}

	if quotesByTopic == nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, quotesByTopic)
}

func GetRandomQuote(c *gin.Context) {
	randomQuote, err := service.GetRandomQuote()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting random quote",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, randomQuote)
}
