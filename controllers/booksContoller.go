package controllers

import (
	"codeles/library/data"
	"net/http"
	"regexp"
)

// handler for requests for /books/
func BookRouter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch {
	case r.Method == http.MethodGet && regexp.MustCompile(`^\/users[\/]*$`).MatchString(r.URL.Path):
		getBooks()
		return
		case 
	}

	return
}

// return all books
func getBooks() []data.Book {
	return data.Books
}

// get a book by title
func getBookByTitle(title string) data.Book {
	book := data.Books[0]

	return book
}
