package main

import (
	"codeles/library/controllers"
	"fmt"
	"html"
	"log"
	"net/http"
)

func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	fmt.Println("Opening the library...")

	http.HandleFunc("/books/", controllers.BookRouter)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
