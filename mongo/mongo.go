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

// * Improtant data which is table or collection
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

// ! Let's get all books
func GetAllBooks() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Panic(err)
	}

	var books []primitive.M

	for cursor.Next(context.Background()) {
		var book primitive.M
		err := cursor.Decode(&book)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}

	defer cursor.Close(context.Background())
	return books
}

// ! Let's insert one book
func InsertMyBook(book model.Book) {
	insertied, err := collection.InsertOne(context.Background(), book)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 Book in db with id: ", insertied.InsertedID)
}

// ! Let's update one book
func UpdateMyBook(bookID string, book model.Book) {

	//? this will convert string into id which mongoDB can accept
	id, _ := primitive.ObjectIDFromHex(bookID)
	newBook := mapper.ConvertStructToBSONMap(book, nil)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": newBook}

	result, err := collection.UpdateMany(context.Background(), filter, update)

	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Total number of values updated are: ", result.ModifiedCount)
}

// ! Let's delete one book
func DeleteMyBook(bookID string) *mongo.DeleteResult {
	id, _ := primitive.ObjectIDFromHex(bookID)
	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Panic(err)
	}
	return deleteCount
}
