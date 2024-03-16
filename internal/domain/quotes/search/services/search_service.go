package services

import (
	"quotes-api/internal/domain/quotes"
	"quotes-api/internal/domain/quotes/search/repository"
)

func GetQuoteByID(quoteID int64) (quotes.Quote, error) {
	quote, err := repository.GetQuoteByID(quoteID)
	if err != nil {
		return quotes.Quote{}, err
	}

	return quote, nil
}

func GetQuotesByKeyword(keyword string) ([]quotes.Quote, error) {
	quote, err := repository.GetQuotesByKeyword(keyword)
	if err != nil {
		return nil, err
	}

	return quote, nil
}

func GetQuotesByAuthor(author string) ([]quotes.Quote, error) {
	quote, err := repository.GetQuotesByAuthor(author)
	if err != nil {
		return nil, err
	}

	return quote, nil
}

func GetQuotesByWork(work string) ([]quotes.Quote, error) {
	quote, err := repository.GetQuotesByWork(work)
	if err != nil {
		return nil, err
	}

	return quote, nil
}
