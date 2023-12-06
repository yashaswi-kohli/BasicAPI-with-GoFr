package main

import (
	"strconv"

	"gofr.dev/pkg/gofr"
)

type Book struct {
	ID     string `json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main() {

	// initialise gofr object
	app := gofr.New()

	books = append(books, Book{ID: "1", Isbn: "438227", Title: "Book One", Author: "Yashaswi Kohli"})
	books = append(books, Book{ID: "2", Isbn: "454555", Title: "Book Two", Author: "Yash Roye"})

	app.GET("/books", getBooks)
	app.GET("/book/{id}", getBook)
	app.POST("/book", createBook)
	app.PUT("/book/{id}", updateBook)
	app.DELETE("/book/{id}", deleteBook)

	// Starts the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Start()
}

func getBooks(ctx *gofr.Context) (interface{}, error) {
	return books, nil
}

func getBook(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	for _, book := range books {
		if book.ID == id {
			return book, nil
		}
	}
	return "Book Not Found", nil
}

func createBook(ctx *gofr.Context) (interface{}, error) {
	var book Book
	ctx.Bind(&book)
	lastIdx := len(books) - 1
	idx, _ := strconv.Atoi(books[lastIdx].ID)
	book.ID = strconv.Itoa(idx + 1)
	books = append(books, book)
	return book, nil
}

func updateBook(ctx *gofr.Context) (interface{}, error) {
	var newBook Book
	ctx.Bind(&newBook)
	id := ctx.PathParam("id")

	for index, book := range books {
		if book.ID == id {
			books[index] = newBook
			return books[index], nil
		}
	}
	return "Book Not Found", nil
}

func deleteBook(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")

	for index, book := range books {
		if book.ID == id {
			books = append(books[:index], books[index+1:]...)
			return books, nil
		}
	}
	return "Book Not Found", nil
}
