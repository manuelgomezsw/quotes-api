package repository

import (
	"errors"
	"fmt"
	"quotes-api/internal/domain/words"
	"quotes-api/internal/util/apierror"
	"quotes-api/internal/util/mysql"

	sqlErr "github.com/go-mysql/errors"
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
