package books

import (
	"encoding/json"
)

type Book struct {
	Title string `json:"title"`
	Author string `json:"author"`
	ISBN string `json:"isbn"`
	Description string `json:"description,omitempty"`
}

func (book *Book) ToJSON() []byte {
	json, err := json.Marshal(book)

	if err != nil {
		panic(err)
	}

	return json
}

func FromJson(data []byte) Book {
	book := new(Book)
	err := json.Unmarshal(data, book)

	if err != nil {
		panic(err)
	}

	return *book
}