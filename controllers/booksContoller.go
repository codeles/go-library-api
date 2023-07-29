package controllers

import (
	"codeles/library/data"
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
)

var (
	listBooksRe    = regexp.MustCompile(`^\/users[\/]*$`)
	findBookByIdRe = regexp.MustCompile(`^\/users\/(\d+)$`)
)

// handler for requests for /books/
func BookRouter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch {
	case r.Method == http.MethodGet && listBooksRe.MatchString(r.URL.Path):
		jsonBytes, err := json.Marshal(getBooks())
		if err != nil {
			notFound(w, r)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
		return
	case r.Method == http.MethodGet && findBookByIdRe.MatchString(r.URL.Path):
		matches := findBookByIdRe.FindStringSubmatch(r.URL.Path)
		if len(matches) < 2 {
			// error situation
			notFound(w, r)
			return
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			//error situation
			notFound(w, r)
			return
		}
		book := getBookById(int32(id))
		jsonBytes, err := json.Marshal(book)

		if err != nil {
			notFound(w, r)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
		return
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

// get a book by id
func getBookById(id int32) data.Book {
	book := data.Books[0]

	return book
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}
