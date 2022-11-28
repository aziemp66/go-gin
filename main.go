package main

import (
	"fmt"

	"github.com/aziemp66/go-gin/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1", func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		fmt.Println("Middleware Layer")
		ctx.Next()
	})

	v1.GET("/", handler.RootHandler)

	v1.GET("/hello", handler.HelloHandler)

	v1.GET("/books/:id/:title", handler.BooksHandler)

	v1.GET("/books", handler.QueryHandler)

	v1.POST("/books", handler.PostBookHandler)

	router.Run(":3000")
}
