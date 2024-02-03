package repository

import (
	"quotes-api/internal/domain"
	"quotes-api/internal/util/mysql"
)

func CreateQuote(newQuote *domain.Quote) error {
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

	return nil
}

func UpdateQuote(quoteID int64, currentQuote *domain.Quote) error {
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
