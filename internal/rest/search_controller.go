package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetQuoteByID(c *gin.Context) {
	c.JSON(http.StatusOK, c.Param("quote_id"))
}

func GetQuotesByKeyword(c *gin.Context) {
	c.JSON(http.StatusOK, c.Param("keyword"))
}

func GetQuotesByAuthor(c *gin.Context) {
	c.JSON(http.StatusOK, c.Param("author"))
}

func GetQuotesByWork(c *gin.Context) {
	c.JSON(http.StatusOK, c.Param("work"))
}
