package registry

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quotes-api/internal/domain/quotes"
	"quotes-api/internal/domain/quotes/registry/services"
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

	if err := services.CreateQuoteService(&newQuote); err != nil {
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

	if err := services.UpdateQuoteService(quoteID, &quote); err != nil {
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

	if err := services.DeleteQuoteService(quoteID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting quote",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, nil)
}
