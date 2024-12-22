package repository

import (
	"quotes-api/internal/domain/quotes"
	"quotes-api/internal/util/mysql"
)

func GetQuoteByID(quoteID int64) (quotes.Quote, error) {
	resultQuote, err := mysql.ClientDB.Query(
		"SELECT quote_id, author, work, phrase, date_created FROM quotes WHERE quote_id = ?", quoteID)
	if err != nil {
		return quotes.Quote{}, err
	}

	var quote quotes.Quote
	for resultQuote.Next() {
		err = resultQuote.Scan(&quote.QuoteID, &quote.Author, &quote.Work, &quote.Phrase, &quote.DateCreated)
		if err != nil {
			return quotes.Quote{}, err
		}
	}

	return quote, nil
}

func GetQuotesByKeyword(keyword string) ([]quotes.Quote, error) {
	resultQuote, err := mysql.ClientDB.Query(
		"SELECT q.quote_id, q.author, q.work, q.phrase, q.date_created FROM quotes q JOIN tags t ON q.quote_id = t.quote_id WHERE t.tag LIKE ?", "%"+keyword+"%")
	if err != nil {
		return nil, err
	}

	var quotesSearched []quotes.Quote
	for resultQuote.Next() {
		var quote quotes.Quote

		err = resultQuote.Scan(&quote.QuoteID, &quote.Author, &quote.Work, &quote.Phrase, &quote.DateCreated)
		if err != nil {
			return nil, err
		}

		quotesSearched = append(quotesSearched, quote)
	}

	return quotesSearched, nil
}

func GetQuotesByAuthor(author string) ([]quotes.Quote, error) {
	resultQuote, err := mysql.ClientDB.Query(
		"SELECT quote_id, author, work, phrase, date_created FROM `quotes`.`quotes` WHERE author LIKE ?", "%"+author+"%")
	if err != nil {
		return nil, err
	}

	var quotesSearched []quotes.Quote
	for resultQuote.Next() {
		var quote quotes.Quote

		err = resultQuote.Scan(&quote.QuoteID, &quote.Author, &quote.Work, &quote.Phrase, &quote.DateCreated)
		if err != nil {
			return nil, err
		}

		quotesSearched = append(quotesSearched, quote)
	}

	return quotesSearched, nil
}

func GetQuotesByWork(work string) ([]quotes.Quote, error) {
	resultQuote, err := mysql.ClientDB.Query(
		"SELECT quote_id, author, work, phrase, date_created FROM `quotes`.`quotes` WHERE work LIKE ?", "%"+work+"%")
	if err != nil {
		return nil, err
	}

	var quotesSearched []quotes.Quote
	for resultQuote.Next() {
		var quote quotes.Quote

		err = resultQuote.Scan(&quote.QuoteID, &quote.Author, &quote.Work, &quote.Phrase, &quote.DateCreated)
		if err != nil {
			return nil, err
		}

		quotesSearched = append(quotesSearched, quote)
	}

	return quotesSearched, nil
}

func GetTopics() ([]quotes.Tag, error) {
	resultTopics, err := mysql.ClientDB.Query("SELECT DISTINCT tag FROM tags t ORDER BY tag")
	if err != nil {
		return nil, err
	}

	var topics []quotes.Tag
	for resultTopics.Next() {
		var topic quotes.Tag

		err = resultTopics.Scan(&topic.ID, &topic.Tag)
		if err != nil {
			return nil, err
		}

		topics = append(topics, topic)
	}

	return topics, nil
}
