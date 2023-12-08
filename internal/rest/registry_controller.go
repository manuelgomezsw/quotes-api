package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quotes-api/internal/registry/domain"
	"quotes-api/internal/registry/services"
)

func CreateQuote(c *gin.Context) {
	var quote domain.Quote
	services.CreateQuoteService(quote)

	c.JSON(http.StatusCreated, quote)
}

func UpdateQuote(c *gin.Context) {
	var quote domain.Quote
	services.UpdateQuoteService(quote)

	c.JSON(http.StatusOK, quote)
}

func DeleteQuote(c *gin.Context) {
	services.DeleteQuoteService(c.GetString("quote_id"))
	c.JSON(http.StatusOK, "Quote deleted")
}
