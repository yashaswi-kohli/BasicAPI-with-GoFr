package main

import (
	"github.com/yashaswi-kohli/BasicAPI/controller"
	"gofr.dev/pkg/gofr"
)

func main() {

	// initialise gofr object
	app := gofr.New()

	app.GET("/books", controller.GetBooks)
	app.GET("/books/{id}", controller.GetBook)
	app.POST("/books", controller.CreateBook)
	app.PUT("/books/{id}", controller.UpdateBook)
	app.DELETE("/books/{id}", controller.DeleteBook)

	// Starts the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Start()
}
