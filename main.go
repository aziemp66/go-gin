package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)

	router.GET("/hello", helloHandler)

	router.GET("/books/:id/:title", booksHandler)

	router.GET("/books", queryHandler)

	router.POST("/books", postBookHandler)

	router.Run(":3000")
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":      "Azie Melza Pratama",
		"ismarried": false,
		"age":       19,
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"content": "Hello World",
	})
}

func booksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	author := c.Query("author")

	c.JSON(http.StatusOK, gin.H{
		"title":  title,
		"author": author,
	})
}

type BookInput struct {
	Title    string `json:"title" binding:"required"`
	Price    int    `json:"price" binding:"required,number"`
	SubTitle string `json:"sub_title" binding:"required"`
}

func postBookHandler(c *gin.Context) {
	var bookinput BookInput

	err := c.ShouldBindJSON(&bookinput)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title":     bookinput.Title,
		"price":     bookinput.Price,
		"sub_title": bookinput.SubTitle,
	})
}
