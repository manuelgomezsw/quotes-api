package repository

import (
	"quotes-api/internal/infraestructure"
	"quotes-api/internal/search/domain"
)

func GetQuoteByID(quoteID int64) (domain.CompleteQuote, error) {
	resultQuote, err := infraestructure.ClientDB.Query(
		"SELECT quote_id, author, work, phrase, date_created FROM `quotes-db`.`quotes` WHERE quote_id = ?", quoteID)
	if err != nil {
		return domain.CompleteQuote{}, err
	}

	var quote domain.CompleteQuote
	for resultQuote.Next() {
		err = resultQuote.Scan(&quote.QuoteID, &quote.Author, &quote.Work, &quote.Phrase, &quote.DateCreated)
		if err != nil {
			return domain.CompleteQuote{}, err
		}
	}

	return quote, nil
}

func GetQuotesByKeyword(keyword string) ([]domain.CompleteQuote, error) {
	resultQuote, err := infraestructure.ClientDB.Query(
		"SELECT quote_id, author, work, phrase, date_created FROM `quotes-db`.`quotes` WHERE phrase LIKE ?", "%"+keyword+"%")
	if err != nil {
		return nil, err
	}

	var quotes []domain.CompleteQuote
	for resultQuote.Next() {
		var quote domain.CompleteQuote

		err = resultQuote.Scan(&quote.QuoteID, &quote.Author, &quote.Work, &quote.Phrase, &quote.DateCreated)
		if err != nil {
			return nil, err
		}

		quotes = append(quotes, quote)
	}

	return quotes, nil
}

func GetQuotesByAuthor(author string) ([]domain.CompleteQuote, error) {
	resultQuote, err := infraestructure.ClientDB.Query(
		"SELECT quote_id, author, work, phrase, date_created FROM `quotes-db`.`quotes` WHERE author LIKE ?", "%"+author+"%")
	if err != nil {
		return nil, err
	}

	var quotes []domain.CompleteQuote
	for resultQuote.Next() {
		var quote domain.CompleteQuote

		err = resultQuote.Scan(&quote.QuoteID, &quote.Author, &quote.Work, &quote.Phrase, &quote.DateCreated)
		if err != nil {
			return nil, err
		}

		quotes = append(quotes, quote)
	}

	return quotes, nil
}

func GetQuotesByWork(work string) ([]domain.CompleteQuote, error) {
	resultQuote, err := infraestructure.ClientDB.Query(
		"SELECT quote_id, author, work, phrase, date_created FROM `quotes-db`.`quotes` WHERE work LIKE ?", "%"+work+"%")
	if err != nil {
		return nil, err
	}

	var quotes []domain.CompleteQuote
	for resultQuote.Next() {
		var quote domain.CompleteQuote

		err = resultQuote.Scan(&quote.QuoteID, &quote.Author, &quote.Work, &quote.Phrase, &quote.DateCreated)
		if err != nil {
			return nil, err
		}

		quotes = append(quotes, quote)
	}

	return quotes, nil
}
