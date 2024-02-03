package services

import (
	"quotes-api/internal/domain"
	"quotes-api/internal/domain/registry/repository"
)

func CreateQuoteService(quote *domain.Quote) error {
	if err := repository.CreateQuote(quote); err != nil {
		return err
	}

	return nil
}

func UpdateQuoteService(quoteID int64, currentQuote *domain.Quote) error {
	if err := repository.UpdateQuote(quoteID, currentQuote); err != nil {
		return err
	}

	return nil
}

func DeleteQuoteService(quoteID int64) error {
	if err := repository.DeleteQuote(quoteID); err != nil {
		return err
	}

	return nil
}
