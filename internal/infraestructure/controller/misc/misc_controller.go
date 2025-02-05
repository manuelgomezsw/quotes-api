package misc

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quotes-api/internal/domain/quotes/service"
)

func GetAuthors(c *gin.Context) {
	authors, err := service.GetAuthors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting authors",
			"error":   err.Error(),
		})
		return
	}

	if authors == nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, authors)
}

func GetWorks(c *gin.Context) {
	works, err := service.GetWorks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting works",
			"error":   err.Error(),
		})
		return
	}

	if works == nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, works)
}
