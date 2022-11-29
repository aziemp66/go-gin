package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/aziemp66/go-gin/book"
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":      "Azie Melza Pratama",
		"ismarried": false,
		"age":       19,
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"content": "Hello World",
	})
}

func BooksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	author := c.Query("author")

	c.JSON(http.StatusOK, gin.H{
		"title":  title,
		"author": author,
	})
}

func PostBookHandler(c *gin.Context) {
	var bookinput book.BookInput

	err := c.ShouldBindJSON(&bookinput)
	if err != nil {
		var ve validator.ValidationErrors

		if !errors.As(err, &ve) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": []string{err.Error()},
			})

			return
		}

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on failed %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title":     bookinput.Title,
		"price":     bookinput.Price,
		"sub_title": bookinput.SubTitle,
	})
}
