package controller

import (
	"github.com/yashaswi-kohli/BasicAPI/model"
	"github.com/yashaswi-kohli/BasicAPI/mongo"
	"gofr.dev/pkg/gofr"
)

func GetBooks(ctx *gofr.Context) (interface{}, error) {
	books := mongo.GetAllBooks()
	return books, nil
}

func GetBook(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	books := mongo.GetMyBook(id)
	return books, nil
}

func CreateBook(ctx *gofr.Context) (interface{}, error) {
	var book model.Book
	ctx.Bind(&book)
	mongo.InsertMyBook(book)
	return book, nil
}

func UpdateBook(ctx *gofr.Context) (interface{}, error) {
	var book model.Book
	ctx.Bind(&book)
	id := ctx.PathParam("id")
	mongo.UpdateMyBook(id, book)
	return book, nil
}

func DeleteBook(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	numberOfItemDeleted := mongo.DeleteMyBook(id)
	return numberOfItemDeleted, nil
}
