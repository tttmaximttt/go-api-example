package books

// fake data
var Books = map[string]Book{
	"0123466": Book{Title: "Hello Go", Author: "Maksym Radko", ISBN: "0123466"},
	"2342332": Book{Title: "Hello Cloud Native", Author: "Maksym Radko", ISBN: "2342332"},
}

func AllBooks() []Book {
	var b []Book
	for _, v := range Books {
		b = append(b, v)
	}
	return b
}

func createBook(book Book) (isbn string, created bool) {
	_, ok := Books[book.ISBN]

	if ok {
		return book.ISBN, false
	}

	Books[book.ISBN] = book
	b, ok := Books[book.ISBN]

	return b.ISBN, ok
}

func updateBook(isbn string, book Book) (exist bool) {
	_, ok := Books[isbn]

	if !ok {
		return ok
	}

	Books[isbn] = book
	return true
}

func deleteBook(isbn string) (deleted bool) {
	_, ok := Books[isbn]

	if !ok {
		return false
	}

	delete(Books, isbn)
	return true
}