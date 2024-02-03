package repository

import (
	"quotes-api/internal/domain"
	"quotes-api/internal/util/mysql"
)

func GetDailyQuote() (domain.Quote, error) {
	resultQuote, err := mysql.ClientDB.Query(
		"SELECT author, work, phrase, date_created FROM `quotes`.`quotes` ORDER BY RAND() LIMIT 1;")
	if err != nil {
		return domain.Quote{}, err
	}

	var quote domain.Quote
	for resultQuote.Next() {
		err = resultQuote.Scan(&quote.Author, &quote.Work, &quote.Phrase, &quote.DateCreated)
		if err != nil {
			return domain.Quote{}, err
		}
	}

	return quote, nil
}
