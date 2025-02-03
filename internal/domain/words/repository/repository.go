package repository

import (
	"errors"
	"fmt"
	sqlErr "github.com/go-mysql/errors"
	"quotes-api/internal/domain/words"
	"quotes-api/internal/util/apierror"
	"quotes-api/internal/util/mysql"
)

func Create(newWord *words.Word) apierror.ApiError {
	newRecord, err := mysql.ClientDB.Exec(
		"INSERT INTO quotes.words (word, meaning) VALUES (?, ?)",
		newWord.Word,
		newWord.Meaning,
	)
	if err != nil {
		if ok, errDb := sqlErr.Error(err); ok {
			if errors.Is(errDb, sqlErr.ErrDupeKey) {
				return apierror.NewConflictApiError(fmt.Sprintf("word [%s] already exist", newWord.Word))
			}
		}

		return apierror.NewInternalServerApiError("error adding word", err)
	}

	newWord.WordID, err = newRecord.LastInsertId()
	if err != nil {
		return apierror.NewInternalServerApiError("error getting last word inserted id", err)
	}

	return nil
}

func Update(currentWord *words.Word) error {
	_, err := mysql.ClientDB.Exec(
		"UPDATE quotes.words SET word = ?, meaning = ?  WHERE word_id = ?",
		currentWord.Word,
		currentWord.Meaning,
		currentWord.WordID,
	)
	if err != nil {
		return err
	}

	return nil
}

func Delete(wordID int64) error {
	_, err := mysql.ClientDB.Exec(
		"DELETE FROM quotes.words WHERE word_id = ?",
		wordID,
	)
	if err != nil {
		return err
	}

	return nil
}

func GetByID(wordID int64) (words.Word, error) {
	resultWord, err := mysql.ClientDB.Query(
		"SELECT word_id, word, meaning, date_created FROM quotes.words WHERE word_id = ?", wordID)
	if err != nil {
		return words.Word{}, err
	}

	var quote words.Word
	for resultWord.Next() {
		err = resultWord.Scan(&quote.WordID, &quote.Word, &quote.Meaning, &quote.DateCreated)
		if err != nil {
			return words.Word{}, err
		}
	}

	return quote, nil
}

func GetByKeyword(word string) ([]words.Word, error) {
	keywordQuery := "%" + word + "%"
	resultWord, err := mysql.ClientDB.Query(
		"SELECT word_id, word, meaning, date_created FROM quotes.words WHERE word LIKE ? OR meaning LIKE ?", keywordQuery, keywordQuery)
	if err != nil {
		return nil, err
	}

	var quotes []words.Word
	for resultWord.Next() {
		var quote words.Word

		err = resultWord.Scan(&quote.WordID, &quote.Word, &quote.Meaning, &quote.DateCreated)
		if err != nil {
			return nil, err
		}

		quotes = append(quotes, quote)
	}

	return quotes, nil
}
