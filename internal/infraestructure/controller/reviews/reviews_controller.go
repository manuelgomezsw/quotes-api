package reviews

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quotes-api/internal/domain/reviews"
	"quotes-api/internal/domain/reviews/service"
	"strconv"
)

func Create(c *gin.Context) {
	var newReview reviews.Review
	if err := c.ShouldBindJSON(&newReview); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing body",
			"error":   err.Error(),
		})
		return
	}

	if err := service.Create(&newReview); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error posting review",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, newReview)
}

func Update(c *gin.Context) {
	var review reviews.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing body",
			"error":   err.Error(),
		})
		return
	}

	reviewID, err := strconv.ParseInt(c.Param("review_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing review_id",
			"error":   err.Error(),
		})
		return
	}
	review.ReviewID = reviewID

	if err := service.Update(&review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating review",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, review)
}

func Delete(c *gin.Context) {
	reviewID, err := strconv.ParseInt(c.Param("review_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing review_id",
			"error":   err.Error(),
		})
		return
	}

	if err := service.Delete(reviewID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting review",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func GetByID(c *gin.Context) {
	reviewID, err := strconv.ParseInt(c.Param("review_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing review_id",
			"error":   err.Error(),
		})
		return
	}

	review, err := service.GetByID(reviewID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting review",
			"error":   err.Error(),
		})
		return
	}

	if review.ReviewID == 0 {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, review)
}

func GetByTitle(c *gin.Context) {
	title := c.Param("title")

	reviewsByTitle, err := service.GetByTitle(title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting review",
			"error":   err.Error(),
		})
		return
	}

	if len(reviewsByTitle) == 0 {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, reviewsByTitle)
}
