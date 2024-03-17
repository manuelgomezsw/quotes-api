package repository

import (
	"quotes-api/internal/domain/quotes"
	"quotes-api/internal/util/mysql"
)

func GetQuoteByID(quoteID int64) (quotes.Quote, error) {
	resultQuote, err := mysql.ClientDB.Query(
		"SELECT quote_id, author, work, phrase, date_created FROM `quotes`.`quotes` WHERE quote_id = ?", quoteID)
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
		"SELECT quote_id, author, work, phrase, date_created FROM `quotes`.`quotes` WHERE phrase LIKE ?", "%"+keyword+"%")
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

func GetTopics() ([]quotes.Topic, error) {
	resultTopics, err := mysql.ClientDB.Query(
		"SELECT DISTINCT CASE WHEN author = '' THEN 'Anónimo' ELSE author END 'value', 'author' AS 'type' FROM quotes.quotes UNION ALL SELECT DISTINCT `work` 'value', 'work' AS 'type' FROM quotes.quotes WHERE `work` != ''")
	if err != nil {
		return nil, err
	}

	var topics []quotes.Topic
	for resultTopics.Next() {
		var topic quotes.Topic

		err = resultTopics.Scan(&topic.Value, &topic.Type)
		if err != nil {
			return nil, err
		}

		topics = append(topics, topic)
	}

	return topics, nil
}
