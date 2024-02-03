package repository

import (
	"quotes-api/internal/domain"
	"quotes-api/internal/util/mysql"
)

func GetQuoteByID(quoteID int64) (domain.Quote, error) {
	resultQuote, err := mysql.ClientDB.Query(
		"SELECT quote_id, author, work, phrase, date_created FROM `quotes`.`quotes` WHERE quote_id = ?", quoteID)
	if err != nil {
		return domain.Quote{}, err
	}

	var quote domain.Quote
	for resultQuote.Next() {
		err = resultQuote.Scan(&quote.QuoteID, &quote.Author, &quote.Work, &quote.Phrase, &quote.DateCreated)
		if err != nil {
			return domain.Quote{}, err
		}
	}

	return quote, nil
}

func GetQuotesByKeyword(keyword string) ([]domain.Quote, error) {
	resultQuote, err := mysql.ClientDB.Query(
		"SELECT quote_id, author, work, phrase, date_created FROM `quotes`.`quotes` WHERE phrase LIKE ?", "%"+keyword+"%")
	if err != nil {
		return nil, err
	}

	var quotes []domain.Quote
	for resultQuote.Next() {
		var quote domain.Quote

		err = resultQuote.Scan(&quote.QuoteID, &quote.Author, &quote.Work, &quote.Phrase, &quote.DateCreated)
		if err != nil {
			return nil, err
		}

		quotes = append(quotes, quote)
	}

	return quotes, nil
}

func GetQuotesByAuthor(author string) ([]domain.Quote, error) {
	resultQuote, err := mysql.ClientDB.Query(
		"SELECT quote_id, author, work, phrase, date_created FROM `quotes`.`quotes` WHERE author LIKE ?", "%"+author+"%")
	if err != nil {
		return nil, err
	}

	var quotes []domain.Quote
	for resultQuote.Next() {
		var quote domain.Quote

		err = resultQuote.Scan(&quote.QuoteID, &quote.Author, &quote.Work, &quote.Phrase, &quote.DateCreated)
		if err != nil {
			return nil, err
		}

		quotes = append(quotes, quote)
	}

	return quotes, nil
}

func GetQuotesByWork(work string) ([]domain.Quote, error) {
	resultQuote, err := mysql.ClientDB.Query(
		"SELECT quote_id, author, work, phrase, date_created FROM `quotes`.`quotes` WHERE work LIKE ?", "%"+work+"%")
	if err != nil {
		return nil, err
	}

	var quotes []domain.Quote
	for resultQuote.Next() {
		var quote domain.Quote

		err = resultQuote.Scan(&quote.QuoteID, &quote.Author, &quote.Work, &quote.Phrase, &quote.DateCreated)
		if err != nil {
			return nil, err
		}

		quotes = append(quotes, quote)
	}

	return quotes, nil
}
