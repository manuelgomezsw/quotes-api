package words

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quotes-api/internal/domain/words"
	"quotes-api/internal/domain/words/service"
	"strconv"
)

func Create(c *gin.Context) {
	var newWord words.Word
	if err := c.ShouldBindJSON(&newWord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing body",
			"error":   err.Error(),
		})
		return
	}

	if err := service.Create(&newWord); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, newWord)
}

func Update(c *gin.Context) {
	var word words.Word
	if err := c.ShouldBindJSON(&word); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing body",
			"error":   err.Error(),
		})
		return
	}

	wordID, err := strconv.ParseInt(c.Param("word_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing word_id",
			"error":   err.Error(),
		})
		return
	}
	word.WordID = wordID

	if err := service.Update(&word); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating word",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, word)
}

func Delete(c *gin.Context) {
	wordID, err := strconv.ParseInt(c.Param("word_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing word_id",
			"error":   err.Error(),
		})
		return
	}

	if err := service.Delete(wordID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting word",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func GetByID(c *gin.Context) {
	wordID, err := strconv.ParseInt(c.Param("word_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing name",
			"error":   err.Error(),
		})
		return
	}

	word, err := service.GetByID(wordID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting word",
			"error":   err.Error(),
		})
		return
	}

	if word.WordID == 0 {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, word)
}

func GetByKeyword(c *gin.Context) {
	if c.Param("keyword") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "word is required",
		})
		return
	}

	keyword, err := service.GetByKeyword(c.Param("keyword"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting word",
			"error":   err.Error(),
		})
		return
	}

	if keyword == nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, keyword)
}

func GetRandomWord(c *gin.Context) {
	randomWord, err := service.GetRandomWord()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting random word",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, randomWord)
}
