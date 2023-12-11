package testing

import (
	"encoding/json"

	"github.com/yashaswi-kohli/BasicAPI/model"
)

func IsJsonValid(book model.Book) string {
	byteBook, err := json.Marshal(book)
	valid := json.Valid(byteBook)

	if !valid || err != nil {
		return "JSON is not in valid format"
	}

	if book.Author == "" {
		return "Book author should be mentioned"
	}

	if book.Title == "" {
		return "Book title should be mentioned"
	}

	if len(book.Isbn) != 17 {
		return "ISBN is not in valid format, It's length should be 17"
	}

	return "JSON is Valid"
}
