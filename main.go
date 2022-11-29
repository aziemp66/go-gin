package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/aziemp66/go-gin/common/env"
	"github.com/aziemp66/go-gin/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := env.LoadConfig()

	dsn := cfg.DB_URL
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database Connected")

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
