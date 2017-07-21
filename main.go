package main

import (
	"net/http"
	"os"
	"github.com/tttmaximttt/learning/api/books"
	"github.com/tttmaximttt/learning/api/index"
	"github.com/tttmaximttt/learning/api/echo"
)

func port() string {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8080"
	}

	return ":" + port
}

func main() {
	http.HandleFunc("/", index.Hello)
	http.HandleFunc("/api/echo", echo.Echo)

	http.HandleFunc("/api/books", books.BooksHandlerFunc)
	http.HandleFunc("/api/books/", books.BookHandlerFunc)

	http.ListenAndServe(port(), nil)
}
