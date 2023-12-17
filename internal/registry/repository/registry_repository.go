package repository

import (
	"quotes-api/internal/infraestructure"
	"quotes-api/internal/registry/domain"
)

func CreateQuote(newQuote *domain.Quote) error {
	newRecord, err := infraestructure.ClientDB.Exec(
		"INSERT INTO `quotes-db`.`quotes` (author, work, phrase) VALUES (?, ?, ?)",
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
	_, err := infraestructure.ClientDB.Exec(
		"UPDATE `quotes-db`.`quotes` SET author = ?, work = ?, phrase = ?  WHERE quote_id = ?",
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
	_, err := infraestructure.ClientDB.Exec(
		"DELETE FROM `quotes-db`.`quotes` WHERE quote_id = ?",
		quoteID,
	)
	if err != nil {
		return err
	}

	return nil
}
