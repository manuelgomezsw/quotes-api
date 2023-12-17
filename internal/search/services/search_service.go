package services

import (
	"quotes-api/internal/search/domain"
	"quotes-api/internal/search/repository"
)

func GetQuoteByID(quoteID int64) (domain.CompleteQuote, error) {
	quote, err := repository.GetQuoteByID(quoteID)
	if err != nil {
		return domain.CompleteQuote{}, err
	}

	return quote, nil
}

func GetQuotesByKeyword(keyword string) ([]domain.CompleteQuote, error) {
	quote, err := repository.GetQuotesByKeyword(keyword)
	if err != nil {
		return nil, err
	}

	return quote, nil
}

func GetQuotesByAuthor(author string) ([]domain.CompleteQuote, error) {
	quote, err := repository.GetQuotesByAuthor(author)
	if err != nil {
		return nil, err
	}

	return quote, nil
}

func GetQuotesByWork(work string) ([]domain.CompleteQuote, error) {
	quote, err := repository.GetQuotesByWork(work)
	if err != nil {
		return nil, err
	}

	return quote, nil
}
