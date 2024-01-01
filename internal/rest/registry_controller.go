package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quotes-api/internal/registry/domain"
	"quotes-api/internal/registry/services"
	"strconv"
)

func CreateQuote(c *gin.Context) {
	var newQuote domain.Quote
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

	c.Header("Access-Control-Allow-Origin", "http://localhost:4200")
	c.JSON(http.StatusOK, newQuote)
}

func UpdateQuote(c *gin.Context) {
	var quote domain.Quote
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

	c.Header("Access-Control-Allow-Origin", "http://localhost:4200")
	c.JSON(http.StatusOK, quote)
}

func DeleteQuote(c *gin.Context) {
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

	c.Header("Access-Control-Allow-Origin", "http://localhost:4200")
	c.JSON(http.StatusOK, nil)
}
