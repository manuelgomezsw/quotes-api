package daily

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quotes-api/internal/domain/quotes/daily/services"
)

func SendDailyQuote(c *gin.Context) {
	confirmationID, err := services.SendDailyQuote(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error sending daily quote",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":          "processed",
		"confirmation_id": confirmationID,
	})
}
