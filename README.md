## Go Book Server with HTMX

This is a simple HTTP server written in Go. It serves a list of books and provides only the functionality to add new books and search for existing ones. It makes use of the HTMX tool.

## Getting Started

To run this project, you need to have Go installed on your machine. If you don't have Go installed, you can download it from the [official website](https://golang.org/dl/).

## Usage

1. Clone this repository to your local machine.
2. Navigate to the project directory.
3. Run `go run main.go` to start the server.
4. Open your web browser and navigate to `http://localhost:8080`.

## Features

### View Books

When you navigate to the root URL (`http://localhost:8080`), the server will respond with a list of books.

### Add a Book

To add a book, make a POST request to `http://localhost:8080/add-book/` with the `title` and `author` parameters in the request body.

### Search for a Book

To search for a book, make a GET request to `http://localhost:8080/search/` with the `query` parameter in the request URL.

## Code Overview

The server is implemented in a single Go file, `main.go`.

### Data Structures

The `Book` struct represents a book with a title and an author.

### Functions

- `loadBooks`: This function reads the `books.json` file and loads the list of books into memory.
- `main`: This is the main function that starts the server and sets up the HTTP routes.

### HTTP Handlers

- The root handler (`/`) responds with the list of books.
- The `/add-book/` handler adds a new book to the list.
- The `/search/` handler searches the list of books for a given query.

## Future Improvements

Currently, the list of books is stored in memory and is not persistent. In the future, this could be improved by storing the books in a database.

