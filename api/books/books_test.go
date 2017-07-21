package books

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBooks_ToJSON(t *testing.T) {
	book := Book{Title: "Hello Go", Author: "Maksym Radko", ISBN: "0123466"}
	jsonBook := book.ToJSON()
	assert.Equal(
		t,
		`{"title":"Hello Go","author":"Maksym Radko","isbn":"0123466"}`,
		string(jsonBook),
		"Book ToJSON fail.")
}

func TestFromJson(t *testing.T) {
	json := []byte(`{"title":"Hello Go","author":"Maksym Radko","isbn":"0123466"}`)

	book := FromJson(json)
	assert.Equal(
		t,
		Book{Title:"Hello Go",Author:"Maksym Radko",ISBN:"0123466"},
		book,
		"Book FromJson fail.")
}