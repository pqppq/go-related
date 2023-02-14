package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pqppq/go-related/frameworks/gin/example/handler"
	"github.com/pqppq/go-related/frameworks/gin/example/repository"
	"go.uber.org/zap"
)

var (
	port   = ":8080"
	dbFile = "./example.db"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	repo := repository.NewBookRepo(dbFile)
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}
	bookHandler := handler.NewBookHandler(repo, logger)

	books := router.Group("/books")
	{
		books.GET("/", bookHandler.ShowBookList)
		books.POST("/", bookHandler.AddBook)
		books.GET("/:id", bookHandler.ShowBook)
		books.PUT("/:id", bookHandler.UpdateBook)
		books.DELETE("/:id", bookHandler.DeleteBook)
	}

	return router
}

func main() {
	router := setupRouter()

	router.Run(port)
}
