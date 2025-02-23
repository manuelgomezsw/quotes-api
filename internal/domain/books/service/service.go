package service

import (
	"quotes-api/internal/domain/books"
	"quotes-api/internal/domain/books/repository"
	"quotes-api/internal/util/customstrings"
)

func Get() ([]books.Book, error) {
	return repository.Get()
}

func Create(newBook *books.Book) error {
	newBook.Synopsis = customstrings.TruncateString(newBook.Synopsis, 200)
	return repository.Create(newBook)
}

func Update(currentBook *books.Book) error {
	currentBook.Synopsis = customstrings.TruncateString(currentBook.Synopsis, 200)
	return repository.Update(currentBook.BookID, currentBook)
}

func Delete(bookID int) error {
	return repository.Delete(bookID)
}
