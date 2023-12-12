package stores

import (
	"encoding/json"

	"github.com/yashaswi-kohli/BasicAPI/model"
	"gofr.dev/pkg/errors"
)

func IsJsonValid(book model.Book) error {
	byteBook, err := json.Marshal(book)
	valid := json.Valid(byteBook)

	if !valid || err != nil {
		return &errors.Response{
			Reason: "Json is not in valid format",
		}
	}

	if book.Author == "" {
		return errors.MissingParam{Param: []string{"author"}}
	}

	if book.Title == "" {
		return errors.MissingParam{Param: []string{"title"}}
	}

	if book.Publisher == "" {
		return errors.MissingParam{Param: []string{"publisher"}}
	}

	if len(book.ISBN) != 17 {
		return &errors.Response{
			Reason: "isbn lenght should be 17",
		}
	}

	return nil
}
