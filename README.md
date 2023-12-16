# 💡 Library Management System

This is a simple and user-friendly REST API for managing a library's book collection. The API provides endpoints to perform various operations, such as retrieving all books, fetching a particular book by ISBN or author, updating existing book information, creating new books, and deleting books from the library. Implemented in Go with GoFr framework and MongoDB, the system showcases CRUD operations, and seamless database integration.

## Getting Started

1. Clone the repository:

    ```bash
    git clone https://github.com/yashaswi-kohli/BasicAPI-with-GoFr

    cd BasicAPI-with-GoFr
    ```

2. Download Dependencies:

    ```bash
    go mod download
    ```

3. Verify Dependencies:

    ```bash
    cat go.sum
    ```

4. Run Project

    ```bash
    go run main.go
    ```

5. Open Server

    ```bash
    http://localhost:8000/books
    ```

This will show all the books present in your database.

## API Requests made using postman

## 1.🚀 GET Request

-   Get all Books

    -   **Endpoint:**

        -   `GET /books`

    -   **Description:**

        -   Retrieves a list of all books.

    -   **Response:**

![image](https://github.com/yashaswi-kohli/BasicAPI-with-GoFr/assets/76786303/18cdc591-18ac-4535-96ee-155dfa2ddf5e)

-   Get all Books of a Author

    -   **Endpoint:**

        -   `GET /books/author/AUTHOR`

    -   **Description:**

        -   `AUTHOR`: Obtains a list of all books with the author's name.

    -   **Response:**

![image](https://github.com/yashaswi-kohli/BasicAPI-with-GoFr/assets/76786303/ebb04457-e3af-4fb3-a504-1c1e604fb6c7)

-   Get a book
-   **Endpoint:**

    -   `GET /books/isbn/ISBN`

-   **Description:**

    -   Retrieves a book with the isbn.
    -   `ISBN`: It is the unique identifier for the book.

-   **Response:**

![image](https://github.com/yashaswi-kohli/BasicAPI-with-GoFr/assets/76786303/51df8d91-2595-4389-87a7-a7996447296d)

## 2.➕ POST Request

-   Add a new book
-   **Endpoint:**

    -   `POST /books`

-   **Description:**

    -   Create a Book and add to database

-   **Response:**

![image](https://github.com/yashaswi-kohli/BasicAPI-with-GoFr/assets/76786303/f64c3d4e-559f-4e07-b6fc-8fd5e9e37ebe)

## 3.🔄 PUT Request

-   To update the book
-   **Endpoint:**

    -   `PUT /books/ISBN`

    -   **Description:**

        -   Update the existing book in the database

    -   **Response:**

![image](https://github.com/yashaswi-kohli/BasicAPI-with-GoFr/assets/76786303/1cad7898-8dd5-4f64-8fc4-8344e8c011c2)

## 4.🗑️ DELETE Request

-   To delete the book
-   **Endpoint:**

    -   `DELETE /books/ISBN`

    -   **Description:**

        -   Delete the book in the database, if it exist

    -   **Response:**

    ![image](https://github.com/yashaswi-kohli/BasicAPI-with-GoFr/assets/76786303/2a4640ef-893e-4ed5-8ca1-58ebef3e79b9)

# UML Diagrams

-   Sequence Diagram

-   ![image](https://github.com/yashaswi-kohli/BasicAPI-with-GoFr/assets/76786303/3daefb00-7535-4bd4-bcff-bbc297739e4a)

# Database 🛢️

The project utilizes a MongoDB database to store books. For MongoDB you can use both Atlas or Compass

# For Atlas

Go to website, make a cluster and set up a database, then copy the Database link. Then go to mongo folder and there in mongo.go file paste the link in connectionString and fillUP the dbName and collectionName.

# For Compass

First install Compass make a database and collection and remember the name, Then install mongosh, if it is already then run it in terminal by typing

```bash
mongosh
```

![image](https://github.com/yashaswi-kohli/BasicAPI-with-GoFr/assets/76786303/85b8f372-7456-458f-b1bf-d789e84e5be9)

Copy the link from 'Connecting to:' and then go to mongo folder and there in mongo.go file paste the link in connectionString and fillUP the dbName and collectionName.

# ⚡️ RUN

-   Now simply run the following command in your terminal to run the application:

```bash
go run cmd/main.go
```
