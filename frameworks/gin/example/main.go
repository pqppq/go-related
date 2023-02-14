package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pqppq/go-related/frameworks/gin/example/handler"
	"github.com/pqppq/go-related/frameworks/gin/example/repository"
)

var port = ":8080"

func main() {
	repo := repository.NewBookRepo("./example.db")
	bookHandler := handler.NewBookHandler(repo)

	route := gin.Default()
	books := route.Group("/books")
	{
		books.GET("/", bookHandler.ShowBookList)
		books.POST("/", bookHandler.AddBook)
		books.GET("/:id", bookHandler.ShowBook)
		books.PUT("/:id", bookHandler.UpdateBook)
		books.DELETE("/:id", bookHandler.DeleteBook)
	}

	route.Run(port)
}
