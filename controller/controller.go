package controller

import (
	"github.com/yashaswi-kohli/BasicAPI/model"
	"github.com/yashaswi-kohli/BasicAPI/mongo"
	"github.com/yashaswi-kohli/BasicAPI/testing"
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

	//? checking whether the given json is valid or not
	err := testing.IsJsonValid(book)
	if err != "JSON is Valid" {
		return err, nil
	}

	mongo.InsertMyBook(book)
	return book, nil
}

func UpdateBook(ctx *gofr.Context) (interface{}, error) {
	var book model.Book
	ctx.Bind(&book)

	//? checking whether the given json is valid or not
	err := testing.IsJsonValid(book)
	if err != "JSON is Valid" {
		return err, nil
	}

	id := ctx.PathParam("id")
	mongo.UpdateMyBook(id, book)
	return book, nil
}

func DeleteBook(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	numberOfItemDeleted := mongo.DeleteMyBook(id)
	return numberOfItemDeleted, nil
}
