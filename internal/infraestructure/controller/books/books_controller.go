package books

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quotes-api/internal/domain/books"
	"quotes-api/internal/domain/books/service"
	"quotes-api/internal/util/conversions"
	"strconv"
)

func Create(c *gin.Context) {
	var newBook books.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing body",
			"error":   err.Error(),
		})
		return
	}

	if err := service.Create(&newBook); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error posting book",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, newBook)
}

func Update(c *gin.Context) {
	var currentBook books.Book
	if err := c.ShouldBindJSON(&currentBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing body",
			"error":   err.Error(),
		})
		return
	}

	bookID, err := strconv.ParseInt(c.Param("book_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing book_id",
			"error":   err.Error(),
		})
		return
	}
	currentBook.BookID, err = conversions.SafeIntConversion(bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting id book",
			"error":   err.Error(),
		})
		return
	}

	if err := service.Update(&currentBook); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating book",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, currentBook)
}

func Delete(c *gin.Context) {
	paramBookID, err := strconv.ParseInt(c.Param("book_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing book_id",
			"error":   err.Error(),
		})
		return
	}
	bookID, err := conversions.SafeIntConversion(paramBookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting id book",
			"error":   err.Error(),
		})
		return
	}

	if err := service.Delete(bookID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting book",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func Get(c *gin.Context) {
	allBooks, err := service.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting book",
			"error":   err.Error(),
		})
		return
	}

	if len(allBooks) == 0 {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, allBooks)
}
