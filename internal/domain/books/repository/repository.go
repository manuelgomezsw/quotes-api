package repository

import (
	"fmt"
	"os"
	"quotes-api/internal/domain/books"
	"quotes-api/internal/util/conversions"
	"quotes-api/internal/util/mysql"
)

const (
	basePathSqlQueries = "sql/books"

	fileSqlGet    = "Get.sql"
	fileSqlCreate = "Create.sql"
	fileSqlUpdate = "Update.sql"
	fileSqlDelete = "Delete.sql"
)

func Get() ([]books.Book, error) {
	query, err := os.ReadFile(fmt.Sprintf("%s/%s", basePathSqlQueries, fileSqlGet))
	if err != nil {
		return nil, err
	}

	result, err := mysql.ClientDB.Query(string(query))
	if err != nil {
		return nil, err
	}

	var allBooks []books.Book
	for result.Next() {
		var book books.Book

		err = result.Scan(&book.BookID, &book.Name, &book.Author, &book.Synopsis, &book.Source)
		if err != nil {
			return nil, err
		}

		allBooks = append(allBooks, book)
	}

	return allBooks, nil
}

func Create(newBook *books.Book) error {
	query, err := os.ReadFile(fmt.Sprintf("%s/%s", basePathSqlQueries, fileSqlCreate))
	if err != nil {
		return err
	}

	newRecord, err := mysql.ClientDB.Exec(
		string(query),
		newBook.Name,
		newBook.Author,
		newBook.Synopsis,
		newBook.Source,
	)
	if err != nil {
		return err
	}

	lastID, err := newRecord.LastInsertId()
	newBook.BookID, err = conversions.SafeIntConversion(lastID)

	if err != nil {
		return err
	}

	return nil
}

func Update(bookID int, currentBook *books.Book) error {
	query, err := os.ReadFile(fmt.Sprintf("%s/%s", basePathSqlQueries, fileSqlUpdate))
	if err != nil {
		return err
	}

	_, err = mysql.ClientDB.Exec(
		string(query),
		currentBook.Name,
		currentBook.Author,
		currentBook.Synopsis,
		currentBook.Source,
		bookID,
	)
	if err != nil {
		return err
	}

	currentBook.BookID = bookID

	return nil
}

func Delete(bookID int) error {
	query, err := os.ReadFile(fmt.Sprintf("%s/%s", basePathSqlQueries, fileSqlDelete))
	if err != nil {
		return err
	}

	_, err = mysql.ClientDB.Exec(
		string(query),
		bookID,
	)
	if err != nil {
		return err
	}

	return nil
}
