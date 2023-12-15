package mongo

import (
	"context"
	"fmt"
	"log"

	"github.com/naamancurtis/mongo-go-struct-to-bson/mapper"
	"github.com/yashaswi-kohli/BasicAPI/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gofr.dev/pkg/errors"
)

const connectionString = "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.1.0"
const dbName = "library"
const collectionName = "bookShelf"

var collection *mongo.Collection

// * Let's  connect with mongo
func init() {
	//? client option
	clientOption := options.Client().ApplyURI(connectionString)

	//* connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(collectionName)
}

// * Let's get all books
func GetAllBooks() ([]primitive.M, error) {

	var books []primitive.M

	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var book primitive.M
		err := cursor.Decode(&book)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	defer cursor.Close(context.Background())
	return books, nil
}

// * Let's get all books with author name
func GetMyBookAuthor(author string) ([]primitive.M, error) {
	filter := bson.D{{Key: "author", Value: author}}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Println("Error querying MongoDB:", err)
		return nil, err
	}
	defer func() {
		if err := cursor.Close(context.Background()); err != nil {
			log.Println("Error closing cursor:", err)
		}
	}()

	var books []primitive.M
	for cursor.Next(context.Background()) {
		var book primitive.M
		err := cursor.Decode(&book)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if len(books) == 0 {
		return nil, &errors.Response{
			Reason: "There is no book present with the given author name",
		}
	}
	return books, nil
}

// * Let's get a single book
func GetMyBookIsbn(bookIsbn string) (primitive.M, error) {
	filter := bson.D{{Key: "isbn", Value: bookIsbn}}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Println("Error querying MongoDB:", err)
		return nil, err
	}
	defer func() {
		if err := cursor.Close(context.Background()); err != nil {
			log.Println("Error closing cursor:", err)
		}
	}()

	var myBook primitive.M
	if cursor.Next(context.Background()) {
		if err := cursor.Decode(&myBook); err != nil {
			log.Println("Error decoding document:", err)
			return nil, err
		}
	}

	if myBook == nil {
		return nil, &errors.Response{
			Reason: "There is no book present with the given ISBN",
		}
	}
	return myBook, nil
}

// * Let's insert one book
func InsertMyBook(book model.Book) error {
	insertied, err := collection.InsertOne(context.Background(), book)

	if err != nil {
		return err
	}
	fmt.Println("Inserted 1 Book in db with id: ", insertied.InsertedID)
	return nil
}

// * Let's update one book
func UpdateMyBook(bookIsbn string, updateItems model.Book) (primitive.M, error) {

	//* this will convert struct into primitive.M
	newBook := mapper.ConvertStructToBSONMap(updateItems, nil)

	update := bson.M{"$set": newBook}
	filter := bson.M{"isbn": bookIsbn}

	result, err := collection.UpdateMany(context.Background(), filter, update)

	if err != nil {
		return nil, err
	}
	fmt.Println("Total number of values updated are: ", result.ModifiedCount)

	book, err := GetMyBookIsbn(bookIsbn)
	if err != nil {
		return nil, err
	}
	return book, nil
}

// * Let's delete one book
func DeleteMyBook(bookIsbn string) (*mongo.DeleteResult, error) {

	filter := bson.M{"isbn": bookIsbn}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return nil, err
	}
	return deleteCount, nil
}
