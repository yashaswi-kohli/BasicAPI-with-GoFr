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
)

const connectionString = "mongodb+srv://yashylibrary:yashy@cluster0.if9dil8.mongodb.net/?retryWrites=true&w=majority"
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

// * Let's get a single book
func GetMyBook(bookID string) (primitive.M, error) {
	id, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{Key: "_id", Value: id}}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var myBook primitive.M
	if cursor.Next(context.Background()) {
		err := cursor.Decode(&myBook)
		if err != nil {
			return nil, err
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
func UpdateMyBook(bookID string, book model.Book) error {

	//? this will convert string into id which mongoDB can accept
	id, _ := primitive.ObjectIDFromHex(bookID)
	newBook := mapper.ConvertStructToBSONMap(book, nil)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": newBook}

	result, err := collection.UpdateMany(context.Background(), filter, update)

	if err != nil {
		return err
	}
	fmt.Println("Total number of values updated are: ", result.ModifiedCount)
	return nil
}

// * Let's delete one book
func DeleteMyBook(bookID string) (*mongo.DeleteResult, error) {
	id, _ := primitive.ObjectIDFromHex(bookID)
	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return nil, err
	}
	return deleteCount, nil
}
