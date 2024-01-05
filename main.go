package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func loadBooks() {
	data, err := os.ReadFile("./books.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &books)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Loaded %d books\n", len(books))
}

func main() {
	fmt.Println("Connecting to server...")

	loadBooks()

	handler := func(w http.ResponseWriter, r *http.Request) { // Handler function; takes a ResponseWriter and a Request

		template := template.Must(template.ParseFiles("./src/index.html")) // Parse the HTML template file

		// books := map[string][]Book{ // type map with string keys and Book values
		// 	"Books": {
		// 		{Title: "Book 1", Author: "Author 1"},
		// 		{Title: "Book 2", Author: "Author 2"},
		// 		{Title: "Book 3", Author: "Author 3"},
		// 	},
		// }

		data := map[string][]Book{"Books": books} // type map with string keys and Book values

		template.Execute(w, data) // Execute the template, passing in the books map

	}

	h2 := func(w http.ResponseWriter, r *http.Request) { // Handler function; takes a ResponseWriter and a Request
		log.Print("Adding book...")
		time.Sleep(1 * time.Second) // Simulate a long-running process for 1 second

		title := r.FormValue("title")   // Get the title value from the form
		author := r.FormValue("author") // Get the author value from the form

		//template fragments instead of htmlString
		template := template.Must(template.ParseFiles("./src/index.html")) // Parse the HTML template file
		template.ExecuteTemplate(w, "book-list-element", Book{Title: title, Author: author})

		log.Print("Book added!")
	}

	search := func(w http.ResponseWriter, r *http.Request) {
		query := strings.ToLower(r.FormValue("query"))
		fmt.Printf("Searching for %s...\n", query)

		var results []Book

		for _, book := range books {
			if strings.Contains(strings.ToLower(book.Title), query) || strings.Contains(strings.ToLower(book.Author), query) {
				results = append(results, book)
			}
		}

		fmt.Printf("Search results: %v\n", results)

		template := template.Must(template.ParseFiles("./src/search-books.html")) // Parse the HTML template file
		template.Execute(w, map[string][]Book{"Books": results})
	}

	http.HandleFunc("/", handler) // Handle all requests to the web root ("/") with the handler function

	http.HandleFunc("/add-book/", h2)

	http.HandleFunc("/search/", search)

	log.Fatal(http.ListenAndServe(":8080", nil)) // Listen on port 8080; if error, log fatal
}
