package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quotes-api/api/domain"
	"quotes-api/api/repository"
)

func GetQuotes(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")

	quotes, err := repository.GetLatestQuotes(5)
	if err != nil {
		apiErr := domain.NewInternalServerApiError(err.Error(), err)
		c.IndentedJSON(apiErr.Status(), apiErr)
		return
	}

	c.IndentedJSON(http.StatusOK, quotes)
}
