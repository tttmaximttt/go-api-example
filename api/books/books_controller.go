package books

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/tttmaximttt/learning/api"
)

func BooksHandlerFunc(w http.ResponseWriter, r *http.Request)  {
	switch method := r.Method; method {
	case http.MethodGet:
		books := AllBooks()
		api.WriteJSON(w, books)

	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		book := FromJson(body)
		isbn, created := createBook(book)

		if created {
			w.Header().Add("Location", "/api/books/"+isbn)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "No method specified")
	}
}

func BookHandlerFunc(w http.ResponseWriter, r *http.Request) {
	isbn := r.URL.Path[len("/api/books/"):]
	switch method := r.Method; method {
	case http.MethodGet:
		b, ok := Books[isbn]

		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "No book specified with isbn '%s'", isbn)
			return
		}

		w.Header().Add("Content-Type", "application/json encoding=utf-8")
		w.WriteHeader(http.StatusOK)
		api.WriteJSON(w, b.ToJSON())
	case http.MethodDelete:
		deleted := deleteBook(isbn)

		if deleted {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		b := FromJson(body)

		exist := updateBook(isbn, b)

		if exist {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "No method specified")
	}
}