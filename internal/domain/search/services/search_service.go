package services

import (
	"quotes-api/internal/domain"
	"quotes-api/internal/domain/search/repository"
)

func GetQuoteByID(quoteID int64) (domain.Quote, error) {
	quote, err := repository.GetQuoteByID(quoteID)
	if err != nil {
		return domain.Quote{}, err
	}

	return quote, nil
}

func GetQuotesByKeyword(keyword string) ([]domain.Quote, error) {
	quote, err := repository.GetQuotesByKeyword(keyword)
	if err != nil {
		return nil, err
	}

	return quote, nil
}

func GetQuotesByAuthor(author string) ([]domain.Quote, error) {
	quote, err := repository.GetQuotesByAuthor(author)
	if err != nil {
		return nil, err
	}

	return quote, nil
}

func GetQuotesByWork(work string) ([]domain.Quote, error) {
	quote, err := repository.GetQuotesByWork(work)
	if err != nil {
		return nil, err
	}

	return quote, nil
}
