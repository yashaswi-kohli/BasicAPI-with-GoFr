package controller

import (
	"encoding/json"
	"fmt"

	"github.com/yashaswi-kohli/BasicAPI/model"
	"github.com/yashaswi-kohli/BasicAPI/mongo"

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

func GetBooksAuthor(ctx *gofr.Context) (interface{}, error) {
	author := ctx.PathParam("author")
	fmt.Println("the author name is", author)
	books, err := mongo.GetMyBookAuthor(author)

	if err != nil {
		return nil, err
	}

	var theBooks []model.Book
	bookByte, _ := json.Marshal(books)
	json.Unmarshal(bookByte, &theBooks)
	return theBooks, nil
}

func GetBookISBN(ctx *gofr.Context) (interface{}, error) {
	isbn := ctx.PathParam("isbn")
	book, err := mongo.GetMyBookIsbn(isbn)

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

	//* checking whether the given json is valid or not
	if err := isJsonValid(book); err != nil {
		return nil, err
	}

	//* this will check whether the book is already present or not
	newBookIsbn := book.ISBN
	_, err := mongo.GetMyBookIsbn(newBookIsbn)

	if err == nil {
		return nil, &errors.Response{
			Reason: "There is already a book present with the given ISBN",
		}
	}

	//* if the book is not present then it will be inserted
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

	//* checking whether the user is trying to update the isbn or not
	id := ctx.PathParam("isbn")
	if book.ISBN != "" {
		return nil, &errors.Response{
			Reason: "ISBN could not be updated, once it it published",
		}
	}

	if book.Author == "" && book.Title == "" && book.Publisher == "" {
		return nil, &errors.Response{
			Reason: "Check the format of data, or the name of the fields, it should be author, title and publisher",
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
	_, err := mongo.GetMyBookIsbn(isbn)

	//* checking whether the book is present or not in the library
	if err != nil {
		return nil, &errors.Response{
			Reason: "There is no book present with the given ISBN",
		}
	}

	//* if the book is present then it will proceed to delete it
	numberOfItemDeleted, err := mongo.DeleteMyBook(isbn)

	if err != nil {
		return nil, err
	}
	return numberOfItemDeleted, err
}

func isJsonValid(book model.Book) error {
	byteBook, _ := json.Marshal(book)
	valid := json.Valid(byteBook)

	if !valid {
		return &errors.Response{
			Reason: "Json is not in valid format",
		}
	}

	//* checking whether the required fields are present or not
	if book.Author == "" {
		return &errors.Response{
			Reason: "The name of the author is missing.",
		}
	}

	if book.Title == "" {
		return &errors.Response{
			Reason: "The title of the book is missing.",
		}
	}

	if book.Publisher == "" {
		return &errors.Response{
			Reason: "The name of the publisher is missing.",
		}
	}

	if len(book.ISBN) != 17 {
		return &errors.Response{
			Reason: "isbn lenght should be 17",
		}
	}

	return nil
}
