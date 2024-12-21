package repository

import (
	"quotes-api/internal/domain/quotes"
	"quotes-api/internal/util/mysql"
	"strings"
)

func CreateQuote(newQuote *quotes.Quote) error {
	newRecord, err := mysql.ClientDB.Exec(
		"INSERT INTO `quotes`.`quotes` (author, work, phrase) VALUES (?, ?, ?)",
		newQuote.Author,
		newQuote.Work,
		newQuote.Phrase,
	)
	if err != nil {
		return err
	}

	newQuote.QuoteID, err = newRecord.LastInsertId()
	if err != nil {
		return err
	}

	if newQuote.Tags != "" {
		if err = createTags(newQuote.QuoteID, newQuote.Tags); err != nil {
			return err
		}
	}

	return nil
}

func UpdateQuote(quoteID int64, currentQuote *quotes.Quote) error {
	_, err := mysql.ClientDB.Exec(
		"UPDATE `quotes`.`quotes` SET author = ?, work = ?, phrase = ?  WHERE quote_id = ?",
		currentQuote.Author,
		currentQuote.Work,
		currentQuote.Phrase,
		quoteID,
	)
	if err != nil {
		return err
	}

	if currentQuote.Tags != "" {
		if err = deleteTags(quoteID); err != nil {
			return err
		}

		if err = createTags(quoteID, currentQuote.Tags); err != nil {
			return err
		}
	}

	return nil
}

func DeleteQuote(quoteID int64) error {
	_, err := mysql.ClientDB.Exec(
		"DELETE FROM `quotes`.`quotes` WHERE quote_id = ?",
		quoteID,
	)
	if err != nil {
		return err
	}

	return nil
}

func createTags(quoteID int64, tags string) error {
	bulkInsert := "INSERT INTO tags (quote_id, tag) VALUES "
	values := []string{}
	args := []interface{}{}

	for _, tag := range strings.Split(tags, `,`) {
		values = append(values, "(?, ?)")
		args = append(args, quoteID, tag)
	}

	bulkInsert += strings.Join(values, ",")
	_, err := mysql.ClientDB.Exec(bulkInsert, args...)
	if err != nil {
		return err
	}
	return nil
}

func deleteTags(quoteID int64) error {
	_, err := mysql.ClientDB.Exec(
		"DELETE FROM tags WHERE quote_id = ?",
		quoteID,
	)
	if err != nil {
		return err
	}

	return nil
}
