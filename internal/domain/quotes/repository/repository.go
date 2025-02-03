package repository

import (
	"quotes-api/internal/domain/quotes"
	"quotes-api/internal/util/mysql"
	"strings"
)

func CreateQuote(newQuote *quotes.Quote) error {
	newRecord, err := mysql.ClientDB.Exec(
		"INSERT INTO quotes (author, work, phrase) VALUES (?, ?, ?)",
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

func GetTopics() ([]string, error) {
	resultTopics, err := mysql.ClientDB.Query("SELECT DISTINCT tag FROM tags t ORDER BY tag")
	if err != nil {
		return nil, err
	}

	var topics []string
	for resultTopics.Next() {
		var topic string

		err = resultTopics.Scan(&topic)
		if err != nil {
			return nil, err
		}

		topics = append(topics, topic)
	}

	return topics, nil
}

func GetDailyQuote() (quotes.Quote, error) {
	resultQuote, err := mysql.ClientDB.Query(
		"SELECT author, work, phrase, date_created FROM `quotes`.`quotes` ORDER BY RAND() LIMIT 1;")
	if err != nil {
		return quotes.Quote{}, err
	}

	var quote quotes.Quote
	for resultQuote.Next() {
		err = resultQuote.Scan(&quote.Author, &quote.Work, &quote.Phrase, &quote.DateCreated)
		if err != nil {
			return quotes.Quote{}, err
		}
	}

	return quote, nil
}

func createTags(quoteID int64, tags string) error {
	bulkInsert := "INSERT INTO tags (quote_id, tag) VALUES "
	var values []string
	var args []interface{}

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
