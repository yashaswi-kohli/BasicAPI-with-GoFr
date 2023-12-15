package main

import (
	"github.com/yashaswi-kohli/BasicAPI/controller"
	"gofr.dev/pkg/gofr"
)

func main() {

	app := gofr.New()

	app.GET("/books", controller.GetBooks)
	app.GET("/books/isbn/{isbn}", controller.GetBookISBN)
	app.GET("/books/author/{author}", controller.GetBooksAuthor)
	app.POST("/books", controller.CreateBook)
	app.PUT("/books/{isbn}", controller.UpdateBook)
	app.DELETE("/books/{isbn}", controller.DeleteBook)

	// It start's the server on port 8000
	app.Start()
}
