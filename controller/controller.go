package controller

import (
	"encoding/json"

	"github.com/yashaswi-kohli/BasicAPI/model"
	"github.com/yashaswi-kohli/BasicAPI/mongo"
	"github.com/yashaswi-kohli/BasicAPI/stores"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

func GetBooks(ctx *gofr.Context) (interface{}, error) {
	books, err := mongo.GetAllBooks()
	if err != nil {
		return nil, err
	}

	var theBooks []model.Book

	bookByte, _ := json.Marshal(books)
	json.Unmarshal(bookByte, &theBooks)
	return theBooks, nil
}

func GetBook(ctx *gofr.Context) (interface{}, error) {
	isbn := ctx.PathParam("isbn")
	book, err := mongo.GetMyBook(isbn)

	if err != nil {
		return nil, err
	}

	var theBook model.Book

	bookByte, _ := json.Marshal(book)
	json.Unmarshal(bookByte, &theBook)
	return theBook, nil
}

func CreateBook(ctx *gofr.Context) (interface{}, error) {
	var book model.Book
	ctx.Bind(&book)

	//? checking whether the given json is valid or not
	if err := stores.IsJsonValid(book); err != nil {
		return nil, err
	}

	newBookIsbn := book.ISBN
	_, err := mongo.GetMyBook(newBookIsbn)

	if err == nil {
		return nil, &errors.Response{
			Reason: "There is already a book present with the given ISBN",
		}
	}

	err = mongo.InsertMyBook(book)
	if err != nil {
		return nil, err
	}

	var theBook model.Book
	bookByte, _ := json.Marshal(book)
	json.Unmarshal(bookByte, &theBook)
	return theBook, nil
}

func UpdateBook(ctx *gofr.Context) (interface{}, error) {
	var book model.Book
	ctx.Bind(&book)

	id := ctx.PathParam("isbn")

	if book.ISBN != "" {
		return nil, &errors.Response{
			Reason: "ISBN could not be updated, once it it published",
		}
	}

	uBook, err := mongo.UpdateMyBook(id, book)
	if err != nil {
		return nil, err
	}

	var theBook model.Book
	bookByte, _ := json.Marshal(uBook)
	json.Unmarshal(bookByte, &theBook)
	return theBook, nil
}

func DeleteBook(ctx *gofr.Context) (interface{}, error) {

	isbn := ctx.PathParam("isbn")
	_, err := mongo.GetMyBook(isbn)

	if err != nil {
		return nil, &errors.Response{
			Reason: "There is no book present with the given ISBN",
		}
	}
	numberOfItemDeleted, err := mongo.DeleteMyBook(isbn)

	if err != nil {
		return nil, err
	}
	return numberOfItemDeleted, err
}
