package controller

import (
	"github.com/yashaswi-kohli/BasicAPI/model"
	"github.com/yashaswi-kohli/BasicAPI/mongo"
	"github.com/yashaswi-kohli/BasicAPI/testing"
	"gofr.dev/pkg/gofr"
)

func GetBooks(ctx *gofr.Context) (interface{}, error) {
	books, err := mongo.GetAllBooks()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func GetBook(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	book, err := mongo.GetMyBook(id)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func CreateBook(ctx *gofr.Context) (interface{}, error) {
	var book model.Book
	ctx.Bind(&book)

	//? checking whether the given json is valid or not
	if err := testing.IsJsonValid(book); err != nil {
		return nil, err
	}

	err := mongo.InsertMyBook(book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func UpdateBook(ctx *gofr.Context) (interface{}, error) {
	var book model.Book
	ctx.Bind(&book)

	//? checking whether the given json is valid or not
	if err := testing.IsJsonValid(book); err != nil {
		return nil, err
	}

	id := ctx.PathParam("id")
	err := mongo.UpdateMyBook(id, book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func DeleteBook(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	numberOfItemDeleted, err := mongo.DeleteMyBook(id)

	if err != nil {
		return nil, err
	}
	return numberOfItemDeleted, err
}
