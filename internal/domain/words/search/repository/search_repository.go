package repository

import (
	"quotes-api/internal/domain/words"
	"quotes-api/internal/util/mysql"
)

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
