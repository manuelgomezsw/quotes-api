package repository

import (
	"fmt"
	"os"
	"quotes-api/internal/domain/quotes"
	"quotes-api/internal/util/mysql"
)

const (
	basePathSqlQueries = "sql/quotes"

	fileSqlCreateQuote        = "CreateQuote.sql"
	fileSqlUpdateQuote        = "UpdateQuote.sql"
	fileSqlDeleteQuote        = "DeleteQuote.sql"
	fileSqlGetQuoteByID       = "GetQuoteByID.sql"
	fileSqlGetQuotesByKeyword = "GetQuotesByKeyword.sql"
	fileSqlGetQuotesByAuthor  = "GetQuotesByAuthor.sql"
	fileSqlGetQuotesByWork    = "GetQuotesByWork.sql"
	fileSqlGetMinMaxQuotes    = "GetMinMaxQuotes.sql"
)

func CreateQuote(newQuote *quotes.Quote) error {
	query, err := os.ReadFile(fmt.Sprintf("%s/%s", basePathSqlQueries, fileSqlCreateQuote))
	if err != nil {
		return err
	}

	newRecord, err := mysql.ClientDB.Exec(
		string(query),
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

func UpdateQuote(quoteID int64, currentQuote *quotes.Quote) error {
	query, err := os.ReadFile(fmt.Sprintf("%s/%s", basePathSqlQueries, fileSqlUpdateQuote))
	if err != nil {
		return err
	}

	_, err = mysql.ClientDB.Exec(
		string(query),
		currentQuote.Author,
		currentQuote.Work,
		currentQuote.Phrase,
		quoteID,
	)
	if err != nil {
		return err
	}

	currentQuote.QuoteID = quoteID

	return nil
}

func DeleteQuote(quoteID int64) error {
	query, err := os.ReadFile(fmt.Sprintf("%s/%s", basePathSqlQueries, fileSqlDeleteQuote))
	if err != nil {
		return err
	}

	_, err = mysql.ClientDB.Exec(
		string(query),
		quoteID,
	)
	if err != nil {
		return err
	}

	return nil
}

func GetQuoteByID(quoteID int64) (quotes.Quote, error) {
	query, err := os.ReadFile(fmt.Sprintf("%s/%s", basePathSqlQueries, fileSqlGetQuoteByID))
	if err != nil {
		return quotes.Quote{}, err
	}

	resultQuote, err := mysql.ClientDB.Query(string(query), quoteID)
	if err != nil {
		return quotes.Quote{}, err
	}

	var quote quotes.Quote
	for resultQuote.Next() {
		err = resultQuote.Scan(&quote.QuoteID, &quote.Author, &quote.Work, &quote.Phrase, &quote.Keywords, &quote.DateCreated)
		if err != nil {
			return quotes.Quote{}, err
		}
	}

	return quote, nil
}

func GetQuotesByKeyword(keyword string) ([]quotes.Quote, error) {
	query, err := os.ReadFile(fmt.Sprintf("%s/%s", basePathSqlQueries, fileSqlGetQuotesByKeyword))
	if err != nil {
		return nil, err
	}

	resultQuote, err := mysql.ClientDB.Query(string(query), "%"+keyword+"%")
	if err != nil {
		return nil, err
	}

	var quotesSearched []quotes.Quote
	for resultQuote.Next() {
		var quote quotes.Quote

		err = resultQuote.Scan(&quote.QuoteID, &quote.Author, &quote.Work, &quote.Phrase, &quote.Keywords, &quote.DateCreated)
		if err != nil {
			return nil, err
		}

		quotesSearched = append(quotesSearched, quote)
	}

	return quotesSearched, nil
}

func GetQuotesByAuthor(author string) ([]quotes.Quote, error) {
	query, err := os.ReadFile(fmt.Sprintf("%s/%s", basePathSqlQueries, fileSqlGetQuotesByAuthor))
	if err != nil {
		return nil, err
	}

	resultQuote, err := mysql.ClientDB.Query(string(query), "%"+author+"%")
	if err != nil {
		return nil, err
	}

	var quotesSearched []quotes.Quote
	for resultQuote.Next() {
		var quote quotes.Quote

		err = resultQuote.Scan(&quote.QuoteID, &quote.Author, &quote.Work, &quote.Phrase, &quote.Keywords, &quote.DateCreated)
		if err != nil {
			return nil, err
		}

		quotesSearched = append(quotesSearched, quote)
	}

	return quotesSearched, nil
}

func GetQuotesByWork(work string) ([]quotes.Quote, error) {
	query, err := os.ReadFile(fmt.Sprintf("%s/%s", basePathSqlQueries, fileSqlGetQuotesByWork))
	if err != nil {
		return nil, err
	}

	resultQuote, err := mysql.ClientDB.Query(string(query), "%"+work+"%")
	if err != nil {
		return nil, err
	}

	var quotesSearched []quotes.Quote
	for resultQuote.Next() {
		var quote quotes.Quote

		err = resultQuote.Scan(&quote.QuoteID, &quote.Author, &quote.Work, &quote.Phrase, &quote.Keywords, &quote.DateCreated)
		if err != nil {
			return nil, err
		}

		quotesSearched = append(quotesSearched, quote)
	}

	return quotesSearched, nil
}

func GetMinMaxQuotes() (int64, int64, error) {
	query, err := os.ReadFile(fmt.Sprintf("%s/%s", basePathSqlQueries, fileSqlGetMinMaxQuotes))
	if err != nil {
		return 0, 0, err
	}

	resultQuotes, err := mysql.ClientDB.Query(string(query))
	if err != nil {
		return 0, 0, err
	}

	var minQuoteID, maxQuoteID int64
	for resultQuotes.Next() {
		err = resultQuotes.Scan(&minQuoteID, &maxQuoteID)
		if err != nil {
			return 0, 0, err
		}
	}

	return minQuoteID, maxQuoteID, nil
}
