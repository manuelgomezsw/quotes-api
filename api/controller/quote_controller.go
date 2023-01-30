package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quotes-api/api/domain"
	"quotes-api/api/repository"
	"strconv"
)

func GetQuotes(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")

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
