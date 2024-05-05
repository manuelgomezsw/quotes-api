package services

import (
	"quotes-api/internal/domain/words"
	"quotes-api/internal/domain/words/search/repository"
)

func GetByID(quoteID int64) (words.Word, error) {
	quote, err := repository.GetByID(quoteID)
	if err != nil {
		return words.Word{}, err
	}

	return quote, nil
}

func GetByKeyword(keyword string) ([]words.Word, error) {
	quote, err := repository.GetByKeyword(keyword)
	if err != nil {
		return nil, err
	}

	return quote, nil
}
