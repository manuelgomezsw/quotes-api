package services

import (
	"quotes-api/internal/registry/domain"
	"quotes-api/internal/registry/repository"
)

func CreateQuoteService(quote domain.Quote) error {
	repository.CreateQuery()
	return nil
}

func UpdateQuoteService(quote domain.Quote) error {
	return nil
}

func DeleteQuoteService(quoteID string) error {
	return nil
}
