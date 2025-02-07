package repository

import (
	"errors"
	"fmt"
	sqlErr "github.com/go-mysql/errors"
	"os"
	"quotes-api/internal/domain/words"
	"quotes-api/internal/util/apierror"
	"quotes-api/internal/util/mysql"
)

const (
	basePathSqlQueries = "sql/words"

	fileSqlCreate         = "Create.sql"
	fileSqlUpdate         = "Update.sql"
	fileSqlDelete         = "Delete.sql"
	fileSqlGetByID        = "GetByID.sql"
	fileSqlGetByKeyword   = "GetByKeyword.sql"
	fileSqlGetMinMaxWords = "GetMinMaxWords.sql"
)

func Create(newWord *words.Word) apierror.ApiError {
	query, err := os.ReadFile(fmt.Sprintf("%s/%s", basePathSqlQueries, fileSqlCreate))
	if err != nil {
		return apierror.NewInternalServerApiError("error getting sql file", err)
	}

	newRecord, err := mysql.ClientDB.Exec(
		string(query),
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
	query, err := os.ReadFile(fmt.Sprintf("%s/%s", basePathSqlQueries, fileSqlUpdate))
	if err != nil {
		return err
	}

	_, err = mysql.ClientDB.Exec(
		string(query),
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
	query, err := os.ReadFile(fmt.Sprintf("%s/%s", basePathSqlQueries, fileSqlDelete))
	if err != nil {
		return err
	}

	_, err = mysql.ClientDB.Exec(
		string(query),
		wordID,
	)
	if err != nil {
		return err
	}

	return nil
}

func GetByID(wordID int64) (words.Word, error) {
	query, err := os.ReadFile(fmt.Sprintf("%s/%s", basePathSqlQueries, fileSqlGetByID))
	if err != nil {
		return words.Word{}, err
	}

	resultWord, err := mysql.ClientDB.Query(string(query), wordID)
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
	query, err := os.ReadFile(fmt.Sprintf("%s/%s", basePathSqlQueries, fileSqlGetByKeyword))
	if err != nil {
		return nil, err
	}

	keywordQuery := "%" + word + "%"
	resultWord, err := mysql.ClientDB.Query(string(query), keywordQuery, keywordQuery)
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

func GetMinMaxWords() (int64, int64, error) {
	query, err := os.ReadFile(fmt.Sprintf("%s/%s", basePathSqlQueries, fileSqlGetMinMaxWords))
	if err != nil {
		return 0, 0, err
	}

	resultWord, err := mysql.ClientDB.Query(string(query))
	if err != nil {
		return 0, 0, err
	}

	var minWordID, maxWordID int64
	for resultWord.Next() {
		err = resultWord.Scan(&minWordID, &maxWordID)
		if err != nil {
			return 0, 0, err
		}
	}

	return minWordID, maxWordID, nil
}
